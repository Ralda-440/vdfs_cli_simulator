// Code generated from FileSysCLI.g4 by ANTLR 4.13.1. DO NOT EDIT.

package Parser // FileSysCLI
import "github.com/antlr4-go/antlr/v4"

type BaseFileSysCLIVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFileSysCLIVisitor) VisitCommands(ctx *CommandsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand(ctx *CommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_execute(ctx *Command_executeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_mkdisk(ctx *Command_mkdiskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_rmdisk(ctx *Command_rmdiskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_fdisk(ctx *Command_fdiskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_mount(ctx *Command_mountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_unmount(ctx *Command_unmountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_mkfs(ctx *Command_mkfsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_login(ctx *Command_loginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_logout(ctx *Command_logoutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_mkgrp(ctx *Command_mkgrpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_rmgrp(ctx *Command_rmgrpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_mkusr(ctx *Command_mkusrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_rmusr(ctx *Command_rmusrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_mkfile(ctx *Command_mkfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_rep(ctx *Command_repContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_cat(ctx *Command_catContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_rm(ctx *Command_rmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_edit(ctx *Command_editContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_rename(ctx *Command_renameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_mkdir(ctx *Command_mkdirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_pause(ctx *Command_pauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_chgrp(ctx *Command_chgrpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_move(ctx *Command_moveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_loss(ctx *Command_lossContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_recovery(ctx *Command_recoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_chown(ctx *Command_chownContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitCommand_chmod(ctx *Command_chmodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFileSysCLIVisitor) VisitValue_parameter(ctx *Value_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}
