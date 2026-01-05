# GS

**GS** 是一个简单、高性能的动态语言，语法和 go 类似，源代码会被编译并优化为字节码后由go虚拟机解释执行，可以无缝的与 go 代码集成。

适用于构建：
- [x] 表达式引擎
- [x] 规则引擎
- [x] 动态配置
- [x] 脚本
- [x] 插件
- [x] 配置校验

**凡是需要动态化的场景，都可以考虑使用 GS，实现高性能不停机热更新代码逻辑。**

开始使用 `go get github.com/ycl2018/gs`

---

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
println(fact(10)) // 3628800
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

错误提示

```golang
program := `            // line 1
func divide(a, b) {     // line 2
    return a / b        // line 3
}                       // line 4
divide(1, 0)            // line 5
`
_, err :=gs.Eval(program, nil)
fmt.Println(err)
/*
panic: division by zero
stack trace:
at divide args:(1, 0) line:3
at main args:() line:5
 */
```

**代码调试**

通过 `DumpCode` 和 `Trace` (需配置build tag `gs_trace`) 选项可以开启代码调试功能，打印字节码和执行 trace。

```text
Dump:
.globals: 0

Constant Pool:2
#0000: func main(args:0, locals:0)
#0001: func fact(args:1, locals:0)

Code:
.def main args=0 locals=0
#0000 iconst     	3
#0001 call       	const#1
#0002 println    	1
#0003 halt       	
.def fact args=1 locals=0
#0000 load       	0
#0001 iconst     	2
#0002 lt         	
#0003 brf        	6
#0004 iconst     	1
#0005 ret        	
#0006 load       	0
#0007 load       	0
#0008 iconst     	1
#0009 sub        	
#0010 call       	const#1
#0011 mul        	
#0012 ret 

trace:
	stack=[ ], calls=[ main ]
->0000: iconst     	3
	stack=[ 3 ], calls=[ main ]
->0001: call       	const#1
	stack=[ ], calls=[ main fact ]
->0000: load       	0
	stack=[ 3 ], calls=[ main fact ]
->0001: iconst     	2
	stack=[ 3 2 ], calls=[ main fact ]
->0002: lt         
	stack=[ false ], calls=[ main fact ]
->0003: brf        	6
	stack=[ ], calls=[ main fact ]
...
	stack=[ 3 2 1 ], calls=[ main fact fact fact ]
->0001: iconst     	2
	stack=[ 3 2 1 2 ], calls=[ main fact fact fact ]
->0002: lt         
	stack=[ 3 2 true ], calls=[ main fact fact fact ]
->0003: brf        	6
	stack=[ 3 2 ], calls=[ main fact fact fact ]
->0004: iconst     	1
	stack=[ 3 2 1 ], calls=[ main fact fact fact ]
->0005: ret        
	stack=[ 3 2 1 ], calls=[ main fact fact ]
->0011: mul        
	stack=[ 3 2 ], calls=[ main fact fact ]
->0012: ret        
	stack=[ 3 2 ], calls=[ main fact ]
->0011: mul        
	stack=[ 6 ], calls=[ main fact ]
->0012: ret        
	stack=[ 6 ], calls=[ main ]
->0002: println    	1
	stack=[ ], calls=[ main ]
->0003: halt       

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

// 自定义函数 add
code, _ := gs.Compile(program, gs.DefineFuncs(gs.Func{
    Name: "add",
	// Fn 可以为任意函数签名
    Fn: func(a, b int) int {
        return a + b
    },
}))

ret, _ := gs.Run(code, &Env{A: 1, B: 2})
fmt.Println(ret.MustInt()) // 3

// 为了更好的调用性能，避免使用反射，建议使用 `gs.DefineFast` 定义函数
gs.DefineFast(gs.FastFunc{
    Name: "add",
    // 定义函数 add，接收 2 个参数，返回 1 个参数
    NumIn: 2,
    NumOut: 1,
    // 函数签名必须为 func(args []any) []any
    Fn: func(args []any) []any {
        return []any{args[0].(int) + args[1].(int)}
    },
})
```

赋值

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

甚至并发

```golang
f1 = go $.Fn("World")           // 异步执行任务 1
f2 = go define("define func")   // 异步执行任务 2
val1, err = f1.Wait()           // 阻塞并获取结果
val2, err = f2.Wait()           // 阻塞并获取结果
println(val1, err)
println(val2, err)
```

异步并发由 `go` 关键字触发，仅支持执行go原生函数。
每个并发任务返回一个 `Future` 对象，通过 `Wait` 方法阻塞并等待任务完成，获取返回值和错误。
并发能力借助 [github.com/ycl2018/go-future](https://github.com/ycl2018/go-future) 提供。

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
    <tr><td><code>int</code></td><td><code>uint,int,uintX,intX</code></td><td>基本类型</td></tr>
    <tr><td><code>float64</code></td><td><code>float32,float64</code></td><td>基本类型</td></tr>
    <tr><td><code>string</code></td><td><code>string</code></td><td>基本类型</td></tr>
    <tr><td><code>true</code></td><td><code>true</code></td><td>基本类型</td></tr>
    <tr><td><code>false</code></td><td><code>false</code></td><td>基本类型</td></tr>
    <tr><td><code>[]</code></td><td><code>[]any</code></td><td>复合类型</td></tr>
    <tr><td><code>{}</code></td><td><code>map[any][any]</code></td><td>复合类型</td></tr>
    <tr><td><code>struct</code></td><td><code>自定义类型无对应</code></td><td>复合类型</td></tr>
    <tr><td><code>nil</code></td><td><code>nil</code></td><td>空类型</td></tr>
  </tbody>
</table>

## 运算

优先级从低到高，跟 go 语言一致。

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:20%">类型</th>
      <th style="width:60%">符号</th>
      <th style="width:20%">优先级</th>
    </tr>
  </thead>
  <tbody>
    <tr><td><code>逻辑或</code></td><td><code>||</code></td><td>低</td></tr>
    <tr><td><code>逻辑与</code></td><td><code>&&</code></td><td></td></tr>
    <tr><td><code>关系运算</code></td><td><code>&gt;,&lt;,!=,==,&gt;=,&lt;=</code></td><td></td></tr>
    <tr><td><code>加减运算</code></td><td><code>+,-,|,^</code></td><td></td></tr>
    <tr><td><code>乘除运算</code></td><td><code>*,/,%,<<,>>,&</code></td><td></td></tr>
    <tr><td><code>一元运算</code></td><td><code>*,!,-</code></td><td>高</td></tr>
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
    <tr><td><code>int</code></td><td><code>十进制整数：12345<br/>下划线：10_000<br/>二进制：0b01010101,0B01010101,0B1110_01111<br/>八进制：0o755333,0O1234_4567<br/>十六进制：0xaB_Cd,0X12_AC</code></td></tr>
    <tr><td><code>float64</code></td><td><code>10.3, .2, 123_456.789, 123.45_6e-7, 123.</code></td></tr>
    <tr><td><code>string</code></td><td><code>单行字符串："abc"<br/>多行字符串：<br/>`<br/>“first line” <br/>“second line”<br/>`</code></td></tr>
    <tr><td><code>[]</code></td><td><code>slice： [1,2,3], ["a","b","c"], ["1",1,true]</code></td></tr>
    <tr><td><code>{}</code></td><td><code>Map： {}, {"1":1,"2",2}, {1:"a",2:"b"}, {true:1,false:0}</code></td></tr>
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

#### 基本语句，对齐 go：

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:30%">语句类型</th>
      <th style="width:70%">示例</th>
    </tr>
  </thead>
  <tbody>
    <tr><td>变量声明必须赋值</td><td><code>x = 10<br/> x, y = 1, 2</code></td></tr>
    <tr><td>一元运算</td><td><code>x-=2, x+=2, x/=2, x*=3</code></td></tr>
    <tr><td>自加/自减</td><td><code>i++, i--</code></td></tr>
    <tr><td>负数</td><td><code>-a, -10</code></td></tr>
    <tr><td>比较</td><td><code>&gt;, &lt;, ==, &gt;=, &lt;=, !=</code></td></tr>
    <tr><td>运算</td><td><code>+,-,*,/,&amp;,&lt;&lt;,^,&gt;&gt;</code></td></tr>
    <tr><td>返回，支持多返回值</td><td><code>return x, y</code></td></tr>
    <tr><td>终止循环</td><td><code>break</code></td></tr>
    <tr><td>继续循环</td><td><code>continue</code></td></tr>
    <tr><td>切片</td><td><code>arr[1:2], arr[1:], arr[:10], arr[:]</code></td></tr>
  </tbody>
</table>

#### for 循环，对齐 go：

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

#### if 语句，对齐 go：

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
// 支持 elseif 语句
i = 0
if i++; i < 1 {
  println("i < 1")
} else if i < 2 {
  println("i < 2")
} else {
  println("i >= 2")
}
```

#### switch 语句，对齐 go：

```
i = 0
for range 3 {
  switch i {
    case 0:
      println("i == 0")
    case 1:
      println("i == 1")
    default:
      println("i != 0 and i != 1")
  }
  i++
}
```

#### 并发支持

`go callExpr`

- `callExpr`必须是外部函数：自定义外部函数或者环境对象中的函数指针
- 限制函数签名返回值必须为两个，第一个为**任意类型**结果，第二个为错误接口`error`
- go 语句返回一个**异步对象**
- 异步对象可以用`Wait`方法等待结果，返回值为类型为`any`和错误`error`

```golang
gs.Eval(`
    f = go fast("World")
    // do something else ...
    ret, err = f.Wait()
    println(ret, err) // Hello World <nil>
    `,
    nil,
    DefineFunc(gs.Func{
        Name: "define",
        Fn: func(name string) (string, error) {
            return "Hello " + name, nil
        },
    })
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

**必须是可寻址的字段才能赋值**

- 由于底层使用反射，所以要求赋值的对象必须是可寻址的，比如 `env` 是指针类型时，才能通过 `$` 给其字段赋值

```golang 
$.String = "string"
```

或从 env 对象中获取的指针类型

```golang
ptr = $.Ptr
ptr.Field = "string"
// 注意： ptr 本身是不可寻址的，不能直接赋值，（ptr 只是一个指针变量，跟一个 uint 值没本质区别）
```

**基本类型**

- vm 中基本类型 -> go 中基本类型会**自动适配和强转**，如`intX/uintX/float32/float64`

**复合类型**

- vm中 `{}` <-> go `map[any][any]`
- vm 中`[]` <-> go `[]any`
- go 中任意类型 -> go 中**相同**类型

**指针类型**

- 基本类型指针可直接赋值，即使是空指针也支持，但**要求类型严格相等**

```golang
*$.Ptr = "hello world"
```

- 指针/map/slice，可通过 initRef 函数初始化，其中 map 可指定容量，slice 可指定长度，如：

```golang
// 初始化 map, 容量为 1
initRef($.Map,1)		
$.Map["foo"] = "bar"
println(len($.Map))

// 初始化 slice, 长度为 1
initRef($.Slice, 1)
$.Slice[0] = "hello"
println(len($.Slice))

// 初始化 Pointer
initRef($.MyEnv)		
$.MyEnv.A = "hello"
println($.MyEnv.A == "hello")
```

### 创建任意指针类型
- 内置函数 `newFromType(typeName)`，返回一个指向该typeName类型零值的指针
  - 支持所有 go 基本类型 `intX/uintX/floatX/string/duration/time`
  - 支持自定义类型

创建基础类型指针

```golang
stringPtr = newFromType("string")
*stringPtr = "hello world"
println(*stringPtr == "hello world") // true
```

自定义类型指针 `AddTypes`

```golang
gs.Eval(`
    myItfPtr = newFromType("myItf")
    myItfPtr.Name = "World"
    println(myItfPtr.Hello() == "Hello World") // true
    `,
    nil,
    gs.AddTypes(map[string]any{
    "myItf": MyItf{},
}))

type MyItf struct {
    Name string
}

func (m *MyItf) Hello() string {
    return "Hello " + m.Name
}
```

## 内置函数支持

内置函数是go标准库中函数的包装，签名完全对齐 go。

### 基本函数

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:30%">函数</th>
      <th style="width:70%">作用</th>
    </tr>
  </thead>
  <tbody>
    <tr><td><code>append</code></td><td>切片追加元素，支持三点符号...</td></tr>
    <tr><td><code>delete</code></td><td>删除元素</td></tr>
    <tr><td><code>len</code></td><td>获取长度</td></tr>
    <tr><td><code>uintX/intX/float32/float64/string/duration</code></td><td>类型强转</td></tr>
    <tr><td><code>print/printf/println</code></td><td>打印变量到标准输出，可自定义输出 io</td></tr>
    <tr><td><code>sprintf</code></td><td>格式化字符串，对齐 go</td></tr>
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
    <tr><td><code>in(arr/map, value)</code></td><td>查询 value 是否在 arr/map 的 key 中</td></tr>
    <tr><td><code>sort</code></td><td>排序，只支持基本类型数组</td></tr>
  </tbody>
</table>


### 字符串

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:30%">函数</th>
      <th style="width:70%">作用</th>
    </tr>
  </thead>
  <tbody>
    <tr><td><code>index</code></td><td>索引</td></tr>
    <tr><td><code>hasPrefix</code></td><td>是否有 prefix</td></tr>
    <tr><td><code>hasSuffix</code></td><td>是否有 suffix</td></tr>
    <tr><td><code>trim</code></td><td>去除两端空白</td></tr>
    <tr><td><code>trimPrefix</code></td><td>去除 prefix</td></tr>
    <tr><td><code>trimSuffix</code></td><td>去除 suffix</td></tr>
    <tr><td><code>trimSpace</code></td><td>去除空白</td></tr>
    <tr><td><code>trimLeft</code></td><td>去除左侧，cutset 匹配</td></tr>
    <tr><td><code>trimRight</code></td><td>去除右侧，cutset 匹配</td></tr>
    <tr><td><code>toLower</code></td><td>转小写</td></tr>
    <tr><td><code>toUpper</code></td><td>转大写</td></tr>
    <tr><td><code>split</code></td><td>拆分</td></tr>
    <tr><td><code>join</code></td><td>合并</td></tr>
    <tr><td><code>atoi</code></td><td>string->int64</td></tr>
    <tr><td><code>itoa</code></td><td>int64->string</td></tr>
  </tbody>
</table>

### 时间

<table style="width:100%">
  <thead>
    <tr>
      <th style="width:30%">函数</th>
      <th style="width:70%">作用</th>
    </tr>
  </thead>
  <tbody>
    <tr><td><code>now</code></td><td>获取当前时间</td></tr>
    <tr><td><code>parseTime</code></td><td>解析字符串为时间</td></tr>
    <tr><td><code>parseDuration</code></td><td>解析字符串为 duration</td></tr>
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
    <tr><td><code>toJson</code></td><td> 序列化为 json 字符串</td></tr>
    <tr><td><code>fromJson</code></td><td> 反序列化为 map[string][any]</td></tr>
    <tr><td><code>unmarshalJson</code></td><td> 对齐 json.Unmarhsal</td></tr>
  </tbody>
</table>


# 性能

```text
goos: darwinench=BenchmarkCompile -benchtime=5s
goarch: arm64
pkg: github.com/ycl2018/gs
cpu: Apple M3 Pro
BenchmarkCompileMapExpr-12              41453820               138.2 ns/op
BenchmarkCompileMapGS-12                28542619               214.0 ns/op
BenchmarkCompileStructExpr-12            8556379               750.5 ns/op
BenchmarkCompileStructGS-12             10276479               586.5 ns/op
PASS
ok      github.com/ycl2018/gs   24.698s
```

# 关于编译优化

- 直接常量表达式优化

```javascript
a = (1*2+3*3) + 100 // 编译后直接为 111
b = "hello" + " world" // 编译后直接为 “hello world”
```

- slice，map字面量

```javascript
arr = ["111","222","333"] // 编译为伪常量为[]string{"111", "222", "333"}，运行时拷贝一份
dict = {"1": 1,"2": 2, "3": 3} // 编译为伪常量为map[string]int{"1": 1, "2": 2, "3": 3}，运行时拷贝一份
```

- 在 in 函数中，slice常量会被编译为map，且运行时不会拷贝

```javascript
has = in(["111","222","333"],"111") // 编译为伪常量为map[string]struct{}{"111": {}, "222": {}, "333": {}}，且运行时不会拷贝
```
