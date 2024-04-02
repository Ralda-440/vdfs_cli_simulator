package structs

import "bytes"

type Journaling struct {
	Comando        [10]byte  // Comando a ejecutar
	Path           [128]byte // Path del archivo o carpeta. 2 bloques * 64 = 128 bytes
	Contenido      [320]byte // Contenido del archivo. 5 bloques * 64 = 320 bytes
	Fecha          [19]byte  // Fecha y hora de ejecucion. 00/00/0000 00:00:00
	ComandoConsola [150]byte // Comando a ejecutar en consola
}

// Graficar Journaling
func (j *Journaling) GraficarJournaling() string {
	comando := bytes.Replace(j.Comando[:], []byte{0}, []byte{}, -1)
	path := bytes.Replace(j.Path[:], []byte{0}, []byte{}, -1)
	contenido := bytes.Replace(j.Contenido[:], []byte{0}, []byte{}, -1)
	if len(contenido) == 0 {
		contenido = []byte("N/A")
	}
	fecha := bytes.Replace(j.Fecha[:], []byte{0}, []byte{}, -1)
	return `
	<tr>
        <td>` + string(comando) + `</td>
        <td>` + string(path) + `</td>
        <td>` + string(contenido) + `</td>
        <td>` + string(fecha) + `</td>
    </tr>
	`
}
