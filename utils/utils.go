package utils

import (
	"fmt"
)

const (
	Debug = false
)

func DebugPrintf(format string, a ...interface{}) {
	if Debug == true {
		fmt.Printf(format, a...)
	}
}
