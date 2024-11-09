package comandos

import (
	"os"
	"strings"
)

type Edit struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// ed *Edid ExprCommand
func (ed *Edit) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar que haya parametros
	if len(ed.Parametros) == 0 {
		ctx.AgregarError("El comando edit no tiene parametros", ed.Linea, ed.Columna)
		return nil
	}
	//Verificar parametros obligatorios -path y -cont
	if _, existe := ed.Parametros["-path"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -path en edit", ed.Linea, ed.Columna)
		return nil
	}
	if _, existe := ed.Parametros["-cont"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -cont en edit", ed.Linea, ed.Columna)
		return nil
	}
	//Obtener parametros
	path := ed.Parametros["-path"]
	cont := ed.Parametros["-cont"]
	//Convertir Path a []string
	pathArray := strings.Split(path, "/")
	//Eliminar el primer elemento
	pathArray = pathArray[1:]
	//Abrir archivo de contenido
	datos, err := os.ReadFile(cont)
	if err != nil {
		ctx.AgregarError("Error: No se pudo abrir el archivo de contenido", ed.Linea, ed.Columna)
		return nil
	}
	//Verificar si hay sesion activa
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay ninguna sesion activa", ed.Linea, ed.Columna)
		return nil
	}
	//Recuperar Sesion Actual
	_, _, idPart := SesionActiva.SesionActual()
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+idPart+"\" no esta montada", ed.Linea, ed.Columna)
		return nil
	}
	//Recuperar SuperBloque
	superBloque, err := GetSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", ed.Linea, ed.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Editar Archivo
	superBloque.EditarContenidoArchivo(ctx, partMontada.DiskName, pathArray, datos)
	//Verirficar Errores
	if ctx.HayErrores() {
		return nil
	}
	//Mensaje de exito
	//fmt.Println("----------Comando EDIT--------------")
	//fmt.Println("Se edito el archivo con exito: "+path, "en la particion con id: "+idPart)
	ctx.AgregarOutput("-------------Comando Edit--------------")
	ctx.AgregarOutput("El archivo: \"" + path + "\" se edito correctamente")
	return nil
}
