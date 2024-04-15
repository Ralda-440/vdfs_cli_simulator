package comandos

import (
	"regexp"
	"strings"
)

type Cat struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// cat *Cat ExprCommand
func (cat *Cat) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar que haya parametros
	if len(cat.Parametros) == 0 {
		ctx.AgregarError("El comando cat no tiene parametros", cat.Linea, cat.Columna)
		return nil
	}
	paths := []string{}
	//Verificar parametros obligatorios -file+numeral y agregar el path
	pattern := regexp.MustCompile(`^(-file\d+)$`)
	for key, value := range cat.Parametros {
		if pattern.MatchString(key) {
			paths = append(paths, value)
		}
	}
	//Verificar si hay sesion activa
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay ninguna sesion activa para el comando Cat", cat.Linea, cat.Columna)
		return nil
	}
	//Recuperar Sesion Actual
	_, _, idPart := SesionActiva.SesionActual()
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+idPart+"\" no esta montada", cat.Linea, cat.Columna)
		return nil
	}
	//Recuperar SuperBloque
	superBloque, err := getSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", cat.Linea, cat.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Leer contenido de los archivos
	contenido := []byte{}
	for _, path := range paths {
		//Convertir path en []string
		path_ := strings.Split(path, "/")
		//Eliminar el primer elemento
		path_ = path_[1:]
		cont := superBloque.LeerContenidoArchivo(ctx, partMontada.DiskName, path_)
		//Agregar salto de linea
		cont = append(cont, []byte("\n")...)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return nil
		}
		contenido = append(contenido, cont...)
	}
	//Convertir contenido a string
	contStr := string(contenido)
	//Imprimir contenido
	//fmt.Println("----------Comando CAT--------------")
	//fmt.Println(contStr)
	//fmt.Println("----------------------------------")
	//fmt.Println("Comando CAT ejecutado correctamente")
	ctx.AgregarOutput("--------------------Comando CAT--------------------")
	for _, path := range paths {
		ctx.AgregarOutput("Archivo: " + path)
	}
	ctx.AgregarOutput("--------------------------------------------------")
	ctx.AgregarOutput(contStr)
	return nil
}
