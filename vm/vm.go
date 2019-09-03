package vm

import (
	"bvm/compiler"
	"bvm/parser"
	"bvm/runtime"
	"unsafe"
)

const (
	CONST_IDX = iota
	VAR_IDX
	VAR_POINTER
	STACK_TEMP
	FUNC_IDX
)

type Value struct {
	Type  int
	Value interface{}
}

type StackItem struct {
	Type  int
	Value interface{}
}

type VM struct {
	Constants []*Value     // 常量
	Vars      []*Value     // 变量
	Funcs     []*Value     // 函数
	Stack     []*StackItem // 栈
	ESP       int          // 栈指针
	EBP       int          // 栈基址指针
}

func Run(cmpl *compiler.CompileEnv) error {
	var (
		valueA *Value
		valueB *Value
	)

	vm := VM{
		Constants: make([]*Value, 0),
		Vars:      make([]*Value, 0),
		Stack:     make([]*StackItem, 256),
		ESP:       0,
		EBP:       0,
	}

	code := cmpl.Code
	length := len(code)

	for i := 0; i < length; i++ {
		switch code[i] {
		case runtime.INITVARS:
			variable := Value{}
			vm.Vars = append(vm.Vars, &variable)

		case runtime.PUSH:
			i++
			vm.ESP++
			// 根据PUSH的操作数获取在常量表中的常量值
			cnst := cmpl.ConstantsTable[code[i]]
			// 将常量值封装为Value结构
			value := Value{
				Type:  cnst.Type,
				Value: cnst.Value,
			}
			// 将常量保存在vm.Constants数组中
			vm.Constants = append(vm.Constants, &value)
			// 获取当前索引
			idx := int64(len(vm.Constants) - 1)
			// 创建stack元素，类型为CONST_IDX，值为索引
			stackItem := StackItem{
				Type:  CONST_IDX,
				Value: idx,
			}
			vm.Stack[vm.ESP] = &stackItem

		case runtime.SETVAR:
			i++
			vm.ESP++
			// 根据SETVAR的操作数获取变量的索引
			varIndex := code[i]
			// 获取索引在vm.Vars中的指针
			varPointer := int64(uintptr(unsafe.Pointer(&vm.Vars[varIndex])))
			// 保存指针到stack元素
			stackItem := StackItem{
				Type:  VAR_POINTER,
				Value: varPointer,
			}
			vm.Stack[vm.ESP] = &stackItem

		case runtime.GETVAR:
			i++
			vm.ESP++
			varIndex := code[i]
			stackItem := StackItem{
				Type:  VAR_IDX,
				Value: varIndex,
			}
			vm.Stack[vm.ESP] = &stackItem

		case runtime.ADD:
			// 从stack获取栈顶的2个元素
			stackItemA := vm.Stack[vm.ESP-1]
			stackItemB := vm.Stack[vm.ESP]

			// 如果2个元素都是常量， 则从常量表中获取Value
			if stackItemA.Type == CONST_IDX {
				valueA = vm.Constants[stackItemA.Value.(int64)]
			}
			if stackItemB.Type == CONST_IDX {
				valueB = vm.Constants[stackItemB.Value.(int64)]
			}

			if stackItemA.Type == STACK_TEMP {
				valueB = stackItemA.Value.(*Value)
			}
			if stackItemB.Type == STACK_TEMP {
				valueB = stackItemB.Value.(*Value)
			}

			if valueA.Type == parser.VInt && valueB.Type == parser.VInt {
				result := valueA.Value.(int64) + valueB.Value.(int64)
				vm.ESP--
				vm.Stack[vm.ESP] = &StackItem{
					Type: STACK_TEMP,
					Value: &Value{
						Type:  parser.VInt,
						Value: result,
					},
				}
			}

		case runtime.MUL:
			// 从stack获取栈顶的2个元素
			stackItemA := vm.Stack[vm.ESP-1]
			stackItemB := vm.Stack[vm.ESP]

			// 如果2个元素都是常量， 则从常量表中获取Value
			if stackItemA.Type == CONST_IDX {
				valueA = vm.Constants[stackItemA.Value.(int64)]
			}
			if stackItemB.Type == CONST_IDX {
				valueB = vm.Constants[stackItemB.Value.(int64)]
			}

			if stackItemA.Type == STACK_TEMP {
				valueB = stackItemA.Value.(*Value)
			}
			if stackItemB.Type == STACK_TEMP {
				valueB = stackItemB.Value.(*Value)
			}

			if valueA.Type == parser.VInt && valueB.Type == parser.VInt {
				result := valueA.Value.(int64) * valueB.Value.(int64)
				vm.ESP--
				vm.Stack[vm.ESP] = &StackItem{
					Type: STACK_TEMP,
					Value: &Value{
						Type:  parser.VInt,
						Value: result,
					},
				}
			}

		// 赋值操作符
		case runtime.ASSIGN:
			stackItemA := vm.Stack[vm.ESP-1]
			stackItemB := vm.Stack[vm.ESP]

			// 如果被赋值的类型是VAR_POINTER
			if stackItemA.Type == VAR_POINTER {
				// 常量索引
				if stackItemB.Type == CONST_IDX {
					valueB = vm.Constants[stackItemB.Value.(int64)]
				}

				// 栈临时变量
				if stackItemB.Type == STACK_TEMP {
					valueB = stackItemB.Value.(*Value)
				}

				// 将新值赋值给变量
				value := Value{
					Type:  parser.VInt,
					Value: valueB.Value.(int64),
				}

				*(*int64)(unsafe.Pointer(uintptr(stackItemA.Value.(int64)))) =
					int64(uintptr(unsafe.Pointer(&value)))
			}

		}
	}

	return nil
}
