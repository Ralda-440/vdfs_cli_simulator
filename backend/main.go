package main

import (
	comandos "app/Interprete/Comandos"
	"app/Parser"
	tools "app/Tools"
	"encoding/base64"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//Verificar que existan los directorios donde se guardan los discos (MIA/P1) y si no existe crearlo
	_, err := os.Stat("./MIA/P1")
	if os.IsNotExist(err) {
		os.MkdirAll("./MIA/P1", os.ModePerm)
	}
	//Verificar que exista el directorio donde se guardan los Reportes (REP) y si no existe crearlo
	_, err = os.Stat("./REP")
	if os.IsNotExist(err) {
		os.MkdirAll("./REP", os.ModePerm)
	}
	//Login Activo
	loginActivo := tools.NewLoginActivo()
	//Contexto de la Consola
	ctx := comandos.NewContexto()
	//Servidor
	app := fiber.New()

	app.Use(cors.New())

	//Endpoint Server Ready
	app.Get("/ready", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"ready": true,
		})
	})

	//Input CLI
	app.Post("/inputCLI", func(c *fiber.Ctx) error {
		input := tools.InputCLI{}

		if err := c.BodyParser(&input); err != nil {
			return c.JSON(&fiber.Map{
				"output": "",
				"errs":   []comandos.Error{{Msg: err.Error(), Linea: 0, Columna: 0}},
			})
		}

		//fmt.Println(input.Input)

		//Nuevo Analizador
		lexer := Parser.NewFileSysCLILexer(antlr.NewInputStream(input.Input))
		lexer.RemoveErrorListeners()
		errorListener := &tools.ErrorListener_{Ctx: ctx}
		lexer.AddErrorListener(errorListener)
		tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		parser := Parser.NewFileSysCLIParser(tokenStream)
		parser.RemoveErrorListeners()
		errorListener = &tools.ErrorListener_{Ctx: ctx}
		parser.AddErrorListener(errorListener)
		raiz := parser.Commands()
		visitor := &comandos.FileSysCLIVisitor_{}
		interpreter := visitor.Visit(raiz).(comandos.ExprCommand)
		interpreter.Ejecutar(ctx)
		return c.JSON(&fiber.Map{
			"output": ctx.GetOutput(),
			"errs":   ctx.GetErrores(),
		})
	})

	//Reset CLI
	app.Get("/resetCLI", func(c *fiber.Ctx) error {
		//Eliminar las carpetas "MIA/P1" y "REP"
		errs := []string{}
		err := os.RemoveAll("./MIA")
		if err != nil {
			errs = append(errs, err.Error())
		}
		err = os.RemoveAll("./REP")
		if err != nil {
			errs = append(errs, err.Error())
		}
		//Crear las carpetas "MIA/P1" y "REP"
		err = os.MkdirAll("./MIA/P1", os.ModePerm)
		if err != nil {
			errs = append(errs, err.Error())
		}
		err = os.MkdirAll("./REP", os.ModePerm)
		if err != nil {
			errs = append(errs, err.Error())
		}
		ctx = comandos.NewContexto()
		loginActivo = tools.NewLoginActivo()
		comandos.SesionActiva = comandos.NewSesionActiva()
		comandos.ParticionesMontadas = comandos.NewPartMounts()
		comandos.CountPartLogicas = 0
		return c.JSON(&fiber.Map{
			"errs": errs,
		})
	})

	//Explorer
	app.Post("/explorer", func(c *fiber.Ctx) error {
		path := tools.PathExplorer{}

		if err := c.BodyParser(&path); err != nil {
			return c.JSON(&fiber.Map{
				"content": nil,
				"errs":    []comandos.Error{{Msg: err.Error(), Linea: 0, Columna: 0}},
			})
		}
		slicePath := strings.Split(path.Path, "/")
		//Eliminar el primer elemento
		slicePath = slicePath[1:]
		//Verificar si la ultima posicion es ""
		if slicePath[len(slicePath)-1] == "" {
			slicePath = slicePath[:len(slicePath)-1]
		}
		//Si la longitud es 0, entonces se esta pidiendo los discos creados
		if len(slicePath) == 0 {
			discos := tools.DiscosCreados(ctx)
			if ctx.HayErrores() {
				return c.JSON(&fiber.Map{
					"content": nil,
					"errs":    ctx.GetErrores(),
				})
			}
			return c.JSON(&fiber.Map{
				"content": discos,
				"errs":    nil,
			})
		} else if len(slicePath) == 1 {
			// Si la longitud es 1, entonces se esta pidiendo las particiones de un disco
			driveletter := strings.Split(slicePath[0], ".")[0]
			mbrDisk, err := comandos.GetMBRDisk(driveletter)
			if err != nil {
				ctx.AgregarError(err.Error(), 0, 0)
				return c.JSON(&fiber.Map{
					"content": nil,
					"errs":    ctx.GetErrores(),
				})
			}
			particiones, err := mbrDisk.GetParticiones(driveletter)
			if err != nil {
				ctx.AgregarError(err.Error(), 0, 0)
				return c.JSON(&fiber.Map{
					"content": nil,
					"errs":    ctx.GetErrores(),
				})
			}
			return c.JSON(&fiber.Map{
				"content": particiones,
				"errs":    nil,
			})
		} else if len(slicePath) >= 2 {
			// Si la longitud es 2, entonces se esta pidiendo los archivos de una particion
			driveletter := strings.Split(slicePath[0], ".")[0]
			partname := slicePath[1]
			superBloque, err := comandos.GetSuperBloque(ctx, driveletter, partname)
			if err != nil {
				ctx.AgregarError(err.Error(), 0, 0)
				return c.JSON(&fiber.Map{
					"content": nil,
					"errs":    ctx.GetErrores(),
				})
			}
			//Verificar si hay errores
			if ctx.HayErrores() {
				return c.JSON(&fiber.Map{
					"content": nil,
					"errs":    ctx.GetErrores(),
				})
			}
			//[]path sin el nombre del disco y de la particion
			path_ := slicePath[2:]
			_, content := superBloque.ReporteLs(ctx, driveletter, path_)
			//Verificar si hay errores
			if ctx.HayErrores() {
				return c.JSON(&fiber.Map{
					"content": nil,
					"errs":    ctx.GetErrores(),
				})
			}
			return c.JSON(&fiber.Map{
				"content": content,
				"errs":    nil,
			})
		}
		ctx.AgregarError("Error: No se pudo leer el path", 0, 0)
		return c.JSON(&fiber.Map{
			"content": nil,
			"errs":    ctx.GetErrores(),
		})
	})

	app.Get("/loginActivo", func(c *fiber.Ctx) error {
		return c.JSON(loginActivo)
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		loginActivo = tools.NewLoginActivo()
		return c.JSON(&fiber.Map{
			"errs": nil,
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		body := &fiber.Map{
			"partName": "",
			"diskName": "",
			"usr":      "",
			"pwd":      "",
		}
		err := c.BodyParser(body)
		if err != nil {
			return c.JSON(&fiber.Map{
				"errs": []comandos.Error{{Msg: err.Error()}},
			})
		}
		//Buscar SuperBloque
		superBloque, err := comandos.GetSuperBloque(ctx, (*body)["diskName"].(string), (*body)["partName"].(string))
		if err != nil {
			return c.JSON(&fiber.Map{
				"errs": []comandos.Error{{Msg: err.Error()}},
			})
		}
		//Verificar si hay errores
		if ctx.HayErrores() {
			return c.JSON(&fiber.Map{
				"errs": ctx.GetErrores(),
			})
		}
		//Recuperar usuarios y grupos de la particion
		usuarios, _, err := superBloque.RecuperarUsuariosGrupos(ctx, (*body)["diskName"].(string))
		if err != nil {
			return c.JSON(&fiber.Map{
				"errs": []comandos.Error{{Msg: err.Error()}},
			})
		}
		//Verificar si hay
		if ctx.HayErrores() {
			return c.JSON(&fiber.Map{
				"errs": ctx.GetErrores(),
			})
		}
		//Verificar si el usuario y contraseña existen
		for _, usr := range usuarios {
			if usr.Usuario == (*body)["usr"].(string) {
				if usr.Pass == (*body)["pwd"].(string) {
					loginActivo.ActivarLogin((*body)["partName"].(string), (*body)["diskName"].(string))
					return c.JSON(&fiber.Map{
						"errs": nil,
					})
				}
			}
		}
		return c.JSON(&fiber.Map{
			"errs": []comandos.Error{{Msg: "Usuario o Contraseña incorrectos"}},
		})
	})

	//Contenido file
	app.Post("/contentFile", func(c *fiber.Ctx) error {
		path := tools.PathExplorer{}

		if err := c.BodyParser(&path); err != nil {
			return c.JSON(&fiber.Map{
				"content": nil,
				"errs":    []comandos.Error{{Msg: err.Error(), Linea: 0, Columna: 0}},
			})
		}
		slicePath := strings.Split(path.Path, "/")
		//Eliminar el primer elemento
		slicePath = slicePath[1:]
		//Verificar si la ultima posicion es ""
		if slicePath[len(slicePath)-1] == "" {
			slicePath = slicePath[:len(slicePath)-1]
		}
		// Si la longitud es 2, entonces se esta pidiendo los archivos de una particion
		driveletter := strings.Split(slicePath[0], ".")[0]
		partname := slicePath[1]
		superBloque, err := comandos.GetSuperBloque(ctx, driveletter, partname)
		if err != nil {
			ctx.AgregarError(err.Error(), 0, 0)
			return c.JSON(&fiber.Map{
				"content": nil,
				"errs":    ctx.GetErrores(),
			})
		}
		//Verificar si hay errores
		if ctx.HayErrores() {
			return c.JSON(&fiber.Map{
				"content": nil,
				"errs":    ctx.GetErrores(),
			})
		}
		//[]path sin el nombre del disco y de la particion
		path_ := slicePath[2:]
		content := superBloque.LeerContenidoArchivo(ctx, driveletter, path_)
		//Verificar si hay errores
		if ctx.HayErrores() {
			return c.JSON(&fiber.Map{
				"content": nil,
				"errs":    ctx.GetErrores(),
			})
		}
		//Convertir []byte a string
		contentStr := string(content)
		return c.JSON(&fiber.Map{
			"content": contentStr,
			"errs":    nil,
		})
	})

	app.Get("/reportes", func(c *fiber.Ctx) error {
		//Leer los archivos de la carpeta REP
		files, err := os.ReadDir("./REP")
		if err != nil {
			ctx.AgregarError("Error: "+err.Error(), 0, 0)
			return c.JSON(&fiber.Map{
				"files": nil,
				"errs":  ctx.GetErrores(),
			})
		}
		//Lista de Reportes
		reportes := make([]tools.ItemReport, 0)
		//Recorrer los archivo
		for key, file := range files {
			//Verificar si es un archivo
			if !file.IsDir() {
				//Obtener el nombre del archivo
				nameReport := file.Name()
				//Obtener base64
				fileContent, err := os.ReadFile("./REP/" + nameReport)
				if err != nil {
					ctx.AgregarError("Error: "+err.Error(), 0, 0)
					return c.JSON(&fiber.Map{
						"files": nil,
						"errs":  ctx.GetErrores(),
					})
				}
				//Obtener la extension
				ext := filepath.Ext(nameReport)
				//Verificar si es un archivo de imagen o plano
				var tipo string
				var stringContent string
				if ext == ".png" || ext == ".jpg" {
					tipo = "image"
					stringContent = "data:@file/png;base64," + base64.StdEncoding.EncodeToString(fileContent)
				} else if ext == ".txt" {
					tipo = "plain"
					stringContent = string(fileContent)
				} else {
					tipo = "other"
					stringContent = ""
				}
				//Agregar a la lista de reportes
				reportes = append(reportes, tools.ItemReport{
					Nombre:  nameReport,
					Tipo:    tipo,
					Content: stringContent,
					Key:     strconv.Itoa(key),
				})
			}
		}
		return c.JSON(&fiber.Map{
			"files": reportes,
			"errs":  nil,
		})
	})

	app.Listen(":4005")
}
