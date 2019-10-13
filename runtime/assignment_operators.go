package runtime

import (
	"bvm/parser"
	"bvm/utils"
	"unsafe"
)

func Assign(vm *VM) error {
	stackItemA := vm.Stack[vm.ESP-1]
	stackItemB := vm.Stack[vm.ESP]

	// 如果被赋值的类型是VAR_POINTER
	if stackItemA.Type == VAR_POINTER {
		valueB := GetValueFromStack(vm, stackItemB)

		if err := CheckValue(valueB); err != nil {
			return err
		}

		// 将新值赋值给变量
		value := TypeLoader(valueB)

		*(*int64)(unsafe.Pointer(uintptr(stackItemA.Value.(int64)))) =
			int64(uintptr(unsafe.Pointer(&value)))
	}
	utils.DebugPrintf("VM> ASSIGN\n")

	return nil
}

func AddAssign(vm *VM) error {
	stackItemVar := vm.Stack[vm.ESP-2]

	stackItemA := vm.Stack[vm.ESP-1]
	stackItemB := vm.Stack[vm.ESP]

	// 如果被赋值的类型是VAR_POINTER
	if stackItemVar.Type == VAR_POINTER {
		valueA := GetValueFromStack(vm, stackItemA)

		if err := CheckValue(valueA); err != nil {
			return err
		}

		valueB := GetValueFromStack(vm, stackItemB)

		if err := CheckValue(valueB); err != nil {
			return err
		}

		result := valueA.Value.(int64) + valueB.Value.(int64)
		value := Value{
			Type:  parser.VInt,
			Value: result,
		}

		*(*int64)(unsafe.Pointer(uintptr(stackItemVar.Value.(int64)))) =
			int64(uintptr(unsafe.Pointer(&value)))
	}
	utils.DebugPrintf("VM> ADD_ASSIGN\n")

	return nil
}
