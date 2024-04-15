package comandos

type Contexto struct {
	ListaErrores []Error
	Output       string
}

func (c *Contexto) AgregarError(msg string, linea int, columna int) {
	c.ListaErrores = append(c.ListaErrores, Error{Msg: msg, Linea: linea, Columna: columna})
}

func (c *Contexto) AgregarOutput(output string) {
	c.Output += output + "\n"
}

func (c *Contexto) HayErrores() bool {
	return len(c.ListaErrores) > 0
}

// Imprimir Errores
func (c *Contexto) ImprimirErrores() {
	for _, err := range c.ListaErrores {
		println(err.Msg, "Linea: ", err.Linea, "Columna: ", err.Columna)
	}
}

// Errores
func (c *Contexto) GetErrores() []Error {
	errs := c.ListaErrores
	c.ListaErrores = []Error{}
	return errs
}

// Output
func (c *Contexto) GetOutput() string {
	ouput := c.Output
	c.Output = ""
	return ouput
}

func NewContexto() *Contexto {
	return &Contexto{ListaErrores: []Error{}, Output: ""}
}

type Error struct {
	Msg     string `json:"msg"`
	Linea   int    `json:"linea"`
	Columna int    `json:"columna"`
}
