package tools

import (
	comandos "app/Interprete/Comandos"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type ItemReport struct {
	Nombre  string `json:"nombre"`
	Tipo    string `json:"tipo"`
	Content string `json:"content"`
}

type InputCLI struct {
	Input string `json:"input"`
}

type PathExplorer struct {
	Path string `json:"path"`
}

type LoginActivo struct {
	Activo   bool   `json:"activo"`
	PartName string `json:"partName"`
	DiskName string `json:"diskName"`
}

// Activar Login
func (la *LoginActivo) ActivarLogin(partName, diskName string) {
	la.Activo = true
	la.PartName = partName
	la.DiskName = diskName
}

// Nuevo Login Activo
func NewLoginActivo() *LoginActivo {
	return &LoginActivo{Activo: false, PartName: ""}
}

type ErrorListener_ struct {
	antlr.ErrorListener
	Ctx *comandos.Contexto
}

func (el *ErrorListener_) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if offendingSymbol != nil {
		token := offendingSymbol.(*antlr.CommonToken)
		el.Ctx.AgregarError("Error: "+"'"+token.GetText()+"' Comando Invalido", line, column)
	} else if e != nil {
		start := strings.Index(msg, "'")
		end := strings.LastIndex(msg, "'")
		lexema := msg[start+1 : end]
		el.Ctx.AgregarError("Error: '"+lexema+"' Lexema no Valido", line, column)
	}
}
