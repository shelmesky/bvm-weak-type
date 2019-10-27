package runtime

const (
	NOP = iota
	PUSH
	INITVARS
	GETVAR
	SETVAR
	ADD                // +
	SUB                // -
	MUL                // *
	DIV                // /
	MOD                // %
	BIT_AND            // a & b
	BIT_OR             // a | b
	BIT_XOR            // a^b
	LEFT_SHIFT         // a<<b
	RIGHT_SHIFT        // a>>b
	POW                // a**b
	ADD_ASSIGN         // +=
	SUB_ASSIGN         // -=
	MUL_ASSIGN         // *=
	DIV_ASSIGN         // /=
	MOD_ASSIGN         // %=
	LEFT_SHIFT_ASSIGN  // <<=
	RIGHT_SHIFT_ASSIGN // >>=
	BIT_AND_ASSIGN     // &=
	BIT_XOR_ASSIGN     // ^=
	BIT_OR_ASSIGN      // |=
	ASSIGN             // =
	LOOP               // loop start
	JMP                // jump to offset
	JZE                // jump if stack top is zero
	RETFUNC            // return from function
	RETURN             // return from contract
	CALLFUNC           // call function
	GETPARAMS          // load function params
	CALLEMBED          // call embed function
	AND                // a && b
	OR                 // a || b
	EQ                 // a == b
	NOTEQ              // a != b
	NOT                // !a
	LT                 // a < b
	GT                 // a > b
	LTE                // a <= b
	GTE                // a >= b
	INITMAP            // init map type
)
