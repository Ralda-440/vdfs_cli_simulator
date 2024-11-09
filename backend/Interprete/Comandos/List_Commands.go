package comandos

type Commands struct {
	Commands []ExprCommand
}

func (c *Commands) Ejecutar(ctx *Contexto) interface{} {
	for _, command := range c.Commands {
		command.Ejecutar(ctx)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil
		}
	}
	return nil
}
