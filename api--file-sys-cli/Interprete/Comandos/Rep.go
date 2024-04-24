package comandos

import (
	structs "app/Interprete/Structs"
	"encoding/binary"
	"os"
	"os/exec"
	"strings"
)

type REP struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

//r *REP ExprCommand

func (r *REP) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(r.Parametros) == 0 {
		ctx.AgregarError("Error: El comando no tiene parametros", r.Linea, r.Columna)
		return nil
	}
	//Verificar parametro obligatorio -name
	if _, ok := r.Parametros["-name"]; !ok {
		ctx.AgregarError("Error: Falta el parametro obligatorio -name", r.Linea, r.Columna)
		return nil
	}
	//Verificar parametro obligatorio -path
	if _, ok := r.Parametros["-path"]; !ok {
		ctx.AgregarError("Error: Falta el parametro obligatorio -path", r.Linea, r.Columna)
		return nil
	}
	//Verificar parametro obligatorio -id
	if _, ok := r.Parametros["-id"]; !ok {
		ctx.AgregarError("Error: Falta el parametro obligatorio -id", r.Linea, r.Columna)
		return nil
	}
	//Obtener el parametro -name
	name := r.Parametros["-name"]
	//Obtener el parametro -path
	path := r.Parametros["-path"]
	//Obtener el parametro -id
	id := r.Parametros["-id"]
	//Parametro Opcional -ruta para reporte file y ls
	ruta := ""
	if _, ok := r.Parametros["-ruta"]; !ok && (name == "file" || name == "ls") {
		ctx.AgregarError("Error: Falta el parametro opcional -ruta", r.Linea, r.Columna)
		return nil
	} else if (name == "file" || name == "ls") && ok {
		ruta = r.Parametros["-ruta"]
	}
	path_ := strings.Split(path, "/")
	nameRep := path_[len(path_)-1]
	//Swith para cada tipo de reporte
	graphviz := ""
	switch name {
	case "mbr":
		graphviz = r.RepMBR(string(id[0]), ctx)
	case "disk":
		graphviz = r.RepDisk(string(id[0]), ctx)
	case "sb":
		graphviz = r.RepSB(id, ctx)
	case "bm_inode":
		r.RepBM(ctx, id, true, path)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil
		}
		//Mensaje de exito
		//fmt.Println("----------Comando REP--------------")
		//fmt.Println("Reporte generado con exito")
		ctx.AgregarOutput("----------Comando REP: " + name + " --------------")
		ctx.AgregarOutput("El Reporte: \"" + nameRep + "\" se genero con exito")
		return nil
	case "bm_block":
		r.RepBM(ctx, id, false, path)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil
		}
		//Mensaje de exito
		//fmt.Println("----------Comando REP--------------")
		//fmt.Println("Reporte generado con exito")
		ctx.AgregarOutput("----------Comando REP: " + name + " --------------")
		ctx.AgregarOutput("El Reporte: \"" + nameRep + "\" se genero con exito")
		return nil
	case "inode":
		graphviz = r.RepInode(ctx, id)
		//Verificar si hay errores
		if ctx.HayErrores() {
			ctx.AgregarError("Error: No se pudo generar el reporte de inodos", r.Linea, r.Columna)
			return nil
		}
	case "tree":
		graphviz = r.RepTree(ctx, id)
		//Verificar si hay errores
		if ctx.HayErrores() {
			ctx.AgregarError("Error: No se pudo generar el reporte del arbol de directorios", r.Linea, r.Columna)
			return nil
		}
	case "block":
		graphviz = r.RepBlock(ctx, id)
		//Verificar si hay errores
		if ctx.HayErrores() {
			ctx.AgregarError("Error: No se pudo generar el reporte de bloques", r.Linea, r.Columna)
			return nil
		}
	case "file":
		r.RepFile(ctx, id, ruta, path)
		//Verificar si hay errores
		if ctx.HayErrores() {
			ctx.AgregarError("Error: No se pudo generar el reporte de file", r.Linea, r.Columna)
			return nil
		}
		ctx.AgregarOutput("----------Comando REP: " + name + " --------------")
		ctx.AgregarOutput("El Reporte: \"" + nameRep + "\" de la ruta: \"" + ruta + "\" se genero con exito")
		return nil
	case "ls":
		graphviz = r.RepLs(ctx, id, ruta)
	case "journaling":
		graphviz = r.RepJournaling(ctx, id)
	default:
		ctx.AgregarError("Error: El nombre del reporte no es valido", r.Linea, r.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo generar el reporte", r.Linea, r.Columna)
		return nil
	}
	//Generar archivo .dot y .png
	r.GenerarArchivoDotYPng(graphviz, path, ctx)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Mensaje de exito
	//fmt.Println("----------Comando REP: " + name + " --------------")
	//fmt.Println("Reporte generado con exito")
	ctx.AgregarOutput("----------Comando REP: " + name + " --------------")
	if ruta == "" {
		ctx.AgregarOutput("El Reporte: \"" + nameRep + "\" se genero con exito")
		return nil
	}
	ctx.AgregarOutput("El Reporte: \"" + nameRep + "\" de la ruta: \"" + ruta + "\" se genero con exito")
	return nil
}

// Rep Journaling
func (r *REP) RepJournaling(ctx *Contexto, id string) string {
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(id)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+id+"\" no esta montada", r.Linea, r.Columna)
		return ""
	}
	//Recuperar SuperBloque
	superBloque, err := GetSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", r.Linea, r.Columna)
		return ""
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return ""
	}
	//Graphviz de journaling
	operacionesJournaling := superBloque.RecuperarJournaling(ctx, partMontada.DiskName, partMontada.Start_SuperBloque)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return ""
	}
	graphviz := ""
	for _, op := range operacionesJournaling {
		graphviz += op.GraficarJournaling()
	}
	graphviz = `
	digraph G {
		// Título que ocupa dos columnas al principio
		repLs [label=< 
				<table border="1" cellspacing="0" cellpadding='10' >
					<tr>
						<td bgcolor="#800080"><font color="white"><b>Operacion</b></font></td>
						<td bgcolor="#800080"><font color="white"><b>Path</b></font></td>
						<td bgcolor="#800080"><font color="white"><b>Contenido</b></font></td>
						<td bgcolor="#800080"><font color="white"><b>Fecha</b></font></td>
					</tr>
					` + graphviz + `
	            </table> 
            > shape=box style=invisible ] 
	}
	`
	return graphviz
}

// Rep ls
func (r *REP) RepLs(ctx *Contexto, id string, ruta string) string {
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(id)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+id+"\" no esta montada", r.Linea, r.Columna)
		return ""
	}
	//Recuperar SuperBloque
	superBloque, err := GetSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", r.Linea, r.Columna)
		return ""
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return ""
	}
	//Convertir ruta a []string
	rutaArray := strings.Split(ruta, "/")
	//Eliminar el primer elemento
	rutaArray = rutaArray[1:]
	//Eliminar el ultimo elemento si es vacio
	if rutaArray[len(rutaArray)-1] == "" {
		rutaArray = rutaArray[:len(rutaArray)-1]
	}
	//Reporte Ls
	graphvisLs, _ := superBloque.ReporteLs(ctx, partMontada.DiskName, rutaArray)
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo generar el reporte de ls", r.Linea, r.Columna)
		return ""
	}
	return graphvisLs
}

// Rep File
func (r *REP) RepFile(ctx *Contexto, id string, ruta string, path string) {
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(id)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+id+"\" no esta montada", r.Linea, r.Columna)
		return
	}
	//Recuperar SuperBloque
	superBloque, err := GetSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", r.Linea, r.Columna)
		return
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return
	}
	//Convertir ruta a []string
	rutaArray := strings.Split(ruta, "/")
	//Eliminar el primer elemento
	rutaArray = rutaArray[1:]
	//Recuperar el archivo
	contenido := superBloque.LeerContenidoArchivo(ctx, partMontada.DiskName, rutaArray)
	if ctx.HayErrores() {
		return
	}
	//Crear archivo con el contenido
	path_ := strings.Split(path, "/")
	nameRep := path_[len(path_)-1]
	file, err := os.Create("./REP/" + nameRep + ".txt")
	if err != nil {
		ctx.AgregarError("Error: No se pudo crear el archivo con el contenido: "+err.Error(), r.Linea, r.Columna)
		return
	}
	defer file.Close()
	//Convertir contenido a string
	contenidoStr := string(contenido)
	//Escribir contenido en el archivo
	_, err = file.WriteString(contenidoStr)
	if err != nil {
		ctx.AgregarError("Error: No se pudo escribir el contenido en el archivo: "+err.Error(), r.Linea, r.Columna)
		return
	}
	//Mensaje de exito
	//fmt.Println("----------Comando REP: file--------------")
	//fmt.Println("Archivo generado con exito")
}

// Rep Block
func (r *REP) RepBlock(ctx *Contexto, id string) string {
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(id)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+id+"\" no esta montada", r.Linea, r.Columna)
		return ""
	}
	//Recuperar SuperBloque
	superBloque, err := GetSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", r.Linea, r.Columna)
		return ""
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return ""
	}
	//Graphviz de bloques
	graphviz, _ := superBloque.ReporteGraphviz(ctx, partMontada.DiskName, 2)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return ""
	}
	return graphviz
}

// Rep Tree
func (r *REP) RepTree(ctx *Contexto, id string) string {
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(id)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+id+"\" no esta montada", r.Linea, r.Columna)
		return ""
	}
	//Recuperar SuperBloque
	superBloque, err := GetSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", r.Linea, r.Columna)
		return ""
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return ""
	}
	//Graphviz del arbol de directorios
	graphviz, _ := superBloque.ReporteGraphviz(ctx, partMontada.DiskName, 0)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return ""
	}
	return graphviz
}

// Rep Inode
func (r *REP) RepInode(ctx *Contexto, id string) string {
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(id)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+id+"\" no esta montada", r.Linea, r.Columna)
		return ""
	}
	//Recuperar SuperBloque
	superBloque, err := GetSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", r.Linea, r.Columna)
		return ""
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return ""
	}
	//Graphviz de inodos
	graphviz, _ := superBloque.ReporteGraphviz(ctx, partMontada.DiskName, 1)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return ""
	}
	return graphviz
}

// Rep MBR
func (r *REP) RepMBR(name string, ctx *Contexto) string {
	//Leer el MBR del disco
	mbr, err := GetMBRDisk(name)
	if err != nil {
		ctx.AgregarError("Error: No se pudo leer el MBR del disco : "+err.Error(), r.Linea, r.Columna)
		return ""
	}
	graphviz, err := mbr.ReporteMBR(name)
	if err != nil {
		ctx.AgregarError("Error: No se pudo generar el reporte del MBR : "+err.Error(), r.Linea, r.Columna)
		return ""
	}
	//Obtener graphviz del MBR
	return graphviz
}

// Rep Disk
func (r *REP) RepDisk(name string, ctx *Contexto) string {
	//Leer el MBR del disco
	mbr, err := GetMBRDisk(name)
	if err != nil {
		ctx.AgregarError("Error: No se pudo leer el MBR del disco : "+err.Error(), r.Linea, r.Columna)
		return ""
	}
	graphviz, err := mbr.ReporteDisk(name)
	if err != nil {
		ctx.AgregarError("Error: No se pudo generar el reporte del disco : "+err.Error(), r.Linea, r.Columna)
		return ""
	}
	return graphviz
}

// Reporte SuperBloque
func (r *REP) RepSB(id string, ctx *Contexto) string {
	//Verificar si la particion esta montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(id)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+id+"\" no esta montada", r.Linea, r.Columna)
		return ""
	}
	//Obtener MBR del disco donde esta la particion
	mbr, err := GetMBRDisk(partMontada.DiskName)
	if err != nil {
		ctx.AgregarError("Error : El disco no existe o hubo error al leer su MBR :"+err.Error(), r.Linea, r.Columna)
		return ""
	}
	//Buscar particion con su nombre
	existe, particion, _, ebr, _, _, _, err := mbr.BuscarParticion(partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error : No se pudo buscar la particion "+err.Error(), r.Linea, r.Columna)
		return ""
	}
	if !existe {
		ctx.AgregarError("Error : No existe la particion con el nombre: "+id, r.Linea, r.Columna)
		return ""
	}
	//Determinar donde leer el super bloque
	var start_part int64
	if particion != nil {
		start_part = particion.Part_start
	} else {
		start_part = ebr.Part_start + int64(binary.Size(structs.EBR{}))
	}
	//Abrir el archivo del disco para leer el super bloque
	file, err := os.OpenFile("./MIA/P1/"+partMontada.DiskName+".dsk", os.O_RDONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error al abrir el disco al leer el super bloque: "+err.Error(), r.Linea, r.Columna)
		return ""
	}
	defer file.Close()
	//Leer el super bloque
	sb := &Super_Bloque{}
	file.Seek(start_part, 0)
	err = binary.Read(file, binary.BigEndian, sb)
	if err != nil {
		ctx.AgregarError("Error al leer el super bloque: "+err.Error(), r.Linea, r.Columna)
		return ""
	}
	return sb.ReporteSuperBloque()
}

// Reporte Bitmap
// tipo = true -> Inodos | false -> Bloques
func (r *REP) RepBM(ctx *Contexto, id string, tipo bool, path string) {
	//Verificar si la particion esta montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(id)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+id+"\" no esta montada", r.Linea, r.Columna)
		return
	}
	//Obtener MBR del disco donde esta la particion
	mbr, err := GetMBRDisk(partMontada.DiskName)
	if err != nil {
		ctx.AgregarError("Error : El disco no existe o hubo error al leer su MBR :"+err.Error(), r.Linea, r.Columna)
		return
	}
	//Buscar particion con su nombre
	existe, particion, _, ebr, _, _, _, err := mbr.BuscarParticion(partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error : No se pudo buscar la particion "+err.Error(), r.Linea, r.Columna)
		return
	}
	if !existe {
		ctx.AgregarError("Error : No existe la particion con el nombre: "+id, r.Linea, r.Columna)
		return
	}
	//Determinar donde leer el super bloque
	var start_part int64
	if particion != nil {
		start_part = particion.Part_start
	} else {
		start_part = ebr.Part_start + int64(binary.Size(structs.EBR{}))
	}
	//Abrir el archivo del disco para leer el super bloque
	file, err := os.OpenFile("./MIA/P1/"+partMontada.DiskName+".dsk", os.O_RDONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error al abrir el disco al leer el super bloque: "+err.Error(), r.Linea, r.Columna)
		return
	}
	defer file.Close()
	//Leer el super bloque
	sb := &Super_Bloque{}
	file.Seek(start_part, 0)
	err = binary.Read(file, binary.BigEndian, sb)
	if err != nil {
		ctx.AgregarError("Error al leer el super bloque: "+err.Error(), r.Linea, r.Columna)
		return
	}
	//Obtener el bitmap
	bmInode := sb.Bitmap(ctx, partMontada.DiskName, tipo)
	//Crear archivo con el path y escribir el bitmap
	path_ := strings.Split(path, "/")
	nameRep := path_[len(path_)-1]
	file, err = os.Create("./REP/" + nameRep + ".txt")
	if err != nil {
		ctx.AgregarError("Error al crear el archivo del bitmap: "+err.Error(), r.Linea, r.Columna)
		return
	}
	defer file.Close()
	//Sustituir por sus valores ASCII
	for i := range bmInode {
		if bmInode[i] == 1 {
			bmInode[i] = '1'
		} else if bmInode[i] == 0 {
			bmInode[i] = '0'
		}
	}
	//Escribir 20 bits por linea
	for i := 0; i < len(bmInode); i += 20 {
		err := error(nil)
		if i+20 < len(bmInode) {
			err = binary.Write(file, binary.BigEndian, bmInode[i:i+20])
		} else {
			err = binary.Write(file, binary.BigEndian, bmInode[i:])
		}
		//Verificar si hay errores
		if err != nil {
			ctx.AgregarError("Error al escribir el bitmap En el Reporte: "+err.Error(), r.Linea, r.Columna)
			return
		}
		err = binary.Write(file, binary.BigEndian, []byte("\n"))
		//Verificar si hay errores
		if err != nil {
			ctx.AgregarError("Error al escribir el bitmap En el Reporte: "+err.Error(), r.Linea, r.Columna)
			return
		}
	}
}

// Generar archivo .dot y .png
func (r *REP) GenerarArchivoDotYPng(graphviz string, path string, ctx *Contexto) {
	//Crear directorios del path si no existen
	/*
		dirs := filepath.Dir(path)
		err := os.MkdirAll(dirs, 0777)
		if err != nil {
			ctx.AgregarError("Error: No se pudo crear el directorio del archivo .dot y .png, "+err.Error(), r.Linea, r.Columna)
			return
		}
	*/
	//Generar archivo .dot
	path_ := strings.Split(path, "/")
	nameRep := path_[len(path_)-1]
	file, err := os.Create("./REP/" + nameRep + ".dot")
	if err != nil {
		ctx.AgregarError("Error: No se pudo crear el archivo .dot, "+err.Error(), r.Linea, r.Columna)
		return
	}
	defer file.Close()
	file.WriteString(graphviz)
	//Generar archivo .png
	cmd := exec.Command("dot", "-Tpng", "./REP/"+nameRep+".dot", "-o", "./REP/"+nameRep+".png")
	err = cmd.Run()
	if err != nil {
		ctx.AgregarError("Error: No se pudo generar el archivo .png, "+err.Error(), r.Linea, r.Columna)
		return
	}
	//Eliminar archivo .dot
	err = os.Remove("./REP/" + nameRep + ".dot")
	if err != nil {
		ctx.AgregarError("Error: No se pudo eliminar el archivo .dot, "+err.Error(), r.Linea, r.Columna)
		return
	}
}
