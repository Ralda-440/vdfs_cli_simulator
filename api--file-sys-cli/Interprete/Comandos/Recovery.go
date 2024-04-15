package comandos

import (
	"app/Parser"
	"bytes"

	"github.com/antlr4-go/antlr/v4"
)

type Recovery struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// recovery *Recovery ExprCommand
func (recovery *Recovery) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar que haya parametros
	if len(recovery.Parametros) == 0 {
		ctx.AgregarError("El comando recovery no tiene parametros", recovery.Linea, recovery.Columna)
		return nil
	}
	//Verificar parametros obligatorios -id
	if _, existe := recovery.Parametros["-id"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -id", recovery.Linea, recovery.Columna)
		return nil
	}
	//Obtener id
	id := recovery.Parametros["-id"]
	//Buscar Particion Montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(id)
	if !existe {
		ctx.AgregarError("Error: No hay ninguna particion montada", recovery.Linea, recovery.Columna)
		return nil
	}
	//Obtener SuperBloque
	superBloque, err := getSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo obtener el super bloque de la particion", recovery.Linea, recovery.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Recuperar Journaling
	ops := superBloque.RecuperarJournaling(ctx, partMontada.DiskName, partMontada.Start_SuperBloque)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Limpiar Journaling
	superBloque.LimpiarJournaling(ctx, partMontada.DiskName, partMontada.Start_SuperBloque)
	//Recuperar todos los comandos
	comandos := ""
	for _, op := range ops {
		comandos += string(bytes.Replace(op.ComandoConsola[:], []byte{0}, []byte(""), -1)) + "\n"
	}
	//Crear Nuevo Analizador para ejecutar los comandos recuperados
	lexer := Parser.NewFileSysCLILexer(antlr.NewInputStream(comandos))
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := Parser.NewFileSysCLIParser(tokenStream)
	raiz := parser.Commands()
	visitor := &FileSysCLIVisitor_{}
	interpreter := visitor.Visit(raiz).(ExprCommand)
	//Ejecutar comandos
	interpreter.Ejecutar(ctx)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//fmt.Println("---------Comando RECOVERY----------")
	//fmt.Println("Comandos recuperados y ejecutados correctamente")
	ctx.AgregarOutput("------------Comando RECOVERY----------")
	ctx.AgregarOutput("La particion con id \"" + id + "\" se ha recuperado correctamente")
	return nil
}
