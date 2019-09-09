package main

import (
	"bvm/compiler"
	"bvm/parser"
	"bvm/runtime"
	"io/ioutil"
	"log"
	"os"
	"strings"
)
import "fmt"

func main() {
	inputFilename := os.Args[1]

	if len(inputFilename) == 0 {
		fmt.Println("need filename")
		os.Exit(1)
	}

	content, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		log.Fatal("ReadFile failed:", err)
	}

	list := strings.Split(string(content), "\n")
	source := make([]string, 0, 32)

	for _, line := range list {
		source = append(source, line)
	}

	contractBody := strings.Join(source, "\r\n")

	root, err := parser.Parser(contractBody)
	if err != nil {
		panic(err.Error())
	}

	cmplResult, err := compiler.Compile(root)
	if cmplResult != nil {
		bytesCodeStream := getByteCodes(cmplResult)
		constantTable := getConstantTable(cmplResult)
		funcList := getFuncList(cmplResult)
		varTableSize := len(cmplResult.VarTable)

		if err != nil {
			panic(err.Error())
		}

		err = runtime.Run(bytesCodeStream, funcList, constantTable, varTableSize)

		if err != nil {
			panic(err.Error())
		}
	}
}

func getByteCodes(cmplResult *compiler.CompileEnv) []uint16 {
	var codeStream []uint16

	for idx := range cmplResult.Code {
		codeStream = append(codeStream, uint16(cmplResult.Code[idx]))
	}

	return codeStream
}

func getConstantTable(cmplResult *compiler.CompileEnv) []runtime.Value {
	var constantTable []runtime.Value

	for idx := range cmplResult.ConstantsTable {
		var cnst runtime.Value
		cnst.Type = cmplResult.ConstantsTable[idx].Type
		cnst.Value = cmplResult.ConstantsTable[idx].Value
		constantTable = append(constantTable, cnst)
	}

	return constantTable
}

func getFuncList(cmplResult *compiler.CompileEnv) []runtime.FuncInfo {
	var FuncList []runtime.FuncInfo

	for idx := range cmplResult.FuncList {
		var fInfo runtime.FuncInfo
		fInfo.Offset = cmplResult.FuncList[idx].Offset
		fInfo.HasReturn = cmplResult.FuncList[idx].HasReturn
		fInfo.Name = cmplResult.FuncList[idx].Name
		fInfo.Index = cmplResult.FuncList[idx].Index
		fInfo.ParamsNum = cmplResult.FuncList[idx].ParamsNum
		fInfo.LocalVarNum = cmplResult.FuncList[idx].LocalVarNum
		FuncList = append(FuncList, fInfo)
	}

	return FuncList
}
