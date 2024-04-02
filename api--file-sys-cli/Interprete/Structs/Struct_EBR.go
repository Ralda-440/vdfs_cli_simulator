package structs

import (
	"bytes"
	"strconv"
)

type EBR struct {
	Part_mount byte     // 0 = Inactiva, 1 = Activada
	Part_fit   byte     // B, F, W
	Part_start int64    // Byte de inicio
	Part_s     int64    // Tamaño de la partición
	Part_next  int64    // Byte de inicio de la siguiente partición
	Part_name  [16]byte // Nombre de la partición
}

// Crear Reporte de EBR
func (ebr *EBR) ReporteEBR() string {
	if ebr.Part_s == -1 {
		return ""
	}
	graphviz := `
	<tr>
        <td colspan="2" bgcolor="#0000CC"><font color="white"><b>Partición lógica</b></font></td>
    </tr>
    <tr>
        <td>part_mount</td><td>` + string(ebr.Part_mount) + `</td>
    </tr>
    <tr>
        <td>part_fit</td><td>` + string(ebr.Part_fit) + `</td>
	</tr>
	<tr>
        <td>part_start</td><td>` + strconv.Itoa(int(ebr.Part_start)) + `</td>
    </tr>
	<tr>
        <td>part_s</td><td>` + strconv.Itoa(int(ebr.Part_s)) + `</td>
    </tr>
	<tr>
        <td>part_next</td><td>` + strconv.Itoa(int(ebr.Part_next)) + `</td>
    </tr>
	<tr>
        <td>part_name</td><td>` + string(bytes.Replace(ebr.Part_name[:], []byte{0}, []byte{}, -1)) + `</td>
    </tr>
	`
	return graphviz
}

// Reporte Disk
func (ebr *EBR) ReporteDisk(sizeDisk int64, driveletter string) (string, error) {
	graphviz := ``
	usado := float64(ebr.Part_s) / float64(sizeDisk) * 100
	graphviz = `
	<td>EBR</td>
	<td>Logica<br></br>` + strconv.FormatFloat(usado, 'f', 2, 64) + ` % del Disco</td>
	`
	return graphviz, nil
}

// Crear e inicializar un EBR
func NewEBR() EBR {
	return EBR{Part_mount: 0, Part_fit: 0, Part_start: 0, Part_s: -1, Part_next: -1, Part_name: [16]byte{}}
}
