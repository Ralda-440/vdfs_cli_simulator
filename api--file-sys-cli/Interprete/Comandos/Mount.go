package comandos

import (
	structs "app/Interprete/Structs"
	"encoding/binary"
	"strconv"
)

type Mount struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

func (m *Mount) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(m.Parametros) == 0 {
		//Si no hay parametros, listar particiones montadas
		m.ListarParticionesMontadas(ctx)
		return nil
	}
	//Verificar Parametros Obligatorios -driveletter -name
	if _, existe := m.Parametros["-driveletter"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -driveletter", m.Linea, m.Columna)
		return nil
	}
	if _, existe := m.Parametros["-name"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -name", m.Linea, m.Columna)
		return nil
	}
	//Obtener parametros
	driveLetter := m.Parametros["-driveletter"]
	name := m.Parametros["-name"]
	//Obtener mbr del Disco
	mbr, err := GetMBRDisk(driveLetter)
	if err != nil {
		ctx.AgregarError("Error : No se puedo leer el MBR del disco"+err.Error(), m.Linea, m.Columna)
		return nil
	}
	//Buscar particion con el nombre
	existe, particion, _, ebr, _, _, _, err := mbr.BuscarParticion(driveLetter, name)
	if err != nil {
		ctx.AgregarError("Error : No se puedo buscar la particion"+err.Error(), m.Linea, m.Columna)
		return nil
	}
	if !existe {
		ctx.AgregarError("Error : No existe la particion con el nombre "+name, m.Linea, m.Columna)
	}
	//Montar particion
	//Si es Primaria o Extendida
	var id string
	if particion != nil {
		if particion.Part_type == 'E' {
			ctx.AgregarError("Error : No se puede montar una particion Extendida", m.Linea, m.Columna)
			return nil
		}
		//Actualizar atributo Part_status
		particion.Part_status = '1'
		id = particion.GenerarID(driveLetter)
		//Agregar particion montada
		ParticionesMontadas.AddParticionPrimaria(driveLetter, id, name, string(particion.Part_fit), particion.Part_start)
		//Escribir particion en el disco
		err := writeMBRInDisk(driveLetter, mbr)
		if err != nil {
			ctx.AgregarError("Error : No se puedo escribir el MBR del disco"+err.Error(), m.Linea, m.Columna)
			return nil
		}
	}
	//Si es Logica
	if ebr != nil {
		//Actualizar atributo Part_mount
		ebr.Part_mount = '1'
		//Escribir EBR en el disco
		err := particion.WriteEBRsInDisk(driveLetter, []structs.EBR{*ebr})
		if err != nil {
			ctx.AgregarError("Error : No se puedo escribir el EBR modificado montado del disco"+err.Error(), m.Linea, m.Columna)
			return nil
		}
		//Agregar particion montada
		id = ParticionesMontadas.AddParticionLogica(driveLetter, name, string(ebr.Part_fit), ebr.Part_start+int64(binary.Size(structs.EBR{})))
	}
	//fmt.Println("----------Comando MOUNT----------")
	//fmt.Println("Comando ejecutado con exito")
	ctx.AgregarOutput("----------Comando MOUNT----------")
	ctx.AgregarOutput("Particion: \"" + name + "\" con ID: \"" + id + "\" Ubicada en el Disco: \"" + driveLetter + "\" Montada con exito")
	return nil
}

// Listar particiones montadas
func (m *Mount) ListarParticionesMontadas(ctx *Contexto) {
	//fmt.Println("----------Particiones Montadas----------")
	ctx.AgregarOutput("----------Particiones Montadas----------")
	if len(ParticionesMontadas.Particiones) == 0 {
		//fmt.Println("No hay particiones montadas")
		ctx.AgregarOutput("No hay particiones montadas")
		return
	}
	for _, particion := range ParticionesMontadas.Particiones {
		//fmt.Println("Disco: " + particion.DiskName + " ID: " + particion.Id + " Nombre: " + particion.PartName)
		ctx.AgregarOutput("Disco: \"" + particion.DiskName + "\" ID: \"" + particion.Id + " \"Nombre: \"" + particion.PartName + "\"")
	}
}

type PartMount struct {
	DiskName          string
	Id                string
	PartName          string
	Fit               string
	Start_SuperBloque int64
}

type PartMounts struct {
	Particiones []PartMount
}

// Buscar particion montada
func (pm *PartMounts) BuscarParticionMontada(id string) (bool, *PartMount) {
	for _, particion := range pm.Particiones {
		if particion.Id == id {
			return true, &particion
		}
	}
	return false, nil
}

// Agregar particion montada Primaria
func (pm *PartMounts) AddParticionPrimaria(diskName string, id string, partName string, fit string, startSuperBloque int64) {
	pm.Particiones = append(pm.Particiones, PartMount{DiskName: diskName, Id: id, PartName: partName, Fit: fit, Start_SuperBloque: startSuperBloque})
}

// Agregar particion montada Logica
func (pm *PartMounts) AddParticionLogica(diskName string, partName string, fit string, startSuperBloque int64) string {
	//Crear ID
	id := pm.GenerarID(diskName)
	//Agregar particion
	pm.Particiones = append(pm.Particiones, PartMount{DiskName: diskName, Id: id, PartName: partName, Fit: fit, Start_SuperBloque: startSuperBloque})
	return id
}

// Desmontar particion
func (pm *PartMounts) UnmountParticion(id string) *PartMount {
	for pos, particion := range pm.Particiones {
		if particion.Id == id {
			pm.Particiones = append(pm.Particiones[:pos], pm.Particiones[pos+1:]...)
			return &particion
		}
	}
	return nil
}

// Generar ID para particion logica
func (pm *PartMounts) GenerarID(diskName string) string {
	CountPartLogicas++
	return diskName + strconv.Itoa(CountPartLogicas) + "16"
}

var CountPartLogicas int = 4

var ParticionesMontadas PartMounts
