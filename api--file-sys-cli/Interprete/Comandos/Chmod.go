package comandos

import (
	"fmt"
	"strconv"
	"strings"
)

type Chmod struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// chm *Chmod ExprCommand
func (chm *Chmod) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar que haya parametros
	if len(chm.Parametros) == 0 {
		ctx.AgregarError("El comando chmod no tiene parametros", chm.Linea, chm.Columna)
		return nil
	}
	//Verificar parametros obligatorios -path -ugo
	if _, existe := chm.Parametros["-path"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -path en chmod", chm.Linea, chm.Columna)
		return nil
	}
	if _, existe := chm.Parametros["-ugo"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -ugo en chmod", chm.Linea, chm.Columna)
		return nil
	}
	//Obtener parametros
	path := chm.Parametros["-path"]
	ugo := chm.Parametros["-ugo"]
	//Verificar si hay sesion activa
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay ninguna sesion activa", chm.Linea, chm.Columna)
		return nil
	}
	//Recuperar Sesion Activa
	user, _, idPart := SesionActiva.SesionActual()
	//Solo root puede cambiar permisos
	if user != "root" {
		ctx.AgregarError("Error: Solo el usuario root puede cambiar permisos", chm.Linea, chm.Columna)
		return nil
	}
	//Buscar Particion Montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: No hay ninguna particion montada", chm.Linea, chm.Columna)
		return nil
	}
	//Obtener SuperBloque
	superBloque, err := getSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo obtener el super bloque de la particion", chm.Linea, chm.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Recuperar Inodo por Path
	//Convertir path a []string
	pathArr := strings.Split(path, "/")
	//Eliminar el primer elemento
	pathArr = pathArr[1:]
	//Recuperar Inodo
	inodo, numInodo, _, _ := superBloque.RecuperarInodoPorPath(ctx, partMontada.DiskName, pathArr, 0, 0, false, true, false)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar formato de ugo
	if len(ugo) != 3 {
		ctx.AgregarError("Error: El parametro -ugo debe tener 3 digitos", chm.Linea, chm.Columna)
		return nil
	}
	//Verificar que cada posicion de ugo sea un digito y este entre 0 y 7
	for i := 0; i < 3; i++ {
		if ugo[i] < 48 || ugo[i] > 55 {
			ctx.AgregarError("Error: El parametro -ugo debe tener digitos entre 0 y 7", chm.Linea, chm.Columna)
			return nil
		}
	}
	//Cambiar permisos
	permU, _ := strconv.Atoi(string(ugo[0]))
	permG, _ := strconv.Atoi(string(ugo[1]))
	permO, _ := strconv.Atoi(string(ugo[2]))
	inodo.I_perm = [3]byte{byte(permU), byte(permG), byte(permO)}
	//Escribir Inodo
	superBloque.EscribirInodo(ctx, partMontada.DiskName, inodo, int64(numInodo))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Mensaje de exito
	fmt.Println("----------Comando CHOWN--------------")
	fmt.Println("Se cambiaron los permisos del archivo o carpeta en: " + path)
	return nil
}
