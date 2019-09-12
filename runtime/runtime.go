package runtime

const (
	NOP = iota
	PUSH
	INITVARS
	GETVAR
	SETVAR
	ADD       // +
	SUB       // -
	MUL       // *
	DIV       // /
	ASSIGN    // =
	JMP       // jump to offset
	JZE       // jump if stack top is zero
	RETFUNC   // return from function
	RETURN    // return from contract
	CALLFUNC  // call function
	GETPARAMS // load function params
	CALLEMBED // call embed function
	AND       // a && b
	OR        // a || b
	EQ        // a == b
	NOTEQ     // a != b
	NOT       // !a
	LT        // a < b
	GT        // a > b
	LTE       // a <= b
	GTE       // a >= b
)
