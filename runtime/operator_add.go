package runtime

import "bvm/parser"
import "fmt"

func Add(vm *VM) error {
	// 从stack获取栈顶的2个元素
	stackItemA := vm.Stack[vm.ESP-1]
	stackItemB := vm.Stack[vm.ESP]

	valueA := GetValueFromStack(vm, stackItemA)
	valueB := GetValueFromStack(vm, stackItemB)

	if err := CheckValue(valueA); err != nil {
		return err
	}

	if err := CheckValue(valueB); err != nil {
		return err
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

	return nil
}
