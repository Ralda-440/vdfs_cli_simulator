package comandos

import (
	"fmt"
	"strings"
)

type Remove struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// rm *Remove ExprCommand
func (rm *Remove) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar que haya parametros
	if len(rm.Parametros) == 0 {
		ctx.AgregarError("El comando rm no tiene parametros", rm.Linea, rm.Columna)
		return nil
	}
	//Verificar parametros obligatorios -path
	if _, existe := rm.Parametros["-path"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -path en rm", rm.Linea, rm.Columna)
		return nil
	}
	//Obtener parametros
	path := rm.Parametros["-path"]
	//Convertir Path a []string
	pathArray := strings.Split(path, "/")
	//Eliminar el primer elemento
	pathArray = pathArray[1:]
	//Verificar si hay sesion activa
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay ninguna sesion activa", rm.Linea, rm.Columna)
		return nil
	}
	//Recuperar Sesion Actual
	_, _, idPart := SesionActiva.SesionActual()
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+idPart+"\" no esta montada", rm.Linea, rm.Columna)
		return nil
	}
	//Recuperar SuperBloque
	superBloque, err := getSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", rm.Linea, rm.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Eliminar Archivo o Carpeta
	seElimino := superBloque.EliminarInodoPorPath(ctx, partMontada.DiskName, pathArray)
	//Verirficar Errores
	if ctx.HayErrores() {
		return nil
	}
	if !seElimino {
		ctx.AgregarError("Error: No se pudo eliminar el archivo o carpeta por falta de Permisos", rm.Linea, rm.Columna)
		return nil
	}
	//Mensaje de Exito
	fmt.Println("----------Comando Remove--------------")
	fmt.Println("El archivo o carpeta se elimino correctamente")
	return nil
}
