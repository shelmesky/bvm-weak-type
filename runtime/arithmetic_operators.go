package runtime

import (
	"bvm/parser"
	"bvm/utils"
)

func Add(vm *VM) error {
	valueA, valueB, err := getValueAB(vm)
	if err != nil {
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
	utils.DebugPrintf("VM> ADD\n")

	return nil
}

func Mul(vm *VM) error {
	valueA, valueB, err := getValueAB(vm)
	if err != nil {
		return err
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
	utils.DebugPrintf("VM> MUL\n")

	return nil
}
