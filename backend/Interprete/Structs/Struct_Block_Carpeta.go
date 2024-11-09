package structs

import (
	"bytes"
	"strconv"
)

// Tamaño total 64 bytes
type B_content struct {
	B_name  [12]byte // Nombre del archivo o carpeta
	B_inodo int32    // Apuntador al inodo del archivo o carpeta
}

type Block_Carpeta struct {
	B_content [4]B_content //4 Apuntadores a Carpetas o Archivos
}

// Generar Graphviz
func (b *Block_Carpeta) GetGraph(numBloque int, primerBloque bool, soloBloque bool) string {
	filas := ``
	for pos, content := range b.B_content {
		name := ""
		if content.B_inodo != -1 {
			name = string(bytes.Replace(content.B_name[:], []byte{0}, []byte{}, -1))
		}
		filas += `
		<tr>
			<td>` + name + `</td><td port="f` + strconv.Itoa(pos) + `">` + strconv.Itoa(int(content.B_inodo)) + `</td>
		</tr>
		`
	}
	graphviz := `
	Bloque` + strconv.Itoa(numBloque) + ` [label=<<table border="1" cellspacing="0">
	<tr>
		<td colspan="2" bgcolor="#FF0000"><font color="white"><b>Bloque ` + strconv.Itoa(numBloque) + `</b></font></td>
	</tr>` + filas + `
	</table>> shape=box style=invisible ]
	`
	if soloBloque {
		return graphviz
	}
	//Conexiones
	for pos, content := range b.B_content {
		if primerBloque && pos < 2 {
			continue
		}
		if content.B_inodo != -1 {
			graphviz += `Bloque` + strconv.Itoa(numBloque) + `:f` + strconv.Itoa(pos) + " -> Inodo" + strconv.Itoa(int(content.B_inodo)) + "\n"
		}
	}
	return graphviz
}

// Nuevo Nodo Carpeta
func NewBlock_Carpeta() *Block_Carpeta {
	block := Block_Carpeta{}
	for i := 0; i < len(block.B_content); i++ {
		block.B_content[i] = NewB_content()
	}
	return &block
}

// Nuevo Contenido de Bloque
func NewB_content() B_content {
	return B_content{[12]byte{}, -1}
}
