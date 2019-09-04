package compiler

import (
	"bvm/parser"
	"bvm/runtime"
	"fmt"
)

type BCode uint16

// 函数信息
type FuncInfo struct {
	Name        string // 函数名
	Offset      int64  // 在指令流中的开始位置
	ParamsNum   int    // 参数数量
	LocalVarNum int    // 局部变量数量
	HasReturn   bool   // 是否有返回值
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
	VarTable       map[string]Variable // 变量表
	FuncTable      map[string]FuncInfo // 函数表
	ConstantsTable []Const             // 常量表
	Code           []BCode             // 字节码
	InFunc         bool                // 是否正在编译函数
}

func (this *CompileEnv) AppendCode(codes ...BCode) {

	if len(codes) > 0 {
		switch codes[0] {
		case runtime.PUSH:
			fmt.Printf("PUSH index:[%d]\n", codes[1])
		case runtime.INITVARS:
			fmt.Println("INITVARS")
		case runtime.GETVAR:
			fmt.Printf("GETVAR index:[%d]\n", codes[1])
		case runtime.SETVAR:
			fmt.Printf("SETVAR index:[%d]\n", codes[1])
		case runtime.ADD:
			fmt.Println("ADD")
		case runtime.SUB:
			fmt.Println("SUB")
		case runtime.MUL:
			fmt.Println("MUL")
		case runtime.DIV:
			fmt.Println("DIV")
		case runtime.ASSIGN:
			fmt.Println("ASSIGN")
		case runtime.JMP:
			fmt.Printf("JMP [%d]\n", codes[1])
		case runtime.RETFUNC:
			fmt.Println("RETFUNC")
		case runtime.RETURN:
			fmt.Println("RETURN")
		}
	}

	for _, code := range codes {
		this.Code = append(this.Code, code)
	}
}

func (this *CompileEnv) InitVars(node *parser.Node, vars []parser.NVar) error {
	if len(vars) == 0 {
		return fmt.Errorf("empty vars")
	}

	for _, v := range vars {
		if _, ok := this.VarTable[v.Name]; ok {
			return fmt.Errorf("variable %s already exists\n", v.Name)
		}

		idx := uint16(len(this.VarTable))
		symbol := Variable{
			Index:    idx,
			Name:     v.Name,
			IsGlobal: false,
		}

		this.VarTable[v.Name] = symbol

		this.AppendCode(runtime.INITVARS)

	}

	return nil
}

func Compile(root *parser.Node) (*CompileEnv, error) {
	cmpl := CompileEnv{
		VarTable:       make(map[string]Variable),
		FuncTable:      make(map[string]FuncInfo, 0),
		ConstantsTable: make([]Const, 0),
		Code:           make([]BCode, 0),
	}

	err := nodeToCode(&cmpl, root)

	if err != nil {
		return &cmpl, fmt.Errorf("compile failed:", err)
	}

	return &cmpl, nil
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
		nBinary := node.Value.(*parser.NBinary)

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
			cmpl.AppendCode(runtime.ADD)
		case parser.SUB:
			cmpl.AppendCode(runtime.SUB)
		case parser.MUL:
			cmpl.AppendCode(runtime.MUL)
		case parser.DIV:
			cmpl.AppendCode(runtime.DIV)
		case parser.ASSIGN:
			cmpl.AppendCode(runtime.ASSIGN)
		}

		// 变量声明: var a

	// 变量定义
	case parser.TVars:

		if err = cmpl.InitVars(node, node.Value.(*parser.NVars).Vars); err != nil {

		}

	// var a = 111 或 a = 111 中的变量a
	case parser.TSetVar:
		name := node.Value.(*parser.NVarValue).Name
		if variable, ok = cmpl.VarTable[name]; !ok {
			return fmt.Errorf("unknow variable: %s\n", name)
		}

		cmpl.AppendCode(runtime.SETVAR, BCode(variable.Index))

	// 表达式中出现的变量
	case parser.TGetVar:
		name := node.Value.(*parser.NVarValue).Name
		if variable, ok = cmpl.VarTable[name]; !ok {
			return fmt.Errorf("unknow variable: %s\n", name)
		}

		cmpl.AppendCode(runtime.GETVAR, BCode(variable.Index))

	// 字面值
	case parser.TValue:
		switch node.Value.(type) {
		case int64:
			cnst := Const{
				Type:  parser.VInt,
				Value: node.Value,
			}

			cmpl.ConstantsTable = append(cmpl.ConstantsTable, cnst)
			cmpl.AppendCode(runtime.PUSH, BCode(len(cmpl.ConstantsTable)-1))

		case bool:
			cnst := Const{
				Type:  parser.VBool,
				Value: node.Value,
			}

			cmpl.ConstantsTable = append(cmpl.ConstantsTable, cnst)
			cmpl.AppendCode(runtime.PUSH, BCode(len(cmpl.ConstantsTable)-1))

		case string:
			cnst := Const{
				Type:  parser.VStr,
				Value: node.Value,
			}

			cmpl.ConstantsTable = append(cmpl.ConstantsTable, cnst)
			cmpl.AppendCode(runtime.PUSH, BCode(len(cmpl.ConstantsTable)-1))

		case float64:
			cnst := Const{
				Type:  parser.VFloat,
				Value: node.Value,
			}

			cmpl.ConstantsTable = append(cmpl.ConstantsTable, cnst)
			cmpl.AppendCode(runtime.PUSH, BCode(len(cmpl.ConstantsTable)-1))
		}

	// return语句
	case parser.TReturn:
		expr := node.Value.(*parser.NReturn).Expr

		// 如果return语句有表达式
		if expr != nil {
			if err = nodeToCode(cmpl, expr); err != nil {
				return err
			}
		}

		if cmpl.InFunc {
			cmpl.AppendCode(runtime.RETFUNC)
		} else {
			cmpl.AppendCode(runtime.RETURN)
		}

	case parser.TFunc:
		nFunc := node.Value.(*parser.NFunc)

		// 不允许嵌套函数定义
		if cmpl.InFunc {
			return fmt.Errorf("Function cannot be defined inside another function")
		}

		finfo := FuncInfo{
			Name:      nFunc.Name,
			ParamsNum: len(nFunc.Params),
		}

		if _, ok := cmpl.FuncTable[nFunc.Name]; ok {
			return fmt.Errorf("Function %s hasn't been defined\n", nFunc.Name)
		}

		start := int64(len(cmpl.Code))
		cmpl.AppendCode(runtime.JMP, 0)
		finfo.Offset = start + 2

		// 正在编译函数
		cmpl.InFunc = true

		if err = cmpl.InitVars(node, nFunc.Params); err != nil {
			return err
		}

		// 编译函数体
		varCount := len(cmpl.VarTable)
		if err = nodeToCode(cmpl, nFunc.Body); err != nil {
			return err
		}

		// 统计函数体局部变量的数量
		localVarNum := len(cmpl.VarTable) - varCount
		finfo.LocalVarNum = localVarNum

		// 离开函数编译
		cmpl.InFunc = false

		// 如果函数最后没有return关键字, 则在指令流中插入RETFUNC
		if cmpl.Code[len(cmpl.Code)-1] != runtime.RETFUNC {
			cmpl.AppendCode(runtime.RETFUNC)
		}

		// 跳出函数定义
		funcEnd := int64(len(cmpl.Code)) - start
		cmpl.Code[start+1] = BCode(funcEnd)

		// 在函数表中保存
		cmpl.FuncTable[nFunc.Name] = finfo

	case parser.TCallFunc:

	}

	return nil
}
