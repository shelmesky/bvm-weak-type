package runtime

import (
	"bvm/parser"
	"bvm/utils"
	"fmt"
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
	} else if valueA.Type == parser.VFloat && valueB.Type == parser.VFloat {
		result := valueA.Value.(float64) + valueB.Value.(float64)
		vm.ESP--
		vm.Stack[vm.ESP] = &StackItem{
			Type: STACK_TEMP,
			Value: &Value{
				Type:  parser.VFloat,
				Value: result,
			},
		}
	} else if valueA.Type == parser.VStr && valueB.Type == parser.VStr {
		result := valueA.Value.(string) + valueB.Value.(string)
		vm.ESP--
		vm.Stack[vm.ESP] = &StackItem{
			Type: STACK_TEMP,
			Value: &Value{
				Type:  parser.VStr,
				Value: result,
			},
		}
	} else {
		return fmt.Errorf("ADD operator only supoort int,float,string\n")
	}
	utils.DebugPrintf("VM> ADD\n")

	return nil
}

func Sub(vm *VM) error {
	valueA, valueB, err := getValueAB(vm)
	if err != nil {
		return err
	}

	if valueA.Type == parser.VInt && valueB.Type == parser.VInt {
		result := valueA.Value.(int64) - valueB.Value.(int64)
		vm.ESP--
		vm.Stack[vm.ESP] = &StackItem{
			Type: STACK_TEMP,
			Value: &Value{
				Type:  parser.VInt,
				Value: result,
			},
		}
	} else if valueA.Type == parser.VFloat && valueB.Type == parser.VFloat {
		result := valueA.Value.(float64) - valueB.Value.(float64)
		vm.ESP--
		vm.Stack[vm.ESP] = &StackItem{
			Type: STACK_TEMP,
			Value: &Value{
				Type:  parser.VFloat,
				Value: result,
			},
		}
	} else {
		return fmt.Errorf("SUB operator only supoort int and float\n")
	}

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
	} else if valueA.Type == parser.VFloat && valueB.Type == parser.VFloat {
		result := valueA.Value.(float64) * valueB.Value.(float64)
		vm.ESP--
		vm.Stack[vm.ESP] = &StackItem{
			Type: STACK_TEMP,
			Value: &Value{
				Type:  parser.VFloat,
				Value: result,
			},
		}
	} else {
		return fmt.Errorf("operator DIV only supoort int and float\n")
	}
	utils.DebugPrintf("VM> MUL\n")

	return nil
}

func Div(vm *VM) error {
	valueA, valueB, err := getValueAB(vm)
	if err != nil {
		return err
	}

	if valueA.Type == parser.VInt && valueB.Type == parser.VInt {
		result := valueA.Value.(int64) / valueB.Value.(int64)
		vm.ESP--
		vm.Stack[vm.ESP] = &StackItem{
			Type: STACK_TEMP,
			Value: &Value{
				Type:  parser.VInt,
				Value: result,
			},
		}
	} else if valueA.Type == parser.VFloat && valueB.Type == parser.VFloat {
		result := valueA.Value.(float64) / valueB.Value.(float64)
		vm.ESP--
		vm.Stack[vm.ESP] = &StackItem{
			Type: STACK_TEMP,
			Value: &Value{
				Type:  parser.VFloat,
				Value: result,
			},
		}
	} else {
		return fmt.Errorf("operator DIV only supoort int and float\n")
	}
	utils.DebugPrintf("VM> MUL\n")

	return nil
}

func Mod(vm *VM) error {
	valueA, valueB, err := getValueAB(vm)
	if err != nil {
		return err
	}

	if valueA.Type == parser.VInt && valueB.Type == parser.VInt {
		result := valueA.Value.(int64) % valueB.Value.(int64)
		vm.ESP--
		vm.Stack[vm.ESP] = &StackItem{
			Type: STACK_TEMP,
			Value: &Value{
				Type:  parser.VInt,
				Value: result,
			},
		}
	} else {
		return fmt.Errorf("operator MOD only supoort int\n")
	}
	utils.DebugPrintf("VM> MUL\n")

	return nil
}
