%yyc c
%yyn c = l.Next()
%yym l.Mark()

%{
package parser

import (
	"strconv"
	"strings"

	"modernc.org/golex/lex"
)

func (l *lexer) scan(lval *yySymType) lex.Char {
	c := l.Enter()
%}

hexint	    	0[xX][0-9a-fA-F]+
digit			[0-9]|{unicodeDigit}
int				[0-9]+
float			{int}\.[0-9]+
identifier		{letter}({letter}|{digit})*
env             ${identifier}
callcontract    @{identifier}\(
call    		{identifier}\(
index           {identifier}\[
string 			\"([^\\"]|\\.)*\"
qstring 		`([^`])*`
letter			[_a-zA-Z]|{unicodeLetter}
unicodeDigit	\x81
unicodeLetter	\x80

%%

	c = l.Rule0()

[ \t\r ]+		 // ignore all whitespace
\/\*(.|\n)*\*\/  // ignore comments
\/\/(.)*\n      return l.char(NEWLINE)

\+=				return l.char(ADD_ASSIGN)
-=				return l.char(SUB_ASSIGN)
\*=				return l.char(MUL_ASSIGN)
\/=				return l.char(DIV_ASSIGN)
%=				return l.char(MOD_ASSIGN)
\<\<=           return l.char(LEFT_SHIFT_ASSIGN)
>>=             return l.char(RIGHT_SHIFT_ASSIGN)
&=              return l.char(BIT_AND_ASSIGN)
\^=             return l.char(BIT_XOR_ASSIGN)
\|=             return l.char(BIT_OR_ASSIGN)
=				return l.char(ASSIGN)

\n				return l.char(NEWLINE)
;				return l.char(NEWLINE)
,[ \t\r]*\n?		return l.char(COMMA)
:				return l.char(COLON)
\.\.			return l.char(DOUBLEDOT)
\.				return l.char(DOT)
\?         	    return l.char(QUESTION)
\([ \t]*\n?     return l.char(LPAREN)
\n[ \t]*\)		return l.char(RPAREN)
\)				return l.char(RPAREN)
@\{[ \t]*\n?  	return l.char(OBJ)
\{[ \t]*\n?  	return l.char(LBRACE)
\}  	return l.char(RBRACE)
\[[ \t]*\n?		return l.char(LBRACKET)
\]				return l.char(RBRACKET)

&&				return l.char(AND)
\|\|			return l.char(OR)

&               return l.char(BIT_AND)
\|              return l.char(BIT_OR)
\^              return l.char(BIT_XOR)
\<\<            return l.char(LEFT_SHIFT)
>>              return l.char(RIGHT_SHIFT)

\*\*            return l.char(POW)

\+[ \t\r]*\n?		return l.char(ADD)
-[ \t\r]*\n?		return l.char(SUB)
\*[ \t\r]*\n?		return l.char(MUL)
\/[ \t\r]*\n?		return l.char(DIV)
%[ \t\r]*\n?		return l.char(MOD)

==[ \t\r]*\n?		return l.char(EQ)
!=[ \t\r]*\n?		return l.char(NOT_EQ)
!				return l.char(NOT)

\<=[ \t\r]*\n?	return l.char(LTE)
>=[ \t\r]*\n?		return l.char(GTE)
\<[ \t\r]*\n?		return l.char(LT)
>[ \t\r]*\n?  	return l.char(GT)

break    		return l.char(BREAK)
continue  		return l.char(CONTINUE)
data    		return l.char(DATA)
contract		return l.char(CONTRACT)
while           return l.char(WHILE)
if				return l.char(IF)
elif			return l.char(ELIF)
else			return l.char(ELSE)
return			return l.char(RETURN)
true			return l.char(TRUE)
false			return l.char(FALSE)
func			return l.char(FUNC)
for				return l.char(FOR)
in				return l.char(IN)
switch  		return l.char(SWITCH)
case  	    	return l.char(CASE)
read    		return l.char(READ)
default  		return l.char(DEFAULT)

var			return l.char(VAR)

{float}			{
					ai, _ := strconv.ParseFloat(string(l.TokenBytes(nil)), 64); 
					lval.f = ai
					return l.char(FLOAT)
				}
{hexint}			{
					val, _ := strconv.ParseInt(string(l.TokenBytes(nil)), 0, 64)
					lval.i = int64(val)
					return l.char(INT)
				}
{int}			{
					ai, _ := strconv.Atoi(string(l.TokenBytes(nil)))
					lval.i = int64(ai)
					return l.char(INT)
				}
{identifier}	{
					lval.s = string(l.TokenBytes(nil))
					return l.char(IDENT)
				}
{env}	{
					lval.s = string(l.TokenBytes(nil))
					lval.s = lval.s[1:]
					return l.char(ENV)
				}
{string}		{
					var err error
					s := strings.Replace( string(l.TokenBytes(nil)), "\n", `\n`, -1 )
					lval.s, err = strconv.Unquote(`"` + s[1:len(s)-1] + `"`)
					if err != nil  {
						l.Error(err.Error())
					}
					return l.char(STRING)
				}				
{qstring}		{
					s := string(l.TokenBytes(nil))
					lval.s = s[1:len(s)-1]
					return l.char(STRING)
				}				
{call}	{
					lval.s = string(l.TokenBytes(nil))
					lval.s = lval.s[:len(lval.s)-1]
					return l.char(CALL)
				}
{callcontract}	{
					lval.s = string(l.TokenBytes(nil))
					lval.s = lval.s[:len(lval.s)-1]
					return l.char(CALLCONTRACT)
				}
{index}	{
					lval.s = string(l.TokenBytes(nil))
					lval.s = lval.s[:len(lval.s)-1]
					return l.char(INDEX)
				}

%%
	if c, ok := l.Abort(); ok {
		return l.char(c)
	}

	goto yyAction
}