package comandos

import (
	"strings"
)

type Rmgrp struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// rmgrp *Rmgrp ExprCommand
func (rmgrp *Rmgrp) Ejecutar(ctx *Contexto) interface{} {
	//Verificar si hay error
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(rmgrp.Parametros) == 0 {
		ctx.AgregarError("El comando rmgrp no tiene parametros", rmgrp.Linea, rmgrp.Columna)
		return nil
	}
	//Verifica parametro obligatorio -name
	if _, existe := rmgrp.Parametros["-name"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -name", rmgrp.Linea, rmgrp.Columna)
		return nil
	}
	//Obtener name
	name := rmgrp.Parametros["-name"]
	//Verificar si hay sesion activa
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay ninguna sesion activa", rmgrp.Linea, rmgrp.Columna)
		return nil
	}
	//Recuperar Sesion Actual
	user, _, idPart := SesionActiva.SesionActual()
	//Verificar que el usuario sea root
	if user != "root" {
		ctx.AgregarError("Error: Solo el usuario root puede eliminar grupos", rmgrp.Linea, rmgrp.Columna)
		return nil
	}
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+idPart+"\" no esta montada", rmgrp.Linea, rmgrp.Columna)
		return nil
	}
	//Recuperar SuperBloque
	superBloque, err := getSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", rmgrp.Linea, rmgrp.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Recuperar usuarios y grupos de la particion
	users, grupos, err := superBloque.RecuperarUsuariosGrupos(ctx, partMontada.DiskName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar los usuarios y grupos", rmgrp.Linea, rmgrp.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Buscar grupo
	existe = false
	for i, g := range grupos {
		if g.Grupo == name && g.GID != "0" {
			//Colocar el GID en 0 para indicar que esta eliminado
			grupos[i].GID = "0"
			existe = true
			break
		}
	}
	if !existe {
		ctx.AgregarError("Error: No existe un grupo con el nombre: "+name, rmgrp.Linea, rmgrp.Columna)
		return nil
	}
	//Convertir Usuarios y Grupos a []byte
	cont := ConvertirUsuariosGrupos(users, grupos)
	//Escribir Usuarios y Grupos en user.txt
	pathStr := "/users.txt"
	path := strings.Split(pathStr, "/")
	//Eliminar el primer elemento
	path = path[1:]
	//Escribir el archivo users.txt
	superBloque.CrearYEscribirEnArchivo(ctx, partMontada.DiskName, path, cont, false)
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo eliminar el grupo", rmgrp.Linea, rmgrp.Columna)
		return nil
	}
	//Escribir SuperBloque en la particion
	superBloque.EscribirSuperBloque(ctx, partMontada.DiskName, partMontada.Start_SuperBloque)
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo eliminar el grupo", rmgrp.Linea, rmgrp.Columna)
		return nil
	}
	//Mensaje de Exito
	//fmt.Println("----------Comando RMGRP--------------")
	//fmt.Println("Grupo eliminado con exito")
	ctx.AgregarOutput("--------------------Comando RMGRP--------------------")
	ctx.AgregarOutput("Grupo: \"" + name + "\" en la particion con id: \"" + idPart + "\" eliminado con exito")
	return nil
}
