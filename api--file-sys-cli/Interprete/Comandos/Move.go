package comandos

import (
	"strings"
)

type Move struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// mv *Move ExprCommand
func (mv *Move) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar que haya parametros
	if len(mv.Parametros) == 0 {
		ctx.AgregarError("El comando mv no tiene parametros", mv.Linea, mv.Columna)
		return nil
	}
	//Verificar parametros obligatorios -path -destino
	if _, existe := mv.Parametros["-path"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -path en mv", mv.Linea, mv.Columna)
		return nil
	}
	if _, existe := mv.Parametros["-destino"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -destino en mv", mv.Linea, mv.Columna)
		return nil
	}
	//Obtener parametros
	path := mv.Parametros["-path"]
	destino := mv.Parametros["-destino"]
	//Verificar si hay sesion activa
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay ninguna sesion activa", mv.Linea, mv.Columna)
		return nil
	}
	//Recuperar Sesion Activa
	_, _, idPart := SesionActiva.SesionActual()
	//Buscar Particion Montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: No hay ninguna particion montada", mv.Linea, mv.Columna)
		return nil
	}
	//Obtener SuperBloque
	superBloque, err := GetSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo obtener el super bloque de la particion", mv.Linea, mv.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Converti path a []byte
	pathArray := strings.Split(path, "/")
	//Eliminar el primer elemento
	pathArray = pathArray[1:]
	//Nombre del inodo
	nombreInodo := pathArray[len(pathArray)-1]
	//Obtener Inodo
	inodo, _, _, _ := superBloque.RecuperarInodoPorPath(ctx, partMontada.DiskName, pathArray, 0, 0, false, true, true)
	if inodo == nil {
		ctx.AgregarError("Error: No se encontro el archivo en el Origen"+path, mv.Linea, mv.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Crear inodo para el destino
	destinoArray := strings.Split(destino, "/")
	//Eliminar el primer elemento
	destinoArray = destinoArray[1:]
	//Agregar el nombre del archivo
	destinoArray = append(destinoArray, nombreInodo)
	//Verificar si el archivo ya existe en el destino
	inode, _, _, _ := superBloque.RecuperarInodoPorPath(ctx, partMontada.DiskName, destinoArray, 0, 0, false, true, false)
	if inode != nil {
		ctx.AgregarError("Error: El archivo ya existe en el Destino", mv.Linea, mv.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Crear Inodo para el destino
	inode_, numInode, _, _ := superBloque.RecuperarInodoPorPath(ctx, partMontada.DiskName, destinoArray, 0, int(inodo.I_type), false, false, false)
	if inode_ == nil {
		ctx.AgregarError("Error: No se pudo crear el inodo para el destino", mv.Linea, mv.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Escribir inodo en el numInode
	superBloque.EscribirInodo(ctx, partMontada.DiskName, inodo, int64(numInode))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Imprimir mensaje de exito
	//fmt.Println("----------Comando MOVE--------------")
	//fmt.Println("Se movio el archivo: " + path + " al destino: " + destino)
	ctx.AgregarOutput("-------------Comando MOVE--------------")
	ctx.AgregarOutput("Se movio el archivo: \"" + path + "\" al destino: \"" + destino + "\" correctamente")
	return nil
}
