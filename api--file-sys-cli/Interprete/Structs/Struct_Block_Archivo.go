package structs

import (
	"bytes"
	"strconv"
)

// Tamaño total 64 bytes
type Block_Archivo struct {
	B_content [64]byte //64 bytes de contenido
}

// Nuevo Bloque Archivo
func NewBlock_Archivo() *Block_Archivo {
	block := Block_Archivo{[64]byte{}}
	return &block
}

// Generar Graphviz
func (b *Block_Archivo) GetGraph(numBloque int) string {
	return `
	Bloque` + strconv.Itoa(numBloque) + ` [shape=box label="Bloque ` + strconv.Itoa(numBloque) + `
----------------------
` + string(bytes.Replace(b.B_content[:], []byte{0}, []byte{}, -1)) + `"]
`
}
