// Generated from /home/ralda/Estudios/lab_archivos/proyecto_1/MIA_P1_202103216/GenerateParser/FileSysCLI.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link FileSysCLIParser}.
 */
public interface FileSysCLIListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link FileSysCLIParser#commands}.
	 * @param ctx the parse tree
	 */
	void enterCommands(FileSysCLIParser.CommandsContext ctx);
	/**
	 * Exit a parse tree produced by {@link FileSysCLIParser#commands}.
	 * @param ctx the parse tree
	 */
	void exitCommands(FileSysCLIParser.CommandsContext ctx);
	/**
	 * Enter a parse tree produced by {@link FileSysCLIParser#command}.
	 * @param ctx the parse tree
	 */
	void enterCommand(FileSysCLIParser.CommandContext ctx);
	/**
	 * Exit a parse tree produced by {@link FileSysCLIParser#command}.
	 * @param ctx the parse tree
	 */
	void exitCommand(FileSysCLIParser.CommandContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Command_execute}
	 * labeled alternative in {@link FileSysCLIParser#execute}.
	 * @param ctx the parse tree
	 */
	void enterCommand_execute(FileSysCLIParser.Command_executeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Command_execute}
	 * labeled alternative in {@link FileSysCLIParser#execute}.
	 * @param ctx the parse tree
	 */
	void exitCommand_execute(FileSysCLIParser.Command_executeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Command_mkdisk}
	 * labeled alternative in {@link FileSysCLIParser#mkdisk}.
	 * @param ctx the parse tree
	 */
	void enterCommand_mkdisk(FileSysCLIParser.Command_mkdiskContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Command_mkdisk}
	 * labeled alternative in {@link FileSysCLIParser#mkdisk}.
	 * @param ctx the parse tree
	 */
	void exitCommand_mkdisk(FileSysCLIParser.Command_mkdiskContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Command_rep}
	 * labeled alternative in {@link FileSysCLIParser#rep}.
	 * @param ctx the parse tree
	 */
	void enterCommand_rep(FileSysCLIParser.Command_repContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Command_rep}
	 * labeled alternative in {@link FileSysCLIParser#rep}.
	 * @param ctx the parse tree
	 */
	void exitCommand_rep(FileSysCLIParser.Command_repContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Parameter}
	 * labeled alternative in {@link FileSysCLIParser#parametro}.
	 * @param ctx the parse tree
	 */
	void enterParameter(FileSysCLIParser.ParameterContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Parameter}
	 * labeled alternative in {@link FileSysCLIParser#parametro}.
	 * @param ctx the parse tree
	 */
	void exitParameter(FileSysCLIParser.ParameterContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Value_parameter}
	 * labeled alternative in {@link FileSysCLIParser#value_parametro}.
	 * @param ctx the parse tree
	 */
	void enterValue_parameter(FileSysCLIParser.Value_parameterContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Value_parameter}
	 * labeled alternative in {@link FileSysCLIParser#value_parametro}.
	 * @param ctx the parse tree
	 */
	void exitValue_parameter(FileSysCLIParser.Value_parameterContext ctx);
}