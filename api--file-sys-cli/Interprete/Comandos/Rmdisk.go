package comandos

import (
	"os"
)

type Rmdisk struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

func (rmd *Rmdisk) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(rmd.Parametros) == 0 {
		ctx.AgregarError("Error: El comando no tiene parametros", rmd.Linea, rmd.Columna)
		return nil
	}
	//Verificar si existe el parametro -driveletter
	name, existe := rmd.Parametros["-driveletter"]
	if !existe {
		ctx.AgregarError("Error: Falta el parametro -driveletter", rmd.Linea, rmd.Columna)
		return nil
	}
	//Verificar si existe el disco con el nombre
	files, err := os.ReadDir("./MIA/P1/")
	if err != nil {
		ctx.AgregarError("Error: No se pudo leer el directorio", rmd.Linea, rmd.Columna)
		return nil
	}
	existe = false
	name = name + ".dsk"
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if file.Name() == name {
			existe = true
			break
		}
	}
	if !existe {
		ctx.AgregarError("Error: No existe el disco con el nombre "+name, rmd.Linea, rmd.Columna)
		return nil
	}
	//Pedir confiramcion para eliminar el disco
	/*
		reader := bufio.NewReader(os.Stdin)
		for {
			println("Desea eliminar el disco " + name + " (S/N)")
			confirmacion, _ := reader.ReadString('\n')
			confirmacion = confirmacion[:len(confirmacion)-1]
			if confirmacion == "S" || confirmacion == "s" {
				break
			} else if confirmacion == "N" || confirmacion == "n" {
				return nil
			}
			println("Error: Opcion no valida")
		}
	*/
	//Eliminar el disco
	err = os.Remove("./MIA/P1/" + name)
	if err != nil {
		ctx.AgregarError("Error: No se pudo eliminar el disco", rmd.Linea, rmd.Columna)
		return nil
	}
	/*
		fmt.Println("----------Comando MKDISK----------")
		fmt.Println("Disco " + name + " eliminado correctamente")
	*/
	ctx.AgregarOutput("----------Comando RMDISK----------")
	ctx.AgregarOutput("Disco \"" + name + "\" eliminado correctamente")
	return nil
}
