package comandos

type Loss struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// loss *Loss ExprCommand
func (loss *Loss) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar que haya parametros
	if len(loss.Parametros) == 0 {
		ctx.AgregarError("El comando loss no tiene parametros", loss.Linea, loss.Columna)
		return nil
	}
	//Verificar parametros obligatorios -id
	if _, existe := loss.Parametros["-id"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -id", loss.Linea, loss.Columna)
		return nil
	}
	//Obtener id
	id := loss.Parametros["-id"]
	//Buscar Particion Montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(id)
	if !existe {
		ctx.AgregarError("Error: No hay ninguna particion montada", loss.Linea, loss.Columna)
		return nil
	}
	//Obtener SuperBloque
	superBloque, err := GetSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo obtener el super bloque de la particion", loss.Linea, loss.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Limpiar Bitmpas, Inodos y Bloques
	superBloque.LimpiarBitmapsInodosBloques(ctx, partMontada.DiskName)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Crear carpetas raiz y archivo users.txt
	superBloque.CrearCarpetaRaiz(ctx, partMontada.DiskName)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Crear archivo users.txt
	superBloque.CrearArchivoUsers(ctx, partMontada.DiskName)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Escribir SuperBloque
	superBloque.EscribirSuperBloque(ctx, partMontada.DiskName, partMontada.Start_SuperBloque)
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Mensaje de exito
	//fmt.Println("----------Comando LOSS--------------")
	//fmt.Println("Bitmpas, Inodos y Bloques limpiados con exito de la particion: " + partMontada.PartName)
	ctx.AgregarOutput("----------Comando LOSS--------------")
	ctx.AgregarOutput("La particion con id \"" + id + "\" se ha formateado correctamente")
	return nil
}
