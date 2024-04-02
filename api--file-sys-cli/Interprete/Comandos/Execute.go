package comandos

import (
	"app/Parser"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

type Execute struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

//p *Parametros ExprCommand

func (e *Execute) Ejecutar(ctx *Contexto) interface{} {
	//Veficiar que no haya Errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(e.Parametros) == 0 {
		ctx.AgregarError("Error: Falta el parametro -path", e.Linea, e.Columna)
		return nil
	}
	//Obtener el path del script
	path_, existe := e.Parametros["-path"]
	//Verificar si existe el parametro -path
	if !existe {
		ctx.AgregarError("Error: Falta el parametro -path", e.Linea, e.Columna)
		return nil
	}
	//Leer contendio del script
	bytes, err := os.ReadFile(path_)
	if err != nil {
		ctx.AgregarError("Error al leer el Script: "+err.Error(), e.Linea, e.Columna)
		return nil
	}
	//Convertir a string
	script := string(bytes)
	//Crear un nuevo interprete
	lexer := Parser.NewFileSysCLILexer(antlr.NewInputStream(script))
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := Parser.NewFileSysCLIParser(tokenStream)
	raiz := parser.Commands()
	visitor := &FileSysCLIVisitor_{}
	interpreter := visitor.Visit(raiz).(ExprCommand)
	//Interpretar el script
	interpreter.Ejecutar(ctx)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	fmt.Println("---------Comando EXECUTE----------")
	fmt.Println("Comando ejecutado correctamente")
	return nil
}
