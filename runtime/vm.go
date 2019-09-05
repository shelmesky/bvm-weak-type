package runtime

import (
	"bvm/parser"
	"fmt"
	"reflect"
	"unsafe"
)

type BCode uint16

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

func Run(byteCodeStream []uint16, constantTable []Value) error {
	var (
		valueA *Value
		valueB *Value
	)
	calls := make([]int64, 1000)
	var coff int64

	vm := VM{
		Constants: make([]*Value, 0),
		Vars:      make([]*Value, 0),
		Stack:     make([]*StackItem, 256),
		ESP:       0,
		EBP:       0,
	}

	code := byteCodeStream
	length := int64(len(code))
	var i int64

	for i < length {
		switch code[i] {
		case INITVARS:
			variable := Value{
				Type: VAR_IDX,
			}
			vm.Vars = append(vm.Vars, &variable)
			fmt.Printf("VM> INITVARS\n")

		case PUSH:
			i++
			vm.ESP++
			// 根据PUSH的操作数获取在常量表中的常量值
			cnst := constantTable[code[i]]
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
			fmt.Printf("VM> PUSH %d\n", code[i])

		case SETVAR:
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
			fmt.Printf("VM> SETVAR %d\n", code[i])

		case GETVAR:
			i++
			vm.ESP++
			varIndex := int64(code[i])
			stackItem := StackItem{
				Type:  VAR_IDX,
				Value: varIndex,
			}
			vm.Stack[vm.ESP] = &stackItem
			fmt.Printf("VM> GETVAR %d\n", code[i])

		case ADD:
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

			if stackItemA.Type == VAR_IDX {
				valueA = vm.Vars[stackItemA.Value.(int64)]
			}

			if stackItemB.Type == VAR_IDX {
				valueB = vm.Vars[stackItemB.Value.(int64)]
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
			fmt.Printf("VM> ADD\n")

		case MUL:
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
			fmt.Printf("VM> MUL")

		// 赋值操作符
		case ASSIGN:
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

				if stackItemB.Type == VAR_IDX {
					valueB = vm.Vars[stackItemB.Value.(int64)]
				}

				// 将新值赋值给变量
				value := Value{
					Type:  parser.VInt,
					Value: valueB.Value.(int64),
				}

				*(*int64)(unsafe.Pointer(uintptr(stackItemA.Value.(int64)))) =
					int64(uintptr(unsafe.Pointer(&value)))
			}
			fmt.Printf("VM> ASSIGN\n")

		case JMP:
			dest := int64(int16(code[i+1]))
			i += dest
			fmt.Printf("VM> JMP %d\n", i)

		case CALLFUNC:
			calls[coff] = int64(i) + 2        // 在coff处将当前指令后的2条指令指针保存
			coff += 1                         // coff变量+1
			offset := int64(int16(code[i+1])) // 跳转到目标函数
			i += offset
			fmt.Printf("VM> CALLFUNC  dest: %d  origin: %d\n", i, calls[coff-1])
			continue

		case RETFUNC:
			coff -= 1
			i = calls[coff] // 从函数返回，恢复指令指针
			fmt.Printf("VM> RETRUNC %d\n", i)
			continue

		case GETPARAMS:
			i++
			paramCount := code[i] // 参数数量
			var paramValue *Value // 实参
			for j := 0; j <= int(paramCount)-1; j++ {
				paramItem := vm.Stack[vm.ESP] // 实参在stack中的元素
				if paramItem.Type == CONST_IDX {
					paramValue = vm.Constants[paramItem.Value.(int64)] // 获取真正的实参
				}
				i++
				argumentVarIdx := int64(code[i]) // 形参在VM变量内存中的索引
				if paramValue != nil {
					vm.Vars[argumentVarIdx] = paramValue // 将实参赋值给形参
				}
				// 运行GETPARAMS指令前, 实参已经由调用者PUSH到栈上
				// 在将这些实参复制给形参后, 栈上的实参已经不再使用
				// 所以栈要回退到PUSH实参之前
				// 回退完毕后, 如果是变量赋值, 则栈顶保存的是SETVAR指令放在栈顶的变量
				vm.ESP--
			}
			fmt.Printf("VM> GETPARAMS %d\n", paramCount)

		case CALLEMBED:
			i++
			funcIdx := code[i]
			embedFunc := Stdlib[funcIdx]

			// 从栈中获取实参
			var funcParams []*Value
			for j := 0; j <= embedFunc.ParamNum-1; j++ {
				stackItem := vm.Stack[vm.ESP-j]
				if stackItem.Type == VAR_IDX {
					param := vm.Vars[stackItem.Value.(int64)]
					funcParams = append(funcParams, param)
				}
			}

			// 将实参赋值给形参, 并封装为reflect.Value类型以方便调用reflect.ValueOf().Call()
			funcArguments := make([]reflect.Value, embedFunc.ParamNum+1)
			funcArguments[0] = reflect.ValueOf(vm)
			for idx := range funcParams {
				funcArguments[idx+1] = reflect.ValueOf(funcParams[idx])
			}

			// 调用函数并接收返回值
			var result []reflect.Value
			result = reflect.ValueOf(embedFunc.Func).Call(funcArguments)

			// 如果函数有返回值, 但是未返回任何值
			if embedFunc.HasReturn && len(result) == 0 {
				return fmt.Errorf("Embed function has return value, but return nothing.\n")
			}

			// 如果被调用函数定义有返回值, 且返回了值
			if embedFunc.HasReturn && len(result) > 0 {
				ret := result[0].Interface()
				retValue := ret.(int)
				stackItem := &StackItem{
					Type:  VAR_IDX,
					Value: retValue,
				}
				vm.ESP++
				vm.Stack[vm.ESP] = stackItem

			} else {
				// 如果函数无返回值或未返回, 就向栈上PUSH一个Void值
				stackItem := &StackItem{
					Type: VAR_IDX,
					Value: Value{
						Type:  parser.VVoid,
						Value: nil,
					},
				}
				vm.ESP++
				vm.Stack[vm.ESP] = stackItem
			}

		default:
			return fmt.Errorf("VM> unknown command %d\n", code[i])

		}
		i++
	}

	return nil
}
