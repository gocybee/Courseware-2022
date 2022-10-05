# gocybee第一节课

代码地址：https://github.com/gocybee/Courseware-2022/

> 以后每次更新代码之后同学们可以直接通过 git 进行更新

## 前言

在这节课的学习中，我们将会学习到Golang的基础语法，从"Hello, World"开始，到函数结束

当然不会那么枯燥的一直讲语法，在课程中后段会让大家上手人生中的第一个项目，并正式开启Golang生涯

## 基础语法

### 你好，gocybee!

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, gocybee!")
}
```

### var

#### 变量

##### 变量类型
变量（Variable）的功能是存储数据。不同的变量保存的数据类型可能会不一样。经过半个多世纪的发展，编程语言已经基本形成了一套固定的类型，常见变量的数据类型有：整型、浮点型、布尔型等。

Go语言中的每一个变量都有自己的类型，并且变量必须经过声明才能开始使用。

##### 变量声明
Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且Go语言的变量声明后必须使用。

##### 变量的初始化
Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值，例如： 整型和浮点型变量的默认值为0。 字符串变量的默认值为空字符串。 布尔型变量默认为false。 切片、函数、指针变量的默认为nil。

```go
var 变量名 类型 = 表达式

var a = "initial" // 类型推导，不指定类型自动判断

var b, c int = 1, 2 // 一次初始化多个变量

var d = true

var e float64 // 普通声明未赋值

f := float32(e) // 短声明

g := a + "apple"
fmt.Println(a, b, c, d, e, f) // initial 1 2 true 0 0
fmt.Println(g)                // initialapple
```

#### 常量

相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。 常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。

```go
const s string = "constant"
const h = 500000000
const i = 3e20 / h
fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
```

### for

```go
for init statement; condition expression; post statement {
    //这里是中间循环体
}
```

`statement`是单次表达式，循环开始时会执行一次这里

`expression`是条件表达式，即循环条件，只要满足循环条件就会执行中间循环体。

`statement`是末尾循环体，每次执行完一遍中间循环体之后会执行一次末尾循环体

执行末尾循环体后将再次进行条件判断，若条件还成立，则继续重复上述循环，当条件不成立时则跳出当下for循环

```go
package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println("loop")
		break // 跳出循环
	}
	
	// 打印7、8
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
			// 当n模2为0时不打印，进到下一次的循环
		}
		fmt.Println(n)
	}
	// 直到i>3
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}
```

### if

```go
if 条件表达式 {
	//当条件表达式结果为true时，执行此处代码   
}

if 条件表达式 {
    //当条件表达式结果为true时，执行此处代码  
} else {
    //当条件表达式结果为false时，执行此处代码  
}
```

```go
package main

import "fmt"

func main() {
	// 条件表达式为false，打印出"7 is odd"
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 条件表达式为ture，打印出"8 is divisible by 4"
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 短声明，效果等效于
	//num := 9
	//if num < 0{
	//	...
	//}
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
```

### switch

当分支过多的时候，使用if-else语句会降低代码的可阅读性，这个时候，我们就可以考虑使用switch语句

- switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
- switch 语句在默认情况下 case 相当于自带 break 语句，匹配一种情况成功之后就不会执行其它的case，这一点和 c/c++ 不同
- 如果我们希望在匹配一条 case 之后，继续执行后面的 case ，可以使用 fallthrough

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	a := 2
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		// 在此打印"two"并跳出
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}

	t := time.Now()
	switch {
	// 根据现在的时间判断是上午还是下午
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
```

### array

数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。

```go
package main

import "fmt"

func main() {
	// 声明了长度为5的数组，数组中的每一个元素都是int类型
	var a [5]int
	// 给数组a的第4位元素赋值为100
	a[4] = 100
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))

	// 在给数组声明的同时赋值
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 声明二位数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

```

### slice

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

```go
var s []int
```
类似与声明一个数组，只不过不用填写它的长度

值得一提的是，切片必须先初始化才能使用！

```go
package main

import "fmt"

func main() {

	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])   // c
	fmt.Println("len:", len(s)) // 3

	// 使用append在尾部添加元素
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s) // [a b c d e f]

	c := make([]string, len(s))
	// 将s复制给c
	copy(c, s)
	fmt.Println(c) // [a b c d e f]
	
	fmt.Println(s[2:5]) // [c d e]
	fmt.Println(s[:5])  // [a b c d e]
	fmt.Println(s[2:])  // [c d e f]

	good := []string{"g", "o", "o", "d"}
	fmt.Println(good) // [g o o d]
}
```

### func

函数是指一段可以直接被另一段程序或代码引用的程序或代码，

一个较大的程序一般应分为若干个程序块，每一个模块用来实现一个特定的功能。

```go
package main

import "fmt"

func add(a int, b int) int {
	// 返回a+b的和
	return a + b
}

// 若类型相同，允许这样写
func add2(a, b int) int {
	return a + b
}

func main() {
	res := add(1, 2)
	fmt.Println(res) // 3
}
```

## 年轻人的第一个GoProject

### 猜数游戏

#### v1

由于是猜数嘛，我们肯定需要先有可以猜的数。所以我们需要生成一个随机数

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
}

```

#### v2

有的同学可能会发现，虽然是产生了随机数，但是每一次生成的数字都是一样的。这是因为我们生成随机数的种子没有改变，需要让种子发生变化才能使每次生成的随机数不同

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	// 使用一直在不断变化的时间作为我们的种子
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
}
```

#### v3

实现输入我们猜的数字

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	var guess int
	_, err := fmt.Scanf("%d", &guess)
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer value")
		return
	}
	fmt.Println("You guess is", guess)
}
```

#### v4

实现完整的猜数逻辑

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	var guess int
	_, err := fmt.Scanf("%d", &guess)
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer value")
		return
	}
	fmt.Println("You guess is", guess)
	if guess > secretNumber {
		fmt.Println("Your guess is bigger than the secret number. Please try again")
	} else if guess < secretNumber {
		fmt.Println("Your guess is smaller than the secret number. Please try again")
	} else {
		fmt.Println("Correct, you Legend!")
	}
}
```

#### v5

加入for循环，项目实现完成

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	for {
		var guess int
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
```

## 作业

袁鑫浩 本名袁神 因原神游戏大火害怕律师函警告而改名 现为勤奋蜂蜂主 喜欢用go语言下围棋 是阿尔法狗的开发者之一 曾登上亚洲一百张最帅面孔排行榜

### LV0 

将袁神讲的知识巩固一遍并自己敲一遍代码

### LV1 

众所周知 袁 与 ⚪ 类似，你为了得到袁神的青睐需要打印跟⚪有关的数，

所以我们定义一个自然数，若它本身含有的0和它的各个数位之和含有的0最多，那么我们称之为⚪数

如100-999的⚪数为505，1000-9999的⚪数为5005。

你所做的事情就是在go支持的数据类型（int或float64）的数据范围内打印尽可能多的⚪数

### LV2 

众所周知，在袁神星球上有许多 UFO。这些 UFO 时常来收集地球上的忠诚的袁神支持者。不幸的是，他们的飞碟每次出行都只能带上一组支持者。因此，他们要用一种聪明的方案让这些小组提前知道谁会被彗星带走。他们为每个彗星起了一个名字，通过这些名字来决定这个小组是不是被带走的那个特定的小组（你认为是谁给这些彗星取的名字呢？）。关于如何搭配的细节会在下面告诉你；你的任务是写一个程序，通过小组名和彗星名来决定这个小组是否能被那颗彗星后面的 UFO 带走。

小组名和彗星名都以下列方式转换成一个数字：最终的数字就是名字中所有字母的积，其中 A 是 1，Z 是 26。例如，YUAN 小组就是 $25\times21\times1\times14=7350$。如果小组的数字 mod 47 等于彗星的数字 mod 47,你就得告诉这个小组需要准备好被带走！（记住“a mod b”是 a 除以 b 的余数，例如 34 mod 10 等于 4）

写出一个程序，读入彗星名和小组名并算出用上面的方案能否将两个名字搭配起来，如果能搭配，就输出 `GO`，否则输出 `STAY`。小组名和彗星名均是没有空格或标点的一串大写字母（不超过 66 个字母）。

#### 样例

##### 输入1

```
COMETQ
HVNGAT
```

##### 输出1

```
GO
```

##### 输入2

```
ABSTAR
YUAN
```

##### 输出2

```
STAY
```

### LV3 
按照自己的想法将猜数游戏升级到v6并说明你的升级内容
>可以是任何的升级，请发挥你们的想象力。如果可以甚至可以将你的猜数游戏部署到网页上(如果真的部署起了，第一个人奖励一杯奶茶哦)



作业完成后将GitHub地址在 **HappyOJ** 提交，

（如果是苹果手机请提交到邮箱 gocybee@gocybee.team，题目格式：第一次作业-20xxxxxx-袁神-lv3）

**截止时间**：下一次上课之前
