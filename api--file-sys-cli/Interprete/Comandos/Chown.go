package comandos

import (
	"fmt"
	"strconv"
	"strings"
)

type Chown struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// chown *Chown ExprCommand
func (chown *Chown) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar que haya parametros
	if len(chown.Parametros) == 0 {
		ctx.AgregarError("El comando chown no tiene parametros", chown.Linea, chown.Columna)
		return nil
	}
	//Verificar parametros obligatorios -path -user
	if _, existe := chown.Parametros["-path"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -path en chown", chown.Linea, chown.Columna)
		return nil
	}
	if _, existe := chown.Parametros["-user"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -usr en chown", chown.Linea, chown.Columna)
		return nil
	}
	//Obtener parametros
	path := chown.Parametros["-path"]
	usr := chown.Parametros["-user"]
	//Verificar si hay sesion activa
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay ninguna sesion activa", chown.Linea, chown.Columna)
		return nil
	}
	//Recuperar Sesion Activa
	_, _, idPart := SesionActiva.SesionActual()
	//Buscar Particion Montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: No hay ninguna particion montada", chown.Linea, chown.Columna)
		return nil
	}
	//Obtener SuperBloque
	superBloque, err := getSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo obtener el super bloque de la particion", chown.Linea, chown.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Recuperar Inodo por Path
	//Convertir path a []string
	arrPath := strings.Split(path, "/")
	//Eliminar el primer elemento
	arrPath = arrPath[1:]
	//Recuperar Inodo
	inodo, numInodo, _, _ := superBloque.RecuperarInodoPorPath(ctx, partMontada.DiskName, arrPath, 0, 0, false, true, false)
	if inodo == nil {
		ctx.AgregarError("Error: No se encontro el inodo para cambiar sus permisos", chown.Linea, chown.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Recuperar Usuarios
	usuario, _, err := superBloque.RecuperarUsuariosGrupos(ctx, partMontada.DiskName)
	if err != nil {
		ctx.AgregarError("Error: No se pudieron recuperar los usuarios y grupos", chown.Linea, chown.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si el usuario existe
	existeUsuario := false
	for _, u := range usuario {
		if u.Usuario == usr {
			existeUsuario = true
			//Cambiar propietario
			uid, err := strconv.Atoi(u.UID)
			if err != nil {
				ctx.AgregarError("Error: No se pudo convertir el uid a entero", chown.Linea, chown.Columna)
				return nil
			}
			inodo.I_uid = int64(uid)
			break
		}
	}
	if !existeUsuario {
		ctx.AgregarError("Error: El usuario \""+usr+"\" no existe. No sepuede cambiar de propietario", chown.Linea, chown.Columna)
		return nil
	}
	//Escribir Inodo
	superBloque.EscribirInodo(ctx, partMontada.DiskName, inodo, int64(numInodo))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	fmt.Println("----------Comando CHOWN--------------")
	fmt.Println("Se cambio el propietario de la carpeta/archivo: " + path + " al usuario: " + usr)
	return nil
}
