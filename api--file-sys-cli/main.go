package main

import (
	comandos "app/Interprete/Comandos"
	"app/Parser"
	tools "app/Tools"
	"os"
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
	//Contexto de la Consola
	ctx := comandos.NewContexto()
	//Servidor
	app := fiber.New()

	app.Use(cors.New())

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
		}
		return nil
	})

	app.Listen(":4005")
}
