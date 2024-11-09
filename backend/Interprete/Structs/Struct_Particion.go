package structs

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"strconv"
)

type Particion struct {
	Part_status      byte     // 0 = Inactiva, 1 = Activa
	Part_type        byte     // P, E
	Part_fit         byte     // B, F, W
	Part_start       int64    // Byte de inicio
	Part_s           int64    // Tamaño de la partición
	Part_name        [16]byte // Nombre de la partición
	Part_correlative int64    // Correlativo de la partición
	Part_id          [4]byte  // Identificador de la partición
}

// Generar ID de particion
func (part *Particion) GenerarID(driveletter string) string {
	part.Part_id = [4]byte{driveletter[0], strconv.Itoa(int(part.Part_correlative))[0], byte('1'), byte('6')}
	return string(part.Part_id[:])
}

// Buscar espacio para una particion logica
// Retorna bool hay espacio, start del nuevo EBR, next
func (part *Particion) BuscarEspacioParticionLogica(size int64, ebrs *[]EBR) (bool, int64, int64) {
	hay_espacio := false
	temp_hay_espacio := false
	start_absolute := int64(-1)
	temp_start_absolute := int64(-1)
	next := int64(-1)
	temp_next := int64(-1)
	espacio := int64(-1)
	temp_espacio := int64(-1)
	var antEBR *EBR
	var temp_antEBR *EBR
	limitSize := int64(-1)
	ebrs_ := *ebrs
	//Iterar sobre los EBRs para buscar espacio
	for indice, ebr := range ebrs_ {
		//Obtener el limite de la particion logica
		if ebr.Part_next == -1 {
			limitSize = part.Part_s + part.Part_start
		} else {
			limitSize = ebr.Part_next
		}
		//Calcular el espacio disponible
		if ebr.Part_s == -1 {
			temp_start_absolute = ebr.Part_start
			temp_espacio = limitSize - temp_start_absolute
		} else {
			temp_start_absolute = ebr.Part_start + ebr.Part_s
			temp_espacio = limitSize - temp_start_absolute
			temp_antEBR = &ebrs_[indice]
		}
		temp_next = ebr.Part_next
		//Si hay espacio suficiente
		if size > temp_espacio {
			continue
		}
		temp_hay_espacio = true
		//Verificar fit
		if (part.Part_fit == 'F') || (part.Part_fit == 'B' && (espacio == -1 || temp_espacio < espacio)) || (part.Part_fit == 'W' && (espacio == -1 || temp_espacio > espacio)) {
			hay_espacio = temp_hay_espacio
			start_absolute = temp_start_absolute
			next = temp_next
			espacio = temp_espacio
			antEBR = temp_antEBR
			if part.Part_fit == 'F' {
				break
			}
		}
	}
	if antEBR != nil && hay_espacio {
		antEBR.Part_next = start_absolute
	}
	return hay_espacio, start_absolute, next
}

// Escribir slice de EBRs en posicion absoluta del disco
func (part *Particion) WriteEBRsInDisk(name string, ebrs []EBR) error {
	//Abrir el archivo del disco
	file, err := os.OpenFile("./MIA/P1/"+name+".dsk", os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	//Escribir el EBR
	for _, ebr := range ebrs {
		//Mover el puntero al EBR actual
		file.Seek(ebr.Part_start, io.SeekStart)
		err = binary.Write(file, binary.BigEndian, &ebr)
		if err != nil {
			return err
		}
	}
	return nil
}

// Leer slice de EBRs en posicion absoluta del disco
func (part *Particion) ReadEBRsInDisk(name string) ([]EBR, error) {
	//Abrir el archivo del disco
	file, err := os.Open("./MIA/P1/" + name + ".dsk")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//Mover el puntero al inicio del archivo
	file.Seek(part.Part_start, io.SeekStart)
	//Leer el EBR
	ebrs := make([]EBR, 0)
	for {
		ebr := NewEBR()
		err = binary.Read(file, binary.BigEndian, &ebr)
		if err != nil {
			return nil, err
		}
		ebrs = append(ebrs, ebr)
		//Si no hay siguiente EBR, terminar
		if ebr.Part_next == -1 {
			break
		}
		//Mover el puntero al siguiente EBR
		file.Seek(ebr.Part_next, io.SeekStart)
	}
	return ebrs, nil
}

// Crear Reporte de Partición
func (part *Particion) ReporteParticion(driveletter string) (string, error) {
	if part.Part_correlative == -1 {
		return "", nil
	}
	//Reporte Partciones Logicas solo si es extendida
	reporteGraphvizLogicas := ""
	if part.Part_type == 'E' {
		reporteGraphvizLogicas = `
		<tr>
        <td colspan="2" align="center" cellpadding="15" >
		`
		rep, err := part.GraphvizParticionLogicas(driveletter)
		if err != nil {
			return "", err
		}
		if rep == "" {
			reporteGraphvizLogicas += "No hay particiones lógicas"
		} else {
			reporteGraphvizLogicas += `<table border="1" cellspacing="0" >` + rep + `</table>
			`
		}
		reporteGraphvizLogicas += `</td>
    	</tr>
		`
	}
	//Graphviz Particiones Logicas (Si es que hay y si es que es extendida)
	graphviz := `
	<tr>
        <td colspan="2" bgcolor="#800080"><font color="white"><b>Partición</b></font></td>
    </tr>
    <tr>
        <td>part_status</td><td>` + string(part.Part_status) + `</td>
    </tr>
	<tr>
		<td>part_type</td><td>` + string(part.Part_type) + `</td>
	</tr>
	<tr>
		<td>part_fit</td><td>` + string(part.Part_fit) + `</td>
	</tr>
	<tr>
		<td>part_start</td><td>` + strconv.Itoa(int(part.Part_start)) + `</td>
	</tr>
	<tr>
		<td>part_s</td><td>` + strconv.Itoa(int(part.Part_s)) + `</td>
	</tr>
	<tr>
		<td>part_name</td><td>` + string(bytes.Replace(part.Part_name[:], []byte{0}, []byte{}, -1)) + `</td>
	</tr>
	<tr>
		<td>part_correlative</td><td>` + strconv.Itoa(int(part.Part_correlative)) + `</td>
	</tr>
	<tr>
		<td>part_id</td><td>` + string(part.Part_id[:]) + `</td>
	</tr>
	` + reporteGraphvizLogicas
	return graphviz, nil

}

// Graphviz particiones
func (part *Particion) GraphvizParticionLogicas(driveletter string) (string, error) {
	graphviz := ""
	ebrs, err := part.ReadEBRsInDisk(driveletter)
	if err != nil {
		return "", err
	}
	for _, ebr := range ebrs {
		graphviz += ebr.ReporteEBR()
	}
	return graphviz, nil
}

// Rep Disk
func (part *Particion) ReporteDisk(sizeDisk int64, driveletter string) (string, error) {
	graphviz := ``
	if part.Part_type == 'P' {
		usado := float64(part.Part_s) / float64(sizeDisk) * 100
		graphviz = `
		<td>Primaria<br></br>` + strconv.FormatFloat(usado, 'f', 2, 64) + ` % del Disco</td>
		`
		return graphviz, nil
	}
	columnas := 0
	//Obtener las particiones logicas
	ebrs, err := part.ReadEBRsInDisk(driveletter)
	if err != nil {
		return "", err
	}
	temp_start := part.Part_start
	for _, ebr := range ebrs {
		if ebr.Part_s == -1 {
			continue
		}
		libre := ebr.Part_start - temp_start
		if libre > 0 {
			porcentaje := float64(libre) / float64(sizeDisk) * 100
			graphviz += `
				<td>Libre<br></br>` + strconv.FormatFloat(porcentaje, 'f', 2, 64) + ` % del Disco</td>
				`
			columnas++
		}
		graphvizParticion, err := ebr.ReporteDisk(sizeDisk, driveletter)
		if err != nil {
			return "", err
		}
		graphviz += graphvizParticion
		columnas += 2
		temp_start = ebr.Part_start + ebr.Part_s
	}
	//Espacio libre al final de la particion extendida
	libre := part.Part_start + part.Part_s - temp_start
	if libre > 0 {
		porcentaje := float64(libre) / float64(sizeDisk) * 100
		graphviz += `
			<td>Libre<br></br>` + strconv.FormatFloat(porcentaje, 'f', 2, 64) + ` % del Disco</td>
			`
		columnas++
	}
	graphviz = `
		<td align="center" textalign="center" cellpadding="2" >
			<table border="1" cellspacing="0">
				<tr>
				<td colspan="` + strconv.Itoa(columnas) + `" bgcolor="#800080"><font color="white"><b>Extendida</b></font></td>
				</tr>
				<tr>
					` + graphviz + `
				</tr>
			</table>
		</td>`
	return graphviz, nil
}

// Crear e inicializar una partición
func NewParticion() Particion {
	return Particion{Part_status: byte('0'), Part_type: byte('0'), Part_fit: byte('0'), Part_start: 0, Part_s: 0, Part_name: [16]byte{}, Part_correlative: -1, Part_id: [4]byte{0, 0, 0, 0}}
}
