package compiler

import (
	"bvm/parser"
	"golang.org/x/exp/errors/fmt"
)

const (
	NOP = iota
	PUSH
	INITVARS
	GETVAR
	SETVAR
	ADD
	SUB
	MUL
	DIV
	ASSIGN
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

		this.AppendCode(INITVARS)

	}

	return nil
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
		err      error
		variable Variable
		ok       bool
	)

	switch node.Type {
	// 智能合约
	case parser.TContract:
		contract := node.Value.(*parser.NContract)
		if err = nodeToCode(cmpl, contract.Block); err != nil {
			return err
		}

	// 代码块
	case parser.TBlock:
		block := node.Value.(*parser.NBlock)

		for _, child := range block.Statements {
			if err = nodeToCode(cmpl, child); err != nil {
				return err
			}
		}

	// 赋值语句和二元表达式
	case parser.TBinary:
		nBinary := node.Value.(parser.NBinary)

		// 递归处理左子树
		if err = nodeToCode(cmpl, nBinary.Left); err != nil {
			return err
		}

		// 如果左子树的类型是var xxx这样的语句
		// 那么左子树的类型是TVars类型, 上面已经处理了TVars类型
		// 这里需要再次处理变量名, 生命SETVAR指令, 将刚才INITVARS指令生成的对象赋值给变量
		if nBinary.Left.Type == parser.TVars {
			nBinary.Left = &parser.Node{
				Type: parser.TSetVar,
				Value: &parser.NVarValue{
					Name: nBinary.Left.Value.(*parser.NVars).Vars[0].Name,
				},
			}

			if err = nodeToCode(cmpl, nBinary.Left); err != nil {
				return err
			}
		}

		// 递归处理右子树
		if err = nodeToCode(cmpl, nBinary.Right); err != nil {
			return nil
		}

		// 处理操作符
		switch nBinary.Oper {
		case parser.ADD:
			cmpl.AppendCode(ADD)
		case parser.SUB:
			cmpl.AppendCode(SUB)
		case parser.MUL:
			cmpl.AppendCode(MUL)
		case parser.DIV:
			cmpl.AppendCode(DIV)
		case parser.ASSIGN:
			cmpl.AppendCode(ASSIGN)
		}

		// 变量声明: var a

	// 变量定义
	case parser.TVars:

		if err = cmpl.InitVars(node, node.Value.(parser.NVars).Vars); err != nil {

		}

	// var a = 111 或 a = 111 中的变量a
	case parser.TSetVar:
		name := node.Value.(*parser.NVarValue).Name
		if variable, ok = cmpl.SymbolTable[name]; !ok {
			return fmt.Errorf("unknow variable: %s\n", name)
		}

		cmpl.AppendCode(SETVAR)
		cmpl.AppendCode(BCode(variable.Index))

	// 表达式中出现的变量
	case parser.TGetVar:
		name := node.Value.(*parser.NVarValue).Name
		if variable, ok = cmpl.SymbolTable[name]; !ok {
			return fmt.Errorf("unknow variable: %s\n", name)
		}

		cmpl.AppendCode(GETVAR, BCode(variable.Index))
	}

	return nil
}
