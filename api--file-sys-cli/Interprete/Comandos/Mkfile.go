package comandos

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mkfile struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// mkf *Mkfile ExprCommand

func (mkf *Mkfile) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar que haya parametros
	if len(mkf.Parametros) == 0 {
		ctx.AgregarError("El comando mkfile no tiene parametros", mkf.Linea, mkf.Columna)
		return nil
	}
	//Verificar parametros obligatorios -path
	if _, existe := mkf.Parametros["-path"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -path en Mkfile", mkf.Linea, mkf.Columna)
		return nil
	}
	//Obtener parametros
	path := mkf.Parametros["-path"]
	//Parametros opcionales
	size := ""
	cont := ""
	//Obtener parametros opcionales
	var haySize, hayCont, r bool
	if _, existe := mkf.Parametros["-r"]; existe {
		r = true
		rValue := mkf.Parametros["-r"]
		//No se permite Valor para r
		if rValue != "" {
			ctx.AgregarError("Error: El parametro -r no tiene que tener valor", mkf.Linea, mkf.Columna)
			return nil
		}
	}
	if _, existe := mkf.Parametros["-size"]; existe {
		haySize = true
		size = mkf.Parametros["-size"]
	}
	if _, existe := mkf.Parametros["-cont"]; existe {
		hayCont = true
		cont = mkf.Parametros["-cont"]
	}
	//Verificar si hay sesion activa
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay ninguna sesion activa", mkf.Linea, mkf.Columna)
		return nil
	}
	//Recuperar Sesion Actual
	_, _, idPart := SesionActiva.SesionActual()
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+idPart+"\" no esta montada", mkf.Linea, mkf.Columna)
		return nil
	}
	//Recuperar SuperBloque
	superBloque, err := getSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", mkf.Linea, mkf.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Obtener Contenido a escribir en el archivo
	var contenido []byte
	if haySize {
		//Veiricar que size sea un entero y mayor a 0
		size_, err := strconv.Atoi(size)
		if err != nil {
			ctx.AgregarError("Error: El parametro -size debe ser un entero en Mkfile", mkf.Linea, mkf.Columna)
			return nil
		}
		if size_ <= 0 {
			ctx.AgregarError("Error: El parametro -size debe ser mayor a 0 en Mkfile", mkf.Linea, mkf.Columna)
			return nil
		}
		//Llenar contenido con numero del 0 al 9 hasta llegar al size
		contenido = make([]byte, size_)
		for i := 0; i < size_; i++ {
			contenido[i] = byte(i % 10)
		}
		//Convertir los bytes binarios a su ASCII correspondiente
		for i := 0; i < size_; i++ {
			contenido[i] = contenido[i] + 48
		}
	} else if hayCont {
		//Obtener Contenido de una archivo de texto. El path es cont
		bytes, err := os.ReadFile(cont)
		if err != nil {
			ctx.AgregarError("Error al leer el archivo de texto en Mkfile: "+err.Error(), mkf.Linea, mkf.Columna)
			return nil
		}
		contenido = bytes
	}
	//Convertir path en []string
	path_ := strings.Split(path, "/")
	//Eliminar el primer elemento
	path_ = path_[1:]
	//Crear Archivo
	superBloque.CrearYEscribirEnArchivo(ctx, partMontada.DiskName, path_, contenido, r)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Imprimir mensaje de exito
	fmt.Println("----------Comando MKFILE--------------")
	fmt.Println("Archivo creado con exito")
	return nil
}
