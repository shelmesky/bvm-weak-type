package runtime

type EmbedFunc struct {
	Index     int
	Name      string      // 函数名称
	ParamNum  int         // 参数数量
	ListArgs  bool        // 是否支持变长参数
	HasReturn bool        // 是否有返回值
	Func      interface{} // 函数对象
}

var (
	Stdlib = []EmbedFunc{
		{Index: 0, Name: `println`, ParamNum: 1, ListArgs: false, HasReturn: false, Func: Println},
	}
)

func GetEmbedFunc(name string) *EmbedFunc {
	var embedFunc *EmbedFunc

	for idx := range Stdlib {
		embedFunc = &Stdlib[idx]
		if embedFunc.Name == name {
			return embedFunc
		}
	}

	return embedFunc
}

func Println() int64 {
	return 0
}
