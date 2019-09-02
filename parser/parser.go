// Code generated by goyacc -o parser.go -v y.output parser.y. DO NOT EDIT.

//line parser.y:2
package parser

import __yyfmt__ "fmt"

//line parser.y:2

func setResult(l yyLexer, v *Node) {
	l.(*lexer).result = v
}

//line parser.y:10
type yySymType struct {
	yys int
	n   *Node
	b   bool
	i   int64
	f   float64
	s   string
	sa  []string
	va  []NVar
}

const IDENT = 57346
const ENV = 57347
const CALL = 57348
const CALLCONTRACT = 57349
const INDEX = 57350
const INT = 57351
const FLOAT = 57352
const STRING = 57353
const QSTRING = 57354
const TRUE = 57355
const FALSE = 57356
const NEWLINE = 57357
const COMMA = 57358
const COLON = 57359
const LPAREN = 57360
const RPAREN = 57361
const OBJ = 57362
const LBRACE = 57363
const RBRACE = 57364
const LBRACKET = 57365
const RBRACKET = 57366
const QUESTION = 57367
const DOUBLEDOT = 57368
const DOT = 57369
const ADD = 57370
const SUB = 57371
const MUL = 57372
const DIV = 57373
const MOD = 57374
const ADD_ASSIGN = 57375
const SUB_ASSIGN = 57376
const MUL_ASSIGN = 57377
const DIV_ASSIGN = 57378
const MOD_ASSIGN = 57379
const ASSIGN = 57380
const AND = 57381
const OR = 57382
const EQ = 57383
const NOT_EQ = 57384
const NOT = 57385
const LT = 57386
const GT = 57387
const LTE = 57388
const GTE = 57389
const BREAK = 57390
const CONTINUE = 57391
const DATA = 57392
const CONTRACT = 57393
const IF = 57394
const ELIF = 57395
const ELSE = 57396
const RETURN = 57397
const WHILE = 57398
const FUNC = 57399
const FOR = 57400
const IN = 57401
const SWITCH = 57402
const CASE = 57403
const READ = 57404
const DEFAULT = 57405
const VAR = 57406
const UNARYMINUS = 57407
const UNARYNOT = 57408

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"ENV",
	"CALL",
	"CALLCONTRACT",
	"INDEX",
	"INT",
	"FLOAT",
	"STRING",
	"QSTRING",
	"TRUE",
	"FALSE",
	"NEWLINE",
	"COMMA",
	"COLON",
	"LPAREN",
	"RPAREN",
	"OBJ",
	"LBRACE",
	"RBRACE",
	"LBRACKET",
	"RBRACKET",
	"QUESTION",
	"DOUBLEDOT",
	"DOT",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"MOD",
	"ADD_ASSIGN",
	"SUB_ASSIGN",
	"MUL_ASSIGN",
	"DIV_ASSIGN",
	"MOD_ASSIGN",
	"ASSIGN",
	"AND",
	"OR",
	"EQ",
	"NOT_EQ",
	"NOT",
	"LT",
	"GT",
	"LTE",
	"GTE",
	"BREAK",
	"CONTINUE",
	"DATA",
	"CONTRACT",
	"IF",
	"ELIF",
	"ELSE",
	"RETURN",
	"WHILE",
	"FUNC",
	"FOR",
	"IN",
	"SWITCH",
	"CASE",
	"READ",
	"DEFAULT",
	"VAR",
	"UNARYMINUS",
	"UNARYNOT",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 831

var yyAct = [...]int{

	67, 92, 45, 111, 73, 159, 161, 74, 163, 6,
	174, 190, 192, 2, 11, 13, 20, 112, 120, 35,
	68, 58, 66, 62, 160, 106, 63, 64, 78, 79,
	76, 77, 80, 71, 58, 12, 59, 89, 142, 61,
	53, 54, 55, 56, 57, 52, 143, 74, 201, 93,
	178, 96, 97, 98, 99, 100, 101, 102, 103, 104,
	105, 119, 76, 77, 80, 91, 90, 144, 169, 142,
	14, 7, 152, 145, 186, 151, 179, 126, 127, 128,
	129, 130, 131, 132, 133, 134, 135, 136, 137, 138,
	116, 114, 123, 141, 140, 78, 79, 76, 77, 80,
	116, 114, 209, 117, 115, 171, 146, 148, 83, 84,
	118, 87, 88, 85, 86, 153, 175, 122, 34, 155,
	156, 33, 8, 3, 149, 150, 165, 158, 95, 65,
	113, 107, 47, 46, 43, 44, 32, 37, 38, 39,
	40, 41, 42, 164, 157, 154, 36, 166, 124, 49,
	69, 48, 20, 20, 70, 60, 170, 50, 4, 5,
	94, 176, 93, 177, 1, 9, 17, 162, 125, 16,
	191, 51, 182, 180, 184, 185, 19, 10, 110, 72,
	189, 0, 0, 0, 20, 181, 20, 0, 183, 199,
	20, 200, 0, 0, 187, 0, 0, 0, 0, 0,
	20, 0, 197, 0, 0, 20, 20, 0, 0, 20,
	0, 202, 203, 20, 0, 0, 31, 206, 28, 29,
	32, 210, 0, 0, 0, 0, 0, 15, 0, 31,
	0, 28, 29, 32, 212, 0, 0, 0, 0, 0,
	15, 0, 0, 0, 0, 0, 0, 211, 31, 0,
	28, 29, 32, 0, 0, 0, 0, 0, 0, 15,
	23, 24, 0, 0, 22, 0, 208, 25, 26, 27,
	30, 0, 18, 23, 24, 0, 21, 22, 0, 0,
	25, 26, 27, 30, 0, 18, 0, 0, 0, 21,
	0, 0, 23, 24, 0, 0, 22, 0, 0, 25,
	26, 27, 30, 31, 18, 28, 29, 32, 21, 0,
	0, 0, 0, 0, 15, 0, 31, 0, 28, 29,
	32, 207, 0, 0, 0, 0, 0, 15, 0, 0,
	0, 0, 0, 0, 204, 31, 0, 28, 29, 32,
	0, 0, 0, 0, 0, 0, 15, 23, 24, 0,
	0, 22, 0, 198, 25, 26, 27, 30, 0, 18,
	23, 24, 0, 21, 22, 0, 0, 25, 26, 27,
	30, 0, 18, 0, 0, 0, 21, 0, 0, 23,
	24, 0, 0, 22, 0, 0, 25, 26, 27, 30,
	31, 18, 28, 29, 32, 21, 0, 0, 0, 0,
	0, 15, 0, 31, 0, 28, 29, 32, 194, 0,
	0, 0, 0, 0, 15, 0, 0, 0, 0, 0,
	0, 193, 31, 0, 28, 29, 32, 0, 0, 0,
	0, 0, 0, 15, 23, 24, 0, 0, 22, 0,
	168, 25, 26, 27, 30, 0, 18, 23, 24, 0,
	21, 22, 0, 0, 25, 26, 27, 30, 31, 18,
	28, 29, 32, 21, 0, 0, 23, 24, 0, 15,
	22, 0, 0, 25, 26, 27, 30, 31, 18, 28,
	29, 32, 21, 0, 0, 0, 0, 0, 15, 0,
	0, 0, 0, 0, 0, 167, 0, 0, 0, 0,
	0, 0, 23, 24, 0, 0, 22, 0, 0, 25,
	26, 27, 30, 0, 18, 0, 0, 0, 21, 0,
	0, 23, 24, 0, 0, 22, 0, 0, 25, 26,
	27, 30, 0, 18, 172, 0, 0, 21, 0, 173,
	0, 78, 79, 76, 77, 80, 0, 0, 0, 0,
	0, 0, 81, 82, 83, 84, 205, 87, 88, 85,
	86, 0, 0, 78, 79, 76, 77, 80, 0, 0,
	0, 0, 0, 0, 81, 82, 83, 84, 196, 87,
	88, 85, 86, 0, 0, 78, 79, 76, 77, 80,
	0, 0, 0, 0, 0, 0, 81, 82, 83, 84,
	195, 87, 88, 85, 86, 0, 0, 78, 79, 76,
	77, 80, 0, 0, 0, 0, 0, 0, 81, 82,
	83, 84, 0, 87, 88, 85, 86, 147, 0, 0,
	0, 78, 79, 76, 77, 80, 0, 0, 0, 0,
	0, 0, 81, 82, 83, 84, 139, 87, 88, 85,
	86, 0, 0, 0, 0, 78, 79, 76, 77, 80,
	0, 0, 0, 0, 0, 0, 81, 82, 83, 84,
	0, 87, 88, 85, 86, 121, 0, 0, 0, 78,
	79, 76, 77, 80, 0, 0, 0, 0, 0, 0,
	81, 82, 83, 84, 109, 87, 88, 85, 86, 0,
	0, 78, 79, 76, 77, 80, 0, 0, 0, 0,
	0, 0, 81, 82, 83, 84, 108, 87, 88, 85,
	86, 0, 0, 78, 79, 76, 77, 80, 0, 0,
	75, 0, 0, 0, 81, 82, 83, 84, 0, 87,
	88, 85, 86, 78, 79, 76, 77, 80, 0, 0,
	0, 0, 0, 0, 81, 82, 83, 84, 0, 87,
	88, 85, 86, 78, 79, 76, 77, 80, 0, 0,
	0, 0, 0, 0, 81, 82, 83, 84, 0, 87,
	88, 85, 86, 47, 46, 43, 44, 32, 37, 38,
	39, 40, 41, 42, 188, 0, 0, 36, 0, 0,
	49, 0, 48, 0, 0, 0, 0, 0, 50, 0,
	0, 78, 79, 76, 77, 80, 0, 0, 0, 0,
	0, 0, 51, 82, 83, 84, 0, 87, 88, 85,
	86,
}
var yyPact = [...]int{

	-38, 108, 154, -1000, -53, 50, -1000, 107, -36, 13,
	-1000, 49, -1000, 454, 106, -1000, -1000, 103, 128, 7,
	-2, 151, 128, -1000, -1000, 128, 128, 123, 128, 146,
	150, -1000, 128, -57, -1000, 715, 128, -1000, -1000, -1000,
	-1000, -1000, -1000, 128, 146, 11, -1000, -1000, 128, 117,
	128, 128, 128, 128, 128, 128, 128, 128, 128, 128,
	-13, 127, 695, 735, 673, 126, 85, 735, 84, 93,
	2, 651, 102, -1000, 144, -1000, 128, 128, 128, 128,
	128, 128, 128, 128, 128, 128, 128, 128, 128, 627,
	75, 74, 22, 735, 51, 89, -1000, -1000, 735, 735,
	735, 735, 735, 735, 603, 735, 128, -1000, -1000, -1000,
	56, -1000, 127, -1000, 128, -1000, 141, -1000, 128, 128,
	140, -1000, -17, 127, -14, -55, -1000, -1000, 32, 32,
	-1000, 783, 67, 0, 0, 0, 0, 0, 0, -1000,
	-1000, -1000, 128, -1000, 115, -1000, 128, -1000, 735, 473,
	418, 47, 126, 735, 88, 735, 513, -49, -1000, 101,
	128, 128, -1000, 29, 735, 59, 735, -1000, -1000, -1000,
	-1000, 128, -1000, 128, 128, -1000, 735, 53, -1000, 779,
	-42, 399, 735, 386, 579, 557, -1000, 331, 128, 735,
	128, -1000, 27, -1000, -1000, -1000, -1000, 312, -1000, 735,
	535, -1000, 299, 244, 87, -1000, 225, -1000, -1000, -1000,
	212, -1000, -1000,
}
var yyPgo = [...]int{

	0, 17, 4, 179, 3, 178, 177, 22, 176, 0,
	173, 170, 169, 168, 167, 166, 15, 20, 165, 164,
	2, 1, 160, 159,
}
var yyR1 = [...]int{

	0, 16, 16, 16, 16, 7, 7, 7, 17, 17,
	17, 8, 20, 20, 11, 11, 10, 10, 13, 13,
	14, 14, 12, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 21, 21, 22, 22, 22, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
	9, 9, 9, 9, 9, 9, 9, 9, 1, 1,
	4, 5, 5, 5, 2, 2, 3, 3, 6, 6,
	18, 23, 23, 19, 19,
}
var yyR2 = [...]int{

	0, 0, 2, 2, 3, 0, 1, 3, 0, 3,
	5, 1, 3, 4, 0, 4, 0, 6, 0, 7,
	0, 4, 5, 3, 3, 3, 3, 3, 3, 3,
	4, 2, 7, 1, 1, 1, 2, 5, 7, 3,
	3, 7, 9, 9, 1, 3, 3, 6, 5, 3,
	1, 1, 1, 1, 1, 1, 3, 3, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 2, 1, 2,
	1, 0, 1, 3, 2, 4, 1, 3, 0, 7,
	2, 0, 1, 7, 2,
}
var yyChk = [...]int{

	-1000, -19, 51, 15, 4, -23, 62, 21, 15, -18,
	-6, 50, 22, -16, 21, 15, -12, -15, 60, -8,
	-20, 64, 52, 48, 49, 55, 56, 57, 6, 7,
	58, 4, 8, 15, 15, -9, 18, 9, 10, 11,
	12, 13, 14, 6, 7, -20, 5, 4, 23, 21,
	29, 43, 38, 33, 34, 35, 36, 37, 23, 38,
	4, -1, -9, -9, -9, 6, -7, -9, -17, 4,
	4, -9, -3, -2, 64, 15, 30, 31, 28, 29,
	32, 39, 40, 41, 42, 46, 47, 44, 45, -9,
	-7, -17, -21, -9, -22, 11, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, 38, 4, 21, 21,
	-5, -4, -1, 4, 16, 19, 16, 19, 17, 59,
	16, 24, 15, -1, 4, -13, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, 19,
	19, 19, 16, 24, 16, 22, 17, 24, -9, -16,
	-16, 19, 16, -9, 4, -9, -9, 4, -2, 22,
	38, 61, -14, 63, -9, 11, -9, 22, 22, 21,
	-4, 17, 21, 26, 59, 15, -9, -21, 21, 17,
	-10, -16, -9, -16, -9, -9, 21, -16, 15, -9,
	53, -11, 54, 22, 22, 21, 21, -16, 22, -9,
	-9, 21, -16, -16, 22, 21, -16, 22, 22, 15,
	-16, 22, 22,
}
var yyDef = [...]int{

	0, -2, 0, 94, 91, 0, 92, 0, 88, 0,
	1, 0, 93, 90, 0, 2, 3, 0, 0, 0,
	0, 0, 0, 33, 34, 35, 0, 0, 5, 8,
	0, 11, 0, 0, 4, 0, 0, 50, 51, 52,
	53, 54, 55, 5, 8, 58, 59, 60, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	78, 31, 0, 36, 0, 81, 0, 6, 0, 0,
	0, 0, 0, 86, 0, 18, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 44, 0, 0, 76, 77, 23, 24,
	25, 26, 27, 28, 0, 29, 0, 79, 1, 1,
	0, 82, 80, 78, 0, 39, 0, 40, 0, 0,
	0, 12, 0, 84, 78, 20, 63, 64, 65, 66,
	67, 68, 69, 70, 71, 72, 73, 74, 75, 49,
	56, 57, 0, 61, 0, 62, 0, 13, 30, 0,
	0, 0, 0, 7, 0, 9, 0, 0, 87, 0,
	0, 0, 22, 0, 45, 0, 46, 16, 37, 1,
	83, 0, 1, 0, 0, 89, 85, 0, 1, 0,
	14, 0, 10, 0, 0, 0, 1, 0, 0, 48,
	0, 32, 0, 38, 41, 1, 1, 0, 21, 47,
	0, 1, 0, 0, 0, 1, 0, 43, 42, 19,
	0, 15, 17,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:130
		{
			yyVAL.n = nil
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:131
		{
			yyVAL.n = yyDollar[1].n
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:132
		{
			yyVAL.n = addStatement(yyDollar[1].n, yyDollar[2].n, yylex)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:133
		{
			yyVAL.n = addStatement(yyDollar[1].n, yyDollar[2].n, yylex)
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:137
		{
			yyVAL.n = nil
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:138
		{
			yyVAL.n = newParam(yyDollar[1].n, yylex)
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:139
		{
			yyVAL.n = addParam(yyDollar[1].n, yyDollar[3].n)
		}
	case 8:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:143
		{
			yyVAL.n = nil
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:144
		{
			yyVAL.n = newContractParam(yyDollar[1].s, yyDollar[3].n, yylex)
		}
	case 10:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:145
		{
			yyVAL.n = addContractParam(yyDollar[1].n, yyDollar[3].s, yyDollar[5].n)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:149
		{
			yyVAL.n = newVarValue(yyDollar[1].s, yylex)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:152
		{
			yyVAL.n = newIndex(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:153
		{
			yyVAL.n = addIndex(yyDollar[1].n, yyDollar[3].n, yylex)
		}
	case 14:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:156
		{
			yyVAL.n = nil
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:157
		{
			yyVAL.n = yyDollar[3].n
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:161
		{
			yyVAL.n = nil
		}
	case 17:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:162
		{
			yyVAL.n = newElif(yyDollar[1].n, yyDollar[3].n, yyDollar[5].n, yylex)
		}
	case 18:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:166
		{
			yyVAL.n = nil
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:167
		{
			yyVAL.n = newCase(yyDollar[1].n, yyDollar[3].n, yyDollar[5].n, yylex)
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:171
		{
			yyVAL.n = nil
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:172
		{
			yyVAL.n = yyDollar[3].n
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:176
		{
			yyVAL.n = newSwitch(yyDollar[2].n, yyDollar[4].n, yyDollar[5].n, yylex)
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:180
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ASSIGN, yylex)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:181
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ADD_ASSIGN, yylex)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:182
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, SUB_ASSIGN, yylex)
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:183
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MUL_ASSIGN, yylex)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:184
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, DIV_ASSIGN, yylex)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:185
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MOD_ASSIGN, yylex)
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:186
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ASSIGN, yylex)
		}
	case 30:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:187
		{
			yyVAL.n = newBinary(newVarDecl([]string{yyDollar[2].s}, yylex), yyDollar[4].n, ASSIGN, yylex)
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:188
		{
			yyVAL.n = newVarDecl(yyDollar[2].sa, yylex)
		}
	case 32:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:189
		{
			yyVAL.n = newIf(yyDollar[2].n, yyDollar[4].n, yyDollar[6].n, yyDollar[7].n, yylex)
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:190
		{
			yyVAL.n = newBreak(yylex)
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:191
		{
			yyVAL.n = newContinue(yylex)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:192
		{
			yyVAL.n = newReturn(nil, yylex)
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:193
		{
			yyVAL.n = newReturn(yyDollar[2].n, yylex)
		}
	case 37:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:194
		{
			yyVAL.n = newWhile(yyDollar[2].n, yyDollar[4].n, yylex)
		}
	case 38:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:195
		{
			yyVAL.n = newFunc(yyDollar[2].s, yyDollar[3].va, yyDollar[6].n, yylex)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:198
		{
			yyVAL.n = newCallFunc(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:199
		{
			yyVAL.n = newCallContract(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 41:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:200
		{
			yyVAL.n = newFor(yyDollar[2].s, yyDollar[4].n, yyDollar[6].n, yylex)
		}
	case 42:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.y:201
		{
			yyVAL.n = newForAll(yyDollar[2].s, yyDollar[4].s, yyDollar[6].n, yyDollar[8].n, yylex)
		}
	case 43:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.y:202
		{
			yyVAL.n = newForInt(yyDollar[2].s, yyDollar[4].n, yyDollar[6].n, yyDollar[8].n, yylex)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:206
		{
			yyVAL.n = newArray(yyDollar[1].n, yylex)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:207
		{
			yyVAL.n = appendArray(yyDollar[1].n, yyDollar[3].n, yylex)
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:211
		{
			yyVAL.n = newMap(yyDollar[1].s, yyDollar[3].n, yylex)
		}
	case 47:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.y:212
		{
			yyVAL.n = appendMap(yyDollar[1].n, yyDollar[3].s, yyDollar[6].n, yylex)
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.y:213
		{
			yyVAL.n = appendMap(yyDollar[1].n, yyDollar[3].s, yyDollar[5].n, yylex)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:217
		{
			yyVAL.n = yyDollar[2].n
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:218
		{
			yyVAL.n = newValue(yyDollar[1].i, yylex)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:219
		{
			yyVAL.n = newValue(yyDollar[1].f, yylex)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:220
		{
			yyVAL.n = newValue(yyDollar[1].s, yylex)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:221
		{
			yyVAL.n = newValue(yyDollar[1].s, yylex)
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:222
		{
			yyVAL.n = newValue(true, yylex)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:223
		{
			yyVAL.n = newValue(false, yylex)
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:224
		{
			yyVAL.n = newCallFunc(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:225
		{
			yyVAL.n = newCallContract(yyDollar[1].s, yyDollar[2].n, yylex)
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:226
		{
			yyVAL.n = yyDollar[1].n
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:227
		{
			yyVAL.n = newEnv(yyDollar[1].s, yylex)
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:228
		{
			yyVAL.n = newGetVar(yyDollar[1].s, yylex)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:229
		{
			yyVAL.n = yyDollar[2].n
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:230
		{
			yyVAL.n = yyDollar[2].n
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:231
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MUL, yylex)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:232
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, DIV, yylex)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:233
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, ADD, yylex)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:234
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, SUB, yylex)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:235
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, MOD, yylex)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:236
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, AND, yylex)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:237
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, OR, yylex)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:238
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, EQ, yylex)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:239
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, NOT_EQ, yylex)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:240
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, LTE, yylex)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:241
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, GTE, yylex)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:242
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, LT, yylex)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:243
		{
			yyVAL.n = newBinary(yyDollar[1].n, yyDollar[3].n, GT, yylex)
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:245
		{
			yyVAL.n = newUnary(yyDollar[2].n, SUB, yylex)
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:246
		{
			yyVAL.n = newUnary(yyDollar[2].n, NOT, yylex)
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:250
		{
			yyVAL.sa = []string{yyDollar[1].s}
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:251
		{
			yyVAL.sa = append(yyDollar[1].sa, yyDollar[2].s)
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:255
		{
			yyVAL.va = newVars(yyDollar[1].sa)
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:259
		{
			yyVAL.va = nil
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:260
		{
			yyVAL.va = yyDollar[1].va
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:261
		{
			yyVAL.va = append(yyDollar[1].va, yyDollar[3].va...)
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:265
		{
			yyVAL.va = newVars(yyDollar[2].sa)
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.y:266
		{
			yyVAL.va = newVarExp(yyDollar[2].s, yyDollar[4].n, yylex)
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:270
		{
			yyVAL.va = yyDollar[1].va
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.y:271
		{
			yyVAL.va = append(yyDollar[1].va, yyDollar[3].va...)
		}
	case 88:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:275
		{
			yyVAL.va = nil
		}
	case 89:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:276
		{
			yyVAL.va = yyDollar[4].va
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.y:280
		{
			yyVAL.n = newBlock(yyDollar[1].va, yyDollar[2].n, yylex)
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.y:286
		{
			yyVAL.b = false
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.y:287
		{
			yyVAL.b = true
		}
	case 93:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.y:291
		{
			yyVAL.n = newContract(yyDollar[2].s, yyDollar[3].b, yyDollar[6].n, yylex)
			setResult(yylex, yyVAL.n)
		}
	}
	goto yystack /* stack new state and value */
}