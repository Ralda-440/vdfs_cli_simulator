package comandos

import (
	"app/Parser"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type FileSysCLIVisitor_ struct{}

//FSCLI *FileSysCLIVisitor_ Parser.FileSysCLIVisitor

func (FSCLI *FileSysCLIVisitor_) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(FSCLI)
}

func (FSCLI *FileSysCLIVisitor_) VisitChildren(node antlr.RuleNode) interface{} {
	panic("not implemented") // TODO: Implement
}

func (FSCLI *FileSysCLIVisitor_) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("not implemented") // TODO: Implement
}

func (FSCLI *FileSysCLIVisitor_) VisitErrorNode(node antlr.ErrorNode) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by FileSysCLIParser#commands.
func (FSCLI *FileSysCLIVisitor_) VisitCommands(ctx *Parser.CommandsContext) interface{} {
	commands := &Commands{Commands: make([]ExprCommand, 0)}
	//Verificar si hay comandos
	if ctx.AllCommand() == nil {
		return commands
	}
	for _, command := range ctx.AllCommand() {
		commands.Commands = append(commands.Commands, command.Accept(FSCLI).(ExprCommand))
	}
	return commands
}

// Visit a parse tree produced by FileSysCLIParser#command.
func (FSCLI *FileSysCLIVisitor_) VisitCommand(ctx *Parser.CommandContext) interface{} {
	if ctx.Execute() != nil {
		return ctx.Execute().Accept(FSCLI)
	} else if ctx.Mkdisk() != nil {
		return ctx.Mkdisk().Accept(FSCLI)
	} else if ctx.Rep() != nil {
		return ctx.Rep().Accept(FSCLI)
	} else if ctx.Rmdisk() != nil {
		return ctx.Rmdisk().Accept(FSCLI)
	} else if ctx.Fdisk() != nil {
		return ctx.Fdisk().Accept(FSCLI)
	} else if ctx.Mount() != nil {
		return ctx.Mount().Accept(FSCLI)
	} else if ctx.Unmount() != nil {
		return ctx.Unmount().Accept(FSCLI)
	} else if ctx.Mkfs() != nil {
		return ctx.Mkfs().Accept(FSCLI)
	} else if ctx.Login() != nil {
		return ctx.Login().Accept(FSCLI)
	} else if ctx.Logout() != nil {
		return ctx.Logout().Accept(FSCLI)
	} else if ctx.Mkgrp() != nil {
		return ctx.Mkgrp().Accept(FSCLI)
	} else if ctx.Rmgrp() != nil {
		return ctx.Rmgrp().Accept(FSCLI)
	} else if ctx.Mkusr() != nil {
		return ctx.Mkusr().Accept(FSCLI)
	} else if ctx.Rmusr() != nil {
		return ctx.Rmusr().Accept(FSCLI)
	} else if ctx.Mkfile() != nil {
		return ctx.Mkfile().Accept(FSCLI)
	} else if ctx.Cat() != nil {
		return ctx.Cat().Accept(FSCLI)
	} else if ctx.Remove() != nil {
		return ctx.Remove().Accept(FSCLI)
	} else if ctx.Edit() != nil {
		return ctx.Edit().Accept(FSCLI)
	} else if ctx.Rename() != nil {
		return ctx.Rename().Accept(FSCLI)
	} else if ctx.Mkdir() != nil {
		return ctx.Mkdir().Accept(FSCLI)
	} else if ctx.Pause() != nil {
		return ctx.Pause().Accept(FSCLI)
	} else if ctx.Chgrp() != nil {
		return ctx.Chgrp().Accept(FSCLI)
	} else if ctx.Move() != nil {
		return ctx.Move().Accept(FSCLI)
	} else if ctx.Loss() != nil {
		return ctx.Loss().Accept(FSCLI)
	} else if ctx.Recovery() != nil {
		return ctx.Recovery().Accept(FSCLI)
	} else if ctx.Chown() != nil {
		return ctx.Chown().Accept(FSCLI)
	} else if ctx.Chmod() != nil {
		return ctx.Chmod().Accept(FSCLI)
	} else {
		return nil
	}
}

// Visit a parse tree produced by FileSysCLIParser#Command_execute.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_execute(ctx *Parser.Command_executeContext) interface{} {
	execute := &Execute{Parametros: make(map[string]string, 0)}
	execute.Linea = int(ctx.GetStart().GetLine())
	execute.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return execute
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			execute.Parametros[clave] = valor
		}
	}
	return execute
}

// Visit a parse tree produced by FileSysCLIParser#Command_mkdisk.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_mkdisk(ctx *Parser.Command_mkdiskContext) interface{} {
	mkd := &Mkdisk{Parametros: make(map[string]string, 0)}
	mkd.Linea = int(ctx.GetStart().GetLine())
	mkd.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return mkd
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			mkd.Parametros[clave] = valor
		}
	}
	return mkd
}

// Visit a parse tree produced by FileSysCLIParser#Command_rmdisk.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_rmdisk(ctx *Parser.Command_rmdiskContext) interface{} {
	rmd := &Rmdisk{Parametros: make(map[string]string, 0)}
	rmd.Linea = int(ctx.GetStart().GetLine())
	rmd.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return rmd
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			rmd.Parametros[clave] = valor
		}
	}
	return rmd
}

// Visit a parse tree produced by FileSysCLIParser#Command_fdisk.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_fdisk(ctx *Parser.Command_fdiskContext) interface{} {
	fdk := &Fdisk{Parametros: make(map[string]string, 0)}
	fdk.Linea = int(ctx.GetStart().GetLine())
	fdk.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return fdk
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			fdk.Parametros[clave] = valor
		}
	}
	return fdk
}

// Visit a parse tree produced by FileSysCLIParser#Command_mount.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_mount(ctx *Parser.Command_mountContext) interface{} {
	mnt := &Mount{Parametros: make(map[string]string, 0)}
	mnt.Linea = int(ctx.GetStart().GetLine())
	mnt.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return mnt
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			mnt.Parametros[clave] = valor
		}
	}
	return mnt
}

// Visit a parse tree produced by FileSysCLIParser#Command_unmount.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_unmount(ctx *Parser.Command_unmountContext) interface{} {
	umt := &Unmount{Parametros: make(map[string]string, 0)}
	umt.Linea = int(ctx.GetStart().GetLine())
	umt.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return umt
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			umt.Parametros[clave] = valor
		}
	}
	return umt
}

// Visit a parse tree produced by FileSysCLIParser#Command_mkfs.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_mkfs(ctx *Parser.Command_mkfsContext) interface{} {
	mks := &Mkfs{Parametros: make(map[string]string, 0)}
	mks.Linea = int(ctx.GetStart().GetLine())
	mks.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return mks
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			mks.Parametros[clave] = valor
		}
	}
	return mks
}

// Visit a parse tree produced by FileSysCLIParser#Command_login.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_login(ctx *Parser.Command_loginContext) interface{} {
	lg := &Login{Parametros: make(map[string]string, 0)}
	lg.Linea = int(ctx.GetStart().GetLine())
	lg.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return lg
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			lg.Parametros[clave] = valor
		}
	}
	return lg
}

// Visit a parse tree produced by FileSysCLIParser#Command_logout.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_logout(ctx *Parser.Command_logoutContext) interface{} {
	lgt := &Logout{Parametros: make(map[string]string, 0)}
	lgt.Linea = int(ctx.GetStart().GetLine())
	lgt.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return lgt
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			lgt.Parametros[clave] = valor
		}
	}
	return lgt
}

// Visit a parse tree produced by FileSysCLIParser#Command_mkgrp.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_mkgrp(ctx *Parser.Command_mkgrpContext) interface{} {
	mkgrp := &Mkgrp{Parametros: make(map[string]string, 0)}
	mkgrp.Linea = int(ctx.GetStart().GetLine())
	mkgrp.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return mkgrp
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			mkgrp.Parametros[clave] = valor
		}
	}
	return mkgrp
}

// Visit a parse tree produced by FileSysCLIParser#Command_rmgrp.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_rmgrp(ctx *Parser.Command_rmgrpContext) interface{} {
	rmgrp := &Rmgrp{Parametros: make(map[string]string, 0)}
	rmgrp.Linea = int(ctx.GetStart().GetLine())
	rmgrp.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return rmgrp
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			rmgrp.Parametros[clave] = valor
		}
	}
	return rmgrp
}

// Visit a parse tree produced by FileSysCLIParser#Command_mkusr.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_mkusr(ctx *Parser.Command_mkusrContext) interface{} {
	mkusr := &Mkusr{Parametros: make(map[string]string, 0)}
	mkusr.Linea = int(ctx.GetStart().GetLine())
	mkusr.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return mkusr
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			mkusr.Parametros[clave] = valor
		}
	}
	return mkusr
}

// Visit a parse tree produced by FileSysCLIParser#Command_rmusr.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_rmusr(ctx *Parser.Command_rmusrContext) interface{} {
	rmusr := &Rmusr{Parametros: make(map[string]string, 0)}
	rmusr.Linea = int(ctx.GetStart().GetLine())
	rmusr.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return rmusr
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			rmusr.Parametros[clave] = valor
		}
	}
	return rmusr
}

// Visit a parse tree produced by FileSysCLIParser#Command_mkfile.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_mkfile(ctx *Parser.Command_mkfileContext) interface{} {
	mkf := &Mkfile{Parametros: make(map[string]string, 0)}
	mkf.Linea = int(ctx.GetStart().GetLine())
	mkf.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return mkf
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			mkf.Parametros[clave] = valor
		}
	}
	return mkf
}

// Visit a parse tree produced by FileSysCLIParser#Command_cat.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_cat(ctx *Parser.Command_catContext) interface{} {
	cat := &Cat{Parametros: make(map[string]string, 0)}
	cat.Linea = int(ctx.GetStart().GetLine())
	cat.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return cat
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			cat.Parametros[clave] = valor
		}
	}
	return cat
}

// Visit a parse tree produced by FileSysCLIParser#Command_rm.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_rm(ctx *Parser.Command_rmContext) interface{} {
	remove := &Remove{Parametros: make(map[string]string, 0)}
	remove.Linea = int(ctx.GetStart().GetLine())
	remove.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return remove
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			remove.Parametros[clave] = valor
		}
	}
	return remove
}

// Visit a parse tree produced by FileSysCLIParser#Command_edit.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_edit(ctx *Parser.Command_editContext) interface{} {
	ed := &Edit{Parametros: make(map[string]string, 0)}
	ed.Linea = int(ctx.GetStart().GetLine())
	ed.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return ed
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			ed.Parametros[clave] = valor
		}
	}
	return ed
}

// Visit a parse tree produced by FileSysCLIParser#Command_rename.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_rename(ctx *Parser.Command_renameContext) interface{} {
	rename := &Rename{Parametros: make(map[string]string, 0)}
	rename.Linea = int(ctx.GetStart().GetLine())
	rename.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return rename
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			rename.Parametros[clave] = valor
		}
	}
	return rename
}

// Visit a parse tree produced by FileSysCLIParser#Command_mkdir.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_mkdir(ctx *Parser.Command_mkdirContext) interface{} {
	mkdir := &Mkdir{Parametros: make(map[string]string, 0)}
	mkdir.Linea = int(ctx.GetStart().GetLine())
	mkdir.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return mkdir
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			mkdir.Parametros[clave] = valor
		}
	}
	return mkdir
}

// Visit a parse tree produced by FileSysCLIParser#Command_pause.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_pause(ctx *Parser.Command_pauseContext) interface{} {
	pause := &Pause{Parametros: make(map[string]string, 0)}
	pause.Linea = int(ctx.GetStart().GetLine())
	pause.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return pause
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			pause.Parametros[clave] = valor
		}
	}
	return pause
}

// Visit a parse tree produced by FileSysCLIParser#Command_chgrp.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_chgrp(ctx *Parser.Command_chgrpContext) interface{} {
	chgrp := &Chgrp{Parametros: make(map[string]string, 0)}
	chgrp.Linea = int(ctx.GetStart().GetLine())
	chgrp.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return chgrp
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			chgrp.Parametros[clave] = valor
		}
	}
	return chgrp
}

// Visit a parse tree produced by FileSysCLIParser#Command_rep.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_rep(ctx *Parser.Command_repContext) interface{} {
	rep := &REP{Parametros: make(map[string]string, 0)}
	rep.Linea = int(ctx.GetStart().GetLine())
	rep.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return rep
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			rep.Parametros[clave] = valor
		}
	}
	return rep
}

// Visit a parse tree produced by FileSysCLIParser#Command_move.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_move(ctx *Parser.Command_moveContext) interface{} {
	move := &Move{Parametros: make(map[string]string, 0)}
	move.Linea = int(ctx.GetStart().GetLine())
	move.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return move
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			move.Parametros[clave] = valor
		}
	}
	return move
}

// Visit a parse tree produced by FileSysCLIParser#Command_loss.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_loss(ctx *Parser.Command_lossContext) interface{} {
	loss := &Loss{Parametros: make(map[string]string, 0)}
	loss.Linea = int(ctx.GetStart().GetLine())
	loss.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return loss
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			loss.Parametros[clave] = valor
		}
	}
	return loss
}

// Visit a parse tree produced by FileSysCLIParser#Command_recovery.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_recovery(ctx *Parser.Command_recoveryContext) interface{} {
	recovery := &Recovery{Parametros: make(map[string]string, 0)}
	recovery.Linea = int(ctx.GetStart().GetLine())
	recovery.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return recovery
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			recovery.Parametros[clave] = valor
		}
	}
	return recovery
}

// Visit a parse tree produced by FileSysCLIParser#Command_chown.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_chown(ctx *Parser.Command_chownContext) interface{} {
	chown := &Chown{Parametros: make(map[string]string, 0)}
	chown.Linea = int(ctx.GetStart().GetLine())
	chown.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return chown
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			chown.Parametros[clave] = valor
		}
	}
	return chown
}

// Visit a parse tree produced by FileSysCLIParser#Command_chmod.
func (FSCLI *FileSysCLIVisitor_) VisitCommand_chmod(ctx *Parser.Command_chmodContext) interface{} {
	chmod := &Chmod{Parametros: make(map[string]string, 0)}
	chmod.Linea = int(ctx.GetStart().GetLine())
	chmod.Columna = int(ctx.GetStart().GetColumn())
	//Verificar si hay parametros
	if ctx.AllParametro() == nil {
		return chmod
	}
	for _, param := range ctx.AllParametro() {
		for clave, valor := range param.Accept(FSCLI).(map[string]string) {
			chmod.Parametros[clave] = valor
		}
	}
	return chmod
}

// Visit a parse tree produced by FileSysCLIParser#Parameter.
func (FSCLI *FileSysCLIVisitor_) VisitParameter(ctx *Parser.ParameterContext) interface{} {
	param := make(map[string]string, 0)
	valorParametro := ""
	if ctx.Value_parametro() != nil {
		valorParametro = ctx.Value_parametro().Accept(FSCLI).(string)
	}
	//Poner en minusculas el parametro

	param[strings.ToLower(ctx.QUOTED_TEXT_().GetText())] = valorParametro
	return param
}

// Visit a parse tree produced by FileSysCLIParser#Value_parameter.
func (FSCLI *FileSysCLIVisitor_) VisitValue_parameter(ctx *Parser.Value_parameterContext) interface{} {
	if ctx.STRING() != nil {
		//Quitar las comillas al principio y al final
		string_ := ctx.STRING().GetText()
		return string_[1 : len(string_)-1]
	} else if ctx.QUOTED_TEXT_() != nil {
		return ctx.QUOTED_TEXT_().GetText()
	} else {
		return nil
	}
}
