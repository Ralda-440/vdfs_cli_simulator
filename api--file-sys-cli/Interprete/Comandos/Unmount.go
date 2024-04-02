package comandos

import (
	structs "app/Interprete/Structs"
	"fmt"
)

type Unmount struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

//umt *Unmount ExprCommand

func (umt *Unmount) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(umt.Parametros) == 0 {
		ctx.AgregarError("Error: Falta el parametro -id", umt.Linea, umt.Columna)
		return nil
	}
	//Verificar Parametros Obligatorios -id
	if _, existe := umt.Parametros["-id"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -id", umt.Linea, umt.Columna)
		return nil
	}
	//Obtener parametros
	id := umt.Parametros["-id"]
	//Desmontar particion
	partMount := ParticionesMontadas.UnmountParticion(id)
	if partMount == nil {
		ctx.AgregarError("Error : No existe la particion montada con el id "+id, umt.Linea, umt.Columna)
		return nil
	}
	//Actualizar atributo Part_status
	//Obtener mbr del Disco
	mbr, err := getMBRDisk(partMount.DiskName)
	if err != nil {
		ctx.AgregarError("Error : El disco de la particion montada ya no Existe"+err.Error(), umt.Linea, umt.Columna)
		return nil
	}
	//Buscar particion con el nombre
	existe, particion, _, ebr, _, _, _, err := mbr.BuscarParticion(partMount.DiskName, partMount.PartName)
	if err != nil {
		ctx.AgregarError("Error : No se puedo buscar la particion al Desmontar"+err.Error(), umt.Linea, umt.Columna)
		return nil
	}
	if !existe {
		ctx.AgregarError("Error : Ya no existe la particion con el nombre \""+partMount.PartName+"\" que se intenta Desmontar", umt.Linea, umt.Columna)
	}
	//Si es Primaria o Extendida
	if particion != nil {
		if particion.Part_type == 'E' {
			ctx.AgregarError("Error : No se puede desmontar una particion Extendida", umt.Linea, umt.Columna)
			return nil
		}
		//Actualizar atributo Part_status
		particion.Part_status = '0'
		//Actualizar MBR
		err = writeMBRInDisk(partMount.DiskName, mbr)
		if err != nil {
			ctx.AgregarError("Error : No se puedo actualizar el MBR del disco al intentar Desmontar la particion"+err.Error(), umt.Linea, umt.Columna)
			return nil
		}
	}
	//Si es Logica
	if ebr != nil {
		//Actualizar atributo Part_mount
		ebr.Part_mount = '0'
		//Actualizar EBR
		err := particion.WriteEBRsInDisk(partMount.DiskName, []structs.EBR{*ebr})
		if err != nil {
			ctx.AgregarError("Error : No se puedo actualizar el EBR del disco al intentar Desmontar la particion"+err.Error(), umt.Linea, umt.Columna)
			return nil
		}
	}
	//Mensaje de exito
	fmt.Println("----------Comando UNMOUNT----------")
	fmt.Println("Particion Desmontada con exito")
	return nil
}
