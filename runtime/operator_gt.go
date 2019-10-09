package runtime

import (
	"bvm/parser"
	"bvm/utils"
)

func Gt(vm *VM) error {
	vm.ESP--
	stackItemA := vm.Stack[vm.ESP]
	stackItemB := vm.Stack[vm.ESP+1]

	valueA := GetValueFromStack(vm, stackItemA)
	valueB := GetValueFromStack(vm, stackItemB)

	if err := CheckValue(valueA); err != nil {
		return err
	}
	if err := CheckValue(valueB); err != nil {
		return err
	}

	var stackItem *StackItem

	if valueA.Value.(int64) > valueB.Value.(int64) {
		stackItem = &StackItem{
			Type: STACK_TEMP,
			Value: &Value{
				Type:  parser.VBool,
				Value: true,
			},
		}
	} else {
		stackItem = &StackItem{
			Type: STACK_TEMP,
			Value: &Value{
				Type:  parser.VBool,
				Value: false,
			},
		}
	}
	vm.Stack[vm.ESP] = stackItem

	utils.DebugPrintf("VM> GT\n")
	return nil
}
