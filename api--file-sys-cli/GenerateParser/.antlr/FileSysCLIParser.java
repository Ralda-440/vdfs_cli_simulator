// Generated from /home/ralda/Estudios/lab_archivos/proyecto_2/MIA_P2_202103216/api--file-sys-cli/GenerateParser/FileSysCLI.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class FileSysCLIParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, ENBLANCO=29, LINE_COMMENT=30, 
		STRING=31, QUOTED_TEXT_=32;
	public static final int
		RULE_commands = 0, RULE_command = 1, RULE_execute = 2, RULE_mkdisk = 3, 
		RULE_rmdisk = 4, RULE_fdisk = 5, RULE_mount = 6, RULE_unmount = 7, RULE_mkfs = 8, 
		RULE_login = 9, RULE_logout = 10, RULE_mkgrp = 11, RULE_rmgrp = 12, RULE_mkusr = 13, 
		RULE_rmusr = 14, RULE_mkfile = 15, RULE_rep = 16, RULE_cat = 17, RULE_remove = 18, 
		RULE_edit = 19, RULE_rename = 20, RULE_mkdir = 21, RULE_pause = 22, RULE_chgrp = 23, 
		RULE_move = 24, RULE_loss = 25, RULE_recovery = 26, RULE_chown = 27, RULE_chmod = 28, 
		RULE_parametro = 29, RULE_value_parametro = 30;
	private static String[] makeRuleNames() {
		return new String[] {
			"commands", "command", "execute", "mkdisk", "rmdisk", "fdisk", "mount", 
			"unmount", "mkfs", "login", "logout", "mkgrp", "rmgrp", "mkusr", "rmusr", 
			"mkfile", "rep", "cat", "remove", "edit", "rename", "mkdir", "pause", 
			"chgrp", "move", "loss", "recovery", "chown", "chmod", "parametro", "value_parametro"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'execute'", "'mkdisk'", "'rmdisk'", "'fdisk'", "'mount'", "'unmount'", 
			"'mkfs'", "'login'", "'logout'", "'mkgrp'", "'rmgrp'", "'mkusr'", "'rmusr'", 
			"'mkfile'", "'rep'", "'cat'", "'remove'", "'edit'", "'rename'", "'mkdir'", 
			"'pause'", "'chgrp'", "'move'", "'loss'", "'recovery'", "'chown'", "'chmod'", 
			"'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, "ENBLANCO", "LINE_COMMENT", "STRING", "QUOTED_TEXT_"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "FileSysCLI.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public FileSysCLIParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CommandsContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(FileSysCLIParser.EOF, 0); }
		public List<CommandContext> command() {
			return getRuleContexts(CommandContext.class);
		}
		public CommandContext command(int i) {
			return getRuleContext(CommandContext.class,i);
		}
		public CommandsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_commands; }
	}

	public final CommandsContext commands() throws RecognitionException {
		CommandsContext _localctx = new CommandsContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_commands);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(65);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 268435454L) != 0)) {
				{
				{
				setState(62);
				command();
				}
				}
				setState(67);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(68);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CommandContext extends ParserRuleContext {
		public MkdiskContext mkdisk() {
			return getRuleContext(MkdiskContext.class,0);
		}
		public RmdiskContext rmdisk() {
			return getRuleContext(RmdiskContext.class,0);
		}
		public FdiskContext fdisk() {
			return getRuleContext(FdiskContext.class,0);
		}
		public ExecuteContext execute() {
			return getRuleContext(ExecuteContext.class,0);
		}
		public MountContext mount() {
			return getRuleContext(MountContext.class,0);
		}
		public UnmountContext unmount() {
			return getRuleContext(UnmountContext.class,0);
		}
		public MkfsContext mkfs() {
			return getRuleContext(MkfsContext.class,0);
		}
		public LoginContext login() {
			return getRuleContext(LoginContext.class,0);
		}
		public LogoutContext logout() {
			return getRuleContext(LogoutContext.class,0);
		}
		public MkgrpContext mkgrp() {
			return getRuleContext(MkgrpContext.class,0);
		}
		public RmgrpContext rmgrp() {
			return getRuleContext(RmgrpContext.class,0);
		}
		public MkusrContext mkusr() {
			return getRuleContext(MkusrContext.class,0);
		}
		public RmusrContext rmusr() {
			return getRuleContext(RmusrContext.class,0);
		}
		public MkfileContext mkfile() {
			return getRuleContext(MkfileContext.class,0);
		}
		public CatContext cat() {
			return getRuleContext(CatContext.class,0);
		}
		public RemoveContext remove() {
			return getRuleContext(RemoveContext.class,0);
		}
		public EditContext edit() {
			return getRuleContext(EditContext.class,0);
		}
		public RenameContext rename() {
			return getRuleContext(RenameContext.class,0);
		}
		public MkdirContext mkdir() {
			return getRuleContext(MkdirContext.class,0);
		}
		public ChgrpContext chgrp() {
			return getRuleContext(ChgrpContext.class,0);
		}
		public RepContext rep() {
			return getRuleContext(RepContext.class,0);
		}
		public MoveContext move() {
			return getRuleContext(MoveContext.class,0);
		}
		public LossContext loss() {
			return getRuleContext(LossContext.class,0);
		}
		public RecoveryContext recovery() {
			return getRuleContext(RecoveryContext.class,0);
		}
		public ChownContext chown() {
			return getRuleContext(ChownContext.class,0);
		}
		public ChmodContext chmod() {
			return getRuleContext(ChmodContext.class,0);
		}
		public PauseContext pause() {
			return getRuleContext(PauseContext.class,0);
		}
		public CommandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_command; }
	}

	public final CommandContext command() throws RecognitionException {
		CommandContext _localctx = new CommandContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_command);
		try {
			setState(97);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
				enterOuterAlt(_localctx, 1);
				{
				setState(70);
				mkdisk();
				}
				break;
			case T__2:
				enterOuterAlt(_localctx, 2);
				{
				setState(71);
				rmdisk();
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 3);
				{
				setState(72);
				fdisk();
				}
				break;
			case T__0:
				enterOuterAlt(_localctx, 4);
				{
				setState(73);
				execute();
				}
				break;
			case T__4:
				enterOuterAlt(_localctx, 5);
				{
				setState(74);
				mount();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 6);
				{
				setState(75);
				unmount();
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 7);
				{
				setState(76);
				mkfs();
				}
				break;
			case T__7:
				enterOuterAlt(_localctx, 8);
				{
				setState(77);
				login();
				}
				break;
			case T__8:
				enterOuterAlt(_localctx, 9);
				{
				setState(78);
				logout();
				}
				break;
			case T__9:
				enterOuterAlt(_localctx, 10);
				{
				setState(79);
				mkgrp();
				}
				break;
			case T__10:
				enterOuterAlt(_localctx, 11);
				{
				setState(80);
				rmgrp();
				}
				break;
			case T__11:
				enterOuterAlt(_localctx, 12);
				{
				setState(81);
				mkusr();
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 13);
				{
				setState(82);
				rmusr();
				}
				break;
			case T__13:
				enterOuterAlt(_localctx, 14);
				{
				setState(83);
				mkfile();
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 15);
				{
				setState(84);
				cat();
				}
				break;
			case T__16:
				enterOuterAlt(_localctx, 16);
				{
				setState(85);
				remove();
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 17);
				{
				setState(86);
				edit();
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 18);
				{
				setState(87);
				rename();
				}
				break;
			case T__19:
				enterOuterAlt(_localctx, 19);
				{
				setState(88);
				mkdir();
				}
				break;
			case T__21:
				enterOuterAlt(_localctx, 20);
				{
				setState(89);
				chgrp();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 21);
				{
				setState(90);
				rep();
				}
				break;
			case T__22:
				enterOuterAlt(_localctx, 22);
				{
				setState(91);
				move();
				}
				break;
			case T__23:
				enterOuterAlt(_localctx, 23);
				{
				setState(92);
				loss();
				}
				break;
			case T__24:
				enterOuterAlt(_localctx, 24);
				{
				setState(93);
				recovery();
				}
				break;
			case T__25:
				enterOuterAlt(_localctx, 25);
				{
				setState(94);
				chown();
				}
				break;
			case T__26:
				enterOuterAlt(_localctx, 26);
				{
				setState(95);
				chmod();
				}
				break;
			case T__20:
				enterOuterAlt(_localctx, 27);
				{
				setState(96);
				pause();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExecuteContext extends ParserRuleContext {
		public ExecuteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_execute; }
	 
		public ExecuteContext() { }
		public void copyFrom(ExecuteContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_executeContext extends ExecuteContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_executeContext(ExecuteContext ctx) { copyFrom(ctx); }
	}

	public final ExecuteContext execute() throws RecognitionException {
		ExecuteContext _localctx = new ExecuteContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_execute);
		int _la;
		try {
			_localctx = new Command_executeContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(99);
			match(T__0);
			setState(103);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(100);
				parametro();
				}
				}
				setState(105);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MkdiskContext extends ParserRuleContext {
		public MkdiskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkdisk; }
	 
		public MkdiskContext() { }
		public void copyFrom(MkdiskContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_mkdiskContext extends MkdiskContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_mkdiskContext(MkdiskContext ctx) { copyFrom(ctx); }
	}

	public final MkdiskContext mkdisk() throws RecognitionException {
		MkdiskContext _localctx = new MkdiskContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_mkdisk);
		int _la;
		try {
			_localctx = new Command_mkdiskContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(106);
			match(T__1);
			setState(110);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(107);
				parametro();
				}
				}
				setState(112);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RmdiskContext extends ParserRuleContext {
		public RmdiskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rmdisk; }
	 
		public RmdiskContext() { }
		public void copyFrom(RmdiskContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_rmdiskContext extends RmdiskContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_rmdiskContext(RmdiskContext ctx) { copyFrom(ctx); }
	}

	public final RmdiskContext rmdisk() throws RecognitionException {
		RmdiskContext _localctx = new RmdiskContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_rmdisk);
		int _la;
		try {
			_localctx = new Command_rmdiskContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(113);
			match(T__2);
			setState(117);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(114);
				parametro();
				}
				}
				setState(119);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FdiskContext extends ParserRuleContext {
		public FdiskContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fdisk; }
	 
		public FdiskContext() { }
		public void copyFrom(FdiskContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_fdiskContext extends FdiskContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_fdiskContext(FdiskContext ctx) { copyFrom(ctx); }
	}

	public final FdiskContext fdisk() throws RecognitionException {
		FdiskContext _localctx = new FdiskContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_fdisk);
		int _la;
		try {
			_localctx = new Command_fdiskContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(120);
			match(T__3);
			setState(124);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(121);
				parametro();
				}
				}
				setState(126);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MountContext extends ParserRuleContext {
		public MountContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mount; }
	 
		public MountContext() { }
		public void copyFrom(MountContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_mountContext extends MountContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_mountContext(MountContext ctx) { copyFrom(ctx); }
	}

	public final MountContext mount() throws RecognitionException {
		MountContext _localctx = new MountContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_mount);
		int _la;
		try {
			_localctx = new Command_mountContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(127);
			match(T__4);
			setState(131);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(128);
				parametro();
				}
				}
				setState(133);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UnmountContext extends ParserRuleContext {
		public UnmountContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unmount; }
	 
		public UnmountContext() { }
		public void copyFrom(UnmountContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_unmountContext extends UnmountContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_unmountContext(UnmountContext ctx) { copyFrom(ctx); }
	}

	public final UnmountContext unmount() throws RecognitionException {
		UnmountContext _localctx = new UnmountContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_unmount);
		int _la;
		try {
			_localctx = new Command_unmountContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			match(T__5);
			setState(138);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(135);
				parametro();
				}
				}
				setState(140);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MkfsContext extends ParserRuleContext {
		public MkfsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkfs; }
	 
		public MkfsContext() { }
		public void copyFrom(MkfsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_mkfsContext extends MkfsContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_mkfsContext(MkfsContext ctx) { copyFrom(ctx); }
	}

	public final MkfsContext mkfs() throws RecognitionException {
		MkfsContext _localctx = new MkfsContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_mkfs);
		int _la;
		try {
			_localctx = new Command_mkfsContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			match(T__6);
			setState(145);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(142);
				parametro();
				}
				}
				setState(147);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LoginContext extends ParserRuleContext {
		public LoginContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_login; }
	 
		public LoginContext() { }
		public void copyFrom(LoginContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_loginContext extends LoginContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_loginContext(LoginContext ctx) { copyFrom(ctx); }
	}

	public final LoginContext login() throws RecognitionException {
		LoginContext _localctx = new LoginContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_login);
		int _la;
		try {
			_localctx = new Command_loginContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(148);
			match(T__7);
			setState(152);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(149);
				parametro();
				}
				}
				setState(154);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LogoutContext extends ParserRuleContext {
		public LogoutContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_logout; }
	 
		public LogoutContext() { }
		public void copyFrom(LogoutContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_logoutContext extends LogoutContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_logoutContext(LogoutContext ctx) { copyFrom(ctx); }
	}

	public final LogoutContext logout() throws RecognitionException {
		LogoutContext _localctx = new LogoutContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_logout);
		int _la;
		try {
			_localctx = new Command_logoutContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(155);
			match(T__8);
			setState(159);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(156);
				parametro();
				}
				}
				setState(161);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MkgrpContext extends ParserRuleContext {
		public MkgrpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkgrp; }
	 
		public MkgrpContext() { }
		public void copyFrom(MkgrpContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_mkgrpContext extends MkgrpContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_mkgrpContext(MkgrpContext ctx) { copyFrom(ctx); }
	}

	public final MkgrpContext mkgrp() throws RecognitionException {
		MkgrpContext _localctx = new MkgrpContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_mkgrp);
		int _la;
		try {
			_localctx = new Command_mkgrpContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			match(T__9);
			setState(166);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(163);
				parametro();
				}
				}
				setState(168);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RmgrpContext extends ParserRuleContext {
		public RmgrpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rmgrp; }
	 
		public RmgrpContext() { }
		public void copyFrom(RmgrpContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_rmgrpContext extends RmgrpContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_rmgrpContext(RmgrpContext ctx) { copyFrom(ctx); }
	}

	public final RmgrpContext rmgrp() throws RecognitionException {
		RmgrpContext _localctx = new RmgrpContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_rmgrp);
		int _la;
		try {
			_localctx = new Command_rmgrpContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
			match(T__10);
			setState(173);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(170);
				parametro();
				}
				}
				setState(175);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MkusrContext extends ParserRuleContext {
		public MkusrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkusr; }
	 
		public MkusrContext() { }
		public void copyFrom(MkusrContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_mkusrContext extends MkusrContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_mkusrContext(MkusrContext ctx) { copyFrom(ctx); }
	}

	public final MkusrContext mkusr() throws RecognitionException {
		MkusrContext _localctx = new MkusrContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_mkusr);
		int _la;
		try {
			_localctx = new Command_mkusrContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(176);
			match(T__11);
			setState(180);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(177);
				parametro();
				}
				}
				setState(182);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RmusrContext extends ParserRuleContext {
		public RmusrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rmusr; }
	 
		public RmusrContext() { }
		public void copyFrom(RmusrContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_rmusrContext extends RmusrContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_rmusrContext(RmusrContext ctx) { copyFrom(ctx); }
	}

	public final RmusrContext rmusr() throws RecognitionException {
		RmusrContext _localctx = new RmusrContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_rmusr);
		int _la;
		try {
			_localctx = new Command_rmusrContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			match(T__12);
			setState(187);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(184);
				parametro();
				}
				}
				setState(189);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MkfileContext extends ParserRuleContext {
		public MkfileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkfile; }
	 
		public MkfileContext() { }
		public void copyFrom(MkfileContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_mkfileContext extends MkfileContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_mkfileContext(MkfileContext ctx) { copyFrom(ctx); }
	}

	public final MkfileContext mkfile() throws RecognitionException {
		MkfileContext _localctx = new MkfileContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_mkfile);
		int _la;
		try {
			_localctx = new Command_mkfileContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(190);
			match(T__13);
			setState(194);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(191);
				parametro();
				}
				}
				setState(196);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RepContext extends ParserRuleContext {
		public RepContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rep; }
	 
		public RepContext() { }
		public void copyFrom(RepContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_repContext extends RepContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_repContext(RepContext ctx) { copyFrom(ctx); }
	}

	public final RepContext rep() throws RecognitionException {
		RepContext _localctx = new RepContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_rep);
		int _la;
		try {
			_localctx = new Command_repContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(197);
			match(T__14);
			setState(201);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(198);
				parametro();
				}
				}
				setState(203);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CatContext extends ParserRuleContext {
		public CatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cat; }
	 
		public CatContext() { }
		public void copyFrom(CatContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_catContext extends CatContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_catContext(CatContext ctx) { copyFrom(ctx); }
	}

	public final CatContext cat() throws RecognitionException {
		CatContext _localctx = new CatContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_cat);
		int _la;
		try {
			_localctx = new Command_catContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(204);
			match(T__15);
			setState(208);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(205);
				parametro();
				}
				}
				setState(210);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RemoveContext extends ParserRuleContext {
		public RemoveContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_remove; }
	 
		public RemoveContext() { }
		public void copyFrom(RemoveContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_rmContext extends RemoveContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_rmContext(RemoveContext ctx) { copyFrom(ctx); }
	}

	public final RemoveContext remove() throws RecognitionException {
		RemoveContext _localctx = new RemoveContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_remove);
		int _la;
		try {
			_localctx = new Command_rmContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(211);
			match(T__16);
			setState(215);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(212);
				parametro();
				}
				}
				setState(217);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EditContext extends ParserRuleContext {
		public EditContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_edit; }
	 
		public EditContext() { }
		public void copyFrom(EditContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_editContext extends EditContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_editContext(EditContext ctx) { copyFrom(ctx); }
	}

	public final EditContext edit() throws RecognitionException {
		EditContext _localctx = new EditContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_edit);
		int _la;
		try {
			_localctx = new Command_editContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(218);
			match(T__17);
			setState(222);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(219);
				parametro();
				}
				}
				setState(224);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RenameContext extends ParserRuleContext {
		public RenameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rename; }
	 
		public RenameContext() { }
		public void copyFrom(RenameContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_renameContext extends RenameContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_renameContext(RenameContext ctx) { copyFrom(ctx); }
	}

	public final RenameContext rename() throws RecognitionException {
		RenameContext _localctx = new RenameContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_rename);
		int _la;
		try {
			_localctx = new Command_renameContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(225);
			match(T__18);
			setState(229);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(226);
				parametro();
				}
				}
				setState(231);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MkdirContext extends ParserRuleContext {
		public MkdirContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkdir; }
	 
		public MkdirContext() { }
		public void copyFrom(MkdirContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_mkdirContext extends MkdirContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_mkdirContext(MkdirContext ctx) { copyFrom(ctx); }
	}

	public final MkdirContext mkdir() throws RecognitionException {
		MkdirContext _localctx = new MkdirContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_mkdir);
		int _la;
		try {
			_localctx = new Command_mkdirContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(232);
			match(T__19);
			setState(236);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(233);
				parametro();
				}
				}
				setState(238);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PauseContext extends ParserRuleContext {
		public PauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pause; }
	 
		public PauseContext() { }
		public void copyFrom(PauseContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_pauseContext extends PauseContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_pauseContext(PauseContext ctx) { copyFrom(ctx); }
	}

	public final PauseContext pause() throws RecognitionException {
		PauseContext _localctx = new PauseContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_pause);
		int _la;
		try {
			_localctx = new Command_pauseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			match(T__20);
			setState(243);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(240);
				parametro();
				}
				}
				setState(245);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ChgrpContext extends ParserRuleContext {
		public ChgrpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_chgrp; }
	 
		public ChgrpContext() { }
		public void copyFrom(ChgrpContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_chgrpContext extends ChgrpContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_chgrpContext(ChgrpContext ctx) { copyFrom(ctx); }
	}

	public final ChgrpContext chgrp() throws RecognitionException {
		ChgrpContext _localctx = new ChgrpContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_chgrp);
		int _la;
		try {
			_localctx = new Command_chgrpContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(246);
			match(T__21);
			setState(250);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(247);
				parametro();
				}
				}
				setState(252);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MoveContext extends ParserRuleContext {
		public MoveContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_move; }
	 
		public MoveContext() { }
		public void copyFrom(MoveContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_moveContext extends MoveContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_moveContext(MoveContext ctx) { copyFrom(ctx); }
	}

	public final MoveContext move() throws RecognitionException {
		MoveContext _localctx = new MoveContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_move);
		int _la;
		try {
			_localctx = new Command_moveContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(253);
			match(T__22);
			setState(257);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(254);
				parametro();
				}
				}
				setState(259);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LossContext extends ParserRuleContext {
		public LossContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loss; }
	 
		public LossContext() { }
		public void copyFrom(LossContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_lossContext extends LossContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_lossContext(LossContext ctx) { copyFrom(ctx); }
	}

	public final LossContext loss() throws RecognitionException {
		LossContext _localctx = new LossContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_loss);
		int _la;
		try {
			_localctx = new Command_lossContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(260);
			match(T__23);
			setState(264);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(261);
				parametro();
				}
				}
				setState(266);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RecoveryContext extends ParserRuleContext {
		public RecoveryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_recovery; }
	 
		public RecoveryContext() { }
		public void copyFrom(RecoveryContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_recoveryContext extends RecoveryContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_recoveryContext(RecoveryContext ctx) { copyFrom(ctx); }
	}

	public final RecoveryContext recovery() throws RecognitionException {
		RecoveryContext _localctx = new RecoveryContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_recovery);
		int _la;
		try {
			_localctx = new Command_recoveryContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(267);
			match(T__24);
			setState(271);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(268);
				parametro();
				}
				}
				setState(273);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ChownContext extends ParserRuleContext {
		public ChownContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_chown; }
	 
		public ChownContext() { }
		public void copyFrom(ChownContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_chownContext extends ChownContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_chownContext(ChownContext ctx) { copyFrom(ctx); }
	}

	public final ChownContext chown() throws RecognitionException {
		ChownContext _localctx = new ChownContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_chown);
		int _la;
		try {
			_localctx = new Command_chownContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(274);
			match(T__25);
			setState(278);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(275);
				parametro();
				}
				}
				setState(280);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ChmodContext extends ParserRuleContext {
		public ChmodContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_chmod; }
	 
		public ChmodContext() { }
		public void copyFrom(ChmodContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Command_chmodContext extends ChmodContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public Command_chmodContext(ChmodContext ctx) { copyFrom(ctx); }
	}

	public final ChmodContext chmod() throws RecognitionException {
		ChmodContext _localctx = new ChmodContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_chmod);
		int _la;
		try {
			_localctx = new Command_chmodContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(281);
			match(T__26);
			setState(285);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==QUOTED_TEXT_) {
				{
				{
				setState(282);
				parametro();
				}
				}
				setState(287);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParametroContext extends ParserRuleContext {
		public ParametroContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametro; }
	 
		public ParametroContext() { }
		public void copyFrom(ParametroContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParameterContext extends ParametroContext {
		public TerminalNode QUOTED_TEXT_() { return getToken(FileSysCLIParser.QUOTED_TEXT_, 0); }
		public Value_parametroContext value_parametro() {
			return getRuleContext(Value_parametroContext.class,0);
		}
		public ParameterContext(ParametroContext ctx) { copyFrom(ctx); }
	}

	public final ParametroContext parametro() throws RecognitionException {
		ParametroContext _localctx = new ParametroContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_parametro);
		int _la;
		try {
			_localctx = new ParameterContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(288);
			match(QUOTED_TEXT_);
			setState(291);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__27) {
				{
				setState(289);
				match(T__27);
				setState(290);
				value_parametro();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Value_parametroContext extends ParserRuleContext {
		public Value_parametroContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value_parametro; }
	 
		public Value_parametroContext() { }
		public void copyFrom(Value_parametroContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Value_parameterContext extends Value_parametroContext {
		public TerminalNode QUOTED_TEXT_() { return getToken(FileSysCLIParser.QUOTED_TEXT_, 0); }
		public TerminalNode STRING() { return getToken(FileSysCLIParser.STRING, 0); }
		public Value_parameterContext(Value_parametroContext ctx) { copyFrom(ctx); }
	}

	public final Value_parametroContext value_parametro() throws RecognitionException {
		Value_parametroContext _localctx = new Value_parametroContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_value_parametro);
		int _la;
		try {
			_localctx = new Value_parameterContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(293);
			_la = _input.LA(1);
			if ( !(_la==STRING || _la==QUOTED_TEXT_) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001 \u0128\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0001\u0000\u0005\u0000@\b\u0000\n\u0000\f\u0000C\t\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001b\b\u0001"+
		"\u0001\u0002\u0001\u0002\u0005\u0002f\b\u0002\n\u0002\f\u0002i\t\u0002"+
		"\u0001\u0003\u0001\u0003\u0005\u0003m\b\u0003\n\u0003\f\u0003p\t\u0003"+
		"\u0001\u0004\u0001\u0004\u0005\u0004t\b\u0004\n\u0004\f\u0004w\t\u0004"+
		"\u0001\u0005\u0001\u0005\u0005\u0005{\b\u0005\n\u0005\f\u0005~\t\u0005"+
		"\u0001\u0006\u0001\u0006\u0005\u0006\u0082\b\u0006\n\u0006\f\u0006\u0085"+
		"\t\u0006\u0001\u0007\u0001\u0007\u0005\u0007\u0089\b\u0007\n\u0007\f\u0007"+
		"\u008c\t\u0007\u0001\b\u0001\b\u0005\b\u0090\b\b\n\b\f\b\u0093\t\b\u0001"+
		"\t\u0001\t\u0005\t\u0097\b\t\n\t\f\t\u009a\t\t\u0001\n\u0001\n\u0005\n"+
		"\u009e\b\n\n\n\f\n\u00a1\t\n\u0001\u000b\u0001\u000b\u0005\u000b\u00a5"+
		"\b\u000b\n\u000b\f\u000b\u00a8\t\u000b\u0001\f\u0001\f\u0005\f\u00ac\b"+
		"\f\n\f\f\f\u00af\t\f\u0001\r\u0001\r\u0005\r\u00b3\b\r\n\r\f\r\u00b6\t"+
		"\r\u0001\u000e\u0001\u000e\u0005\u000e\u00ba\b\u000e\n\u000e\f\u000e\u00bd"+
		"\t\u000e\u0001\u000f\u0001\u000f\u0005\u000f\u00c1\b\u000f\n\u000f\f\u000f"+
		"\u00c4\t\u000f\u0001\u0010\u0001\u0010\u0005\u0010\u00c8\b\u0010\n\u0010"+
		"\f\u0010\u00cb\t\u0010\u0001\u0011\u0001\u0011\u0005\u0011\u00cf\b\u0011"+
		"\n\u0011\f\u0011\u00d2\t\u0011\u0001\u0012\u0001\u0012\u0005\u0012\u00d6"+
		"\b\u0012\n\u0012\f\u0012\u00d9\t\u0012\u0001\u0013\u0001\u0013\u0005\u0013"+
		"\u00dd\b\u0013\n\u0013\f\u0013\u00e0\t\u0013\u0001\u0014\u0001\u0014\u0005"+
		"\u0014\u00e4\b\u0014\n\u0014\f\u0014\u00e7\t\u0014\u0001\u0015\u0001\u0015"+
		"\u0005\u0015\u00eb\b\u0015\n\u0015\f\u0015\u00ee\t\u0015\u0001\u0016\u0001"+
		"\u0016\u0005\u0016\u00f2\b\u0016\n\u0016\f\u0016\u00f5\t\u0016\u0001\u0017"+
		"\u0001\u0017\u0005\u0017\u00f9\b\u0017\n\u0017\f\u0017\u00fc\t\u0017\u0001"+
		"\u0018\u0001\u0018\u0005\u0018\u0100\b\u0018\n\u0018\f\u0018\u0103\t\u0018"+
		"\u0001\u0019\u0001\u0019\u0005\u0019\u0107\b\u0019\n\u0019\f\u0019\u010a"+
		"\t\u0019\u0001\u001a\u0001\u001a\u0005\u001a\u010e\b\u001a\n\u001a\f\u001a"+
		"\u0111\t\u001a\u0001\u001b\u0001\u001b\u0005\u001b\u0115\b\u001b\n\u001b"+
		"\f\u001b\u0118\t\u001b\u0001\u001c\u0001\u001c\u0005\u001c\u011c\b\u001c"+
		"\n\u001c\f\u001c\u011f\t\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0003"+
		"\u001d\u0124\b\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0000\u0000\u001f"+
		"\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a"+
		"\u001c\u001e \"$&(*,.02468:<\u0000\u0001\u0001\u0000\u001f \u013f\u0000"+
		"A\u0001\u0000\u0000\u0000\u0002a\u0001\u0000\u0000\u0000\u0004c\u0001"+
		"\u0000\u0000\u0000\u0006j\u0001\u0000\u0000\u0000\bq\u0001\u0000\u0000"+
		"\u0000\nx\u0001\u0000\u0000\u0000\f\u007f\u0001\u0000\u0000\u0000\u000e"+
		"\u0086\u0001\u0000\u0000\u0000\u0010\u008d\u0001\u0000\u0000\u0000\u0012"+
		"\u0094\u0001\u0000\u0000\u0000\u0014\u009b\u0001\u0000\u0000\u0000\u0016"+
		"\u00a2\u0001\u0000\u0000\u0000\u0018\u00a9\u0001\u0000\u0000\u0000\u001a"+
		"\u00b0\u0001\u0000\u0000\u0000\u001c\u00b7\u0001\u0000\u0000\u0000\u001e"+
		"\u00be\u0001\u0000\u0000\u0000 \u00c5\u0001\u0000\u0000\u0000\"\u00cc"+
		"\u0001\u0000\u0000\u0000$\u00d3\u0001\u0000\u0000\u0000&\u00da\u0001\u0000"+
		"\u0000\u0000(\u00e1\u0001\u0000\u0000\u0000*\u00e8\u0001\u0000\u0000\u0000"+
		",\u00ef\u0001\u0000\u0000\u0000.\u00f6\u0001\u0000\u0000\u00000\u00fd"+
		"\u0001\u0000\u0000\u00002\u0104\u0001\u0000\u0000\u00004\u010b\u0001\u0000"+
		"\u0000\u00006\u0112\u0001\u0000\u0000\u00008\u0119\u0001\u0000\u0000\u0000"+
		":\u0120\u0001\u0000\u0000\u0000<\u0125\u0001\u0000\u0000\u0000>@\u0003"+
		"\u0002\u0001\u0000?>\u0001\u0000\u0000\u0000@C\u0001\u0000\u0000\u0000"+
		"A?\u0001\u0000\u0000\u0000AB\u0001\u0000\u0000\u0000BD\u0001\u0000\u0000"+
		"\u0000CA\u0001\u0000\u0000\u0000DE\u0005\u0000\u0000\u0001E\u0001\u0001"+
		"\u0000\u0000\u0000Fb\u0003\u0006\u0003\u0000Gb\u0003\b\u0004\u0000Hb\u0003"+
		"\n\u0005\u0000Ib\u0003\u0004\u0002\u0000Jb\u0003\f\u0006\u0000Kb\u0003"+
		"\u000e\u0007\u0000Lb\u0003\u0010\b\u0000Mb\u0003\u0012\t\u0000Nb\u0003"+
		"\u0014\n\u0000Ob\u0003\u0016\u000b\u0000Pb\u0003\u0018\f\u0000Qb\u0003"+
		"\u001a\r\u0000Rb\u0003\u001c\u000e\u0000Sb\u0003\u001e\u000f\u0000Tb\u0003"+
		"\"\u0011\u0000Ub\u0003$\u0012\u0000Vb\u0003&\u0013\u0000Wb\u0003(\u0014"+
		"\u0000Xb\u0003*\u0015\u0000Yb\u0003.\u0017\u0000Zb\u0003 \u0010\u0000"+
		"[b\u00030\u0018\u0000\\b\u00032\u0019\u0000]b\u00034\u001a\u0000^b\u0003"+
		"6\u001b\u0000_b\u00038\u001c\u0000`b\u0003,\u0016\u0000aF\u0001\u0000"+
		"\u0000\u0000aG\u0001\u0000\u0000\u0000aH\u0001\u0000\u0000\u0000aI\u0001"+
		"\u0000\u0000\u0000aJ\u0001\u0000\u0000\u0000aK\u0001\u0000\u0000\u0000"+
		"aL\u0001\u0000\u0000\u0000aM\u0001\u0000\u0000\u0000aN\u0001\u0000\u0000"+
		"\u0000aO\u0001\u0000\u0000\u0000aP\u0001\u0000\u0000\u0000aQ\u0001\u0000"+
		"\u0000\u0000aR\u0001\u0000\u0000\u0000aS\u0001\u0000\u0000\u0000aT\u0001"+
		"\u0000\u0000\u0000aU\u0001\u0000\u0000\u0000aV\u0001\u0000\u0000\u0000"+
		"aW\u0001\u0000\u0000\u0000aX\u0001\u0000\u0000\u0000aY\u0001\u0000\u0000"+
		"\u0000aZ\u0001\u0000\u0000\u0000a[\u0001\u0000\u0000\u0000a\\\u0001\u0000"+
		"\u0000\u0000a]\u0001\u0000\u0000\u0000a^\u0001\u0000\u0000\u0000a_\u0001"+
		"\u0000\u0000\u0000a`\u0001\u0000\u0000\u0000b\u0003\u0001\u0000\u0000"+
		"\u0000cg\u0005\u0001\u0000\u0000df\u0003:\u001d\u0000ed\u0001\u0000\u0000"+
		"\u0000fi\u0001\u0000\u0000\u0000ge\u0001\u0000\u0000\u0000gh\u0001\u0000"+
		"\u0000\u0000h\u0005\u0001\u0000\u0000\u0000ig\u0001\u0000\u0000\u0000"+
		"jn\u0005\u0002\u0000\u0000km\u0003:\u001d\u0000lk\u0001\u0000\u0000\u0000"+
		"mp\u0001\u0000\u0000\u0000nl\u0001\u0000\u0000\u0000no\u0001\u0000\u0000"+
		"\u0000o\u0007\u0001\u0000\u0000\u0000pn\u0001\u0000\u0000\u0000qu\u0005"+
		"\u0003\u0000\u0000rt\u0003:\u001d\u0000sr\u0001\u0000\u0000\u0000tw\u0001"+
		"\u0000\u0000\u0000us\u0001\u0000\u0000\u0000uv\u0001\u0000\u0000\u0000"+
		"v\t\u0001\u0000\u0000\u0000wu\u0001\u0000\u0000\u0000x|\u0005\u0004\u0000"+
		"\u0000y{\u0003:\u001d\u0000zy\u0001\u0000\u0000\u0000{~\u0001\u0000\u0000"+
		"\u0000|z\u0001\u0000\u0000\u0000|}\u0001\u0000\u0000\u0000}\u000b\u0001"+
		"\u0000\u0000\u0000~|\u0001\u0000\u0000\u0000\u007f\u0083\u0005\u0005\u0000"+
		"\u0000\u0080\u0082\u0003:\u001d\u0000\u0081\u0080\u0001\u0000\u0000\u0000"+
		"\u0082\u0085\u0001\u0000\u0000\u0000\u0083\u0081\u0001\u0000\u0000\u0000"+
		"\u0083\u0084\u0001\u0000\u0000\u0000\u0084\r\u0001\u0000\u0000\u0000\u0085"+
		"\u0083\u0001\u0000\u0000\u0000\u0086\u008a\u0005\u0006\u0000\u0000\u0087"+
		"\u0089\u0003:\u001d\u0000\u0088\u0087\u0001\u0000\u0000\u0000\u0089\u008c"+
		"\u0001\u0000\u0000\u0000\u008a\u0088\u0001\u0000\u0000\u0000\u008a\u008b"+
		"\u0001\u0000\u0000\u0000\u008b\u000f\u0001\u0000\u0000\u0000\u008c\u008a"+
		"\u0001\u0000\u0000\u0000\u008d\u0091\u0005\u0007\u0000\u0000\u008e\u0090"+
		"\u0003:\u001d\u0000\u008f\u008e\u0001\u0000\u0000\u0000\u0090\u0093\u0001"+
		"\u0000\u0000\u0000\u0091\u008f\u0001\u0000\u0000\u0000\u0091\u0092\u0001"+
		"\u0000\u0000\u0000\u0092\u0011\u0001\u0000\u0000\u0000\u0093\u0091\u0001"+
		"\u0000\u0000\u0000\u0094\u0098\u0005\b\u0000\u0000\u0095\u0097\u0003:"+
		"\u001d\u0000\u0096\u0095\u0001\u0000\u0000\u0000\u0097\u009a\u0001\u0000"+
		"\u0000\u0000\u0098\u0096\u0001\u0000\u0000\u0000\u0098\u0099\u0001\u0000"+
		"\u0000\u0000\u0099\u0013\u0001\u0000\u0000\u0000\u009a\u0098\u0001\u0000"+
		"\u0000\u0000\u009b\u009f\u0005\t\u0000\u0000\u009c\u009e\u0003:\u001d"+
		"\u0000\u009d\u009c\u0001\u0000\u0000\u0000\u009e\u00a1\u0001\u0000\u0000"+
		"\u0000\u009f\u009d\u0001\u0000\u0000\u0000\u009f\u00a0\u0001\u0000\u0000"+
		"\u0000\u00a0\u0015\u0001\u0000\u0000\u0000\u00a1\u009f\u0001\u0000\u0000"+
		"\u0000\u00a2\u00a6\u0005\n\u0000\u0000\u00a3\u00a5\u0003:\u001d\u0000"+
		"\u00a4\u00a3\u0001\u0000\u0000\u0000\u00a5\u00a8\u0001\u0000\u0000\u0000"+
		"\u00a6\u00a4\u0001\u0000\u0000\u0000\u00a6\u00a7\u0001\u0000\u0000\u0000"+
		"\u00a7\u0017\u0001\u0000\u0000\u0000\u00a8\u00a6\u0001\u0000\u0000\u0000"+
		"\u00a9\u00ad\u0005\u000b\u0000\u0000\u00aa\u00ac\u0003:\u001d\u0000\u00ab"+
		"\u00aa\u0001\u0000\u0000\u0000\u00ac\u00af\u0001\u0000\u0000\u0000\u00ad"+
		"\u00ab\u0001\u0000\u0000\u0000\u00ad\u00ae\u0001\u0000\u0000\u0000\u00ae"+
		"\u0019\u0001\u0000\u0000\u0000\u00af\u00ad\u0001\u0000\u0000\u0000\u00b0"+
		"\u00b4\u0005\f\u0000\u0000\u00b1\u00b3\u0003:\u001d\u0000\u00b2\u00b1"+
		"\u0001\u0000\u0000\u0000\u00b3\u00b6\u0001\u0000\u0000\u0000\u00b4\u00b2"+
		"\u0001\u0000\u0000\u0000\u00b4\u00b5\u0001\u0000\u0000\u0000\u00b5\u001b"+
		"\u0001\u0000\u0000\u0000\u00b6\u00b4\u0001\u0000\u0000\u0000\u00b7\u00bb"+
		"\u0005\r\u0000\u0000\u00b8\u00ba\u0003:\u001d\u0000\u00b9\u00b8\u0001"+
		"\u0000\u0000\u0000\u00ba\u00bd\u0001\u0000\u0000\u0000\u00bb\u00b9\u0001"+
		"\u0000\u0000\u0000\u00bb\u00bc\u0001\u0000\u0000\u0000\u00bc\u001d\u0001"+
		"\u0000\u0000\u0000\u00bd\u00bb\u0001\u0000\u0000\u0000\u00be\u00c2\u0005"+
		"\u000e\u0000\u0000\u00bf\u00c1\u0003:\u001d\u0000\u00c0\u00bf\u0001\u0000"+
		"\u0000\u0000\u00c1\u00c4\u0001\u0000\u0000\u0000\u00c2\u00c0\u0001\u0000"+
		"\u0000\u0000\u00c2\u00c3\u0001\u0000\u0000\u0000\u00c3\u001f\u0001\u0000"+
		"\u0000\u0000\u00c4\u00c2\u0001\u0000\u0000\u0000\u00c5\u00c9\u0005\u000f"+
		"\u0000\u0000\u00c6\u00c8\u0003:\u001d\u0000\u00c7\u00c6\u0001\u0000\u0000"+
		"\u0000\u00c8\u00cb\u0001\u0000\u0000\u0000\u00c9\u00c7\u0001\u0000\u0000"+
		"\u0000\u00c9\u00ca\u0001\u0000\u0000\u0000\u00ca!\u0001\u0000\u0000\u0000"+
		"\u00cb\u00c9\u0001\u0000\u0000\u0000\u00cc\u00d0\u0005\u0010\u0000\u0000"+
		"\u00cd\u00cf\u0003:\u001d\u0000\u00ce\u00cd\u0001\u0000\u0000\u0000\u00cf"+
		"\u00d2\u0001\u0000\u0000\u0000\u00d0\u00ce\u0001\u0000\u0000\u0000\u00d0"+
		"\u00d1\u0001\u0000\u0000\u0000\u00d1#\u0001\u0000\u0000\u0000\u00d2\u00d0"+
		"\u0001\u0000\u0000\u0000\u00d3\u00d7\u0005\u0011\u0000\u0000\u00d4\u00d6"+
		"\u0003:\u001d\u0000\u00d5\u00d4\u0001\u0000\u0000\u0000\u00d6\u00d9\u0001"+
		"\u0000\u0000\u0000\u00d7\u00d5\u0001\u0000\u0000\u0000\u00d7\u00d8\u0001"+
		"\u0000\u0000\u0000\u00d8%\u0001\u0000\u0000\u0000\u00d9\u00d7\u0001\u0000"+
		"\u0000\u0000\u00da\u00de\u0005\u0012\u0000\u0000\u00db\u00dd\u0003:\u001d"+
		"\u0000\u00dc\u00db\u0001\u0000\u0000\u0000\u00dd\u00e0\u0001\u0000\u0000"+
		"\u0000\u00de\u00dc\u0001\u0000\u0000\u0000\u00de\u00df\u0001\u0000\u0000"+
		"\u0000\u00df\'\u0001\u0000\u0000\u0000\u00e0\u00de\u0001\u0000\u0000\u0000"+
		"\u00e1\u00e5\u0005\u0013\u0000\u0000\u00e2\u00e4\u0003:\u001d\u0000\u00e3"+
		"\u00e2\u0001\u0000\u0000\u0000\u00e4\u00e7\u0001\u0000\u0000\u0000\u00e5"+
		"\u00e3\u0001\u0000\u0000\u0000\u00e5\u00e6\u0001\u0000\u0000\u0000\u00e6"+
		")\u0001\u0000\u0000\u0000\u00e7\u00e5\u0001\u0000\u0000\u0000\u00e8\u00ec"+
		"\u0005\u0014\u0000\u0000\u00e9\u00eb\u0003:\u001d\u0000\u00ea\u00e9\u0001"+
		"\u0000\u0000\u0000\u00eb\u00ee\u0001\u0000\u0000\u0000\u00ec\u00ea\u0001"+
		"\u0000\u0000\u0000\u00ec\u00ed\u0001\u0000\u0000\u0000\u00ed+\u0001\u0000"+
		"\u0000\u0000\u00ee\u00ec\u0001\u0000\u0000\u0000\u00ef\u00f3\u0005\u0015"+
		"\u0000\u0000\u00f0\u00f2\u0003:\u001d\u0000\u00f1\u00f0\u0001\u0000\u0000"+
		"\u0000\u00f2\u00f5\u0001\u0000\u0000\u0000\u00f3\u00f1\u0001\u0000\u0000"+
		"\u0000\u00f3\u00f4\u0001\u0000\u0000\u0000\u00f4-\u0001\u0000\u0000\u0000"+
		"\u00f5\u00f3\u0001\u0000\u0000\u0000\u00f6\u00fa\u0005\u0016\u0000\u0000"+
		"\u00f7\u00f9\u0003:\u001d\u0000\u00f8\u00f7\u0001\u0000\u0000\u0000\u00f9"+
		"\u00fc\u0001\u0000\u0000\u0000\u00fa\u00f8\u0001\u0000\u0000\u0000\u00fa"+
		"\u00fb\u0001\u0000\u0000\u0000\u00fb/\u0001\u0000\u0000\u0000\u00fc\u00fa"+
		"\u0001\u0000\u0000\u0000\u00fd\u0101\u0005\u0017\u0000\u0000\u00fe\u0100"+
		"\u0003:\u001d\u0000\u00ff\u00fe\u0001\u0000\u0000\u0000\u0100\u0103\u0001"+
		"\u0000\u0000\u0000\u0101\u00ff\u0001\u0000\u0000\u0000\u0101\u0102\u0001"+
		"\u0000\u0000\u0000\u01021\u0001\u0000\u0000\u0000\u0103\u0101\u0001\u0000"+
		"\u0000\u0000\u0104\u0108\u0005\u0018\u0000\u0000\u0105\u0107\u0003:\u001d"+
		"\u0000\u0106\u0105\u0001\u0000\u0000\u0000\u0107\u010a\u0001\u0000\u0000"+
		"\u0000\u0108\u0106\u0001\u0000\u0000\u0000\u0108\u0109\u0001\u0000\u0000"+
		"\u0000\u01093\u0001\u0000\u0000\u0000\u010a\u0108\u0001\u0000\u0000\u0000"+
		"\u010b\u010f\u0005\u0019\u0000\u0000\u010c\u010e\u0003:\u001d\u0000\u010d"+
		"\u010c\u0001\u0000\u0000\u0000\u010e\u0111\u0001\u0000\u0000\u0000\u010f"+
		"\u010d\u0001\u0000\u0000\u0000\u010f\u0110\u0001\u0000\u0000\u0000\u0110"+
		"5\u0001\u0000\u0000\u0000\u0111\u010f\u0001\u0000\u0000\u0000\u0112\u0116"+
		"\u0005\u001a\u0000\u0000\u0113\u0115\u0003:\u001d\u0000\u0114\u0113\u0001"+
		"\u0000\u0000\u0000\u0115\u0118\u0001\u0000\u0000\u0000\u0116\u0114\u0001"+
		"\u0000\u0000\u0000\u0116\u0117\u0001\u0000\u0000\u0000\u01177\u0001\u0000"+
		"\u0000\u0000\u0118\u0116\u0001\u0000\u0000\u0000\u0119\u011d\u0005\u001b"+
		"\u0000\u0000\u011a\u011c\u0003:\u001d\u0000\u011b\u011a\u0001\u0000\u0000"+
		"\u0000\u011c\u011f\u0001\u0000\u0000\u0000\u011d\u011b\u0001\u0000\u0000"+
		"\u0000\u011d\u011e\u0001\u0000\u0000\u0000\u011e9\u0001\u0000\u0000\u0000"+
		"\u011f\u011d\u0001\u0000\u0000\u0000\u0120\u0123\u0005 \u0000\u0000\u0121"+
		"\u0122\u0005\u001c\u0000\u0000\u0122\u0124\u0003<\u001e\u0000\u0123\u0121"+
		"\u0001\u0000\u0000\u0000\u0123\u0124\u0001\u0000\u0000\u0000\u0124;\u0001"+
		"\u0000\u0000\u0000\u0125\u0126\u0007\u0000\u0000\u0000\u0126=\u0001\u0000"+
		"\u0000\u0000\u001eAagnu|\u0083\u008a\u0091\u0098\u009f\u00a6\u00ad\u00b4"+
		"\u00bb\u00c2\u00c9\u00d0\u00d7\u00de\u00e5\u00ec\u00f3\u00fa\u0101\u0108"+
		"\u010f\u0116\u011d\u0123";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}