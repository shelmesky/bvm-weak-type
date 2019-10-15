%{
package parser

func setResult(l yyLexer, v *Node) {
  l.(*lexer).result = v
}

%}

%union {
    n       *Node
    b       bool
    i       int64
    f       float64
    s       string
    sa      []string
    va      []NVar
}

// Identifiers + literals
%token<s> IDENT  // foobar
%token<s> ENV   // $foobar
%token<s> CALL  // foobar(
%token<s> CALLCONTRACT  // @foobar(
%token<s> INDEX  // foobar[
%token<i> INT    // 314
%token<f> FLOAT    // 3.14
%token<s> STRING  // "string"
%token<s> QSTRING  // `string`
%token<b> TRUE   // true
%token<b> FALSE  // false

// Delimiters
%token NEWLINE // \n
%token COMMA   // ,
%token COLON   // :
%token LPAREN  // (
%token RPAREN  // )
%token OBJ  // @{
%token LBRACE  // {
%token RBRACE  // }
%token LBRACKET // [
%token RBRACKET // ]
%token QUESTION // ?
%token DOUBLEDOT   // ..
%token DOT   // .

// Operators
%token ADD // +
%token SUB // -
%token MUL // *
%token DIV // /
%token MOD // %

%token BIT_NOT
%token BIT_AND
%token BIT_OR
%token BIT_XOR
%token LEFT_SHIFT
%token RIGHT_SHIFT
%token POW

%token ADD_ASSIGN // +=
%token SUB_ASSIGN // -=
%token MUL_ASSIGN // *=
%token DIV_ASSIGN // /=
%token MOD_ASSIGN // %=
%token LEFT_SHIFT_ASSIGN // <<=
%token RIGHT_SHIFT_ASSIGN // >>=
%token BIT_AND_ASSIGN // &=
%token BIT_XOR_ASSIGN // ^=
%token BIT_OR_ASSIGN // |=
%token ASSIGN // =

%token AND // &&
%token OR  // ||

%token EQ     // ==
%token NOT_EQ // !=
%token NOT    // !

%token LT     // <
%token GT     // >
%token LTE    // <=
%token GTE    // >=

// Keywords
%token BREAK      // break
%token CONTINUE   // continue
%token DATA       // data
%token CONTRACT   // contract
%token IF       // if
%token ELIF     // elif
%token ELSE     // else
%token RETURN   // return
%token WHILE   // while
%token FUNC    // func
%token FOR     // for
%token IN      // in
%token SWITCH  // switch
%token CASE    // case
%token READ    // read
%token DEFAULT // default
%token VAR     // var


%type <sa> ident_list
%type <va> var_declaration
%type <va> var_declarations
%type <va> par_declaration
%type <va> par_declarations
%type <va> contract_data
%type <n> params
%type <n> variable
%type <n> expr
%type <n> elif
%type <n> else
%type <n> switch
%type <n> case
%type <n> default
%type <n> statement
%type <n> statements
%type <n> cntparams
%type <n> contract_body
%type <n> contract_declaration
%type <n> index
%type <n> exprlist
%type <n> exprmaplist
%type <b> contract_read

%left AND 
%left OR
%left LTE GTE LT GT EQ NOT_EQ
%left ADD SUB
%left MUL DIV MOD BIT_AND BIT_OR BIT_XOR LEFT_SHIFT RIGHT_SHIFT
%right POW
%right UNARYMINUS UNARYNOT

%start contract_declaration

%%

statements
    : /*empty*/ { $$ = nil }
    | statements NEWLINE { $$ = $1 }
    | statements switch { $$ = addStatement($1, $2, yylex)}
    | statements statement NEWLINE { $$ = addStatement($1, $2, yylex)}
    ;

params
    : /*empty*/ { $$ = nil }
    | expr { $$ = newParam( $1, yylex ) }
    | params COMMA expr { $$ = addParam($1, $3)}
    ;

cntparams
    : /*empty*/ { $$ = nil }
    | IDENT COLON expr { $$ = newContractParam( $1, $3, yylex ) }
    | cntparams COMMA IDENT COLON expr { $$ = addContractParam($1, $3, $5)}
    ;

variable
    : IDENT { $$ = newVarValue($1, yylex); }

index 
    : INDEX expr RBRACKET { $$ = newIndex($1, $2, yylex);}
    | index LBRACKET expr RBRACKET { $$ = addIndex($1, $3, yylex);}

else 
   : /*empty*/ {$$ = nil}
   | ELSE LBRACE statements RBRACE { $$ = $3 }
   ;

elif
   : /*empty*/ {$$ = nil}
   | elif ELIF expr LBRACE statements RBRACE { $$ = newElif($1, $3, $5, yylex) }
   ;

case
   : /*empty*/ {$$ = nil}
   | case CASE exprlist LBRACE statements RBRACE NEWLINE { $$ = newCase($1, $3, $5, yylex) }
   ;

default 
   : /*empty*/ {$$ = nil}
   | DEFAULT LBRACE statements RBRACE { $$ = $3 }
   ;

switch
    : SWITCH expr NEWLINE case default { $$ = newSwitch( $2, $4, $5, yylex ) }
    ;

statement 
    : variable ASSIGN expr { $$ = newBinary($1, $3, ASSIGN, yylex) }
    | index ASSIGN expr { $$ = newBinary($1, $3, ASSIGN, yylex) }
    | VAR IDENT ASSIGN expr { $$ = newBinary( newVarDecl([]string{$2}, yylex ), $4, ASSIGN, yylex) }

    | variable ADD_ASSIGN expr { $$ = newBinary($1, $3, ADD_ASSIGN, yylex); }
    | index ADD_ASSIGN expr { $$ = newBinary($1, $3, ADD_ASSIGN, yylex) }
    | VAR IDENT ADD_ASSIGN expr { $$ = newBinary( newVarDecl([]string{$2}, yylex ), $4, ADD_ASSIGN, yylex) }

    | variable SUB_ASSIGN expr { $$ = newBinary($1, $3, SUB_ASSIGN, yylex); }
    | index SUB_ASSIGN expr { $$ = newBinary($1, $3, SUB_ASSIGN, yylex) }
    | VAR IDENT SUB_ASSIGN expr { $$ = newBinary( newVarDecl([]string{$2}, yylex ), $4, SUB_ASSIGN, yylex) }

    | variable MUL_ASSIGN expr { $$ = newBinary($1, $3, MUL_ASSIGN, yylex); }
    | index MUL_ASSIGN expr { $$ = newBinary($1, $3, MUL_ASSIGN, yylex) }
    | VAR IDENT MUL_ASSIGN expr { $$ = newBinary( newVarDecl([]string{$2}, yylex ), $4, MUL_ASSIGN, yylex) }


    | variable DIV_ASSIGN expr { $$ = newBinary($1, $3, DIV_ASSIGN, yylex); }
    | index DIV_ASSIGN expr { $$ = newBinary($1, $3, DIV_ASSIGN, yylex) }
    | VAR IDENT DIV_ASSIGN expr { $$ = newBinary( newVarDecl([]string{$2}, yylex ), $4, DIV_ASSIGN, yylex) }

    | variable MOD_ASSIGN expr { $$ = newBinary($1, $3, MOD_ASSIGN, yylex); }
    | index MOD_ASSIGN expr { $$ = newBinary($1, $3, MOD_ASSIGN, yylex) }
    | VAR IDENT MOD_ASSIGN expr { $$ = newBinary( newVarDecl([]string{$2}, yylex ), $4, MOD_ASSIGN, yylex) }

    | variable LEFT_SHIFT_ASSIGN expr { $$ = newBinary($1, $3, LEFT_SHIFT_ASSIGN, yylex); }
    | index LEFT_SHIFT_ASSIGN expr { $$ = newBinary($1, $3, LEFT_SHIFT_ASSIGN, yylex) }
    | VAR IDENT LEFT_SHIFT_ASSIGN expr { $$ = newBinary( newVarDecl([]string{$2}, yylex ), $4, LEFT_SHIFT_ASSIGN, yylex) }

    | variable RIGHT_SHIFT_ASSIGN expr { $$ = newBinary($1, $3, RIGHT_SHIFT_ASSIGN, yylex); }
    | index RIGHT_SHIFT_ASSIGN expr { $$ = newBinary($1, $3, RIGHT_SHIFT_ASSIGN, yylex) }
    | VAR IDENT RIGHT_SHIFT_ASSIGN expr { $$ = newBinary( newVarDecl([]string{$2}, yylex ), $4, RIGHT_SHIFT_ASSIGN, yylex) }

    | variable BIT_AND_ASSIGN expr { $$ = newBinary($1, $3, BIT_AND_ASSIGN, yylex); }
    | index BIT_AND_ASSIGN expr { $$ = newBinary($1, $3, BIT_AND_ASSIGN, yylex) }
    | VAR IDENT BIT_AND_ASSIGN expr { $$ = newBinary( newVarDecl([]string{$2}, yylex ), $4, BIT_AND_ASSIGN, yylex) }

    | variable BIT_XOR_ASSIGN expr { $$ = newBinary($1, $3, BIT_XOR_ASSIGN, yylex); }
    | index BIT_XOR_ASSIGN expr { $$ = newBinary($1, $3, BIT_XOR_ASSIGN, yylex) }
    | VAR IDENT BIT_XOR_ASSIGN expr { $$ = newBinary( newVarDecl([]string{$2}, yylex ), $4, BIT_XOR_ASSIGN, yylex) }

    | variable BIT_OR_ASSIGN expr { $$ = newBinary($1, $3, BIT_OR_ASSIGN, yylex); }
    | index BIT_OR_ASSIGN expr { $$ = newBinary($1, $3, BIT_OR_ASSIGN, yylex) }
    | VAR IDENT BIT_OR_ASSIGN expr { $$ = newBinary( newVarDecl([]string{$2}, yylex ), $4, BIT_OR_ASSIGN, yylex) }

    | VAR ident_list { $$ = newVarDecl($2, yylex )}

    | IF expr LBRACE statements RBRACE elif else { $$ = newIf( $2, $4, $6, $7, yylex )}
    | BREAK { $$ = newBreak(yylex); }
    | CONTINUE { $$ = newContinue(yylex); }
    | RETURN { $$ = newReturn(nil, yylex); }
    | RETURN expr { $$ = newReturn($2, yylex); }
    | WHILE expr LBRACE statements RBRACE { $$ = newWhile( $2, $4, yylex )}

    | FUNC CALL par_declarations RPAREN LBRACE statements RBRACE {
           $$ = newFunc($2, $3, $6, yylex)
           }
    | CALL params RPAREN { $$ = newCallFunc($1, $2, yylex);}
    | CALLCONTRACT cntparams RPAREN { $$ = newCallContract($1, $2, yylex);}

    | FOR IDENT IN expr LBRACE statements RBRACE { $$ = newFor( $2, $4, $6, yylex )}
    | FOR IDENT COMMA IDENT IN expr LBRACE statements RBRACE { $$ = newForAll( $2, $4, $6, $8, yylex )}
    | FOR IDENT IN expr DOUBLEDOT expr LBRACE statements RBRACE { $$ = newForInt( $2, $4, $6, $8, yylex )}
    ;

exprlist
    : expr { $$ = newArray($1, yylex); }
    | exprlist COMMA expr { $$ = appendArray($1, $3, yylex);}
    ;   

exprmaplist
    : STRING COLON expr { $$ = newMap($1, $3, yylex); }
    | exprmaplist COMMA STRING COLON NEWLINE expr { $$ = appendMap($1, $3, $6, yylex); }
    | exprmaplist COMMA STRING COLON expr { $$ = appendMap($1, $3, $5, yylex); }
    ;   

expr
    : LPAREN expr RPAREN { $$ = $2; }
    | INT { $$ = newValue($1, yylex);}
    | FLOAT { $$ = newValue($1, yylex);}
    | STRING { $$ = newValue($1, yylex);}
    | QSTRING { $$ = newValue($1, yylex);}
    | TRUE { $$ = newValue(true, yylex);}
    | FALSE { $$ = newValue(false, yylex);}

    | CALL params RPAREN { $$ = newCallFunc($1, $2, yylex);}
    | CALLCONTRACT cntparams RPAREN { $$ = newCallContract($1, $2, yylex);}

    | index { $$ = $1}
    | ENV { $$ = newEnv($1, yylex);}
    | IDENT { $$ = newGetVar($1, yylex);}
    | LBRACKET exprlist RBRACKET { $$ = $2;}	// [expr, expr ...]
    | LBRACE exprmaplist RBRACE { $$ = $2;}	// {"string": expr, "string": expr}

    | expr MUL expr { $$ = newBinary($1, $3, MUL, yylex); }
    | expr DIV expr { $$ = newBinary($1, $3, DIV, yylex);  }
    | expr ADD expr { $$ = newBinary($1, $3, ADD, yylex); }
    | expr SUB expr { $$ = newBinary($1, $3, SUB, yylex);}
    | expr MOD expr { $$ = newBinary($1, $3, MOD, yylex); } 
    | expr AND expr { $$ = newBinary($1, $3, AND, yylex); }
    | expr OR expr { $$ = newBinary($1, $3, OR, yylex);  }
    | expr EQ expr { $$ = newBinary($1, $3, EQ, yylex); }
    | expr NOT_EQ expr { $$ = newBinary($1, $3, NOT_EQ, yylex);}
    | expr LTE expr { $$ = newBinary($1, $3, LTE, yylex); }
    | expr GTE expr { $$ = newBinary($1, $3, GTE, yylex);  }
    | expr LT expr { $$ = newBinary($1, $3, LT, yylex); }
    | expr GT expr { $$ = newBinary($1, $3, GT, yylex);}
    | expr BIT_AND expr { $$ = newBinary($1, $3, BIT_AND, yylex);}
    | expr BIT_OR expr { $$ = newBinary($1, $3, BIT_OR, yylex);}
    | expr BIT_XOR expr { $$ = newBinary($1, $3, BIT_XOR, yylex);}
    | expr LEFT_SHIFT expr { $$ = newBinary($1, $3, LEFT_SHIFT, yylex);}
    | expr RIGHT_SHIFT expr { $$ = newBinary($1, $3, RIGHT_SHIFT, yylex);}
    | expr POW expr { $$ = newBinary($1, $3, POW, yylex);}

    | SUB expr %prec UNARYMINUS { $$ = newUnary($2, SUB, yylex)}
    | NOT expr %prec UNARYNOT { $$ = newUnary($2, NOT, yylex)}
    ;

ident_list
    : IDENT { $$ = []string{$1} }
    | ident_list IDENT { $$ = append($1, $2) }
    ;

par_declaration
    : ident_list { $$ = newVars($1)}
    ;

par_declarations
    : /*empty*/ {$$=nil}
    | par_declaration { $$ = $1}
    | par_declarations COMMA par_declaration { $$ = append($1, $3...) }
    ;

var_declaration
    : VAR ident_list { $$ = newVars($2) }
    | VAR IDENT ASSIGN expr { $$ = newVarExp($2, $4, yylex) }
    ;

var_declarations
    : var_declaration { $$=$1 }
    | var_declarations NEWLINE var_declaration { $$ = append($1, $3...) }
    ;

contract_data 
    : /*empty*/ { $$ = nil }
    | DATA LBRACE NEWLINE var_declarations NEWLINE RBRACE NEWLINE { $$ = $4 }
    ;

contract_body 
    : contract_data statements {
        $$ = newBlock($1, $2, yylex)
    }
    ;

contract_read
    : /*empty*/ { $$ = false }
    | READ { $$ = true }
    ;

contract_declaration 
    : CONTRACT IDENT contract_read LBRACE NEWLINE contract_body RBRACE { 
        $$ = newContract($2, $3, $6, yylex)
        setResult(yylex, $$)
        }
    | contract_declaration NEWLINE
    ;  
