package runtime

import (
	"bvm/parser"
	"bvm/utils"
	"unsafe"
)

func Assign(vm *VM) error {
	stackItemVar := vm.Stack[vm.ESP-1]
	stackItemExprResult := vm.Stack[vm.ESP]

	// 如果被赋值的类型是VAR_POINTER
	if stackItemVar.Type == VAR_POINTER {
		valueA := GetValueFromStack(vm, stackItemExprResult)

		if err := CheckValue(valueA); err != nil {
			return err
		}

		// 将新值赋值给变量
		value := TypeLoader(valueA)

		*(*int64)(unsafe.Pointer(uintptr(stackItemVar.Value.(int64)))) =
			int64(uintptr(unsafe.Pointer(&value)))
	}
	utils.DebugPrintf("VM> ASSIGN\n")

	return nil
}

/*
实现+=指令：stackItemVar是栈上保存的变量的指针
valueA是将要被赋值的变量的复制品
a+=1最终在VM中被转换为a=a+1
*/
func AddAssign(vm *VM) error {
	stackItemVar := vm.Stack[vm.ESP-2]

	// 如果被赋值的类型是VAR_POINTER
	if stackItemVar.Type == VAR_POINTER {
		valueA, valueB, err := getValueAB(vm)
		if err != nil {
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
