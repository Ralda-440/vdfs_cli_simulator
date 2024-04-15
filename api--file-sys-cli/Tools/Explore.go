package tools

import (
	comandos "app/Interprete/Comandos"
	"os"
)

// Discos Creados
func DiscosCreados(ctx *comandos.Contexto) map[string]string {
	files, err := os.ReadDir("./MIA/P1")
	if err != nil {
		ctx.AgregarError("Error: No se pudo leer la carpeta de discos", 0, 0)
		return nil
	}
	discos := make(map[string]string)
	for _, file := range files {
		discos[file.Name()] = "disk"
	}
	return discos
}
