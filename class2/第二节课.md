# Go组第二节课



## 🌸前言

 第二节课 我们要重点讲 结构体 指针 **接口** 等



## 🚀结构体

Go语言中没有像`C++` `Java`等 “类”的概念，也不支持“类”的继承,多态等面向对象的特征(不是纯`oop`的语言)。但是Go语言可以通过 结构体和接口的使用实现比向对象具有更高的扩展性和灵活性的功能。



### 类型别名和自定义类型



#### 自定义类型

 除了Go语言基本的数据类型string、整型、浮点型、布尔等数据类型, 我们可以通过type关键词 自定义类型。

```go
type wxgg int
```

 通过`type `定义的`wxgg`就是一种**新的类型**，它只是具有int的特性,但**不是int类型**。



#### 类型别名 

类型别名是一个类型的别称，本质上还是同一个类型。

```go
type wxgg=int
```
常见的`rune`,`byte`,`any`就是**类型别名**。
```go
// byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
// used, by convention, to distinguish byte values from 8-bit unsigned
// integer values.
type byte = uint8

// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.
type rune = int32

// any is an alias for interface{} and is equivalent to interface{} in all ways.
type any = interface{}
```



#### 区别

其实最主要的区别是 自定义类型是新类型，类型别名依然是旧类型，只不过有了一个别名 新的名称 。

~~**ps**：袁神 也是yxh哥哥  wxgg和袁神一样强 但他们是不同两位人噢~~

```go
package main

import (
	"fmt"
	"reflect"
)

type wxjj int

type wxgg = int

func main() {
	var a wxjj

	var b wxgg

	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)

	rfTypeOf(a)
	rfTypeOf(b)

	TypeOf(a)
	TypeOf(b)
}

func rfTypeOf(data interface{}) {
	of := reflect.TypeOf(data)
	fmt.Println(of)
}

func TypeOf(data interface{}) {
	switch data.(type) {
	case wxgg:
		fmt.Println("Type is int")
	case nil:
		fmt.Println("Type is nil")
	default:
		fmt.Println("Type Not Found")
	}
}
```



### 结构体



#### 结构体的定义


使用`type`和`struct`关键字来定义结构体：

```go
    type 类型名 struct {
        字段名 字段类型
        字段名 字段类型
    }
```

#### 结构体实例

 只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
结合前面所说 **结构体本身也是一种类型**，我们可以像声明内置类型一样使用`var`关键字声明结构体类型 或者是**:=**语法糖（匿名结构体）。

**匿名结构体**

> 在以后的 学习过程中 匿名结构体也经常出现在全局变量管理 模板渲染 数据测试 等 

```go
//var
var WxJJ struct{Name string; Age int} ; WxJJ.Name = "wxjj" ; WxJJ.Age = 18

//推荐
WxGG := struct {
		Name string
		Age int
	}{
		"wxgg",
		18,
	}
```



```go
package main

import (
	"fmt"
)

type WxGG struct {
	Name string
	Age  int
}

func main() {

	//最常见的方式
	a := WxGG{
		Name: "wxgg1",
		Age:  18,
	}

	var b WxGG
	b.Name = "wxgg2"
	b.Age = 18

	//var WxJJ struct {
	//	Name string
	//	Age  int
	//}
	//
	//WxJJ.Name = "wxjj1"
	//WxJJ.Age = 18

	WxJJ := struct {
		Name string
		Age  int
	}{
		"wxjj2",
		18,
	}

	//  类比
	//	type  yxh int
	//	god :=yxh(55)
	
	
	c:=NewWxGG("wxgg tql",18)
	
	
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", WxJJ)
	fmt.Printf("%#v\n", c)

}

//构造函数
func NewWxGG(name string, age int) *WxGG {
	return &WxGG{
		Name: name,
		Age:  age,
	}
}
```



#### 结构体的访问

  我们通过上述 代码可以轻而易举的知道 给以直接拿到整 个结构体变量 但是我们应该怎么获取 结构体的成员呢？

  我们可以通过 `.`  操作来进行访问。

```go
  Wxjj := struct {
		Name string
		Age int
	}{
		"wxjj2",
		18,
	}

fmt.Println(wxjj)
fmt.Println(wxjj.Name)
fmt.Println(wxjj.Age)
```

但 并不是 所有的成员函数 都可以访问

在Go中 没有public、protected、private等访问控制修饰符，它是通过字母大小写来控制可见性的，如果定义的常量、变量、类型、接口、结构、函数等的名称是大写字母开头表示能被其它包访问或调用（相当于public），非大写开头就只能在包内使用（相当于private，变量或常量也可以**下划线**开头）



#### 嵌套结构体

我们把一个结构体中可以嵌套包含另一个结构体或结构体指针 叫做 **嵌套结构体**或者是**结构体内嵌**

```go
type People struct {
    Name string
    Age  int
}

type Bankend struct {
    Name string
    member People 
}

```



#### 方法和接收者

> Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。  方法的定义格式如下：

```go
  func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
        函数体
    }
```



#### 举一个🌰

就拿 上面的wxGG来举🌰 

```go
type WxGG struct {
	Name  string
	Age   int
	Books []Book
}

type Book struct {
	Name string
}

func (w WxGG) PrintName() {
	fmt.Println(w.Name)
}
func (w WxGG) PrintAge() {
	fmt.Println(w.Age)
}

func (w WxGG) PrintBook() {
	fmt.Println(w.Books)
}

func (b Book) PrintBookName() {
	fmt.Println(b.Name)
}
```





##  🧷指针

> 区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是**安全指针**。
>
> 要搞明白Go语言中的指针需要先知道3个概念：`指针地址`、`指针类型`和`指针取值`。



###  Go语言中的指针

> go语言中的函数传参都是值传递，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。传递数据使用指针，而无须拷贝数据。类型指针不能进行偏移和运算。Go语言中的指针操作非常简单，只需要记住两个符号：`&`（取地址）和`*`（根据地址取值）。
>
> 具体可以 看看[go语言中的函数传参都是值传递](https://juejin.cn/post/6892678846223974407) 



### 指针地址和指针类型

​    每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用&字符放在变量前面对变量进行“取地址”操作。 Go语言中的值类型`（int、float、bool、string、array、struct）`都有对应的指针类型，如：`*int、*int64、*string`等



我们可以通过 **&**来获取 **地址**

 ```go
  ptr := &v    // v的类型为
 ```



举个🌰 

```go
func main() {
    a := 10
    b := &a
    fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
    fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
    fmt.Println(&b)                    // 0xc00000e018
}
```



### 🕳空指针

> - 当一个指针被定义后没有分配到任何变量时，它的值为 nil
> - 空指针的判断

 ```go
package main

import "fmt"

func main() {
    var p *string
    fmt.Println(p)
    fmt.Printf("p的值是%v\n", p)
    if p != nil {
        fmt.Println("非空")
    } else {
        fmt.Println("空值")
    }
}
 ```





### 🌐new和make



在这之前 我们先来看一个 🌰 大家 不忙猜测 一下 运行程序会出现 什么问题 

```go
func main() {
	var a *string
	*a = "无香的一刀"
	fmt.Println(*a)

	var b map[string]string
	b["袁神"] = "YYDS"
	fmt.Println(b)
}
```





#### 🆕new

> new是一个内置的函数，它的函数签名如下：

```go
func new(Type) *Type
```

对于之前的🌰 

```go
    var a *string
    a=new(string)
	*a = "无香的一刀"
	fmt.Println(*a)
```



#### 🤖make

> make也是用于内存分配的，区别于new，它只用于`slice`、`map`以及`chan`的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。make函数的函数签名如下：



```go
func make(t Type, size ...IntegerType) Type
```



对于之前的🌰 

```go
	var b map[string]string
	b = make(map[string]string)
	b["袁神"] = "YYDS"
	fmt.Println(b)
```



#### 🤔思考一下

**New 和Make 有那些区别 **

大家 可以课后 去查查



## 🎉接口

> 接口（interface) 可以定义一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节



###  接口类型

>在Go语言中接口（interface）是一种类型，一种抽象的类型。
>
>interface是**方法**（1.18以后 加入了泛型 准确说 是类型 因为 方法也是一种类型）的集合



### 为什么要用接口

在讲解之前我们来 看一下以下🌰 

````go
type Cat struct{}

func (c Cat) Say() string { return "喵喵喵" }

type Dog struct{}

func (d Dog) Say() string { return "汪汪汪" }

func main() {
    c := Cat{}
    fmt.Println("猫:", c.Say())
    d := Dog{}
    fmt.Println("狗:", d.Say())
}
````

 上面的代码中定义了🐶and🐱，然后它们都会叫，你会发现有重复的代码，如果我们后续再加上其他动物的话，我们的代码还会一直重复下去，为了代码 的扩展性，那我们能不能把这动物 都 归为 **"会叫的动物"**

除了 上面 这些还有很多🌰 

比如一个网上商城可能使用支付宝、微信、银联等方式去在线支付，我们可以把他当成 “支付方式” 来一起处理。



### 接口的定义

```go
type 接口类型名 interface{
        方法名1( 参数列表1 ) 返回值列表1
        方法名2( 参数列表2 ) 返回值列表2
    }
```



### 🔍接口的实现

我们 可以 根据上面的🌰  实现一个 `sayer` 接口

```go
type Sayer interface {
	Say()
}

type dog struct {
}

type cat struct {
}

func (d dog) Say() {
	fmt.Println("汪汪汪")
}

func (c cat) Say() {
	fmt.Println("喵喵喵")
}
func main() {
	var x Sayer
	a := cat{}
	b := dog{}
	x = a
	x.Say()
	x = b
	x.Say()
}
```



### 🕳空接口

> 空接口是指没有定义**任何方法**的接口。因此任何类型都实现了空接口。空接口类型的变量可以存储任意类型的变量
>
> 使用空接口实现可以接收任意类型的函数参数。



```go
   var x interface{}
    s := "YuanShen"
    x = s

func NilInterface(x  interface{}){
}
```





### 类型断言

 上面我们提到了 空接口可以存储任意类型，但是我们怎么才能知道他的类型呢？？？

 欸嘿 其实 我们在讲**类型别名**的时候就用到了。 

````go
func TypeOf(data interface{}) {
	switch data.(type) {
	case wxgg:
		fmt.Println("Type is int")
	case nil:
		fmt.Println("Type is nil")
	default:
		fmt.Println("Type Not Found")
	}
}

func main() {
    var x interface{}
    s := "YuanShen"
    x = s
    v, ok := x.(string)
    if ok {
        fmt.Println(v)
    } else {
        fmt.Println("类型断言失败")
    }
}
````



## 📖拓展阅读和学习

https://juejin.cn/post/6892678846223974407

https://www.zhihu.com/question/318138275

https://www.bilibili.com/video/BV1iZ4y1T7zF/

## 📚作业



## Lv0

复习，将上课所讲的代码都敲一遍。





## Lv1



打开 知乎 我们可以 看到很多 question 尽可能的 完善一个 question 的结构体

![](https://s3.bmp.ovh/imgs/2022/10/13/3bb63e8c83b24b00.png)

类似这种 尽量补充完整

````go
type struct Question struct {
	Id        int64
	content   string
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
}
````

## Lv2

**看看课外阅读 尽量理解**













### make

在 Go 语言中，内置函数 `make` 仅支持 `slice`、`map`、`channel` 三种数据类型的内存创建，**其返回值是所创建类型的本身，而不是新的指针引用**。





### new

在 Go 语言中，内置函数 `new` 可以对类型进行内存创建和初始化。**其返回值是所创建类型的指针引用**，与 `make` 函数在实质细节上存在区别。





本质上在于 make 函数在初始化时，会初始化 slice、chan、map 类型的内部数据结构，new 函数并不会。

在 map 类型中，合理的长度（len）和容量（cap）可以提高效率和减少开销。

更进一步的区别：



**make 函数**

能够分配并初始化类型所需的内存空间和结构，**返回引用类型的本身**。

具有使用范围的局限性，仅支持 channel、map、slice 三种类型。

具有独特的优势，make 函数会对三种类型的内部数据结构（长度、容量等）赋值。

**new 函数：**

能够分配类型所需的内存空间，返回指针引用（指向内存的指针）。

可被替代，能够通过字面值快速初始化。


