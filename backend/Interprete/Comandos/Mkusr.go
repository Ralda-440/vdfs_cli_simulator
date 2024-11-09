package comandos

import (
	"strconv"
	"strings"
)

type Mkusr struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// mkusr *Mkusr ExprCommand
func (mkusr *Mkusr) Ejecutar(ctx *Contexto) interface{} {
	//Verificar si hay error
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(mkusr.Parametros) == 0 {
		ctx.AgregarError("El comando mkusr no tiene parametros", mkusr.Linea, mkusr.Columna)
		return nil
	}
	//Verifica parametro obligatorio -user -pass -grp
	if _, existe := mkusr.Parametros["-user"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -user", mkusr.Linea, mkusr.Columna)
		return nil
	}
	if _, existe := mkusr.Parametros["-pass"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -pass", mkusr.Linea, mkusr.Columna)
		return nil
	}
	if _, existe := mkusr.Parametros["-grp"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -grp", mkusr.Linea, mkusr.Columna)
		return nil
	}
	//Obtener user, pass, grp
	user := mkusr.Parametros["-user"]
	pass := mkusr.Parametros["-pass"]
	grp := mkusr.Parametros["-grp"]
	//Verificar si hay sesion activa
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay ninguna sesion activa", mkusr.Linea, mkusr.Columna)
		return nil
	}
	//Recuperar Sesion Actual
	userSesion, _, idPart := SesionActiva.SesionActual()
	//Verificar que el usuario sea root
	if userSesion != "root" {
		ctx.AgregarError("Error: Solo el usuario root puede crear usuarios", mkusr.Linea, mkusr.Columna)
		return nil
	}
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+idPart+"\" no esta montada", mkusr.Linea, mkusr.Columna)
		return nil
	}
	//Recuperar SuperBloque
	superBloque, err := GetSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", mkusr.Linea, mkusr.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Recuperar usuarios y grupos de la particion
	users, grupos, err := superBloque.RecuperarUsuariosGrupos(ctx, partMontada.DiskName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar los usuarios y grupos de la particion", mkusr.Linea, mkusr.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si el usuario ya existe
	for _, usuario := range users {
		if usuario.Usuario == user && usuario.UID != "0" {
			ctx.AgregarError("Error: El usuario \""+user+"\" ya existe", mkusr.Linea, mkusr.Columna)
			return nil
		}
	}
	//Verificar si el grupo existe
	existeGrupo := false
	for _, grupo := range grupos {
		if grupo.Grupo == grp && grupo.GID != "0" {
			existeGrupo = true
			break
		}
	}
	if !existeGrupo {
		ctx.AgregarError("Error: El grupo \""+grp+"\" no existe", mkusr.Linea, mkusr.Columna)
		return nil
	}
	//Crear usuario
	UID := superBloque.GenerarUID(users)
	usuario := &UsuarioParticion{strconv.Itoa(UID), "U", grp, user, pass}
	//Agregar usuario
	users = append(users, *usuario)
	//Convertir usuarios y grupos a []byte
	cont := ConvertirUsuariosGrupos(users, grupos)
	//Recuperar el inodo del archivo users.txt
	pathStr := "/users.txt"
	path := strings.Split(pathStr, "/")
	//Eliminar el primer elemento
	path = path[1:]
	//Escribir el archivo users.txt
	superBloque.CrearYEscribirEnArchivo(ctx, partMontada.DiskName, path, cont, false)
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo crear el usuario", mkusr.Linea, mkusr.Columna)
		return nil
	}
	//Escribir SuperBloque en la particion
	superBloque.EscribirSuperBloque(ctx, partMontada.DiskName, partMontada.Start_SuperBloque)
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo crear el usuario", mkusr.Linea, mkusr.Columna)
		return nil
	}
	//Mensaje de exito
	//fmt.Println("----------Comando MKUSR--------------")
	//fmt.Println("Usuario creado con exito: " + user)
	ctx.AgregarOutput("--------------------Comando MKUSR--------------------")
	ctx.AgregarOutput("Usuario: \"" + user + "\" con UID: \"" + strconv.Itoa(UID) + "\" y Grupo: \"" + grp + "\" creado con exito en la particion con Id: \"" + idPart + "\"")
	return nil
}
