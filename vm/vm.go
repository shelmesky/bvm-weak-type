package vm

import (
	"bvm/compiler"
	"unsafe"
)

type Value struct {
	Type  int
	Value interface{}
}

type VM struct {
	Constants []Value // 常量
	Vars      []Value // 变量
	Stack     []int64 // 栈
	ESP       int
	EBP       int
}

func Run(cmpl *compiler.CompileEnv) error {

	vm := VM{
		Constants: make([]Value, 0),
		Vars:      make([]Value, 0, 1024),
		Stack:     make([]int64, 100),
		ESP:       0,
		EBP:       0,
	}

	code := cmpl.Code
	length := len(code)

	for i := 0; i < length; i++ {
		switch code[i] {
		case compiler.INITVARS:
			variable := Value{}
			vm.Vars = append(vm.Vars, variable)

		case compiler.PUSH:
			cnst := cmpl.Constants[code[i]]
			value := Value{
				Type:  cnst.Type,
				Value: cnst.Value,
			}
			vm.Constants = append(vm.Constants, value)
			vm.Stack[vm.ESP] = int64(len(vm.Constants) - 1)
			vm.ESP += 1

		case compiler.SETVAR:
			i += 1
			varIndex := code[i]
			vm.Stack[vm.ESP] = int64(uintptr(unsafe.Pointer(&vm.Vars[varIndex])))
			vm.ESP += 1

		case compiler.GETVAR:
			i += 1
			varIndex := code[i]
			vm.Stack[vm.ESP] = int64(varIndex)
			vm.ESP += 1
		}
	}

	return nil
}
