package comandos

import (
	"fmt"
	"strings"
)

type Chgrp struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// chgrp *Chgrp ExprCommand
func (chgrp *Chgrp) Ejecutar(ctx *Contexto) interface{} {
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Veiricar que el parametro no tenga parametros
	if len(chgrp.Parametros) == 0 {
		ctx.AgregarError("Error: El comando Chgrp no tiene parametros", chgrp.Linea, chgrp.Columna)
		return nil
	}
	//Verificar parametros obligatorios -user -grp
	if _, ok := chgrp.Parametros["-user"]; !ok {
		ctx.AgregarError("Error: El comando Chgrp necesita el parametro -user ", chgrp.Linea, chgrp.Columna)
		return nil
	}
	if _, ok := chgrp.Parametros["-grp"]; !ok {
		ctx.AgregarError("Error: El comando Chgrp necesita el parametro -grp ", chgrp.Linea, chgrp.Columna)
		return nil
	}
	//Obtener valores de los parametros
	usuario := chgrp.Parametros["-user"]
	grupo := chgrp.Parametros["-grp"]
	//Verificar que haya sesion iniciada
	if !SesionActiva.HaySesionActiva() {
		ctx.AgregarError("Error: No hay sesion Activa", chgrp.Linea, chgrp.Columna)
	}
	//Recuperar Sesion Activa
	_, _, idPart := SesionActiva.SesionActual()
	//Buscar Particion Montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(idPart)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+idPart+"\" no esta montada", chgrp.Linea, chgrp.Columna)
		return nil
	}
	//Recuperar SuperBloque
	superBloque, err := getSuperBloque(ctx, partMontada.DiskName, partMontada.PartName)
	if err != nil {
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Recuperar usuarios y grupos de la particion
	usuarios, grupos, err := superBloque.RecuperarUsuariosGrupos(ctx, partMontada.DiskName)
	if err != nil {
		return nil
	}
	//Verificar si hay errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar que el grupo exista
	existeGrup := false
	for _, gr := range grupos {
		if gr.Grupo == grupo && gr.GID != "0" {
			existeGrup = true
			break
		}
	}
	if !existeGrup {
		ctx.AgregarError("Error: El grupo: "+grupo+" no Existe", chgrp.Linea, chgrp.Columna)
		return nil
	}
	//Verificar que el usuario exista
	existeUs := false
	for pos, us := range usuarios {
		if us.Usuario == usuario && us.UID != "0" {
			//Cambiar grupo
			usuarios[pos].Grupo = grupo
			existeUs = true
			break
		}
	}
	if !existeUs {
		ctx.AgregarError("Error: El usuario: "+usuario+" no Existe", chgrp.Linea, chgrp.Columna)
		return nil
	}
	//Convertir usuarios a []byte
	usersGroups := ConvertirUsuariosGrupos(usuarios, grupos)
	//Escribir usuarios y grupos
	pathStr := "/users.txt"
	path := strings.Split(pathStr, "/")
	//Eliminar el primer elemento
	path = path[1:]
	//Escribir el archivo users.txt
	superBloque.CrearYEscribirEnArchivo(ctx, partMontada.DiskName, path, usersGroups, false)
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo cambiar el grupo del usuario", chgrp.Linea, chgrp.Columna)
		return nil
	}
	//Escribir superBloque en la particion
	superBloque.EscribirSuperBloque(ctx, partMontada.DiskName, partMontada.Start_SuperBloque)
	//Verificar si hay errores
	if ctx.HayErrores() {
		ctx.AgregarError("Error: No se pudo cambiar el grupo del usuario", chgrp.Linea, chgrp.Columna)
		return nil
	}
	//Mensaje de exito
	fmt.Println("----------Comando CHGRP--------------")
	fmt.Println("Grupo del usuario: " + usuario + " cambiado a: " + grupo)
	return nil
}
