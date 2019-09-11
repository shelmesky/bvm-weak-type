# BVM虚拟机和编译器

##  词法分析和语法分析

词法分析器使用**flex**生成, 此法规则文件保存在```parser/lex.l```. 

语法生成器使用**bison**生成, 语法文件保存在```parser/parser.y```. 

## 语法示例

```go
contract mycnt {

    func myfunc0(b) {
        var x = b
        return x
    }

    func myfunc1(a, c) {
        var r = myfunc0(333)
        return a+c+r
    }

    var y = myfunc1(111, 222)
    println(y)
}
```

## 功能列表
 - 内部函数调用
 - 外部函数调用
 - if/else for switch while
 - 数据类型: 整数 字符串 map list
 - 逻辑和关系: && || ! >= <= > <
 - 运算符: + - * / += -= *= /= % ^
 
 ## 安装方式:
 
 1. 建立golang开发环境
 2. 建立开发目录: 在当前目录执行 ```mkdir -p bvm/src/bvm```
 3. 设置GOPATH: 进入```bvm```目录执行: ```export GOPATH=$(pwd)```
 4. 克隆代码: 进入```src/bvm```, 执行: ```git clone https://gitlab.com/bottos-project/bvm.git .```
 5. 编译代码: 进入代码的```cmd```目录下, 执行: ```go build main.go```
 6. 运行测试代码: 在代码的cmd目录下运行: ```./main test1.contract```
 