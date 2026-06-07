// Code generated from MonkeyParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MonkeyParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MonkeyParser struct {
	*antlr.BaseParser
}

var MonkeyParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func monkeyparserParserInit() {
	staticData := &MonkeyParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'fn'", "'if'", "'else'", "'while'", "'let'", "'return'", "'='",
		"'+'", "'-'", "'!'", "'*'", "'/'", "'%'", "';'", "'<'", "'>'", "'>='",
		"'<='", "'=='", "'!='", "','", "':'", "", "", "", "", "", "'('", "')'",
		"'{'", "'}'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "FN", "IF", "ELSE", "WHILE", "LET", "RETURN", "ASSIGN", "PLUS",
		"MINUS", "BANG", "STAR", "SLASH", "PERCENT", "END", "LT", "GT", "GE",
		"LE", "EQ", "NOT_EQ", "COMMA", "COLON", "BooleanLiteral", "StringLiteral",
		"Identifier", "FloatLiteral", "IntegerLiteral", "LPAREN", "RPAREN",
		"LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "statement", "expression", "identifier", "literal", "expression_list",
		"expression_pair", "array_literal", "hash_literal", "function_literal",
		"params", "if_expression", "call_expression", "let_statement",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 34, 184, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 4, 0, 30, 8, 0, 11,
		0, 12, 0, 31, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 53, 8, 1,
		5, 1, 55, 8, 1, 10, 1, 12, 1, 58, 9, 1, 1, 1, 1, 1, 1, 1, 3, 1, 63, 8,
		1, 3, 1, 65, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 81, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 97, 8, 2,
		10, 2, 12, 2, 100, 9, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5,
		109, 8, 5, 10, 5, 12, 5, 112, 9, 5, 1, 5, 1, 5, 3, 5, 116, 8, 5, 3, 5,
		118, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 5, 8, 132, 8, 8, 10, 8, 12, 8, 135, 9, 8, 1, 8, 1, 8, 3, 8,
		139, 8, 8, 3, 8, 141, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 10, 1, 10, 5, 10, 153, 8, 10, 10, 10, 12, 10, 156, 9, 10, 1, 10,
		1, 10, 3, 10, 160, 8, 10, 3, 10, 162, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 3, 11, 171, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 0, 1, 4, 14, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 0, 4, 1, 0, 11, 13, 1, 0,
		8, 9, 1, 0, 15, 20, 2, 0, 23, 24, 26, 27, 199, 0, 29, 1, 0, 0, 0, 2, 64,
		1, 0, 0, 0, 4, 80, 1, 0, 0, 0, 6, 101, 1, 0, 0, 0, 8, 103, 1, 0, 0, 0,
		10, 110, 1, 0, 0, 0, 12, 119, 1, 0, 0, 0, 14, 123, 1, 0, 0, 0, 16, 127,
		1, 0, 0, 0, 18, 144, 1, 0, 0, 0, 20, 154, 1, 0, 0, 0, 22, 163, 1, 0, 0,
		0, 24, 172, 1, 0, 0, 0, 26, 177, 1, 0, 0, 0, 28, 30, 3, 2, 1, 0, 29, 28,
		1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0,
		32, 1, 1, 0, 0, 0, 33, 34, 5, 6, 0, 0, 34, 35, 3, 4, 2, 0, 35, 36, 5, 14,
		0, 0, 36, 65, 1, 0, 0, 0, 37, 38, 5, 5, 0, 0, 38, 39, 5, 25, 0, 0, 39,
		40, 5, 7, 0, 0, 40, 41, 3, 4, 2, 0, 41, 42, 5, 14, 0, 0, 42, 65, 1, 0,
		0, 0, 43, 44, 5, 4, 0, 0, 44, 45, 5, 28, 0, 0, 45, 46, 3, 4, 2, 0, 46,
		47, 5, 29, 0, 0, 47, 48, 3, 2, 1, 0, 48, 65, 1, 0, 0, 0, 49, 56, 5, 30,
		0, 0, 50, 52, 3, 2, 1, 0, 51, 53, 5, 14, 0, 0, 52, 51, 1, 0, 0, 0, 52,
		53, 1, 0, 0, 0, 53, 55, 1, 0, 0, 0, 54, 50, 1, 0, 0, 0, 55, 58, 1, 0, 0,
		0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 59, 1, 0, 0, 0, 58, 56,
		1, 0, 0, 0, 59, 65, 5, 31, 0, 0, 60, 62, 3, 4, 2, 0, 61, 63, 5, 14, 0,
		0, 62, 61, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 65, 1, 0, 0, 0, 64, 33,
		1, 0, 0, 0, 64, 37, 1, 0, 0, 0, 64, 43, 1, 0, 0, 0, 64, 49, 1, 0, 0, 0,
		64, 60, 1, 0, 0, 0, 65, 3, 1, 0, 0, 0, 66, 67, 6, 2, -1, 0, 67, 81, 3,
		22, 11, 0, 68, 81, 3, 18, 9, 0, 69, 81, 3, 24, 12, 0, 70, 81, 3, 6, 3,
		0, 71, 81, 3, 8, 4, 0, 72, 81, 3, 14, 7, 0, 73, 81, 3, 16, 8, 0, 74, 75,
		5, 28, 0, 0, 75, 76, 3, 4, 2, 0, 76, 77, 5, 29, 0, 0, 77, 81, 1, 0, 0,
		0, 78, 79, 5, 10, 0, 0, 79, 81, 3, 4, 2, 1, 80, 66, 1, 0, 0, 0, 80, 68,
		1, 0, 0, 0, 80, 69, 1, 0, 0, 0, 80, 70, 1, 0, 0, 0, 80, 71, 1, 0, 0, 0,
		80, 72, 1, 0, 0, 0, 80, 73, 1, 0, 0, 0, 80, 74, 1, 0, 0, 0, 80, 78, 1,
		0, 0, 0, 81, 98, 1, 0, 0, 0, 82, 83, 10, 5, 0, 0, 83, 84, 7, 0, 0, 0, 84,
		97, 3, 4, 2, 6, 85, 86, 10, 4, 0, 0, 86, 87, 7, 1, 0, 0, 87, 97, 3, 4,
		2, 5, 88, 89, 10, 3, 0, 0, 89, 90, 7, 2, 0, 0, 90, 97, 3, 4, 2, 4, 91,
		92, 10, 10, 0, 0, 92, 93, 5, 32, 0, 0, 93, 94, 3, 4, 2, 0, 94, 95, 5, 33,
		0, 0, 95, 97, 1, 0, 0, 0, 96, 82, 1, 0, 0, 0, 96, 85, 1, 0, 0, 0, 96, 88,
		1, 0, 0, 0, 96, 91, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0,
		98, 99, 1, 0, 0, 0, 99, 5, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 102, 5,
		25, 0, 0, 102, 7, 1, 0, 0, 0, 103, 104, 7, 3, 0, 0, 104, 9, 1, 0, 0, 0,
		105, 106, 3, 4, 2, 0, 106, 107, 5, 21, 0, 0, 107, 109, 1, 0, 0, 0, 108,
		105, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 110, 111,
		1, 0, 0, 0, 111, 117, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113, 115, 3, 4,
		2, 0, 114, 116, 5, 21, 0, 0, 115, 114, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0,
		116, 118, 1, 0, 0, 0, 117, 113, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118,
		11, 1, 0, 0, 0, 119, 120, 3, 4, 2, 0, 120, 121, 5, 22, 0, 0, 121, 122,
		3, 4, 2, 0, 122, 13, 1, 0, 0, 0, 123, 124, 5, 32, 0, 0, 124, 125, 3, 10,
		5, 0, 125, 126, 5, 33, 0, 0, 126, 15, 1, 0, 0, 0, 127, 133, 5, 30, 0, 0,
		128, 129, 3, 12, 6, 0, 129, 130, 5, 21, 0, 0, 130, 132, 1, 0, 0, 0, 131,
		128, 1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134,
		1, 0, 0, 0, 134, 140, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 138, 3, 12,
		6, 0, 137, 139, 5, 21, 0, 0, 138, 137, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0,
		139, 141, 1, 0, 0, 0, 140, 136, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141,
		142, 1, 0, 0, 0, 142, 143, 5, 31, 0, 0, 143, 17, 1, 0, 0, 0, 144, 145,
		5, 1, 0, 0, 145, 146, 5, 28, 0, 0, 146, 147, 3, 20, 10, 0, 147, 148, 5,
		29, 0, 0, 148, 149, 3, 2, 1, 0, 149, 19, 1, 0, 0, 0, 150, 151, 5, 25, 0,
		0, 151, 153, 5, 21, 0, 0, 152, 150, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154,
		152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 161, 1, 0, 0, 0, 156, 154,
		1, 0, 0, 0, 157, 159, 5, 25, 0, 0, 158, 160, 5, 21, 0, 0, 159, 158, 1,
		0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 162, 1, 0, 0, 0, 161, 157, 1, 0, 0,
		0, 161, 162, 1, 0, 0, 0, 162, 21, 1, 0, 0, 0, 163, 164, 5, 2, 0, 0, 164,
		165, 5, 28, 0, 0, 165, 166, 3, 4, 2, 0, 166, 167, 5, 29, 0, 0, 167, 170,
		3, 2, 1, 0, 168, 169, 5, 3, 0, 0, 169, 171, 3, 2, 1, 0, 170, 168, 1, 0,
		0, 0, 170, 171, 1, 0, 0, 0, 171, 23, 1, 0, 0, 0, 172, 173, 3, 6, 3, 0,
		173, 174, 5, 28, 0, 0, 174, 175, 3, 10, 5, 0, 175, 176, 5, 29, 0, 0, 176,
		25, 1, 0, 0, 0, 177, 178, 5, 5, 0, 0, 178, 179, 3, 6, 3, 0, 179, 180, 5,
		7, 0, 0, 180, 181, 3, 4, 2, 0, 181, 182, 5, 14, 0, 0, 182, 27, 1, 0, 0,
		0, 18, 31, 52, 56, 62, 64, 80, 96, 98, 110, 115, 117, 133, 138, 140, 154,
		159, 161, 170,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MonkeyParserInit initializes any static state used to implement MonkeyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMonkeyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MonkeyParserInit() {
	staticData := &MonkeyParserParserStaticData
	staticData.once.Do(monkeyparserParserInit)
}

// NewMonkeyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMonkeyParser(input antlr.TokenStream) *MonkeyParser {
	MonkeyParserInit()
	this := new(MonkeyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MonkeyParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MonkeyParser.g4"

	return this
}

// MonkeyParser tokens.
const (
	MonkeyParserEOF            = antlr.TokenEOF
	MonkeyParserFN             = 1
	MonkeyParserIF             = 2
	MonkeyParserELSE           = 3
	MonkeyParserWHILE          = 4
	MonkeyParserLET            = 5
	MonkeyParserRETURN         = 6
	MonkeyParserASSIGN         = 7
	MonkeyParserPLUS           = 8
	MonkeyParserMINUS          = 9
	MonkeyParserBANG           = 10
	MonkeyParserSTAR           = 11
	MonkeyParserSLASH          = 12
	MonkeyParserPERCENT        = 13
	MonkeyParserEND            = 14
	MonkeyParserLT             = 15
	MonkeyParserGT             = 16
	MonkeyParserGE             = 17
	MonkeyParserLE             = 18
	MonkeyParserEQ             = 19
	MonkeyParserNOT_EQ         = 20
	MonkeyParserCOMMA          = 21
	MonkeyParserCOLON          = 22
	MonkeyParserBooleanLiteral = 23
	MonkeyParserStringLiteral  = 24
	MonkeyParserIdentifier     = 25
	MonkeyParserFloatLiteral   = 26
	MonkeyParserIntegerLiteral = 27
	MonkeyParserLPAREN         = 28
	MonkeyParserRPAREN         = 29
	MonkeyParserLBRACE         = 30
	MonkeyParserRBRACE         = 31
	MonkeyParserLBRACKET       = 32
	MonkeyParserRBRACKET       = 33
	MonkeyParserWS             = 34
)

// MonkeyParser rules.
const (
	MonkeyParserRULE_prog             = 0
	MonkeyParserRULE_statement        = 1
	MonkeyParserRULE_expression       = 2
	MonkeyParserRULE_identifier       = 3
	MonkeyParserRULE_literal          = 4
	MonkeyParserRULE_expression_list  = 5
	MonkeyParserRULE_expression_pair  = 6
	MonkeyParserRULE_array_literal    = 7
	MonkeyParserRULE_hash_literal     = 8
	MonkeyParserRULE_function_literal = 9
	MonkeyParserRULE_params           = 10
	MonkeyParserRULE_if_expression    = 11
	MonkeyParserRULE_call_expression  = 12
	MonkeyParserRULE_let_statement    = 13
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MonkeyParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5897192566) != 0) {
		{
			p.SetState(28)
			p.Statement()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WhileStatementContext struct {
	StatementContext
}

func NewWhileStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStatementContext {
	var p = new(WhileStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserWHILE, 0)
}

func (s *WhileStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLPAREN, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRPAREN, 0)
}

func (s *WhileStatementContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompoundStatementContext struct {
	StatementContext
}

func NewCompoundStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompoundStatementContext {
	var p = new(CompoundStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *CompoundStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACE, 0)
}

func (s *CompoundStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACE, 0)
}

func (s *CompoundStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *CompoundStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CompoundStatementContext) AllEND() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserEND)
}

func (s *CompoundStatementContext) END(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserEND, i)
}

func (s *CompoundStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterCompoundStatement(s)
	}
}

func (s *CompoundStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitCompoundStatement(s)
	}
}

func (s *CompoundStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitCompoundStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type LetStatementContext struct {
	StatementContext
}

func NewLetStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetStatementContext {
	var p = new(LetStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *LetStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetStatementContext) LET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLET, 0)
}

func (s *LetStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIdentifier, 0)
}

func (s *LetStatementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserASSIGN, 0)
}

func (s *LetStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LetStatementContext) END() antlr.TerminalNode {
	return s.GetToken(MonkeyParserEND, 0)
}

func (s *LetStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterLetStatement(s)
	}
}

func (s *LetStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitLetStatement(s)
	}
}

func (s *LetStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitLetStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionStatementContext struct {
	StatementContext
}

func NewExpressionStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) END() antlr.TerminalNode {
	return s.GetToken(MonkeyParserEND, 0)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitExpressionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStatementContext struct {
	StatementContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRETURN, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) END() antlr.TerminalNode {
	return s.GetToken(MonkeyParserEND, 0)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MonkeyParserRULE_statement)
	var _la int

	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Match(MonkeyParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.expression(0)
		}
		{
			p.SetState(35)
			p.Match(MonkeyParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewLetStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(MonkeyParserLET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Match(MonkeyParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(MonkeyParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.expression(0)
		}
		{
			p.SetState(41)
			p.Match(MonkeyParserEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewWhileStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)
			p.Match(MonkeyParserWHILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Match(MonkeyParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.expression(0)
		}
		{
			p.SetState(46)
			p.Match(MonkeyParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Statement()
		}

	case 4:
		localctx = NewCompoundStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(49)
			p.Match(MonkeyParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5897192566) != 0 {
			{
				p.SetState(50)
				p.Statement()
			}
			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MonkeyParserEND {
				{
					p.SetState(51)
					p.Match(MonkeyParserEND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(59)
			p.Match(MonkeyParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewExpressionStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(60)
			p.expression(0)
		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(61)
				p.Match(MonkeyParserEND)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentifierExpressionContext struct {
	ExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegatedExpressionContext struct {
	ExpressionContext
}

func NewNegatedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegatedExpressionContext {
	var p = new(NegatedExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NegatedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegatedExpressionContext) BANG() antlr.TerminalNode {
	return s.GetToken(MonkeyParserBANG, 0)
}

func (s *NegatedExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NegatedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterNegatedExpression(s)
	}
}

func (s *NegatedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitNegatedExpression(s)
	}
}

func (s *NegatedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitNegatedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewRelationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationExpressionContext {
	var p = new(RelationExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RelationExpressionContext) GetOp() antlr.Token { return s.op }

func (s *RelationExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RelationExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RelationExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserGT, 0)
}

func (s *RelationExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLT, 0)
}

func (s *RelationExpressionContext) GE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserGE, 0)
}

func (s *RelationExpressionContext) LE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLE, 0)
}

func (s *RelationExpressionContext) EQ() antlr.TerminalNode {
	return s.GetToken(MonkeyParserEQ, 0)
}

func (s *RelationExpressionContext) NOT_EQ() antlr.TerminalNode {
	return s.GetToken(MonkeyParserNOT_EQ, 0)
}

func (s *RelationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterRelationExpression(s)
	}
}

func (s *RelationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitRelationExpression(s)
	}
}

func (s *RelationExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitRelationExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpressionContext struct {
	ExpressionContext
}

func NewParenExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpressionContext {
	var p = new(ParenExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLPAREN, 0)
}

func (s *ParenExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRPAREN, 0)
}

func (s *ParenExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterParenExpression(s)
	}
}

func (s *ParenExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitParenExpression(s)
	}
}

func (s *ParenExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitParenExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionExpressionContext struct {
	ExpressionContext
}

func NewFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) Function_literal() IFunction_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_literalContext)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitFunctionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type StarSlashExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewStarSlashExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StarSlashExpressionContext {
	var p = new(StarSlashExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StarSlashExpressionContext) GetOp() antlr.Token { return s.op }

func (s *StarSlashExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *StarSlashExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarSlashExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *StarSlashExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StarSlashExpressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSTAR, 0)
}

func (s *StarSlashExpressionContext) SLASH() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSLASH, 0)
}

func (s *StarSlashExpressionContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPERCENT, 0)
}

func (s *StarSlashExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterStarSlashExpression(s)
	}
}

func (s *StarSlashExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitStarSlashExpression(s)
	}
}

func (s *StarSlashExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitStarSlashExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExpressionContext struct {
	ExpressionContext
}

func NewCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExpressionContext {
	var p = new(CallExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExpressionContext) Call_expression() ICall_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_expressionContext)
}

func (s *CallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterCallExpression(s)
	}
}

func (s *CallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitCallExpression(s)
	}
}

func (s *CallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayLiteralExpressionContext struct {
	ExpressionContext
}

func NewArrayLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLiteralExpressionContext {
	var p = new(ArrayLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArrayLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralExpressionContext) Array_literal() IArray_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_literalContext)
}

func (s *ArrayLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterArrayLiteralExpression(s)
	}
}

func (s *ArrayLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitArrayLiteralExpression(s)
	}
}

func (s *ArrayLiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitArrayLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type PlusMinusExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewPlusMinusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusMinusExpressionContext {
	var p = new(PlusMinusExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PlusMinusExpressionContext) GetOp() antlr.Token { return s.op }

func (s *PlusMinusExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *PlusMinusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusMinusExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PlusMinusExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PlusMinusExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserPLUS, 0)
}

func (s *PlusMinusExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MonkeyParserMINUS, 0)
}

func (s *PlusMinusExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterPlusMinusExpression(s)
	}
}

func (s *PlusMinusExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitPlusMinusExpression(s)
	}
}

func (s *PlusMinusExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitPlusMinusExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type HashLiteralExpressionContext struct {
	ExpressionContext
}

func NewHashLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HashLiteralExpressionContext {
	var p = new(HashLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *HashLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HashLiteralExpressionContext) Hash_literal() IHash_literalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHash_literalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHash_literalContext)
}

func (s *HashLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterHashLiteralExpression(s)
	}
}

func (s *HashLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitHashLiteralExpression(s)
	}
}

func (s *HashLiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitHashLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexExpressionContext struct {
	ExpressionContext
}

func NewIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExpressionContext {
	var p = new(IndexExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IndexExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexExpressionContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACKET, 0)
}

func (s *IndexExpressionContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACKET, 0)
}

func (s *IndexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterIndexExpression(s)
	}
}

func (s *IndexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitIndexExpression(s)
	}
}

func (s *IndexExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitIndexExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfExpressionContext struct {
	ExpressionContext
}

func NewIfExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfExpressionContext {
	var p = new(IfExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IfExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExpressionContext) If_expression() IIf_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_expressionContext)
}

func (s *IfExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterIfExpression(s)
	}
}

func (s *IfExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitIfExpression(s)
	}
}

func (s *IfExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitIfExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExpressionContext struct {
	ExpressionContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *MonkeyParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, MonkeyParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(67)
			p.If_expression()
		}

	case 2:
		localctx = NewFunctionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Function_literal()
		}

	case 3:
		localctx = NewCallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.Call_expression()
		}

	case 4:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(70)
			p.Identifier()
		}

	case 5:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(71)
			p.Literal()
		}

	case 6:
		localctx = NewArrayLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(72)
			p.Array_literal()
		}

	case 7:
		localctx = NewHashLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(73)
			p.Hash_literal()
		}

	case 8:
		localctx = NewParenExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(74)
			p.Match(MonkeyParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.expression(0)
		}
		{
			p.SetState(76)
			p.Match(MonkeyParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewNegatedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(78)
			p.Match(MonkeyParserBANG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.expression(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewStarSlashExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expression)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(83)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*StarSlashExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14336) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*StarSlashExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(84)
					p.expression(6)
				}

			case 2:
				localctx = NewPlusMinusExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expression)
				p.SetState(85)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(86)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*PlusMinusExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserPLUS || _la == MonkeyParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*PlusMinusExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(87)
					p.expression(5)
				}

			case 3:
				localctx = NewRelationExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expression)
				p.SetState(88)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(89)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2064384) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(90)
					p.expression(4)
				}

			case 4:
				localctx = NewIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expression)
				p.SetState(91)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(92)
					p.Match(MonkeyParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(93)
					p.expression(0)
				}
				{
					p.SetState(94)
					p.Match(MonkeyParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIdentifier, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MonkeyParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(MonkeyParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntegerLiteral() antlr.TerminalNode
	FloatLiteral() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode
	BooleanLiteral() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIntegerLiteral, 0)
}

func (s *LiteralContext) FloatLiteral() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFloatLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(MonkeyParserStringLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(MonkeyParserBooleanLiteral, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MonkeyParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&226492416) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpression_listContext is an interface to support dynamic dispatch.
type IExpression_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExpression_listContext differentiates from other interfaces.
	IsExpression_listContext()
}

type Expression_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_listContext() *Expression_listContext {
	var p = new(Expression_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expression_list
	return p
}

func InitEmptyExpression_listContext(p *Expression_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expression_list
}

func (*Expression_listContext) IsExpression_listContext() {}

func NewExpression_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_listContext {
	var p = new(Expression_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_expression_list

	return p
}

func (s *Expression_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_listContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Expression_listContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expression_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserCOMMA)
}

func (s *Expression_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOMMA, i)
}

func (s *Expression_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterExpression_list(s)
	}
}

func (s *Expression_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitExpression_list(s)
	}
}

func (s *Expression_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitExpression_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Expression_list() (localctx IExpression_listContext) {
	localctx = NewExpression_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MonkeyParserRULE_expression_list)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(105)
				p.expression(0)
			}
			{
				p.SetState(106)
				p.Match(MonkeyParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5897192454) != 0 {
		{
			p.SetState(113)
			p.expression(0)
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MonkeyParserCOMMA {
			{
				p.SetState(114)
				p.Match(MonkeyParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpression_pairContext is an interface to support dynamic dispatch.
type IExpression_pairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COLON() antlr.TerminalNode

	// IsExpression_pairContext differentiates from other interfaces.
	IsExpression_pairContext()
}

type Expression_pairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_pairContext() *Expression_pairContext {
	var p = new(Expression_pairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expression_pair
	return p
}

func InitEmptyExpression_pairContext(p *Expression_pairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expression_pair
}

func (*Expression_pairContext) IsExpression_pairContext() {}

func NewExpression_pairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_pairContext {
	var p = new(Expression_pairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_expression_pair

	return p
}

func (s *Expression_pairContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_pairContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Expression_pairContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expression_pairContext) COLON() antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOLON, 0)
}

func (s *Expression_pairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_pairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_pairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterExpression_pair(s)
	}
}

func (s *Expression_pairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitExpression_pair(s)
	}
}

func (s *Expression_pairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitExpression_pair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Expression_pair() (localctx IExpression_pairContext) {
	localctx = NewExpression_pairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MonkeyParserRULE_expression_pair)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.expression(0)
	}
	{
		p.SetState(120)
		p.Match(MonkeyParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_literalContext is an interface to support dynamic dispatch.
type IArray_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	Expression_list() IExpression_listContext
	RBRACKET() antlr.TerminalNode

	// IsArray_literalContext differentiates from other interfaces.
	IsArray_literalContext()
}

type Array_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_literalContext() *Array_literalContext {
	var p = new(Array_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_array_literal
	return p
}

func InitEmptyArray_literalContext(p *Array_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_array_literal
}

func (*Array_literalContext) IsArray_literalContext() {}

func NewArray_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_literalContext {
	var p = new(Array_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_array_literal

	return p
}

func (s *Array_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_literalContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACKET, 0)
}

func (s *Array_literalContext) Expression_list() IExpression_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_listContext)
}

func (s *Array_literalContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACKET, 0)
}

func (s *Array_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterArray_literal(s)
	}
}

func (s *Array_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitArray_literal(s)
	}
}

func (s *Array_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitArray_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Array_literal() (localctx IArray_literalContext) {
	localctx = NewArray_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MonkeyParserRULE_array_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(MonkeyParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Expression_list()
	}
	{
		p.SetState(125)
		p.Match(MonkeyParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHash_literalContext is an interface to support dynamic dispatch.
type IHash_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllExpression_pair() []IExpression_pairContext
	Expression_pair(i int) IExpression_pairContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsHash_literalContext differentiates from other interfaces.
	IsHash_literalContext()
}

type Hash_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHash_literalContext() *Hash_literalContext {
	var p = new(Hash_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_hash_literal
	return p
}

func InitEmptyHash_literalContext(p *Hash_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_hash_literal
}

func (*Hash_literalContext) IsHash_literalContext() {}

func NewHash_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Hash_literalContext {
	var p = new(Hash_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_hash_literal

	return p
}

func (s *Hash_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Hash_literalContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLBRACE, 0)
}

func (s *Hash_literalContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRBRACE, 0)
}

func (s *Hash_literalContext) AllExpression_pair() []IExpression_pairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpression_pairContext); ok {
			len++
		}
	}

	tst := make([]IExpression_pairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpression_pairContext); ok {
			tst[i] = t.(IExpression_pairContext)
			i++
		}
	}

	return tst
}

func (s *Hash_literalContext) Expression_pair(i int) IExpression_pairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_pairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_pairContext)
}

func (s *Hash_literalContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserCOMMA)
}

func (s *Hash_literalContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOMMA, i)
}

func (s *Hash_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Hash_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Hash_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterHash_literal(s)
	}
}

func (s *Hash_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitHash_literal(s)
	}
}

func (s *Hash_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitHash_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Hash_literal() (localctx IHash_literalContext) {
	localctx = NewHash_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MonkeyParserRULE_hash_literal)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(MonkeyParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(128)
				p.Expression_pair()
			}
			{
				p.SetState(129)
				p.Match(MonkeyParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5897192454) != 0 {
		{
			p.SetState(136)
			p.Expression_pair()
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MonkeyParserCOMMA {
			{
				p.SetState(137)
				p.Match(MonkeyParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(142)
		p.Match(MonkeyParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_literalContext is an interface to support dynamic dispatch.
type IFunction_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Params() IParamsContext
	RPAREN() antlr.TerminalNode
	Statement() IStatementContext

	// IsFunction_literalContext differentiates from other interfaces.
	IsFunction_literalContext()
}

type Function_literalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_literalContext() *Function_literalContext {
	var p = new(Function_literalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_function_literal
	return p
}

func InitEmptyFunction_literalContext(p *Function_literalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_function_literal
}

func (*Function_literalContext) IsFunction_literalContext() {}

func NewFunction_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_literalContext {
	var p = new(Function_literalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_function_literal

	return p
}

func (s *Function_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_literalContext) FN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserFN, 0)
}

func (s *Function_literalContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLPAREN, 0)
}

func (s *Function_literalContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *Function_literalContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRPAREN, 0)
}

func (s *Function_literalContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Function_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterFunction_literal(s)
	}
}

func (s *Function_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitFunction_literal(s)
	}
}

func (s *Function_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitFunction_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Function_literal() (localctx IFunction_literalContext) {
	localctx = NewFunction_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MonkeyParserRULE_function_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(MonkeyParserFN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(MonkeyParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Params()
	}
	{
		p.SetState(147)
		p.Match(MonkeyParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Statement()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserIdentifier)
}

func (s *ParamsContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserIdentifier, i)
}

func (s *ParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserCOMMA)
}

func (s *ParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserCOMMA, i)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitParams(s)
	}
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MonkeyParserRULE_params)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(150)
				p.Match(MonkeyParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(151)
				p.Match(MonkeyParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MonkeyParserIdentifier {
		{
			p.SetState(157)
			p.Match(MonkeyParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MonkeyParserCOMMA {
			{
				p.SetState(158)
				p.Match(MonkeyParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_expressionContext is an interface to support dynamic dispatch.
type IIf_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	ELSE() antlr.TerminalNode

	// IsIf_expressionContext differentiates from other interfaces.
	IsIf_expressionContext()
}

type If_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_expressionContext() *If_expressionContext {
	var p = new(If_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_if_expression
	return p
}

func InitEmptyIf_expressionContext(p *If_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_if_expression
}

func (*If_expressionContext) IsIf_expressionContext() {}

func NewIf_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_expressionContext {
	var p = new(If_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_if_expression

	return p
}

func (s *If_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *If_expressionContext) IF() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIF, 0)
}

func (s *If_expressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLPAREN, 0)
}

func (s *If_expressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *If_expressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRPAREN, 0)
}

func (s *If_expressionContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *If_expressionContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *If_expressionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MonkeyParserELSE, 0)
}

func (s *If_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterIf_expression(s)
	}
}

func (s *If_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitIf_expression(s)
	}
}

func (s *If_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitIf_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) If_expression() (localctx IIf_expressionContext) {
	localctx = NewIf_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MonkeyParserRULE_if_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(MonkeyParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.Match(MonkeyParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.expression(0)
	}
	{
		p.SetState(166)
		p.Match(MonkeyParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Statement()
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(168)
			p.Match(MonkeyParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Statement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICall_expressionContext is an interface to support dynamic dispatch.
type ICall_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	LPAREN() antlr.TerminalNode
	Expression_list() IExpression_listContext
	RPAREN() antlr.TerminalNode

	// IsCall_expressionContext differentiates from other interfaces.
	IsCall_expressionContext()
}

type Call_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_expressionContext() *Call_expressionContext {
	var p = new(Call_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_call_expression
	return p
}

func InitEmptyCall_expressionContext(p *Call_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_call_expression
}

func (*Call_expressionContext) IsCall_expressionContext() {}

func NewCall_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_expressionContext {
	var p = new(Call_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_call_expression

	return p
}

func (s *Call_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_expressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Call_expressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLPAREN, 0)
}

func (s *Call_expressionContext) Expression_list() IExpression_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_listContext)
}

func (s *Call_expressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserRPAREN, 0)
}

func (s *Call_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterCall_expression(s)
	}
}

func (s *Call_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitCall_expression(s)
	}
}

func (s *Call_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitCall_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Call_expression() (localctx ICall_expressionContext) {
	localctx = NewCall_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MonkeyParserRULE_call_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Identifier()
	}
	{
		p.SetState(173)
		p.Match(MonkeyParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Expression_list()
	}
	{
		p.SetState(175)
		p.Match(MonkeyParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILet_statementContext is an interface to support dynamic dispatch.
type ILet_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	LET() antlr.TerminalNode
	Identifier() IIdentifierContext
	Expression() IExpressionContext
	END() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode

	// IsLet_statementContext differentiates from other interfaces.
	IsLet_statementContext()
}

type Let_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyLet_statementContext() *Let_statementContext {
	var p = new(Let_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_let_statement
	return p
}

func InitEmptyLet_statementContext(p *Let_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_let_statement
}

func (*Let_statementContext) IsLet_statementContext() {}

func NewLet_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Let_statementContext {
	var p = new(Let_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_let_statement

	return p
}

func (s *Let_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Let_statementContext) GetOp() antlr.Token { return s.op }

func (s *Let_statementContext) SetOp(v antlr.Token) { s.op = v }

func (s *Let_statementContext) LET() antlr.TerminalNode {
	return s.GetToken(MonkeyParserLET, 0)
}

func (s *Let_statementContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Let_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Let_statementContext) END() antlr.TerminalNode {
	return s.GetToken(MonkeyParserEND, 0)
}

func (s *Let_statementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(MonkeyParserASSIGN, 0)
}

func (s *Let_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Let_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Let_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.EnterLet_statement(s)
	}
}

func (s *Let_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyParserListener); ok {
		listenerT.ExitLet_statement(s)
	}
}

func (s *Let_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MonkeyParserVisitor:
		return t.VisitLet_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *MonkeyParser) Let_statement() (localctx ILet_statementContext) {
	localctx = NewLet_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MonkeyParserRULE_let_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(MonkeyParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Identifier()
	}
	{
		p.SetState(179)

		var _m = p.Match(MonkeyParserASSIGN)

		localctx.(*Let_statementContext).op = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.expression(0)
	}
	{
		p.SetState(181)
		p.Match(MonkeyParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *MonkeyParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MonkeyParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
