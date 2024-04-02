package comandos

import "fmt"

type Login struct {
	Parametros map[string]string
	Linea      int
	Columna    int
}

// lg *Login ExprCommand
func (lg *Login) Ejecutar(ctx *Contexto) interface{} {
	//Verificar que no haya errores
	if ctx.HayErrores() {
		return nil
	}
	//Verificar que haya parametros
	if len(lg.Parametros) == 0 {
		ctx.AgregarError("El comando login no tiene parametros", lg.Linea, lg.Columna)
		return nil
	}
	//Verificar parametros obligatorios -user -pass -id
	if _, existe := lg.Parametros["-user"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -user", lg.Linea, lg.Columna)
		return nil
	}
	if _, existe := lg.Parametros["-pass"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -pass", lg.Linea, lg.Columna)
		return nil
	}
	if _, existe := lg.Parametros["-id"]; !existe {
		ctx.AgregarError("Error: Falta el parametro -id", lg.Linea, lg.Columna)
		return nil
	}
	//Obtener parametros
	user := lg.Parametros["-user"]
	pass := lg.Parametros["-pass"]
	id := lg.Parametros["-id"]
	//Verificar que no haya una sesion activa
	if SesionActiva.Activa {
		ctx.AgregarError("Error: Ya hay una sesion activa", lg.Linea, lg.Columna)
		return nil
	}
	//Verificar que el Id de la particion exista y este montada
	existe, partMontada := ParticionesMontadas.BuscarParticionMontada(id)
	if !existe {
		ctx.AgregarError("Error: La particion con id \""+id+"\" no esta montada", lg.Linea, lg.Columna)
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
	//Verificar que el usuario exista
	for _, usuario := range usuarios {
		if usuario.UID == "0" {
			continue
		}
		if usuario.Usuario == user {
			//Verificar que la contraseña sea correcta
			if usuario.Pass == pass {
				GUI := ""
				//Recuperar el GID del usuario
				for _, grupo := range grupos {
					if grupo.Grupo == usuario.Grupo {
						GUI = grupo.GID
					}
				}
				//Iniciar sesion
				SesionActiva.IniciarSesion(ctx, user, pass, id, usuario.UID, GUI)
				//Mensaje de exito
				fmt.Println("----------Comando LOGIN--------------")
				fmt.Println("Sesion iniciada con exito")
				return nil
			}
			ctx.AgregarError("Error: La contraseña es incorrecta", lg.Linea, lg.Columna)
			return nil
		}
	}
	ctx.AgregarError("Error: El usuario no existe", lg.Linea, lg.Columna)
	return nil
}

// Sesion de Usuario
var SesionActiva Sesion

type Sesion struct {
	Usuario     string //Nombre de usuario
	Password    string //Contraseña
	IdParticion string //Id de la particion montada
	UID         string //Id del usuario
	GID         string //Id del grupo
	Activa      bool   //Sesion activa
}

// Iniciar Sesion
func (sesion *Sesion) IniciarSesion(ctx *Contexto, user string, pass string, id string, uid string, gid string) {
	//Verificar que no haya una sesion activa
	if sesion.HaySesionActiva() {
		ctx.AgregarError("Error: Ya hay una sesion activa", 0, 0)
		return
	}
	//Iniciar sesion
	sesion.Usuario = user
	sesion.Password = pass
	sesion.IdParticion = id
	sesion.UID = uid
	sesion.GID = gid
	sesion.Activa = true
}

// HaySesionActiva verifica si hay una sesion activa
func (sesion *Sesion) HaySesionActiva() bool {
	return SesionActiva.Activa
}

// CerrarSesion cierra la sesion activa
func (sesion *Sesion) CerrarSesion() {
	sesion.Usuario = ""
	sesion.Password = ""
	sesion.IdParticion = ""
	sesion.Activa = false
}

// Sesion Actual
func (sesion *Sesion) SesionActual() (string, string, string) {
	return sesion.Usuario, sesion.Password, sesion.IdParticion
}
