package comandos

type Logout struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// lgt *Logout ExprCommand

func (lgt *Logout) Ejecutar(ctx *Contexto) interface{} {
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(lgt.Parametros) > 0 {
		ctx.AgregarError("El comando logout no acepta parametros", lgt.Linea, lgt.Columna)
		return nil
	}
	//Verificar si hay una sesion activa
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay una sesion activa para cerrar sesion", lgt.Linea, lgt.Columna)
		return nil
	}
	//Cerrar sesion
	usr, idPart := SesionActiva.CerrarSesion()
	//Mensaje de exito
	//fmt.Println("----------Comando LOGOUT--------------")
	//fmt.Println("Sesion cerrada con exito")
	ctx.AgregarOutput("--------------------Comando LOGOUT--------------------")
	ctx.AgregarOutput("Usuario: \"" + usr + "\" ha cerrado sesion en la particion con id: " + idPart)
	return nil
}
