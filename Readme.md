# GS

**GS**是一个简单、高效的动态语言，语法和 go 类似，源代码被编译为字节码后由go虚拟机解释执行，可以无缝的与 go 代码集成。

## examples

斐波那契函数

```golang
program := `
func fact(n) {
    if n < 2 {
        return 1
    }
    return n * fact(n-1)
}
println(fact(10))
`
gs.Eval(program, nil)
```

循环

```golang
program := `
arr = []
for i = range 10 {
    arr = append(arr, i)
}
println(arr) // [0 1 2 3 4 5 6 7 8 9]
`
gs.Eval(program, nil)
```

### 与go代码集成

表达式计算

```golang
program := `
return $.A + $.B
`
code, _ := gs.Compile(program)
ret, _ := gs.Run(code, &Env{A: 1, B: 2})
fmt.Println(ret.MustInt()) // 3
```

条件判断

```golang
program := `
return $.A == 1 && $.B == 2
`
code, _ := gs.Compile(program)
ret, _ := gs.Run(code, &Env{A: 1, B: 2})
fmt.Println(ret.MustBool()) // true

```

自定义函数

```golang
program := `
return add($.A + $.B)
`
code, _ := gs.Compile(program, gs.DefineFuncs(gs.Func{
    Name: "add",
    Fn: func(a, b int) int {
        return a + b
    },
}))
ret, _ := gs.Run(code, &Env{A: 1, B: 2})
fmt.Println(ret.MustInt()) // 3
```

甚至赋值

```golang
program := `
$.Sum = add($.A + $.B)
`
code, _ := gs.Compile(program, gs.DefineFuncs(gs.Func{
    Name: "add",
    Fn: func(a, b int) int {
        return a + b
    },
}))
env := &Env{A: 1, B: 2, Sum: 0}
gs.Run(code, env)
fmt.Println(env.Sum) // 3
```

## 数据类型

|类型|对应 go 语言类型|分类|
|--|--|--|
|int|uint,int,uintX,intX|基本类型|
|float64|float32,float64|基本类型|
|string|string|基本类型|
|true|true|基本类型|
|false|false|基本类型|
|[]|[]any|复合类型|
|{}|map\[any]\[any]|复合类型|
|struct|自定义类型无对应|复合类型|
|nil|nil|空类型|

## 运算

优先级从低到高

|类型| 符号                      |优先级|
|--|-------------------------|--|
|或| \|                      |低|
|与| &&                      ||
|比较| >,<,=,!=,==,>=,<=,>>,<< ||
|位运算| &,\|,^                  ||
|加减| +,-                     ||
|乘除| *,/                     ||
|指针解引用（来自 go 的指针）| *                       |高|

## 语法定义

**GS**语法类似 go，主要差异：

- 简化了类型系统，采用动态类型，不支持指针，都是引用
- 简化了作用域系统，作用域只有两级：函数内和全局

### 字面量

|类型| 示例                                                                                                                          |
|--|-----------------------------------------------------------------------------------------------------------------------------|
|int| 十进制整数：12345<br/>下划线：10_000<br/>二进制：0b01010101,0B01010101,0B1110_01111<br/>八进制：0o755333,0O1234_4567<br/>十六进制：0xaB_Cd,0X12_AC |
|float| 10.3, .2, 123_456.789, 123.45_6e-7, 123.                                                                                    |
|string| 单行字符串："abc"<br/>多行字符串：\`<br/>“first line” <br/>“second line”<br/>\`                                                         |
|[]| 切片 [1,2,3\], ["a","b","c"\], ["1",1,true\]                                                                                  |
|{}| Map {}, {"1":1,"2",2}, {1:"a",2:"b"}, {true:1,false:0}                                                                      |

### 类型定义

函数定义，类似 go：

1. 无需声明参数类型和返回值
2. 支持返回多个值

```
func swap(x,y) {
  return y, x
}
```

结构体定义，类似 go:

```
type MyStruct struct {Field0, Filed1, Field2, ...}
```

### 语句

基本语句，对齐 go：

|语句类型|示例|
|--|--|
|变量声明必须赋值|x = 10<br/>x, y  =  1, 2|
|一元元算|x-=2, x+=2, x/=2, x*=3|
|自加/自减|i++, i--|
|负数|-a, -10|
|比较|>, < , ==, >=, <=, !=|
|运算|+, -, *, /, &, \|, ^|
|返回，支持多返回值|return x, y|
|终止循环|break|
|继续循环|continue|
|切片|arr[1:2], arr[1:], arr[:10], arr[:]|

for 循环，对齐 go：

```
// C 风格
for i = 0; i < 10; i++ {
  println(i)
}
// for range 支持整数，slice，map
for i = range 10 {
  println(i)
}
for i, v = range [1,2,3,4,5] {
  println(i,v)
}
for k, v = range {"a": 1, "b": 2, "c": 3} {
  println(k, v)
}
// 单条件循环
for i < 10 {
  println(i)
  i++
}
```

if 语句，对齐 go：

```
if x > 10 {
  println(x, " is less than 10")
} else {
  println(x, " is not less than 10"
}
// 支持添加初始化语句
if sum = a + b; sum < 10 {
  println("a + b  is less than 10")
} else {
  println("a + b  is not less than 10")
}
```

## 与go类型系统的集成

#### 访问go值

- 访问结构体用点`.`
- 访问数组/Map 用方括号`[]`
- 访问指针类型的成员：支持自动解引用，如 *someStruct.Field，可以直接写 someStruct.Field

#### 基本类型自动转换

go 原生计算和比较，要求两个类型必须相同，比如 int 和 uint 不能直接计算，为了简化，虚拟机中在部分场景中支持自动转换:

- 所有基本类型的计算，比较
- 给 go 基本类型赋值
- 内置 in 函数，入参为数组时
- 内置 append 函数

#### 赋值go 值

当把 vm 中的值赋值给 go 类型时，遵循下面的原则：

- vm 中基本类型 -> go 中基本类型会**自动适配和强转**，如`intX/uintX/float32/float64`
- 复合类型赋值支持：
	- vm中 `{}` <-> go `map[any][any]`
	- vm 中`[]` <-> go `[]any`
	- go 中任意类型 -> go 中**相同**类型
	- 赋值给 go 中指针，除了类型要**严格**相同外，`*`操作符要求是对 env 的**直接操作**：即"$"开头，如：
		- `*$.A.B = []{1,2,3}`
		- `*$.C = "aaaa"`

	  赋值给指针，支持给空指针解引用赋值，但要求类型严格相等，即使是基本类型，也不会自动强转

## 内置函数支持

### 基本函数

|函数|作用|
|--|--|
|append|切片追加元素，支持三点符号...|
|delete|删除元素|
|len|获取长度|
|uint/uint8/uintX/int/int8/intX/float32/float64/string/duration|类型强转|
|print/printf/println|打印变量到标准输出，可自定义输出 io|
|sprintf|格式化字符串，对齐 go|

### 集合

|函数|作用|
|--|--|
|in(arr/map, value)|查询 value 是否在 arr/map 的 key 中|
|sort|排序，只支持基本类型数组|
|||

### 字符串

签名完全对齐 go

|函数|作用|
|--|--|
| index| 索引|
| hasPrefix|是否有 prefix|
|hasSuffix|有否有 suffix|
|trim|去除两端空白|
|trimPrefix|去除 prefix|
|trimSuffix|去除 suffix|
|trimSpace|去除空白|
|trimLeft|去除左侧，cutset 匹配|
|trimRight|去除右侧，cutset 匹配|
|toLower|转小写|
|toUpper|转大写|
|split|拆分|
|join|合并|
|atoi|string->int64|
|itoa|int64->string|

### 时间

签名完全对齐 go

|函数|作用|
|--|--|
|now|获取当前时间|
|parseTime|解析字符串为时间|
|parseDuration|解析字符串为 duration|

### json

|函数|作用|
|--|--|
|toJson| 序列化为 json 字符串|
|fromJson| 反序列化为 map[string\][any\]|
|unmarshalJson| 对齐 json.Unmarhsal|
