package comandos

type ExprCommand interface {
	Ejecutar(ctx *Contexto) interface{}
}
