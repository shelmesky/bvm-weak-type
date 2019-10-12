package runtime

import (
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
