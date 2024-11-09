package comandos

import (
	structs "app/Interprete/Structs"
	"encoding/binary"
	"io"
	"os"
	"sort"
	"strconv"
	"time"
)

type Mkdisk struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

func (mkd *Mkdisk) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(mkd.Parametros) == 0 {
		ctx.AgregarError("Error: El comando no tiene Parametros", mkd.Linea, mkd.Columna)
		return nil
	}
	//Verificar si existe el parametro -size
	value, existe := mkd.Parametros["-size"]
	sizeOrigen := value
	if !existe {
		ctx.AgregarError("Error: Falta el parametro -size", mkd.Linea, mkd.Columna)
		return nil
	}
	//Convertir el valor a entero
	size, err := strconv.Atoi(value)
	if err != nil {
		ctx.AgregarError("Error: El valor del parametro -size no es valido", mkd.Linea, mkd.Columna)
		return nil
	}
	//Verificar que el tamaño sea mayor a 0
	if size <= 0 {
		ctx.AgregarError("Error: El valor del parametro -size debe ser mayor a 0", mkd.Linea, mkd.Columna)
		return nil
	}
	//Obtener parametro -fit
	value, existe = mkd.Parametros["-fit"]
	if !existe {
		value = "F"
	}
	//Guardar fit
	fit := value
	//Verificar que el valor sea B, F o W
	if fit != "B" && fit != "F" && fit != "W" {
		ctx.AgregarError("Error: El valor del parametro -fit no es valido", mkd.Linea, mkd.Columna)
		return nil
	}
	//Obtener parametro -unit
	value, existe = mkd.Parametros["-unit"]
	if !existe {
		value = "M"
	}
	//Guardar unit
	unit := value
	//Verificar que el valor sea M o K
	if unit != "M" && unit != "K" {
		ctx.AgregarError("Error: El valor del parametro -unit no es valido", mkd.Linea, mkd.Columna)
		return nil
	}
	//Calcular el tamaño del disco
	if unit == "K" {
		size = size * 1024
	} else if unit == "M" {
		size = size * 1024 * 1024
	}
	//####Crear el disco
	nameDisk := mkd.NameDisk(ctx)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Crear el archivo del disco
	file, err := os.Create("./MIA/P1/" + nameDisk)
	if err != nil {
		ctx.AgregarError("Error al crear el disco: "+err.Error(), mkd.Linea, mkd.Columna)
		return nil
	}
	defer file.Close()
	//Escribir '/0' en el archivo hasta completar el tamaño
	content := make([]byte, size)
	_, err = file.Write(content)
	if err != nil {
		ctx.AgregarError("Error al escribir el disco: "+err.Error(), mkd.Linea, mkd.Columna)
		return nil
	}
	//Crear MBR del disco
	newMBR := structs.NewMBR()
	//Tamaño del disco
	newMBR.Mbr_tamanio = int64(size)
	//Fecha y Hora
	now := time.Now()
	fecha_hora := now.Format("2006-01-02 15:04:05")
	copy(newMBR.Mbr_fecha_creacion[:], fecha_hora)
	//Firma del disco (Numero random)
	newMBR.Mbr_disk_signature = generateRandomSignature()
	//Fit
	newMBR.Mbr_disk_fit = fit[0]
	//Escribir MBR en el disco
	/*
		file, err = os.OpenFile("./MIA/P1/A.dsk", os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Error al abrir el disco: ", err)
			return nil
		}
		defer file.Close()
	*/
	//Mover el puntero al inicio del archivo
	file.Seek(0, io.SeekStart)
	//Escribir el MBR
	err = binary.Write(file, binary.BigEndian, &newMBR)
	if err != nil {
		ctx.AgregarError("Error al escribir el MBR: "+err.Error(), mkd.Linea, mkd.Columna)
		return nil
	}
	ctx.AgregarOutput("----------Comando MKDISK----------")
	ctx.AgregarOutput("Disco \"" + nameDisk + "\" creado con exito. Con Size: " + sizeOrigen + " " + unit)
	return nil
}

func (mkd *Mkdisk) NameDisk(ctx *Contexto) string {
	files, err := os.ReadDir("./MIA/P1/")
	if err != nil {
		ctx.AgregarError("Error al leer el directorio: "+err.Error(), mkd.Linea, mkd.Columna)
		return ""
	}
	var nombresArchivos []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		nombresArchivos = append(nombresArchivos, file.Name())
	}

	if len(nombresArchivos) == 0 {
		return "A.dsk"
	}

	sort.Strings(nombresArchivos)
	ultimoNombre := nombresArchivos[len(nombresArchivos)-1]
	ultimaLetraASCII := ultimoNombre[0]
	sigLetra := string(ultimaLetraASCII + 1)
	return sigLetra + ".dsk"
}
