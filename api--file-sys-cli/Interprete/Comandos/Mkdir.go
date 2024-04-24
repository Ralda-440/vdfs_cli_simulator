package comandos

import (
	"strings"
)

type Mkdir struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// mkdir *Mkdir ExprCommand
func (mkdir *Mkdir) Ejecutar(ctx *Contexto) interface{} {
	//Verificar si hay error
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(mkdir.Parametros) == 0 {
		ctx.AgregarError("Error : El comando mkdir no tiene parametros", mkdir.Linea, mkdir.Columna)
		return nil
	}
	//Verifica parametro obligatorio -path
	if _, existe := mkdir.Parametros["-path"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -path", mkdir.Linea, mkdir.Columna)
		return nil
	}
	//Obtener path
	path := mkdir.Parametros["-path"]
	//Veriricar parametros opcionales -r
	r := false
	if _, existe := mkdir.Parametros["-r"]; existe {
		r = true
		rValue := mkdir.Parametros["-r"]
		//No se permite Valor para r
		if rValue != "" {
			ctx.AgregarError("Error: El parametro -r no tiene que tener valor", mkdir.Linea, mkdir.Columna)
			return nil
		}
	}
	//Verificar si hay sesion activa
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay ninguna sesion activa", mkdir.Linea, mkdir.Columna)
		return nil
	}
	//Recuperar Sesion Actual
	_, _, idPart := SesionActiva.SesionActual()
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+idPart+"\" no esta montada", mkdir.Linea, mkdir.Columna)
		return nil
	}
	//Recuperar SuperBloque
	superBloque, err := GetSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", mkdir.Linea, mkdir.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Crear carpeta
	//Convertir path a []string
	path_ := strings.Split(path, "/")
	//Eliminar el primer elemento
	path_ = path_[1:]
	//Crear carpeta
	superBloque.CrearCarpeta(ctx, partMontada.DiskName, path_, r)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Si es ext3 actualizar Journal
	if superBloque.S_filesystem_type == 1 {
		comandoConsola := "mkdir "
		//Agregar los Parametros
		for k, v := range mkdir.Parametros {
			if v == "" {
				comandoConsola += k + " "
				continue
			}
			comandoConsola += k + "=" + v + " "
		}
		superBloque.RegistrarJournaling(ctx, partMontada.DiskName, partMontada.Start_SuperBloque, "mkdir", path, "", comandoConsola)
		//Verificar si hay errores
		if ctx.HayErrores() {
			ctx.AgregarError("Error: No se pudo registrar en el journal", mkdir.Linea, mkdir.Columna)
			return nil
		}
	}
	//Mensaje de exito
	//fmt.Println("----------Comando MKDIR--------------")
	//fmt.Println("Se creo la carpeta con exito: " + path + " en la particion con id: " + idPart)
	ctx.AgregarOutput("-------------Comando MKDIR--------------")
	ctx.AgregarOutput("La carpeta: \"" + path + "\" se creo correctamente")
	return nil
}
