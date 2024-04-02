// Generated from /home/ralda/Estudios/Lab ARCHIVOS/Proyectos/Proyecto 1/app/GenerateParser/FileSysCLILex.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class FileSysCLILex extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		ENBLANCO=1, LINE_COMMENT=2, SIZE=3, DECIMAL=4, ENTERO=5, STRING=6, QUOTED_TEXT_=7;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"ENBLANCO", "LINE_COMMENT", "SIZE", "DIG", "LETRA", "DECIMAL", "ENTERO", 
			"STRING", "QUOTED_TEXT_", "Quoted_text_item_", "QUOTED_TEXT", "Quoted_text_item", 
			"Escaped_character"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, "'-size='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "ENBLANCO", "LINE_COMMENT", "SIZE", "DECIMAL", "ENTERO", "STRING", 
			"QUOTED_TEXT_"
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


	public FileSysCLILex(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "FileSysCLILex.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\u0007c\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0001\u0000\u0004\u0000\u001d\b\u0000\u000b"+
		"\u0000\f\u0000\u001e\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0005\u0001\'\b\u0001\n\u0001\f\u0001*\t\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001"+
		"\u0004\u0001\u0005\u0004\u0005:\b\u0005\u000b\u0005\f\u0005;\u0001\u0005"+
		"\u0001\u0005\u0004\u0005@\b\u0005\u000b\u0005\f\u0005A\u0001\u0006\u0004"+
		"\u0006E\b\u0006\u000b\u0006\f\u0006F\u0001\u0007\u0001\u0007\u0003\u0007"+
		"K\b\u0007\u0001\u0007\u0001\u0007\u0001\b\u0004\bP\b\b\u000b\b\f\bQ\u0001"+
		"\t\u0001\t\u0003\tV\b\t\u0001\n\u0004\nY\b\n\u000b\n\f\nZ\u0001\u000b"+
		"\u0001\u000b\u0003\u000b_\b\u000b\u0001\f\u0001\f\u0001\f\u0000\u0000"+
		"\r\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0000\t\u0000\u000b\u0004"+
		"\r\u0005\u000f\u0006\u0011\u0007\u0013\u0000\u0015\u0000\u0017\u0000\u0019"+
		"\u0000\u0001\u0000\u0007\u0003\u0000\t\n\r\r  \u0002\u0000\n\n\r\r\u0001"+
		"\u000009\u0002\u0000AZaz\u0005\u0000\n\n\r\r  \"\"\\\\\u0004\u0000\n\n"+
		"\r\r\"\"\\\\\u0005\u0000\"\"\\\\nnrrttf\u0000\u0001\u0001\u0000\u0000"+
		"\u0000\u0000\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001\u0000\u0000"+
		"\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000"+
		"\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000\u0000"+
		"\u0001\u001c\u0001\u0000\u0000\u0000\u0003\"\u0001\u0000\u0000\u0000\u0005"+
		"-\u0001\u0000\u0000\u0000\u00074\u0001\u0000\u0000\u0000\t6\u0001\u0000"+
		"\u0000\u0000\u000b9\u0001\u0000\u0000\u0000\rD\u0001\u0000\u0000\u0000"+
		"\u000fH\u0001\u0000\u0000\u0000\u0011O\u0001\u0000\u0000\u0000\u0013U"+
		"\u0001\u0000\u0000\u0000\u0015X\u0001\u0000\u0000\u0000\u0017^\u0001\u0000"+
		"\u0000\u0000\u0019`\u0001\u0000\u0000\u0000\u001b\u001d\u0007\u0000\u0000"+
		"\u0000\u001c\u001b\u0001\u0000\u0000\u0000\u001d\u001e\u0001\u0000\u0000"+
		"\u0000\u001e\u001c\u0001\u0000\u0000\u0000\u001e\u001f\u0001\u0000\u0000"+
		"\u0000\u001f \u0001\u0000\u0000\u0000 !\u0006\u0000\u0000\u0000!\u0002"+
		"\u0001\u0000\u0000\u0000\"#\u0005/\u0000\u0000#$\u0005/\u0000\u0000$("+
		"\u0001\u0000\u0000\u0000%\'\b\u0001\u0000\u0000&%\u0001\u0000\u0000\u0000"+
		"\'*\u0001\u0000\u0000\u0000(&\u0001\u0000\u0000\u0000()\u0001\u0000\u0000"+
		"\u0000)+\u0001\u0000\u0000\u0000*(\u0001\u0000\u0000\u0000+,\u0006\u0001"+
		"\u0000\u0000,\u0004\u0001\u0000\u0000\u0000-.\u0005-\u0000\u0000./\u0005"+
		"s\u0000\u0000/0\u0005i\u0000\u000001\u0005z\u0000\u000012\u0005e\u0000"+
		"\u000023\u0005=\u0000\u00003\u0006\u0001\u0000\u0000\u000045\u0007\u0002"+
		"\u0000\u00005\b\u0001\u0000\u0000\u000067\u0007\u0003\u0000\u00007\n\u0001"+
		"\u0000\u0000\u00008:\u0003\u0007\u0003\u000098\u0001\u0000\u0000\u0000"+
		":;\u0001\u0000\u0000\u0000;9\u0001\u0000\u0000\u0000;<\u0001\u0000\u0000"+
		"\u0000<=\u0001\u0000\u0000\u0000=?\u0005.\u0000\u0000>@\u0003\u0007\u0003"+
		"\u0000?>\u0001\u0000\u0000\u0000@A\u0001\u0000\u0000\u0000A?\u0001\u0000"+
		"\u0000\u0000AB\u0001\u0000\u0000\u0000B\f\u0001\u0000\u0000\u0000CE\u0003"+
		"\u0007\u0003\u0000DC\u0001\u0000\u0000\u0000EF\u0001\u0000\u0000\u0000"+
		"FD\u0001\u0000\u0000\u0000FG\u0001\u0000\u0000\u0000G\u000e\u0001\u0000"+
		"\u0000\u0000HJ\u0005\"\u0000\u0000IK\u0003\u0015\n\u0000JI\u0001\u0000"+
		"\u0000\u0000JK\u0001\u0000\u0000\u0000KL\u0001\u0000\u0000\u0000LM\u0005"+
		"\"\u0000\u0000M\u0010\u0001\u0000\u0000\u0000NP\u0003\u0013\t\u0000ON"+
		"\u0001\u0000\u0000\u0000PQ\u0001\u0000\u0000\u0000QO\u0001\u0000\u0000"+
		"\u0000QR\u0001\u0000\u0000\u0000R\u0012\u0001\u0000\u0000\u0000SV\u0003"+
		"\u0019\f\u0000TV\b\u0004\u0000\u0000US\u0001\u0000\u0000\u0000UT\u0001"+
		"\u0000\u0000\u0000V\u0014\u0001\u0000\u0000\u0000WY\u0003\u0017\u000b"+
		"\u0000XW\u0001\u0000\u0000\u0000YZ\u0001\u0000\u0000\u0000ZX\u0001\u0000"+
		"\u0000\u0000Z[\u0001\u0000\u0000\u0000[\u0016\u0001\u0000\u0000\u0000"+
		"\\_\u0003\u0019\f\u0000]_\b\u0005\u0000\u0000^\\\u0001\u0000\u0000\u0000"+
		"^]\u0001\u0000\u0000\u0000_\u0018\u0001\u0000\u0000\u0000`a\u0005\\\u0000"+
		"\u0000ab\u0007\u0006\u0000\u0000b\u001a\u0001\u0000\u0000\u0000\u000b"+
		"\u0000\u001e(;AFJQUZ^\u0001\u0000\u0001\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}