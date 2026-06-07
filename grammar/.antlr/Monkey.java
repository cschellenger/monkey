// Generated from /home/chad/dev/antlr_monkey/grammar/Monkey.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Monkey extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		Prog=1, Statement=2, Literal=3, Expression=4, FunctionLiteral=5, CallExpression=6, 
		WhileStatement=7, IfExpression=8, ReturnStatement=9, LetStatement=10, 
		StringLiteral=11, Identifier=12, FloatLiteral=13, IntegerLiteral=14, WS=15, 
		BOOLEAN=16, TRUE=17, FALSE=18, FN=19, IF=20, ELSE=21, WHILE=22, LET=23, 
		RETURN=24, ASSIGN=25, PLUS=26, MINUS=27, BANG=28, STAR=29, SLASH=30, END=31, 
		LT=32, GT=33, GE=34, LE=35, EQ=36, NOT_EQ=37, COMMA=38, COLON=39, LPAREN=40, 
		RPAREN=41, LBRACE=42, RBRACE=43, LBRACKET=44, RBRACKET=45;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"Prog", "Statement", "Literal", "Expression", "FunctionLiteral", "CallExpression", 
			"WhileStatement", "IfExpression", "ReturnStatement", "LetStatement", 
			"StringLiteral", "Identifier", "FloatLiteral", "IntegerLiteral", "WS", 
			"DIGIT", "ALPHA", "BOOLEAN", "TRUE", "FALSE", "FN", "IF", "ELSE", "WHILE", 
			"LET", "RETURN", "ASSIGN", "PLUS", "MINUS", "BANG", "STAR", "SLASH", 
			"END", "LT", "GT", "GE", "LE", "EQ", "NOT_EQ", "COMMA", "COLON", "LPAREN", 
			"RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, "'true'", "'false'", "'fn'", "'if'", "'else'", 
			"'while'", "'let'", "'return'", "'='", "'+'", "'-'", "'!'", "'*'", "'/'", 
			"';'", "'<'", "'>'", "'>='", "'<='", "'=='", "'!='", "','", "':'", "'('", 
			"')'", "'{'", "'}'", "'['", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "Prog", "Statement", "Literal", "Expression", "FunctionLiteral", 
			"CallExpression", "WhileStatement", "IfExpression", "ReturnStatement", 
			"LetStatement", "StringLiteral", "Identifier", "FloatLiteral", "IntegerLiteral", 
			"WS", "BOOLEAN", "TRUE", "FALSE", "FN", "IF", "ELSE", "WHILE", "LET", 
			"RETURN", "ASSIGN", "PLUS", "MINUS", "BANG", "STAR", "SLASH", "END", 
			"LT", "GT", "GE", "LE", "EQ", "NOT_EQ", "COMMA", "COLON", "LPAREN", "RPAREN", 
			"LBRACE", "RBRACE", "LBRACKET", "RBRACKET"
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


	public Monkey(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Monkey.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2/\u0140\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\3\2\6\2c\n\2\r\2\16\2d\3\3\3\3\3\3\3"+
		"\3\5\3k\n\3\3\4\3\4\3\4\5\4p\n\4\3\5\3\5\3\5\3\5\3\5\5\5w\n\5\3\6\3\6"+
		"\3\6\3\6\3\6\5\6~\n\6\7\6\u0080\n\6\f\6\16\6\u0083\13\6\3\6\3\6\3\6\7"+
		"\6\u0088\n\6\f\6\16\6\u008b\13\6\3\6\3\6\3\7\3\7\3\7\3\7\5\7\u0093\n\7"+
		"\7\7\u0095\n\7\f\7\16\7\u0098\13\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00b0\n\t\3\n"+
		"\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\7\f\u00be\n\f\f\f\16"+
		"\f\u00c1\13\f\3\f\3\f\3\r\3\r\3\r\7\r\u00c8\n\r\f\r\16\r\u00cb\13\r\3"+
		"\16\6\16\u00ce\n\16\r\16\16\16\u00cf\3\16\3\16\6\16\u00d4\n\16\r\16\16"+
		"\16\u00d5\3\17\6\17\u00d9\n\17\r\17\16\17\u00da\3\20\6\20\u00de\n\20\r"+
		"\20\16\20\u00df\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\5\23\u00ea\n\23"+
		"\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26"+
		"\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\35"+
		"\3\35\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3%\3&"+
		"\3&\3&\3\'\3\'\3\'\3(\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3/\3/"+
		"\3\60\3\60\2\2\61\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31"+
		"\16\33\17\35\20\37\21!\2#\2%\22\'\23)\24+\25-\26/\27\61\30\63\31\65\32"+
		"\67\339\34;\35=\36?\37A C!E\"G#I$K%M&O\'Q(S)U*W+Y,[-]._/\3\2\6\5\2\f\f"+
		"\17\17$$\5\2\13\f\17\17\"\"\3\2\62;\5\2C\\aac|\2\u0155\2\3\3\2\2\2\2\5"+
		"\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2"+
		"\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33"+
		"\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2"+
		"+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2"+
		"\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2"+
		"C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3"+
		"\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2"+
		"\2\2]\3\2\2\2\2_\3\2\2\2\3b\3\2\2\2\5j\3\2\2\2\7o\3\2\2\2\tv\3\2\2\2\13"+
		"x\3\2\2\2\r\u008e\3\2\2\2\17\u009b\3\2\2\2\21\u00a4\3\2\2\2\23\u00b1\3"+
		"\2\2\2\25\u00b5\3\2\2\2\27\u00bb\3\2\2\2\31\u00c4\3\2\2\2\33\u00cd\3\2"+
		"\2\2\35\u00d8\3\2\2\2\37\u00dd\3\2\2\2!\u00e3\3\2\2\2#\u00e5\3\2\2\2%"+
		"\u00e9\3\2\2\2\'\u00eb\3\2\2\2)\u00f0\3\2\2\2+\u00f6\3\2\2\2-\u00f9\3"+
		"\2\2\2/\u00fc\3\2\2\2\61\u0101\3\2\2\2\63\u0107\3\2\2\2\65\u010b\3\2\2"+
		"\2\67\u0112\3\2\2\29\u0114\3\2\2\2;\u0116\3\2\2\2=\u0118\3\2\2\2?\u011a"+
		"\3\2\2\2A\u011c\3\2\2\2C\u011e\3\2\2\2E\u0120\3\2\2\2G\u0122\3\2\2\2I"+
		"\u0124\3\2\2\2K\u0127\3\2\2\2M\u012a\3\2\2\2O\u012d\3\2\2\2Q\u0130\3\2"+
		"\2\2S\u0132\3\2\2\2U\u0134\3\2\2\2W\u0136\3\2\2\2Y\u0138\3\2\2\2[\u013a"+
		"\3\2\2\2]\u013c\3\2\2\2_\u013e\3\2\2\2ac\5\5\3\2ba\3\2\2\2cd\3\2\2\2d"+
		"b\3\2\2\2de\3\2\2\2e\4\3\2\2\2fk\5\23\n\2gk\5\25\13\2hk\5\17\b\2ik\5\t"+
		"\5\2jf\3\2\2\2jg\3\2\2\2jh\3\2\2\2ji\3\2\2\2k\6\3\2\2\2lp\5\35\17\2mp"+
		"\5\33\16\2np\5\27\f\2ol\3\2\2\2om\3\2\2\2on\3\2\2\2p\b\3\2\2\2qw\5\7\4"+
		"\2rw\5\21\t\2sw\5\13\6\2tw\5\r\7\2uw\5\31\r\2vq\3\2\2\2vr\3\2\2\2vs\3"+
		"\2\2\2vt\3\2\2\2vu\3\2\2\2w\n\3\2\2\2xy\5+\26\2yz\5\31\r\2z\u0081\5U+"+
		"\2{}\5\31\r\2|~\5Q)\2}|\3\2\2\2}~\3\2\2\2~\u0080\3\2\2\2\177{\3\2\2\2"+
		"\u0080\u0083\3\2\2\2\u0081\177\3\2\2\2\u0081\u0082\3\2\2\2\u0082\u0084"+
		"\3\2\2\2\u0083\u0081\3\2\2\2\u0084\u0085\5W,\2\u0085\u0089\5Y-\2\u0086"+
		"\u0088\5\5\3\2\u0087\u0086\3\2\2\2\u0088\u008b\3\2\2\2\u0089\u0087\3\2"+
		"\2\2\u0089\u008a\3\2\2\2\u008a\u008c\3\2\2\2\u008b\u0089\3\2\2\2\u008c"+
		"\u008d\5[.\2\u008d\f\3\2\2\2\u008e\u008f\5\31\r\2\u008f\u0096\5U+\2\u0090"+
		"\u0092\5\31\r\2\u0091\u0093\5Q)\2\u0092\u0091\3\2\2\2\u0092\u0093\3\2"+
		"\2\2\u0093\u0095\3\2\2\2\u0094\u0090\3\2\2\2\u0095\u0098\3\2\2\2\u0096"+
		"\u0094\3\2\2\2\u0096\u0097\3\2\2\2\u0097\u0099\3\2\2\2\u0098\u0096\3\2"+
		"\2\2\u0099\u009a\5W,\2\u009a\16\3\2\2\2\u009b\u009c\5\61\31\2\u009c\u009d"+
		"\5U+\2\u009d\u009e\5\t\5\2\u009e\u009f\5W,\2\u009f\u00a0\5Y-\2\u00a0\u00a1"+
		"\5\5\3\2\u00a1\u00a2\5[.\2\u00a2\u00a3\5C\"\2\u00a3\20\3\2\2\2\u00a4\u00a5"+
		"\5-\27\2\u00a5\u00a6\5U+\2\u00a6\u00a7\5\t\5\2\u00a7\u00a8\5W,\2\u00a8"+
		"\u00a9\5Y-\2\u00a9\u00af\5\t\5\2\u00aa\u00ab\5/\30\2\u00ab\u00ac\5Y-\2"+
		"\u00ac\u00ad\5\t\5\2\u00ad\u00ae\5[.\2\u00ae\u00b0\3\2\2\2\u00af\u00aa"+
		"\3\2\2\2\u00af\u00b0\3\2\2\2\u00b0\22\3\2\2\2\u00b1\u00b2\5\65\33\2\u00b2"+
		"\u00b3\5\t\5\2\u00b3\u00b4\5C\"\2\u00b4\24\3\2\2\2\u00b5\u00b6\5\63\32"+
		"\2\u00b6\u00b7\5\31\r\2\u00b7\u00b8\5\67\34\2\u00b8\u00b9\5\t\5\2\u00b9"+
		"\u00ba\5C\"\2\u00ba\26\3\2\2\2\u00bb\u00bf\7$\2\2\u00bc\u00be\n\2\2\2"+
		"\u00bd\u00bc\3\2\2\2\u00be\u00c1\3\2\2\2\u00bf\u00bd\3\2\2\2\u00bf\u00c0"+
		"\3\2\2\2\u00c0\u00c2\3\2\2\2\u00c1\u00bf\3\2\2\2\u00c2\u00c3\7$\2\2\u00c3"+
		"\30\3\2\2\2\u00c4\u00c9\5#\22\2\u00c5\u00c8\5#\22\2\u00c6\u00c8\5!\21"+
		"\2\u00c7\u00c5\3\2\2\2\u00c7\u00c6\3\2\2\2\u00c8\u00cb\3\2\2\2\u00c9\u00c7"+
		"\3\2\2\2\u00c9\u00ca\3\2\2\2\u00ca\32\3\2\2\2\u00cb\u00c9\3\2\2\2\u00cc"+
		"\u00ce\5!\21\2\u00cd\u00cc\3\2\2\2\u00ce\u00cf\3\2\2\2\u00cf\u00cd\3\2"+
		"\2\2\u00cf\u00d0\3\2\2\2\u00d0\u00d1\3\2\2\2\u00d1\u00d3\7\60\2\2\u00d2"+
		"\u00d4\5!\21\2\u00d3\u00d2\3\2\2\2\u00d4\u00d5\3\2\2\2\u00d5\u00d3\3\2"+
		"\2\2\u00d5\u00d6\3\2\2\2\u00d6\34\3\2\2\2\u00d7\u00d9\5!\21\2\u00d8\u00d7"+
		"\3\2\2\2\u00d9\u00da\3\2\2\2\u00da\u00d8\3\2\2\2\u00da\u00db\3\2\2\2\u00db"+
		"\36\3\2\2\2\u00dc\u00de\t\3\2\2\u00dd\u00dc\3\2\2\2\u00de\u00df\3\2\2"+
		"\2\u00df\u00dd\3\2\2\2\u00df\u00e0\3\2\2\2\u00e0\u00e1\3\2\2\2\u00e1\u00e2"+
		"\b\20\2\2\u00e2 \3\2\2\2\u00e3\u00e4\t\4\2\2\u00e4\"\3\2\2\2\u00e5\u00e6"+
		"\t\5\2\2\u00e6$\3\2\2\2\u00e7\u00ea\5\'\24\2\u00e8\u00ea\5)\25\2\u00e9"+
		"\u00e7\3\2\2\2\u00e9\u00e8\3\2\2\2\u00ea&\3\2\2\2\u00eb\u00ec\7v\2\2\u00ec"+
		"\u00ed\7t\2\2\u00ed\u00ee\7w\2\2\u00ee\u00ef\7g\2\2\u00ef(\3\2\2\2\u00f0"+
		"\u00f1\7h\2\2\u00f1\u00f2\7c\2\2\u00f2\u00f3\7n\2\2\u00f3\u00f4\7u\2\2"+
		"\u00f4\u00f5\7g\2\2\u00f5*\3\2\2\2\u00f6\u00f7\7h\2\2\u00f7\u00f8\7p\2"+
		"\2\u00f8,\3\2\2\2\u00f9\u00fa\7k\2\2\u00fa\u00fb\7h\2\2\u00fb.\3\2\2\2"+
		"\u00fc\u00fd\7g\2\2\u00fd\u00fe\7n\2\2\u00fe\u00ff\7u\2\2\u00ff\u0100"+
		"\7g\2\2\u0100\60\3\2\2\2\u0101\u0102\7y\2\2\u0102\u0103\7j\2\2\u0103\u0104"+
		"\7k\2\2\u0104\u0105\7n\2\2\u0105\u0106\7g\2\2\u0106\62\3\2\2\2\u0107\u0108"+
		"\7n\2\2\u0108\u0109\7g\2\2\u0109\u010a\7v\2\2\u010a\64\3\2\2\2\u010b\u010c"+
		"\7t\2\2\u010c\u010d\7g\2\2\u010d\u010e\7v\2\2\u010e\u010f\7w\2\2\u010f"+
		"\u0110\7t\2\2\u0110\u0111\7p\2\2\u0111\66\3\2\2\2\u0112\u0113\7?\2\2\u0113"+
		"8\3\2\2\2\u0114\u0115\7-\2\2\u0115:\3\2\2\2\u0116\u0117\7/\2\2\u0117<"+
		"\3\2\2\2\u0118\u0119\7#\2\2\u0119>\3\2\2\2\u011a\u011b\7,\2\2\u011b@\3"+
		"\2\2\2\u011c\u011d\7\61\2\2\u011dB\3\2\2\2\u011e\u011f\7=\2\2\u011fD\3"+
		"\2\2\2\u0120\u0121\7>\2\2\u0121F\3\2\2\2\u0122\u0123\7@\2\2\u0123H\3\2"+
		"\2\2\u0124\u0125\7@\2\2\u0125\u0126\7?\2\2\u0126J\3\2\2\2\u0127\u0128"+
		"\7>\2\2\u0128\u0129\7?\2\2\u0129L\3\2\2\2\u012a\u012b\7?\2\2\u012b\u012c"+
		"\7?\2\2\u012cN\3\2\2\2\u012d\u012e\7#\2\2\u012e\u012f\7?\2\2\u012fP\3"+
		"\2\2\2\u0130\u0131\7.\2\2\u0131R\3\2\2\2\u0132\u0133\7<\2\2\u0133T\3\2"+
		"\2\2\u0134\u0135\7*\2\2\u0135V\3\2\2\2\u0136\u0137\7+\2\2\u0137X\3\2\2"+
		"\2\u0138\u0139\7}\2\2\u0139Z\3\2\2\2\u013a\u013b\7\177\2\2\u013b\\\3\2"+
		"\2\2\u013c\u013d\7]\2\2\u013d^\3\2\2\2\u013e\u013f\7_\2\2\u013f`\3\2\2"+
		"\2\25\2djov}\u0081\u0089\u0092\u0096\u00af\u00bf\u00c7\u00c9\u00cf\u00d5"+
		"\u00da\u00df\u00e9\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}