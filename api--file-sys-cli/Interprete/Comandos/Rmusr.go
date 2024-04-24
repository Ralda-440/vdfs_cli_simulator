package comandos

import (
	"strings"
)

type Rmusr struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// rmusr *Rmusr ExprCommand
func (rmusr *Rmusr) Ejecutar(ctx *Contexto) interface{} {
	//Verificar si hay error
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(rmusr.Parametros) == 0 {
		ctx.AgregarError("El comando rmusr no tiene parametros", rmusr.Linea, rmusr.Columna)
		return nil
	}
	//Verifica parametro obligatorio -user
	if _, existe := rmusr.Parametros["-user"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -user", rmusr.Linea, rmusr.Columna)
		return nil
	}
	//Obtener user
	user := rmusr.Parametros["-user"]
	//Verificar si hay sesion activa
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay ninguna sesion activa", rmusr.Linea, rmusr.Columna)
		return nil
	}
	//Recuperar Sesion Actual
	userSesion, _, idPart := SesionActiva.SesionActual()
	//Verificar que el usuario sea root
	if userSesion != "root" {
		ctx.AgregarError("Error: Solo el usuario root puede eliminar usuarios", rmusr.Linea, rmusr.Columna)
		return nil
	}
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+idPart+"\" no esta montada", rmusr.Linea, rmusr.Columna)
		return nil
	}
	//Recuperar SuperBloque
	superBloque, err := GetSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", rmusr.Linea, rmusr.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Recuperar usuarios y grupos de la particion
	usuarios, grupos, err := superBloque.RecuperarUsuariosGrupos(ctx, partMontada.DiskName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar los usuarios y grupos de la particion", rmusr.Linea, rmusr.Columna)
		return nil
	}
	//Veirificar si el usuario existe
	existe = false
	for pos, u := range usuarios {
		if u.Usuario == user && u.UID != "0" {
			//Marcar UID como "0" para indicar que esta eliminado
			usuarios[pos].UID = "0"
			existe = true
			break
		}
	}
	if !existe {
		ctx.AgregarError("Error: No existe el usuario "+user, rmusr.Linea, rmusr.Columna)
		return nil
	}
	//Convertir usuarios y grupos a []byte
	cont := ConvertirUsuariosGrupos(usuarios, grupos)
	//Recuperar Inodo de users.txt
	pathStr := "/users.txt"
	path := strings.Split(pathStr, "/")
	//Eliminar el primer elemento
	path = path[1:]
	//Escribir contenido en users.txt
	superBloque.CrearYEscribirEnArchivo(ctx, partMontada.DiskName, path, cont, false)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Escribir SuperBloque en la particion
	superBloque.EscribirSuperBloque(ctx, partMontada.DiskName, partMontada.Start_SuperBloque)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Mensaje de exito
	//fmt.Println("----------Comando RMUSR--------------")
	//fmt.Println("Usuario \"" + user + "\" eliminado con exito")
	ctx.AgregarOutput("--------------------Comando RMUSR--------------------")
	ctx.AgregarOutput("Usuario \"" + user + "\" eliminado con exito en la particion con Id: \"" + idPart + "\"")
	return nil
}
