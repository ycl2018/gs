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

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:20%">类型</th>
      <th style="width:50%">对应 go 语言类型</th>
      <th style="width:30%">分类</th>
    </tr>
  </thead>
  <tbody>
    <tr><td>int</td><td>uint,int,uintX,intX</td><td>基本类型</td></tr>
    <tr><td>float64</td><td>float32,float64</td><td>基本类型</td></tr>
    <tr><td>string</td><td>string</td><td>基本类型</td></tr>
    <tr><td>true</td><td>true</td><td>基本类型</td></tr>
    <tr><td>false</td><td>false</td><td>基本类型</td></tr>
    <tr><td>[]</td><td>[]any</td><td>复合类型</td></tr>
    <tr><td>{}</td><td>map[any][any]</td><td>复合类型</td></tr>
    <tr><td>struct</td><td>自定义类型无对应</td><td>复合类型</td></tr>
    <tr><td>nil</td><td>nil</td><td>空类型</td></tr>
  </tbody>
</table>

## 运算

优先级从低到高

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:20%">类型</th>
      <th style="width:60%">符号</th>
      <th style="width:20%">优先级</th>
    </tr>
  </thead>
  <tbody>
    <tr><td>或</td><td>|</td><td>低</td></tr>
    <tr><td>与</td><td>&&</td><td></td></tr>
    <tr><td>比较</td><td>&gt;,&lt;,=,!=,==,&gt;=,&lt;=,&gt;&gt;,&lt;&lt;</td><td></td></tr>
    <tr><td>位运算</td><td>&amp;,|,^</td><td></td></tr>
    <tr><td>加减</td><td>+,-</td><td></td></tr>
    <tr><td>乘除</td><td>*,/</td><td></td></tr>
    <tr><td>指针解引用（来自 go 的指针）</td><td>*</td><td>高</td></tr>
  </tbody>
</table>

## 语法定义

**GS**语法类似 go，主要差异：

- 简化了类型系统，采用动态类型，不支持指针，都是引用
- 简化了作用域系统，作用域只有两级：函数内和全局

### 字面量

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:20%">类型</th>
      <th style="width:80%">示例</th>
    </tr>
  </thead>
  <tbody>
    <tr><td>int</td><td>十进制整数：12345<br/>下划线：10_000<br/>二进制：0b01010101,0B01010101,0B1110_01111<br/>八进制：0o755333,0O1234_4567<br/>十六进制：0xaB_Cd,0X12_AC</td></tr>
    <tr><td>float</td><td>10.3, .2, 123_456.789, 123.45_6e-7, 123.</td></tr>
    <tr><td>string</td><td>单行字符串："abc"<br/>多行字符串：`<br/>“first line” <br/>“second line”<br/>`</td></tr>
    <tr><td>[]</td><td>切片 [1,2,3], ["a","b","c"], ["1",1,true]</td></tr>
    <tr><td>{}</td><td>Map {}, {"1":1,"2",2}, {1:"a",2:"b"}, {true:1,false:0}</td></tr>
  </tbody>
</table>

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

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:30%">语句类型</th>
      <th style="width:70%">示例</th>
    </tr>
  </thead>
  <tbody>
    <tr><td>变量声明必须赋值</td><td>x = 10<br/>x, y = 1, 2</td></tr>
    <tr><td>一元运算</td><td>x-=2, x+=2, x/=2, x*=3</td></tr>
    <tr><td>自加/自减</td><td>i++, i--</td></tr>
    <tr><td>负数</td><td>-a, -10</td></tr>
    <tr><td>比较</td><td>&gt;, &lt;, ==, &gt;=, &lt;=, !=</td></tr>
    <tr><td>运算</td><td>+, -, *, /, &amp;, |, ^</td></tr>
    <tr><td>返回，支持多返回值</td><td>return x, y</td></tr>
    <tr><td>终止循环</td><td>break</td></tr>
    <tr><td>继续循环</td><td>continue</td></tr>
    <tr><td>切片</td><td>arr[1:2], arr[1:], arr[:10], arr[:]</td></tr>
  </tbody>
</table>

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

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:30%">函数</th>
      <th style="width:70%">作用</th>
    </tr>
  </thead>
  <tbody>
    <tr><td>append</td><td>切片追加元素，支持三点符号...</td></tr>
    <tr><td>delete</td><td>删除元素</td></tr>
    <tr><td>len</td><td>获取长度</td></tr>
    <tr><td>uint/uint8/uintX/int/int8/intX/float32/float64/string/duration</td><td>类型强转</td></tr>
    <tr><td>print/printf/println</td><td>打印变量到标准输出，可自定义输出 io</td></tr>
    <tr><td>sprintf</td><td>格式化字符串，对齐 go</td></tr>
  </tbody>
</table>

### 集合

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:30%">函数</th>
      <th style="width:70%">作用</th>
    </tr>
  </thead>
  <tbody>
    <tr><td>in(arr/map, value)</td><td>查询 value 是否在 arr/map 的 key 中</td></tr>
    <tr><td>sort</td><td>排序，只支持基本类型数组</td></tr>
  </tbody>
</table>


### 字符串

签名完全对齐 go

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:30%">函数</th>
      <th style="width:70%">作用</th>
    </tr>
  </thead>
  <tbody>
    <tr><td>index</td><td>索引</td></tr>
    <tr><td>hasPrefix</td><td>是否有 prefix</td></tr>
    <tr><td>hasSuffix</td><td>是否有 suffix</td></tr>
    <tr><td>trim</td><td>去除两端空白</td></tr>
    <tr><td>trimPrefix</td><td>去除 prefix</td></tr>
    <tr><td>trimSuffix</td><td>去除 suffix</td></tr>
    <tr><td>trimSpace</td><td>去除空白</td></tr>
    <tr><td>trimLeft</td><td>去除左侧，cutset 匹配</td></tr>
    <tr><td>trimRight</td><td>去除右侧，cutset 匹配</td></tr>
    <tr><td>toLower</td><td>转小写</td></tr>
    <tr><td>toUpper</td><td>转大写</td></tr>
    <tr><td>split</td><td>拆分</td></tr>
    <tr><td>join</td><td>合并</td></tr>
    <tr><td>atoi</td><td>string->int64</td></tr>
    <tr><td>itoa</td><td>int64->string</td></tr>
  </tbody>
</table>

### 时间

签名完全对齐 go

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:30%">函数</th>
      <th style="width:70%">作用</th>
    </tr>
  </thead>
  <tbody>
    <tr><td>now</td><td>获取当前时间</td></tr>
    <tr><td>parseTime</td><td>解析字符串为时间</td></tr>
    <tr><td>parseDuration</td><td>解析字符串为 duration</td></tr>
  </tbody>
</table>

### json

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:30%">函数</th>
      <th style="width:70%">作用</th>
    </tr>
  </thead>
  <tbody>
    <tr><td>toJson</td><td> 序列化为 json 字符串</td></tr>
    <tr><td>fromJson</td><td> 反序列化为 map[string][any]</td></tr>
    <tr><td>unmarshalJson</td><td> 对齐 json.Unmarhsal</td></tr>
  </tbody>
</table>
