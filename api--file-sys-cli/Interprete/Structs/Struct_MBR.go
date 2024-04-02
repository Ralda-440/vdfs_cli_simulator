package structs

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

type MBR struct {
	Mbr_tamanio        int64        //Tamaño del disco
	Mbr_fecha_creacion [19]byte     //00/00/0000 00:00:00
	Mbr_disk_signature int64        //Numero random
	Mbr_disk_fit       byte         // B, F, W
	Particiones        [4]Particion //Particiones
}

// Agregar Particion a lugar vacio
func (mbr *MBR) AgregarParticion(particion_ Particion) bool {
	for pos, particion := range mbr.Particiones {
		if particion.Part_correlative == -1 {
			mbr.Particiones[pos] = particion_
			return true
		}
	}
	return false
}

// Ordenar Particiones por byte de inicio y las no usadas al final
func (mbr *MBR) OrdenarParticiones() {
	particiones := make([]Particion, 0)
	for _, particion := range mbr.Particiones {
		if particion.Part_correlative == -1 {
			continue
		}
		particiones = append(particiones, particion)
	}
	for i := 0; i < len(particiones); i++ {
		for j := 0; j < len(particiones)-1; j++ {
			if particiones[j].Part_start > particiones[j+1].Part_start {
				particiones[j], particiones[j+1] = particiones[j+1], particiones[j]
			}
		}
	}
	for i := 0; i < len(mbr.Particiones); i++ {
		if i < len(particiones) {
			mbr.Particiones[i] = particiones[i]
		} else {
			mbr.Particiones[i] = NewParticion()
		}
	}
}

// Buscar Particion Extendida
func (mbr *MBR) BuscarParticionExtendida() *Particion {
	for _, particion := range mbr.Particiones {
		if particion.Part_correlative == -1 {
			continue
		}
		if particion.Part_type == 'E' {
			return &particion
		}
	}
	return nil

}

// Buscar particion por nombre. Retorna si existe, Particion P o E (Si lo es), Particion Logica que coincide con el nombre (Si lo es),
// Particion Logica Anterior a la Particion Logica que coincide con el nombre, la cantidad de particiones primarias y la cantidad de particiones extendidas y error
func (mbr *MBR) BuscarParticion(driveletter string, name string) (bool, *Particion, int, *EBR, *EBR, int, int, error) {
	cantPartPrimarias := 0
	cantPartExtendidas := 0
	for pos, particion := range mbr.Particiones {
		if particion.Part_correlative == -1 {
			continue
		}
		if particion.Part_type == 'P' {
			cantPartPrimarias++
		}
		if particion.Part_type == 'E' {
			cantPartExtendidas++
		}
		if string(bytes.Replace(particion.Part_name[:], []byte{0}, []byte{}, -1)) == name {
			return true, &mbr.Particiones[pos], pos, nil, nil, cantPartPrimarias, cantPartExtendidas, nil
		}
		//Buscar particion en extendida
		if particion.Part_type == 'E' {
			//Leer EBRs
			ebrs, err := particion.ReadEBRsInDisk(driveletter)
			if err != nil {
				return false, nil, -1, nil, nil, cantPartPrimarias, cantPartExtendidas, err
			}
			//Buscar particion logica
			for pos, ebr := range ebrs {
				if ebr.Part_s == -1 {
					continue
				}
				if string(bytes.Replace(ebr.Part_name[:], []byte{0}, []byte{}, -1)) == name {
					var antEBR *EBR = nil
					if pos > 0 {
						antEBR = &ebrs[pos-1]
					}
					return true, nil, -1, &ebr, antEBR, cantPartPrimarias, cantPartExtendidas, nil
				}
			}
		}
	}
	return false, nil, -1, nil, nil, cantPartPrimarias, cantPartExtendidas, nil
}

// Buscar Espacio para una particion
// Retorna bool (si hay espacio), int64 (byte de inicio)
func (mbr *MBR) BuscarEspacio(size int64) (bool, int64) {
	start_absolute := int64(binary.Size(MBR{}))
	//Nuevo Slice con solo las particiones ocupadas
	particiones := make([]Particion, 0)
	//Ordenar Partciones por byte de inicio (Ascendente)
	mbr.OrdenarParticiones()
	//Llenar Slice de particiones ocupadas
	for _, particion := range mbr.Particiones {
		if particion.Part_correlative == -1 {
			continue
		}
		particiones = append(particiones, particion)

	}
	//Buscar espacio
	hayEspacio := false
	auxStart_absolute := start_absolute
	auxEspacio := int64(-1)
	for _, particion := range particiones {
		if string(mbr.Mbr_disk_fit) == "F" {
			if auxStart_absolute+size <= particion.Part_start {
				start_absolute = auxStart_absolute
				hayEspacio = true
				//auxStart_absolute = particion.Part_start + particion.Part_s
				//break
				return hayEspacio, start_absolute
			}
		} else if string(mbr.Mbr_disk_fit) == "B" {
			if auxStart_absolute+size <= particion.Part_start {
				if auxEspacio == -1 || particion.Part_start-auxStart_absolute < auxEspacio {
					auxEspacio = particion.Part_start - auxStart_absolute
					start_absolute = auxStart_absolute
				}
				hayEspacio = true
			}
		} else if string(mbr.Mbr_disk_fit) == "W" {
			if auxStart_absolute+size <= particion.Part_start {
				if auxEspacio == -1 || particion.Part_start-auxStart_absolute > auxEspacio {
					auxEspacio = particion.Part_start - auxStart_absolute
					start_absolute = auxStart_absolute
				}
				hayEspacio = true
			}
		}
		auxStart_absolute = particion.Part_start + particion.Part_s
	}
	//Verificar si hay espacio al final
	if auxStart_absolute+size <= mbr.Mbr_tamanio {
		if string(mbr.Mbr_disk_fit) == "F" {
			start_absolute = auxStart_absolute
			hayEspacio = true
		} else if string(mbr.Mbr_disk_fit) == "B" {
			if auxEspacio == -1 || mbr.Mbr_tamanio-auxStart_absolute < auxEspacio {
				auxEspacio = mbr.Mbr_tamanio - auxStart_absolute
				start_absolute = auxStart_absolute
			}
			hayEspacio = true
		} else if string(mbr.Mbr_disk_fit) == "W" {
			if auxEspacio == -1 || mbr.Mbr_tamanio-auxStart_absolute > auxEspacio {
				auxEspacio = mbr.Mbr_tamanio - auxStart_absolute
				start_absolute = auxStart_absolute
			}
			hayEspacio = true
		}
	}
	return hayEspacio, start_absolute
}

// Buscar posicion de particion vacia
func (mbr *MBR) NuevoCorrelativo() int64 {
	//Guardar Correlativos usados
	correlativos := make([]int64, 0)
	for _, particion := range mbr.Particiones {
		if particion.Part_correlative != -1 {
			correlativos = append(correlativos, particion.Part_correlative)
		}
	}
	//Buscar correlativo que no se haya usado
	allCorrelaitvos := []int64{1, 2, 3, 4}
	for _, allallCorrelaitvos := range allCorrelaitvos {
		used := false
		for _, correlativo := range correlativos {
			if allallCorrelaitvos == correlativo {
				used = true
				break
			}
		}
		if !used {
			return allallCorrelaitvos
		}
	}
	return int64(-1)
}

// Crear Reporte de MBR
func (mbr *MBR) ReporteMBR(driveletter string) (string, error) {
	particiones := ""
	for _, particion := range mbr.Particiones {
		if particion.Part_correlative == -1 {
			continue
		}
		part, err := particion.ReporteParticion(driveletter)
		if err != nil {
			return "", err
		}
		particiones += part
	}
	if particiones == "" {
		particiones = "No hay particiones"
	} else {
		particiones = `
		<table border="1" cellspacing="0">
			` + particiones + `
		</table>`
	}
	graphviz := `
	digraph G {
		table [label=< 
				<table border="1" cellspacing="0">
				<tr>
					<td colspan="2" bgcolor="#000000"><font color="white"><b>Reporte MBR</b></font></td>
				</tr>
				<tr>
					<td>mbr_tamaño</td><td>` + strconv.Itoa(int(mbr.Mbr_tamanio)) + `</td>
				</tr>
				<tr>
					<td>mbr_fecha_creacion</td><td>` + string(mbr.Mbr_fecha_creacion[:]) + `</td>
				</tr>
				<tr>
					<td>mbr_disk_signature</td><td>` + strconv.Itoa(int(mbr.Mbr_disk_signature)) + `</td>
				</tr>
				<tr>
					<td>mbr_disk_fit</td><td>` + string(mbr.Mbr_disk_fit) + `</td>
				</tr>
				<tr>
					<td colspan="2" align="center" textalign="center" cellpadding="15" >` + particiones + `</td>
				</tr>
				</table> 
				> shape=box style=invisible ] 
	}
	`
	return graphviz, nil
}

// Reporte Graphviz Disk
func (mbr *MBR) ReporteDisk(driveletter string) (string, error) {
	graphivz := ``
	columnas := 1
	//Slice de particiones ocupadas
	particiones := make([]Particion, 0)
	for _, particion := range mbr.Particiones {
		if particion.Part_correlative == -1 {
			continue
		}
		particiones = append(particiones, particion)
	}
	temp_start := int64(binary.Size(MBR{}))
	for _, particion := range particiones {
		libre := particion.Part_start - temp_start
		if libre > 0 {
			porcentaje := float64(libre) / float64(mbr.Mbr_tamanio) * 100
			graphivz += `
			<td>Libre<br></br>` + strconv.FormatFloat(porcentaje, 'f', 2, 64) + ` % del Disco</td>
			`
			columnas++
		}
		graphivzParticion, err := particion.ReporteDisk(mbr.Mbr_tamanio, driveletter)
		if err != nil {
			return "", err
		}
		graphivz += graphivzParticion
		temp_start = particion.Part_start + particion.Part_s
		columnas++
	}
	libre := mbr.Mbr_tamanio - temp_start
	if libre > 0 {
		porcentaje := float64(libre) / float64(mbr.Mbr_tamanio) * 100
		graphivz += `
		<td>Libre<br></br>` + strconv.FormatFloat(porcentaje, 'f', 2, 64) + ` % del Disco</td>
		`
		columnas++
	}
	graphivz = `
	digraph G {
		disk [label=< 
				<table border="1" cellspacing="0">
				<tr>
					<td colspan="` + strconv.Itoa(columnas) + `" bgcolor="#000000"><font color="white"><b>` + driveletter + `.dsk</b></font></td>
				</tr>
				<tr>
					<td>MBR</td>
				` + graphivz + `
				</tr>
				</table>
			> shape=box style=invisible ]
				}`
	return graphivz, nil
}

// Crear e inicializar un MBR
func NewMBR() MBR {
	return MBR{Mbr_tamanio: 0, Mbr_fecha_creacion: [19]byte{}, Mbr_disk_signature: 0, Mbr_disk_fit: 0, Particiones: [4]Particion{NewParticion(), NewParticion(), NewParticion(), NewParticion()}}
}
