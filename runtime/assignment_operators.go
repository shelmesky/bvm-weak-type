package runtime

import (
	"bvm/parser"
	"bvm/utils"
)

/*
实现ASSIGN指令：将右值的类型和值，直接复制给左值.
*/
func Assign(vm *VM) error {
	left, right, err := getValueAB(vm)
	if err != nil {
		return err
	}

	left.Type = right.Type
	left.Value = right.Value

	utils.DebugPrintf("VM> ASSIGN\n")

	return nil
}

/*
实现+=指令：stackItemVar是栈上保存的变量的Var Index
valueA是将要被赋值的变量的复制品
a+=1最终在VM中被转换为a=a+1
*/
func AddAssign(vm *VM) error {
	stackItemVar := vm.Stack[vm.ESP-2]

	if stackItemVar.Type == VAR_IDX {

		valueA, valueB, err := getValueAB(vm)
		if err != nil {
			return err
		}

		// 检查左边的变量是否未初始化
		// 否则将左值初始化为右值同样的值和类型
		checkEmptyValue(valueA, valueB)

		if valueA.Type == parser.VInt && valueB.Type == parser.VInt {
			result := valueA.Value.(int64) + valueB.Value.(int64)
			value := &Value{
				Type:  parser.VInt,
				Value: result,
			}

			// 执行完毕ADD_ASSIGN后: 变量 = 变量 + stackItem, 这3个栈上的元素不再需要
			vm.ESP -= 3

			leftValue := GetValueFromStack(vm, stackItemVar)
			leftValue.Value = value.Value
		}

		if valueA.Type == parser.VStr && valueB.Type == parser.VStr {
			result := valueA.Value.(string) + valueB.Value.(string)
			value := &Value{
				Type:  parser.VInt,
				Value: result,
			}

			// 执行完毕ADD_ASSIGN后: 变量 = 变量 + stackItem, 这3个栈上的元素不再需要
			vm.ESP -= 3

			leftValue := GetValueFromStack(vm, stackItemVar)
			leftValue.Value = value.Value
		}
	}

	utils.DebugPrintf("VM> ADD_ASSIGN\n")

	return nil
}
