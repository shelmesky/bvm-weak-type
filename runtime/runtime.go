package runtime

const (
	NOP = iota
	PUSH
	INITVARS
	GETVAR
	SETVAR
	ADD    // +
	SUB    // -
	MUL    // *
	DIV    // /
	ASSIGN // =
	JMP
	RETFUNC // return from function
	RETURN  // return from contract
)
