# MySQL+Gorm
#### 代码地址：https://github.com/gocybee/Courseware-2022/tree/main/class6
# 1. MySQL简介
MySQL的使用场景非常广，学习MySQL的相关命令就**需要在本机安装MySQL**（主要是方便）。
~~这个需要大家自己在网站上搜相关的安装方法（**一定要学会自己解决问题**），然后看官方文档或者博客~~。
## 1.1 数据类型
### 1.1.1 数值类型
| 类型 | 字节大小 |
| --- | --- |
| TINIY | 1 |
| SMALLINT | 2 |
| MEDIUMINT | 3 |
| INT(INTEGER) | 4 |
| BIGINT | 8 |
| FLOAT | 4 |
| DOUBLE | 8 |
| DECIMAL（小数） | 对DECIMAL(M,D)中 M>D 则为M+2，否则为D+2 |

### 1.1.2 日期和时间类型
| 类型 | 字节大小 | 格式形式 | 用途 |
| --- | --- | --- | --- |
| DATE | 3 | YYY-MM-DD | 日期值 |
| TIME | 3 | HH:MM:SS | 时间值或持续时间 |
| YEAR | 1 | YYYY | 年份值 |
| DATETIME | 8 | YYYY-MM-DD
hh:mm:ss | 混合日期值和时间值 |
| TIMESTAMP | 4 | YYYY-MM-DD
hh:mm:ss | 混合日期值和时间值和时间戳 |

### 1.1.3 字符串类型
| 类型 | 字节大小 | 用途 |
| --- | --- | --- |
| **CHAR** | 0-255 | 定长字符串 |
| **VARCHAR** | 0-65535 | 变长字符串 |
| TINYBLOB | 0-255 | <=255个字符的二进制字符串 |
| TINYTEXT | 0-255 | 短文本字符串 |
| BLOB | 0-65535 | 二进制形式的长文本数据 |
| TEXT | 0-65535 | 长文本数据 |
| MEDIUMBLOB | 0-16777215 | 二进制形式的中等长度文本数据 |
| MEDIUMTEXT | 0-16777215 | 中等长度文本数据 |
| LONGBLOB | 0-4294967295 | 二进制形式的极大文本数据 |
| LONGTEXT | 0-4294967295 | 极大文本数据 |

### 1.1.3 注意
>**CHAR(n)和VARCHAR(n)**中的**n**代表的是**字节个数**。  
>**CHAR**和**VARCHAR**类似，但是储存和检索过程中**不作大小写转换**。  
>**CHAR**申明时**会向右填充空格**。   
>**VARCHAR**只是声明最长内容的大小，与实质内容无关。  
## 1.2 MyQSL的基本使用
当然，操作数据库之前是能进入数据库，~~安装数据库的时候应该就会了叭。~~
```sql
-- 1. 显示数据库信息列表
show databases;

-- 2. 创建数据库
create database 库名;

-- 3. 删除数据库
drop database 库名;

-- 4. 进入数据库
use 库名;

-- 5. 查看表格列表
show tables;

-- 6. 创建表格
creat table 表名(字段列表);
 --例子
 /*
   CREATE TABLE `user` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(60) DEFAULT '',
  `account`VARCHAR(60)NOT NULL,
  `password` VARCHAR(60) NOT NULL,
  `question` varchar(60) default '',
  `answer` varchar(60) default '',
  PRIMARY KEY(`id`)
  )ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
*/

-- 7. 删除表格
drop table 表名;

-- 8. 查看表格字段信息
describe 表名;
desc 表名;

-- 9. 删除表格所有信息
delete from 表名；

-- 10. 详细显示表中所有信息
select * from 表名;

-- 11. 导出数据
mysqldump --opt --database test > mysql.test;
 -- "--opt"包含 "-u root -p密码" 

-- 12. 导入数据
mysqlimport -u root -p密码 < musql.dbname;

-- 13. 将文本数据导入数据库（文本中的字段数据用tab键隔开）
use 数据库名
load data local infile "文件名" into table 表名;
```
~~更加详尽的用法及细节一定要**看文档**~~！
## 1.3 事务处理
**前置条件**：必须使用**innodb引擎**的**数据库**或者数据库的**表**才能支持事务处理。
**目的**：使事务中的所有语句要么全部都执行，要么全部不执行（如果在执行的过程中出现了错误，一般使用**回滚**使得表中的数据回到初始状态）。
### 1.3.1 事务 ACID
事务必须满足4个条件：**原子性**，**一致性**，**隔离性**，**持久性**。
**原子性**：事务中的所有操作，要么全部完成，要么全部不完成。
**一致性**：在事务开始之前和事务结束以后，数据库的完整性不会被破坏（写入的数据符合所有的预设规则）
**隔离性**：数据库允许多个并发事务同时操作，隔离性可以防止多个事务并发执行而导致的数据不一致。_**事务的隔离分为不同级别：读未提交，读提交，可重复读，串行化。**_
**持久性**：事务处理结束后，对数据的修改时持久的。
### 1.3.2 事务处理原生语句
```sql
begin;  # 开始事务

insert into runoob_transaction_test value(5);

insert into runoob_transaction_test value(6);

#  SAVEPOINT savepoint_name;    # 声明一个 savepoint

#  ROLLBACK TO savepoint_name;  # 回滚到savepoint

commit; # 提交事务
```
当然，也可以通过一些命令行为即将执行的事务设置隔离级别。~~这个各位一定要去看文档~~
## 1.4 Mysql的并发问题
在很多的情况下，同一个表会被多个事务同时访问，这也造成了很多~~细思极恐~~的问题。
### 1.3.1 脏写
定义：有两个事务A和B，**同时更新某一条数据**（这条数据的初始值为NULL）,而后，当事务Ｂ更新完成后，事务Ａ突然回滚（将这条数据调回到**事务开始之前的只值NULL**），从而导致了**事务B更新了值之后值并没有改变的现象**，被称为**脏写**。
当然。这段时间的log信息也会对应不上。
### 1.3.2 脏读
定义：假设事务**A写入了一条数据**，然后事务**B读取了此条记录**，并进行了相关的操作，然而，此时**A进行了回滚**，**B再次查询**相关的值的时候，**词条信息变回了NULL**。
### 1.3.3 不可重复读
定义：一次事务**多次查询一条数据所读到的都是不同的值**。
### 1.3.4 幻读
#### 快照读
即读取快照中的数据。
#### 当前读
即读取的是最新版本的信息，并且**对读取的记录加锁**，阻塞其他事务的同时改动事务
而**幻读只可能在当前读的情况下出现**。

**幻读定义**：就是**每次事务A查询的时候都会多读到数据。**
产生原因：每一个锁只能保护**一行的内容不被修改**，但是**不能保证**整个数据库的其他行（或者数据库本身）不被操作。~~（如果直接给数据库的链接也就是DB加锁的话，会更慢）~~
# 2 GO环境下的使用
~~首先是要干嘛呢，当然时下载第三方库的驱动啊。~~
```
// golang的终端输入

go get github.com/go-sql-driver/mysql 
```
## 2.1 一般操作
首先我们直接看一段代码：
```go
package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql" // _ 表示不使用相关函数但是会自动执行init方法
	"log"
)
var db *sql.DB //定义全局变量供函数使用

//initDB :初始化db
func initDB()(err error){
	//链接数据库
	dsn:="root:@tcp(127.0.0.1:3306)/text"
	db,err =sql.Open("mysql",dsn) //不会检验用户名和密码是否正确
	//dsn格式不正确的时候才会报错
	if err!=nil{
		return
	}
	err = db.Ping() //尝试和数据库建立连接
	if err!=nil{
	    return
	}
	return err
}
func main() {
	err := initDB()
	if err!=nil{
		log.Printf("Open failed,err: %v",err)
	}
	fmt.Println("数据库链接成功！")
}
```
## 2.3 Prepare
```go
func Insert(){
    sqlStr:="insert into users(name,age) values (?,?)"
    stmt,err := db.Prepare(sqlStr) //先将不完整的mysql语句发送给服务器
    if err!=nil{
        fmt.Printf("Prepare failed,err; %v",err)
        return
    }
    //后只需要用stmt执行操作
    var m = map[string]int{
        "Alice":18,
        "Peter":40,
    }
    //将数据一个个传入数据库中
    for k,v := range m{
        _,_ = stmt.Exec(k,v) //补充原本的占位字符(?)
    }
    _ = stmt.Close()  //要关闭链接！！
}
```
## 2.2 Transaction 事务处理
```go
func transaction(){
    //正式开始事务
    Tx,err:=db.Begin()
    if err !=nil{
        fmt.Printf("Begin failed,err: %v\n",err)
    }
    //执行多个操作
    sqlStr1:= "update users set age=age-2 where id=1"
    sqlStr2:= "update users set age=age+2 where id=2"
    //执行sql语句1
    _,err1 :=Tx.Exec(sqlStr1)
    if err1 !=nil{
        //出现错误要回滚
        _ = Tx.Rollback()    //回滚表示取消此次执行的所有进度
        fmt.Println("执行sql语句1出错，需回滚")
        return
    }
    //执行sql语句2
    _,err2 :=Tx.Exec(sqlStr2)
    if err2 !=nil{
        //出现错误要回滚
        _ = Tx.Rollback()
        fmt.Println("执行sql语句2出错，需回滚")
        return
    }
    //均执行成功，则提交事务
    _ = Tx.Commit()
    fmt.Println("均执行完成!")
}

```
# 3 Gorm 框架
[参考Gorm中文文档](https://gorm.io/zh_CN/docs/context.html)
这里我们主要将的是gorm与mysql的联动，当然gorm也支持很多其他种类的数据库。
**首先一样是安装驱动**

```
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql
```
## 3.1 链接数据库
首先需要向框架提供数据库的相关信息：
```go
 dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
```
然后发送给框架：
```go
//最常见的版本
func main() {
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
//其中gorm.Config结构体中提供了更多的高级配置-详情见官方文档
//多讲一点：禁用复数表名的方法：
//1. 数据库的设置
db.SingularTable(true)
//2. 自主设置table name--写一个结构体专有的函数即可
type table struct{	//table是建表所需的结构体。
    Id int `gorm:"id"`//在后方的tag中描述数据库对应字段名
}	
func (c *table)TableName(){return "啊吧啊吧啊吧"}
```
## 3.2 gorm提供的接口
### 3.2.1 创建(插入)一条信息记录
```go
user := User{Name: "Jinzhu", Age: 18}

result := db.Create(&user) // 通过数据的指针来创建
```
			当然也支持创建多条信息（传递一个`slice`给`Create`函数），分批创建（`CreateInBatches()`）
### 3.2.2 创建钩子
Gorm允许用户自定义的钩子有`BeforeSave` `BeforeCreate` `AfterSave` `AfterCreate` 四种(其实是函数名)
```go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.UUID = uuid.New()

    if u.Role == "admin" {
        return errors.New("invalid role")
    }
    return
}
```
### 3.2.3 检索（查询）-部分
```go
db.Model("所使用的表对应的结构体").Where("你的条件").Find("结果储存的位置的地址")
// 获取第一条记录（主键升序）
db.First(&user)

// 获取一条记录，没有指定排序字段
db.Take(&user)

// 获取最后一条记录（主键降序）
db.Last(&user)
// SELECT * FROM users ORDER BY id DESC LIMIT 1;

result := db.First(&user)
result.RowsAffected // 返回找到的记录数
result.Error        // returns error or nil

// 检查 ErrRecordNotFound 错误
errors.Is(result.Error, gorm.ErrRecordNotFound)
```
### 3.2.4 更新-Update
```go
// save会储存所有的字段-即使字段是零值
db.Save(&user)

// 条件更新
db.Model(&User{}).Where("active = ?", true).Update("name", "hello")

// User 的 ID 是 `111`
db.Model(&user).Update("name", "hello")

// 根据条件和 model 的值进行更新
db.Model(&user).Where("active = ?", true).Update("name", "hello")

```
更新的知识点很多，~~目前只需要了解这一点~~（建议都看看）
### 3.2.5 删除-Delete
```go
// Email 的 ID 是 `10`
db.Delete(&email)
// DELETE from emails where id = 10;

// 带额外条件的删除
db.Where("name = ?", "jinzhu").Delete(&email)
```
当然还有很多的接口可以调用，这里只介绍一小部分**（sql**、**sql生成器**、**关联**等等，甚至可以和**Context**联动**）**，详细了解一定要**多看文档**，**多敲代码**！
## 3.3 实例分析
下面咱就直接看项目里面的gorm的使用示例(群里的QQbot)
~~虽然很拉，但是可以借鉴一下，还会继续迭代的！~~
```go
package dao

import (
	"QQbot/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.Mysql.User, global.Mysql.Password, global.Mysql.Address, global.Mysql.DbName,
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 禁用复数
	db.SingularTable(true)

	// 判断是否有聊天白名单
	if !db.HasTable(&global.ChatWhiteListStruct{}) {
		err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&global.ChatWhiteListStruct{}).Error
		if err != nil {
			return err
		}
	}

	// 判断是否有回答黑名单
	if !db.HasTable(&global.BannedAnswerListStruct{}) {
		err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&global.BannedAnswerListStruct{}).Error
		if err != nil {
			return err
		}
	}

	// 判断是否有信息记录
	if !db.HasTable(&global.AnswerAndIdStruct{}) {
		err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(&global.AnswerAndIdStruct{}).Error
		if err != nil {
			return err
		}
	}

	global.DB = db

	//将聊天白名单写入数据库
	err = writChatWhiteList(global.ChatList)
	if err != nil {
		return err
	}

	return nil
}
```
```go
package dao

import (
	"QQbot/global"
	"github.com/jinzhu/gorm"
)

var number = 0 //记录信息条数

// Banned 在所有情况下设置不能说的句子
func Banned(msgId string) error {
	// 在全局记录中寻找id对应的信息
	var c global.AnswerAndIdStruct
	err := global.DB.Model(&global.AnswerAndIdStruct{}).Where("msg_id = ?", msgId).Find(&c).Error
	if err != nil {
		return err
	}
	// 将其写入回答黑名单
	var t = global.BannedAnswerListStruct{
		Baned: c.Content,
	}
	return global.DB.Model(&global.BannedAnswerListStruct{}).Create(&t).Error
}

// Filter 将rasa的回答过滤一遍，
func Filter(answer *string) {
	var t global.BannedAnswerListStruct
	if !global.DB.Model(&global.BannedAnswerListStruct{}).Where("baned = ?", *answer).First(&t).RecordNotFound() {
		*answer = "这。。。不好说"
	}
}

// CanChatWith 是否在聊天白名单内
func CanChatWith(opp string) bool {
	// debug标志
	if global.Debug == true {
		return true
	}
	var t global.ChatWhiteListStruct

	//连着搜两次就有问题
	flag := global.DB.Model(&global.ChatWhiteListStruct{}).Where("uid = ?", opp).Find(&t).RecordNotFound()

	return !flag
}

// WriteIdAndAnswer 将信息写入数据库-只存500条
func WriteIdAndAnswer(x global.AnswerAndIdStruct) error {
	var err error
	number++
	if number >= 300 {
		err = global.DB.Model(&global.AnswerAndIdStruct{}).Where("id=?", number-299).Update(&x).Error
	} else {
		err = global.DB.Model(&global.AnswerAndIdStruct{}).Create(&x).Error
	}
	if err != nil {
		return err
	}
	return nil
}

// witChatWhiteList 初始化数据库时初始化白名单
func writChatWhiteList(uid []string) error {
	if len(uid) == 0 {
		return nil
	}

	//删除已经初始化的信息-测试时不用删库
	global.DB.Model(&global.ChatWhiteListStruct{}).Delete(&global.ChatWhiteListStruct{})

	err := global.DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range uid {
			t := global.ChatWhiteListStruct{Uid: v}
			err := global.DB.Model(&global.ChatWhiteListStruct{}).Create(&t).Error
			if err != nil {
				return err
			}
		}
		// 返回 nil 提交事务
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

```

## 作业
~~不写注释的我把你刀了~~

### Level 1
>使用**mysql的原生语句**创建一个数据库和名为user的表格(元素数量不少于3，并且包含主键)，提交截图包括**所有相关的原生语句**和**表格的描述(desc)**。
### Level 2
>使用**gorm框架**，链接level1中的数据库并对**user表**中的数据进行操作，实现**判断表格存在**、**创建表格**，并在main函数中使用**增删改查**的函数，并提交源码。
### Level 3
>基于Gorm写一个银行转账的系统，基本功能就是转账呗**(事务处理)**，其他的可以自己加。

### Extra(不用提交)
>最后一段示例代码其实有一个问题，试着找一下[**doge**]，一定要会**读代码**哦。  
_提示：和变量number有关_
