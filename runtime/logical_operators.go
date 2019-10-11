package runtime

import (
	"bvm/parser"
	"bvm/utils"
)

func Gt(vm *VM) error {
	valueA, valueB, err := getValueAB(vm)
	if err != nil {
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

func Lt(vm *VM) error {
	valueA, valueB, err := getValueAB(vm)
	if err != nil {
		return err
	}

	var stackItem *StackItem

	if valueA.Value.(int64) < valueB.Value.(int64) {
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

	utils.DebugPrintf("VM> LT\n")

	return nil
}
