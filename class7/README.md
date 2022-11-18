# 第七节课

## Redis初识

Redis全称为：Remote Dictionary Server（远程数据服务），该软件使用C语言编写，Redis是一个key-value存储系统，其有以下特点。

- 1:性能高,单线程非常适合**读多写少**的场景,可以减轻数据库压力,
- 2.集群分布式存储,可以横向拓展,可以做到高可用.
- 3.数据结构类型丰富数据类型 
- 4.支持数据持久化
- 5.支持事务（一般不用）一般通过lua脚本去实现

而在上一节课我们学习了关系型数据库**MySQL**，今天我们学习的是非关系型据库（NoSQL）**Redis ** 。

>NoSQL= Not Only SQL （不仅仅是SQL)泛指非关系型数据库 



#### 那什么是关系型数据库，什么是非关系型数据库呢❔



**关系型据库：**关系型数据库最典型的**数据结构是表**，由**二维表**及其**之间的联系**所组成的一个数据组织，其遵循ACID原则

>**优点：**
>1、易于维护：都是使用表结构，格式一致；2、使用方便：有通用的SQL语言；3、支持复杂操作：多个表之间做繁杂的查询。
>**缺点**：
>1、读写性能比较差，尤其是海量数据的高效率读写；2、固定的表结构，灵活度稍欠；3、高并发读写有IO瓶颈。

常见的有 **MySQL/Oracle/SQL Server/Sqlite/TiDB/PostgreSQL/MariaDB**

**非关系型据库**：

>**优点：**
>1、格式灵活：存储数据的格式可以是key,value形式、文档形式、图片形式。2、速度快：NoSQL可以使用硬盘或者随机存储器作为载体，而关系型数据库只能使用硬盘；3、高扩展性；4、成本低。
>
>**缺点：**
>1、一般而言没有太强的事务处理；2、数据结构相对复杂，复杂查询方面稍欠。

常见的有 **Redis / HBase /MongoDB /CouchDB /Neo4J**



## Redis数据类型和命令

Redis 目前有9种数据类型和

**5种常见类型：**String（字符串），Hash（哈希），List（列表），Set（集合）、Zset（有序集合）

**4种新增类型：**BitMap（2.2 版新增）、HyperLogLog（2.8 版新增）、GEO（3.2 版新增）、Stream（5.0 版新增）##



### **基本命令**

`redis-cli` 进入Redis命令行界面

```
root@bc2c25de5155:/data# redis-cli
127.0.0.1:6379>
```

`auth [username] password ` 默认用户是`default`

```
127.0.0.1:6379> auth default 123456
OK
127.0.0.1:6379> auth 123456
OK
```

`dbsize` 擦看数据库的key数量

```
127.0.0.1:6379> dbsize
(integer) 1
```

`select index`选择数据库 默认是0

```
127.0.0.1:6379> select 0
OK
```

`help @基本数据类型` 查看命令 

```
127.0.0.1:6379> help @string
```

`flushall` 删除所有key

`keys *`  查看所有的key

`exit` or ` quit`退出命令行

### 字符串 

String 是最基本的 key-value 结构，key 是唯一标识，value 是具体的值，value其实不仅是[字符串](https://so.csdn.net/so/search?q=字符串&spm=1001.2101.3001.7020)， 也可以是数字（整数或浮点数），value 最多可以容纳的数据长度是 `512M`。

**底层实现：**String 类型的底层的数据结构实现主要是 int 和 SDS（简单动态字符串）。



#### 基本命令



##### SET

语法 (不强制大小写)

`SET key value [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|KEEPTTL] [NX|XX] [GET]`

**EX**:表明过期时间，单位是秒 和 **setex** 相同

**PX**:单位毫秒

**EXAT**:设置时间到某个时间戳（秒级）  

**PXAT**:设置时间到某个时间戳（毫秒级）

**NX** ：当键k不存在时，设置键；设置成功返回ok，不成功时返回nil；和**SETNX**等价

**XX**： 与NX相反只在键已经存在时， 才对键进行设置操作

**GET**: 返回之前设的值


```
127.0.0.1:6379> SET wx 666
OK
127.0.0.1:6379> GET wx
"666"
127.0.0.1:6379> TTL wx
(integer) -1
```

>**ttl** 查看Key的过期时间(ms)   -1是永久 -2是没有这个值

```
127.0.0.1:6379> SET wx 666  EX 666
OK
127.0.0.1:6379> TTL wx
(integer) 662
127.0.0.1:6379> SET wx 666 EXAT 1668931126576
OK
127.0.0.1:6379> TTL wx
(integer) 1667262367568
127.0.0.1:6379> SET wx 666 PXAT 1668931126576
OK
```

**XX与NX选项**: NX当键k不存在时，设置键 ;  XX则在这个k存在时设置

```
127.0.0.1:6379> SET wx 999  NX
(nil)
127.0.0.1:6379> SET yxh 999 NX
OK
127.0.0.1:6379> GET yxh
"999"
127.0.0.1:6379> GET wx
"666"
127.0.0.1:6379> SET wx 999 XX
OK
127.0.0.1:6379> GET wx
"999"
```

GET 会返回之前的value 如果key不存在SET不会失败

```
127.0.0.1:6379> SET yxh 999
OK
127.0.0.1:6379> SET yxh 1 GET
"999"
127.0.0.1:6379> SET lmj 1 GET
(nil)  
127.0.0.1:6379> GET lmj
"1"
```

SETEX

语法:` setex key seconds value`

和 set key value ex 一样



##### PSETEX

语法：`psetex key milliseconds value`

和 set key value px 一样



##### MSET 

语法：` MSET key value [key value ...]`

同时设置多个键值对 返回OK

```
127.0.0.1:6379> MSET wx  1 lmj 2
OK
```



##### MGET

```
127.0.0.1:6379> MGET wx  lmj
1) "1"
2) "2"
```

##### MSETNX

语法：`msetnx key value [key value ...]`

msetnx a 1 b 2 c 3 批量设置键值对，当所有key都不存在时返回1，否则返回0

```
127.0.0.1:6379> msetnx a 1 b 2 c 3
(integer) 1
127.0.0.1:6379> msetnx a 1 b 2 c 3
(integer) 0
```



##### GET

语法：`GET key`

````
127.0.0.1:6379> GET lmj
"1"
````



##### GETEX

 语法：`GETEX key [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|PERSIST]`

- 用于获取k的值，并设置或者移除过期时间，参数类似于set
- EX,PX,EXAT,PXAT 和set的相同
- PERSIST移除k的过期时间

相当于或值并且重新设置过期时间

```
127.0.0.1:6379> GETEX lmj EX 1
"2"
127.0.0.1:6379> GET lmj
(nil)
```



##### GETSET

语法 `getset key value`

和 set key value get 相同



##### GETRANGE

语法：`getrange key start end`

截取字符串 下标从0开始

```
127.0.0.1:6379> getrange lanshan 0 2
"lan"
```



##### SETRANGE

语法： `setrange key offset value`

setrange  用value从偏移量（offset）开始(包括offset)，覆盖key的值

如 

```
127.0.0.1:6379> set key abc
OK
127.0.0.1:6379> setrange key 2 def
(integer) 5
127.0.0.1:6379> get key
"abdef"
```



##### STRLEN

语法 `strlen key`

strlen key 返回键key存储的值的长度，**不存在的k返回0**

```
127.0.0.1:6379> strlen lmj
(integer) 3
```

 

##### INCR

语法`incr key`

incr  对key的值加1，并返回加1后的结果；如果k的值是字符串，无法加1，则提示错误

```
127.0.0.1:6379> set k 1
OK
127.0.0.1:6379> incr k
(integer) 2
127.0.0.1:6379> set k wx
OK
127.0.0.1:6379> incr k
(error) ERR value is not an integer or out of range
```



##### INCRBY

语法 `incrby key increment`

类似 INCR 只不过可以设置任意数

```
127.0.0.1:6379> get k
"8"
127.0.0.1:6379> incrby k  2
(integer) 10
```



##### INCRBYFLOAT

语法：`incrbyfloat key increment`

增加浮点值或者加整型值

```
127.0.0.1:6379> get k
"0.3"
127.0.0.1:6379> incrbyfloat k 0.2
"0.5"
127.0.0.1:6379> incrbyfloat k 1
"1.5"
```



##### DECR

语法 `decr key`

值减一 类似incr



##### DECRBY

语法 : `decrby key decrement`

自定义减值 类似incrby



##### APPEND

语法 `append key value`

```
127.0.0.1:6379> set k 12
OK
127.0.0.1:6379> append k 345
(integer) 5
127.0.0.1:6379> get k
"12345"
```

**DEL**

语法 `del key` 删除值

```
127.0.0.1:6379> del k
(integer) 1
127.0.0.1:6379> get k
(nil)
```



#### 应用场景

##### 缓存对象

使用 String 来缓存对象有两种方式：

- 直接缓存整个对象的 JSON，命令例子： `SET user:1 '{"name":"wxgg", "age":18}'`。
- 采用将 key 进行分离为 user:ID:属性，采用 MSET 存储，用 MGET 获取各属性值，命令例子： `MSET user:1:name wxgg user:1:age 18 user:2:name wxjj user:2:age 18`。

##### 常规计数

因为 Redis 处理命令是单线程，所以执行命令的过程是原子的。因此 String 数据类型适合计数场景，比如计算访问次数、点赞、转发、库存数量等等。

##### 分布式锁

SET 命令有个 NX 参数可以实现「key不存在才插入」，可以用它来实现分布式锁：

- 如果 key 不存在，则显示插入成功，可以用来表示加锁成功；
- 如果 key 存在，则会显示插入失败，可以用来表示加锁失败。



### 哈希

Hash 是一个键值对（key - value）集合，其中 value 的形式入： `value=[{field1，value1}，...{fieldN，valueN}]`。Hash 特别适合用于存储对象。

#### 内部实现

Hash 类型的底层数据结构是由**压缩列表或哈希表**实现的：

- 如果哈希类型元素个数小于 `512` 个（默认值，可由 `hash-max-ziplist-entries` 配置），所有值小于 `64` 字节（默认值，可由 `hash-max-ziplist-value` 配置）的话，Redis 会使用**压缩列表**作为 Hash 类型的底层数据结构；
- 如果哈希类型元素不满足上面条件，Redis 会使用**哈希表**作为 Hash 类型的 底层数据结构。

**在 Redis 7.0 中，压缩列表数据结构已经废弃了，交由 listpack 数据结构来实现了**。

Hash 类型的底层数据结构是由**压缩列表或哈希表**实现的：

- 如果哈希类型元素个数小于 `512` 个（默认值，可由 `hash-max-ziplist-entries` 配置），所有值小于 `64` 字节（默认值，可由 `hash-max-ziplist-value` 配置）的话，Redis 会使用**压缩列表**作为 Hash 类型的底层数据结构；
- 如果哈希类型元素不满足上面条件，Redis 会使用**哈希表**作为 Hash 类型的 底层数据结构。

**在 Redis 7.0 中，压缩列表数据结构已经废弃了，交由 listpack 数据结构来实现了**。

#### 基本命令

##### HSET

语法 `hset key field value`

一个key的值可以有多个 field但一个 field只能有一个 value 可以和 go中的map 一起理解

```
127.0.0.1:6379> hset gocybee wx 666
(integer) 1
127.0.0.1:6379> hset gocybee yxh 666
(integer) 1
```



##### HSETNX

语法 `hsetnx key field value`

设置哈希的一个字段，当指定的字段不存在时才会被设置

```
127.0.0.1:6379> hset wx age 18
(integer) 1
127.0.0.1:6379> hsetnx wx age 18
(integer) 0
127.0.0.1:6379> hsetnx wx weight 70
(integer) 1
1234
```

在上面的命令中，age 字段已经存在于 wx这个 key 中，因此 hsetnx 命令的执行并没有添加或修改 age，而 weight 在 wx这个 key 中是不存在的，因此 weight 被添加到 wx中。



##### HMSET

语法： `hmset key field value [field value ...]`

```
127.0.0.1:6379> hset wx age 18  weight 60
(integer) 1
```



##### HGET

语法： ` hget key field`

```
127.0.0.1:6379> hget wx age
"18"
```



##### HGETALL

获取哈希的所有字段的值，该命令的用法如下：

`hgetall key`

```
127.0.0.1:6379> hgetall wx
1) "age"
2) "18"
3) "weight"
4) "60"
```

前一个是field 后面是value



##### HkEYS

语法：`hkeys key`

该命令的作用是：获取哈希的所有字段，但是不获取值

```
127.0.0.1:6379> hkeys wx
1) "age"
2) "weight"
```



##### HEXISTS

语法：`hexists key field`

该命令的作用是：判断字段是否存在于指定哈希中 

```
127.0.0.1:6379> hexists wx age
(integer) 1
127.0.0.1:6379> hexists wx tall
(integer) 0
```



##### HLEN 

语法： `hlen key`

获取指定哈希中字段的数量

```
127.0.0.1:6379> hlen wx
(integer) 2
```



##### HMGET

语法 ： ` hmget key field [field ...]`

获取指定哈希中的多个字段

```
127.0.0.1:6379> hmget wx age weight
1) "18"
2) "60"
```



##### HSTRLEN

语法 ： `hstrlen key field`

该命令的作用是：获取指定哈希中字段的长度，

```
127.0.0.1:6379> hstrlen wx age
(integer) 2
```



##### HVALS

语法： `hvals key`

该命令的作用是：获取指定哈希的所有值，不获取字段名称

```
127.0.0.1:6379> hvals wx
1) "18"
2) "60"
```



##### HINCRBY

语法： `hincrby key field increment`

和 incryby类似 将指定哈希中的指定字段的值加一个指定的整型值



##### HINCRBYFLOAT

语法： `hincrbyfloat key field increment`

将指定哈希中的指定字段的值加一个指定的浮点型值 也可以加整型值

```
127.0.0.1:6379> hincrbyfloat wx age 0.5
"18.5"
127.0.0.1:6379> hget wx age
"18.5"
127.0.0.1:6379> hincrbyfloat wx age 1
"19.5"
```



##### HDEL

语法: `hdel key`

删除哈希中指定一个或多个字段

````
127.0.0.1:6379> hdel wx b
(integer) 1
127.0.0.1:6379> hkeys wx
1) "age"
2) "weight"
127.0.0.1:6379> hdel wx age weight
(integer) 2
127.0.0.1:6379> hkeys wx
(empty array)
````





#### 应用场景

#### 缓存对象

Hash 类型的 （key，field， value） 的结构与对象的（对象id， 属性， 值）的结构相似，也可以用来存储对象。

| id   | name | age  |
| ---- | ---- | ---- |
| 1    | lmj  | 18   |
| 2    | wx   | 18   |
| 3    | yxh  | 18   |

我们可以这样来存储对象 
```
 HSET uid:1 name lmj age 18
 HSET uid:2 name wx  age 18
 HSET uid:2 name yxh age 18
```



### 列表

List 列表是简单的字符串列表，**按照插入顺序排序**，可以从头部或尾部向 List 列表添加元素。

列表的最大长度为 `2^32 - 1`，也即每个列表支持超过 `40 亿`个元素。

#### 内部实现

List 类型的底层数据结构是由**双向链表或压缩列表**实现的：

- 如果列表的元素个数小于 `512` 个（默认值，可由 `list-max-ziplist-entries` 配置），列表每个元素的值都小于 `64` 字节（默认值，可由 `list-max-ziplist-value` 配置），Redis 会使用**压缩列表**作为 List 类型的底层数据结构；
- 如果列表的元素不满足上面的条件，Redis 会使用**双向链表**作为 List 类型的底层数据结构；

但是**在 Redis 3.2 版本之后，List 数据类型底层数据结构就只由 quicklist（快速表） 实现了，替代了双向链表和压缩列表**。



#### 基本命令



##### LPUSH

语法： ` lpush key element [element ...]`

当key存在但是却不是列表会返回错误

````
127.0.0.1:6379> lpush gocybee 1 2 3 4 5 6
(integer) 6
````



##### LRANGE

语法： `LRANGE key start stop`

-1是倒数第一个 依次类推

````
127.0.0.1:6379> lrange gocybee 0 -1
1) "6"
2) "5"
3) "4"
4) "3"
5) "2"
6) "1"
````



##### LPOP

语法：`lpop key [count]`

移除列表key的表头元素，出队列 counts是出队次数

```
127.0.0.1:6379> lpush gocybee 1
(integer) 1
```



##### RPOP

语法：`RPOP key [count]`

移除列表key的尾元素，出队列 counts是出队次数

```
127.0.0.1:6379> lrange gocybee 0 -1
1) "twh"
2) "tr"
3) "lmj"
4) "000"
127.0.0.1:6379> RPOP gocybee 1
1) "000"
```



##### LSET

语法： `lset key index element`

通过索引设置列表元素的值 

```
127.0.0.1:6379> lrange gocybee 0 -1
1) "twh"
2) "lmj"
3) "wx"
4) "hhz"
127.0.0.1:6379> lset gocybee 1 tr
OK
127.0.0.1:6379> lrange gocybee 0 -1
1) "twh"
2) "tr"
3) "wx"
4) "hhz"
```



##### LLEN

语法：`llen key`

获取列表长度

```
127.0.0.1:6379> llen gocybee
(integer) 3
```



##### BLPOP

语法： `blpop key [key ...] timeout`

移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止

timeout （s）

````
127.0.0.1:6379> lpush gocybee 'hello gocybee'
(integer) 1
````

```
127.0.0.1:6379> blpop gocybee 50
1) "gocybee"
2) "hello gocybee"
(22.08s)
```



其他命令自行了解............... help @list 解决 

#### 应用场景

消息队列  以后会学 现在不讲



### 集合

Set 类型是一个无序并唯一的键值集合，它的存储顺序不会按照插入的先后顺序进行存储。

一个集合最多可以存储 `2^32-1` 个元素。概念和数学中个的集合基本类似，可以交集，并集，差集等等，所以 Set 类型除了支持集合内的增删改查，同时还支持多个集合取交集、并集、差集。

Set 类型和 List 类型的区别如下：

- List 可以存储重复元素，Set 只能存储非重复元素；
- List 是按照元素的先后顺序存储元素的，而 Set 则是无序方式存储元素的。



#### 内部实现

Set 类型的底层数据结构是由**哈希表或整数集合**实现的：

- 如果集合中的元素都是整数且元素个数小于 `512` （默认值，`set-maxintset-entries`配置）个，Redis 会使用**整数集合**作为 Set 类型的底层数据结构；
- 如果集合中的元素不满足上面条件，则 Redis 使用**哈希表**作为 Set 类型的底层数据结构。



#### 基本命令

##### SADD

语法：`SADD key member [member ...]`

 往集合key中存入元素，元素存在则忽略，若key不存在则新建

```
127.0.0.1:6379> SADD gocybee 1
(integer) 1
127.0.0.1:6379> SADD gocybee 1
(integer) 0
```



##### SREM 

语法：`SREM key member [member ...] `

从集合key中删除元素

```
127.0.0.1:6379> SREM gocybee 1
(integer) 1
127.0.0.1:6379> SREM gocybee 1
(integer) 0
```



##### SMEMBERS 

语法：`SMEMBERS key`

查看所有的members

```
127.0.0.1:6379> smembers gocybee
1) "wx"
```



##### SCARD

语法：`SCARD key`

查看元素个数

```
127.0.0.1:6379> SCARD gocybee
(integer) 1
```



##### SISMEMBER 

语法：`SISMEMBER key member`

判断元素是否是在集合中

```
127.0.0.1:6379> SISmember gocybee wx
(integer) 1
```



##### SRANDMEMBER 

语法： `SRANDMEMBER key [count]`

从集合key中随机选出count个元素，元素不从key中删除

```
127.0.0.1:6379> Srandmember gocybee 3
1) "yxh"
2) "hhz"
3) "mj"
```



##### SPOP

 语法： ` SPOP key [count]`

从集合key中随机选出count个元素，元素从key中删除

```
127.0.0.1:6379> Smembers gocybee
1) "yxh"
2) "hhz"
3) "wx"
127.0.0.1:6379> SPOP gocybee 1
1) "wx"
127.0.0.1:6379> Smembers gocybee
1) "yxh"
2) "hhz"
```



##### SINTER

语法： `SINTER key [key ...]`

交集运算返回两个set的交集

```
127.0.0.1:6379> Smembers gocybee
1) "yxh"
2) "hhz"
3) "mj"
127.0.0.1:6379> Smembers lanshan
1) "lpc"
2) "yxh"
3) "hhz"
4) "mj"
127.0.0.1:6379> SINTER  gocybee lanshan
1) "yxh"
2) "hhz"
3) "mj"
```



##### SINTERSTORE 

语法： `SINTERSTORE destination key [key ...]`

将交集结果存入新集合destination中

```
127.0.0.1:6379> SINTERSTORE common gocybee lanshan
(integer) 3
127.0.0.1:6379> smembers common
1) "yxh"
2) "hhz"
3) "mj"
```



##### SUNION

语法： `SUNION gocybee lanshan`

并集运算

````
127.0.0.1:6379> SUNION gocybee lanshan
1) "hhz"
2) "mj"
3) "lpc"
4) "yxh"
````



##### SUNIONSTORE

语法： `SUNIONSTORE destination key [key ...]`

将并集结果存入新集合destination中

```
127.0.0.1:6379> SUNIONSTORE union gocybee lanshan
(integer) 4
127.0.0.1:6379> Smembers union
1) "mj"
2) "lpc"
3) "yxh"
4) "hhz"
```



##### SDIFF 

差集运算 `SDIFF key [key ...]` 



##### SDIFFSTORE 

将差集结果存入新集合destination中 `SDIFFSTORE destination key [key ...]`



#### 应用场景



集合的主要几个特性，无序、不可重复、支持并交差等操作。

因此 Set 类型比较适合用来数据去重和保障数据的唯一性，还可以用来统计多个集合的交集、错集和并集等，当我们存储的数据是无序并且需要去重的情况下，比较适合使用集合类型进行存储。

但是有一个潜在的风险。**Set 的差集、并集和交集的计算复杂度较高，在数据量较大的情况下，如果直接执行这些计算，会导致 Redis 实例阻塞**。

在主从集群中，为了避免主库因为 Set 做聚合计算（交集、差集、并集）时导致主库被阻塞，我们可以选择一个从库完成聚合统计，或者把数据返回给客户端，由客户端来完成聚合统计。

##### 点赞

Set 类型可以保证一个用户只能点一个赞，这里举例子一个场景，key 是文章id，value 是用户id。

`uid:1` 、`uid:2`、`uid:3` 三个用户分别对 article:1 文章点赞了。

```
127.0.0.1:6379> SADD article:1 uid:1
(integer) 1
127.0.0.1:6379> SADD article:1 uid:2
(integer) 1
127.0.0.1:6379> SADD article:1 uid:3
(integer) 1
127.0.0.1:6379> SCARD article:1
(integer) 3
127.0.0.1:6379> SREM article:1 uid:1
(integer) 1
127.0.0.1:6379> SCARD article:1
(integer) 2
```



##### 共同关注

Set 类型支持交集运算，所以可以用来计算共同关注的好友、公众号等。

key 可以是用户id，value 则是已关注的公众号的id。

`uid:1` 用户关注频道号 id 为 5、6、7、8、9，`uid:2` 用户关注 频道号id 为 7、8、9、10、11。

````
127.0.0.1:6379> SADD uid:1 5 6 7 8 9
(integer) 5
127.0.0.1:6379> SADD uid:2 7 8 9 10 11
(integer) 5
````

我们可以查看 `uid:1` 和 `uid:2` 共同关注的频道

可以向`uid:2`推荐 `uid:1` 的频道等等... 



##### 抽奖活动

存储某活动中中奖的用户名 ，Set 类型因为有去重功能，可以保证同一个用户不会中奖两次。

key为抽奖活动名，value为员工名称，把所有员工名称放入抽奖箱 

再利用 `SRANDRM `或者是 `SPOP `





### 有序集合

Zset 类型（有序集合类型）相比于 Set 类型多了一个排序属性 score（分值），对于有序集合 ZSet 来说，每个存储元素相当于有两个值组成的，一个是有序结合的元素值，一个是排序值。

有序集合保留了集合不能有重复成员的特性（分值可以重复），但不同的是，有序集合中的元素可以排序。



##### 内部实现

Zset 类型的底层数据结构是由**压缩列表或跳表**实现的：

- 如果有序集合的元素个数小于 `128` 个，并且每个元素的值小于 `64` 字节时，Redis 会使用**压缩列表**作为 Zset 类型的底层数据结构；
- 如果有序集合的元素不满足上面的条件，Redis 会使用**跳表**作为 Zset 类型的底层数据结构；

**在 Redis 7.0 中，压缩列表数据结构已经废弃了，交由 listpack 数据结构来实现了。**



#### 基本命令



##### ZADD

语法 ：`ZADD key score member [[score member]...]   `

往有序集合key中加入带分值元素 ,core必须是浮点数或者整型，添加成功后返回被成功添加的新成员的数量

```
127.0.0.1:6379> ZADD movie 1 zl 2 lldq
(integer) 2
```



##### ZREM

语法 ：` ZREM key member [member...] `

往有序集合key中删除元素

```
127.0.0.1:6379> ZREM movie zl
(integer) 1
```



##### ZSCORE

语法： `ZSCORE key member`

返回有序集合key中元素member的分值

```
127.0.0.1:6379> ZSCORE movie lldq
"2"
```



##### ZCARD

语法： `ZCARD key `

返回有序集合个数

```
127.0.0.1:6379> ZCARD movie
(integer) 1
```



##### ZINCRBY

`ZINCRBY key increment member `

为有序集合key中元素member的分值加上increment

````
127.0.0.1:6379> ZINCRBY movie 3 lldq
"5"
````



##### ZRANGE

`ZRANGE key start stop [WITHSCORES]`

正序获取有序集合key从start下标到stop下标的元素

````
127.0.0.1:6379> ZRANGE movie 0 -1
1) "zl"
2) "lldq"

````



##### ZREVRANGE

`ZREVRANGE key start stop [WITHSCORES]`

倒序获取有序集合key从start下标到stop下标的元素 

```
127.0.0.1:6379> ZREVRANGE movie  0 -1
1) "lldq"
2) "zl"
```



##### ZRANGENYSCORE

`ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]`

返回有序集合中指定分数区间内的成员，分数由低到高排序。

```
127.0.0.1:6379> ZRANGEBYSCORE movie 0 100
1) "zl"
2) "lldq"
```



##### ZREVRANGEBYSCORE 

`ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]`

返回有序集合中指定分数区间内的成员，分数由高到低排序。

```
127.0.0.1:6379> ZREVRANGEBYSCORE movie 100 0
1) "lldq"
2) "zl"
127.0.0.1:6379> ZSCORE movie lldq
"5"
127.0.0.1:6379> ZSCORE movie zl
"4"
```



##### ZRANGEBYLEX

`ZRANGEBYLEX key min max  [LIMIT offset count]`

返回指定成员区间内的成员，按字典正序排列, 分数必须相同

```
127.0.0.1:6379> ZRANGEBYLEX movie - [zl
1) "zl"
2) "lldq"
```



##### ZREVRANGEBYLEX

`ZREVRANGEBYLEX key max min [LIMIT offset count]`

返回指定成员区间内的成员，按字典倒序排列, 分数必须相同



> **+ - ** 表示正无限 和负无限
>
> ( 不包含  [ 包含



#### 应用场景

Zset 类型（Sorted Set，有序集合） 可以根据元素的权重来排序，我们可以自己来决定每个元素的权重值。比如说，我们可以根据元素插入 Sorted Set 的时间确定权重值，先插入的元素权重小，后插入的元素权重大。

在面对需要展示最新列表、排行榜等场景时，如果数据更新频繁或者需要分页显示，可以优先考虑使用 Sorted Set。

##### 排行榜

有序集合比较典型的使用场景就是排行榜。例如学生成绩的排名榜、游戏积分排行榜、视频播放排名、电商系统中商品的销量排名等。

##### 电话、姓名排序

使用有序集合的 `ZRANGEBYLEX` 或 `ZREVRANGEBYLEX` 可以帮助我们实现电话号码或姓名的排序，我们以 `ZRANGEBYLEX` （返回指定成员区间内的成员，按 key 正序排列，分数必须相同）为例。



### 位图

Bitmap，即位图，是一串连续的二进制数组（0和1），可以通过偏移量（offset）定位元素。BitMap通过最小的单位bit来进行`0|1`的设置，表示某个元素的值或者状态，时间复杂度为O(1)。

由于 bit 是计算机中最小的单位，使用它进行储存将非常节省空间，特别适合一些数据量大且使用**二值统计的场景**。



#### 内部实现

Bitmap 本身是用 String 类型作为底层数据结构实现的一种统计二值状态的数据类型。

String 类型是会保存为二进制的字节数组，所以，Redis 就把字节数组的每个 bit 位利用起来，用来表示一个元素的二值状态，你可以把 Bitmap 看作是一个 bit 数组。



#### 常用命令

 自己去看



#### 应用场景

1.签到统计 2.判断用户登陆态 3.连续签到用户总数  等等



### GEO

Redis GEO 是 Redis 3.2 版本新增的数据类型，主要用于存储地理位置信息，并对存储的信息进行操作。

在日常生活中，我们越来越依赖搜索“附近的餐馆”、在打车软件上叫车，这些都离不开基于位置信息服务（Location-Based Service，LBS）的应用。LBS 应用访问的数据是和人或物关联的一组经纬度信息，而且要能查询相邻的经纬度范围，GEO 就非常适合应用在 LBS 服务的场景中



#### 内部实现

GEO 本身并没有设计新的底层数据结构，而是直接使用了 Sorted Set 集合类型。

GEO 类型使用 GeoHash 编码方法实现了经纬度到 Sorted Set 中元素权重分数的转换，这其中的两个关键机制就是「对二维地图做区间划分」和「对区间进行编码」。一组经纬度落在某个区间后，就用区间的编码值来表示，并把编码值作为 Sorted Set 元素的权重分数。

这样一来，我们就可以把经纬度保存到 Sorted Set 中，利用 Sorted Set 提供的“按权重进行有序范围查找”的特性，实现 LBS 服务中频繁使用的“搜索附近”的需求。



#### 应用场景

1.查找附加的人 车 物品



### 其他



## Go 操作Redis



我这里使用的是 `github.com/go-redis/redis/v8` 这个库

当然也可以用 ` github.com/gomodule/redigo/redis` 



### 链接



连接池以及链接设置

```go
var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		//连接信息
		Network:  "tcp",            //网络类型，tcp or unix，默认tcp
		Addr:     "127.0.0.1:6379", //主机名+冒号+端口，默认localhost:6379
		Password: "123456",         //密码
		DB:       0,                // redis数据库index

		//连接池容量及闲置连接数量
		PoolSize:     15, // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
		MinIdleConns: 10, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。

		//超时
		DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		WriteTimeout: 3 * time.Second, //写超时，默认等于读超时
		PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

		//闲置连接检查包括IdleTimeout，MaxConnAge
		IdleCheckFrequency: 60 * time.Second, //闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
		IdleTimeout:        5 * time.Minute,  //闲置超时，默认5分钟，-1表示取消闲置超时检查
		MaxConnAge:         0 * time.Second,  //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接

		//命令执行失败时的重试策略
		MaxRetries:      0,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔

		//可自定义连接函数
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			netDialer := &net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 5 * time.Minute,
			}
			return netDialer.Dial("tcp", "127.0.0.1:6379")
		},

		//钩子函数
		OnConnect: func(ctx context.Context, conn *redis.Conn) error { //仅当客户端执行命令时需要从连接池获取连接时，如果连接池需要新建连接时则会调用此钩子函数
			fmt.Printf("conn=%v\n", conn)
			return nil
		},
	})

	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Print(err)
	}
	fmt.Println("redis 链接成功")
}
```

一般这样就行

```
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456", 
		DB:       0,        
	})
```





### 字符串



````go
func GetRedisValue(ctx context.Context, key string) (string, error) {
	GetKey := Rdb.Get(ctx, key)
	if GetKey.Err() != nil {
		return "", GetKey.Err()
	}
	return GetKey.Val(), nil
}

func SetRedisValue(ctx context.Context, key string, value string, expiration time.Duration) error {
	SetKV := Rdb.Set(ctx, key, value, expiration)
	return SetKV.Err()
}

````



### 集合

```GO
type RedisSet struct {
	Id      int64
	Object  string
	Conn    *redis.Client
	Context context.Context
}

func NewRedisSet(context context.Context, Objet string, Id int64, Conn *redis.Client) *RedisSet {
	return &RedisSet{
		Id:      Id,
		Object:  Objet,
		Conn:    Conn,
		Context: context,
	}
}


func Set() {
	rs := NewRedisSet(context.Background(), "article:1", 1, Rdb)
	_, err := rs.Conn.SAdd(rs.Context, rs.Object, rs.Id).Result()
	if err != nil {
		fmt.Println(err)
	}
}

```





其他都很简单 只要会redis 命令 就和这个一样的



## 作业

### Lv0

将课上的redis 命令 和代码 自己敲一遍 

### Lv1

结合Gin 的作业

写一个获取用户信息的接口 做缓存层的处理

### Lv2

结合gin redis 写一个点赞功能的接口（只能点赞一次）


作业格式 Lv2-姓名-学号 <br>
发送到gocybee@gocybee.team 也可以发送到 limingjie@gocybee.team
### 



参考链接 https://blog.csdn.net/qq_34827674/article/details/125259739





 



