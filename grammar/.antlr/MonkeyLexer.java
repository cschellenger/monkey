// Generated from /home/chad/dev/monkey/grammar/MonkeyLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class MonkeyLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		FN=1, IF=2, ELSE=3, WHILE=4, LET=5, RETURN=6, ASSIGN=7, PLUS=8, MINUS=9, 
		BANG=10, STAR=11, SLASH=12, PERCENT=13, END=14, LT=15, GT=16, GE=17, LE=18, 
		EQ=19, NOT_EQ=20, COMMA=21, COLON=22, BooleanLiteral=23, StringLiteral=24, 
		Identifier=25, FloatLiteral=26, IntegerLiteral=27, LPAREN=28, RPAREN=29, 
		LBRACE=30, RBRACE=31, LBRACKET=32, RBRACKET=33, WS=34;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"FN", "IF", "ELSE", "WHILE", "LET", "RETURN", "ASSIGN", "PLUS", "MINUS", 
			"BANG", "STAR", "SLASH", "PERCENT", "END", "LT", "GT", "GE", "LE", "EQ", 
			"NOT_EQ", "COMMA", "COLON", "DIGIT", "ALPHA", "TRUE", "FALSE", "BooleanLiteral", 
			"StringLiteral", "Identifier", "FloatLiteral", "IntegerLiteral", "LPAREN", 
			"RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'fn'", "'if'", "'else'", "'while'", "'let'", "'return'", "'='", 
			"'+'", "'-'", "'!'", "'*'", "'/'", "'%'", "';'", "'<'", "'>'", "'>='", 
			"'<='", "'=='", "'!='", "','", "':'", null, null, null, null, null, "'('", 
			"')'", "'{'", "'}'", "'['", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "FN", "IF", "ELSE", "WHILE", "LET", "RETURN", "ASSIGN", "PLUS", 
			"MINUS", "BANG", "STAR", "SLASH", "PERCENT", "END", "LT", "GT", "GE", 
			"LE", "EQ", "NOT_EQ", "COMMA", "COLON", "BooleanLiteral", "StringLiteral", 
			"Identifier", "FloatLiteral", "IntegerLiteral", "LPAREN", "RPAREN", "LBRACE", 
			"RBRACE", "LBRACKET", "RBRACKET", "WS"
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


	public MonkeyLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "MonkeyLexer.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2$\u00dc\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\3\2\3\2\3\2\3\3\3\3\3\3\3"+
		"\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16"+
		"\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\23\3\24"+
		"\3\24\3\24\3\25\3\25\3\25\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32"+
		"\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34\5\34\u00a1"+
		"\n\34\3\35\3\35\7\35\u00a5\n\35\f\35\16\35\u00a8\13\35\3\35\3\35\3\36"+
		"\3\36\3\36\7\36\u00af\n\36\f\36\16\36\u00b2\13\36\3\37\5\37\u00b5\n\37"+
		"\3\37\6\37\u00b8\n\37\r\37\16\37\u00b9\3\37\3\37\6\37\u00be\n\37\r\37"+
		"\16\37\u00bf\3 \5 \u00c3\n \3 \6 \u00c6\n \r \16 \u00c7\3!\3!\3\"\3\""+
		"\3#\3#\3$\3$\3%\3%\3&\3&\3\'\6\'\u00d7\n\'\r\'\16\'\u00d8\3\'\3\'\2\2"+
		"(\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20"+
		"\37\21!\22#\23%\24\'\25)\26+\27-\30/\2\61\2\63\2\65\2\67\319\32;\33=\34"+
		"?\35A\36C\37E G!I\"K#M$\3\2\6\3\2\62;\5\2C\\aac|\5\2\f\f\17\17$$\5\2\13"+
		"\f\17\17\"\"\2\u00e1\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2"+
		"\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3"+
		"\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2"+
		"\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2"+
		"\2-\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2"+
		"A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3"+
		"\2\2\2\3O\3\2\2\2\5R\3\2\2\2\7U\3\2\2\2\tZ\3\2\2\2\13`\3\2\2\2\rd\3\2"+
		"\2\2\17k\3\2\2\2\21m\3\2\2\2\23o\3\2\2\2\25q\3\2\2\2\27s\3\2\2\2\31u\3"+
		"\2\2\2\33w\3\2\2\2\35y\3\2\2\2\37{\3\2\2\2!}\3\2\2\2#\177\3\2\2\2%\u0082"+
		"\3\2\2\2\'\u0085\3\2\2\2)\u0088\3\2\2\2+\u008b\3\2\2\2-\u008d\3\2\2\2"+
		"/\u008f\3\2\2\2\61\u0091\3\2\2\2\63\u0093\3\2\2\2\65\u0098\3\2\2\2\67"+
		"\u00a0\3\2\2\29\u00a2\3\2\2\2;\u00ab\3\2\2\2=\u00b4\3\2\2\2?\u00c2\3\2"+
		"\2\2A\u00c9\3\2\2\2C\u00cb\3\2\2\2E\u00cd\3\2\2\2G\u00cf\3\2\2\2I\u00d1"+
		"\3\2\2\2K\u00d3\3\2\2\2M\u00d6\3\2\2\2OP\7h\2\2PQ\7p\2\2Q\4\3\2\2\2RS"+
		"\7k\2\2ST\7h\2\2T\6\3\2\2\2UV\7g\2\2VW\7n\2\2WX\7u\2\2XY\7g\2\2Y\b\3\2"+
		"\2\2Z[\7y\2\2[\\\7j\2\2\\]\7k\2\2]^\7n\2\2^_\7g\2\2_\n\3\2\2\2`a\7n\2"+
		"\2ab\7g\2\2bc\7v\2\2c\f\3\2\2\2de\7t\2\2ef\7g\2\2fg\7v\2\2gh\7w\2\2hi"+
		"\7t\2\2ij\7p\2\2j\16\3\2\2\2kl\7?\2\2l\20\3\2\2\2mn\7-\2\2n\22\3\2\2\2"+
		"op\7/\2\2p\24\3\2\2\2qr\7#\2\2r\26\3\2\2\2st\7,\2\2t\30\3\2\2\2uv\7\61"+
		"\2\2v\32\3\2\2\2wx\7\'\2\2x\34\3\2\2\2yz\7=\2\2z\36\3\2\2\2{|\7>\2\2|"+
		" \3\2\2\2}~\7@\2\2~\"\3\2\2\2\177\u0080\7@\2\2\u0080\u0081\7?\2\2\u0081"+
		"$\3\2\2\2\u0082\u0083\7>\2\2\u0083\u0084\7?\2\2\u0084&\3\2\2\2\u0085\u0086"+
		"\7?\2\2\u0086\u0087\7?\2\2\u0087(\3\2\2\2\u0088\u0089\7#\2\2\u0089\u008a"+
		"\7?\2\2\u008a*\3\2\2\2\u008b\u008c\7.\2\2\u008c,\3\2\2\2\u008d\u008e\7"+
		"<\2\2\u008e.\3\2\2\2\u008f\u0090\t\2\2\2\u0090\60\3\2\2\2\u0091\u0092"+
		"\t\3\2\2\u0092\62\3\2\2\2\u0093\u0094\7v\2\2\u0094\u0095\7t\2\2\u0095"+
		"\u0096\7w\2\2\u0096\u0097\7g\2\2\u0097\64\3\2\2\2\u0098\u0099\7h\2\2\u0099"+
		"\u009a\7c\2\2\u009a\u009b\7n\2\2\u009b\u009c\7u\2\2\u009c\u009d\7g\2\2"+
		"\u009d\66\3\2\2\2\u009e\u00a1\5\63\32\2\u009f\u00a1\5\65\33\2\u00a0\u009e"+
		"\3\2\2\2\u00a0\u009f\3\2\2\2\u00a18\3\2\2\2\u00a2\u00a6\7$\2\2\u00a3\u00a5"+
		"\n\4\2\2\u00a4\u00a3\3\2\2\2\u00a5\u00a8\3\2\2\2\u00a6\u00a4\3\2\2\2\u00a6"+
		"\u00a7\3\2\2\2\u00a7\u00a9\3\2\2\2\u00a8\u00a6\3\2\2\2\u00a9\u00aa\7$"+
		"\2\2\u00aa:\3\2\2\2\u00ab\u00b0\5\61\31\2\u00ac\u00af\5\61\31\2\u00ad"+
		"\u00af\5/\30\2\u00ae\u00ac\3\2\2\2\u00ae\u00ad\3\2\2\2\u00af\u00b2\3\2"+
		"\2\2\u00b0\u00ae\3\2\2\2\u00b0\u00b1\3\2\2\2\u00b1<\3\2\2\2\u00b2\u00b0"+
		"\3\2\2\2\u00b3\u00b5\5\23\n\2\u00b4\u00b3\3\2\2\2\u00b4\u00b5\3\2\2\2"+
		"\u00b5\u00b7\3\2\2\2\u00b6\u00b8\5/\30\2\u00b7\u00b6\3\2\2\2\u00b8\u00b9"+
		"\3\2\2\2\u00b9\u00b7\3\2\2\2\u00b9\u00ba\3\2\2\2\u00ba\u00bb\3\2\2\2\u00bb"+
		"\u00bd\7\60\2\2\u00bc\u00be\5/\30\2\u00bd\u00bc\3\2\2\2\u00be\u00bf\3"+
		"\2\2\2\u00bf\u00bd\3\2\2\2\u00bf\u00c0\3\2\2\2\u00c0>\3\2\2\2\u00c1\u00c3"+
		"\5\23\n\2\u00c2\u00c1\3\2\2\2\u00c2\u00c3\3\2\2\2\u00c3\u00c5\3\2\2\2"+
		"\u00c4\u00c6\5/\30\2\u00c5\u00c4\3\2\2\2\u00c6\u00c7\3\2\2\2\u00c7\u00c5"+
		"\3\2\2\2\u00c7\u00c8\3\2\2\2\u00c8@\3\2\2\2\u00c9\u00ca\7*\2\2\u00caB"+
		"\3\2\2\2\u00cb\u00cc\7+\2\2\u00ccD\3\2\2\2\u00cd\u00ce\7}\2\2\u00ceF\3"+
		"\2\2\2\u00cf\u00d0\7\177\2\2\u00d0H\3\2\2\2\u00d1\u00d2\7]\2\2\u00d2J"+
		"\3\2\2\2\u00d3\u00d4\7_\2\2\u00d4L\3\2\2\2\u00d5\u00d7\t\5\2\2\u00d6\u00d5"+
		"\3\2\2\2\u00d7\u00d8\3\2\2\2\u00d8\u00d6\3\2\2\2\u00d8\u00d9\3\2\2\2\u00d9"+
		"\u00da\3\2\2\2\u00da\u00db\b\'\2\2\u00dbN\3\2\2\2\r\2\u00a0\u00a6\u00ae"+
		"\u00b0\u00b4\u00b9\u00bf\u00c2\u00c7\u00d8\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}