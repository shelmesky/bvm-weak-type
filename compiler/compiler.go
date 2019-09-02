package compiler

import (
	"bvm/parser"
	"golang.org/x/exp/errors/fmt"
)

const (
	NOP = iota
	LOADCONST
	INITVARS
	ADD
	SUB
	MUL
	DIV
)

type BCode uint16

// 函数信息
type FuncInfo struct {
	Offset       uint64
	Name         string
	ParamsLength int
}

// 常量
type Const struct {
	Type  int
	Value interface{}
}

// 变量
type Variable struct {
	Index    uint16
	Name     string
	IsGlobal bool
}

// 编译环境
type CompileEnv struct {
	SymbolTable map[string]Variable
	FuncTable   map[string]FuncInfo
	Constants   []Const
	Code        []BCode
}

func (this *CompileEnv) AppendCode(codes ...BCode) {
	for _, code := range codes {
		this.Code = append(this.Code, code)
	}
}

func (this *CompileEnv) InitVars(node *parser.Node, vars []parser.NVar) error {
	if len(vars) == 0 {
		return fmt.Errorf("empty vars")
	}

	for _, v := range vars {
		if _, ok := this.SymbolTable[v.Name]; ok {
			return fmt.Errorf("variable %s already exists\n", v.Name)
		}

		idx := uint16(len(this.SymbolTable))
		symbol := Variable{
			Index:    idx,
			Name:     v.Name,
			IsGlobal: false,
		}

		this.SymbolTable[v.Name] = symbol

		this.AppendCode()

	}
}

func Compile(root *parser.Node) error {
	cmpl := CompileEnv{
		SymbolTable: make(map[string]Variable),
		FuncTable:   make(map[string]FuncInfo, 32),
		Constants:   make([]Const, 32),
		Code:        make([]BCode, 32),
	}

	err := nodeToCode(&cmpl, root)

	if err != nil {
		return fmt.Errorf("compile failed:", err)
	}

	return nil
}

func nodeToCode(cmpl *CompileEnv, node *parser.Node) error {
	var (
		err error
	)

	switch node.Type {
	case parser.TContract:
		contract := node.Value.(*parser.NContract)
		if err = nodeToCode(cmpl, contract.Block); err != nil {
			return err
		}

	case parser.TBlock:
		block := node.Value.(*parser.NBlock)

		for _, child := range block.Statements {
			if err = nodeToCode(cmpl, child); err != nil {
				return err
			}
		}

		// 变量声明: var a
	case parser.TVars:

		if err = cmpl.InitVars(); err != nil {

		}

	}

	return nil
}
