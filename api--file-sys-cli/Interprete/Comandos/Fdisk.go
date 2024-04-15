package comandos

import (
	structs "app/Interprete/Structs"
	"strconv"
)

type Fdisk struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// f *fdk ExprCommand
func (f *Fdisk) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(f.Parametros) == 0 {
		ctx.AgregarError("Error: El comando no tiene parametros", f.Linea, f.Columna)
		return nil
	}
	//##Verificacion de parametro driveletter y name obligatorios para todas las operaciones de fdisk
	//Verificar parametro obligatorio -driveletter
	if _, ok := f.Parametros["-driveletter"]; !ok {
		ctx.AgregarError("Error: Falta el parametro obligatorio -driveletter", f.Linea, f.Columna)
		return nil
	}
	//Verificar parametro obligatorio -name
	if _, ok := f.Parametros["-name"]; !ok {
		ctx.AgregarError("Error: Falta el parametro obligatorio -name", f.Linea, f.Columna)
		return nil
	}
	//Verificar que parametro viene primero(-size, -delete, -add) para saber que operacion realizar(Crear, Eliminar, Modificar)
	for key := range f.Parametros {
		if key == "-size" {
			//##Crear Particion
			f.CrearParticion(ctx)
			return nil
		} else if key == "-delete" {
			//##Eliminar Particion
			f.EliminarParticion(ctx)
			return nil
		} else if key == "-add" {
			//##Modificar Particion
			f.ModificarParticion(ctx)
			return nil
		}
	}
	//Si no viene ningun parametro valido
	ctx.AgregarError("Error: No se puede realizar ninguna operacion (Crear, Eliminar, Modificar) con los parametros actuales", f.Linea, f.Columna)
	return nil
}

// ##Crear Particion
func (f *Fdisk) CrearParticion(ctx *Contexto) {
	//Parametros obligatorios para crear particion
	size := f.Parametros["-size"]
	drivelleter := f.Parametros["-driveletter"]
	name := f.Parametros["-name"]
	//Parametros opcionales para crear particion
	unit := "K"
	if value, ok := f.Parametros["-unit"]; ok {
		unit = value
	}
	type_ := "P"
	if value, ok := f.Parametros["-type"]; ok {
		type_ = value
	}
	fit := "W"
	if value, ok := f.Parametros["-fit"]; ok {
		fit = value
	}
	//Obtener mbr del disco donde se creara la particion
	mbr, err := GetMBRDisk(drivelleter)
	if err != nil {
		ctx.AgregarError("Error: No se pudo leer el MBR del disco : "+err.Error(), f.Linea, f.Columna)
		return
	}
	//Buscar particion por nombre
	existe, _, _, _, _, cantPartPrimarias, cantPartExtendidas, err := mbr.BuscarParticion(drivelleter, name)
	if err != nil {
		ctx.AgregarError("Error: Error al buscar la Particion : "+err.Error(), f.Linea, f.Columna)
		return
	}
	if existe {
		ctx.AgregarError("Error: Ya existe una particion con el nombre "+name, f.Linea, f.Columna)
		return
	}
	//Verificar si se puede crear una particion
	if cantPartPrimarias+cantPartExtendidas == 4 && type_ != "L" {
		ctx.AgregarError("Error: No se puede crear mas particiones en el disco", f.Linea, f.Columna)
		return
	}
	//No se verifica si se puede crear una particion primaria, ya que puede haber hasta 4 particiones primarias
	//Verificar si se puede crear una particion extendida
	if type_ == "E" && cantPartExtendidas == 1 {
		ctx.AgregarError("Error: No se puede crear mas particiones extendidas", f.Linea, f.Columna)
		return
	}
	//Verificar si se puede crear una particion logica
	if type_ == "L" && cantPartExtendidas == 0 {
		ctx.AgregarError("Error: No se puede crear una particion logica sin una particion extendida", f.Linea, f.Columna)
		return
	}
	//Convertir size a int64
	sizeInt, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		ctx.AgregarError("Error: El valor del parametro -size no es valido", f.Linea, f.Columna)
		return
	}
	//Size mayor a 0
	if sizeInt <= 0 {
		ctx.AgregarError("Error: El valor del parametro -size debe ser mayor a 0", f.Linea, f.Columna)
		return
	}
	//Calcular size de la particion (unit = B|K|M)
	if unit == "B" {
		sizeInt = sizeInt * 1
	} else if unit == "K" {
		sizeInt = sizeInt * 1024
	} else if unit == "M" {
		sizeInt = sizeInt * 1024 * 1024
	} else {
		ctx.AgregarError("Error: El valor del parametro -unit no es valido", f.Linea, f.Columna)
		return
	}
	//Verificar que el valor de fit sea B, F o W
	if fit != "B" && fit != "F" && fit != "W" {
		ctx.AgregarError("Error: El valor del parametro -fit no es valido", f.Linea, f.Columna)
		return
	}
	//Verificar que tipo de particion se creara
	if type_ == "L" {
		//##Crear Particion Logica
		f.CrearParticionLogica(ctx, mbr, drivelleter, name, sizeInt, unit, fit, size)
	} else if type_ == "P" || type_ == "E" {
		//##Crear Particion Primaria o Extendida
		f.CrearParticionPrimariaOExtendida(ctx, mbr, drivelleter, name, sizeInt, unit, fit, type_, size)
	} else {
		ctx.AgregarError("Error: El valor del parametro -type no es valido", f.Linea, f.Columna)
		return
	}
}

// ##Crear Particion Primaria o Extendida
func (f *Fdisk) CrearParticionPrimariaOExtendida(ctx *Contexto, mbr *structs.MBR, drivelleter string, name string, size int64, unit string, fit string, type_ string, sizeStr string) {
	//Buscar espacio para la particion
	existe, start_absolute := mbr.BuscarEspacio(size)
	if !existe {
		ctx.AgregarError("Error: No hay espacio suficiente para crear la particion", f.Linea, f.Columna)
		return
	}
	//Crear particion
	particion := structs.NewParticion()
	particion.Part_status = byte('0')
	particion.Part_type = byte(type_[0])
	particion.Part_fit = byte(fit[0])
	particion.Part_start = start_absolute
	particion.Part_s = size
	copy(particion.Part_name[:], name)
	particion.Part_correlative = mbr.NuevoCorrelativo()
	particion.Part_id = [4]byte{45, 45, 45, 45}
	//Agregar particion al MBR
	mbr.AgregarParticion(particion)
	//Ordenar particiones
	mbr.OrdenarParticiones()
	//Escribir MBR en el disco
	err := writeMBRInDisk(drivelleter, mbr)
	if err != nil {
		ctx.AgregarError("Error: No se pudo escribir el MBR en el disco : "+err.Error(), f.Linea, f.Columna)
		return
	}
	//Escribir Primer EBR en el disco si es una particion extendida
	if type_ == "E" {
		ebr := structs.NewEBR()
		ebr.Part_mount = byte(0)
		ebr.Part_fit = byte(0)
		ebr.Part_start = start_absolute
		ebr.Part_s = -1
		ebr.Part_next = -1
		ebr.Part_name = [16]byte{}
		//Escribir EBR en el disco
		err = particion.WriteEBRsInDisk(drivelleter, []structs.EBR{ebr})
		if err != nil {
			ctx.AgregarError("Error: No se pudo escribir el EBR en el disco : "+err.Error(), f.Linea, f.Columna)
			return
		}
	}
	//Mensaje de exito
	//fmt.Println("----------Comando FDISK----------")
	//fmt.Println("Particion \"" + name + "\" creada en el disco \"" + drivelleter + "\" con exito. Tamaño: " + sizeStr + " " + unit)
	ctx.AgregarOutput("----------Comando FDISK----------")
	ctx.AgregarOutput("Particion " + type_ + " \"" + name + "\" creada en el disco \"" + drivelleter + "\" con exito. Tamaño: " + sizeStr + " " + unit)
}

// ##Crear Particion Logica
func (f *Fdisk) CrearParticionLogica(ctx *Contexto, mbr *structs.MBR, drivelleter string, name string, size int64, unit string, fit string, sizeStr string) {
	//Buscar particion extendida
	extendida := mbr.BuscarParticionExtendida()
	if extendida == nil {
		ctx.AgregarError("Error: No se encontro una particion extendida", f.Linea, f.Columna)
		return
	}
	//Obtener EBRs del disco
	ebrs, err := extendida.ReadEBRsInDisk(drivelleter)
	if err != nil {
		ctx.AgregarError("Error: No se pudieron leer los EBRs del disco : "+err.Error(), f.Linea, f.Columna)
		return
	}
	//Buscar espacio para la particion logica
	hay_espacio, start_absolute, next := extendida.BuscarEspacioParticionLogica(size, &ebrs)
	if !hay_espacio {
		ctx.AgregarError("Error: No hay espacio suficiente para crear la particion logica", f.Linea, f.Columna)
		return
	}
	//Crear particion logica
	particion := structs.NewEBR()
	particion.Part_mount = byte('0')
	particion.Part_fit = byte(fit[0])
	particion.Part_start = start_absolute
	particion.Part_s = size
	particion.Part_next = next
	copy(particion.Part_name[:], name)
	//Escribir EBR en el disco
	//Agregar nueva particion logica a la lista de EBRs
	ebrs = append(ebrs, particion)
	err = extendida.WriteEBRsInDisk(drivelleter, ebrs)
	if err != nil {
		ctx.AgregarError("Error: No se pudo escribir el EBR en el disco : "+err.Error(), f.Linea, f.Columna)
		return
	}
	//Mensaje de exito
	//fmt.Println("----------Comando FDISK----------")
	//fmt.Println("Particion Extendida \"" + name + "\" creada en el disco \"" + drivelleter + "\" con exito. Tamaño: " + sizeStr + " " + unit)
	ctx.AgregarOutput("----------Comando FDISK----------")
	ctx.AgregarOutput("Particion Logica \"" + name + "\" creada en la particion extendida con exito. Tamaño: " + sizeStr + " " + unit)
}

// ##Eliminar Particion
func (f *Fdisk) EliminarParticion(ctx *Contexto) {
	//Parametros obligatorios para eliminar una particion
	delete := f.Parametros["-delete"]
	//Verificar que delete tenga el valor Full
	if delete != "Full" {
		ctx.AgregarError("Error: El valor del parametro -delete no es valido", f.Linea, f.Columna)
		return
	}
	drivelleter := f.Parametros["-driveletter"]
	name := f.Parametros["-name"]
	//Obtener mbr del disco donde se eliminara la particion
	mbr, err := GetMBRDisk(drivelleter)
	if err != nil {
		ctx.AgregarError("Error: No se pudo leer el MBR del disco : "+err.Error(), f.Linea, f.Columna)
		return
	}
	//Buscar particion por nombre
	existe, particion, _, ebr, antEBR, _, _, err := mbr.BuscarParticion(drivelleter, name)
	if err != nil {
		ctx.AgregarError("Error: Error al buscar la Particion : "+err.Error(), f.Linea, f.Columna)
		return
	}
	if !existe {
		ctx.AgregarError("Error: No existe una particion con el nombre "+name+" para eliminar", f.Linea, f.Columna)
		return
	}
	//Si es primaria o extendida
	if particion != nil {
		//Eliminar particion
		particion.Part_correlative = -1
		//Ordenar particiones
		mbr.OrdenarParticiones()
		//Escribir MBR en el disco
		err = writeMBRInDisk(drivelleter, mbr)
		if err != nil {
			ctx.AgregarError("Error: No se pudo escribir el MBR en el disco : "+err.Error(), f.Linea, f.Columna)
			return
		}
		//Mensaje de exito
		//fmt.Println("----------Comando FDISK----------")
		//fmt.Println("Particion \"" + name + "\" eliminada del disco \"" + drivelleter + "\" con exito")
		ctx.AgregarOutput("----------Comando FDISK----------")
		ctx.AgregarOutput("Particion \"" + name + "\" eliminada del disco \"" + drivelleter + "\" con exito")
		return
	}
	//Si es logica
	//Nuevo EBRs para los EBRs modificados para eliminar la particion logica
	ebrs := make([]structs.EBR, 0)
	//##Cambiar next del EBR anterior si la particion logica a eliminar no es la primera
	if antEBR != nil {
		antEBR.Part_next = ebr.Part_next
		ebrs = append(ebrs, *antEBR)
	}
	//##Cambiar part_S del EBR a eliminar para eliminar la particion logica
	ebr.Part_s = -1
	ebrs = append(ebrs, *ebr)
	//Escribir EBRs modificados en el disco
	err = particion.WriteEBRsInDisk(drivelleter, ebrs)
	if err != nil {
		ctx.AgregarError("Error: No se pudo escribir los EBRs modificados en el disco : "+err.Error(), f.Linea, f.Columna)
		return
	}
	//Mensaje de exito
	//fmt.Println("----------Comando FDISK----------")
	//fmt.Println("Particion Logica \"" + name + "\" eliminada de la particion extendida con exito")
	ctx.AgregarOutput("----------Comando FDISK----------")
	ctx.AgregarOutput("Particion Logica \"" + name + "\" eliminada de la particion extendida con exito")
}

// ##Modificar Particion
func (f *Fdisk) ModificarParticion(ctx *Contexto) {
	//Parametros obligatorios para modificar una particion
	add := f.Parametros["-add"] //Valor entero positivo o negativo
	add_, err := strconv.Atoi(add)
	if err != nil {
		ctx.AgregarError("Error: El valor del parametro -add no es valido", f.Linea, f.Columna)
		return
	}
	drivelleter := f.Parametros["-driveletter"]
	name := f.Parametros["-name"]
	//Parametro opcional para modificar una particion
	unit := "K"
	if value, ok := f.Parametros["-unit"]; ok {
		unit = value
	}
	//Obtener mbr del disco donde se modificara la particion
	mbr, err := GetMBRDisk(drivelleter)
	if err != nil {
		ctx.AgregarError("Error: No se pudo leer el MBR del disco : "+err.Error(), f.Linea, f.Columna)
		return
	}
	//Buscar particion por nombre
	existe, particion, posPart, ebr, _, _, _, err := mbr.BuscarParticion(drivelleter, name)
	if err != nil {
		ctx.AgregarError("Error: Error al buscar la Particion : "+err.Error(), f.Linea, f.Columna)
		return
	}
	if !existe {
		ctx.AgregarError("Error: No existe una particion con el nombre "+name+" para modificar", f.Linea, f.Columna)
		return
	}
	//Calcular add de la particion (unit = B|K|M)
	if unit == "B" {
		add_ = add_ * 1
	} else if unit == "K" {
		add_ = add_ * 1024
	} else if unit == "M" {
		add_ = add_ * 1024 * 1024
	} else {
		ctx.AgregarError("Error: El valor del parametro -unit no es valido", f.Linea, f.Columna)
		return
	}
	//Si es primaria o extendida
	if particion != nil {
		//SizeMax
		var sizeMax int64
		if posPart < (len(mbr.Particiones) - 1) {
			partSig := mbr.Particiones[posPart+1]
			sizeMax = partSig.Part_start
		} else {
			sizeMax = mbr.Mbr_tamanio
		}
		//Verificar si se puede modificar la particion
		if add_ > 0 && particion.Part_start+particion.Part_s+int64(add_) <= sizeMax {
			particion.Part_s = particion.Part_s + int64(add_)
		} else if add_ < 0 && particion.Part_start+particion.Part_s+int64(add_) > particion.Part_start {
			particion.Part_s = particion.Part_s + int64(add_)
		} else {
			if add_ > 0 {
				ctx.AgregarError("Error: No hay espacio para agregar espacio a la Particion", f.Linea, f.Columna)
				return

			}
			ctx.AgregarError("Error: El tamaño a reducir es mas grande que la particion misma", f.Linea, f.Columna)
			return
		}
		//Escribir MBR en el disco
		err = writeMBRInDisk(drivelleter, mbr)
		if err != nil {
			ctx.AgregarError("Error: No se pudo escribir el MBR modificado en el disco : "+err.Error(), f.Linea, f.Columna)
			return
		}
	}
	//Si es logica
	if ebr != nil {
		//Buscar particion extendida
		extendida := mbr.BuscarParticionExtendida()
		if extendida == nil {
			ctx.AgregarError("Error: No se encontro una particion extendida", f.Linea, f.Columna)
			return
		}
		//SizeMax
		var sizeMax int64
		if ebr.Part_next == -1 {
			sizeMax = extendida.Part_start + extendida.Part_s
		} else {
			sizeMax = ebr.Part_next
		}
		//Verificar si se puede modificar la particion
		if add_ > 0 && ebr.Part_start+ebr.Part_s+int64(add_) <= sizeMax {
			ebr.Part_s = ebr.Part_s + int64(add_)
		} else if add_ < 0 && ebr.Part_start+ebr.Part_s+int64(add_) > ebr.Part_start {
			ebr.Part_s = ebr.Part_s + int64(add_)
		} else {
			ctx.AgregarError("Error: No se puede modificar la particion logica", f.Linea, f.Columna)
			return
		}
		//Escribir EBR en el disco
		err = extendida.WriteEBRsInDisk(drivelleter, []structs.EBR{*ebr})
		if err != nil {
			ctx.AgregarError("Error: No se pudo escribir el EBR modificado en el disco : "+err.Error(), f.Linea, f.Columna)
			return
		}
	}
	//Mensaje de exito
	//fmt.Println("----------Comando FDISK----------")
	//fmt.Println("Particion \"" + name + "\" modificada en el disco " + drivelleter + " con exito")
	ctx.AgregarOutput("----------Comando FDISK----------")
	ctx.AgregarOutput("Particion \"" + name + "\" modificada en el disco " + drivelleter + " con exito")
}
