package main

import (
	comandos "app/Interprete/Comandos"
	"app/Parser"
	"bufio"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	//Verificar que existan los directorios donde se guardan los discos (MIA/P1) y si no existe crearlo
	_, err := os.Stat("./MIA/P1")
	if os.IsNotExist(err) {
		os.MkdirAll("./MIA/P1", os.ModePerm)
	}
	//
	for {
		fmt.Println("-------------------------------------------------")
		fmt.Println("Ingrese Comando")
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		lexer := Parser.NewFileSysCLILexer(antlr.NewInputStream(command))
		tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		parser := Parser.NewFileSysCLIParser(tokenStream)
		raiz := parser.Commands()
		visitor := &comandos.FileSysCLIVisitor_{}
		interpreter := visitor.Visit(raiz).(comandos.ExprCommand)
		ctx := comandos.NewContexto()
		interpreter.Ejecutar(ctx)
		fmt.Println("-------------------------------------------------")
		ctx.ImprimirErrores()
	}
}
