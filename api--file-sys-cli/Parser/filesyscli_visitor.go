// Code generated from FileSysCLI.g4 by ANTLR 4.13.1. DO NOT EDIT.

package Parser // FileSysCLI
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by FileSysCLIParser.
type FileSysCLIVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FileSysCLIParser#commands.
	VisitCommands(ctx *CommandsContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#command.
	VisitCommand(ctx *CommandContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_execute.
	VisitCommand_execute(ctx *Command_executeContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_mkdisk.
	VisitCommand_mkdisk(ctx *Command_mkdiskContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_rmdisk.
	VisitCommand_rmdisk(ctx *Command_rmdiskContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_fdisk.
	VisitCommand_fdisk(ctx *Command_fdiskContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_mount.
	VisitCommand_mount(ctx *Command_mountContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_unmount.
	VisitCommand_unmount(ctx *Command_unmountContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_mkfs.
	VisitCommand_mkfs(ctx *Command_mkfsContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_login.
	VisitCommand_login(ctx *Command_loginContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_logout.
	VisitCommand_logout(ctx *Command_logoutContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_mkgrp.
	VisitCommand_mkgrp(ctx *Command_mkgrpContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_rmgrp.
	VisitCommand_rmgrp(ctx *Command_rmgrpContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_mkusr.
	VisitCommand_mkusr(ctx *Command_mkusrContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_rmusr.
	VisitCommand_rmusr(ctx *Command_rmusrContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_mkfile.
	VisitCommand_mkfile(ctx *Command_mkfileContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_rep.
	VisitCommand_rep(ctx *Command_repContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_cat.
	VisitCommand_cat(ctx *Command_catContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_rm.
	VisitCommand_rm(ctx *Command_rmContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_edit.
	VisitCommand_edit(ctx *Command_editContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_rename.
	VisitCommand_rename(ctx *Command_renameContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_mkdir.
	VisitCommand_mkdir(ctx *Command_mkdirContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_pause.
	VisitCommand_pause(ctx *Command_pauseContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_chgrp.
	VisitCommand_chgrp(ctx *Command_chgrpContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_move.
	VisitCommand_move(ctx *Command_moveContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_loss.
	VisitCommand_loss(ctx *Command_lossContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_recovery.
	VisitCommand_recovery(ctx *Command_recoveryContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_chown.
	VisitCommand_chown(ctx *Command_chownContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Command_chmod.
	VisitCommand_chmod(ctx *Command_chmodContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by FileSysCLIParser#Value_parameter.
	VisitValue_parameter(ctx *Value_parameterContext) interface{}
}
