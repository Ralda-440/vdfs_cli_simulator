package comandos

import (
	structs "app/Interprete/Structs"
	"encoding/binary"
	"io"
	"os"
	"time"
)

type Mkfs struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

//mks *Mkfs ExprCommand

func (mks *Mkfs) Ejecutar(ctx *Contexto) interface{} {
	//Verificar si hay error
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(mks.Parametros) == 0 {
		ctx.AgregarError("El comando mkfs no tiene parametros", mks.Linea, mks.Columna)
		return nil
	}
	//Verifica parametro obligatorio -id
	if _, existe := mks.Parametros["-id"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -id", mks.Linea, mks.Columna)
		return nil
	}
	//Obtener id
	id := mks.Parametros["-id"]
	//Obtener parametros opcionales
	type_ := "Full"
	if val, existe := mks.Parametros["-type"]; existe {
		type_ = val
	}
	fs := "2fs"
	if val, existe := mks.Parametros["-fs"]; existe {
		fs = val
	}
	//Verificar valor correcto de -type
	if type_ != "Full" {
		ctx.AgregarError("Error: El parametro -type solo puede ser Full", mks.Linea, mks.Columna)
		return nil
	}
	//Verificar valor correcto de -fs
	if fs != "2fs" && fs != "3fs" {
		ctx.AgregarError("Error: El parametro -fs solo puede ser 2fs o 3fs", mks.Linea, mks.Columna)
		return nil
	}
	//Verificar que la particion ya este montada para poder formatearla con el sistema de archivos
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(id)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+id+"\" no esta montada", mks.Linea, mks.Columna)
		return nil
	}
	//Obtener MBR del disco donde esta la particion
	mbr, err := GetMBRDisk(partMontada.DiskName)
	if err != nil {
		ctx.AgregarError("Error : El disco no existe o hubo error al leer su MBR :"+err.Error(), mks.Linea, mks.Columna)
		return nil
	}
	//Buscar particion con el id
	existe, particion, _, ebr, _, _, _, err := mbr.BuscarParticion(partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error : No se pudo buscar la particion "+err.Error(), mks.Linea, mks.Columna)
		return nil
	}
	if !existe {
		ctx.AgregarError("Error : No existe la particion con el nombre: "+id, mks.Linea, mks.Columna)
		return nil
	}
	var start_part int64
	var size_part int64
	var tipo int64 //0 -> ext2 | 1 -> ext3
	//Obtener el byte de inicio y el tamaño de la particion
	if particion != nil {
		//Es particion primaria
		start_part = particion.Part_start
		size_part = particion.Part_s
	} else {
		//Es particion logica
		start_part = ebr.Part_start + int64(binary.Size(structs.EBR{}))
		size_part = ebr.Part_s - int64(binary.Size(structs.EBR{}))
	}
	//Tipo de sistema de archivos
	if fs == "3fs" {
		tipo = 1
	} else if fs == "2fs" {
		tipo = 0
	}
	//Limpiar la particion
	err = Clear(partMontada.DiskName, start_part, size_part)
	if err != nil {
		ctx.AgregarError("Error : No se pudo limpiar la particion Antes de Formater el Sistema de Archivos"+err.Error(), mks.Linea, mks.Columna)
		return nil
	}
	//Calcular Super Bloque
	superBloque := mks.CalcularSuperBloque(tipo, start_part, size_part)
	//Si es ext3 escribir un byte 0 al final del super bloque. Este indica cuantas operaciones se han realizado
	if tipo == 1 {
		file, err := os.OpenFile("./MIA/P1/"+partMontada.DiskName+".dsk", os.O_WRONLY, 0644)
		if err != nil {
			ctx.AgregarError("Error : No se pudo abrir el disco :"+err.Error(), mks.Linea, mks.Columna)
			return nil
		}
		defer file.Close()
		file.Seek(start_part+int64(binary.Size(Super_Bloque{})), io.SeekStart)
		binary.Write(file, binary.BigEndian, byte(0))
	}
	//Crear Carpeta Raiz
	superBloque.CrearCarpetaRaiz(ctx, partMontada.DiskName)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Crear Archivo de usuarios y grupos users.txt
	superBloque.CrearArchivoUsers(ctx, partMontada.DiskName)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Escribir Super Bloque
	superBloque.EscribirSuperBloque(ctx, partMontada.DiskName, start_part)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Mensaje de exito
	//fmt.Println("----------Comando MKFS--------------")
	//fmt.Println("Sistema de Archivos Formateado con exito")
	ctx.AgregarOutput("----------Comando MKFS--------------")
	ctx.AgregarOutput("Particion: \"" + partMontada.PartName + "\" con ID: \"" + id + "\" Formateada con exito con el Sistema de Archivos: " + fs)
	return nil
}

// Calcular numero de bloques ext2
// tipo = 0 -> ext2 | tipo = 1 -> ext3
func (mks *Mkfs) CalcularBloques(tipo int64, sizePart int64) int64 {
	if tipo == 1 {
		return (sizePart - int64(binary.Size(Super_Bloque{}))) / (4 + int64(binary.Size(structs.Journaling{})) + int64(binary.Size(structs.Inodes{})) + (3 * int64(binary.Size(structs.Block_Archivo{}))))
	}
	//n = int64(math.Floor(float64(n))) No es necesario ya que el resultado es entero
	return (sizePart - int64(binary.Size(Super_Bloque{}))) / (4 + int64(binary.Size(structs.Inodes{})) + (3 * int64(binary.Size(structs.Block_Archivo{}))))
}

// Calcular Super Bloque
func (mks *Mkfs) CalcularSuperBloque(tipo int64, start_part int64, sizePart int64) *Super_Bloque {
	//Calcular numero de bloques
	n := mks.CalcularBloques(tipo, sizePart)
	//Start despues del Super Bloque
	start := start_part + int64(binary.Size(Super_Bloque{}))
	//Sumar espacio del journaling si es ext3
	if tipo == 1 {
		start += n * int64(binary.Size(structs.Journaling{}))
	}
	//Inicio Bitmap Inodos
	start_bitmap_inodos := start
	//Inicio Bitmap Bloques
	start_bitmap_bloques := start_bitmap_inodos + n
	//Inicio Inodos
	start_inodos := start_bitmap_bloques + (3 * n)
	//Inicio Bloques
	start_bloques := start_inodos + (n * int64(binary.Size(structs.Inodes{})))
	//Crear Super Bloque
	superBloque := Super_Bloque{}
	superBloque.S_filesystem_type = tipo
	superBloque.S_inodes_count = n
	superBloque.S_blocks_count = n * 3
	superBloque.S_free_blocks_count = n * 3
	superBloque.S_free_inodes_count = n
	copy(superBloque.S_mtime[:], time.Now().Format("2006-01-02 15:04:05"))
	copy(superBloque.S_umtime[:], "....-..-.. --:--:--")
	superBloque.S_mnt_count = 1
	superBloque.S_magic = 0xEF53
	superBloque.S_inode_s = int64(binary.Size(structs.Inodes{}))
	superBloque.S_block_s = int64(binary.Size(structs.Block_Archivo{}))
	superBloque.S_first_ino = 0
	superBloque.S_first_blo = 0
	superBloque.S_bm_inode_start = start_bitmap_inodos
	superBloque.S_bm_block_start = start_bitmap_bloques
	superBloque.S_inode_start = start_inodos
	superBloque.S_block_start = start_bloques
	return &superBloque
}
