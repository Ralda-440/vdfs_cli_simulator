package comandos

import (
	"strconv"
	"strings"
)

type Mkgrp struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

//mkgrp *Mkgrp ExprCommand

func (mkgrp *Mkgrp) Ejecutar(ctx *Contexto) interface{} {
	//Verificar si hay error
	if ctx.HayErrores() {
		return nil
	}
	//Verificar si hay parametros
	if len(mkgrp.Parametros) == 0 {
		ctx.AgregarError("El comando mkgrp no tiene parametros", mkgrp.Linea, mkgrp.Columna)
		return nil
	}
	//Verifica parametro obligatorio -name
	if _, existe := mkgrp.Parametros["-name"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -name", mkgrp.Linea, mkgrp.Columna)
		return nil
	}
	//Obtener name
	name := mkgrp.Parametros["-name"]
	//Verificar si hay sesion activa
	if !SesionActiva.Activa {
		ctx.AgregarError("Error: No hay ninguna sesion activa", mkgrp.Linea, mkgrp.Columna)
		return nil
	}
	//Recuperar Sesion Actual
	user, _, idPart := SesionActiva.SesionActual()
	//Verificar que el usuario sea root
	if user != "root" {
		ctx.AgregarError("Error: Solo el usuario root puede crear grupos", mkgrp.Linea, mkgrp.Columna)
		return nil
	}
	//Buscar particion montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+idPart+"\" no esta montada", mkgrp.Linea, mkgrp.Columna)
		return nil
	}
	//Recuperar SuperBloque
	superBloque, err := getSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar el superBloque", mkgrp.Linea, mkgrp.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Recuperar usuarios y grupos de la particion
	users, grupos, err := superBloque.RecuperarUsuariosGrupos(ctx, partMontada.DiskName)
	if err != nil {
		ctx.AgregarError("Error: No se pudo recuperar los usuarios y grupos de la particion", mkgrp.Linea, mkgrp.Columna)
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar que no exista el grupo
	for _, grupo := range grupos {
		if grupo.Grupo == name && grupo.GID != "0" {
			ctx.AgregarError("Error: Ya existe un grupo con el nombre: "+name, mkgrp.Linea, mkgrp.Columna)
			return nil
		}
	}
	//Generar GID para el grupo
	gid := superBloque.GenerarGID(grupos)
	//Crear Grupo
	grupo := GrupoParticion{GID: strconv.Itoa(gid), Tipo: "G", Grupo: name}
	//Agregar Grupo a grupos
	grupos = append(grupos, grupo)
	//Convertir user y grupo a []byte
	cont := ConvertirUsuariosGrupos(users, grupos)
	//Recuperar el inodo del archivo users.txt
	pathStr := "/users.txt"
	path := strings.Split(pathStr, "/")
	//Eliminar el primer elemento
	path = path[1:]
	//Escribir el archivo users.txt
	superBloque.CrearYEscribirEnArchivo(ctx, partMontada.DiskName, path, cont, false)
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo crear el grupo", mkgrp.Linea, mkgrp.Columna)
		return nil
	}
	//Escribir SuperBloque en la particion
	superBloque.EscribirSuperBloque(ctx, partMontada.DiskName, partMontada.Start_SuperBloque)
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo crear el grupo", mkgrp.Linea, mkgrp.Columna)
		return nil
	}
	//Mensaje de exito
	//fmt.Println("----------Comando MKGRP--------------")
	//fmt.Println("Grupo creado con exito: " + name)
	ctx.AgregarOutput("--------------------Comando MKGRP--------------------")
	ctx.AgregarOutput("Grupo: \"" + name + "\" con GID: " + strconv.Itoa(gid) + " creado con exito en la particion con Id: " + idPart)
	return nil
}
