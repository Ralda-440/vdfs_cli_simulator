package comandos

import (
	structs "app/Interprete/Structs"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Super_Bloque struct {
	S_filesystem_type   int64    // 0 = EXT2, 1 = EXT3
	S_inodes_count      int64    // Cantidad total inodos
	S_blocks_count      int64    // Cantidad total bloques
	S_free_blocks_count int64    // Cantidad bloques libres
	S_free_inodes_count int64    // Cantidad inodos libres
	S_mtime             [19]byte // Ultima fecha de montaje
	S_umtime            [19]byte // Ultima fecha de desmontaje
	S_mnt_count         int64    // Cantidad de montajes
	S_magic             int64    // Numero magico 0xEF53
	S_inode_s           int64    // Tamaño de un inodo
	S_block_s           int64    // Tamaño de un bloque
	S_first_ino         int64    // Primer inodo libre
	S_first_blo         int64    // Primer bloque libre
	S_bm_inode_start    int64    // Inicio del bitmap de inodos
	S_bm_block_start    int64    // Inicio del bitmap de bloques
	S_inode_start       int64    // Inicio de la tabla de inodos
	S_block_start       int64    // Inicio de la tabla de bloques
}

// Limpiar Journaling
func (sb *Super_Bloque) LimpiarJournaling(ctx *Contexto, driveLetter string, start_Part int64) {
	//La cantidad de Journaling total es la cantidad de inodos totales
	//Crear []byte del tamaño de la cantidad de inodos
	journaling := make([]byte, sb.S_inodes_count*int64(binary.Size(structs.Journaling{})))
	//Escribir el journaling en el disco
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_WRONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error: No se pudo abrir el disco", 0, 0)
		return
	}
	defer file.Close()
	//Mover el puntero al inicio del journaling
	file.Seek(start_Part+int64(binary.Size(Super_Bloque{})), io.SeekStart)
	//Escribir el journaling
	binary.Write(file, binary.BigEndian, journaling)
}

// Recuperar Journaling
func (sb *Super_Bloque) RecuperarJournaling(ctx *Contexto, driveLetter string, start_Part int64) []structs.Journaling {
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_RDWR, 0644)
	if err != nil {
		ctx.AgregarError("Error: No se pudo abrir el disco", 0, 0)
		return nil
	}
	defer file.Close()
	//Obtener la cantidad de operaciones en el journaling. Es el primer byte despues del super bloque
	file.Seek(start_Part+int64(binary.Size(Super_Bloque{})), io.SeekStart)
	operaciones := make([]byte, 1)
	binary.Read(file, binary.BigEndian, &operaciones)
	//Mover el puntero al inicio del journaling
	file.Seek(start_Part+int64(binary.Size(Super_Bloque{}))+int64(binary.Size(operaciones)), io.SeekStart)
	//Recuperar todo el Journaling
	ops := []structs.Journaling{}
	for i := 0; i < int(operaciones[0]); i++ {
		journaling := structs.Journaling{}
		binary.Read(file, binary.BigEndian, &journaling)
		ops = append(ops, journaling)
	}
	return ops
}

// Registrar en Journaling
func (sb *Super_Bloque) RegistrarJournaling(ctx *Contexto, driveLetter string, start_Part int64, comando string, path string, contenido string, comandoConsola string) {
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_RDWR, 0644)
	if err != nil {
		ctx.AgregarError("Error: No se pudo abrir el disco", 0, 0)
		return
	}
	defer file.Close()
	//Obtener la cantidad de operaciones en el journaling. Es el primer byte despues del super bloque
	file.Seek(start_Part+int64(binary.Size(Super_Bloque{})), io.SeekStart)
	operaciones := make([]byte, 1)
	binary.Read(file, binary.BigEndian, &operaciones)
	//Mover el puntero al inicio del journaling
	file.Seek(start_Part+int64(binary.Size(Super_Bloque{}))+int64(binary.Size(operaciones)+(int(operaciones[0])*binary.Size(structs.Journaling{}))), io.SeekStart)
	//Crear Journaling
	journaling := structs.Journaling{}
	copy(journaling.Comando[:], comando)
	copy(journaling.Path[:], path)
	copy(journaling.Contenido[:], contenido)
	copy(journaling.ComandoConsola[:], comandoConsola)
	copy(journaling.Fecha[:], time.Now().Format("2006-01-02 15:04:05"))
	//Escribir Journaling
	binary.Write(file, binary.BigEndian, &journaling)
	//Aumentar la cantidad de operaciones
	operaciones[0]++
	//Mover el puntero al inicio del journaling
	file.Seek(start_Part+int64(binary.Size(Super_Bloque{})), io.SeekStart)
	//Escribir la cantidad de operaciones
	binary.Write(file, binary.BigEndian, operaciones)
}

// Limpiar Bitmaps, Area de Inodos y Bloques
func (sb *Super_Bloque) LimpiarBitmapsInodosBloques(ctx *Contexto, driveLetter string) {
	//[]byte del tamaño de la cantidad de bloques
	bitmapBloques := make([]byte, sb.S_blocks_count)
	//[]byte del tamaño de la cantidad de inodos
	bitmapInodos := make([]byte, sb.S_inodes_count)
	//[]byte del tamaño de un inodo * cantidad de inodos
	areaInodos := make([]byte, sb.S_inode_s*sb.S_inodes_count)
	//[]byte del tamaño de un bloque * cantidad de bloques
	areaBloques := make([]byte, sb.S_block_s*sb.S_blocks_count)
	//Escribir los bitmap y areas en el disco
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_WRONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error: No se pudo abrir el disco", 0, 0)
		return
	}
	defer file.Close()
	//Mover el puntero al inicio del bitmap de bloques
	file.Seek(sb.S_bm_block_start, io.SeekStart)
	//Escribir el bitmap de bloques
	binary.Write(file, binary.BigEndian, bitmapBloques)
	//Mover el puntero al inicio del bitmap de inodos
	file.Seek(sb.S_bm_inode_start, io.SeekStart)
	//Escribir el bitmap de inodos
	binary.Write(file, binary.BigEndian, bitmapInodos)
	//Mover el puntero al inicio de la tabla de inodos
	file.Seek(sb.S_inode_start, io.SeekStart)
	//Escribir el area de inodos
	binary.Write(file, binary.BigEndian, areaInodos)
	//Mover el puntero al inicio de la tabla de bloques
	file.Seek(sb.S_block_start, io.SeekStart)
	//Escribir el area de bloques
	binary.Write(file, binary.BigEndian, areaBloques)
	//Reiniciar cantidad de bloques e inodos libres
	sb.S_free_blocks_count = sb.S_blocks_count
	sb.S_free_inodes_count = sb.S_inodes_count
	//Reiniciar primer bloque e inodo libre
	sb.S_first_blo = 0
	sb.S_first_ino = 0
}

// Reporte ls
func (sb *Super_Bloque) ReporteLs(ctx *Contexto, driveLetter string, path []string) (string, map[string]string) {
	//Recuperar inodo del archivo
	var inodo *structs.Inodes = nil
	if len(path) == 0 {
		inodo = sb.RecuperarInodo(ctx, driveLetter, 0)
	} else {
		inodo, _, _, _ = sb.RecuperarInodoPorPath(ctx, driveLetter, path, 0, 1, false, true, false)
	}
	if inodo == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No existe el path para listar el archivo", 0, 0)
		return "", nil
	}
	//Verificar que el inodo ser Carpeta
	if inodo.I_type != 0 {
		ctx.AgregarError("Error: El path no es de una carpeta", 0, 0)
		return "", nil
	}
	//Obtener Graphviz
	graphviz, content := sb.ReporteLsInodo(ctx, driveLetter, inodo.I_block[:])
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo listar el archivo", 0, 0)
		return "", nil
	}
	graphviz = `
	digraph G {
		repLs [label=< 
				<table border="1" cellspacing="0" cellpadding='10' >
					<tr>
						<td bgcolor="#2F4F4F"><font color="white"><b>Permisos</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Owner</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Grupo</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Size (en bytes)</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Fecha</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Hora</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Tipo</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Name</b></font></td>
					</tr>
					` + graphviz + `
				</table> 
				> shape=box style=invisible ] 
	}
	`
	return graphviz, content
}

// Reporte ls inodo
func (sb *Super_Bloque) ReporteLsInodo(ctx *Contexto, driveLetter string, apuntadores []int32) (string, map[string]string) {
	graphviz := ""
	content := map[string]string{}
	//Iterar todos los apuntadores
	for pos, numBloque := range apuntadores {
		if numBloque == -1 {
			continue
		}
		if pos == 12 {
			//Graphviz del Bloque indirecto simple
			graphviz_, content_ := sb.ReporteLsInodoApuntadorIndirectoSimple(ctx, driveLetter, int(numBloque))
			//Verificar si hay errores
			if ctx.HayErrores() {
				return "", nil
			}
			graphviz += graphviz_
			for k, v := range content_ {
				content[k] = v
			}
			continue
		} else if pos == 13 {
			//Graphviz del Bloque indirecto doble
			graphviz_, content_ := sb.ReporteLsInodoApuntadorIndirectoDoble(ctx, driveLetter, int(numBloque))
			//Verificar si hay errores
			if ctx.HayErrores() {
				return "", nil
			}
			graphviz += graphviz_
			for k, v := range content_ {
				content[k] = v
			}
			continue
		} else if pos == 14 {
			//Graphviz del Bloque indirecto triple
			graphviz_, content_ := sb.ReporteLsInodoApuntadorIndirectoTriple(ctx, driveLetter, int(numBloque))
			//Verificar si hay errores
			if ctx.HayErrores() {
				return "", nil
			}
			graphviz += graphviz_
			for k, v := range content_ {
				content[k] = v
			}
			continue
		}
		//Recuperar bloque
		bloque := sb.RecuperarBloqueCarpeta(ctx, driveLetter, int(numBloque))
		if bloque == nil || ctx.HayErrores() {
			ctx.AgregarError("Error: No se pudo recuperar el bloque", 0, 0)
			return "", nil
		}
		//Iterar todos los archivos y carpetas
		for _, contenido := range bloque.B_content {
			if contenido.B_inodo == -1 {
				continue
			}
			nombre := string(bytes.Replace(contenido.B_name[:], []byte{0}, []byte{}, -1))
			if nombre == "." || nombre == ".." {
				continue
			}
			//Recuperar inodo
			inodo := sb.RecuperarInodo(ctx, driveLetter, int(contenido.B_inodo))
			//Verificar si hay errores
			if ctx.HayErrores() {
				return "", nil
			}
			//Buscar UID y GID del inodo
			owner := ""
			grupo := ""
			infUsuario, infGrupo := sb.BuscarUsuarioGrupoPorID(ctx, driveLetter, inodo.I_uid, inodo.I_gid)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return "", nil
			}
			if infUsuario != nil {
				owner = infUsuario.Usuario
			} else {
				owner = "-------"
			}
			if infGrupo != nil {
				grupo = infGrupo.Grupo
			} else {
				grupo = "-------"
			}
			//Agregar al graphviz
			_, graphvizLs := inodo.GetGraph(int(contenido.B_inodo), nombre, true, owner, grupo)
			graphviz += graphvizLs
			//Tipo content
			var tipo string
			if inodo.I_type == 0 {
				tipo = "dir"
			} else {
				tipo = "file"
			}
			content[nombre] = tipo
		}
	}
	return graphviz, content
}

// Reporte ls inodo apuntador indirecto simple
func (sb *Super_Bloque) ReporteLsInodoApuntadorIndirectoSimple(ctx *Contexto, driveLetter string, numBloque int) (string, map[string]string) {
	graphviz := ""
	content := map[string]string{}
	//Recuperar bloque
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, numBloque)
	if bloque == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo recuperar el bloque", 0, 0)
		return "", nil
	}
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		if numBloque == -1 {
			continue
		}
		//Recuperar bloque
		bloque := sb.RecuperarBloqueCarpeta(ctx, driveLetter, int(numBloque))
		if bloque == nil || ctx.HayErrores() {
			ctx.AgregarError("Error: No se pudo recuperar el bloque", 0, 0)
			return "", nil
		}
		//Iterar todos los archivos y carpetas
		for _, contenido := range bloque.B_content {
			if contenido.B_inodo == -1 {
				continue
			}
			nombre := string(bytes.Replace(contenido.B_name[:], []byte{0}, []byte{}, -1))
			//Recuperar inodo
			inodo := sb.RecuperarInodo(ctx, driveLetter, int(contenido.B_inodo))
			//Verificar si hay errores
			if ctx.HayErrores() {
				return "", nil
			}
			//Buscar UID y GID del inodo
			owner := ""
			grupo := ""
			infUsuario, infGrupo := sb.BuscarUsuarioGrupoPorID(ctx, driveLetter, inodo.I_uid, inodo.I_gid)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return "", nil
			}
			if infUsuario != nil {
				owner = infUsuario.Usuario
			} else {
				owner = "-------"
			}
			if infGrupo != nil {
				grupo = infGrupo.Grupo
			} else {
				grupo = "-------"
			}
			//Agregar al graphviz
			_, graphvizLs := inodo.GetGraph(int(contenido.B_inodo), nombre, true, owner, grupo)
			graphviz += graphvizLs
			//Tipo content
			var tipo string
			if inodo.I_type == 0 {
				tipo = "dir"
			} else {
				tipo = "file"
			}
			content[nombre] = tipo
		}
	}
	return graphviz, content
}

// Reporte ls inodo apuntador indirecto doble
func (sb *Super_Bloque) ReporteLsInodoApuntadorIndirectoDoble(ctx *Contexto, driveLetter string, numBloque int) (string, map[string]string) {
	graphviz := ""
	content := map[string]string{}
	//Recuperar bloque
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, numBloque)
	if bloque == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo recuperar el bloque", 0, 0)
		return "", nil
	}
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		if numBloque == -1 {
			continue
		}
		//Agregar al graphviz
		graphviz_, content_ := sb.ReporteLsInodoApuntadorIndirectoSimple(ctx, driveLetter, int(numBloque))
		//Verificar si hay errores
		if ctx.HayErrores() {
			return "", nil
		}
		graphviz += graphviz_
		content = content_
	}
	return graphviz, content
}

// Reporte ls inodo apuntador indirecto triple
func (sb *Super_Bloque) ReporteLsInodoApuntadorIndirectoTriple(ctx *Contexto, driveLetter string, numBloque int) (string, map[string]string) {
	graphviz := ""
	content := map[string]string{}
	//Recuperar bloque
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, numBloque)
	if bloque == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo recuperar el bloque", 0, 0)
		return "", nil
	}
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		if numBloque == -1 {
			continue
		}
		//Agregar al graphviz
		graphviz_, content_ := sb.ReporteLsInodoApuntadorIndirectoDoble(ctx, driveLetter, int(numBloque))
		//Verificar si hay errores
		if ctx.HayErrores() {
			return "", nil
		}
		graphviz += graphviz_
		content = content_
	}
	return graphviz, content
}

// Reporte Tree Graphviz
// 0 = Reporte Tree, 1 = Reporte Solo Inodos, 2 = Reporte Solo Bloques
func (sb *Super_Bloque) ReporteGraphviz(ctx *Contexto, driveLetter string, tipo int) (string, string) {
	graphviz, graphvizLs := sb.ReporteInodoGraphviz(ctx, driveLetter, 0, "Raiz", tipo)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return "", ""
	}
	horientacion := "TB"
	if tipo == 0 {
		horientacion = "LR"
	}
	graphviz = `
	digraph {
		node [shape=plaintext]
		rankdir=` + horientacion + `
	` +
		graphviz + `
	}`
	graphvizLs = `
	digraph G {
		repLs [label=< 
				<table border="1" cellspacing="0" cellpadding='10' >
					<tr>
						<td bgcolor="#2F4F4F"><font color="white"><b>Permisos</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Owner</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Grupo</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Size (en bytes)</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Fecha</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Hora</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Tipo</b></font></td>
						<td bgcolor="#2F4F4F"><font color="white"><b>Name</b></font></td>
					</tr>
					` + graphvizLs + `
				</table> 
				> shape=box style=invisible ] 
	}
	`
	return graphviz, graphvizLs
}

// Reporte Inodo Archivo Graphviz
func (sb *Super_Bloque) ReporteInodoGraphviz(ctx *Contexto, driveLetter string, numInodo int, nombre string, tipo int) (string, string) {
	//Recuperar Inodo
	inodo := sb.RecuperarInodo(ctx, driveLetter, numInodo)
	if inodo == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo recuperar el inodo", 0, 0)
		return "", ""
	}
	graphviz := ""
	graphvizls := ""
	//Buscar UID y GID del inodo
	owner := ""
	grupo := ""
	infUsuario, infGrupo := sb.BuscarUsuarioGrupoPorID(ctx, driveLetter, inodo.I_uid, inodo.I_gid)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return "", ""
	}
	if infUsuario != nil {
		owner = infUsuario.Usuario
	} else {
		owner = "-------"
	}
	if infGrupo != nil {
		grupo = infGrupo.Grupo
	} else {
		grupo = "-------"
	}
	//Graphviz del Inodo
	if tipo == 0 {
		graphviz, graphvizls = inodo.GetGraph(numInodo, nombre, false, owner, grupo)
	} else if tipo == 1 {
		graphviz, graphvizls = inodo.GetGraph(numInodo, nombre, true, owner, grupo)
	}
	//Iterar todos los bloques
	for pos, numBloque := range inodo.I_block {
		if numBloque == -1 {
			continue
		}
		if pos == 12 {
			//Graphviz del Bloque indirecto simple
			graphviz += sb.ReporteGraphvizBloqueIndirectoSimple(ctx, driveLetter, int(numBloque), inodo.I_type, tipo)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return "", ""
			}
			continue
		} else if pos == 13 {
			//Graphviz del Bloque indirecto doble
			graphviz += sb.ReporteGraphvizBloqueIndirectoDoble(ctx, driveLetter, int(numBloque), inodo.I_type, tipo)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return "", ""
			}
			continue
		} else if pos == 14 {
			//Graphviz del Bloque indirecto triple
			graphviz += sb.ReporteGraphvizBloqueIndirectoTriple(ctx, driveLetter, int(numBloque), inodo.I_type, tipo)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return "", ""
			}
			continue
		}
		//Verificar si es inodo Carpeta o Archivo
		if inodo.I_type == 1 {
			//Recuperar Bloque
			bloque := sb.RecuperarBloqueArchivo(ctx, driveLetter, int(numBloque))
			if bloque == nil || ctx.HayErrores() {
				ctx.AgregarError("Error: No se pudo recuperar el bloque", 0, 0)
				return "", ""
			}
			//Graphviz del Bloque
			if tipo == 0 || tipo == 2 {
				graphviz += bloque.GetGraph(int(numBloque))
			}
			continue
		}
		//Recuperar Bloque
		bloque := sb.RecuperarBloqueCarpeta(ctx, driveLetter, int(numBloque))
		if bloque == nil || ctx.HayErrores() {
			ctx.AgregarError("Er(numBloqueror: No se pudo recuperar el bloque", 0, 0)
			return "", ""
		}
		//Graphviz del Bloque
		primerBloque := false
		if pos == 0 {
			primerBloque = true
		}
		//Graphviz del Bloque
		if tipo == 0 {
			graphviz += bloque.GetGraph(int(numBloque), primerBloque, false)
		} else if tipo == 2 {
			graphviz += bloque.GetGraph(int(numBloque), primerBloque, true)
		}
		//Iterar todos los archivos y carpetas
		for _, contenido := range bloque.B_content {
			if contenido.B_inodo == -1 {
				continue
			}
			nombre := string(bytes.Replace(contenido.B_name[:], []byte{0}, []byte{}, -1))
			if nombre == "." || nombre == ".." {
				continue
			}
			//Graphviz del Inodo
			graphviz_, graphvizls_ := sb.ReporteInodoGraphviz(ctx, driveLetter, int(contenido.B_inodo), nombre, tipo)
			graphviz += graphviz_
			graphvizls += graphvizls_
			//Verificar si hay errores
			if ctx.HayErrores() {
				return "", ""
			}
		}
	}
	return graphviz, graphvizls
}

// Reporte Graphviz Bloque indirecto simple
func (sb *Super_Bloque) ReporteGraphvizBloqueIndirectoSimple(ctx *Contexto, driveLetter string, numBloque int, tipo byte, tipoRep int) string {
	//Recuperar Bloque indirecto simple
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, numBloque)
	if bloque == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo recuperar el bloque indirecto simple", 0, 0)
		return ""
	}
	//Graphviz del Bloque
	graphviz := ""
	graphvizls := ""
	if tipoRep == 0 {
		graphviz = bloque.GetGraph(numBloque, false)
	} else if tipoRep == 1 {
		graphviz = bloque.GetGraph(numBloque, true)
	}
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		if numBloque == -1 {
			continue
		}
		//Verificar si es inodo Carpeta o Archivo
		if tipo == 1 {
			//Recuperar Bloque
			bloque := sb.RecuperarBloqueArchivo(ctx, driveLetter, int(numBloque))
			if bloque == nil || ctx.HayErrores() {
				ctx.AgregarError("Error: No se pudo recuperar el bloque", 0, 0)
				return ""
			}
			//Graphviz del Bloque
			if tipoRep == 0 || tipoRep == 2 {
				graphviz += bloque.GetGraph(int(numBloque))
			}
			continue
		}
		//Recuperar Bloque
		bloque := sb.RecuperarBloqueCarpeta(ctx, driveLetter, int(numBloque))
		if bloque == nil || ctx.HayErrores() {
			ctx.AgregarError("Error: No se pudo recuperar el bloque", 0, 0)
			return ""
		}
		//Graphviz del Bloque
		if tipoRep == 0 {
			graphviz += bloque.GetGraph(int(numBloque), false, false)
		} else if tipoRep == 2 {
			graphviz += bloque.GetGraph(int(numBloque), false, true)
		}
		//Iterar todos los archivos y carpetas
		for _, contenido := range bloque.B_content {
			if contenido.B_inodo == -1 {
				continue
			}
			nombre := string(bytes.Replace(contenido.B_name[:], []byte{0}, []byte{}, -1))
			//Graphviz del Inodo
			graphviz_, graphvizls_ := sb.ReporteInodoGraphviz(ctx, driveLetter, int(contenido.B_inodo), nombre, tipoRep)
			graphviz += graphviz_
			graphvizls += graphvizls_
			//Verificar si hay errores
			if ctx.HayErrores() {
				return ""
			}
		}
	}
	return graphviz
}

// Reporte Graphviz Bloque indirecto doble
func (sb *Super_Bloque) ReporteGraphvizBloqueIndirectoDoble(ctx *Contexto, driveLetter string, numBloque int, tipo byte, tipoRep int) string {
	//Recuperar Bloque indirecto doble
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, numBloque)
	if bloque == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo recuperar el bloque indirecto doble", 0, 0)
		return ""
	}
	//Graphviz del Bloque
	graphviz := bloque.GetGraph(numBloque, false)
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		if numBloque == -1 {
			continue
		}
		//Graphviz del Bloque indirecto simple
		graphviz += sb.ReporteGraphvizBloqueIndirectoSimple(ctx, driveLetter, int(numBloque), tipo, tipoRep)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return ""
		}
	}
	return graphviz
}

// Reporte Graphviz Bloque indirecto triple
func (sb *Super_Bloque) ReporteGraphvizBloqueIndirectoTriple(ctx *Contexto, driveLetter string, numBloque int, tipo byte, tipoRep int) string {
	//Recuperar Bloque indirecto triple
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, numBloque)
	if bloque == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo recuperar el bloque indirecto triple", 0, 0)
		return ""
	}
	//Graphviz del Bloque
	graphviz := bloque.GetGraph(numBloque, false)
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		if numBloque == -1 {
			continue
		}
		//Graphviz del Bloque indirecto doble
		graphviz += sb.ReporteGraphvizBloqueIndirectoDoble(ctx, driveLetter, int(numBloque), tipo, tipoRep)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return ""
		}
	}
	return graphviz
}

// Renombrar Archivo o Carpeta
func (sb *Super_Bloque) RenombrarArchivoOCarpeta(ctx *Contexto, driveLetter string, path []string, nombreNuevo string) {
	//Verificar si el nuevo nombre ya existe
	nuevoPath := make([]string, len(path))
	copy(nuevoPath, path)
	nuevoPath[len(nuevoPath)-1] = nombreNuevo
	inode, numInode, _, _ := sb.RecuperarInodoPorPath(ctx, driveLetter, nuevoPath, 0, 1, false, true, false)
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo verificar si el nuevo nombre ya existe", 0, 0)
		return
	}
	if inode != nil && numInode != -1 {
		ctx.AgregarError("Error: El nuevo nombre ya existe: "+nombreNuevo, 0, 0)
		return
	}
	//Recuperar inodo del archivo
	inodo, numInodo, numBlock, posBlock := sb.RecuperarInodoPorPath(ctx, driveLetter, path, 0, 1, false, true, false)
	if inodo == nil || numInodo == -1 || numBlock == -1 || posBlock == -1 || ctx.HayErrores() {
		ctx.AgregarError("Error: No existe el path para renombrar el archivo o carpeta", 0, 0)
		return
	}
	//Editar la posicion del bloque que contiene el archivo o carpeta
	bloque := sb.RecuperarBloqueCarpeta(ctx, driveLetter, int(numBlock))
	if bloque == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo recuperar el bloque para renombrar el archivo o carpeta", 0, 0)
		return
	}
	//Cambiar el nombre del archivo o carpeta
	//Limpiar el nombre
	bloque.B_content[posBlock].B_name = [12]byte{}
	copy(bloque.B_content[posBlock].B_name[:], nombreNuevo)
	//Escribir el bloque en el disco
	sb.EscribirBloque(ctx, driveLetter, nil, bloque, nil, int64(numBlock))
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo renombrar el archivo o carpeta", 0, 0)
		return
	}
}

// Eliminar inodo y sus bloques
func (sb *Super_Bloque) EliminarInodoPorPath(ctx *Contexto, driveLetter string, path []string) bool {
	copia_path := make([]string, len(path))
	copy(copia_path, path)
	//Recuperar inodo del archivo
	inodo, numInodo, _, _ := sb.RecuperarInodoPorPath(ctx, driveLetter, path, 0, 1, false, true, false)
	if inodo == nil || numInodo == -1 || ctx.HayErrores() {
		ctx.AgregarError("Error: No existe el path para eliminar el archivo", 0, 0)
		return false
	}
	//Eliminar inodo y sus bloques
	seElimino := sb.EliminarInodo(ctx, driveLetter, int32(numInodo))
	if ctx.HayErrores() {
		return false
	}
	if !seElimino {
		ctx.AgregarError("Error: No se pudo eliminar el archivo :"+copia_path[len(copia_path)-1]+" por falta de Permisos", 0, 0)
		return false
	}
	//Eliminar el inodo del padre
	inodo, numInodo, _, _ = sb.RecuperarInodoPorPath(ctx, driveLetter, copia_path, 0, 1, false, true, true)
	if inodo == nil || numInodo == -1 || ctx.HayErrores() {
		ctx.AgregarError("Error: No existe el path para eliminar el archivo", 0, 0)
		return false
	}
	//Eliminar el inodo del padre)
	return seElimino
}

// Eliminar Inodo
func (sb *Super_Bloque) EliminarInodo(ctx *Contexto, driveLetter string, numInodo int32) bool {
	//Recuperar inodo
	inodo := sb.RecuperarInodo(ctx, driveLetter, int(numInodo))
	if inodo == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo recuperar el inodo para eliminarlo", 0, 0)
		return false
	}
	//Verificar que tipo de inodo es
	if inodo.I_type == 1 {
		//Es un archivo. Para eliminarlo, escribir en el 0 bytes para que se liberen los bloques
		sb.EscribirEnBloquesArchivo(ctx, driveLetter, inodo.I_block[:], make([]byte, 0))
		if ctx.HayErrores() {
			ctx.AgregarError("Error: No se pudo eliminar el archivo", 0, 0)
			return false
		}
		return true
	}
	//Iterar todos los apuntadores
	for pos, numBloque := range inodo.I_block {
		if numBloque == -1 {
			continue
		}
		if pos == 12 {
			seElimino := sb.EliminarBloqueIndirectoSimpleInodo(ctx, driveLetter, numBloque)
			if !seElimino {
				ctx.AgregarError("Error: No se pudo eliminar el bloque indirecto simple", 0, 0)
				return false
			}
			continue
		} else if pos == 13 {
			seElimino := sb.EliminarBloqueIndirectoDobleInodo(ctx, driveLetter, numBloque)
			if !seElimino {
				ctx.AgregarError("Error: No se pudo eliminar el bloque indirecto doble", 0, 0)
				return false
			}
			continue
		} else if pos == 14 {
			seElimino := sb.EliminarBloqueIndirectoTripleInodo(ctx, driveLetter, numBloque)
			if !seElimino {
				ctx.AgregarError("Error: No se pudo eliminar el bloque indirecto triple", 0, 0)
				return false
			}
			continue
		}
		//Recuperar bloque
		bloque := sb.RecuperarBloqueCarpeta(ctx, driveLetter, int(numBloque))
		if bloque == nil || ctx.HayErrores() {
			ctx.AgregarError("Error: No se pudo recuperar el bloque para eliminarlo", 0, 0)
			return false
		}
		//Iterar todos los archivos y carpetas
		for _, contenido := range bloque.B_content {
			if contenido.B_inodo == -1 {
				continue
			}
			//No hacer nada si es "." o ".."
			if string(bytes.Replace(contenido.B_name[:], []byte{0}, []byte{}, -1)) == "." || string(bytes.Replace(contenido.B_name[:], []byte{0}, []byte{}, -1)) == ".." {
				continue
			}
			seElimino := sb.EliminarInodo(ctx, driveLetter, contenido.B_inodo)
			if !seElimino {
				ctx.AgregarError("Error: No se pudo eliminar el inodo", 0, 0)
				return false
			}
			//Liberar el inodo
			sb.ModificarBitEnBitMap(ctx, driveLetter, contenido.B_inodo, true, false)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return false
			}
		}
		//Eliminar bloque
		sb.ModificarBitEnBitMap(ctx, driveLetter, numBloque, false, false)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return false
		}
	}
	//Liberar el inodo
	sb.ModificarBitEnBitMap(ctx, driveLetter, int32(numInodo), true, false)
	//Verificar si hay errores
	return !ctx.HayErrores()
}

// Eliminar Bloque indirecto simple de un inodo
func (sb *Super_Bloque) EliminarBloqueIndirectoSimpleInodo(ctx *Contexto, driveLetter string, numBloque int32) bool {
	//Recuperar bloque indirecto simple
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, int(numBloque))
	if bloque == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo recuperar el bloque indirecto simple", 0, 0)
		return false
	}
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		if numBloque == -1 {
			continue
		}
		//Recuperar Bloque Carpeta
		bloqueCarpeta := sb.RecuperarBloqueCarpeta(ctx, driveLetter, int(numBloque))
		if bloqueCarpeta == nil || ctx.HayErrores() {
			ctx.AgregarError("Error: No se pudo recuperar el bloque carpeta", 0, 0)
			return false
		}
		//Iterar todos los archivos y carpetas
		for _, contenido := range bloqueCarpeta.B_content {
			if contenido.B_inodo == -1 {
				continue
			}
			seElimino := sb.EliminarInodo(ctx, driveLetter, contenido.B_inodo)
			if !seElimino {
				ctx.AgregarError("Error: No se pudo eliminar el inodo", 0, 0)
				return false
			}
			//Liberar el inodo
			sb.ModificarBitEnBitMap(ctx, driveLetter, contenido.B_inodo, true, false)
		}
		//Eliminar bloque
		sb.ModificarBitEnBitMap(ctx, driveLetter, numBloque, false, false)
	}
	//Liberar el bloque indirecto simple
	sb.ModificarBitEnBitMap(ctx, driveLetter, int32(numBloque), false, false)
	//Verificar si hay errores
	return !ctx.HayErrores()
}

// Eliminar Bloque indirecto doble de un inodo
func (sb *Super_Bloque) EliminarBloqueIndirectoDobleInodo(ctx *Contexto, driveLetter string, numBloque int32) bool {
	//Recuperar bloque indirecto doble
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, int(numBloque))
	if bloque == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo recuperar el bloque indirecto doble", 0, 0)
		return false
	}
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		if numBloque == -1 {
			continue
		}
		seElimino := sb.EliminarBloqueIndirectoSimpleInodo(ctx, driveLetter, numBloque)
		if !seElimino {
			ctx.AgregarError("Error: No se pudo eliminar el bloque indirecto simple", 0, 0)
			return false
		}
		//Liberar el bloque
		sb.ModificarBitEnBitMap(ctx, driveLetter, numBloque, false, false)
	}
	//Liberar el bloque indirecto doble
	sb.ModificarBitEnBitMap(ctx, driveLetter, int32(numBloque), false, false)
	//Verificar si hay errores
	return !ctx.HayErrores()
}

// Eliminar Bloque indirecto triple de un inodo
func (sb *Super_Bloque) EliminarBloqueIndirectoTripleInodo(ctx *Contexto, driveLetter string, numBloque int32) bool {
	//Recuperar bloque indirecto triple
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, int(numBloque))
	if bloque == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo recuperar el bloque indirecto triple", 0, 0)
		return false
	}
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		if numBloque == -1 {
			continue
		}
		seElimino := sb.EliminarBloqueIndirectoDobleInodo(ctx, driveLetter, numBloque)
		if !seElimino {
			ctx.AgregarError("Error: No se pudo eliminar el bloque indirecto doble", 0, 0)
			return false
		}
		//Liberar el bloque
		sb.ModificarBitEnBitMap(ctx, driveLetter, numBloque, false, false)
	}
	//Liberar el bloque indirecto triple
	sb.ModificarBitEnBitMap(ctx, driveLetter, int32(numBloque), false, false)
	//Verificar si hay errores
	return !ctx.HayErrores()
}

// Crear Inodo
func (sb *Super_Bloque) CrearInodo(ctx *Contexto, driveLetter string, padre int32, tipo int) int32 {
	//Obtener siguiente inodo libre
	numInodo := sb.SiguienteBitLibre(ctx, driveLetter, true)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return -1
	}
	//Marcar el inodo como ocupado
	sb.ModificarBitEnBitMap(ctx, driveLetter, numInodo, true, true)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return -1
	}
	//Crear nuevo inodo
	inodo := structs.NewInode()
	//Asignar valores al inodo
	UID, err := strconv.Atoi(SesionActiva.UID)
	if err != nil {
		ctx.AgregarError("Error: No se pudo convertir el UID de la sesion activa", 0, 0)
		return -1
	}
	GUI, err := strconv.Atoi(SesionActiva.GID)
	if err != nil {
		ctx.AgregarError("Error: No se pudo convertir el GID de la sesion activa", 0, 0)
		return -1
	}
	inodo.I_uid = int64(UID)
	inodo.I_gid = int64(GUI)
	inodo.I_s = 0
	copy(inodo.I_atime[:], time.Now().Format("2006-01-02 15:04:05"))
	copy(inodo.I_ctime[:], time.Now().Format("2006-01-02 15:04:05"))
	copy(inodo.I_mtime[:], time.Now().Format("2006-01-02 15:04:05"))
	inodo.I_type = byte(tipo)
	inodo.I_perm = [3]byte{6, 6, 4}
	//Si es una carpeta, agregar "." y ".."
	if tipo == 0 {
		//Obtener siguiente bloque libre
		numBloque := sb.SiguienteBitLibre(ctx, driveLetter, false)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return -1
		}
		//Marcar el bloque como ocupado
		sb.ModificarBitEnBitMap(ctx, driveLetter, numBloque, false, true)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return -1
		}
		//Crear nuevo Bloque
		bloque := structs.NewBlock_Carpeta()
		//Escribir "." y ".." en el bloque
		copy(bloque.B_content[0].B_name[:], ".")
		bloque.B_content[0].B_inodo = int32(numInodo)
		copy(bloque.B_content[1].B_name[:], "..")
		bloque.B_content[1].B_inodo = padre
		//Escribir bloque en el disco
		sb.EscribirBloque(ctx, driveLetter, nil, bloque, nil, int64(numBloque))
		//Verificar si hay errores
		if ctx.HayErrores() {
			return -1
		}
		//Asignar el bloque al inodo
		inodo.I_block[0] = int32(numBloque)
	}
	//Escribir inodo en el disco
	sb.EscribirInodo(ctx, driveLetter, inodo, int64(numInodo))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return -1
	}
	return numInodo
}

// Leer Contenido de Archivo
func (sb *Super_Bloque) LeerContenidoArchivo(ctx *Contexto, driveLetter string, path []string) []byte {
	//Recuperar inodo del archivo
	inodo, _, _, _ := sb.RecuperarInodoPorPath(ctx, driveLetter, path, 0, 1, false, true, false)
	if inodo == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No existe el path para leer el archivo", 0, 0)
		return nil
	}
	//Leer contenido del archivo
	contenido := sb.LeerBloquesArchivo(ctx, driveLetter, inodo.I_block[:])
	if contenido == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo leer el contenido del archivo", 0, 0)
		return nil
	}
	return contenido
}

// Editar Contenido de Archivo
func (sb *Super_Bloque) EditarContenidoArchivo(ctx *Contexto, driveLetter string, path []string, contenido []byte) {
	//Recuperar inodo del archivo
	inodo, numInodo, _, _ := sb.RecuperarInodoPorPath(ctx, driveLetter, path, 0, 1, false, true, false)
	if inodo == nil || numInodo == -1 || ctx.HayErrores() {
		ctx.AgregarError("Error: No existe el path para editar el archivo", 0, 0)
		return
	}
	//Escribir en Inodo de Archivo
	apuntadoresModificados, contenidoRestante := sb.EscribirEnBloquesArchivo(ctx, driveLetter, inodo.I_block[:], contenido)
	if apuntadoresModificados == nil || contenidoRestante == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo escribir en el inodo del archivo", 0, 0)
		return
	}
	//Warning si hay contenido restante
	if len(contenidoRestante) > 0 {
		//No agregar error, solo warning
		fmt.Println("Warning: No se pudo escribir todo el contenido en el archivo")
	}
	//Actualizar apuntadores en el inodo
	copy(inodo.I_block[:], apuntadoresModificados)
	//Actualizar size del inodo
	inodo.I_s = int64(len(contenido))
	//Actualizar el inodo en el disco
	sb.EscribirInodo(ctx, driveLetter, inodo, int64(numInodo))
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo actualizar el inodo del archivo", 0, 0)
		return
	}
}

// Crear Carpeta
func (sb *Super_Bloque) CrearCarpeta(ctx *Contexto, driveLetter string, path []string, r bool) {
	//Recuperar inodo del archivo
	inodo, numInodo, _, _ := sb.RecuperarInodoPorPath(ctx, driveLetter, path, 0, 0, r, false, false)
	if inodo == nil || numInodo == -1 || ctx.HayErrores() {
		ctx.AgregarError("Error: No existe el path para crear la carpeta", 0, 0)
		return
	}
	//Escribir inodo en el disco
	sb.EscribirInodo(ctx, driveLetter, inodo, int64(numInodo))
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo crear la carpeta", 0, 0)
		return
	}
}

// Escribir en Archivo
func (sb *Super_Bloque) CrearYEscribirEnArchivo(ctx *Contexto, driveLetter string, path []string, contenido []byte, r bool) {
	//Verificar si hay Inodos y Bloques libres
	if sb.S_free_blocks_count < 1 || sb.S_free_inodes_count < 1 {
		ctx.AgregarError("Error: No hay inodos o bloques libres", 0, 0)
		return
	}
	//Recuperar inodo del archivo
	inodo, numInodo, _, _ := sb.RecuperarInodoPorPath(ctx, driveLetter, path, 0, 1, r, false, false)
	if inodo == nil || numInodo == -1 || ctx.HayErrores() {
		ctx.AgregarError("Error: No existe el path para crear el archivo", 0, 0)
		return
	}
	//Escribir en Inodo de Archivo
	apuntadoresModificados, contenidoRestante := sb.EscribirEnBloquesArchivo(ctx, driveLetter, inodo.I_block[:], contenido)
	if apuntadoresModificados == nil || contenidoRestante == nil || ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo escribir en el inodo del archivo", 0, 0)
		return
	}
	//Warning si hay contenido restante
	if len(contenidoRestante) > 0 {
		//No agregar error, solo warning
		fmt.Println("Warning: No se pudo escribir todo el contenido en el archivo")
	}
	//Actualizar apuntadores en el inodo
	copy(inodo.I_block[:], apuntadoresModificados)
	//Actualizar size del inodo
	inodo.I_s = int64(len(contenido))
	//Actualizar el inodo en el disco
	sb.EscribirInodo(ctx, driveLetter, inodo, int64(numInodo))
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo actualizar el inodo del archivo", 0, 0)
		return
	}
}

// Escribir en Inodo de Archivo
// Retorna los apuntadores modificados del inodo y el contenido restante por escribir
func (sb *Super_Bloque) EscribirEnBloquesArchivo(ctx *Contexto, driveLetter string, apuntadores []int32, contenido []byte) ([]int32, []byte) {
	//Contenido restante por escribir
	var contenidoRestante []byte = contenido
	//Iterar todos los apuntadores
	for pos, numBloque := range apuntadores {
		if pos < 12 {
			numBloque, contenidoRestante = sb.EscribirEnBloqueArchivo(ctx, driveLetter, numBloque, contenidoRestante)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return nil, nil
			}
			apuntadores[pos] = numBloque
		} else if pos == 12 {
			numBloque, contenidoRestante = sb.EscribirEnBloqueIndirectoSimpleArchivo(ctx, driveLetter, numBloque, contenidoRestante)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return nil, nil
			}
			apuntadores[pos] = numBloque
		} else if pos == 13 {
			numBloque, contenidoRestante = sb.EscribirEnBloqueIndirectoDobleArchivo(ctx, driveLetter, numBloque, contenidoRestante)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return nil, nil
			}
			apuntadores[pos] = numBloque
		} else if pos == 14 {
			numBloque, contenidoRestante = sb.EscribirEnBloqueIndirectoTripleArchivo(ctx, driveLetter, numBloque, contenidoRestante)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return nil, nil
			}
			apuntadores[pos] = numBloque
		}
	}
	return apuntadores, contenidoRestante
}

func (sb *Super_Bloque) EscribirEnBloqueArchivo(ctx *Contexto, driveletter string, numBloque int32, contenido []byte) (int32, []byte) {
	if len(contenido) == 0 {
		//Liberar el bloque
		if numBloque != -1 {
			sb.ModificarBitEnBitMap(ctx, driveletter, numBloque, false, false)
		}
		return -1, make([]byte, 0)
	}
	var bloque *structs.Block_Archivo
	if numBloque != -1 {
		bloque = sb.RecuperarBloqueArchivo(ctx, driveletter, int(numBloque))
		if ctx.HayErrores() {
			return -1, nil
		}
	} else {
		//Obtener siguiente bloque libre
		numBloque = sb.SiguienteBitLibre(ctx, driveletter, false)
		if ctx.HayErrores() {
			return -1, nil
		}
		//Marcar el bloque como ocupado
		sb.ModificarBitEnBitMap(ctx, driveletter, numBloque, false, true)
		if ctx.HayErrores() {
			return -1, nil
		}
		//Crear nuevo bloque
		bloque = structs.NewBlock_Archivo()
	}
	//Calcular cantidad de contenido a escribir
	cantidadContenido := len(contenido)
	if cantidadContenido > len(bloque.B_content) {
		cantidadContenido = len(bloque.B_content)
	}
	//Limpiar contenido del bloque
	bloque.B_content = [64]byte{}
	//Escribir contenido en el bloque
	copy(bloque.B_content[:cantidadContenido], contenido[:cantidadContenido])
	//Actualizar el contenido restante
	contenidoRestante := contenido[cantidadContenido:]
	//Escribir bloque en el disco
	sb.EscribirBloque(ctx, driveletter, bloque, nil, nil, int64(numBloque))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return -1, nil
	}
	return numBloque, contenidoRestante
}

// Escribir en Bloque indirecto Simple Archivo
func (sb *Super_Bloque) EscribirEnBloqueIndirectoSimpleArchivo(ctx *Contexto, driveletter string, numBloque int32, contenido []byte) (int32, []byte) {
	//Contenido restante por escribir
	var contenidoRestante []byte = contenido
	//Recuperar bloque indirecto
	var bloque *structs.Block_Indirecto
	if numBloque != -1 {
		bloque = sb.RecuperarBloqueIndirecto(ctx, driveletter, int(numBloque))
		if ctx.HayErrores() {
			return -1, nil
		}
	} else if numBloque == -1 && len(contenidoRestante) > 0 {
		//Obtener siguiente bloque libre
		numBloque = sb.SiguienteBitLibre(ctx, driveletter, false)
		if ctx.HayErrores() {
			return -1, nil
		}
		//Marcar el bloque como ocupado
		sb.ModificarBitEnBitMap(ctx, driveletter, numBloque, false, true)
		if ctx.HayErrores() {
			return -1, nil
		}
		//Crear nuevo bloque
		bloque = structs.NewBlock_Indirecto()
	} else {
		return -1, make([]byte, 0)
	}
	if len(contenidoRestante) == 0 {
		//Liberar el bloque
		if numBloque != -1 {
			sb.ModificarBitEnBitMap(ctx, driveletter, numBloque, false, false)
		}
		numBloque = -1
	}
	//Iterar todos los apuntadores
	for pos, numBloque := range bloque.B_pointers {
		var numBloque_ int32
		numBloque_, contenidoRestante = sb.EscribirEnBloqueArchivo(ctx, driveletter, numBloque, contenidoRestante)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return -1, nil
		}
		numBloque = int32(numBloque_)
		bloque.B_pointers[pos] = numBloque
	}
	//Escribir bloque indirecto en el disco
	sb.EscribirBloque(ctx, driveletter, nil, nil, bloque, int64(numBloque))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return -1, nil
	}
	return numBloque, contenidoRestante
}

// Escribir en Bloque indirecto Doble Archivo
func (sb *Super_Bloque) EscribirEnBloqueIndirectoDobleArchivo(ctx *Contexto, driveletter string, numBloque int32, contenido []byte) (int32, []byte) {
	//Contenido restante por escribir
	var contenidoRestante []byte = contenido
	//Recuperar bloque indirecto
	var bloque *structs.Block_Indirecto
	if numBloque != -1 {
		bloque = sb.RecuperarBloqueIndirecto(ctx, driveletter, int(numBloque))
		if ctx.HayErrores() {
			return -1, nil
		}
	} else if numBloque == -1 && len(contenidoRestante) > 0 {
		//Obtener siguiente bloque libre
		numBloque = sb.SiguienteBitLibre(ctx, driveletter, false)
		if ctx.HayErrores() {
			return -1, nil
		}
		//Marcar el bloque como ocupado
		sb.ModificarBitEnBitMap(ctx, driveletter, numBloque, false, true)
		if ctx.HayErrores() {
			return -1, nil
		}
		//Crear nuevo bloque
		bloque = structs.NewBlock_Indirecto()
	} else {
		return -1, make([]byte, 0)
	}
	if len(contenidoRestante) == 0 {
		//Liberar el bloque
		if numBloque != -1 {
			sb.ModificarBitEnBitMap(ctx, driveletter, numBloque, false, false)
		}
		numBloque = -1
	}
	//Iterar todos los apuntadores
	for pos, numBloque := range bloque.B_pointers {
		var numBloque_ int32
		numBloque_, contenidoRestante = sb.EscribirEnBloqueIndirectoSimpleArchivo(ctx, driveletter, numBloque, contenidoRestante)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return -1, nil
		}
		numBloque = int32(numBloque_)
		bloque.B_pointers[pos] = numBloque
	}
	//Escribir bloque indirecto en el disco
	sb.EscribirBloque(ctx, driveletter, nil, nil, bloque, int64(numBloque))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return -1, nil
	}
	return numBloque, contenidoRestante
}

// Escribir en Bloque indirecto Triple Archivo
func (sb *Super_Bloque) EscribirEnBloqueIndirectoTripleArchivo(ctx *Contexto, driveletter string, numBloque int32, contenido []byte) (int32, []byte) {
	//Contenido restante por escribir
	var contenidoRestante []byte = contenido
	//Recuperar bloque indirecto
	var bloque *structs.Block_Indirecto
	if numBloque != -1 {
		bloque = sb.RecuperarBloqueIndirecto(ctx, driveletter, int(numBloque))
		if ctx.HayErrores() {
			return -1, nil
		}
	} else if numBloque == -1 && len(contenidoRestante) > 0 {
		//Obtener siguiente bloque libre
		numBloque = sb.SiguienteBitLibre(ctx, driveletter, false)
		if ctx.HayErrores() {
			return -1, nil
		}
		//Marcar el bloque como ocupado
		sb.ModificarBitEnBitMap(ctx, driveletter, numBloque, false, true)
		if ctx.HayErrores() {
			return -1, nil
		}
		//Crear nuevo bloque
		bloque = structs.NewBlock_Indirecto()
	} else {
		return -1, make([]byte, 0)
	}
	if len(contenidoRestante) == 0 {
		//Liberar el bloque
		if numBloque != -1 {
			sb.ModificarBitEnBitMap(ctx, driveletter, numBloque, false, false)
		}
		numBloque = -1
	}
	//Iterar todos los apuntadores
	for pos, numBloque := range bloque.B_pointers {
		var numBloque_ int32
		numBloque_, contenidoRestante = sb.EscribirEnBloqueIndirectoDobleArchivo(ctx, driveletter, numBloque, contenidoRestante)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return -1, nil
		}
		numBloque = int32(numBloque_)
		bloque.B_pointers[pos] = numBloque
	}
	//Escribir bloque indirecto en el disco
	sb.EscribirBloque(ctx, driveletter, nil, nil, bloque, int64(numBloque))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return -1, nil
	}
	return numBloque, contenidoRestante
}

// Buscar Bloques Contiguos
func (sb *Super_Bloque) BuscarBloquesContiguos(ctx *Contexto, driveLetter string, bloquesNecesarios int, fit string) (bool, []int64) {
	//Recuperar bitmap de bloques
	bitmap := sb.Bitmap(ctx, driveLetter, false)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return false, nil
	}
	//Agregar un bit 1 al final del bitmap para funcionar como centinela
	bitmap = append(bitmap, 1)
	//Contador de bloques contiguos
	var bloquesContados int = -1
	var auxBloquesContados int = 0
	numsBlocksContiguos := make([]int64, 0)
	auxNumsBlocksContiguos := make([]int64, 0)
	hayBloquesContiguos := false
	//Iterar el bitmap
	for pos, bit := range bitmap {
		if bit == 0 {
			auxBloquesContados++
			auxNumsBlocksContiguos = append(auxNumsBlocksContiguos, int64(pos))
			continue
		}
		if bloquesNecesarios > auxBloquesContados {
			auxBloquesContados = 0
			auxNumsBlocksContiguos = make([]int64, 0)
			continue
		}
		if fit == "F" {
			return true, auxNumsBlocksContiguos
		} else if fit == "B" {
			if bloquesContados == -1 || bloquesContados > auxBloquesContados {
				bloquesContados = auxBloquesContados
				numsBlocksContiguos = auxNumsBlocksContiguos
				hayBloquesContiguos = true
			}
		} else if fit == "W" {
			if bloquesContados == -1 || bloquesContados < auxBloquesContados {
				bloquesContados = auxBloquesContados
				numsBlocksContiguos = auxNumsBlocksContiguos
				hayBloquesContiguos = true
			}
		}
		auxBloquesContados = 0
		auxNumsBlocksContiguos = make([]int64, 0)
	}
	return hayBloquesContiguos, numsBlocksContiguos
}

// Contar Bloques de Archivos usados por un inodo
func (sb *Super_Bloque) ContarBloquesUsados(ctx *Contexto, driveletter string, inodo *structs.Inodes) int {
	//Contador de bloques
	var contador int = 0
	//Iterar todos los bloques
	for pos, numBloque := range inodo.I_block {
		if numBloque == -1 {
			continue
		}
		if pos < 12 {
			contador++
		} else if pos == 12 {
			contador += sb.ContarBloquesUsadosIndirectoSimple(ctx, driveletter, numBloque)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return -1
			}
		} else if pos == 13 {
			contador += sb.ContarBloquesUsadosIndirectoDoble(ctx, driveletter, numBloque)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return -1
			}
		} else if pos == 14 {
			contador += sb.ContarBloquesUsadosIndirectoTriple(ctx, driveletter, numBloque)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return -1
			}
		}
	}
	return contador
}

// Contar Bloques de Archivos usados por un inodo con indirecto simple
func (sb *Super_Bloque) ContarBloquesUsadosIndirectoSimple(ctx *Contexto, driveLetter string, numBloque int32) int {
	//Contador de bloques
	var contador int = 0
	//Recuperar bloque indirecto simple
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, int(numBloque))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return -1
	}
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		if numBloque == -1 {
			continue
		}
		contador++
	}
	return contador
}

// Contar Bloques de Archivos usados por un inodo con indirecto doble
func (sb *Super_Bloque) ContarBloquesUsadosIndirectoDoble(ctx *Contexto, driveLetter string, numBloque int32) int {
	//Contador de bloques
	var contador int = 0
	//Recuperar bloque indirecto doble
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, int(numBloque))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return -1
	}
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		if numBloque == -1 {
			continue
		}
		contador += sb.ContarBloquesUsadosIndirectoSimple(ctx, driveLetter, numBloque)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return -1
		}
	}
	return contador
}

// Contar Bloques de Archivos usados por un inodo con indirecto triple
func (sb *Super_Bloque) ContarBloquesUsadosIndirectoTriple(ctx *Contexto, driveLetter string, numBloque int32) int {
	//Contador de bloques
	var contador int = 0
	//Recuperar bloque indirecto triple
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, int(numBloque))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return -1
	}
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		if numBloque == -1 {
			continue
		}
		contador += sb.ContarBloquesUsadosIndirectoDoble(ctx, driveLetter, numBloque)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return -1
		}
	}
	return contador
}

// Generar GID para un grupo
func (sb *Super_Bloque) GenerarGID(grupos []GrupoParticion) int {
	//Iterar todos los grupos para encontrar el GID mas alto
	var maxGID int = 0
	for _, grupo := range grupos {
		gid, _ := strconv.Atoi(grupo.GID)
		if gid > maxGID {
			maxGID = gid
		}
	}
	return maxGID + 1
}

// Generar UID para un usuario
func (sb *Super_Bloque) GenerarUID(usuarios []UsuarioParticion) int {
	//Iterar todos los usuarios para encontrar el UID mas alto
	var maxUID int = 0
	for _, usuario := range usuarios {
		uid, _ := strconv.Atoi(usuario.UID)
		if uid > maxUID {
			maxUID = uid
		}
	}
	return maxUID + 1
}

// Leer Bloques de Archivos
func (sb *Super_Bloque) LeerBloquesArchivo(ctx *Contexto, driveLetter string, apuntadores []int32) []byte {
	//Contenido de los bloques
	contenido := make([]byte, 0)
	//Iterar todos los apuntadores
	for pos, numBloque := range apuntadores {
		//Verificar si el numero de bloque es -1
		if numBloque == -1 {
			continue
		}
		contAux := make([]byte, 0)
		if pos < 12 {
			//Recuperar bloque
			bloque := sb.RecuperarBloqueArchivo(ctx, driveLetter, int(numBloque))
			contAux = bloque.B_content[:]
		} else if pos == 12 {
			//Recuperar bloque indirecto simple
			contAux = sb.LeerBloqueIndirectoSimple(ctx, driveLetter, numBloque)
		} else if pos == 13 {
			//Recuperar bloque indirecto doble
			contAux = sb.LeerBloqueIndirectoDoble(ctx, driveLetter, numBloque)
		} else if pos == 14 {
			//Recuperar bloque indirecto triple
			contAux = sb.LeerBloqueIndirectoTriple(ctx, driveLetter, numBloque)
		}
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil
		}
		//Agregar el contenido del bloque al contenido total
		contenido = append(contenido, contAux...)
	}
	//Eliminar los bytes nulos
	contenido = bytes.Replace(contenido, []byte{0}, []byte{}, -1)
	return contenido
}

// Leer Bloque indirecto simple de un inodo
func (sb *Super_Bloque) LeerBloqueIndirectoSimple(ctx *Contexto, driveLetter string, numBloque int32) []byte {
	cotenido := make([]byte, 0)
	//Recuperar bloque indirecto simple
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, int(numBloque))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		//Verificar si el numero de bloque es -1
		if numBloque == -1 {
			continue
		}
		//Recuperar bloque
		bloque := sb.RecuperarBloqueArchivo(ctx, driveLetter, int(numBloque))
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil
		}
		//Agregar el contenido del bloque al contenido total
		cotenido = append(cotenido, bloque.B_content[:]...)
	}
	//Bytes leidos
	return cotenido
}

// Leer Bloque indirecto doble de un inodo
func (sb *Super_Bloque) LeerBloqueIndirectoDoble(ctx *Contexto, driveLetter string, numBloque int32) []byte {
	contenido := make([]byte, 0)
	//Recuperar bloque indirecto doble
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, int(numBloque))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		//Verificar si el numero de bloque es -1
		if numBloque == -1 {
			continue
		}
		//Recuperar bloque indirecto simple
		bloqueSimple := sb.LeerBloqueIndirectoSimple(ctx, driveLetter, numBloque)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil
		}
		//Agregar el contenido del bloque al contenido total
		contenido = append(contenido, bloqueSimple...)
	}
	//Bytes leidos
	return contenido
}

// Leer Bloque indirecto triple de un inodo
func (sb *Super_Bloque) LeerBloqueIndirectoTriple(ctx *Contexto, driveLetter string, numBloque int32) []byte {
	contenido := make([]byte, 0)
	//Recuperar bloque indirecto triple
	bloque := sb.RecuperarBloqueIndirecto(ctx, driveLetter, int(numBloque))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Iterar todos los apuntadores
	for _, numBloque := range bloque.B_pointers {
		//Verificar si el numero de bloque es -1
		if numBloque == -1 {
			continue
		}
		//Recuperar bloque indirecto doble
		bloqueDoble := sb.LeerBloqueIndirectoDoble(ctx, driveLetter, numBloque)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil
		}
		//Agregar el contenido del bloque al contenido total
		contenido = append(contenido, bloqueDoble...)
	}
	//Bytes leidos
	return contenido
}

// Buscar Usuario y grupo por UID y GID
func (sb *Super_Bloque) BuscarUsuarioGrupoPorID(ctx *Contexto, driveLetter string, UID int64, GID int64) (*UsuarioParticion, *GrupoParticion) {
	//Recuperar Usuarios y Grupos
	usuarios, grupos, err := sb.RecuperarUsuariosGrupos(ctx, driveLetter)
	if err != nil {
		ctx.AgregarError("Error: No se pudieron recuperar los usuarios y grupos", 0, 0)
		return nil, nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil, nil
	}
	//Buscar Usuario
	var usuario *UsuarioParticion = nil
	for _, user := range usuarios {
		uid, _ := strconv.Atoi(user.UID)
		if int64(uid) == UID {
			usuario = &user
			break
		}
	}
	//Buscar Grupo
	var grupo *GrupoParticion = nil
	for _, gr := range grupos {
		gid, _ := strconv.Atoi(gr.GID)
		if int64(gid) == GID {
			grupo = &gr
			break
		}
	}
	return usuario, grupo
}

// Recuperar todos los usuarios y grupos de la particion
func (sb *Super_Bloque) RecuperarUsuariosGrupos(ctx *Contexto, driveletter string) ([]UsuarioParticion, []GrupoParticion, error) {
	//Recuperar inodos de usuarios y grupos
	inodoUsers := sb.RecuperarInodo(ctx, driveletter, 1)
	if inodoUsers == nil {
		return nil, nil, nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil, nil, nil
	}
	//Leer todos los bloques del archivo users.txt
	contenidoUsers := sb.LeerBloquesArchivo(ctx, driveletter, inodoUsers.I_block[:])
	//Eliminar los bytes nulos
	contenidoUsers = bytes.Replace(contenidoUsers, []byte{0}, []byte{}, -1)
	//Split por salto de linea
	users := bytes.Split(contenidoUsers, []byte("\n"))
	//Eliminar el ultimo elemento que es un string vacio
	users = users[:len(users)-1]
	//Iterar el split para obtener los usuarios y grupos
	var usuarios []UsuarioParticion
	var grupos []GrupoParticion
	for _, user := range users {
		//Split por coma
		userData := bytes.Split(user, []byte(","))
		if len(userData) == 5 {
			usuarios = append(usuarios, UsuarioParticion{
				UID:     string(userData[0]),
				Tipo:    string(userData[1]),
				Grupo:   string(userData[2]),
				Usuario: string(userData[3]),
				Pass:    string(userData[4]),
			})
		} else if len(userData) == 3 {
			grupos = append(grupos, GrupoParticion{
				GID:   string(userData[0]),
				Tipo:  string(userData[1]),
				Grupo: string(userData[2]),
			})
		}

	}
	return usuarios, grupos, nil
}

// Generar Inodo Raiz
func (sb *Super_Bloque) CrearCarpetaRaiz(ctx *Contexto, driveletter string) {
	//Primer bloque libre
	numeroBloque := sb.S_first_blo
	if numeroBloque == -1 {
		ctx.AgregarError("Error: No hay bloques libres", 0, 0)
		return
	}
	//Crear Bloque
	bloque := structs.NewBlock_Carpeta()
	copy(bloque.B_content[0].B_name[:], ".")
	bloque.B_content[0].B_inodo = 0
	copy(bloque.B_content[1].B_name[:], "..")
	bloque.B_content[1].B_inodo = 0
	//Colocar Inodo del archivo de usuarios users.txt en el bloque
	copy(bloque.B_content[2].B_name[:], "users.txt")
	bloque.B_content[2].B_inodo = 1
	//Primer inodo libre
	numeroInodo := sb.S_first_ino
	if numeroInodo == -1 {
		ctx.AgregarError("Error: No hay inodos libres", 0, 0)
		return
	}
	//Crear Inodo
	inodo := structs.NewInode()
	inodo.I_uid = 1
	inodo.I_gid = 1
	inodo.I_s = 0
	copy(inodo.I_atime[:], time.Now().Format("2006-01-02 15:04:05"))
	copy(inodo.I_ctime[:], time.Now().Format("2006-01-02 15:04:05"))
	copy(inodo.I_mtime[:], time.Now().Format("2006-01-02 15:04:05"))
	inodo.I_type = 0
	inodo.I_perm = [3]byte{7, 7, 7}
	//Asignar Bloque al Inodo
	inodo.I_block[0] = int32(numeroBloque)
	//Marcar Inodo y Bloque en el Bitmap
	sb.ModificarBitEnBitMap(ctx, driveletter, int32(numeroInodo), true, true)
	sb.ModificarBitEnBitMap(ctx, driveletter, int32(numeroBloque), false, true)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return
	}
	//Escribir Inodo y Bloque en el Disco
	sb.EscribirInodo(ctx, driveletter, inodo, numeroInodo)
	sb.EscribirBloque(ctx, driveletter, nil, bloque, nil, numeroBloque)
}

// Crear Archivo users.txt
func (sb *Super_Bloque) CrearArchivoUsers(ctx *Contexto, driveLetter string) {
	//Primer bloque libre
	numeroBloque := sb.S_first_blo
	if numeroBloque == -1 {
		ctx.AgregarError("Error: No hay bloques libres", 0, 0)
		return
	}
	//Crear Bloque
	bloque := structs.NewBlock_Archivo()
	//Contenido del archivo
	contenido := "1,G,root\n1,U,root,root,123\n"
	copy(bloque.B_content[:], contenido)
	//Primer inodo libre
	numeroInodo := sb.S_first_ino
	if numeroInodo == -1 {
		ctx.AgregarError("Error: No hay inodos libres", 0, 0)
		return
	}
	//Crear Inodo
	inodo := structs.NewInode()
	inodo.I_uid = 1
	inodo.I_gid = 1
	inodo.I_s = int64(len(contenido))
	copy(inodo.I_atime[:], time.Now().Format("2006-01-02 15:04:05"))
	copy(inodo.I_ctime[:], time.Now().Format("2006-01-02 15:04:05"))
	copy(inodo.I_mtime[:], time.Now().Format("2006-01-02 15:04:05"))
	inodo.I_type = 1
	inodo.I_perm = [3]byte{7, 7, 4}
	//Asignar Bloque al Inodo
	inodo.I_block[0] = int32(numeroBloque)
	//Marcar Inodo y Bloque en el Bitmap
	sb.ModificarBitEnBitMap(ctx, driveLetter, int32(numeroInodo), true, true)
	sb.ModificarBitEnBitMap(ctx, driveLetter, int32(numeroBloque), false, true)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return
	}
	//Escribir Inodo y Bloque en el Disco
	sb.EscribirInodo(ctx, driveLetter, inodo, numeroInodo)
	sb.EscribirBloque(ctx, driveLetter, bloque, nil, nil, numeroBloque)
}

// Marcar o desmarcar un inodo o bloque en el bitmap
// numero = Numero de inodo o bloque a marcar
// tipo = true -> Inodo | tipo = false -> Bloque
// marcar = true -> Marcar | marcar = false -> Desmarcar
func (sb *Super_Bloque) ModificarBitEnBitMap(ctx *Contexto, driveLetter string, numero int32, tipo bool, marcar bool) {
	var inicio_bitmap int64
	if tipo {
		inicio_bitmap = sb.S_bm_inode_start
	} else {
		inicio_bitmap = sb.S_bm_block_start
	}
	//Abrir el archivo del disco
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_WRONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error al abrir el disco al Marcar en el Bitmap: "+err.Error(), 0, 0)
		return
	}
	defer file.Close()
	//Mover el puntero al inicio del bitmap
	file.Seek(inicio_bitmap+int64(numero), io.SeekStart)
	var bit byte
	if marcar {
		bit = 1
	} else {
		bit = 0
	}
	//Escribir el bit
	err = binary.Write(file, binary.BigEndian, bit)
	if err != nil {
		ctx.AgregarError("Error al Marcar en el Bitmap: "+err.Error(), 0, 0)
		return
	}
	//Modificar el Super Bloque
	if tipo {
		if marcar {
			sb.S_free_inodes_count--
		} else {
			sb.S_free_inodes_count++
		}
		sb.S_first_ino = int64(sb.SiguienteBitLibre(ctx, driveLetter, true))
	} else {
		if marcar {
			sb.S_free_blocks_count--
		} else {
			sb.S_free_blocks_count++
		}
		sb.S_first_blo = int64(sb.SiguienteBitLibre(ctx, driveLetter, false))
	}
}

// Obtener el siguiente inodo o bloque libre
// tipo = true -> Inodo | tipo = false -> Bloque
func (sb *Super_Bloque) SiguienteBitLibre(ctx *Contexto, driveLetter string, tipo bool) int32 {
	var inicio_bitmap int64
	if tipo {
		inicio_bitmap = sb.S_bm_inode_start
	} else {
		inicio_bitmap = sb.S_bm_block_start
	}
	//Abrir el archivo del disco
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_RDONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error al abrir el disco al obtener el siguiente bit libre: "+err.Error(), 0, 0)
		return -1
	}
	defer file.Close()
	//Mover el puntero al inicio del bitmap
	file.Seek(inicio_bitmap, io.SeekStart)
	//Leer el bitmap
	var sizeBitmap int64
	if tipo {
		sizeBitmap = sb.S_inodes_count
	} else {
		sizeBitmap = sb.S_blocks_count
	}
	bitMap := make([]byte, sizeBitmap)
	err = binary.Read(file, binary.BigEndian, &bitMap)
	if err != nil {
		ctx.AgregarError("Error al leer el bitmap: "+err.Error(), 0, 0)
		return -1
	}
	//Buscar el siguiente bit libre
	for pos, bit := range bitMap {
		if bit == 0 {
			return int32(pos)
		}
	}
	//No hay bits libres
	ctx.AgregarError("Error: No hay bits libres en el bitmap", 0, 0)
	return -1
}

// Escribir inodo en el disco
func (sb *Super_Bloque) EscribirInodo(ctx *Contexto, driveLetter string, inodo *structs.Inodes, numero int64) {
	//Abrir el archivo del disco
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_WRONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error al abrir el disco al escribir el inodo: "+err.Error(), 0, 0)
		return
	}
	defer file.Close()
	//Mover el puntero al inicio del inodo
	file.Seek(sb.S_inode_start+(numero*sb.S_inode_s), io.SeekStart)
	//Escribir el inodo
	err = binary.Write(file, binary.BigEndian, inodo)
	if err != nil {
		ctx.AgregarError("Error al escribir el inodo: "+err.Error(), 0, 0)
		return
	}
}

// Escribir bloque en el disco
func (sb *Super_Bloque) EscribirBloque(ctx *Contexto, driveLetter string, bloqueArchivo *structs.Block_Archivo, bloqueCarpeta *structs.Block_Carpeta, bloqueIndirecto *structs.Block_Indirecto, numero int64) {
	//Abrir el archivo del disco
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_WRONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error al abrir el disco al escribir el bloque: "+err.Error(), 0, 0)
		return
	}
	defer file.Close()
	//Mover el puntero al inicio del bloque
	file.Seek(sb.S_block_start+(numero*sb.S_block_s), io.SeekStart)
	//Escribir el bloque segun su tipo
	if bloqueArchivo != nil {
		err = binary.Write(file, binary.BigEndian, bloqueArchivo)
	} else if bloqueCarpeta != nil {
		err = binary.Write(file, binary.BigEndian, bloqueCarpeta)
	} else if bloqueIndirecto != nil {
		err = binary.Write(file, binary.BigEndian, bloqueIndirecto)
	}
	if err != nil {
		ctx.AgregarError("Error al escribir el bloque: "+err.Error(), 0, 0)
		return
	}
}

// Escribir Super Bloque en el disco
func (sb *Super_Bloque) EscribirSuperBloque(ctx *Contexto, driveLetter string, start_part int64) {
	//Abrir el archivo del disco
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_WRONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error al abrir el disco al escribir el Super Bloque: "+err.Error(), 0, 0)
		return
	}
	defer file.Close()
	//Mover el puntero al inicio del Super Bloque
	file.Seek(start_part, io.SeekStart)
	//Escribir el Super Bloque
	err = binary.Write(file, binary.BigEndian, sb)
	if err != nil {
		ctx.AgregarError("Error al escribir el Super Bloque: "+err.Error(), 0, 0)
		return
	}
}

// Recuperar Inodo
func (sb *Super_Bloque) RecuperarInodo(ctx *Contexto, driveLetter string, numero int) *structs.Inodes {
	//Abrir el archivo del disco
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_RDONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error al abrir el disco al recuperar el inodo: "+err.Error(), 0, 0)
		return nil
	}
	defer file.Close()
	//Mover el puntero al inicio del inodo
	file.Seek(sb.S_inode_start+(int64(numero)*sb.S_inode_s), io.SeekStart)
	//Leer el inodo
	inodo := &structs.Inodes{}
	err = binary.Read(file, binary.BigEndian, inodo)
	if err != nil {
		ctx.AgregarError("Error al leer el inodo: "+err.Error(), 0, 0)
		return nil
	}
	return inodo
}

// Insertar Inodo en Inodo
func (sb *Super_Bloque) InsertarInodoEnInodo(ctx *Contexto, driveLetter string, inodo *structs.Inodes, inodoPadre *structs.Inodes, nombre string) *structs.Inodes {
	//Iterar los apuntadores para insertar el inodo
	return nil
}

// Recuperar Inodo por Path
func (sb *Super_Bloque) RecuperarInodoPorPath(ctx *Contexto, driveLetter string, path []string, numInodo int, tipo int, r bool, soloLectura bool, delete bool) (*structs.Inodes, int, int, int) {
	//Recuperar inodo
	inodo := sb.RecuperarInodo(ctx, driveLetter, numInodo)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil, -1, -1, -1
	}
	//Leer Bloque Carpeta por Path
	inode, numInode, block, numBlock := sb.LeerBloqueCarpetaPorPath(ctx, driveLetter, path, inodo.I_block[:], tipo, r, soloLectura, int32(numInodo), delete)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil, -1, -1, -1
	}
	//Actualizar el "inodo" en disco
	sb.EscribirInodo(ctx, driveLetter, inodo, int64(numInodo))
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil, -1, -1, -1
	}
	return inode, numInode, block, numBlock
}

// Leer Bloque Carpeta por Path
func (sb *Super_Bloque) LeerBloqueCarpetaPorPath(ctx *Contexto, driveLetter string, path []string, apuntadores []int32, tipo int, r bool, soloLectura bool, padre int32, delete bool) (*structs.Inodes, int, int, int) {
	//Iterar sobre los bloques del inodo
	//For de Lectura
	for pos, numBloque := range apuntadores {
		//Leer Bloque indirecto simple
		if pos == 12 && len(apuntadores) == 15 {
			inodo, numInodo, block, ind := sb.LeerBloqueIndirectoSimplePorPath(ctx, driveLetter, path, apuntadores, pos, tipo, r, true, padre, delete)
			if inodo != nil {
				return inodo, numInodo, block, ind
			}
			continue
		} else if pos == 13 && len(apuntadores) == 15 {
			inodo, numInodo, block, ind := sb.LeerBloqueIndirectoDoblePorPath(ctx, driveLetter, path, apuntadores, pos, tipo, r, true, padre, delete)
			if inodo != nil {
				return inodo, numInodo, block, ind
			}
			continue
		} else if pos == 14 && len(apuntadores) == 15 {
			inodo, numInodo, block, ind := sb.LeerBloqueIndirectoTriplePorPath(ctx, driveLetter, path, apuntadores, pos, tipo, r, true, padre, delete)
			if inodo != nil {
				return inodo, numInodo, block, ind
			}
			continue
		}
		//Verificar si el numero de bloque es -1
		if numBloque == -1 {
			continue
		}
		//Recuperar bloque
		bloque := sb.RecuperarBloqueCarpeta(ctx, driveLetter, int(numBloque))
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil, -1, -1, -1
		}
		//Iterar sobre el contenido del bloque
		for ind, contenido := range bloque.B_content {
			//Si es un inodo Carpeta, no leer los primeros 2 registros del primer apuntador directo
			/*
				if len(apuntadores) == 15 && pos == 0 && ind == 1 {
					padre = contenido.B_inodo
				}
			*/
			if pos == 0 && ind < 2 && len(apuntadores) == 15 {
				continue
			}
			if contenido.B_inodo == -1 {
				continue
			}
			//Verificar si el nombre del contenido es igual al nombre del path
			//Quitar valores nulos del nombre y convertir a string
			nombre := string(bytes.Replace(contenido.B_name[:], []byte{0}, []byte{}, -1))
			if nombre == path[0] {
				//Si el path tiene un solo elemento, retornar el inodo
				if len(path) == 1 {
					numInodo := int(contenido.B_inodo)
					if delete {
						sb.ModificarBitEnBitMap(ctx, driveLetter, int32(contenido.B_inodo), true, false)
						//Verificar si hay errores
						if ctx.HayErrores() {
							return nil, -1, -1, -1
						}
						bloque.B_content[ind].B_inodo = -1
						//Verificar si el bloque Carpeta ya no apunta a nada
						bloqueVacio := true
						for _, contenido := range bloque.B_content {
							if contenido.B_inodo != -1 {
								bloqueVacio = false
								break
							}
						}
						if bloqueVacio {
							sb.ModificarBitEnBitMap(ctx, driveLetter, numBloque, false, false)
							//Verificar si hay errores
							if ctx.HayErrores() {
								return nil, -1, -1, -1
							}
							//Actualizar el apuntador del bloque en el inodo
							apuntadores[pos] = -1
						}
						sb.EscribirBloque(ctx, driveLetter, nil, bloque, nil, int64(numBloque))
						//Verificar si hay errores
						if ctx.HayErrores() {
							return nil, -1, -1, -1
						}

					}
					return sb.RecuperarInodo(ctx, driveLetter, numInodo), numInodo, int(numBloque), ind
				}
				//Si el path tiene mas de un elemento, buscar el inodo en la carpeta
				return sb.RecuperarInodoPorPath(ctx, driveLetter, path[1:], int(contenido.B_inodo), tipo, r, soloLectura, delete)
			}
		}
	}
	//Verificar si es solo lectura
	if soloLectura {
		return nil, -1, -1, -1
	}
	//Si la flag r es falsa, no se crea el inodo
	if !r && len(path) > 1 {
		return nil, -1, -1, -1
	}
	//Si la flag r es verdadera, crear el inodo
	//Iterar sobre los bloques del inodo para encontrar un bloque libre
	for pos, numBloque := range apuntadores {
		if pos == 12 && len(apuntadores) == 15 {
			inodo, numInodo, block, ind := sb.LeerBloqueIndirectoSimplePorPath(ctx, driveLetter, path, apuntadores, pos, tipo, r, false, padre, delete)
			if inodo != nil {
				return inodo, numInodo, block, ind
			}
			continue
		}
		if pos == 13 && len(apuntadores) == 15 {
			inodo, numInodo, block, ind := sb.LeerBloqueIndirectoDoblePorPath(ctx, driveLetter, path, apuntadores, pos, tipo, r, false, padre, delete)
			if inodo != nil {
				return inodo, numInodo, block, ind
			}
			continue
		} else if pos == 14 && len(apuntadores) == 15 {
			inodo, numInodo, block, ind := sb.LeerBloqueIndirectoTriplePorPath(ctx, driveLetter, path, apuntadores, pos, tipo, r, false, padre, delete)
			if inodo != nil {
				return inodo, numInodo, block, ind
			}
			continue
		}
		var bloqueCarpeta *structs.Block_Carpeta
		if numBloque != -1 {
			bloqueCarpeta = sb.RecuperarBloqueCarpeta(ctx, driveLetter, int(numBloque))
		} else {
			//Obtener el siguiente bloque libre
			numBloque = sb.SiguienteBitLibre(ctx, driveLetter, false)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return nil, -1, -1, -1
			}
			//Marcar el bloque en el bitmap
			sb.ModificarBitEnBitMap(ctx, driveLetter, numBloque, false, true)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return nil, -1, -1, -1
			}
			//Crear nuevo bloque carpeta
			bloqueCarpeta = structs.NewBlock_Carpeta()
			//Asignar el bloque al inodo
			apuntadores[pos] = numBloque
		}
		//Iterar sobre el contenido del bloque
		for ind, contenido := range bloqueCarpeta.B_content {
			if contenido.B_inodo != -1 {
				continue
			}
			//Crear nuevo inodo
			tipoInodo := 0
			if len(path) == 1 {
				if tipo == 0 {
					tipoInodo = 0
				} else {
					tipoInodo = 1
				}
			}
			numInodo := sb.CrearInodo(ctx, driveLetter, padre, tipoInodo)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return nil, -1, -1, -1
			}
			//Asignar el inodo al contenido del bloque
			bloqueCarpeta.B_content[ind].B_inodo = int32(numInodo)
			//Asignar el nombre al contenido del bloque
			copy(bloqueCarpeta.B_content[ind].B_name[:], path[0])
			//Escribir el bloque en el disco
			sb.EscribirBloque(ctx, driveLetter, nil, bloqueCarpeta, nil, int64(numBloque))
			//Verificar si hay errores
			if ctx.HayErrores() {
				return nil, -1, -1, -1
			}
			//Si el path tiene un solo elemento, retornar el inodo
			if len(path) == 1 {
				//fmt.Println("Inodo creado: "+path[0], numInodo, numBloque, ind)
				return sb.RecuperarInodo(ctx, driveLetter, int(numInodo)), int(numInodo), int(numBloque), ind
			} else {
				//Si el path tiene mas de un elemento, buscar el inodo en la carpeta
				return sb.RecuperarInodoPorPath(ctx, driveLetter, path[1:], int(numInodo), tipo, r, soloLectura, delete)
			}
		}
	}
	//Si no se encontro un bloque libre, agregar un error
	ctx.AgregarError("Error: No hay bloques libres para crear : "+path[0], 0, 0)
	return nil, -1, -1, -1
}

// Leer Bloque Indirecto Simple por Path
func (sb *Super_Bloque) LeerBloqueIndirectoSimplePorPath(ctx *Contexto, driveLetter string, path []string, apuntadores []int32, pos int, tipo int, r bool, soloLectura bool, padre int32, delete bool) (*structs.Inodes, int, int, int) {
	//Verificar si es solo lectura
	if soloLectura && apuntadores[pos] == -1 {
		return nil, -1, -1, -1
	}
	//Recuperar bloque indirecto simple
	var bloque *structs.Block_Indirecto
	if apuntadores[pos] == -1 {
		//Crear nuevo bloque indirecto simple
		numBloque := sb.SiguienteBitLibre(ctx, driveLetter, false)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil, -1, -1, -1
		}
		//Marcar el bloque en el bitmap
		sb.ModificarBitEnBitMap(ctx, driveLetter, numBloque, false, true)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil, -1, -1, -1
		}
		//Crear nuevo bloque indirecto simple
		bloque = structs.NewBlock_Indirecto()
		//Asignar el bloque al inodo
		apuntadores[pos] = numBloque
	} else {
		bloque = sb.RecuperarBloqueIndirecto(ctx, driveLetter, int(apuntadores[pos]))
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil, -1, -1, -1
	}
	//Leer Bloque Carpeta por Path
	inodo, numInodo, block, ind := sb.LeerBloqueCarpetaPorPath(ctx, driveLetter, path, bloque.B_pointers[:], tipo, r, soloLectura, padre, delete)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil, -1, -1, -1
	}
	if delete {
		//Verificar si el bloque Indirecto ya no apunta a nada
		bloqueVacio := true
		for _, contenido := range bloque.B_pointers {
			if contenido != -1 {
				bloqueVacio = false
				break
			}
		}
		if bloqueVacio {
			sb.ModificarBitEnBitMap(ctx, driveLetter, int32(apuntadores[pos]), false, false)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return nil, -1, -1, -1
			}
			//Actualizar el apuntador del bloque en el inodo
			apuntadores[pos] = -1
		} else {
			sb.EscribirBloque(ctx, driveLetter, nil, nil, bloque, int64(apuntadores[pos]))
			//Verificar si hay errores
			if ctx.HayErrores() {
				return nil, -1, -1, -1
			}
		}
	} else {
		//Escribir el bloque en el disco
		sb.EscribirBloque(ctx, driveLetter, nil, nil, bloque, int64(apuntadores[pos]))
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil, -1, -1, -1
		}
	}
	return inodo, numInodo, block, ind
}

// Leer Bloque Indirecto Doble por Path
func (sb *Super_Bloque) LeerBloqueIndirectoDoblePorPath(ctx *Contexto, driveLetter string, path []string, apuntadores []int32, pos int, tipo int, r bool, soloLectura bool, padre int32, delete bool) (*structs.Inodes, int, int, int) {
	//Verificar si es solo lectura
	if soloLectura && apuntadores[pos] == -1 {
		return nil, -1, -1, -1
	}
	//Recuperar bloque indirecto doble
	var bloque *structs.Block_Indirecto
	if apuntadores[pos] == -1 {
		//Crear nuevo bloque indirecto doble
		numBloque := sb.SiguienteBitLibre(ctx, driveLetter, false)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil, -1, -1, -1
		}
		//Marcar el bloque en el bitmap
		sb.ModificarBitEnBitMap(ctx, driveLetter, numBloque, false, true)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil, -1, -1, -1
		}
		//Crear nuevo bloque indirecto doble
		bloque = structs.NewBlock_Indirecto()
		//Asignar el bloque al inodo
		apuntadores[pos] = numBloque
	} else {
		bloque = sb.RecuperarBloqueIndirecto(ctx, driveLetter, int(apuntadores[pos]))
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil, -1, -1, -1
	}
	//Iterar todos los apuntadores
	for pos := range bloque.B_pointers {
		//Leer Bloque Indirecto Simple por Path
		inodo, numInodo, block, ind := sb.LeerBloqueIndirectoSimplePorPath(ctx, driveLetter, path, bloque.B_pointers[:], pos, tipo, r, soloLectura, padre, delete)
		if ctx.HayErrores() {
			return nil, -1, -1, -1
		}
		if inodo != nil {
			if delete {
				//Verificar si el bloque Indirecto ya no apunta a nada
				bloqueVacio := true
				for _, contenido := range bloque.B_pointers {
					if contenido != -1 {
						bloqueVacio = false
						break
					}
				}
				if bloqueVacio {
					sb.ModificarBitEnBitMap(ctx, driveLetter, int32(apuntadores[pos]), false, false)
					//Verificar si hay errores
					if ctx.HayErrores() {
						return nil, -1, -1, -1
					}
					//Actualizar el apuntador del bloque en el inodo
					apuntadores[pos] = -1
				} else {
					sb.EscribirBloque(ctx, driveLetter, nil, nil, bloque, int64(apuntadores[pos]))
					//Verificar si hay errores
					if ctx.HayErrores() {
						return nil, -1, -1, -1
					}
				}
			} else {
				//Escribir el bloque en el disco
				sb.EscribirBloque(ctx, driveLetter, nil, nil, bloque, int64(apuntadores[pos]))
				//Verificar si hay errores
				if ctx.HayErrores() {
					return nil, -1, -1, -1
				}
			}
			return inodo, numInodo, block, ind
		}
	}
	return nil, -1, -1, -1
}

// Leer Bloque Indirecto Triple por Path
func (sb *Super_Bloque) LeerBloqueIndirectoTriplePorPath(ctx *Contexto, driveLetter string, path []string, apuntadores []int32, pos int, tipo int, r bool, soloLectura bool, padre int32, delete bool) (*structs.Inodes, int, int, int) {
	//Verificar si es solo lectura
	if soloLectura && apuntadores[pos] == -1 {
		return nil, -1, -1, -1
	}
	//Recuperar bloque indirecto triple
	var bloque *structs.Block_Indirecto
	if apuntadores[pos] == -1 {
		//Crear nuevo bloque indirecto triple
		numBloque := sb.SiguienteBitLibre(ctx, driveLetter, false)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil, -1, -1, -1
		}
		//Marcar el bloque en el bitmap
		sb.ModificarBitEnBitMap(ctx, driveLetter, numBloque, false, true)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil, -1, -1, -1
		}
		//Crear nuevo bloque indirecto triple
		bloque = structs.NewBlock_Indirecto()
		//Asignar el bloque al inodo
		apuntadores[pos] = numBloque
	} else {
		bloque = sb.RecuperarBloqueIndirecto(ctx, driveLetter, int(apuntadores[pos]))
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil, -1, -1, -1
	}
	//Iterar todos los apuntadores
	for pos := range bloque.B_pointers {
		//Leer Bloque Indirecto Doble por Path
		inodo, numInodo, block, ind := sb.LeerBloqueIndirectoDoblePorPath(ctx, driveLetter, path, bloque.B_pointers[:], pos, tipo, r, soloLectura, padre, delete)
		if ctx.HayErrores() {
			return nil, -1, -1, -1
		}
		if inodo != nil {
			if delete {
				//Verificar si el bloque Indirecto ya no apunta a nada
				bloqueVacio := true
				for _, contenido := range bloque.B_pointers {
					if contenido != -1 {
						bloqueVacio = false
						break
					}
				}
				if bloqueVacio {
					sb.ModificarBitEnBitMap(ctx, driveLetter, int32(apuntadores[pos]), false, false)
					//Verificar si hay errores
					if ctx.HayErrores() {
						return nil, -1, -1, -1
					}
					//Actualizar el apuntador del bloque en el inodo
					apuntadores[pos] = -1
				} else {
					sb.EscribirBloque(ctx, driveLetter, nil, nil, bloque, int64(apuntadores[pos]))
					//Verificar si hay errores
					if ctx.HayErrores() {
						return nil, -1, -1, -1
					}
				}
			} else {
				//Escribir el bloque en el disco
				sb.EscribirBloque(ctx, driveLetter, nil, nil, bloque, int64(apuntadores[pos]))
				//Verificar si hay errores
				if ctx.HayErrores() {
					return nil, -1, -1, -1
				}
			}
			return inodo, numInodo, block, ind
		}
	}
	return nil, -1, -1, -1
}

// Recuperar Bloque Archivo
func (sb *Super_Bloque) RecuperarBloqueArchivo(ctx *Contexto, driveLetter string, numero int) *structs.Block_Archivo {
	//Abrir el archivo del disco
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_RDONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error al abrir el disco al recuperar el bloque: "+err.Error(), 0, 0)
		return nil
	}
	defer file.Close()
	//Mover el puntero al inicio del bloque
	file.Seek(sb.S_block_start+(int64(numero)*sb.S_block_s), io.SeekStart)
	//Leer el bloque
	bloque := &structs.Block_Archivo{}
	err = binary.Read(file, binary.BigEndian, bloque)
	if err != nil {
		ctx.AgregarError("Error al leer el bloque: "+err.Error(), 0, 0)
		return nil
	}
	return bloque
}

// Recuperar Bloque Carpeta
func (sb *Super_Bloque) RecuperarBloqueCarpeta(ctx *Contexto, driveLetter string, numero int) *structs.Block_Carpeta {
	//Abrir el archivo del disco
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_RDONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error al abrir el disco al recuperar el bloque: "+err.Error(), 0, 0)
		return nil
	}
	defer file.Close()
	//Mover el puntero al inicio del bloque
	file.Seek(sb.S_block_start+(int64(numero)*sb.S_block_s), io.SeekStart)
	//Leer el bloque
	bloque := &structs.Block_Carpeta{}
	err = binary.Read(file, binary.BigEndian, bloque)
	if err != nil {
		ctx.AgregarError("Error al leer el bloque: "+err.Error(), 0, 0)
		return nil
	}
	return bloque
}

// Recuperar Bloque Indirecto
func (sb *Super_Bloque) RecuperarBloqueIndirecto(ctx *Contexto, driveLetter string, numero int) *structs.Block_Indirecto {
	//Abrir el archivo del disco
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_RDONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error al abrir el disco al recuperar el bloque: "+err.Error(), 0, 0)
		return nil
	}
	defer file.Close()
	//Mover el puntero al inicio del bloque
	file.Seek(sb.S_block_start+(int64(numero)*sb.S_block_s), io.SeekStart)
	//Leer el bloque
	bloque := &structs.Block_Indirecto{}
	err = binary.Read(file, binary.BigEndian, bloque)
	if err != nil {
		ctx.AgregarError("Error al leer el bloque: "+err.Error(), 0, 0)
		return nil
	}
	return bloque
}

// Graphviz del Super Bloque
func (sb *Super_Bloque) ReporteSuperBloque() string {
	return `
	digraph G {
		// Título que ocupa dos columnas al principio
		table [label=< 
				<table border="1" cellspacing="0" >
					<tr>
						<td colspan="2" bgcolor="#0000CC"><font color="white"><b>Super Bloque</b></font></td>
					</tr>
					<tr>
						<td>s_filesystem_type</td><td>` + strconv.Itoa(int(sb.S_filesystem_type)) + `</td>
					</tr>
					<tr>
						<td>s_inodes_count</td><td>` + strconv.Itoa(int(sb.S_inodes_count)) + `</td>
					</tr>
					<tr>
						<td>s_blocks_count</td><td>` + strconv.Itoa(int(sb.S_blocks_count)) + `</td>
					</tr>
					<tr>
						<td>s_free_blocks_count</td><td>` + strconv.Itoa(int(sb.S_free_blocks_count)) + `</td>
					</tr>
					<tr>
						<td>s_free_inodes_count</td><td>` + strconv.Itoa(int(sb.S_free_inodes_count)) + `</td>
					</tr>
					<tr>
						<td>s_mtime</td><td>` + string(sb.S_mtime[:]) + `</td>
					</tr>
					<tr>
						<td>s_umtime</td><td>` + string(sb.S_umtime[:]) + `</td>
					</tr>
					<tr>
						<td>s_mnt_count</td><td>` + strconv.Itoa(int(sb.S_mnt_count)) + `</td>
					</tr>
					<tr>
						<td>s_magic</td><td>` + strconv.Itoa(int(sb.S_magic)) + `</td>
					</tr>
					<tr>
						<td>s_inode_s</td><td>` + strconv.Itoa(int(sb.S_inode_s)) + `</td>
					</tr>
					<tr>
						<td>s_block_s</td><td>` + strconv.Itoa(int(sb.S_block_s)) + `</td>
					</tr>
					<tr>
						<td>s_first_ino</td><td>` + strconv.Itoa(int(sb.S_first_ino)) + `</td>
					</tr>
					<tr>
						<td>s_first_blo</td><td>` + strconv.Itoa(int(sb.S_first_blo)) + `</td>
					</tr>
					<tr>
						<td>s_bm_inode_start</td><td>` + strconv.Itoa(int(sb.S_bm_inode_start)) + `</td>
					</tr>
					<tr>
						<td>s_bm_block_start</td><td>` + strconv.Itoa(int(sb.S_bm_block_start)) + `</td>
					</tr>
					<tr>
						<td>s_inode_start</td><td>` + strconv.Itoa(int(sb.S_inode_start)) + `</td>
					</tr>
					<tr>
						<td>s_block_start</td><td>` + strconv.Itoa(int(sb.S_block_start)) + `</td>
					</tr>
					</table> 
				> shape=box style=invisible ] 
	}
	`
}

// Bitmap de Inodos
// tipo = true -> Inodes | tipo = false -> Bloques
func (sb *Super_Bloque) Bitmap(ctx *Contexto, driveLetter string, tipo bool) []byte {
	//Start y count del bitmap
	var start int64
	var count int64
	if tipo {
		start = sb.S_bm_inode_start
		count = sb.S_inodes_count
	} else {
		start = sb.S_bm_block_start
		count = sb.S_blocks_count

	}
	//Abrir el archivo del disco
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_RDONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error al abrir el disco al obtener el Bitmap de Inodos: "+err.Error(), 0, 0)
		return nil
	}
	defer file.Close()
	//Mover el puntero al inicio del bitmap de inodos
	file.Seek(start, io.SeekStart)
	//Leer el bitmap de inodos
	bitMap := make([]byte, count)
	err = binary.Read(file, binary.BigEndian, &bitMap)
	if err != nil {
		ctx.AgregarError("Error al leer el Bitmap de Inodos: "+err.Error(), 0, 0)
		return nil
	}
	return bitMap
}
