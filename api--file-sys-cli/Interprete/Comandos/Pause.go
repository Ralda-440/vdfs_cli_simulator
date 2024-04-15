package comandos

type Pause struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// ps *Pause ExprCommand
func (ps *Pause) Ejecutar(ctx *Contexto) interface{} {
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Veiricar que el parametro no tenga parametros
	if len(ps.Parametros) > 0 {
		return "El comando pause no recibe parametros"
	}
	//Pausar y esperar a que el usuario presione enter
	//fmt.Println("Presiona Enter para continuar...")
	//ps.waitForEnter()
	//fmt.Println("Continuando con la ejecución del programa...")
	ctx.AgregarOutput("--------------------Comando PAUSE--------------------")
	ctx.AgregarOutput("PAUSE")
	return nil
}

/*
func (ps *Pause) waitForEnter() {
	var input string
	fmt.Scanln(&input)
	cmd := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1")
	cmd.Stdin = os.Stdin
	_ = cmd.Run()
}
*/
