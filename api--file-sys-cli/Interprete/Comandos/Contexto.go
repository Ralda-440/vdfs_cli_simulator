package comandos

type Contexto struct {
	ListaErrores []Error
}

func (c *Contexto) AgregarError(msg string, linea int, columna int) {
	c.ListaErrores = append(c.ListaErrores, Error{Msg: msg, Linea: linea, Columna: columna})
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

func NewContexto() *Contexto {
	return &Contexto{ListaErrores: []Error{}}
}

type Error struct {
	Msg     string
	Linea   int
	Columna int
}
