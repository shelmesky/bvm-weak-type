package runtime

import (
	"bvm/parser"
	"bvm/utils"
	"fmt"
	"reflect"
)

const (
	StackSize = 64
)

type BCode uint16

// 栈元素的类型
const (
	CONST_IDX   = iota // 常量
	VAR_IDX            // 变量内存
	VAR_POINTER        // 指针的整数形式
	STACK_TEMP         // 栈中的临时元素
	FUNC_IDX
)

// 值类型, 可保存类型任何值
type Value struct {
	Type  int         // 可以是VVoid/VInt/VStr/VFloat等类型
	Value interface{} // 实际的值
}

// 栈元素
type StackItem struct {
	Type  int         // 可以是CONST_IDX/VAR_IDX/VAR_POINTER等类型
	Value interface{} // 实际的值
}

// 函数信息
type FuncInfo struct {
	Index       int    // 函数在列表中的索引
	Name        string // 函数名
	Offset      int64  // 在指令流中的开始位置
	ParamsNum   int    // 参数数量
	LocalVarNum int    // 局部变量数量
	HasReturn   bool   // 是否有返回值
}

// 调用栈帧保存函数调用的信息
// 在函数返回时恢复
type CallFrame struct {
	ReturnAddress    int64
	IdxOfCallingFunc int
	ESP              int
}

/*
vm.Stack中元素为StackItem,
StackItem的Value大部分情况下是Value类型,
但也有可能是简单类型, 例如整数或字符串等.
*/

type VM struct {
	Constants        []*Value     // 常量
	Vars             []*Value     // 变量
	Funcs            []*Value     // 函数
	Stack            []*StackItem // 栈
	ESP              int          // 栈指针
	EBP              int          // 栈基址指针
	IdxOfCallingFunc int          // 正在调用的函数
}

func Run(byteCodeStream []uint16, FuncList []FuncInfo, constantTable []Value, varTableSize int) error {
	calls := make([]CallFrame, 1000)
	var coff int64
	var loopStack int

	vm := &VM{
		Constants: make([]*Value, 0),
		Vars:      make([]*Value, 0),
		Stack:     make([]*StackItem, StackSize),
		ESP:       0,
		EBP:       0,
	}

	code := byteCodeStream
	length := int64(len(code))
	var i int64

	for j := 0; j < varTableSize; j++ {
		variable := Value{
			Type: VAR_IDX,
		}
		vm.Vars = append(vm.Vars, &variable)
	}

	for i < length {
		switch code[i] {
		/*
			case INITVARS:
				variable := Value{
					Type: VAR_IDX,
				}
				vm.Vars = append(vm.Vars, &variable)
				fmt.Printf("VM> INITVARS\n")
		*/

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
			utils.DebugPrintf("VM> PUSH %d\n", code[i])

		case SETVAR:
			i++
			vm.ESP++
			// 根据SETVAR的操作数获取变量的索引
			varIndex := int64(code[i])
			utils.DebugPrintf("VM> SETVAR %d\n", code[i])
			// 获取索引在vm.Vars中的指针
			//varPointer := int64(uintptr(unsafe.Pointer(&vm.Vars[varIndex])))
			// 保存指针到stack元素
			stackItem := StackItem{
				Type:  VAR_IDX,
				Value: varIndex,
			}
			vm.Stack[vm.ESP] = &stackItem

		case GETVAR:
			i++
			vm.ESP++
			varIndex := int64(code[i])
			stackItem := StackItem{
				Type:  VAR_IDX,
				Value: varIndex,
			}
			vm.Stack[vm.ESP] = &stackItem
			utils.DebugPrintf("VM> GETVAR %d\n", code[i])

		case ADD:
			if err := Add(vm); err != nil {
				return err
			}

		case SUB:
			if err := Sub(vm); err != nil {
				return err
			}

		case MUL:
			if err := Mul(vm); err != nil {
				return err
			}

		case DIV:
			if err := Div(vm); err != nil {
				return err
			}

		case MOD:
			if err := Mod(vm); err != nil {
				return err
			}

		case BIT_AND:
			utils.DebugPrintf("VM> BIT_AND\n")
		case BIT_OR:
			utils.DebugPrintf("VM> BIT_OR\n")
		case BIT_XOR:
			utils.DebugPrintf("VM> BIT_XOR\n")
		case LEFT_SHIFT:
			utils.DebugPrintf("VM> LEFT_SHIFT\n")
		case RIGHT_SHIFT:
			utils.DebugPrintf("VM> RIGHT_SHIFT\n")
		case POW:
			utils.DebugPrintf("VM> POW\n")

		case ADD_ASSIGN:
			if err := AddAssign(vm); err != nil {
				return err
			}
		case SUB_ASSIGN:
			utils.DebugPrintf("VM> SUB_ASSIGN\n")
		case MUL_ASSIGN:
			utils.DebugPrintf("VM> MUL_ASSIGN\n")
		case DIV_ASSIGN:
			utils.DebugPrintf("VM> DIV_ASSIGN\n")
		case MOD_ASSIGN:
			utils.DebugPrintf("VM> MOD_ASSIGN\n")
		case LEFT_SHIFT_ASSIGN:
			utils.DebugPrintf("VM> LEFT_SHIFT_ASSIGN\n")
		case RIGHT_SHIFT_ASSIGN:
			utils.DebugPrintf("VM> RIGHT_SHIFT_ASSIGN\n")
		case BIT_AND_ASSIGN:
			utils.DebugPrintf("VM> BIT_AND_ASSIGN\n")
		case BIT_XOR_ASSIGN:
			utils.DebugPrintf("VM> BIT_XOR_ASSIGN\n")
		case BIT_OR_ASSIGN:
			utils.DebugPrintf("VM> BIT_OR_ASSIGN\n")
		// 赋值操作符
		case ASSIGN:
			if err := Assign(vm); err != nil {
				return err
			}

		case LOOP:
			utils.DebugPrintf("VM> LOOP\n")
			// 循环开始前记录栈的大小
			loopStack = vm.ESP

		case JMP:
			dest := int64(int16(code[i+1]))
			i = dest
			// 跳转前检查
			if loopStack >= 0 {
				vm.ESP = loopStack
				loopStack = -1
			}
			utils.DebugPrintf("VM> JMP %d\n", i)
			continue

		case JZE:
			i++
			offset := code[i]
			// 这里栈顶-1是因为JZE指令获取栈顶的值后
			// 这个值不再被后续的指令需要, 需要弹出栈
			vm.ESP--

			topStackItem := vm.Stack[vm.ESP+1]
			value := GetValueFromStack(vm, topStackItem)
			if value.Type != parser.VBool {
				panic("JZE need VBool type")
			}

			// 如果逻辑或关系运算符的结果为true
			if value.Value.(bool) == false {
				i = int64(offset)
				utils.DebugPrintf("VM> JZE %d\n", offset)
				continue
			}

			utils.DebugPrintf("VM> JZE %d\n", offset)

		case CALLFUNC:
			var callFrame CallFrame
			callFrame.ReturnAddress = int64(i) + 2 // 将当前指令后的2条指令指针保存
			funcIndex := int(code[i+1])
			callFrame.IdxOfCallingFunc = funcIndex // 记录当前正在调用的函数索引
			callFrame.ESP = vm.ESP                 // 保存栈指针

			fInfo := FuncList[funcIndex]
			// 如果被调函数有参数, 就在保存调用栈时预先减去参数的数量
			// 因为GETPARAMS指令在复制实参到形参后, 会将实参出栈, 栈会缩小
			if fInfo.ParamsNum > 0 {
				callFrame.ESP -= fInfo.ParamsNum
			}
			calls[coff] = callFrame // 保存当前调用栈帧
			coff += 1               // coff变量+1

			offset := fInfo.Offset
			i = offset
			utils.DebugPrintf("VM> CALLFUNC  dest: %d  origin: %d\n", i, calls[coff-1])
			continue

		case RETFUNC:
			coff -= 1
			// 获取之前的调用栈帧
			callFrame := calls[coff]

			// 在跳转回调用函数之前
			// 将被调函数放在栈顶的表达式返回值复制到之前保存的栈顶+1的地方
			topStackItem := vm.Stack[vm.ESP]
			vm.Stack[callFrame.ESP+1] = topStackItem

			// 函数返回前, 将栈顶设置为调用之前的大小+1
			// +1是因为函数返中有最后一个表达式的返回值
			// 如果函数无返回值
			// 下面的代码也会在栈顶强行插入一个Void类型
			vm.ESP = callFrame.ESP + 1

			i++
			hasReturnExpr := code[i]
			// 如果从函数中没有返回表达式
			// 则在栈顶放置一个VVoid类型
			if hasReturnExpr == 0 {
				vm.Stack[vm.ESP] = &StackItem{
					Type: STACK_TEMP,
					Value: &Value{
						Type:  parser.VVoid,
						Value: nil,
					},
				}
			}
			i = callFrame.ReturnAddress // 从函数返回，恢复指令指针
			utils.DebugPrintf("VM> RETRUNC %d\n", i)
			continue

		case GETPARAMS:
			i++
			paramCount := code[i] // 参数数量
			var paramValue *Value // 实参
			for j := 0; j <= int(paramCount)-1; j++ {
				paramItem := vm.Stack[vm.ESP]                 // 实参在stack中的元素
				paramValue = GetValueFromStack(vm, paramItem) //  根据栈元素类型获取Value
				i++
				argumentVarIdx := int64(code[i]) // 形参在VM变量内存中的索引
				if paramValue != nil {
					vm.Vars[argumentVarIdx] = paramValue // 将实参赋值给形参
				}
				// 运行GETPARAMS指令前, 实参已经由调用者PUSH到栈上
				// 在使用这些实参设置形参后, 栈上的实参已经不再使用
				// 所以栈要回退到PUSH实参之前
				// 回退完毕后, 如果是变量赋值, 则栈顶保存的是SETVAR指令放在栈顶的变量
				vm.ESP--
			}
			utils.DebugPrintf("VM> GETPARAMS %d\n", paramCount)

		case CALLEMBED:
			i++
			funcIdx := code[i]
			embedFunc := Stdlib[funcIdx]
			utils.DebugPrintf("VM> CALLEMBED %s\n", embedFunc.Name)

			// 从栈中获取实参
			var funcParams []*Value
			for j := 0; j <= embedFunc.ParamNum-1; j++ {
				stackItem := vm.Stack[vm.ESP-j]
				paramValue := GetValueFromStack(vm, stackItem)
				funcParams = append(funcParams, paramValue)
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
			// 返回值已经保存在vm.Vars中, 返回的是索引
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
				// 如果函数无返回值或未返回,
				// 就向栈上PUSH一个临时的Void值
				stackItem := &StackItem{
					Type: STACK_TEMP,
					Value: &Value{
						Type:  parser.VVoid,
						Value: nil,
					},
				}
				vm.ESP++
				vm.Stack[vm.ESP] = stackItem
			}

		case AND:
			utils.DebugPrintf("VM> AND\n")
		case OR:
			utils.DebugPrintf("VM> OR\n")
		case EQ:
			utils.DebugPrintf("VM> EQ\n")
		case NOTEQ:
			utils.DebugPrintf("VM> NOTEQ\n")
		case NOT:
			utils.DebugPrintf("VM> NOT\n")

		case LT:
			if err := Lt(vm); err != nil {
				return err
			}

		case GT:
			if err := Gt(vm); err != nil {
				return err
			}

		case LTE:
			utils.DebugPrintf("VM> LTE\n")
		case GTE:
			utils.DebugPrintf("VM> GTE\n")

		case INITMAP:
			i++
			kvCount := int64(code[i])
			imap := make(map[string]*Value)
			for k := int64(0); k < kvCount; k++ {
				currentIdx := int64(vm.ESP) - 2*(kvCount-k) + 1

				// 在栈上获取key和value元素
				keyStackItem := vm.Stack[currentIdx]
				key := GetValueFromStack(vm, keyStackItem)
				valueStackItem := vm.Stack[currentIdx+1]
				value := GetValueFromStack(vm, valueStackItem)

				if err := CheckValue(key); err != nil {
					return err
				}
				if err := CheckValue(value); err != nil {
					return err
				}

				// 如果是key是字符串类型
				// TODO: 增加int类型的key
				if key.Type == parser.VStr {
					keyStr := key.Value.(string)
					// 保存key/value
					imap[keyStr] = value
				} else {
					return fmt.Errorf("map only support string key")
				}
			}

			// 缩小栈
			vm.ESP -= int(kvCount) * 2

			imapValue := &Value{
				Type:  parser.VMap,
				Value: imap,
			}
			vm.Vars = append(vm.Vars, imapValue)

			// 将map作为变量，保存在vm.Vars当中
			stackItem := &StackItem{
				Type:  VAR_IDX,
				Value: int64(len(vm.Vars) - 1),
			}
			vm.ESP++
			vm.Stack[vm.ESP] = stackItem

			utils.DebugPrintf("VM> INITMAP %d\n", kvCount)

		case GETINDEX:
			utils.DebugPrintf("VM> GETINDEX\n")
			object, key, err := getValueAB(vm)
			if err != nil {
				return err
			}

			vm.ESP -= 2

			if object.Type == parser.VMap {
				// 如果对象是map
				keyType := key.Type
				if keyType == parser.VStr {
					// 获取map对象
					imap := object.Value.(map[string]*Value)
					// 根据key取value
					value := imap[key.Value.(string)]

					// 将value作为临时变量保存在栈顶
					stackItem := &StackItem{
						Type:  STACK_TEMP,
						Value: value,
					}
					vm.ESP++
					vm.Stack[vm.ESP] = stackItem
				}
				// 如果是list类型
			} else if object.Type == parser.VArr {

				// 如果既不是map也不是list则提示不支持下标操作
			} else {
				return fmt.Errorf("variable not support index operation\n")
			}

		default:
			return fmt.Errorf("VM> unknown command %d\n", code[i])

		}
		i++
	}

	return nil
}

// 根据栈元素的类型, 在常量表/变量表/栈元素本身中获取Value类型
func GetValueFromStack(vm *VM, stackItem *StackItem) *Value {
	var value *Value

	if stackItem.Type == CONST_IDX {
		idx := stackItem.Value.(int64)
		value = vm.Constants[idx]
	}

	if stackItem.Type == VAR_IDX {
		idx := stackItem.Value.(int64)
		value = vm.Vars[idx]
	}

	if stackItem.Type == STACK_TEMP {
		value = stackItem.Value.(*Value)
	}

	return value
}

// 检查value是否为空值
func CheckValue(value *Value) error {
	if value.Type == parser.VVoid {
		return fmt.Errorf("Void type is not allowed\n")
	}

	return nil
}

// 获取栈顶的2个元素，并获取它们的值
// 检查不是空值后返回
func getValueAB(vm *VM) (*Value, *Value, error) {
	var (
		valueA *Value
		valueB *Value
		err    error
	)

	// 从stack获取栈顶的2个元素
	stackItemA := vm.Stack[vm.ESP-1]
	stackItemB := vm.Stack[vm.ESP]

	valueA = GetValueFromStack(vm, stackItemA)
	valueB = GetValueFromStack(vm, stackItemB)

	if err := CheckValue(valueA); err != nil {
		return valueA, valueB, err
	}

	if err := CheckValue(valueB); err != nil {
		return valueA, valueB, err
	}

	return valueA, valueB, err
}

/*
	如果变量声明时没有指定初始值, 则ValueA的值为nil.
	应该根据valueB的类型, 给valueA赋予初始值.
*/
func checkEmptyValue(valueA, valueB *Value) {

	if valueB.Type == parser.VInt && valueA.Value == nil {
		valueA.Type = parser.VInt
		valueA.Value = int64(0)
	}
	if valueB.Type == parser.VStr && valueA.Value == nil {
		valueA.Type = parser.VStr
		valueA.Value = ""
	}
	if valueB.Type == parser.VFloat && valueA.Value == nil {
		valueA.Type = parser.VFloat
		valueA.Value = float64(0)
	}
	if valueB.Type == parser.VBool && valueA.Value == nil {
		valueA.Type = parser.VBool
		valueA.Value = false
	}
}
