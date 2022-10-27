# 第四节课 map与标准库函数

## map

### 哈希表（散列表）

![asd](https://yourbasic.org/algorithms/hash-table.png)

在编程实现中，常常面临着两个问题：存储和查找，存储和查找的效率往往决定了整个程序的效率。

- 数组的特点是：寻址容易，插入和删除困难；
- 而链表的特点是：寻址困难，插入和删除容易。

那么我们能不能综合两者的特性，做出一种寻址容易，插入删除也容易的数据结构？答案是肯定的，这就是哈希表

![](https://muku-store.com/user_data/img/item/mobel/wch04.jpg)

脑补下，你在家里忘记了指甲刀放在哪里，通常要在你家所有抽屉中顺序寻找，直到找到，最差情况下，有N个抽屉，你就要打开N个抽屉。这种存储方式叫数组，查找方法称为「遍历」。

脑补下，你是一个整理控，所有物品必须分门别类放入整理箱，再将整理箱编号，比如1号放入针线，2号放入证件，3号放入细软。这种存储和查找方式称为「哈希」，如果这个时候要查找护照，你不许要再翻所有抽屉，直接可在2号整理箱中获取，通常只用一次查找即可，如何编号整理箱，称为哈希算法。

同样是查找，差距怎么那么大捏~，假设我们有100亿条数据记录，那差距就变得明显，遍历需要查找最多100亿次，最少1次，哈希只需1次。

让我们正式介绍哈希和哈希算法，哈希也称散列，哈希表是一种与数组、链表等不同的数据结构，与他们需要不断的遍历比较来查找的办法，哈希表设计了一个映射关系f(key)= address，根据key来计算存储地址address，这样可以1次查找，f既是存储数据过程中用来指引数据存储到什么位置的函数，也是将来查找这个位置的算法，叫做哈希算法。

哈希表hashtable(key，value) 就是把Key通过一个固定的算法函数既所谓的哈希函数转换成一个整型数字，然后就将该数字对数组长度进行取余，取余结果就当作数组的下标，将value存储在以该数字为下标的数组空间里。（或者：把任意长度的输入（又叫做预映射， pre-image），通过散列算法，变换成固定长度的输出，该输出就是散列值。）

这种转换是一种压缩映射，也就是，散列值的空间通常远小于输入的空间，不同的输入可能会散列成相同的输出，而不可能从散列值来唯一的确定输入值。简单的说就是一种将任意长度的消息压缩到某一固定长度的消息摘要的函数。

而当使用哈希表进行查询的时候，就是再次使用哈希函数将key转换为对应的数组下标，并定位到该空间获取value，如此一来，就可以充分利用到数组的定位性能进行数据定位。

### map的使用

![](https://technobeans.com/wp-content/uploads/2019/02/golang-maps.png)

在 Go 语言中，`map` 是散列表的引用，`map` 的类型是 `map[K]V`，其中 K 和 V 是字典的键和值对应的数据类型。`map` 中所有的键都拥有相同的数据类型，同时所有的值也都拥有相同的数据类型，但是键的类型和值的类型不一定相同。键的类型 K，必须是可以通过操作符 == 来进行比较的数据类型，所以 `map` 可以检测某一个键是否存在。

#### 初始化

内置函数 `make` 可以用来创建一个 map：

```go
ageMap := make(map[string]int)
```

也可以使用 map 的字面量来新建一个带初始化键值对元素的字典：

```go
ageMap := map[string]int{
	"yuanshen": 31,
	"wx": 22,
}
```

这个等价于：

```go
ageMap := make(map[string]int)
ageMap["yuanshen"]=31
ageMap["wx"]=22
```

因此，新的空 `map` 的另外一种表达式是： `map[string]int{}`

#### 访问

`map` 的元素访问是通过 `key` 的形式访问：

```3
ageMap["hhz"]=18
fmt.Println(ageMap["hhz"]) // 32
```

但是 `map` 元素不是一个变量，不可以获取它的地址，比如这样是不对的：

```go
_ = &ageMap["wx"] // 编译错误，无法获取 map 元素的地址
```



当从 `map` 中读取一个不存在的 `key` 的时候，返回零值，有时候很麻烦，所以可以用 go 语言的另一个用法

```go
type Node struct {
    Next  *Node
    Value interface{}
}

var first *Node

func main(){
	visited := make(map[*Node]bool)
    for n := first; n != nil; n = n.Next {
        if visited[n] {
            fmt.Println("cycle detected")
            break
        }
        visited[n] = true
        fmt.Println(n.Value)
    }
}
```

#### 删除

可以使用内置函数 `delete` 来从字典中根据键移除一个元素：

```go
delete(ageMap, "jjz")
```

即使键不在 `map` 中，上面的操作也是安全的。`map` 使用给定的键来查找元素，如果对应的元素不存在，就返回值类型的零值。例如，下面的代码同样可以工作，尽管`"jjz"`还不是 `map` 的键，因为 `ageMap["jjz"]`的值是 0。

#### 赋值

```go
ageMap["wx"] = ageMap["wx"] + 1
```

快捷赋值方式（如`x+=y`和`x++`）对 map 中的元素同样使用，所以上面的代码还可以写成：

```go
ageMap["wx"] += 1
```

或者更简洁的：

```go
ageMap["wx"]++
```

#### 遍历

可以使用 `for` 循环（结合 `range` 关键字）来遍历 `map` 中所有的键和对应的值，就像遍历 `slice` 一样。循环语句的连续迭代将会使得变量 `name` 和 `age` 被赋予 `map` 中的下一个键值对。

```go
for name, age := range ageMap{
	fmt.Printf("%s\t%s\n", name, age)
}
```

#### 线程安全

##### 例子

- slice

    我们使用多个 `goroutine` 对类型为 `slice` 的变量进行操作，看看结果会变的怎么样。

    ```go
    func main() {
         var s []string
         for i := 0; i < 9999; i++ {
              go func() {
                   s = append(s, "袁神又在玩原了！")
              }()
         }
    
         fmt.Printf("玩 %d 次原神", len(s))
    }
    ```

- map

    同样针对 `map` 也如法炮制一下。重复针对类型为 `map` 的变量进行写入。

    ```go
    func main() {
    	s := make(map[string]string)
    	for i := 0; i < 99; i++ {
    		go func() {
    			s["袁神"] = "原"
    		}()
    	}
    
    	fmt.Printf("玩了 %d 次原神", len(s))
    }
    ```

##### 如何保证并发安全

- 加锁

    ```go
    var counter = struct{
        sync.RWMutex
        m map[string]int
    }{m: make(map[string]int)}
    ```

    这条语句声明了一个变量，它是一个匿名结构（struct）体，包含一个原生和一个嵌入读写锁 `sync.RWMutex`。

    要想从变量中中读出数据，则调用读锁：

    ```go
    counter.RLock()
    n := counter.m["原"]
    counter.RUnlock()
    fmt.Println("原:", n)
    ```

    要往变量中写数据，则调用写锁：

    ```go
    counter.Lock()
    counter.m["原"]++
    counter.Unlock()
    ```

    这就是一个最常见的 Map 支持并发读写的方式了。

- 使用 `sync.Map`

    Go 语言的 `sync.Map` 支持并发读写 map，采取了 “空间换时间” 的机制，冗余了两个数据结构，分别是：read 和 dirty，减少加锁对性能的影响：

    ```go
    type Map struct {
    	mu Mutex
    	read atomic.Value // readOnly
    	dirty map[interface{}]*entry
    	misses int
    }
    ```

    其是专门为 `append-only` 场景设计的，也就是适合读多写少的场景。这是他的优点之一。

    若出现写多/并发多的场景，会导致 read map 缓存失效，需要加锁，冲突变多，性能急剧下降。这是他的重大缺点。

    提供了以下常用方法：

    ```go
    func (m *Map) Delete(key interface{})
    func (m *Map) Load(key interface{}) (value interface{}, ok bool)
    func (m *Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool)
    func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
    func (m *Map) Range(f func(key, value interface{}) bool)
    func (m *Map) Store(key, value interface{})
    ```

    - `Delete`：删除某一个键的值。
    - `Load`：返回存储在 map 中的键的值，如果没有值，则返回 nil。ok 结果表示是否在 map 中找到了值。
    - `LoadAndDelete`：删除一个键的值，如果有的话返回之前的值。
    - `LoadOrStore`：如果存在的话，则返回键的现有值。否则，它存储并返回给定的值。如果值被加载，加载的结果为 true，如果被存储，则为 false。
    - `Range`：递归调用，对 map 中存在的每个键和值依次调用闭包函数 `f`。如果 `f` 返回 false 就停止迭代。
    - `Store`：存储并设置一个键的值。

    实际运行例子如下：

    ```go
    var m sync.Map
    
    func main() {
     //写入
     data := []string{"jjz", "hhz", "rrz"}
     for i := 0; i < 4; i++ {
      go func(i int) {
       m.Store(i, data[i])
      }(i)
     }
     time.Sleep(time.Second)
    
     //读取
     v, ok := m.Load(0)
     fmt.Printf("Load: %v, %v\n", v, ok)
    
     //删除
     m.Delete(1)
    
     //读或写
     v, ok = m.LoadOrStore(1, "xxz")
     fmt.Printf("LoadOrStore: %v, %v\n", v, ok)
    
     //遍历
     m.Range(func(key, value interface{}) bool {
      fmt.Printf("Range: %v, %v\n", key, value)
      return true
     })
    }
    ```

## 标准库函数

### fmt

![](https://i.ytimg.com/vi/GQ880MlHBBE/maxresdefault.jpg)

`fmt`是 `Go `中最常用的包之一，主要实现了格式化I/O(输入/输出)。

#### 输出

输出部分包括三个系列和一个独立的函数：`Print系列`，`Fprint系列`，`Sprint系列`，以及`Errorf()`。这些函数的使用场景如下：

- 向终端输出一些信息的时候，使用`Print系列`。（使用程度：频繁）
- 将信息写入文件中时，使用`Fprint系列`。（使用程度：一般）
- 在程序中获取格式化字符串中时，使用`Sprint系列`。（使用程度：一般）
- 在程序中获取包含格式化字符串的错误时，使用`Errorf()`。（使用程度：一般）

##### Print系列

Print系列中包含三个重要的函数：`Print()`，`Println()`，`Printf()`。

```go
func Print(a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
```

- Print()

    ```go
    var a string = "test Print"
    fmt.Print(a)
    
    fmt.Println()
    
    fmt.Print(a,a,a)
    ```

- Println()

    ```go
    var b string = "test Println"
    fmt.Println(b)
    
    fmt.Println()
    
    fmt.Println(b,b,b)
    ```

- Printf()

    ```go
    var date string = "2022-10-29"
    fmt.Printf("今天的日期是： %s", date)
    ```

##### Sprint系列

`Sprint系列`函数会把传入的参数生成并返回一个字符串。
`Sprint系列`包含三个函数：`Sprint()`,`Sprintf()`,`Sprintln()`。

```go
func Sprint(a ...interface{}) string
func Sprintln(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
```

`Sprint系列`与`Print系列`的区别在于输出的对象不同，`Sprint系列`的输出对象为字符串。利用time包，举一个输出日期和时间的例子：

```go
year, month, day := time.Now().Date()
hour, min, sec := time.Now().Clock()

//将格式化字符串写入变量s1中
s1 := fmt.Sprintf("今天的日期的是：%d年%d月%d日，现在的时间是：%d:%d:%d\n",year,month,day,hour,min,sec)
fmt.Println(s1)
```

##### Fprint系列

`Fprint系列`包含三个函数：`Fprint()`,`Fprintf()`,`Fprintln()`。

```go
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
```

`Fprint系列`与`Print系列`相比多了一个`io.Writer`接口类型的参数`w`。`Fprint系列`函数会将内容输出到参数`w`中。只要参数类型实现了`io.Writer`接口，则都可以实现写入。

```go
// 打开test.txt文件
fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
if err != nil {
    fmt.Println("打开文件错误：err:", err)
    return
}
s := "test"
// 向打开的文件中写入格式化字符串
fmt.Fprintf(fileObj, "往文件中写如信息：%s", s)
```

实际上，`Print系列`其实就是通过封装了`Fprint系列`来实现的。`Print()`的源代码如下：

```go
func Print(a ...interface{}) (n int, err error) {
    return Fprint(os.Stdout, a...)
}
```

调用`Print()`后返回了一个`Fprint()`，而`os.Stdout`代表标准输出。因此我们可以用`Fprint()`来实现与`Print()`。

```go
var a string = "test Fprint"
fmt.Fprintln(os.Stdout, a)
```

##### Errorf函数

`Errorf()`根据format参数生成格式化字符串并返回一个包含该字符串的错误。

```go
func Errorf(format string, a ...interface{}) error
```

```go
error := "未知"
err := fmt.Errorf("这个错误类型为：%s", error)
fmt.Println(err)
```

它的底层是通过`error包`的`new()`中传入`Sprintf()`来实现的：

```go
func Errorf(format string, a ...interface{}) error {
    return errors.New(Sprintf(format, a...))
}
```

这也就是为什么`Errorf()`没有它的兄弟`Error()`和`Errorln()`的原因了。因为我们可以直接通过`errors.New()`来生成一个非格式化字符串的错误。

#### 输入

**Scan** 从标准输入中读取数据，并将数据用空白分割并解析后存入 a 提供的变量中（换行符会被当作空白处理），变量必须以指针传入。

当读到 EOF 或所有变量都填写完毕则停止扫描。

返回成功解析的参数数量。

```go
func Scan(a ...interface{}) (n int, err error)
```

**Scanln** 和 **Scan** 类似，只不过遇到换行符就停止扫描。

```go
func Scanln(a ...interface{}) (n int, err error)
```

Scanf 从标准输入中读取数据，并根据格式字符串 format 对数据进行解析，将解析结果存入参数 a 所提供的变量中，变量必须以指针传入。

输入端的换行符必须和 format 中的换行符相对应（如果格式字符串中有换行符，则输入端必须输入相应的换行符）。

占位符 %c 总是匹配下一个字符，包括空白，比如空格符、制表符、换行符。

返回成功解析的参数数量。

```go
func Scanf(format string, a ...interface{}) (n int, err error)
```

以下三个函数功能同上面三个函数，只不过从 r 中读取数据。

```go
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
```

以下三个函数功能同上面三个函数，只不过从 str 中读取数据。

```go
func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
```

实例:

```go
// 对于 Scan 而言，回车视为空白
func main() {
    a, b, c := "", 0, false
    fmt.Scan(&a, &b, &c)
    fmt.Println(a, b, c)
    // 在终端执行后，输入 abc 1 回车 true 回车
    // 结果 abc 1 true
}

// 对于 Scanln 而言，回车结束扫描
func main() {
    a, b, c := "", 0, false
    fmt.Scanln(&a, &b, &c)
    fmt.Println(a, b, c)
    // 在终端执行后，输入 abc 1 true 回车
    // 结果 abc 1 true
}

// 格式字符串可以指定宽度
func main() {
    a, b, c := "", 0, false
    fmt.Scanf("%4s%d%t", &a, &b, &c)
    fmt.Println(a, b, c)
    // 在终端执行后，输入 1234567true 回车
    // 结果 1234 567 true
}
```

#### 格式化符

Go 的 `fmt` 相关的函数支持一些占位符，最常见的是字符串占位符的 `%s`，整型占位符 `%d`，以及浮点型占位符 `%f`。现在让我们探究一些其他的占位符。

##### 通用

|  占位符  | 说明                                 | 示例                  | 输出                   |
| :------: | :----------------------------------- | :-------------------- | :--------------------- |
| **`%v`** | 相应值的默认格式                     | Printf("%v",person )  | {zhangsan}             |
|  `%+v`   | 类似%v，但输出结构体时会添加字段名式 | Printf("%+v",person ) | {Name:zhangsan}        |
|  `%#v`   | 相应值的Go语法表示                   | Printf("#v",person )  | main.Person={zhangsan} |
| **`%T`** | 相应值的类型的Go语法表示             | Printf("%T",person )  | main.Person            |
|   `%%`   | 字面上的百分号，并非值的占位符       | Printf("%%")          | %                      |

| 布尔占位符 |      说明       |       示例        | 输出 |
| :--------: | :-------------: | :---------------: | :--: |
|  **`%t`**  | 单词true或false | Printf("%t",true) | true |

`%v` 占位符将会打印出 Go 的值，如果此占位符以 `+` 作为前缀，将会打印出结构体的字段名，如果以 `#` 作为前缀，那么它会打印出结构体的字段名和类型。

```go
// Point is a 2D point
type Point struct {
    X int
    Y int
}

func main() {
    p := &Point{1, 2}
    fmt.Printf("%v %+v %#v \n", p, p, p)
}
```

##### 整数

| 占位符 | 说明                                       | 示例                |  输出  |
| :----: | :----------------------------------------- | :------------------ | :----: |
|  `%b`  | 二进制表示                                 | Printf("%b",5)      |  101   |
|  `%c`  | 该值对应的unicode码值                      | Printf("%c",0x4E2d) |   中   |
|   %d   | 十进制表示                                 | Printf("%d",0x12)   |   18   |
|  `%o`  | 八进制表示                                 | Printf("%o",10)     |   12   |
|  `%q`  | 单引号围绕的字符字面值，由Go语法安全的转译 | Printf("%q",0x4E2d) |  '中'  |
|  `%x`  | 十六进制表示，字母形式为小写a-f            | Printf("%x",13)     |   d    |
|  `%X`  | 十六进制表示，字母形式为大写A-F            | Printf("%X",13)     |   D    |
|  `%U`  | 表示为Unicode格式：U+1234，等价于"U+%04X"  | Printf("%U",0x4E2d) | U+4E2D |

##### 浮点数与复数的两个组分

|  占位符  |                             说明                             |        示例        |         输出         |
| :------: | :----------------------------------------------------------: | :----------------: | :------------------: |
|   `%b`   | 无小数部分、指数为二的幂的科学计数法，与strconv.FormatFloat的'b'转换格式一致。 | Printf("%b",10.20) | 5742089524897382p-49 |
|   `%e`   |                 科学计数法，如-1234.456e+78                  | Printf("%e",10.20) |     1.020000e+01     |
|   `%E`   |                 科学计数法，如-1234.456E+78                  | Printf("%E",10.20) |     1.020000E+01     |
| **`%f`** |              有小数部分但无指数部分，如123.456               | Printf("%f",10.20) |      10.200000       |
|   `%g`   |    根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）    | Printf("%g",10.20) |         10.2         |
|   `%G`   |    根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）    | Printf("%G",10.20) |      (10.2+2i)       |

##### 字符串和[]byte

|  占位符  |                  说明                  |             示例              |     输出     |
| :------: | :------------------------------------: | :---------------------------: | :----------: |
| **`%s`** |   输出字符串表示(string类型或[]byte)   | Printf("%s",[]byte("Go语言")) |    Go语言    |
|   `%q`   | 双引号围绕的字符串，由Go语法安全的转译 |     Printf("%q","Go语言")     |   "Go语言"   |
|   `%x`   |   十六进制，小写字母，每字节两个字符   |     Printf("%x","golang")     | 676f6c616e67 |
|   `%X`   |   十六进制，大写字母，每字节两个字符   |     Printf("%X","golang")     | 676F6C616E67 |

##### 指针

| 占位符 |         说明          |         示例         |     输出     |
| :----: | :-------------------: | :------------------: | :----------: |
|   %P   | 十六进制表示，前缀 0x | Printf("%p",&person) | 0xc0420341c0 |

##### 其他

| 占位符 | 说明                                                         | 示例                      | 输出           |
| :----: | :----------------------------------------------------------- | :------------------------ | :------------- |
|  `+`   | 总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义） | Printf("%+q","中文")      | "\u4e2d\u6587" |
|  `-`   | 在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）； |                           |                |
|  `#`   | 切换格式：八进制数前加0（%#o）                               | Printf("%#0",46)          |                |
|   #x   | 十六进制数前加0x（%#x）或0X（%#X）                           | Printf("%#x",46)          | 0x2e           |
|   #p   | 指针去掉前面的0x（%#p）；）                                  | fmt.Printf("%#p",&person) | c0420441b0     |
|   #q   | 对%q（%#q），如果strconv.CanBackquote返回真会输出反引号括起来的未转义字符串； | Printf("%#q",'中')        | '中'           |
|   #U   | 对%U（%#U），如果字符是可打印的，会在输出Unicode格式、空格、单引号括起来的go字面值； | Printf("%#U",'中')        | U+4E2D '中'    |
| `' '`  | (空格)为数值中省略的正负号流出空白(% d);                     | Printf("% d",16)          | 16             |
| `` x`  | 以十六进制(% x,% X)打印字符串或切片时，在字节之间用空格隔开  | Printf("% x","abc")       | 61 62 63       |
|  `0`   | 使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面   |                           |                |

##### 宽度

你可以为一个打印的数值指定宽段，比如：

```go
fmt.Printf("%10d\n", 353)  // will print "       353"
```

你还可以通过将宽度指定为 `*` 来将宽度当作 `Printf` 的参数，例如：

```go
fmt.Printf("%*d\n", 10, 353)  // will print "       353"
```

当你打印出数字列表而且希望它们能够靠右对齐时，这非常的有用。

```go
// alignSize return the required size for aligning all numbers in nums
func alignSize(nums []int) int {
    size := 0
    for _, n := range nums {
        if s := int(math.Log10(float64(n))) + 1; s > size {
            size = s
        }
    }

    return size
}

func main() {
    nums := []int{12, 237, 3878, 3}
    size := alignSize(nums)
    for i, n := range nums {
        fmt.Printf("%02d %*d\n", i, size, n)
    }
}
```

将会打印出：

```bash
00   12
01  237
02 3878
03    3
```

这使得我们更加容易比较数字。

##### 通过位置引用

如果你在一个格式化的字符串中多次引用一个变量，你可以使用 `%[n]`，其中 `n` 是你的参数索引（位置，从 1 开始）。

```go
fmt.Printf("The price of %[1]s was $%[2]d. $%[2]d! imagine that.\n", "carrot", 23)
```

这将会打印出：

```bash
The price of carrot was $23. $23! imagine that.
```

### time

time 包提供了时间显示和测量等所用的函数

#### 时间类型

Go 语言中使用`time.Time`类型表示时间。我们可以通过`time.Now`函数获取当前的时间对象，然后从时间对象中可以获取到年、月、日、时、分、秒等信息。

```go
// timeDemo 时间对象的年月日时分秒
func timeDemo() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Println(year, month, day, hour, minute, second)
```

#### 时间戳

时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）。

基于时间对象获取时间戳的示例代码如下：

```go
// timestampDemo 时间戳
func timestampDemo() {
	now := time.Now()        // 获取当前时间
	timestamp := now.Unix()  // 秒级时间戳
	milli := now.UnixMilli() // 毫秒时间戳 Go1.17+
	micro := now.UnixMicro() // 微秒时间戳 Go1.17+
	nano := now.UnixNano()   // 纳秒时间戳
	fmt.Println(timestamp, milli, micro, nano)
}
```

使用time.Unix()函数可以将时间戳转为时间格式。

```go
func timestampDemo2(timestamp int64) {
    timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
    fmt.Println(timeObj)
    year := timeObj.Year()     //年
    month := timeObj.Month()   //月
    day := timeObj.Day()       //日
    hour := timeObj.Hour()     //小时
    minute := timeObj.Minute() //分钟
    second := timeObj.Second() //秒
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
```

#### 时间间隔

time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。time.Duration表示一段时间间隔，可表示的最长时间段大约290年。

time包中定义的时间间隔类型的常量如下：

```go
const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)
```

例如：time.Duration表示1纳秒，time.Second表示1秒。

#### Location和time zone

Go 语言中使用 location 来映射具体的时区。时区（Time Zone）是根据世界各国家与地区不同的经度而划分的时间定义，全球共分为24个时区。中国差不多跨5个时区，但为了使用方便只用东八时区的标准时即北京时间为准。

下面的示例代码中使用`beijing`来表示东八区8小时的偏移量，其中time.FixedZonv 和 time.LoadLocation`这两个函数则是用来获取location信息。

```go
// timezoneDemo 时区示例
func timezoneDemo() {
	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
	// 对于很多其他国家需要考虑夏令时。
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location。
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 如果当前系统有时区数据库，则可以加载一个位置得到对应的时区
	// 例如，加载纽约所在的时区
	newYork, err := time.LoadLocation("America/New_York") // UTC-05:00
	if err != nil {
		fmt.Println("load America/New_York location failed", err)
		return
	}
	fmt.Println()
	// 加载上海所在的时区
	//shanghai, err := time.LoadLocation("Asia/Shanghai") // UTC+08:00
	// 加载东京所在的时区
	//tokyo, err := time.LoadLocation("Asia/Tokyo") // UTC+09:00

	// 创建时间对象需要指定位置。常用的位置是 time.Local（当地时间） 和 time.UTC（UTC时间）。
	//timeInLocal := time.Date(2009, 1, 1, 20, 0, 0, 0, time.Local)  // 系统本地时间
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	sameTimeInNewYork := time.Date(2009, 1, 1, 7, 0, 0, 0, newYork)

	// 北京时间（东八区）比UTC早8小时，所以上面两个时间看似差了8小时，但表示的是同一个时间
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)

	// 纽约（西五区）比UTC晚5小时，所以上面两个时间看似差了5小时，但表示的是同一个时间
	timesAreEqual = timeInUTC.Equal(sameTimeInNewYork)
	fmt.Println(timesAreEqual)
}
```

在日常编码过程中使用时间对象的时候一定要注意其时区信息。

#### 时间操作

##### Add

我们在日常的编码过程中可能会遇到要求时间+时间间隔的需求，Go语言的时间对象有提供Add方法如下：

```
    func (t Time) Add(d Duration) Time
```

举个例子，求一个小时之后的时间：

```go
func main() {
    now := time.Now()
    later := now.Add(time.Hour) // 当前时间加1小时后的时间
    fmt.Println(later)
}
```

##### Sub

求两个时间之间的差值：

```
    func (t Time) Sub(u Time) Duration
```

返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)。

##### Equal

```
    func (t Time) Equal(u Time) bool
```

判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。

##### Before

```
    func (t Time) Before(u Time) bool
```

如果t代表的时间点在u之前，返回真；否则返回假。

##### After

```
    func (t Time) After(u Time) bool
```

如果t代表的时间点在u之后，返回真；否则返回假。

##### Timer与Ticker

使用`time.Tick(时间间隔)`来设置定时器，定时器的本质上是一个通道（channel）。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    timer := time.NewTimer(10 * time.Second)

    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    for {
        select {
        case <-timer.C:
            fmt.Println("Done")
            return
        case t := <-ticker.C:
            fmt.Println("ticker at: ", t)
        }
    }
}
```

##### 时间格式化

`time.Format`函数能够将一个时间对象格式化输出为指定布局的文本表示形式，需要注意的是 Go 语言中时间格式化的布局不是常见的`Y-m-d H:M:S`，而是使用 `2006-01-02 15:04:05.000`（记忆口诀为2006 1 2 3 4 5）。

其中：

- 2006：年（Y）
- 01：月（m）
- 02：日（d）
- 15：时（H）
- 04：分（M）
- 05：秒（S）

**补充**

- 如果想格式化为12小时格式，需在格式化布局中添加`PM`。
- 小数部分想保留指定位数就写0，如果想省略末尾可能的0就写 9。

```go
// formatDemo 时间格式化
func formatDemo() {
	now := time.Now()
	// 格式化的模板为 2006-01-02 15:04:05

	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	fmt.Println(now.Format("2006/01/02 15:04:05.000")) // 2022/02/27 00:10:42.960
	// 小数点后写9，会省略末尾可能出现的0
	fmt.Println(now.Format("2006/01/02 15:04:05.999")) // 2022/02/27 00:10:42.96

	// 只格式化时分秒部分
	fmt.Println(now.Format("15:04:05"))
	// 只格式化日期部分
	fmt.Println(now.Format("2006.01.02"))
}
```

##### 解析字符串格式的时间

对于从文本的时间表示中解析出时间对象，`time`包中提供了`time.Parse`和`time.ParseInLocation`两个函数。

其中`time.Parse`在解析时不需要额外指定时区信息。

```go
// parseDemo 指定时区解析时间
func parseDemo() {
	// 在没有时区指示符的情况下，time.Parse 返回UTC时间
	timeObj, err := time.Parse("2006/01/02 15:04:05", "2022/10/05 11:25:20")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2022-10-05 11:25:20 +0000 UTC

	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	// RFC3339     = "2006-01-02T15:04:05Z07:00"
	timeObj, err = time.Parse(time.RFC3339, "2022-10-05T11:25:20+08:00")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) // 2022-10-05 11:25:20 +0800 CST
}
```

`time.ParseInLocation`函数需要在解析时额外指定时区信息。

```go
// parseDemo 解析时间
func parseDemo() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2022/10/05 11:25:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}
```

### strings

#### 包含判断

##### 前后缀

```go
func main() {
    //前后缀包含
    var str string = "This is a fat cat"
    var str1 string = "Hello Boss!"
    var str2 string = "Lucy is name My"
    fmt.Printf("T/F? Does the string \"%s\" have prefix %s?\n", str, "is") //str的前缀是否包含"is"?
    fmt.Printf("%t\n", strings.HasPrefix(str, "is"))                       //使用HasPrefix方法判断前缀
    fmt.Printf("T/F Does the string \"%s\" have prefix %s?\n", str1, "He")
    fmt.Printf("%t\n", strings.HasPrefix(str1, "He"))
    fmt.Printf("T/F Does the string \"%s\" have prefix %s?\n", str2, "My") //str2的后缀是否包含"y"
    fmt.Printf("%t\n,", strings.HasSuffix(str2, "My"))                     //使用HasSuffix方法判断后缀
}
```

##### 子字符串包含

除了可以检查前后缀，还可以使用**Contains**方法判断一个**字符串的中间**是否包含指定子字符串。

```go
func main() {

    //使用ContainsAny方法查询字符串是否包含多个字符，
    fmt.Println(strings.ContainsAny("team", "e"))     //true
    fmt.Println(strings.ContainsAny("chain", ""))     //false
    fmt.Println(strings.ContainsAny("name", "n & m")) //true

    //使用Contains方法查找某个字符串是否存在这个字符串中，若存在，返回true
    fmt.Println(strings.Contains("team", "t & a")) //false
    fmt.Println(strings.Contains("chain", ""))     //true
    fmt.Println(strings.Contains("name", "n & m")) //false
}
```

#### 索引

```go
func main() {
    str := "Hello, I'm StellarisW, Yep"
    str1 := "我是袁鑫浩，我爱玩原神"

    fmt.Printf("The position of \"StellarisW\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "StellarisW")) //使用Index方法返回str的第一个字符"Hello"的索引

    fmt.Printf("The position of the last instance of \"Hi\" is: ")
    fmt.Printf("%d\n", strings.LastIndex(str, "Yep")) //使用LastIndex方法返回str的最后一个字符"Yep"出现的索引

    fmt.Printf("The position of \"yuan\" is: ")
    fmt.Printf("%d\n", strings.Index(str, "yuan")) //不存在相应的字符串，返回-1

    //适用于非ASCII编码的字符，使用strings.IndexRune()对字符进行定位
    fmt.Printf("The position of \"袁\" is: ")
    fmt.Printf("%d\n", strings.IndexRune(str1, '袁')) //一个中文占用3字节长度，第三个字符'满'的索引值是9。
}
```

#### 替换

在Go语言中，经常使用strings.Replace函数对字符串进行处理。

```go
func main() {
    str := "我是袁鑫浩"                       //原字符串
    new := "袁神"                              //替换的新内容
    old := "袁鑫浩"                                //被替换的字符串
    n := 1                                     //表示替换到第n个old
    println(strings.Replace(str, old, new, n)) //使用Replace函数替换字符串
}
```

strings.Replace(str, old, new, n)是函数一共有四个参数，**str是原字符串，new是需要替换的内容，old是被替换的字符串，n表示匹配到第n个old**。如果n等于-1，则表示匹配所有。

#### 统计

- 出现频率

在strings包中，可以借助strings.Count(str, manyO)统计字符串出现的频率。

```go
func main() {
   
    str := "Golang is cool, right?"
    str1 := "一起玩原神吗铁汁"
    var manyG string = "o"
    
    //使用strings.Count()统计的字符串出现频率
    fmt.Printf("%d\n", strings.Count(str, manyG))
    fmt.Printf("%d\n", strings.Count(str, "oo"))
    fmt.Printf("%d\n", strings.Count(str1, "袁神"))
}
/* 
3
1
1
*/
```

- 字符数量

统计字符串的字符数量可以用两种方法：第一种是使用len([]rune(str))，先把字符串通过[]rune转化，然后再调用len()统计字符串长度。第二种则是通过utf8包中的RuneCountInString函数统计。

```go
func main() {
    str := "一起玩嘛？袁神"
    fmt.Printf("%d\n", len([]rune(str)))     //使用len([]rune())方法
    fmt.Println(utf8.RuneCountInString(str)) //使用RuneCountInString()方法
}
```

一般来说，Go语言官方更推荐第二种统计方式，但如果只是单纯统计，就没有必要引入utf8包，只需要使用第一种方式即可。以下是代码演示。

#### 大小写转换

```go
func main() {

    str := "Hello World!\n"
    fmt.Printf("%s", str)
    fmt.Printf(strings.ToLower(str)) //使用ToLower将字符串全部转换成小写
    fmt.Printf(strings.ToUpper(str)) //使用ToUpper将字符串全部转换成大写
}
/*
Hello World!
hello world!
HELLO WORLD!
*/
```

#### 修剪

使用strings.Trim()函数对字符串去掉一些不必要的内容，这一过程被称为修剪。

```go
func main() {
    str := "袁神 Hello Go 袁神"
    str1 := "          请使用TrimSpace()修剪空白字符       "
    fmt.Println(str)
    fmt.Printf("%q\n", strings.Trim(str, "袁神")) //修剪前缀和后缀
    fmt.Printf("%q\n", strings.Trim(str, "袁神 "))

    fmt.Printf("%q\n", strings.TrimLeft(str, "袁神 "))  //修剪左边前缀的字符
    fmt.Printf("%q\n", strings.TrimRight(str, " 袁神")) //修剪右边后缀的字符
    fmt.Printf("%q\n", strings.TrimSpace(str1))       //修剪前缀和后缀的空白字符
}
```

#### 分割

使用分割函数strings.Split()，函数返回的是一个切片(slice)。切片的形式是用 [ ] 括起来，在后续的复合数据类型中将会进一步学习切片的使用。

```go
func main() {
    str := "This is Golang Project"
    fmt.Println(str)
    fmt.Printf("%q\n", strings.Split(str, "Project")) //分割指定字符
    fmt.Printf("%q\n", strings.Split(str, " "))       //分割空白字符
}
/*
This is Golang Project
["This is Golang " ""]
["This" "is" "Golang" "Project"]
*/
```

### os

权限，目录，文件的打开和关闭，文件的读 \ 写、进程相关和环境相关

#### 权限

权限 perm，在创建文件时才需要指定，不需要创建新文件时可以将其设定为０

权限项	文件类型	读	写	执行	读	写	执行	读	写	执行
字符表示	（d|l|c|s|p）	r	w	x	r	w	x	r	w	x
数字表示		4	2	1	4	2	1	4	2	1
权限分配		文件所有者	文件所有者	文件所有者	文件所属组用户	文件所属组用户	文件所属组用户	其他用户	其他用户	其他用户

go 语言在 syscall 包中定义了很多关于文件操作的权限的常量例如 (部分)；

```go
const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)
```

#### 目录

os.Create 方法创建文件

```go
func Create(name string) (*File, error)
```

实例：

```go
//CreateFile 创建文件
func CreateFile(name string) {
    file, err := os.Create(name)
    if err != nil {
        fmt.Printf("err: %v\n", err)
    } else {
        fmt.Printf("file:%v\n", file)
    }
}
```

```go
func main() {
    name := "/Users/feng/go/src/StudyGin/OSlearn/hello_test.txt"
    CreateFile(name)
}
```

##### 创建目录

创建的单个目录：

```go
func Mkdir(name string, perm FileMode) error
```

实例：

```go
//CreateDir 创建单个目录
func CreateDir(name string) {
    err := os.Mkdir(name, os.ModePerm)
    if err != nil {
        fmt.Printf("err:%v\n", err)
    }
}
```

创建多级目录:

```go
func MkdirAll(path string, perm FileMode) error
```

实例：

```go
//CreateDirAll 创建多级目录
func CreateDirAll(name string) {
    err := os.MkdirAll(name, os.ModePerm)
    if err != nil {
        fmt.Printf("err:%v\n", err)
    }
}
```

##### 删除目录

删除单个空目录或文件:

```go
func Remove(name string) error
```

实例：

```go
//RemoveDir 删除单个空目录或文件
func RemoveDir(name string) {
    err := os.Remove(name)
    if err != nil {
        fmt.Printf("err:%v\n", err)
    }
}
```

强制删除当前目录下所有目录：

```go
func RemoveAll(path string) error
```

实例：

```go
//RemoveDirAll 强制删除所有目录
func RemoveDirAll(name string) {
    err := os.RemoveAll(name)
    if err != nil {
        fmt.Printf("err:%v\n", err)
    }
}
```

os.Getwd () 获取当前目录

```go
func Getwd() (dir string, err error)
```

实例：

```go
//GetWd 获取当前目录
func GetWd() {
    dir, err := os.Getwd()
    if err != nil {
        fmt.Printf("err:%v\n", err)
    } else {
        fmt.Printf("dir:%v\n", dir)
    }
}
```

os.Chdir () 修改当前工作目录

```go
func Chdir(dir string) error
```

实例：

```go
//ChWd 修改当前工作目录
func ChWd() {
    err := os.Chdir("d:/")
    if err != nil {
        fmt.Printf("err: %v\n", err)
    }
    fmt.Println(os.Getwd())
}

```

os.TempDir () 获取临时文件

```go
func TempDir() string
```

实例：

```go
//TemDir 获取临时文件目录
func TemDir() {
    fmt.Println(os.TempDir())
}
```

os.Rename () 修改文件名

```go
func Rename(oldpath string, newpath string) error
```

实例：

```go
//RenameFile 修改文件名
func RenameFile(oldpath, Newpath string) {
    err := os.Rename(oldpath, Newpath)
    if err != nil {
        fmt.Printf("err:%v\n", err)
    }
}
```

##### 文件读操作

对于文件的读 / 写操作我们，我们需要拥有相关权限，才能对文件进行读 / 写

这里我们定义一些关于权限的常量：

```go
const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)
```

打开文件

默认打开方式：

```go
func Open(name string) (*File, error)
```

os.Open() 返回一个 File 对象的指针

实例：

```go
//Open 只读文件，不能写
func Open(name string) {
    file, err := os.Open(name)
    if err != nil {
        fmt.Printf("err:%v\n", err)
    } else {
        fmt.Printf("file:%s\n", file)
    }
}
```

以指定权限打开文件

```go
func OpenFile(name string, flag int, perm FileMode) (*File, error)
```

实例：

```go
//OpenFile 以指定权限打开
func OpenFile(name string, perm os.FileMode) {
    file, err := os.OpenFile(name, O_RDONLY, perm)  //只读模式打开文件
    if err != nil {
        fmt.Printf("err:%v\n", err)
    } else {
        fmt.Printf("file:%v\n", file.Name())
    }
}
```

关闭文件

```go
func (f *File) Close() error
```

在对文件操作结束后我们需要文件关闭，所以以经常和 defer 一起使用

实例：

```go
func ReadFile(name){
  file, err := os.Open(name)
    if err != nil {
        fmt.Printf("err:%v\n", err)
        return
    }

  defer file.Close()

  //将文件读取到缓冲区
    buf := make([]byte, 100)
    n, err := file.Read(buf)
}
```

读取文件

```go
func (f *File) Read(b []byte) (n int, err error)
```

file.Read 中需要一个 byte 类型的切片做缓冲区，并将文件数据读入这个缓冲区，然后返回读取到的数据长度和 error

实例：

```go
//ReadFile 读取文件
func ReadFile(name string) {
  //打开文件
    file, err := os.Open(name)
    if err != nil {
        fmt.Printf("err:%v\n", err)
        return
    }
    for {
        //将文件读取到缓冲区
        buf := make([]byte, 5)  //设置一个缓冲区,一次读取5个字节
        n, err := file.Read(buf)
        fmt.Printf("buf:%v\n", string(buf))
        fmt.Printf("数字:%d\n", n)
        if err == io.EOF { //表示文件读取完毕
            break
        }
    }
    file.Close()
}
```

从指定位置开始读取

```go
func (f *File) ReadAt(b []byte, off int64) (n int, err error)
```

实例：

```go
//FileReadAt 从指定位置开始读取
func FileReadAt(name string) {
    file, err := os.Open(name)
    if err != nil {
        fmt.Printf("err:%v\n", err)
    }

	buf := make([]byte, 10)       // 设置一个缓冲区，一次读10个字节
	n, err := file.ReadAt(buf, 5) //从第五个开始读取
	fmt.Printf("buf:%v\n", string(buf))
	fmt.Printf("n:%s\n", n)

	file.Close()
}
```

读取目录

```go
func (f *File) ReadDir(n int) ([]DirEntry, error)
```

返回当前目录下的所有目录及文件放入 []DirEntry 中

实例：

```go
//ReadDir 获取目录
func ReadDir(name string) {
    file, err := os.Open(name)
    if err != nil {
        fmt.Printf("err:%v\n", err)
    }
    dir, err := file.ReadDir(-1)
    if err != nil {
        fmt.Printf("err:%v\n", err)
    }
    for key, value := range dir {
        fmt.Printf("dir:  key:%v, value: %v\n", key, value)
    }
}
```

设置偏移量

```go
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
```

实例：

```go
func Seek(name string) {
    file, err := os.Open(name)   //打开文件光标默认为文件开头
    if err != nil {
        fmt.Println("err:", err)
    }
    buf := make([]byte, 5)
    file.Seek(3, 0) // 从索引值为3处开始读
    n, err := file.Read(buf)
    if err != nil {
        fmt.Printf("err:%v\n", err)
    }
    fmt.Printf("info:%s\n", string(buf))
    fmt.Printf("n:%s\n", n)
    file.Close()
}
```

file.Stat () 获取文件信息

```go
func (f *File) Stat() (FileInfo, error)
```

实例：

```go
//StatFile 获取文件信息
func StatFile(name string) {
    file, err := os.Open(name)
    if err != nil {
        fmt.Printf("err:%v\n", err)
    }
    fInfo, err := file.Stat()
    if err != nil {
        fmt.Printf("err:%v\n", err)
    }
    fmt.Printf("f是否是一个文件: %v\n", fInfo.IsDir())
    fmt.Printf("f文件的修改时间: %v\n", fInfo.ModTime().String())
    fmt.Printf("f文件的名称: %v\n", fInfo.Name())
    fmt.Printf("f文件的大小: %v\n", fInfo.Size())
    fmt.Printf("f文件的权限: %v\n", fInfo.Mode().String())
}
```

##### 文件写操作

file.Write () 数据写入

```go
func (f *File) Write(b []byte) (n int, err error)
```

实例:

```go
//WritFile 写入数据
func WritFile(name string) {
    file, err := os.OpenFile(name, O_RDWR, 0775) // 以读写模式打开文件，并且打开时清空文件
    if err != nil {
        fmt.Printf("err:%v\n", err)
    }
    for i := 0; i < 10; i++ {
        file.Write([]byte(fmt.Sprintf("hello golang 我是%d\n  ", i)))
    }
    file.Close()
}
```

file.WriteString () 写入字符串

```go
func (f *File) WriteString(s string) (n int, err error）
```

实例:

```go
//WriteStringFile 写入字符串
func WriteStringFile(name string) {
    file, err := os.OpenFile(name, O_RDWR, 0775) // 以读写模式打开文件，并且打开时清空文件
    if err != nil {
        fmt.Printf("err:%v\n")
    }
    file.WriteString("您好 golang")
}
```

file.WriteAt () 写入指定位置

```go
func (f *File) WriteAt(b []byte, off int64) (n int, err error)
```

实例：

```go
//WriteFileAt 写入指定位置
func WriteFileAt(name string) {
    file, err := os.OpenFile(name, O_RDWR, 0775) // 以读写模式打开文件，并且打开时清空文件
    if err != nil {
        fmt.Printf("err:%v\n")
    }
    file.WriteAt([]byte("学习使我快乐"), 10) // 从索引值为10的地方开始写入并覆盖原来当前位置的数据
}
```

实例：使用缓冲区

```go
//WriteFile 使用缓冲区
func WriteFile(name string) {
    file, err := os.OpenFile(name, O_RDWR, 0775) // 以读写模式打开文件，并且打开时清空文件
    if err != nil {
        fmt.Printf("err:%v\n")
    }
    defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	writefile := bufio.NewWriter(file)
	for i := 0; i < 50; i++ {
	    writefile.WriteString(fmt.Sprintf("您好，我是第%d个帅哥  \n", i))
	}
	//Flush将缓存的文件真正写入到文件中
	writefile.Flush()
}
```

逐行读取
os 包没有给我们提供逐行读取的方法，这需要我们自己实现:

```go
//ReadLine 逐行读
func ReadLine(fileName string) error {
    f, err := os.Open(fileName)
    if err != nil {
        return err
    }
    //将文件读到缓冲区
    buf := bufio.NewReader(f)

	for {
	    line, err := buf.ReadString('\n')
	    line = strings.TrimSpace(line)
	    fmt.Printf("行数据:%v\n", line)
	    if err != nil {
	        if err == io.EOF {
	            return nil
	        }
	        return err
	    }
	}
	return nil
}
```

#### 进程相关

```go
func Exit(code int) // 让当前程序以给出的状态码（code）退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行。

func Getuid() int // 获取调用者的用户id

func Geteuid() int // 获取调用者的有效用户id

func Getgid() int // 获取调用者的组id

func Getegid() int // 获取调用者的有效组id

func Getgroups() ([]int, error) // 获取调用者所在的所有组的组id

func Getpid() int // 获取调用者所在进程的进程id

func Getppid() int // 获取调用者所在进程的父进程的进程id
```

#### 环境相关

```go
func Hostname() (name string, err error) // 获取主机名

func Getenv(key string) string // 获取某个环境变量

func Setenv(key, value string) error // 设置一个环境变量,失败返回错误，经测试当前设置的环境变量只在 当前进程有效（当前进程衍生的所以的go程都可以拿到，子go程与父go程的环境变量可以互相获取）；进程退出消失

func Clearenv() // 删除当前程序已有的所有环境变量。不会影响当前电脑系统的环境变量，这些环境变量都是对当前go程序而言的
```

### 其他包

path/filepath,math,flag,bytes,errors,log

## 作业

### Lv.0

复习，将上课所讲的代码都敲一遍。

### Lv.1

利用`map`实现统计数字的功能

#### 输入格式

共n+1行。

第一行是整数n，表示自然数的个数；

第2至n+1每行一个自然数。

#### 输出格式

共m行（m为n个自然数中不相同数的个数），按照自然数从小到大的顺序输出。

每行输出2个整数，分别是自然数和该数出现的次数，其间用一个空格隔开。

#### 输入输出样例

**输入 #1**

```
8
2
4
2
4
5
100
2
100
```

**输出 #1**

```
2 3
4 2
5 1
100 2
```

### Lv.2

灵活使用fmt和time包，channel，goroutine实现日程提醒组件（提醒就是在控制台打印你设置的提醒的内容）

功能包含如下：

- 可以设定提醒时间
- 可以设定提醒的内容
- 可以设定多个提醒的日程
- 可以设定重复日程（每天的多久提醒一次，每周的多久提醒一次）