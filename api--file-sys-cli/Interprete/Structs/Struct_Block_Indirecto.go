package structs

import "strconv"

// Tamaño total 64 bytes
type Block_Indirecto struct {
	B_pointers [16]int32 //16 Apuntadores a bloques de carpetas o archivos
}

// Generar Graphviz
func (b *Block_Indirecto) GetGraph(numBloque int, soloBloque bool) string {
	filas := ``
	for pos, content := range b.B_pointers {
		filas += `
		<tr>
			<td>` + strconv.Itoa(pos) + `</td><td port="f` + strconv.Itoa(pos) + `">` + strconv.Itoa(int(content)) + `</td>
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
	for pos, content := range b.B_pointers {
		if content != -1 {
			graphviz += `Bloque` + strconv.Itoa(numBloque) + `:f` + strconv.Itoa(pos) + " -> Bloque" + strconv.Itoa(int(content)) + "\n"
		}
	}
	return graphviz
}

// Nuevo Bloque Indirecto
func NewBlock_Indirecto() *Block_Indirecto {
	block := Block_Indirecto{[16]int32{}}
	//Cada apuntador a bloque se inicializa en -1
	for i := 0; i < 16; i++ {
		block.B_pointers[i] = -1
	}
	return &block
}
