package comandos

import (
	"strings"
)

type Rename struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// rn *Rename ExprCommand
func (rn *Rename) Ejecutar(ctx *Contexto) interface{} {
	//Verificar si hay error
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(rn.Parametros) == 0 {
		ctx.AgregarError("El comando rename no tiene parametros", rn.Linea, rn.Columna)
		return nil
	}
	//Verifica parametro obligatorio -path y -name
	if _, existe := rn.Parametros["-path"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -path", rn.Linea, rn.Columna)
		return nil
	}
	if _, existe := rn.Parametros["-name"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -name", rn.Linea, rn.Columna)
		return nil
	}
	//Obtener path y name
	path := rn.Parametros["-path"]
	name := rn.Parametros["-name"]
	//Convertir path a []string
	path_ := strings.Split(path, "/")
	//Eliminar el primer elemento
	path_ = path_[1:]
	//Verificar si hay sesion activa
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay ninguna sesion activa", rn.Linea, rn.Columna)
		return nil
	}
	//Recuperar Sesion Actual
	_, _, idPart := SesionActiva.SesionActual()
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+idPart+"\" no esta montada", rn.Linea, rn.Columna)
		return nil
	}
	//Recuperar SuperBloque
	superBloque, err := getSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", rn.Linea, rn.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Renombrar Archivo o Carpeta
	superBloque.RenombrarArchivoOCarpeta(ctx, partMontada.DiskName, path_, name)
	//Verirficar Errores
	if ctx.HayErrores() {
		return nil
	}
	//Mensaje de exito
	//fmt.Println("----------Comando RENAME--------------")
	//fmt.Println("El archivo/carpeta: \"" + path + "\" se ha renombrado a: \"" + name + "\" correctamente")
	ctx.AgregarOutput("-------------Comando Rename--------------")
	ctx.AgregarOutput("El archivo/carpeta: \"" + path + "\" se ha renombrado a: \"" + name + "\" correctamente")
	return nil
}
