package compiler

import (
	"bvm/parser"
	"bvm/runtime"
	"bvm/utils"
	"fmt"
	"math"
)

type BCode uint16

// 函数信息
type FuncInfo struct {
	Index       int    // 函数在列表中的索引
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

type Jumps struct {
	Breaks    []int
	Continues []int
}

// 编译环境
type CompileEnv struct {
	VarTable       map[string]Variable // 变量表
	FuncTable      map[string]FuncInfo // 函数表
	FuncList       []FuncInfo          // 有序函数表
	ConstantsTable []Const             // 常量表
	Code           []BCode             // 字节码
	InFunc         bool                // 是否正在编译函数
	Jumps          []*Jumps
}

func (this *CompileEnv) AppendCode(codes ...BCode) {

	if len(codes) > 0 {
		switch codes[0] {
		case runtime.PUSH:
			utils.DebugPrintf("Compile>  PUSH index:[%d]\n", codes[1])
		case runtime.INITVARS:
			utils.DebugPrintf("Compile>  INITVARS")
		case runtime.GETVAR:
			utils.DebugPrintf("Compile>  GETVAR index:[%d]\n", codes[1])
		case runtime.SETVAR:
			utils.DebugPrintf("Compile>  SETVAR index:[%d]\n", codes[1])
		case runtime.ADD:
			utils.DebugPrintln("Compile>  ADD")
		case runtime.SUB:
			utils.DebugPrintln("Compile>  UB")
		case runtime.MUL:
			utils.DebugPrintln("Compile>  MUL")
		case runtime.DIV:
			utils.DebugPrintln("Compile>  DIV")
		case runtime.ASSIGN:
			utils.DebugPrintf("Compile>  ASSIGN\n")
		case runtime.LOOP:
			utils.DebugPrintln("Compile>  LOOP")
		case runtime.JMP:
			utils.DebugPrintf("Compile>  JMP [%d]\n", codes[1])
		case runtime.JZE:
			utils.DebugPrintf("Compile>  JZE [%d]\n", codes[1])
		case runtime.RETFUNC:
			utils.DebugPrintf("Compile>  RETFUNC expr:[%d]\n", codes[1])
		case runtime.RETURN:
			utils.DebugPrintln("Compile>  RETURN")
		case runtime.CALLFUNC:
			utils.DebugPrintf("Compile>  CALLFUNC offset:[%d]\n", codes[1])
		case runtime.GETPARAMS:
			utils.DebugPrintf("Compile>  GETPARAMS count:[%d] ", codes[1])
			varIdxList := codes[2:]
			if len(varIdxList) > 0 {
				utils.DebugPrintf("varIdxList:[ ")
				for i := 0; i < len(varIdxList); i++ {
					utils.DebugPrintf("%d ", varIdxList[i])
				}
				utils.DebugPrintf("]")
			}
			utils.DebugPrintf("\n")

		case runtime.CALLEMBED:
			utils.DebugPrintf("Compile>  CALLEMBED %d\n", codes[1])

		case runtime.AND:
			utils.DebugPrintf("Compile>  AND\n")
		case runtime.OR:
			utils.DebugPrintf("Compile>  OR\n")
		case runtime.EQ:
			utils.DebugPrintf("Compile>  EQ\n")
		case runtime.NOTEQ:
			utils.DebugPrintf("Compile>  NOTEQ\n")
		case runtime.NOT:
			utils.DebugPrintf("Compile>  NOT\n")
		case runtime.LT:
			utils.DebugPrintf("Compile>  LT\n")
		case runtime.GT:
			utils.DebugPrintf("Compile>  GT\n")
		case runtime.LTE:
			utils.DebugPrintf("Compile>  LTE\n")
		case runtime.GTE:
			utils.DebugPrintf("Compile>  GTE\n")
		}
	}

	for _, code := range codes {
		this.Code = append(this.Code, code)
	}
}

func (this *CompileEnv) InitVars(node *parser.Node, vars []parser.NVar) ([]BCode, error) {
	var varIdxList []BCode
	if len(vars) == 0 {
		return varIdxList, fmt.Errorf("empty vars")
	}

	for _, v := range vars {
		if _, ok := this.VarTable[v.Name]; ok {
			return varIdxList, fmt.Errorf("variable %s already exists\n", v.Name)
		}

		idx := uint16(len(this.VarTable))
		symbol := Variable{
			Index:    idx,
			Name:     v.Name,
			IsGlobal: false,
		}

		varIdxList = append(varIdxList, BCode(idx))

		this.VarTable[v.Name] = symbol

		//this.AppendCode(runtime.INITVARS)
	}

	return varIdxList, nil
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
		return &cmpl, fmt.Errorf("compile failed: %s\n", err.Error())
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

		if len(block.Statements) > 0 {
			for _, child := range block.Statements {
				if err = nodeToCode(cmpl, child); err != nil {
					return err
				}
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
			return err
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
		case parser.AND:
			cmpl.AppendCode(runtime.AND)
		case parser.OR:
			cmpl.AppendCode(runtime.OR)
		case parser.EQ:
			cmpl.AppendCode(runtime.EQ)
		case parser.NOT_EQ:
			cmpl.AppendCode(runtime.NOTEQ)
		case parser.NOT:
			cmpl.AppendCode(runtime.NOT)
		case parser.LT:
			cmpl.AppendCode(runtime.LT)
		case parser.GT:
			cmpl.AppendCode(runtime.GT)
		case parser.LTE:
			cmpl.AppendCode(runtime.LTE)
		case parser.GTE:
			cmpl.AppendCode(runtime.GTE)
		}

	// 变量定义
	case parser.TVars:
		if _, err = cmpl.InitVars(node, node.Value.(*parser.NVars).Vars); err != nil {
			return err
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
			if expr != nil {
				// 如果return有表达式
				cmpl.AppendCode(runtime.RETFUNC, 1)
			} else {
				cmpl.AppendCode(runtime.RETFUNC, 0)
			}
		} else {
			cmpl.AppendCode(runtime.RETURN)
		}

	// 函数实现
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

		var varIdxList []BCode
		if varIdxList, err = cmpl.InitVars(node, nFunc.Params); err != nil {
			return err
		}

		start := int64(len(cmpl.Code))
		cmpl.AppendCode(runtime.JMP, 0)
		finfo.Offset = start + 2

		// 正在编译函数
		cmpl.InFunc = true

		if len(nFunc.Params) > 0 {
			var getParamIns []BCode
			getParamIns = append(getParamIns, runtime.GETPARAMS)
			getParamIns = append(getParamIns, BCode(len(nFunc.Params)))
			getParamIns = append(getParamIns, varIdxList...)
			cmpl.AppendCode(getParamIns...)
		}

		// 编译函数体
		varCount := len(cmpl.VarTable)
		if nFunc.Body != nil {
			if err = nodeToCode(cmpl, nFunc.Body); err != nil {
				return err
			}
		}

		// 统计函数体局部变量的数量
		localVarNum := len(cmpl.VarTable) - varCount
		finfo.LocalVarNum = localVarNum

		// 离开函数编译
		cmpl.InFunc = false

		// 如果函数最后没有return关键字, 则在指令流中插入RETFUNC
		// -2是因为RETFUNC有一个参数
		if cmpl.Code[len(cmpl.Code)-2] != runtime.RETFUNC {
			cmpl.AppendCode(runtime.RETFUNC, 0)
		}

		// 跳出函数定义
		funcEnd := int64(len(cmpl.Code))
		cmpl.Code[start+1] = BCode(funcEnd)

		// 在函数表中保存
		finfo.Index = len(cmpl.FuncTable)
		cmpl.FuncTable[nFunc.Name] = finfo
		cmpl.FuncList = append(cmpl.FuncList, finfo)

	// 函数调用
	case parser.TCallFunc:
		nFunc := node.Value.(*parser.NCallFunc)

		// 优先在标准库中查找函数
		embedFunc, err := runtime.GetEmbedFunc(nFunc.Name)
		if err == nil {
			// 编译实参
			if nFunc.Params != nil {
				paramsList := nFunc.Params.Value.(*parser.NParams)
				for _, expr := range paramsList.Expr {
					if err = nodeToCode(cmpl, expr); err != nil {
						return err
					}
				}

				// 如果提供的参数数量不等于实际数量
				if len(paramsList.Expr) != embedFunc.ParamNum {
					return fmt.Errorf("Call embed function [%s] need %d arguments, got %d.\n",
						nFunc.Name, embedFunc.ParamNum, len(paramsList.Expr))
				}
			}

			cmpl.AppendCode(runtime.CALLEMBED, BCode(embedFunc.Index))

		} else {

			var fInfo FuncInfo
			if fInfo, ok = cmpl.FuncTable[nFunc.Name]; !ok {
				return fmt.Errorf("Function %s hasn't been defined\n", nFunc.Name)
			}

			// 编译实参
			if nFunc.Params != nil {
				paramsList := nFunc.Params.Value.(*parser.NParams)
				for _, expr := range paramsList.Expr {
					if err = nodeToCode(cmpl, expr); err != nil {
						return err
					}
				}

				// 如果提供的参数数量不等于实际函数的参数
				if len(paramsList.Expr) != fInfo.ParamsNum {
					return fmt.Errorf("Call function [%s] need %d arguments, got %d.\n",
						nFunc.Name, fInfo.ParamsNum, len(paramsList.Expr))
				}
			} else {
				// 如果调用时未提供参数, 但函数有参数则报错
				if fInfo.ParamsNum > 0 {
					return fmt.Errorf("Call function [%s] need %d arguments, got 0.\n",
						nFunc.Name, fInfo.ParamsNum)
				}
			}

			//offset := fInfo.Offset - int64(len(cmpl.Code))
			idx := fInfo.Index
			cmpl.AppendCode(runtime.CALLFUNC, BCode(idx))
		}

	// if/elif/else语句
	case parser.TIf:
		ends := make([]int, 0, 16)
		nIf := node.Value.(*parser.NIf)

		// 编译if的表达式
		_, sizeCond, err := cmpl.ConditionCode(nIf.Cond)
		if err != nil {
			return err
		}
		cmpl.AppendCode(runtime.JZE, 0)

		if nIf.IfBody != nil {
			// 编译if的代码块
			if err = nodeToCode(cmpl, nIf.IfBody); err != nil {
				return err
			}
		}

		ends = append(ends, len(cmpl.Code))
		cmpl.AppendCode(runtime.JMP, 0)

		var off BCode
		if off, err = cmpl.JumpOff(node, len(cmpl.Code)); err != nil {
			return err
		}
		cmpl.Code[sizeCond+1] = off

		// 如果含有多个elif表达式
		if nIf.ElifBody != nil {
			nElif := nIf.ElifBody.Value.(*parser.NElif)
			for _, child := range nElif.List {
				_, sizeCond, err = cmpl.ConditionCode(child.Cond)
				if err != nil {
					return err
				}

				cmpl.AppendCode(runtime.JZE, 0)

				if err = nodeToCode(cmpl, child.Body); err != nil {
					return err
				}

				ends = append(ends, len(cmpl.Code))

				cmpl.AppendCode(runtime.JMP, 0)

				if off, err = cmpl.JumpOff(node, len(cmpl.Code)); err != nil {
					return err
				}

				cmpl.Code[sizeCond+1] = off
			}
		}

		// 如果含有else语句
		if nIf.ElseBody != nil {
			if err = nodeToCode(cmpl, nIf.ElseBody); err != nil {
				return err
			}
		}

		size := len(cmpl.Code)
		for _, end := range ends {
			if off, err = cmpl.JumpOff(node, size); err != nil {
				return err
			}

			cmpl.Code[end+1] = off
		}

	case parser.TWhile:
		nWhile := node.Value.(*parser.NWhile)

		cmpl.AppendCode(runtime.LOOP)

		cmpl.Jumps = append(cmpl.Jumps, &Jumps{})
		// 处理while的条件表达式
		// sizeCode为条件代码结束的位置, sizeCond为条件代码开始的位置
		sizeCode, sizeCond, err := cmpl.ConditionCode(nWhile.Cond)
		if err != nil {
			return err
		}

		// 紧跟着条件表达式插入JZE指令
		cmpl.AppendCode(runtime.JZE, 0)

		// 编译while的主体代码
		if err = nodeToCode(cmpl, nWhile.Body); err != nil {
			return err
		}

		// 循环处理Body中出现的continue语句
		// 给continue关键字生成的JMP指令设置为条件语句开始的地方
		for _, b := range cmpl.Jumps[len(cmpl.Jumps)-1].Continues {
			cmpl.Code[b] = BCode(sizeCode)
		}

		// while中一次循环运行完毕, 跳转到条件表达式开始的地方
		var off BCode
		if off, err = cmpl.JumpOff(node, sizeCode-1); err != nil {
			return err
		}
		cmpl.AppendCode(runtime.JMP, off)

		// 设置紧跟着条件表达式的JZE指令参数
		// 如果JZE检查出条件为false, 则直接跳到while语句结束的地方运行
		if off, err = cmpl.JumpOff(node, len(cmpl.Code)); err != nil {
			return err
		}
		cmpl.Code[sizeCond+1] = off

		// 设置代码中出现的break关键字的跳转位置为代码结束处
		for _, b := range cmpl.Jumps[len(cmpl.Jumps)-1].Breaks {
			cmpl.Code[b] = BCode(len(cmpl.Code))
		}

		cmpl.Jumps = cmpl.Jumps[:len(cmpl.Jumps)-1]

	case parser.TBreak:
		if len(cmpl.Jumps) == 0 {
			return fmt.Errorf("break must be inside of while or for")
		}

		// 生成JMP指令
		// 使用Jump.Breaks记录当前的位置
		// while中会给这里生成的JMP指令设置真正需要跳转的位置
		cmpl.AppendCode(runtime.JMP, 0)
		cmpl.Jumps[len(cmpl.Jumps)-1].Breaks = append(cmpl.Jumps[len(cmpl.Jumps)-1].Breaks,
			len(cmpl.Code)-1)

	case parser.TContinue:
		if len(cmpl.Jumps) == 0 {
			return fmt.Errorf("continue must be inside of while or for")
		}

		// 生成JMP指令
		// 使用Jump.Continues记录当前的代码位置
		cmpl.AppendCode(runtime.JMP, 0)
		cmpl.Jumps[len(cmpl.Jumps)-1].Continues = append(cmpl.Jumps[len(cmpl.Jumps)-1].Continues,
			len(cmpl.Code)-1)
	}

	return nil
}

func (this *CompileEnv) ConditionCode(node *parser.Node) (before, after int, err error) {
	before = len(this.Code)
	if err = nodeToCode(this, node); err != nil {
		return
	}

	after = len(this.Code)
	return
}

func (this *CompileEnv) JumpOff(node *parser.Node, off int) (BCode, error) {
	if off < math.MinInt16 || off > math.MaxInt16 {
		return runtime.NOP, fmt.Errorf("Too big relative jump\n")
	}

	return BCode(off), nil
}
