# Golang

# DOS命令

1、切换盘符

c: d: e: 没有大小写之分

2、显示详细信息**dir**

3、清屏cls

4、切换历史命令：上下箭头

5、补全：table

6、创建\删除目录:md\rd

7、复制：copy 文件 目的目录

8、删除文件：del

# 第二章 基本变量与类型

## 常量

### 1.介绍

- 常量使用const修改

- 常量在定义时必须初始化

- 常量不能修改

- 常量只能修饰bool、数值类型、string类型

- 语法：const identifier [type] = value

- 例如
  
  - const name = "tom"
  
  - const tax float64 = 0.8
  
  - const c = getVal()  //error, 函数返回值不确定

### 2. 使用注意事项

```go
// （1）简洁写法
const (
    a = 1
    b = 2
)

// （2）专业写法
const (
    a = iota // 表示给a初始化为0
    b        // 依次＋1
    c        // 所有b=1, c=2
)

// （3）Go中没有常量名字母必须大写的规范，比如TAX_RATE
// （4）仍然通过首字母大小写控制常量的访问范围
```

## 变量

### 1、介绍

变量相当于内存中一个数据存储空间的表示

### 2、使用步骤

- 声明

- 赋值

- 使用

```go
import "fmt"

func main() {
    var age int               //声明int
    age = 18
    fmt.Println("age = ", age)

    //声明和赋值可以合成一句
    var size int = 100
    fmt.Println("size is ", size)

    /*变量重复定义会报错*/

    //不可以在赋值的时候给不匹配的类型
    var age3 int = 12.98
    fmt.Println("age3 is ", age3)
}
```

### 4种使用形式

```go
import "fmt"

func main() {
    //第一种
    var age int               //声明int
    age = 18
    fmt.Println("age = ", age)

    //第二种：指定变量类型不赋值，使用默认值
    var num int
    fmt.Println("num = ", num)

    //第三种：如果没有写变量类型，根据等号=后面值90判断变量类型（自动类型推断）
    var size = 90    
    fmt.Println(size)

    //第四种：省略var，注意:= 
    sex := "男"
    fmt.Println(sex)
}
```

### 支持一次性声明多种变量(多种变量同时声明)

```go
import "fmt"

//全局变量：定义在函数外的变量
var n7 = 90

//一次性声明
var (
    n9 = 92
    str1 = "Bob"
)

func main() {
    var n1, n2, n3
    fmt.Println(n1, n2, n3)

    var n4, name, n5 = 10, "Alen", 8.0
    fmt.Println(n4, name, n5)

    n6, height := 90, 89.23
    fmt.Println(n6, height) 

    fmt.Println(n7)
    fmt.Println(n9)
    fmt.Println(str1) 
}
```

## 数据类型介绍

- 变量的数据类型

![](C:\Users\12143\AppData\Roaming\marktext\images\2023-03-08-15-37-07-image.png)

### 1、整数类型

默认整数声明类型int

选择：

在整型变量使用中，遵守保小不保大的原则，在保证程序正确运行下，尽量使用占用空间小的数据类型

### 2、浮点类型

| 类型      | 存储空间 | 范围                     |
| ------- | ---- | ---------------------- |
| float32 | 4字节  | -3.403E38 ~  3.403E38  |
| float64 | 8字节  | -1.798E308 ~ 1.798E308 |

**PS:**

    底层存储空间和操作系统无关

    浮点类型底层存储：符号位+指数位+尾数位，尾数位不准，可能出现精度损失

```go
package main
import "fmt"

func main(){
    //定义浮点类型数据
    var num1 float32 = 3.14
    fmt.Println(num1)

    var num2 float32 = -3.14
    fmt.Println(num2)

    var num3 float32 = 314E-2
    fmt.Println(num3)

    var num4 float32 = 314E+2
    fmt.Println(num4)

    var num5 float64 = 314e+2   //E大小写都可以
    fmt.Println(num5)

    //float32有精度损失，通常用float64
    //golang默认浮点类型为float64

}
```

# 第三章 运算符

- 运算符是一种特殊符号，表示数据的运算、赋值和比较等

![](C:\Users\12143\AppData\Roaming\marktext\images\2023-03-08-20-16-41-image.png)

### 算术运算符

- 介绍：对数值类型变量进行计算

- +加号：相加操作，字符串拼接

```go
package main
import "fmt"

func main(){
    //加号
    var s1 string = "abc" + "def"
    fmt.Println(s1)

    var a int =10
    a++
    fmt.Println(a)
}
```

### 关系运算符

- 运算结果都是bool型，要么是true，要么是false

### 逻辑运算符

逻辑运算符：&&（逻辑与/短路与），||（逻辑或/短路或），！（逻辑非）

### 其它运算符

- &：返回变量的存储地址

- *：取指针变量对应的数值

```go
package main

import "fmt"

func main() {
    //定义一个变量
    var age int = 18
    fmt.Println("age address is: ", &age)

    var ptr *int = &age
    fmt.Println("指针*ptr指向的数值为: \n", *ptr) //*ptr指向值
    fmt.Println("ptr指向的数值为: \n", ptr) //ptr是age的地址

}
```

# 第四章 流程控制

- 流程控制语句是用来控制程序中各语句执行顺序的语句，可以把语句组合成能完成一定功能的小逻辑模块。

- 分类
  
  - 顺序、选择和循环
  
  - 顺序：先执行a，再执行b的串行
  
  - 选择：如果......则......if else
  
  - 循环for

## 分支

### if分支

#### 1.单分支

- 基本语法

    if 条件表达式 {

        逻辑代码

}

表达式为true则执行

PS:

- 条件表达式左右的（）建议不写

- if和表达式之间有空格

- Golang中{}必须有

```go
package main

import "fmt"

func main() {
    //定义一个变量
    var age int = 18

    if age >= 18 {
        fmt.Println("adult\n")
    }

    if count := 20;count < 30 {
        fmt.Println("count is low.\n")
    }

}
```

#### 2.多分支

    if 条件表达式1 {

        逻辑代码1

    } else if 条件表达式2  {

    逻辑代码2

    } else {

    }

#### 3.双分支

-     
  
  if 条件表达式 {
  
          逻辑代码1
  
      } else {
  
      逻辑代码2
  
      }

### switch

- 基本语法

    switch 表达式 {

            case 值1, 值2, ... :

                    语句块1

            case 值3, 值4, ... :

                    语句块2

            ...

            default:

                    语句块

}

**PS：**

- switch后是一个表达式（常量值、变量、有返回值的函数都行）

- case后**各个值的数据类型**，必须和switch***表达式数据类型一致***

- case后可带多个表达式，用，隔开

- case后表达式如果是常量值，不能重复

- case后不需要带break，匹配到case执行语句块后，退出switch，都匹配不到执行default

- default不是必须的

- switch后可以不带表达式，当if分支用

- switch后也可以直接声明\定义一个变量，分号结束，*不推荐*

- **switch穿透**，利用fallthrough关键字，如果case语句块后增加fallthrough，会继续执行下一个case。

## 循环结构

- 语法
  
  for (初始表达式; 布尔表达式; 迭代因子){
  
          循环体
  
  }

- 介绍
  
  - for循环是支持迭代的一种通用结构。在初始化之前，先执行初始表达式；再对布尔表达式进行判断；再每一次迭代时都执行迭代因子。

**PS:**

- 初始表达式，不能用var定义变量的形式，要用":="

# 第五章 函数

- 为完成某一功能的程序指令（语句）的集合，称为函数。

- 基本语法
  
  func 函数名 （形参列表）(返回值类型列表){
     执行语句
     return 返回值
  
  }

## 注意事项

- 函数名
  
  遵循标识符命名规范:见名知意，驼峰命名
  
  首字母不能是数字
  
  首字母大写该函数可以被本包文件和其他包文件调用(类似public)
  
  首字母小写只能被本包文件使用，其他包文件不可调用(类似private)

- 形参列表

        可以是1个或n个参数，也可以是0个

- 返回值类型列表

        对于多个返回值，如果只想接收第一个，后面用_代替

        

```go
package main

import "fmt"

func Sum(x int, y int) (sum int, value int) {
    sum = x + y
    value = y - x
    return sum, value
}

func main() {
    rel, _ := Sum(20, 30)

    fmt.Println("rel is \n", rel)
}
```

- Go函数不支持重载

  ![](C:\Users\12143\AppData\Roaming\marktext\images\2023-03-09-15-23-45-image.png)

- Go函数支持可变参数(如果希望函数带可变数量的参数)

```go
package main

import "fmt
"
//可变参数可以是0个和任意个
func Sum(args ...int) (sum int) {
    for i := 0; i < len(args); i++ {
        sum += args[i]
    }
    return sum
}

func main() {
    rel := Sum(20, 30)

    fmt.Println("rel is \n", rel)
}
```

- 基本数据类型和数组默认以值传递，即值拷贝。在函数内修改不会影响原来的值

- 以值传递的数据类型，如果希望修改该值，可以传入地址&，其它函数以指针方式修改。

- Go函数也是数据类型，可以赋值给变量，该变量就变为函数类型的变量。通过该变量可以调用函数

```go
package main

import "fmt"

func Sum(args ...int) (sum int) {
    for i := 0; i < len(args); i++ {
        sum += args[i]
    }
    return sum
}

func main() {
    rel := Sum
    fmt.Printf("rel type is: %T, Sum type is: %T \n", rel, Sum)
}
```

- 函数是数据类型，则可以作为形参调用

- 简化数据类型定义，Go支持自定义数据类型
  
  - 基本语法：type 自定义数据类型名 数据类型
  
  - 相当于起别名
  
  - 例如：type mylnt int -----> 这时mylnt等价于int
  
  - 例如：type mySum func(int, int) -----> 这时mySum就等价函数类型func(int, int)

- 支持对函数返回值命名

## 字符串常用的系统函数

| 函数名                                                                                                                              | 功能                              |
| -------------------------------------------------------------------------------------------------------------------------------- | ------------------------------- |
| len(str)                                                                                                                         | 统计字符串长度                         |
| r:=[]rune(str)                                                                                                                   | 字符串遍历，同时处理有中文的问题                |
| n, err := strconvAtoi["12"]                                                                                                      | 字符串转整数                          |
| str := strconv.Itoa[12345]                                                                                                       | 整数转字符串                          |
| var bytes = []byte["hello go"]                                                                                                   | 字符串转[]byte                      |
| str = string([]byte{9, 10})                                                                                                      | []byte转字符串                      |
| str = strconv.FormatInt(123, 2)     // 2 -> 8, 16                                                                                | 10进制转2，8，16进制                   |
| strings.Contains("Seafood", "foo")     //true                                                                                    | 查找字符串"Seafood"是否包含子串"foo"       |
| strings.Count("Cheese", "e")             //4                                                                                     | 统计字符串有几个指定的子串                   |
| fmt.Println(strings.EqualFold("abc", "Abc"))  //true                                                                             | 不区分大小写的字符串比较(==是区分字母大小写)        |
| string.Index("NLT_abc", "abc")       //4                                                                                         | 返回子串在字符串第一次出现的index值            |
| strings.Lastindex("go golang", "go")                                                                                             | 返回子串在字符串最后一次出现的index，如果没有则返回-1  |
| strings.Replace("go go hello", "go",  "go语言", n)    //n可以指定希望替换几个，如果n=-1就全部替换<br/>//用"go语言"替换"go"<br/><br/>//原字符串并未改变，生成了新的字符串数组 | 将指定的子串替换为另外一个子串                 |
| strings.Split("hello, world, ok", ",")                                                                                           | 按照指定的某个字符，作为分割标识，将一个字符串拆分成字符串数组 |
| strings.Tolower("Go")             //go strings.ToUpper("Go")           //GO                                                      | 字符串字母大小写转换                      |
| strings.Trim("!hello!", "!")  // ["hello"], 将左右的！和""去掉                                                                           | 去掉两边指定的字符                       |
| strings.TrimLeft("!hello!", "!")                                                                                                 | 将左边指定字符去掉                       |
| strings.TrimRight("!hello!", "!")                                                                                                | 将右边指定字符去掉                       |
| strings.TrimSpace("!hello  !", "!")                                                                                              | 去掉字符串左右两边空格                     |
| strings.HasPrefix("ftp://192.168.0.1", "ftp")        //true                                                                      | 判断字符串是否以指定的字符串开头                |
| strings.HasSuffix("ftp://192.168.0.1", "ftp")        //false                                                                     | 判断字符串是否以指定的字符串结束                |

```go
    //golang的编码统一为utf-8,(ascii的字符(字母和数字)占1个字符, 汉字占3个字符)
    //统计字符长度
    str := "hello苏"
    fmt.Println(len(str))

    //字符串遍历
    //转换为rune的切片
    str2 := []rune(str)
    for i := 0; i < len(str2); i++ {
        fmt.Printf("字符=%c\n", str2[i])
    }
```

## 时间和日期相关函数

- 编程中经常使用到日期相关的函数，如统计1段代码运行时间
1. 要导入time包

2. time.Time类型，表示时间

3. 获取当前时间：
   
   1. now := time.Now()     //now的类型就是time.TIME

4. 获取其他的日期信息

```go
//获取年月日时分秒
time.Now().Year()
time.Now().Month()
time.Now().Day()
time.Now().Hour()
time.Now().Minute()
time.Now().Second()
```

5. 格式化 日期时间

```go
//1.第一种方式：按照4去格式化
now = time.Now()

//第二种
fmt.Printf(now.Format("2006/01/02 16:23:26"))


//"2006/01/02 16:23:26" 这个字符串的各个数字是固定的，必须这样写
```

6. 时间的常量

```go
const (
    Nanosecond     Duration = 1            //纳秒
    Microsecond    = 1000 * Nanosecond     //微秒
    Millisecond    = 1000 * Microsecond    //毫秒
    Second         = 1000 * Millisecond    //秒
    Minute         = 1000 * Second         //分
    Hour           = 1000 * Minute         //时
)
```

常量的作用：在程序中获取指定时间单位的时间。

                        如获得100毫秒：100 * time.Millisecond

7. 休眠

func Sleep(d Duration)

案例：time.Sleep(100 * time.Millisecond)   //休眠100毫秒

8. 获取当前unix时间戳和unixnano时间戳(作用是获取随机数字)

unix时间戳: now.Unix()

unixnano时间戳: now.Unixnano()

## 内置函数(builtin)

- 目的是为了编程方便

| 函数   | 功能                                          |
| ---- | ------------------------------------------- |
| len  | 求长度                                         |
| new  | 分配内存，主要用来分配值类型，如int、float32、struct...返回的是指针 |
| make | 分配内存，主要用来分配引用类型，如chan、map、slice。            |

## 包

### 1.使用包的原因

- 不能把所以函数放在同一个源文件中，可以分门别类的把函数放在不同源文件中

- 解决同名问题：定义了同名函数，可以用包来区分

### 2.包的使用

- package进行包的声明，建议：包的声明这个包和所在的文件夹同名

- main包是程序的入口包，一般main函数放在这个包下
  
  - main函数必须放在main包，否则无法编译执行

- 打包语法：package 包名

- 引入包的语法：import "包的路径"

        包名是从$GOPATH/src后开始计算，使用/进行路径隔离

- 引用多个包，格式：

        import(

                "fmt"

                "xx/xx"

)

- 在函数调用时候前面要定位到所在的包

- 首字母大写，函数才能被其他包访问

- 一个目录下不能有重复的函数

- 一个目录下的同级文件归属一个包

- 包到底是什么：
  
  - 程序层面：所有使用相同 package 包名 的源文件组成的代码模块
  
  - 源文件层面：就是一个文件夹

# 第六章 错误处理

## defer + recover机制处理错误

- 异常处理/捕获机制：
  
  引入defer + recover机制处理错误

```go
package main

import "fmt"

func main() {
    test()
    fmt.Println("除法执行成功")
    fmt.Println("正常执行下面的逻辑")
}

func test() {
    defer func() {
        err := recover()
        if err != nil {
            fmt.Println("err is \n", err)
        }
    }()
    num1 := 0
    num2 := 10
    rel := num2 / num1
    fmt.Println(rel)

}
```

## 自定义错误

Go支持自定义错误，使用errors.New和panic内置函数

1. errors.New("错误说明")，会返回一个error类型的值，表示一个错误

2. panic内置函数，接受一个interface()类型的值(就是任何值)作为参数，可以**接收error类型的变量，输出错误信息**，并退出程序。

```go
func test() (err error) { //声明返回的类型
    defer func() {
        err := recover()
        if err != nil {
            fmt.Println("err is \n", err)
        }
    }()
    num1 := 10
    num2 := 0

    if num2 == 0 {
        //调用errors.New函数输出自定义错误
        return errors.New("num2 can not be zero")
    } else {
        rel := num2 / num1
        fmt.Println(rel)
        return nil
    }
}
```

另一种情况：

程序出现错误，不执行后面的代码，让其中断，退出程序，**借助：builtin包下内置函数：panic**

```textile
func panic
func panic(v interface{})
内建函数panic停止当前Go程的正常执行。
当函数F调用panic时，F的正常执行就会立刻停止。
F中defer的所有函数先入后出执行后，F返回给其调用者G。
G如同F一样行动，层层返回，直到该Go程中所有函数都按相反的顺序停止执行。
之后，程序被终止，而错误情况会被报告，包括引发该恐慌的实参值，此终止序列称为恐慌过程。
```

```go
func main() {
    err := test()
    if err != nil {
        fmt.Println("自定义错误")
        panic(err) //panic使用
    } else {
        fmt.Println("除法执行成功")
        fmt.Println("正常执行下面的逻辑")
    }
}
```

# 第七章 数组

- 定义

```textile
数组定义格式：
var 数组名 [数组大小]数据类型
```

## 数组的遍历

1.普通的for循环

2.键值循环

（键值循环）for range结构是GO特有的一种迭代结构。它可以遍历数组、切片、字符串、map和通道。for range类似foreach，一般形式为

```go
for key, val := range coll {
    ...
}
```

**注意：**

1.coll就是要遍历的内容

2.每次遍历得到的索引用key接收，值用val

3.key，val名字随便起

4.key，val属于循环中的局部变量

### 数组的初始化方式

```go
package main

import "fmt"

func main() {
    //数组的初始化方式
    //第一种：
    var arr1 [5]int = [5]int{72, 34, 29}
    fmt.Println(arr1)

    //第二种：
    var arr2 = [5]int{72, 34, 29}
    fmt.Println(arr2)

    //第三种：
    var arr3 = [...]int{72, 34, 29}
    fmt.Println(arr3)

    //第四种：
    var arr4 = [...]int{2:66, 9:20, 12:40, 13:20}
    fmt.Println(arr4)
}
```

### 注意事项

- 长度属于类型的一部分

- Go中数组属值类型，在默认情况下是值传递，因此会进行值拷贝

- 其他函数修改数组，使用指针方式（引用传递）

## 二维数组

```go
package main

import "fmt"

func main() {
    //二维数组
    var arr [2][3]int
    fmt.Println(arr) //[[0 0 0] [0 0 0]]
}
```

- 二维数组初始化，与一维数组方式相同

- 二维数组的遍历

```go
package main

import "fmt"

func main() {
    //二维数组
    var arr [2][3]int = [2][3]int{{1, 3, 5}, {8, 9, 10}}
    for key, val := range arr {
        fmt.Println(key, val)
        for key2, val2 := range val {
            fmt.Println(key2, val2)
        }
    }
}
```

## 数组的使用细节

1. 数组是多个**相同类型数据**的组合，一旦声明或者定义，**长度固定，不能动态变化。**

2. var arr []int，此时arr是一个slice切片。

3. 数组中元素类型任意，但不能混用。

4. 创建数组没有赋值，则会有默认值
   
   1. 数值类型数组：默认值0
   
   2. 字符串类型数组：默认值为""
   
   3. bool类型数组：默认值为false

5. 使用数组的步骤
   
   1. 声明数组并开辟空间
   
   2. 给数组元素赋值
   
   3. 使用数组

6. 数组下标从0开始

7. 数组下标必须在指定范围内使用，否则报panic：数组越界。

8. Go的数组是值类型，默认值传递，即值拷贝

9. 在其他函数中修改原数组，通常使用指针

## 数组反转

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    //随机生成5个数，将其反转打印
    //1.随机生成5个数，rand.Intn()
    //2.放到数组
    //3.反转打印
    var array [10]int

    //为了每次生成的随机数不同，都需要一个seed值
    rand.Seed(time.Now().Unix())
    for i := 0; i < len(array); i++ {
        array[i] = rand.Intn(1000) //0 <= n < 1000
    }
    fmt.Println(array)

    for i := 0; i < len(array)/2; i++ {
        temp := array[len(array)-1-i]  //第n个元素和倒数第n个元素交换
        array[len(array)-1-i] = array[i]
        array[i] = temp
    }
    fmt.Println(array)
}
```

## 排序和查找

- 排序：将一组数据按指定顺序排列的过程。

- 排序分类
  
  - 内部排序：
    
    - 将需要处理的所有数据加载到内部存储器中进行排序(交换、选择和插入)
  
  - 外部排序
    
    - 数据量过大，无法全部加载到内存，需要借助外部存储进行排序(合并和直接合并)
1. 交换式排序

        属于内部排序，比较数据值，依据判断规则对数据位置进行交换达到排序目的。

       两种：**冒泡和快速(Bubble sort和quick sort)**

     1.1 冒泡

        一共会经过arr.length-1次的轮数比较

        每一轮确定一个数的位置

        每一轮比较次数减小

```go
    for i := 0; i < (len(array) - 1); i++ {
        for j := i + 1; j < len(array); j++ {
            if array[i] <= array[j] {
                continue
            } else {
                temp = array[i]
                array[i] = array[j]
                array[j] = temp
            }
        }
        fmt.Printf("第%d轮排序后的结果是%d\n", i, array)
    }
```

    1.2 快排

            介绍：

快速排序（Quick Sort）是一种基于比较的排序算法，通过将数组或列表分成较小的部分，从而递归地排序整个数组。

下面是快速排序的详细步骤：

1. 选择一个基准元素（pivot），一般选择第一个元素或最后一个元素。
2. 将数组分成两个子数组，小于等于基准元素的放在左边，大于基准元素的放在右边。
3. 对于左右两个子数组，递归地调用快速排序算法，直到子数组只有一个元素为止。
4. 将所有子数组的排序结果合并起来，就得到了最终的排序结果。

在快速排序的过程中，关键是如何将数组分成两个子数组。通常有两种方法：

1. 挖坑填数法：先选定一个基准元素，将该元素从数组中移除，然后从数组的最右边开始扫描，找到第一个小于基准元素的数，并将其填入到基准元素的位置；接着从数组的最左边开始扫描，找到第一个大于基准元素的数，并将其填入到上一个数的位置；如此循环进行，直到左右指针相遇，此时基准元素左边的子数组全部小于等于基准元素，右边的子数组全部大于基准元素。
2. 左右指针法：同样是先选定一个基准元素，然后左右指针分别从数组的左右两端开始扫描，左指针向右移动找到大于等于基准元素的数，右指针向左移动找到小于等于基准元素的数，然后交换两个数的位置。直到左右指针相遇，此时基准元素左边的子数组全部小于等于基准元素，右边的子数组全部大于基准元素。

两种方法的思想相同，都是利用左右指针不断地交换数组元素的位置，从而实现将数组分成两个子数组的目的。快速排序的时间复杂度为 O(nlogn) ~ O(n^2)，取决于基准元素的选择以及数组的分布情况。在实际应用中，快速排序是一种高效的排序算法，被广泛应用于各种编程语言和数据结构中。

### 查找

- 常用查找：
  
  - 顺序
  
  - 二分(数组有序)

# 第八章 切片

## 切片的引入

- 切片是go一种数据类型

- 数组有些呆板（长度固定不变），在Go中不常见，切片是建立在数组类型之上的抽象，构建在数组之上且更强大。

- 切片(slice)是对数组一个连续片段的引用，所以***切片是一个引用类型***。这个片段可以是整个数组，或者是由起始和终止索引标识的一些项的子集。**注意索引标识项不在子集内*。*切片提供了一个相关数组的动态窗口。

语法：var 切片名 []数据类型 = 原数组[start:end]

```go
package main

import "fmt"

func main() {
    //二维数组
    var arr = [6]int{1, 3, 5, 8, 9, 10}

    //切片构建在数组之上
    //定义一个切片slice，[]动态变化的数组长度不写，int类型，arr是原数组
    var slice []int = arr[2:6]
    fmt.Println(slice)
    //切片长度
    fmt.Println(len(slice))
    //切片容量
    fmt.Println(cap(slice))
}
```

## 内存分析

切片有3个字段的数据结构：

- 一个是指向底层数组的指针

- 一个是切片长度

- 一个是切片容量

## 切片的定义

```go
//定义一个切片，然后去引用一个已经创建好的数组
var slice []int = arr[2:6]


//make内置函数创建切片。
//基本语法：var切片名[type = make([], len, [cap])]
slice1 := make([]int, 4, 20)

//定义一个切片，直接指定具体数组，使用类似make的方式
slice2 := []int{1, 4, 7}
```

- make创建一个底层数组，对外不可见，所以不可以直接操作这个数组，要通过slice去间接访问各个元素， 不能直接对这个数组进行维护\操作

## 切片的注意事项

1. 切片定义后不可直接使用，需要将切片引用到一个数组，或者make一个空间供切片使用

2. 切片使用不能越界

3. 简写方式：
   
   ```textile
   1.var slice = arr[0:end] ----->> var slice = arr[:end]
   2.var slice = arr[start:len(arr)] ----->> var slice = arr[start:]
   3.var slice = arr[0:len(err)] ----->> var slice = arr[:]
   ```

4. 切片可以继续切片

5. 切片可以动态增长**append()**
   
   slice2 := append(slice, 99, 8) 
   
   //底层追加元素是对老数组扩容为新数组
   
   //创建新数组，将老数组值复制到新数组中，再追加99，8
   
   //slice2指向的是新数组
   
   //使用追加的时候是给slice追加的话，赋值给slice：slice := append(slice, 99, 8)
   
   //底层新数组不能直接维护，还是用切片去维护操作
   
   5.1 通过append将切片添加给切片
   
   slice = append(slice, slice2...)
   
   //**注意，...是必须的**

6. 切片的拷贝**copy()**
   
   ```go
   package main
   
   import "fmt"
   
   func main() {
       //二维数组
       var slice []int = []int{1, 3, 5, 8, 9, 10}
       var slice_01 []int = make([]int, 10)
   
       //拷贝
       copy(slice_01, slice) //将slice数组中的值拷贝到slice_01
       fmt.Println(slice_01)
   }
   ```

## string和slice

1. string底层是一个byte数组，因此string也可以进行切片处理

2. string和切片在内存的存在形式：切片生成新的string数组

3. string是不可变的，不可通过string[0]='z'方式修改字符串

4. 如果需要修改字符串，先将string->[]byte / 或者 []rune ->修改 ->重写转成string

```go
//"hello" -> "iello"
str := "hello"
arr := []byte(str)
arr[0] = 'i'

str = string(arr)

//细节，转成[]byte可以处理数字和英文，但是不能处理中文
//原因是[]byte字节来处理，而一个汉字，是3个字节，因此就会出现乱码
//解决方法是将string 转成 []rune即可，因为[]rune是按字符处理，兼容汉字
arr := []rune(str)
arr[0] = "北"

str = string(arr)
```

## 切片练习

- 编写函数fbn(n int)，要求：
  
  - 可以接收一个n int
  
  - 能够将Fibnacci数列放入切片

```go
// arr[0]=1;arr[1]=1;arr[2]=2;arr[3]=4
func fbn
```

# 第九章 映射

## 1.map的引入

键值对key-value

- Go内置一种类型，将键值对相关联，可以通过键获取对应的值value。

- 类似python的字典

- 基本语法
  
  - var map变量名 map[keytype]valuetype
  
  - key, value的类型：bool、数字、string、指针、channel、还可以是包含上面类型的接口、结构体、数组
  
  - key通常为int、string类型，value一般为数字、string、map、结构体
  
  - 作为key不能是slice、map、function

- map的特点：
  
  - map使用前一定要make，让内存分配空间
  
  - map的key-value是无序的
  
  - 如果key重复，会覆盖前面的键值对
  
  - value可以重复
  
  - make的第2个参数size可以省略，默认分配一个数据类型的空间

```go
package main

import "fmt"

func main() {
    //定义map
    var a map[int]string
    //只声明map内存没有分配空间
    //必须通过make函数对其初始化，内存才会分配空间
    a = make(map[int]string, 10) //map可以存放10个键值对
    //将键值存入map中
    a[12] = "sadwe"
    a[14] = "alen"

    fmt.Println(a)
}
```

## 2.map的三种创建方式

```go
package main

import "fmt"

func main() {
    //方式1
    var a map[int]string
    a = make(map[int]string, 10) //map可以存放10个键值对

    //方式2
    b := make(map[int]string)

    //方式3
    c := map[int]string{
            1:"tsg",
            20:"sda",
         }

    //多重map
    studentMap := make(map[string]map[string][string])
    studentMap["stu01"] = make(map[string]string, 3) //这句话不能少
    studentMap["stu01"]["name"] = "Will"
    studentMap["stu01"]["sex"] = "男"
    studentMap["stu01"]["address"] = "虎丘"

}
```

## 3.map的操作

1. 增加和更新操作
   
   1. map["key"] = value ----->> 如果key还没有，就增加，可以存在就是修改对应的value

2. 删除操作
   
   1. delete(map, "key") ，delete是内置函数，如果key存在则删除键值对，key不存在也不会报错

3. 清空操作
   
   1. 如果要删除map所有的key，没有一个专门的方法一次删除，一般是遍历key去删除
   
   2. map = make(...)，make一个新的，原来的被gc回收

4. 查找操作

        value, bool = map[key]

        value为返回的value， bool为true\false，表示返回成功\失败

## 4.map的遍历

- map使用for-range的结构遍历

```go
package main

import "fmt"

func main() {
    //使用for-range遍历map
    cities := make(map[string]string)
    cities["no1"] = "北京"
    cities["no2"] = "上海"
    cities["no3"] = "苏州"

    for _, val := range cities {
        fmt.Printf("city is %s\n", val)
    }
}
```

## 5.map切片

- 基本介绍

    切片数据类型如果是map，则称为slice of map， map切片，这样使用map个数可以动态变化。

- 案例演示
  
  使用map记录monster的信息name和age，一个monster对应一个map，且monster的个数可以动态的增加 ==> map切片

```go
package main

import "fmt"

func main() {
    //使用map记录monster的信息name和age，一个monster对应一个map，且monster的个数可以动态的增加 ==> map切片
    //map切片的使用

    //1.声明一个map切片
    var monsters []map[string]string
    monsters = make([]map[string]string, 2) //放2个monster

    //2.增加妖怪信息
    if monsters[0] == nil {
        monsters[0] = make(map[string]string, 2)
        monsters[0]["name"] = "东皇太一"
        monsters[0]["race"] = "妖"
    }
    fmt.Println(monsters)

    //如果增加多个monster，则需要切片的append函数，动态增加monster
    //1.先定义monster信息
    newMonster := map[string]string{
        "name": "孙悟空",
        "race": "神",
    }
    monsters = append(monsters, newMonster)
    fmt.Println(monsters)
}
```

## 6.map排序

- 基本介绍

        1. go没有一个专门的方法针对map的key进行排序

        2.go的map默认无序，注意也不是按照添加的顺序存放，每次遍历，得到的输出可能不太一样。

        3.map的排序可以先将key排序，再根据key值遍历输出即可。

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    //map排序
    map1 := make(map[int]string, 10)
    map1[0] = "唐僧"
    map1[1] = "孙悟空"
    map1[2] = "猪八戒"
    map1[3] = "沙僧"
    map1[4] = "白龙马"
    map1[5] = "金翅大鹏"
    map1[20] = "玉兔"
    map1[15] = "锦毛鼠"

    fmt.Println(map1)
    //按照map的key顺序排序输出
    //1.先将map的key放入切片
    //2.对切片排序
    //3.遍历切片，按照key输出map的value
    var keys []int
    for key, _ := range map1 {
        keys = append(keys, key)
    }
    fmt.Println(keys)

    //排序
    sort.Ints(keys)
    fmt.Println(keys)

    for _, k := range keys {
        fmt.Println(map1[k])
    }

}
```

## 7.map使用细节

1. map是引用类型，遵守引用类型传递的机制，在一个函数接收map，修改后会直接修改原来的map

2. map的容量达到后，再想map增加元素，会自动扩容，并不会发送panic，就是说map能动态增长键值对(key-value)

3. map的value也经常使用struct类型，更适合管理复杂的数据。比如value为结构体。

## 8.练习

1. 使用map[string]map[string]string的map

2. key：用户名，唯一不可重复

3. 如果用户名存在，将密码修改为"888"，存在则增加该用户信息(包括nickname和passwd)

4. 编写一个函数modifyUser(users map[string]map[string]string, name string)完成上述功能

```go
func modifyUser(users map[string]map[string]string, name string) {
    defalt_passwd := "888"

    if _, ok := users[name]; ok {
        if defalt_passwd != users[name]["passwd"] {
            users[name]["passwd"] = defalt_passwd
        }
    } else {
        users[name] = make(map[string]string, 2)
        users[name]["passwd"] = defalt_passwd
        users[name]["nickname"] = "昵称~" + name

    }
    fmt.Println(users)
```

# 第十章 Go面向对象编程说明

## 1.引入

- 1.Go支持面向对象编程，但不是纯粹的。准确说法是支持面向对象编程特性。
- 2.Go没有类，结构体(struct)和其他语言的类(class)地位相同。可以理解Go是基于struct实现oop特性的。
- 3.Go面向对象编程简洁，去掉了传统OOP语言的方法重载、构造函数和析构函数、隐藏的this指针等等
- 4.GO仍然有面向对象编程的继承，封装和多态的特性，只是实现方式和其他OOP语言不同.
  - 比如继承：Golang没有extends关键字，继承是通过匿名字段来实现。
- OOP本身是语言类型系统的一部分，通过接口(interface)关联，耦合性低，也非常灵活。Go中**面向接口编程**是非常重要的特性。

## 2.结构体的声明和使用陷阱

### 2.1 定义

```go
  package main
  import "fmt"

  //定义结构体，将各个属性统一放入结构体中管理：
  type Teacher struct{
      //变量名字大写外界可以访问这个结构体
      Name   string
      Age    int
      School string
  }

  func main(){
      //创建具体的结构体的实例、对象、变量
      var ma Teacher                   //不赋字段初始值默认{0}，所以结构体是值类型
      ma.Name = "Mark"
      ma.Age  = 18
      ma.scholl = "Suzhou"
  }
```

### 2.2 字段/属性

- 注意事项和说明
1. 字段声明语法同变量

2. 字段类型可以是：基本类型、数组或引用类型

3. 在创建结构体变量后，没有赋值会有默认值，规则如下：
   
   1. bool类型是false，整型是0，字符串类型是""
   
   2. 数组类型默认值和元素类型相关，如score [3]int就是[0,0,0]
   
   3. 指针，slice，map的默认值都是nil，因为还没分配空间

4. 不同结构体变量字段独立，互不影响，一个结构体变量字段更改不影响另一个。

## 3.结构体和结构体实例的区别

1. 结构体是自定义的数据类型

2. 实例是使用结构体具体的变量

## 4.结构体变量的声明

```go
//1.直接声明
var person Person

//2.{}
stu := Stu{"Jin", 20, "男"}
    fmt.Println(stu)

    //或者是
    stu.Name = "Cheng"
    stu.age = 22
    stu.sex = "男"

    //方式3：直接返回指针
    var stu2 *Stu = new(Stu)
    //stu2是一个指针，因此给字段赋值的标准方式是：
    //底层会进行处理，去掉*
    (*stu2).Name = "Hehai"
    stu2.age = 100

    //方式4
    var stu3 *Stu = &Stu{"Fav", 20, "男"}

    //或者
    var stu4 *Stu = &Stu{}
    //因为stu4是一个指针，用标准的访问字段的方法
    //也可以直接写为
    stu4.age = 29

    fmt.Println(stu2, stu3, stu4)
```

## 5.结构体的注意事项和使用细节

1. 结构体的所有字段在内存中连续

## 6.结构体之间的转换

- 6.1 结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的字段(名字、个数和类型)

- 6.2 结构体进行type重新定义(相当于取别名)，Go认为是新的数据类型，但是相互间可以强转

```go
 package main
 type Student struct {
 Age int
}

type Person struct {
 Age int
}

//6.2 取别名强转
type Stu Student

func main() {
 var s Student = Student{18}
 var p Person = Person{18}

 //6.1 转换要有完整字段
p = Person(s)

//6.2 取别名强转
var s2 Stu = Stu(s)
 }

//直接赋值
var ma Teacher = Teacher{"Mark", 18, "Suzhou"}

//方式二：返回的是结构体指针
var t *Teacher = new(Teacher) //t是指针，指向的是结构体地址，应该给这个地址指向的对象的字段赋值
(*t).Name = "Huan"

//Go提供了简化的赋值
t.Age = 26 //Go底层对t.Age转化为(*t).Age

//方式三：返回的是结构体指针
var t *Teacher = &Teacher{"Mark", 18, "Suzhou"}
```

- 6.3 struct的每个字段，可以写上1个tag，该tag可以通过**反射机制**获取，常见场景是序列号和反序列化。

- 使用场景：

```go
//将struct变量进行json处理
/*
问题：json处理后的字段名也是首字母大写，如果将json后的字符串给其他程序，
比如query，php等，可能不习惯这种命名方式，怎么办？

解决方案：
1.将monster字段首字母小写，行不通，处理后返回的是空字符串，因为json.Marshal相当于
  在其他包访问monster结构体，首字母小写就不能在其他包访问
2.使用tag标签解决
*/
package main

import (
    "encoding/json"
    "fmt"
)

// 定义学生结构体
type Monster struct {
    Name  string `json:"name"` //`json:"name"`就是struct tag
    Age   uint16 `json:"age"`
    Skill string `json:"skill"`
}

func main() {
    var monster Monster
    monster.Name = "红孩儿"
    monster.Age = 800
    monster.Skill = "三昧真火"

    //将其进行json格式的序列号处理
    //json.Marshal使用到反射
    data, err := json.Marshal(monster)
    if err != nil {
        fmt.Println("json encoding err\n", err)
        return
    }
    fmt.Println("Json后的数据是: ", string(data)) //Json后的数据是:  {"Name":"红孩儿"}, N是大写的

}
```

## 7 方法的引入

- 1.方法是作用在指定的数据类型上，和指定的数据类型绑定，因此自定义类型，都可以有方法，而不仅仅是struct

- 2.方法的声明和调用格式：
   声明：
  
  ```
    type A struct {
        Num int
    }
    func (a A) test() {
        fmt.Println(a.NUM)
    }
  ```
  
  调用：
  
  ```
    var a A
    atest()
  ```
  
  解释：
  
  ```
    (1)func(a A) test()相当于A结构体有一个方法叫test
    (2)(a A)体现方法test和结构体A绑定关系
  ```
  
  - 3.代码层面
    
    ```go
      package main
      import "fmt"
    
      type Person struct{
          Name string
      }
    
      //绑定方法
      func(per Person) test(){ //参数名字无所谓
          fmt.Prinln("test, p.Name is ", p.Name)
      }
    
      func main(){
          //创建结构体对象
          var p Person
          p.Name = "Asua"
          p.test()
      }
    ```
    
    - 注意：
       (1) test方法中参数名字无所谓
       (2) 结构体Person和test方法绑定，调用test方法必须靠指定的类型：Person
       (3) 如果其他类型变量调用test会报错
       (4) 结构体对象传入test方法中，值传递，和函数传参一致
    
            (5) func(p Person) test() {}...
    
                  p表示哪个Person变量调用，p就是它的副本，见4

### 7.1 方法的注意事项

1. 结构体类型是值类型，在方法调用中，遵循值传递机制，是值拷贝传递方式

2. 如果想改变结构体变量的值，可以通过结构体指针的方式处理
   
   ```go
   package main
   import "fmt"
   
   type Person struct{
      Name string
   }
   
   //绑定方法
   func(per *Person) test(){
      p.Name = "Lulu"
      fmt.Prinln("test, p.Name is ", p.Name)
   }
   
   func main(){
      //创建结构体对象
      var p Person
      p.Name = "Asua"
      p.test()
      fmt.Println(p.Name) //输出Lulu，注意绑定方法中的*Person，
                          //此时是结构体指针，底层会自动加上&,*
   }
   ```
   
   注意：
   方法中是结构体指针的话, 底层会自动加上&,*

3. Go中方法作用在指定的数据类型上, 和指定的数据类型绑定, 
   因此自定义类型, 都可以有方法, 而不仅仅是struct, 比如int, float32都有方法
   
   ```go
   package main
   import "fmt"
   
   type integer int
   
   //绑定方法
   func(i integer) print(){
      fmt.Prinln("i =  ", i)
   }
   
   func main(){
      var i integer = 200
      i.print()
   }
   ```

4. 方法的访问范围控制规则和函数一样，首字母小写只能在本包访问; 首字母大写，可以在本包和其他包访问。

5. 如果一个变量实现了String()这个方法，那么fmt.Println会默认调用这个变量的String()进行输出，

### 7.2 方法快速入门

```go
// 定义结构体
type Monster struct {
    Name  string `json:"name"` //`json:"name"`就是struct tag
    Age   uint16 `json:"age"`
    Skill string `json:"skill"`
}
```

1. 给Monster结构体添加speak方法

```go
func (m Monster) speak() {
 fmt.Printf("抓住唐僧\n")
}
```

2. 给Monster结构体添加calculate方法，可以计算连加的结果，说明方法体内可以像函数一样

```go
// 给Monster结构体添加calculate方法
func (m Monster) calculate(start int, end int) int {
    res := 0
    res = (start + end) * ((end - start + 1) / 2)
    return res
}
```

### 7.3 方法的调用和传参机制原理

- 说明：方法的调用和传参机制和函数基本一样，

- 不同的是：变量调用方法时，变量monster本身也会作为一个参数传递到方法(变量是值类型则值拷贝，是引用类型则地址拷贝)。

![](C:\Users\12143\AppData\Roaming\marktext\images\2023-03-16-10-29-35-image.png)

### 7.4 方法的声明和定义

```
func (receiver type) MethodName(参数列表) (返回值列表) {
    方法体
    return 返回值
}

(1)参数列表: 表示方法输入
(2)receiver type: 表示这个方法和type这个类型绑定，或者这个方法作用于这个type类型
(3)receiver type: type可以是结构体，也可以是其他的自定义类型
(4)receiver: type类型的一个实例，如Monster结构体的实例monster
(5)返回值列表: 表示返回的值，可以是多个
(6)方法主体: 代码块
(7)return不是必须的
```

### 7.4 练习

```go
//1.编写一个方法，提供参数m,n,打印一个m*n的矩形
func (m Matrix) plot(Length uint16, Width uint16) {
    for i := 0; i < int(Length); i++ {
        fmt.Printf("+\t")
    }
    fmt.Println()
    for i := 0; i < int(Width); i++ {
        fmt.Printf("+\t")
        for z := 0; z < (int(Length) - 2); z++ {
            fmt.Printf(" \t")
        }
        fmt.Printf("+\t")
        fmt.Println()
    }
    for i := 0; i < int(Length); i++ {
        fmt.Printf("+\t")
    }
    fmt.Println()
}

//2.编写方法计算矩形面积并返回
func (m Matrix) calculate(Length uint16, Width uint16) uint16 {
    res := Length * Width
    return res
}
```

### 7.5 工厂模式

- Go结构体没有构造函数，使用工厂模式解决该问题

结构体定义中，首字母大写可以被其他包调用，小写想调用是不行的。

所以用**工厂模式**解决。

案例：

1. 如果model包的结构体变量首字母大写，引入后，直接使用没有问题

2. 首字母小写，用工厂模式解决

```go
type stu struct {
    Name string
    Sex  string
}

//通过工厂模式解决首字母小写无法被其他包调用的问题
func Newstu(n string, s string) *stu {
    return &stu{
        Name : n,
        Sex  : s,
    }
}

// 如果Sex是小写的sex
// 首字母小写不能访问，使用方法设置成可以访问
func (s *stu) GetSex() string {
    return s.sex
}

//在main.go中调用时
var stu = mode.Newstu{
    Name:  "Selen",
    Sex: "女",
}
```

## 8.抽象

- 定义结构体实际上是把一类事物共有属性和行为（字段和方法）抽象出来形成模型。这周研究方法叫**抽象**。

# 第十一章 面向对象编程三大特性--继承、封装和多态

## 1.基本介绍

Go也有继承、封装和多态的特性，和其他oop不太一样。

## 2.封装

- 基本介绍

        封装(encapsolution)是把抽象出的字段和对字段的操作封装在一起，数据被保护在内部，程序的其它包只有通过被授权的操作(方法)，才能对字段进行操作。

- 封装的理解和好处

        (1) 隐藏实现细节

        (2) 可对数据进行验证，保证安全合理

- 如何体现封装

        (1) 对结构体属性进行封装

        (2) 通过方法，包实现封装

### 2.1 封装实现步骤

1. 将结构体、字段(属性)的**首字母小写**(不能导出被其他包使用，类似private)

2. 给结构体所在包提供一个**工厂模式函数**，**首字母大写**。类似一个构造函数

3. 提供一个**首字母大写**的Set方法(类似其他语言的public)，用于对属性判断并赋值

```
func (var struct_type_name) Setxxx(paras_list)(return_list){
    //加入数据验证的业务逻辑
    var.字段=参数
}
```

4. 提供一个**首字母大写**的Get方法(类似其他语言的public)，用于获取属性的值

```
func (var struct_type_name) Getxxx()(return_list){
    //加入数据验证的业务逻辑
    return var.字段
}
```

特别说明：

        在Go中没有特别强调封装，是对面向对象的特性做了简化。

### 2.2 快速入门

不能查看人的年龄、工资等隐私，并对输入的年龄进行合理性验证

```go
//main.go
package main

import (
    "fmt"
    "gocode/project01/unit10/encapsolution/model"
)

func main() {
    p := model.NewPerson("Smith")
    p.SetAge(28)
    p.SetSal(6890.78)

    fmt.Println(p)
    fmt.Printf("%T\n", p)

}
//model/model.go
package model

import "fmt"

type person struct {
    Name string
    age  uint8
    sal  float64
}

// 创建工厂模式函数
func NewPerson(name string) *person {
    return &person{ // &person{}，此时是一个指向person类型结构体实例的指针。
        Name: name,
    }
}

// 访问age和sal，编写Set和Get方法
func (p *person) SetAge(age uint8) {
    if age > 0 && age < 100 {
        p.age = age
    } else {
        fmt.Println("年龄不正确...")
    }
}

func (p *person) GetAge() uint8 {
    return p.age
}

func (p *person) SetSal(sal float64) {
    if sal > 0 {
        p.sal = sal
    } else {
        fmt.Println("工资不是负数...")
    }
}

func (p *person) GetSal() float64 {
    return p.sal
}
```

### 2.3 练习（待做）p207

## 3. 继承

- 引入继承的原因
1. 结构体字段和方法类似，却写了相同代码

2. 出现了代码冗余，且代码**不利于维护**，不利于**功能的扩展**

3. 通过**继承**的方式解决

### 3.1 继承的基本语法

当多个结构体存在相同的属性(字段)和方法时，可以从这些结构体中抽象出结构体，在该结构体中定义这些属性和方法。

其他结构体不需要重新定义，只需嵌套这个匿名结构体即可。

即：在Go中，如果一个结构体struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而体现继承特性。

- 基本语法

```
type Goods struct {
    Name  string
    Price int
}

type Book struct {
    Goods        //这里就是嵌套匿名结构体Goods
    Writer string
}
```

### 3.2 快速入门

```go
package main

import (
    "fmt"
)

// 编写一个学生考试系统
type Student struct {
    Name  string
    Age   uint8
    Score float64
}

//将pupil和graduate共有方法绑定到* Student

func (stu *Student) ShowInfo() {
    fmt.Printf("姓名: %v, 年龄: %v, 分数: %v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score float64) {
    stu.Score = score
}

func (stu *Student) ShowScore() float64 {
    return stu.Score
}

type Pupil struct {
    Student //嵌入Student匿名结构体
}

func (stu *Pupil) testing() {
    fmt.Println("小学生正在考试中....")
}

type Graduate struct {
    Student
}

func (stu *Graduate) testing() {
    fmt.Println("大学生正在考试中....")
}

func main() {
    pupil := &Pupil{
        Student{
            Name: "Tom",
            Age:  10,
        },
    }

    pupil.testing()
    pupil.Student.SetScore(90.88)
    pupil.Student.ShowInfo()

}
```

### 3.3 深入讨论

1. 结构体可以使用嵌套匿名结构体的所有字段和方法，即：首字母大、小写的字段、方法都可以使用

2. 匿名结构体字段访问可以简化

        如b.A.name可以简化为b.name

3. 当结构体和匿名结构体具有相同的字段或方法时，编译器采用就近访问原则。如果希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分。

4. 结构体嵌入2个或多个匿名结构体，如2个匿名结构体有相同的字段和方法(结构体本身没有相同的字段和方法)，在访问时，**必须明确指定匿名结构体名字**。

5. 如果struct嵌套了一个有名结构体，这种模式叫**组合**。如果组合关系，那么在访问组合的结构体字段或方法时，必须带上结构体名字。

```go
type A struct {
    Name string
    Age int
}

type C struct {
    a A     //有名结构体
}


C.Name = "tom" //报错
```

6. 嵌套匿名结构体后，也可以在创建结构体变量(实例)，**直接指定各个匿名结构体字段的值**。

```go
type Goods struct {
    Name  string
    Price float64
}


type Brand struct {
    Name string
    Addr string
}

type TV struct {
    Goods
    Brand
}

type TV2 struct {
    *Goods
    *Brand
}


func main() {
    tv := TV{
             Goods{"海尔", 3798.98},
             Brand{"海尔", "青岛"},
            }

    tv2 := TV2{
          &Goods{"海尔", 3798.98},
          &Brand{"海尔", "青岛"},
          }
    fmt.Println(*tv2.Goods, *tv2.Brand)
}
```

注意：

结构体还能添加基本数据类型

```go
type E struct {
    int
}
 E.int = 20
```

但是有一个int的匿名字段 ，就不能有第二个

如果需要多个int字段，需要指定名称

### 3.4 多重继承

结构体嵌入2个或多个匿名结构体，该结构体可以直接访问嵌套匿名结构体的字段和方法，从而实现多重继承。

- 嵌入匿名结构体有相同字段或方法，访问时需要通过匿名结构体类型名区分

- 为了代码简洁，尽量不使用多重继承

## 4. 多态

### 4.1基本介绍

变量(实例)具有多种形态。面向对象第三大特性，多态特征通过接口实现。可以按照统一的接口来调用不同的实现。此时接口变量就呈现不同形态。

- 快速入门见十二章Usb实例

- 接口体现多态特征
1. 多态参数

        在Usb接口案例中，usb Usb可以接收手机变量，也可以接收Camera变量，体现了Usb接口多态

2. 多态数组

        案例：给Usb数组中，存放Phone结构体和Camera结构体变量，Phone还有一个特有方法call()，请遍历Usb数组，如果是Phone变量，除了调用Usb接口声明，还需要调用Phone特有方法call

```go
func main() {
    //声明一个Usb接口数组，可以存放Phone和Camera结构体变量
    //体现多态数组
    var UsbArray [3]Usb
    UsbArray[0] = Phone{"vivo"}
    UsbArray[1] = Camera{"Canno"}
}
```

# 第十二章 接口(interface)

```go
package main

import (
    "fmt"
)

// 声明一个接口
type Usb interface {
    //声明了两个没有实现的方法
    Start()
    Stop()
}

type MobilePhone struct {
}

type Camera struct {
}

type Computer struct {
}

// MobilePhone实现Usb接口的方法
func (p MobilePhone) Start() {
    fmt.Println("手机启动...")
}

func (p MobilePhone) Stop() {
    fmt.Println("手机关闭...")
}

// Camera实现Usb接口的方法
func (p Camera) Start() {
    fmt.Println("相机启动...")
}

func (p Camera) Stop() {
    fmt.Println("相机关闭...")
}

// 给Computer编写一个方法working
// 所谓实现Usb接口，指实现Usb接口声明的所有方法
func (c Computer) Working(usb Usb) {
    //通过usb接口变量调用start和stop方法
    usb.Start()
    usb.Stop()
}

func main() {
    computer := Computer{}
    phone := MobilePhone{}
    camera := Camera{}

    //关键点
    computer.Working(phone)
    computer.Working(camera)
}
```

## 1. 基本介绍

interface类型可以定义一组方法，但是不需要实现。并且interface不能包含任何变量。到某个自定义类型(比如结构体Phone)要使用时，根据情况实现方法。

## 2. 基本语法

```
type 接口名 interface {
    method1(参数列表)返回值列表
    method2(参数列表)返回值列表
    ...
}

//实现接口所有方法
func (t 自定义类型)method1(参数列表)返回值列表{
    代码块
}

func (t 自定义类型)method2(参数列表)返回值列表{
    代码块
}
...
```

- 小结

（1）接口里的所有方法都没有方法体，即接口的方法都没有实现。体现了程序设计<font face="黑体" color=red>多态和高内聚低耦合</font>的思想。

(2) Go中的接口，不需要显式实现。只要一个变量，含有接口中所有方法，这个变量就实现了这个接口。因此，<font color="red">Go没有implement这样的关键字</font>。

## 3.注意事项和细节

1. 接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量(实例)

```go
    var stu Student //结构体变量实现了Say() 实现了AintFace
    var a AintFace = stu
    a.Say()
```

1. 接口中所有的方法都没有方法体(代码块)，都是没有实现的方法

2. 在Go中，一个自定义类型需要将某个接口的所有方法都实现，称这个自定义类型实现了该接口

3. 一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋给接口类型

4. 一个自定义数据类型，就可以实现接口，不仅仅是struct

```go
type integer int

func (i integer) Say() {
    fmt.Println("interger say i = ", i)
}
```

1. 一个自定义类型可以实现多个接口

```go
type Bintface interface {
    Hello()
}

func (m Monster) Say() {
    fmt.Println(m, "Say")
}

func (m Monster) Hello() {
    fmt.Println("Hello")
}

 //Monster实现了Aintface和Bintface
var monster Monster

var A AintFace = monster
var B Bintface = monster
A.Say()
B.Hello()
```

1. <font color="red">Go接口中不能有任何变量</font>

2. 一个接口(比如A接口)可以继承多个别的接口(比如B,C接口)，如果实现A接口，必须把B,C接口的方法也实现

```go
package main

import "fmt"

type AintFace interface {
    test01()
}

type Bintface interface {
    test02()
}

type Cintface interface {
    AintFace
    Bintface
    test03()
}

// 如果实现C接口，需要将A,B接口的方法都实现
type Stu struct{}

func (stu Stu) test01() {
    fmt.Println("test01")
}

func (stu Stu) test02() {
    fmt.Println("test02")
}

// func (stu Stu) test03() {
//     fmt.Println("test03")
// }

func main() {
    var stu Stu
    var a Cintface = stu 
//cannot use stu (variable of type Stu) as Cintface value in 
//variable declaration: Stu does not implement Cintface 
//(missing method test03)
    a.test01()
    a.test02()
    // a.test03()
}
```

1. interface默认是一个指针，如果没有对interface初始化就使用，会输出nil

2. 空接口interface{} 没有任何方法，所有类型都实现了空接口。

## 4.实践

- 实现对Hero结构体切片的排序：sort.Sort(data interface)

```go
package main

import (
    "fmt"
    "math/rand"
    "sort"
)

func bubble_sort(slice []int) []int {
    var temp int
    for i := 0; i < len(slice); i++ {
        for j := i + 1; j < len(slice); j++ {
            if slice[i] < slice[j] {
                temp = slice[i]
                slice[i] = slice[j]
                slice[j] = temp
            }
        }
    }
    return slice
}

type Hero struct {
    Name string
    age  int
}
type HeroSlice []Hero

// 实现接口
func (hs HeroSlice) Len() int {
    return len(hs)
}

// Less方法决定使用什么标准进行排序
// 标准是按照年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
    if hs[i].age < hs[j].age {
        return true
    } else {
        return false
    }
}

func (hs HeroSlice) Swap(i, j int) {
    // temp := hs[i]
    // hs[i] = hs[j]
    // hs[j] = temp
    hs[i], hs[j] = hs[j], hs[i]
}

func main() {
    //定义一个数组/切片
    //var arr [5]int
    var slice = []int{0, -20, -10, 2, 98}

    //对slice进行冒泡排序
    // bubble_sort(slice)
    // fmt.Println(slice)

    sort.IntSlice.Sort(slice)
    fmt.Println(slice)

    //对结构体切片进行排序
    var heros HeroSlice
    for i := 0; i < 10; i++ {
        hero := Hero{
            Name: fmt.Sprintf("Hero~%d", rand.Intn(100)),
            age:  rand.Intn(100),
        }
        heros = append(heros, hero)
    }
    //排序前
    for _, val := range heros {
        fmt.Println(val)
    }

    //排序后
    sort.Sort(heros)
    fmt.Printf("-----------排序后-----------\n")
    //fmt.Println(heros)
    for _, val := range heros {
        fmt.Println(val)
    }
}
```

## 5. 接口 vs 继承

```go
package main

import (
    "fmt"
)

// 猴子不仅能继承老猴子的技能
// 还可以通过接口实现鸟的飞行和鱼的游泳
// 定义Monkey struct
type Monkey struct {
    Name string
}

func (m *Monkey) climb() {
    fmt.Println(m.Name, "生来会爬树")
}

type Bird interface {
    fly()
}

type Fish interface {
    Swim()
}

func (m *Monkey) fly() {
    fmt.Println(m.Name, "学习后会飞")
}

func (m *Monkey) Swim() {
    fmt.Println(m.Name, "学习后游泳")
}

// 定义little_monkey struct
type LittleMonkey struct {
    Monkey
}

func main() {
    //创建Little_Monkey实例
    var monkey LittleMonkey
    monkey.Name = "小猴子"
    monkey.climb()
    monkey.fly()

}
```

接口不会破坏原来继承的关系。

1. 当A结构体继承B结构体，就自动继承B的字段和方法，可以直接使用

2. A需要扩展功能，同时不希望破坏继承关系，可以使用接口实现

3. 可以认为，接口是对继承的补充

### 5.1 继承vs接口

- 解决问题不同
  
  - 继承价值在于：解决代码**复用性和可维护性**
  
  - 接口价值在于：设计好规范(方法)，让其他自定义类型去实现

- 接口比继承更加灵活
  
  - 继承是is - a的关系，接口是like - a的关系

- 接口在一定程度上实现了代码解耦

## 6.类型断言

![](C:\Users\12143\AppData\Roaming\marktext\images\2023-03-17-21-51-20-image.png)

- b=a.(Point)就是类型断言，表示判断a是否指向Point类型的变量，如果是就转成Point类型并赋值给b变量，否则报错。

### 6.1 基本介绍

类型断言，由于接口是一般类型，不知道具体类型，如果转成具体类型，就需要使用类型断言。如下：

```go
    var t float32
    var x interface{}
    x = t
    //转成float，带检查
    _, ok := x.(float32)
    if ok == true {
        fmt.Print("convert success")
    } else {
        fmt.Println("convert fail")
    }
```

代码说明：

- 在进行类型断言时，类型不匹配会报错panic

- 因此进行类型断言时，确保原来空接口指向的就是断言的类型

### 6.2 最佳实践

1. 对11.4.1中多态数组改进

 案例：给Usb数组中，存放Phone结构体和Camera结构体变量，Phone还有一个特有方法call()，请遍历Usb数组，如果是Phone变量，除了调用Usb接口声明，还需要调用Phone特有方法call

```go
//Phone结构体增加特有方法call(), 当Usb接口接收的时Phone变量时, 还需调用call方法
func (p Phone) call() {
    fmt.Println(p.Name, "打电话")
}

type Computer struct {

}

func (computer Computer) Working(usb Usb) {
    usb.Start()
    //如果usb指向Phone，还需要调用Call方法
    //类型断言
    if phone, ok := usb.(Phone); ok == true{
        phone.call()
    }
    usb.Stop()
}

func main() {
    //声明一个Usb接口数组，可以存放Phone和Camera结构体变量
    //体现多态数组
    var UsbArray [3]Usb
    UsbArray[0] = Phone{"vivo"}
    UsbArray[1] = Camera{"Canno"}

    //调用call--》类型断言
    for key, val := range(UsbArray) {
        computer.Working(v)
    }

}
```

2. 循环判断传入参数的类型

```go
func TypeJudge(items ...interface{}) {// ...表示任意多个参数
    for i, x := range items {
        switch x.(type) {
        case bool:
            fmt.Printf("param #%d is a bool, 值是%v\n", i, x)
        case float64, float32:
            fmt.Printf("param #%d is a float64, 值是%v\n", i, x)
        case uint, uint8, uint16, uint32, uint64:
            fmt.Printf("param #%d is a uint, 值是%v\n", i, x)
        case int, int8, int16, int32, int64:
            fmt.Printf("param #%d is a int, 值是%v\n", i, x)
        case string:
            fmt.Printf("param #%d is a string, 值是%v\n", i, x)
        case nil:
            fmt.Printf("param #%d is a nil, 值是%v\n", i, x)
        default:
            fmt.Printf("param #%d is unknown, 值是%v\n", i, x)
        }
    }
}
```

# 第十三章 项目开发

## 1.项目p227

编写文件TestMyAccount.go完成基本功能，要求：

1.1 完成显示主菜单，且能退出

1.2 完成显示明细的功能

思路：

(1) 需要明示明细，需要定义一个变量details string来记录

(2)需要定义变量记录余额rest、每次收支的金额money，每次收支的说明note

1.3 完成登记收入的功能

1.4 完成登记支出的功能

2. 用面向对象的方法实现

2.1 把功能封装到结构体中，通过调用结构体实现记账和收支。

2.2 在main中创建这个结构体的实例，实现功能的实现

## 2.客户信息管理系统

- View.go
  1.显示界面
  2.可以接收用户的输入
  3.根据用户输入完成客户管理[增删改查]
  
  - 编写方法list，调用CustomerService的List方法，并显示客户列表
  
  - add方法调用CustomerService的Add方法
  
  - del方法调用CustomerService的DEL方法

- Service.go(业务逻辑)
  1.完成对用户的各种操作
  2.对用户增删改查
  3.声明customer的切片
  
  - 显示客户列表
    编写方法List[返回客户信息(切片)]
    编写NewCustomerService返回CustomerService实例
  
  - 添加客户
    编写方法Add, 将新客户加入到customers切片中
  
  - 删除客户
    编写方法Del(Id int)删除一个客户
    编写方法FindbyId(Id int), 返回Id号对应的切片的下标

- model.go
  1.表示一个客户
  2.客户各种字段
  
  - customer表示一个客户信息，因此必须含有
    客户的各个字段
    type Customer struct {
    Id int
    Name string
    ...
    }

<font color="red">从下往上编程，即model-》service-》view</font>

# 第十四章 文件操作

## 1.文件的基本介绍

文件是数据源的一种。最主要的作用是保存数据。

- 文件在程序中以流的形式操作

![](C:\Users\12143\AppData\Roaming\marktext\images\2023-03-19-16-57-53-image.png)

**流：**

数据在数据流(文件)和程序(内存)之间经历的路径。

输入流：数据从文件到内存的路径

输出流：数据从内存到文件的路径

- os.File封装所有文件相关操作，File是一个结构体

## 2.打开和关闭文件

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // 打开一个文件
    /*
        概念说明: file的叫法
        1.file对象
        2.file指针
        3.file 文件句柄
    */
    file, err := os.Open("F:/code/jin.txt")
    if err != nil {
        fmt.Println("Opne file err: ", err)
    }

    fmt.Printf("file is \n%v", file)
    //关闭文件
    err = file.Close()
    if err != nil {
        fmt.Println("Close file err: ", err)
    }
}
```

## 3. 读文件

### 3.1 应用实例

- 读文件的内容并显示在终端(**带缓冲区的方式**)os.Open, file.Close, bufio.NewReader(), reader.ReadString函数和方法

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func main() {
    // 打开一个文件
    /*
        概念说明: file的叫法
        1.file对象
        2.file指针
        3.file 文件句柄
    */
    file, err := os.Open("F:/code/jin.txt")
    if err != nil {
        fmt.Println("Opne file err: ", err)
    }

    //关闭文件
    defer file.Close() //要及时关闭file句柄，否则会造成内存泄漏

    // 创建一个*Reader, 是带缓冲的
    //默认缓冲区是4096个字节
    /*const (
    defaultBufSize = 4096
    )
    */
    reader := bufio.NewReader(file)
    //循环读取文件内容
    for {
        info, err := reader.ReadString('\n')
        if err == io.EOF {
            fmt.Print(info)
            break
        } else if err != nil {
            fmt.Println("读取错误, err: ", err)
        }
        //输出内容
        fmt.Print(info)
    }
    fmt.Println("\n文件读取结束")
}
```

- 读文件的内容并显示在终端(**使用ioutil一次将整个文件读入内存**)，适合文件不大的情况。相关方法和函数(ioutil.ReadFile)

```go
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    // ioutil
    file := "F:/code/jin.txt"
    info, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println("读取错误, err: ", err)
    }
    //info此时是[]byte类型
    fmt.Printf("%v", string(info))
    // 没有显式地Open文件，因此不需要显式关闭文件
}
```

## 4. 写文件

### 4.1 基本介绍

func OpenFile(name string, flag int, perm FileMode) (file *File, err error)

说明：os.OpenFile是一个更一般性的文件打开函数，会使用指定的选项(如O_EDONLY等)、指定的模式(如0666等)打开指定名称的文件。

如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError

- 第一个参数name：文件名

- 第二个参数flag：文件打开模式(可以组合)

- 第三个参数perm：Linux文件权限

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    /*
        1.创建1个新文件，写入5句"hello, go"
        2.打开1个存在的文件，将原来的内容覆盖成新的语句"学而时习之, 不亦乐乎"
        3.打开1个存在的文件，在原来的内容下追加"吾日三省吾身"
        4.打开1个存在的文件，将原来的内容显示在终端，并且追加五句"君子不器"
        使用os.OpenFile(), bufio.NewWriter(), *Writer的方法WriteString完成
    */

    //1.创建1个新文件，写入5句"hello, go"
    file := "F:/code/test.txt"
    info, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
    // 2.打开1个存在的文件，将原来的内容覆盖成新的语句"学而时习之, 不亦乐乎"
    // 改为info, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0666)

    // 3.打开1个存在的文件，在原来的内容下追加"吾日三省吾身"
    // info, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0666)

    // 4.打开1个存在的文件，将原来的内容显示在终端，并且追加五句"君子不器"
    // info, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND, 0666)
    // 读取文件内容，输出到终端
    // reader := bufio.NewReader(info)
    // for {
    //     str, err := reader.ReadString('\n')
    //     fmt.Print(str)
    //     if err == io.EOF {
    //         break
    //     }
    //     fmt.Print(str)
    // }

    if err != nil {
        fmt.Println("Opne file err: ", err)
        return
    }
    // 及时关闭file句柄
    defer info.Close()
    str := "hello, go\n"
    // 写入时使用带缓存的*Writer
    writer := bufio.NewWriter(info)
    for i := 0; i < 5; i++ {
        writer.WriteString(str)
    }
    // 因为writer带缓存，因此在调用WriteString方法时
    // 内容是先写入缓存，所有需要调用Flush方法，将缓存内容真正写入文件中
    // 否则文件中会丢失数据
    writer.Flush()



}
```

**重要**

// 内容是先写入缓存，所有需要调用Flush方法，将缓存内容真正写入文件中
 // 否则文件中会丢失数据
 writer.Flush()

- 实例2

编写一个程序，将一个文件内容写入另一个文件。注：这2个文件已经存在。

说明：

（1）使用ioutil.ReadFile / ioutil.WriteFile

```go
    src_content, err := ioutil.ReadFile(src_file)
    if err != nil {
        fmt.Println("Opne file err: ", err)
        return
    }
    err = os.WriteFile(dst_file, src_content, 0666)
    if err != nil {
        fmt.Println("Write file err: ", err)
        return
    }
```

### 4.2 判断文件是否存在

Go判断<font color="red">文件或文件夹是否存在</font>的方法是使用os.Stat()函数返回的错误值进行判断：

1. 如果返回的错误是nil，说明文件或文件夹存在

2. 如果返回的错误类型使用os.IsNotExist()判断为true，说明文件或文件夹不存在

3. 如果返回的错误是其他类型，则不确定是否存在

```go
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    } else if os.IsNotExist(err) {
        return false, nil
    } else {
        return false, nil
    }
}
```

### 4.3 拷贝文件

- 将一张图片拷贝到另一个目录下，如果不存在则创建

- io包func Copy

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

// 编写函数，接收两个文件路径src_path, dst_path
func CopyFile(src_path string, dst_path string) (written int64, err error) {
    src_file, err1 := os.Open(src_path)
    if err1 != nil {
        fmt.Println("Open file error: ", err)
    }
    defer src_file.Close()

    reader := bufio.NewReader(src_file)

    // 打开dst_path
    dst_file, err2 := os.OpenFile(dst_path, os.O_WRONLY|os.O_CREATE, 0666)
    if err2 != nil {
        fmt.Println("Open file error: ", err)
        return
    }
    defer dst_file.Close()

    writer := bufio.NewWriter(dst_file)
    return io.Copy(writer, reader)
}
```

### 4.4 统计不同类型的字符个数

- 统计英文、数字、空格和其他字符数量

- 说明：统计一个文件中英文、数字、空格和其他字符数量

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

// 编写函数，接收两个文件路径src_path, dst_path
func CopyFile(src_path string, dst_path string) (written int64, err error) {
    src_file, err1 := os.Open(src_path)
    if err1 != nil {
        fmt.Println("Open file error: ", err)
    }
    defer src_file.Close()

    reader := bufio.NewReader(src_file)

    // 打开dst_path
    dst_file, err2 := os.OpenFile(dst_path, os.O_WRONLY|os.O_CREATE, 0666)
    if err2 != nil {
        fmt.Println("Open file error: ", err)
        return
    }
    defer dst_file.Close()

    writer := bufio.NewWriter(dst_file)
    return io.Copy(writer, reader)
}

// 定义1个结构体，用于保存统计类型的数目
type Statics struct {
    Letter_num int
    Num        int
    Space_num  int
    Other_num  int
}

func main() {
    /*
      思路:
      打开一个文件，创建一个reader
      每读取一行，就去统计改行有多少个英文、数字、空格和其他字符
      将结果保存在一个结构体中
    */
    src_path := "F:/code/self_intro.txt"
    src_file, err := os.Open(src_path)
    if err != nil {
        fmt.Println("Open file error: ", err)
    }
    defer src_file.Close()

    reader := bufio.NewReader(src_file)

    // 定义Statics实例
    var count Statics
    for {
        str, err1 := reader.ReadString('\n')
        if err1 == io.EOF {
            break
        }

        // str2 := []rune(str)
        // 遍历str统计
        for _, v := range str {
            switch {
            case v >= 'a' && v <= 'z':
                fallthrough
            case v >= 'A' && v <= 'Z':
                count.Letter_num++
            case v == ' ' || v == '\t':
                count.Space_num++
            case v >= '0' && v <= '9':
                count.Num++
            default:
                count.Other_num++
            }
        }
    }

    fmt.Printf("字符个数为: %v, 数字个数为: %v, 空格个数为: %v, 其他字符个数为: %v", count.Letter_num,
        count.Num, count.Space_num, count.Other_num)
}
```

## 5. 命令行参数

### 5.1 基本介绍

- 希望能够获得命令行输入的各种参数

- 基本介绍：os.Args是一个string的切片，用来存储所有的命令行参数

### 5.2 应用实例

编写代码获取命令行各个参数

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("命令行参数有: ", len(os.Args))

    // 遍历os.Args切片，就可以得到所有的命令行输入参数值
    for k, v := range os.Args {
        fmt.Printf("args[%v]是%v\n", k, v)
    }
}
```

### 5.3 flag包解析命令行参数

- 5.2 是比较原生的方式，对解析参数不是特别方便，特别是带有指定参数形式的命令行。
  
  比如：cmd>main.exe -f c:/aaa.txt -p 200 -u root这样形式命令行，go提供了flag包，可以方便解析命令行参数，<font color="red">而且参数顺序可以调整</font>

- 应用实例

```go
package main

import (
    "flag"
    "fmt"
)

// cmd>main.exe -f c:/aaa.txt -p 200 -u root

func main() {
    // 定义变量用于接收命令行参数值
    var user string
    var pwd string
    var host string
    var port int

    // &user就是接收命令行输入的-u后的参数值
    // "u" 接收-u 后面的指定参数
    // "" 默认值
    // "用户名, 默认为空" 说明
    flag.StringVar(&user, "u", "", "用户名, 默认为空")
    flag.StringVar(&pwd, "p", "", "密码, 默认为空")
    flag.StringVar(&host, "h", "localhost", "用户名, 默认为localhost")
    flag.IntVar(&port, "port", 3306, "用户名, 默认为3306")

    // 重要操作, 转换, 必须调用该操作
    flag.Parse()

    // 输出结果
    fmt.Printf("user=%v, pwd=%v, host=%v, port=%v\n", user, pwd, host, port)
}
```

## 6. json

### 6.1 基本介绍

- 概述
  
  JSON(JavaScript Object Notation)是一种轻量级的数据交换格式。易于人阅读和编写。同时也易于机器解析和生成。
  
  JSON易于机器解析和生成，并有效提升网络传输效率，通常程序在**网络传输时会先将数据(结构体、map等)序列化为json字符串，到接收方得到json字符串时，再反序列化恢复成原来的数据类型(结构体、map等)。**
  
  ![](C:\Users\12143\AppData\Roaming\marktext\images\2023-03-20-22-21-41-image.png)

### 6.2 json数据格式说明

**js语言中，一切皆对象**。任何支持的类型都可以通过json表示，例如字符串、数字、对象、数组等。

json键值对是用来保存 数据的方式。

键/值对组合中的键名写在前面并用双引号""包裹，使用：分离，然后是值：

如：

{"firstName": "json"}, {"name":"tom", "age":"18", address:["北京", "上海"]}

[{"name":"tom", "age":"18", address:["北京", "上海"]},

{"name":"mary", "age":"18", address:["南京", "上海"]}]

### 6.3 json数据在线解析

https://www.json.cn 可以验证json数据书写是否正确。

### 6.4 json序列化

- 介绍
  
  json序列化是指，将key-value结构的数据类型(结构体、map等)序列化为json字符串的操作。

- 案例: 结构体、map、切片的序列化

```go
package main

import (
    "encoding/json"
    "fmt"
)

// 演示将结构体、map和切片序列化
type Monster struct {
    Name  string
    Age   int
    Birth string
    Sal   string
    Skill string
}

func testStruct() {
    monster := Monster{
        Name:  "牛魔王",
        Age:   5000,
        Birth: "十万年前",
        Sal:   "5000",
        Skill: "法天象地",
    }

    // 序列化
    data, err := json.Marshal(&monster)
    if err != nil {
        fmt.Printf("序列化失败 err=%v\n", err)
    }
    fmt.Printf("序列化结果是%v\n", string(data))
}

// 序列化map
func testMap() {
    var a map[string]interface{}
    // 使用map前需要先make
    a = make(map[string]interface{})
    a["name"] = "红孩儿"
    a["age"] = 2000
    a["addr"] = "火云洞"

    data, err := json.Marshal(a)
    if err != nil {
        fmt.Printf("序列化失败 err=%v\n", err)
    }
    fmt.Printf("序列化结果是%v\n", string(data))
}

// 切片序列化
func testSlice() {
    var slice []map[string]interface{}
    var m1 map[string]interface{}

    // 使用map前先make
    m1 = make(map[string]interface{})
    m1["name"] = "孙悟空"
    m1["age"] = 5000
    m1["addr"] = "花果山"
    //slice = append(slice, m1)

    var m2 map[string]interface{}
    m2 = make(map[string]interface{})
    m2["name"] = "猪八戒"
    m2["age"] = 5200
    m2["addr"] = [2]string{"云栈洞", "高老庄"}

    slice = append(slice, m1, m2)

    data, err := json.Marshal(slice)
    if err != nil {
        fmt.Printf("序列化失败 err=%v\n", err)
    }
    fmt.Printf("序列化结果是%v\n", string(data))
}

// 基本数据类型序列化
func testFloat64() {
    var num1 float64 = 98.237937
    data, err := json.Marshal(num1)
    if err != nil {
        fmt.Printf("序列化失败 err=%v\n", err)
    }
    fmt.Printf("序列化结果是%v\n", string(data))
}

func main() {
    testStruct()
    testMap()
    testSlice()
    testFloat64()
}
```

### 6.5 json反序列化

json反序列化是指将json字符串反序列化恢复成原来的数据类型(结构体、map等)

- 案例

```go
// 演示将结构体、map和切片序列化
// 使用tag
type Monster struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Birth string
    Sal   string
    Skill string
}
```

**对于结构体的序列化，如果希望序列化后的key的名字由自己指定，可以给结构体加tag标签，如上。**

```go
package main

import (
    "encoding/json"
    "fmt"
)

// 演示将结构体、map和切片序列化
type Monster struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Birth string
    Sal   string
    Skill string
}

// 将json字符串反序列化成struct
func unMarshalStruct() {
    str := "{\"name\":\"牛魔王\",\"age\":5000,\"Birth\":\"十万年前\",\"Sal\":\"5000\",\"Skill\":\"法天象地\"}"
    // 真实情况下str是由网络传输得到

    // 声明monster实例
    var monster Monster
    err := json.Unmarshal([]byte(str), &monster)
    if err != nil {
        fmt.Printf("反序列化失败 err=%v\n", err)
    }
    fmt.Printf("反序列化后的结果是%v\n", monster)
}

// 将json字符串反序列化成map
func unMarshalMap() {
    str := "{\"addr\":\"火云洞\",\"age\":2000,\"name\":\"红孩儿\"}"

    var a map[string]interface{}
    a = make(map[string]interface{})
    err := json.Unmarshal([]byte(str), &a)
    if err != nil {
        fmt.Printf("反序列化失败 err=%v\n", err)
    }
    fmt.Printf("反序列化后的结果是%v\n", a)
}

func unMarshalSlice() {
    str := "[{\"addr\":\"花果山\",\"age\":5000,\"name\":\"孙悟空\"},{\"addr\":[\"云栈洞\",\"高老庄\"],\"age\":5200,\"name\":\"猪八戒\"}]"

    var slice []map[string]interface{}
    err := json.Unmarshal([]byte(str), &slice)
    if err != nil {
        fmt.Printf("反序列化失败 err=%v\n", err)
    }
    fmt.Printf("反序列化后的结果是%v\n", slice)
}

func main() {
    unMarshalStruct()
    unMarshalMap()
    unMarshalSlice()
}
```

- 小结
  
  1. 反序列化1个json字符串时，**确保反序列化后的数据类型和原序列化前的数据类型一致**(字段完全一样也可以)
  
  2. 如果json字符串是通过程序获取到的，则不需要再进行转义处理(\")因为那时转义已经处理好了

# 第十五章 单元测试

如何确认一个函数、一个模块是否正确。

- 传统方法
  
  在main函数中，调用addUpper函数，看实际输出结果是否和预期结果一致，如果一致，则说明函数正确；否则函数有错误需要修改。
  
  - 缺点分析
    
    1. 不方便，在main函数中测试出问题需要停止main函数，可能去停止项目。
    
    2. 不利于管理，当测试多个函数或模块时，都需要写在main函数中，不利于管理。
    
    3. 引出单元测试。---->testing测试框架

## 1. 基本介绍

Go自带轻量级测试框架testing和自带go test命令来实现单元测试和性能测试，testing框架和其他语言测试框架类似，可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例。通过单元测试可以解决如下问题:

1. 确保每个函数可以运行且结果正确

2. 确保代码性能良好

3. 单元测试能及时发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决，性能测试重点在于发现程序设计上的一些问题，让程序能够在高并发的情况下还能保持稳定。

## 2.快速入门

使用Go的单元测试，测试addUpper和sub函数

- 测试用例文件名必须以<font color="red">_test.go</font>结尾。比如cal_test.go，cal不是固定的

- 测试用例函数必须以<font color="red">Test</font>开头，一般是Test+被测试函数名。比如TestAddUpper()

- TestAddUpper(t testing.T)的形参必须是testing.T

- 一个测试用例文件可以有多个被测函数

- 运行指令
  
  - go test [如果运行正确，无日志，错误会输出日志]
  
  - go test -v [运行正确，错误都会输出日志]

- 当出现错误是，可以用t.Fatalf格式化输出错误信息，并退出程序

- t.Logf方法输出相应日志

- 测试用例函数不放在main函数中也执行

- PASS表示测试用例运行成功，FAIL表示运行失败

- 测试单个文件，一定要带上被测试的原文件
  
  - go test -v cal_test.go cal.go

- 测试单个方法
  
  - go test -v -run TestAddUpper

# 第十六章 goroutine和channel

## 1. goroutine

## 1.1 进程和线程说明

1. 进程是程序在操作系统中的一次执行过程，是系统进行资源分配和调度的基本单位

2. 线程是进程的一个执行实例，是程序执行的最小单元，是比进程更小的能独立运行的基本单位

3. 一个进程可以创建和销毁多个线程，同一个进程的多个线程可以并发执行

4. 一个程序至少一个进程，一个进程至少一个线程

### 1.2 并发和并行

1. 多线程程序在单核上运行，就是并发

2. 多线程程序在多核上运行，就是并行

并发：在1个CPU上，有10个线程，每个线程执行10ms（轮询操作），从某个时间点来看其实只有一个线程在执行，这就是并发。

并行：在多个CPU上，比如10个CPU，有10个线程，每个线程执行10ms（在不同CPU上执行），从某个时间点来看10个线程在执行，这就是并行。

### 1.3 Go协程goroutine和Go主线程

1. Go主线程：一个Go线程上，可以起多个协程，协程可以理解为轻量级的线程

2. **Go线程的特点**
   
   1. 独立栈空间
   
   2. 共享程序堆空间
   
   3. 调度由用户控制
   
   4. 协程是轻量级的线程

### 1.4 快速入门

编写程序完成如下功能：

1. 在主线程中开启goroutine，改协程每隔1秒输出"hello world"

2. 在主线程每隔1秒输出"hello golang"，输出10次后退出程序

3. 要求主线程和goroutine同时执行

4. 画出主线程和协程执行流程图

```go
func test() {
    for i := 1; i < 10; i++ {
        fmt.Println("test(), hello world" + strconv.Itoa(i))
        time.Sleep(time.Second)
    }
}

func main() {
    go test() //开启一个协程
    for i := 1; i < 10; i++ {
        fmt.Println("main(), hello golang" + strconv.Itoa(i))
        time.Sleep(time.Second)
    }
}
```

![](C:\Users\12143\AppData\Roaming\marktext\images\2023-03-22-22-10-43-image.png)

输出说明主线程和协程同时执行

![](C:\Users\12143\AppData\Roaming\marktext\images\2023-03-22-22-16-01-image.png)

**小结**：

1. 主线程是一个物理线程，直接作用在CPU上。是重量级的，十分耗费CPU资源

2. 协程是从主线程开启的，是轻量级的线程，是逻辑态。对资源消耗相对小

3. Go的协程机制是重要特点，可以轻松开启上万协程。其他语言开发机制一般基于线程，开启过多线程资源耗费大。凸显Golang的并发优势。

### 1.5 goroutine的调度模型(信息需要补全)

MPG模式的基本介绍

1. M：操作系统的主线程

2. P：协程执行需要的上下文

3. G：协程

**MPG模式运行的状态**

### 1.6 设置Golang运行的CPU数(runtime包)

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    // 获取当前系统CPU数目
    num := runtime.NumCPU()

    runtime.GOMAXPROCS(num)
    fmt.Println("num=", num)
}
```

go1.8后默认CPU数目最大

## 2.channal(管道)

### 2.1 问题引入

需求：

计算1-200的阶乘，并且把每个数的阶乘放入map中。最后显示出来，使用goroutine完成

使用goroutine会出现并发/并行安全问题

![](C:\Users\12143\AppData\Roaming\marktext\images\2023-03-22-22-44-37-image.png)

在运行程序时，如何知道是否存在资源竞争问题。在编译程序时增加参数 -race

#### 不同goroutine之间如何通讯

1. 全局变量加锁同步

2. channal

```go
package main

import (
    "fmt"
    "sync"
)

// 1.编写函数计算阶乘， 放入map
// 2.启动多个协程，将结果放入map
// 3.map应该是全局的

var (
    myMap = make(map[int]int, 10)
    // 声明一个全局互斥锁
    // sync 是包, syncchornized
    // Mutex: 互斥
    lock sync.Mutex
)

func Step(n int) map[int]int {
    res := 1
    for i := 1; i <= n; i++ {
        res *= i
    }

    lock.Lock()
    myMap[n] = res
    lock.Unlock()

    return myMap
}

func main() {
    for i := 1; i <= 80; i++ {
        go Step(i)
    }

    // 读的时候加锁原因
    // 但主线程并不知道，因此底层仍然可能出现资源争夺，因此加入互斥锁即可解决问题
    lock.Lock()
    for i, v := range myMap {
        fmt.Printf("map[%d]=%d\n", i, v)
    }
    lock.Unlock()
}
```

### 2.2 channel的介绍

1. channel本质是一个数据结构-队列

2. 数据先进先出

3. 线程安全，多goroutine访问时，不需要加锁，即channel本身就是线程安全的

4. channel是有类型的，一个string的channel只能存放string类型数据。

### 2.3 定义/声明

var 变量名 chan 数据类型

举例：

var intChan chan int (intChan用于存放int数据)

var mapChan chan map[int]string (mapChan用于存放map[int]string类型)

var perChan chan Person

var perChan2 chan *Person

...

说明

1. channel是引用类型

2. channel必须初始化才能写入数据, 即make后才能使用

3. 管道是有类型的，intChan只能写入整数int

### 2.4 基本使用

1. channel初始化
   
   使用make进行初始化
   
   var intChan chan int
   
   intChan = make(chan int, 10)

2. 向channel中写入(存放)数据
   
   var intChan chan int
   
   intChan = make(chan int, 10)
   
   num := 999
   
   intChan <- 10
   
   intChan <- num

**注意:**

管道写入数据不能超过容量

```go
package main

import "fmt"

func main() {
    // 演示管道的使用
    var intChan chan int
    intChan = make(chan int, 3)

    //2. 看看intChan是什么
    // intChan的值是0xc00001e180, intChan本身的值是0xc00000a028
    fmt.Printf("intChan的值是%v, intChan本身的值是%p\n", intChan, &intChan)

    // 3.向管道写入数据
    intChan <- 10
    num := 211

    intChan <- num
    <-intChan

    // 4.查看管道的长度和容量
    // 管道写入数据不能超过容量
    fmt.Printf("channel len is %v, cap is %v\n", len(intChan), cap(intChan))

    // 5.从channel读取数据
    var num2 int
    num2 = <-intChan
    fmt.Println(num2)
    fmt.Printf("channel len is %v, cap is %v\n", len(intChan), cap(intChan))

    // 6.在没有协程的情况下，若管道数据全部取出，再取就会报告deadlock

}
```

在没有协程的情况下，若管道数据全部取出，再取就会报告deadlock

### 2.5 注意事项

1. channel只能存放指定的数据类型

2. 管道写入数据不能超过容量

3. 从channel取出数据后，可以继续放入

4. 在没有协程的情况下，若管道数据全部取出，再取就会报告deadlock

```go
// 创建一个allChan，最多可放10个任意数据类型变量

var allChan chan interface{} // 声明一个空接口即可

// 举例
func main() {
    var allChan chan interface{}
    allChan = make(chan interface{}, 10)

    cat1 := Cat{Name: "tom", Age: 18}
    cat2 := Cat{Name: "bob", Age: 8}
    allChan <- cat1
    allChan <- cat2
    allChan <- 10

    //取出
    cat := <-allChan
    // cat的类型未知，需要使用断言
    // fmt.Println(cat.Name) //如果未使用断言，cat.Name undefined (type interface{} has no field or method Name)
    a := cat.(Cat)
    fmt.Println(a.Name)
}
```

### 2.6 channel的遍历和关闭

- channel的关闭
  
  使用内置函数close可以关闭channel，当channel关闭后，就不能再向channel写入数据，但还可以从channel中读取数据

- channel的遍历
  
  channel支持for-range的方式进行遍历，注意
  
  1. 遍历时，channel如果没有关闭，则出现deadlock错误
  
  2. 遍历时，channel已经关闭，会正常遍历数据，遍历结束退出
  
  3. 遍历管道不能用普通的for循环

```go
package main

import "fmt"

func main() {
    intChan := make(chan int, 10)

    intChan <- 100
    intChan <- 200
    close(intChan)

    fmt.Println("OK~")

    // 管道关闭后，读取仍然可用
    num := <-intChan
    fmt.Println(num)

    // 遍历管道
    intChan2 := make(chan int, 100)
    for i := 0; i < 100; i++ {
        intChan2 <- i * i
    }
    // 必须关闭管道，不然出现死锁
    close(intChan2)

    for val := range intChan2 {
        fmt.Println("v=", val)
    }
}
```

### 2.7 groutine和channel结合

1. 应用案例

（1）开启一个WriteData协程，向管道intChan中写入50个整数

（2）开启一个ReadData协程，从管道intChan读取WriteData写入的数据

（3）注意：WriteData和ReadData操作的是同一个管道

（4）主线程需要等待WriteData和ReadData协程都完成才能退出【管道】

```go
package main

import "fmt"

func WriteData(intChan chan int) {
    for i := 0; i < 50; i++ {
        intChan <- i
        fmt.Printf("WriteData()写入的数据=%v\n", i)
    }
    close(intChan)
}

func ReadData(intChan chan int, exitChan chan bool) {

    for {
        v, ok := <-intChan
        if !ok {
            break
        }
        fmt.Printf("ReadData()读到的数据=%v\n", v)
    }
    // ReadData读取完数据后，任务完成
    exitChan <- true
    close(exitChan)
}

func main() {
    intChan := make(chan int, 100)
    exitChan := make(chan bool, 1)
    go WriteData(intChan)
    go ReadData(intChan, exitChan)

    for {
        _, ok := <-exitCn
        if !ok {
            break
        }
    }
}
```

练习

（1）启动一个协程，将1-2000的数放入到一个channel中，比如numChan

（2）启动8个协程，从numChan取出数，并计算1~n之和，存放到resChan

（3）最后8个线程都完成后，再遍历resChan，显示结果如res[1]=1，res[2]=3...

（4）考虑resChan chan int是否合适

2. 实例2-阻塞

将1代码中go ReadData(intChan, exitChan)注销

只向管道写入数据而没有读取，就会出现阻塞而dead lock，原因是intChan容量有限，会阻塞在WriteData的ch <- i

3. 实例3

需求：统计1-200000的数字哪些是素数。

要求：使用并发/并行的方式将统计素数的任务分配给多个goroutine完成。

### 2.8 channel使用细节和注意事项

1. channel可以声明为只读或只写形式

```go
    // 2.声明为只写
    var chan2 chan<- int
    chan2 = make(chan int, 3)
    chan2 <- 20

    // 3.声明为只读
    var chan3 <-chan int
    chan3 = make(chan int, 3)
    num2 := <-chan3
```

2. channel只读和只写的最佳实践

```go
func send(ch chan<- int, exitChan chan struct{}) {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)
    var a struct{}
    exitChan<- a
}


func recv(ch chan<- int, exitChan chan struct{}) {
    for {
        v, ok := <-ch
        if !ok {
            break
        }
        fmt.Println(v)
    }
    var a struct{}
    exitChan<- a
}

func main() {
    var ch chan int
    ch = make(chan, int, 10)
    exitChan := make(chan struct{}, 2)

    go send(ch, exitChan)
    go recv(ch, exitChan)

    var total = 0
    for _ := range exitChan {
        total++
        if total == 2 {
            break
        }
    }
    fmt.Println("结束...")
}
```

3. 使用select可以解决从管道读取数据的阻塞问题

```go
package main

import (
    "fmt"
)

func main() {
    // 使用select解决从管道取数据的阻塞问题

    // 1.定义1个管道，10个数据 int
    intChan := make(chan int, 10)
    for i := 0; i < 10; i++ {
        intChan <- i
    }
    // 2.定义1个管道, 5个string
    strChan := make(chan string, 5)
    for i := 0; i < 5; i++ {
        strChan <- "hello" + fmt.Sprintf("%d", i)
    }

    // 传统方法遍历管道, 如果不关闭管道会阻塞导致deadlock
    // 在实际开发中, 可能不好确定何时关闭管道
    // 可以使用select方式
label:
    for {
        select {
        case v := <-intChan: // 如果intChan一直没有关闭，不会一直阻塞而deadlock
            // 会自动到下一个case匹配
            fmt.Printf("从intChan读取数据%d\n", v)
        case v := <-strChan:
            fmt.Printf("从strChan读取数据%s\n", v)
        default:
            fmt.Println("读取结束....")
            break label
        }
    }
}
```

4. goroutine中使用recover，解决协程出现panic，导致程序崩溃问题。
   
   说明：如果起了一个协程，协程出现了panic，如果没有捕获到panic，整个程序会崩溃，此时可以在goroutine中使用recover来捕获panic进行处理。这样即使协程出现问题，主线程仍然不受影响，可以继续执行。

```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    for i := 0; i < 10; i++ {
        time.Sleep(time.Second)
        fmt.Println("hello world")
    }
}

func test() {
    // 使用defer+recover捕获panic
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("test() 发生错误", err)
        }
    }()
    var myMap map[int]string

    myMap[0] = "golang"
}

func main() {
    go sayHello()
    go test()

    for i := 0; i < 10; i++ {
        time.Sleep(time.Second)
        fmt.Println("main() runing....")
    }
}
```

# 第十七章 反射

使用反射机制，编写函数适配器(桥连接)

## 1. 基本介绍

- 反射可以在运行时动态获取变量各种信息，比如变量类型、类别

- 如果是结构体变量，还可以获取到结构体本身的信息(字段、方法)

- 通过反射，可以修改变量值，可以调用关联的方法

- 使用反射，import ("reflect")

## 2. 反射重要的函数和概念

（1）refelct.TypeOf(变量名)，获取变量类型，返回reflect.Type类型

（2）reflect.ValueOf(变量名)，获取变量的值，返回reflect.Value类型是一个结构体类型。通过reflect.Value，可以获得该变量的很多信息。

（3）变量、interface{}和reflect.Value是可以相互转换的，在实际开发中经常使用到。

```go
// 专门做反射
func test(b interface{}) {
    // 1.将interface{}转成reflect.Value
    rVal := reflect.ValueOf(b)

    // 2.将reflect.Value转成interface{}
    iVal := rVal.interface()

    // 3.如何将interface{}转成原来的变量类型
    // 使用类型断言
    v := iVal.(Stu)
}
```

![](C:\Users\12143\AppData\Roaming\marktext\images\2023-03-24-22-38-59-image.png)

## 3. 反射应用场景

1. 不知道接口调用哪个函数，根据传入参数在运行时确定调用的具体接口，这周需要对函数或方法反射。例如下面桥接模式
   
   func bridge(funcPtr interface{}, args ...interface{})
   
   第一个参数funcPtr以接口形式传入函数指针，函数参数args以可变参数形式传入，bridge函数中可以使用反射来动态执行funcPtr函数

2. 结构体序列化时，结构体指定Tag，也会使用反射生成对应字符串。

## 4. 快速入门

4.1 编写案例，演示对(基本数据类型、interface{}、reflect.Value)进行基本的反射操作

```go
func test01(b interface{}) {
    t := reflect.TypeOf(b)
    fmt.Println("Type is ", t)

    v := reflect.ValueOf(b)
    k := v.Kind()
    fmt.Println("kind=", k)
    fmt.Println("value=", v.Int())
}
```

4.2 编写案例，演示对(结构体类型、interface{}、reflect.Value)进行基本的反射操作

```go
package main

import (
    "fmt"
    "reflect"
)

type Stu struct {
    Name string
    Age  int
}

func ReflectStruct(b interface{}) {
    t := reflect.TypeOf(b)
    fmt.Println("Type is ", t)
    fmt.Println("Kind is ", t.Kind())

    v := reflect.ValueOf(b)
    k := v.Kind()
    fmt.Println("kind=", k)
    fmt.Printf("v is %v, v's type is %T\n", v, v)

    // 将v转换为interface{}
    iv := v.Interface()

    // 将interface{}断言转换为需要的类型
    student, ok := iv.(Stu)
    if !ok {
        fmt.Println("error... ")
    }
    fmt.Println(student)
}

func main() {
    // 定义结构体实例
    stu := Stu{
        Name: "Tom",
        Age:  20,
    }

    ReflectStruct(stu)

}
```

## 5.注意事项和细节说明

（1）reflect.Value.Kind，获取变量类别，返回的是一个常量

（2）Type是类型，Kind是类别，Type和Kind可能相同，也可能不同

            比如：

            var num int = 10 num的Kind和Type都是int

            var stu Student  stu的Type是包名.Student，Kind是struct

（3）通过反射可以让变量在interface{}和Reflect.Value之间相互转换

（4）使用反射方式获取变量的值(<font color="red">并返回对应的类型</font>)，要求数据类型匹配，比如x是int，就应该使用reflect.Value(x).Int()，而不是使用其他，否则报panic

（5）通过反射修改变量，注意使用SetXxx方法设置，需要通过<font color="red">对应的指针类型</font>来完成，这样才能改变传入的变量值，同时需要使用reflect.Value.Elem()方法

![](C:\Users\12143\AppData\Roaming\marktext\images\2023-03-25-13-41-27-image.png)

```go
func reflect01(b interface{}) {
    rVal := reflect.ValueOf(b)

    // 改变rVal的值
    rVal.Elem().SetInt(20)
}
```

（6）reflect.Value.Elem()讲解

Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。

```go
num := 9
ptr *int = &num
num2 := ptr    // 类似rVal.Elem()
```

## 6. 最佳实践

- 使用反射遍历结构体字段，调用结构体方法，并获取结构体标签的值

```go
package main

import (
    "fmt"
    "reflect"
)

type Monster struct {
    Name  string `json:"name"`
    Age   int    `json:"monster_age"`
    Score float32
    Sex   string
}

func (s Monster) Print() {
    fmt.Println("---start---")
    fmt.Println(s)
    fmt.Println("---end---")
}

func (s Monster) Getsum(n1, n2 int) int {
    return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
    s.Name = name
    s.Age = age
    s.Score = score
    s.Sex = sex
}

func TestStruct(a interface{}) {
    typ := reflect.TypeOf(a)
    val := reflect.ValueOf(a)
    kd := val.Kind()

    if kd != reflect.Struct {
        fmt.Println("Expect struct")
        return
    }

    num := val.NumField()
    for i := 0; i < num; i++ {
        fmt.Printf("Field %d: 值为%v\n", i, val.Field(i))

        // 获取struct标签，通过reflect.Type来获取tag标签的值
        tagVal := typ.Field(i).Tag.Get("json")
        if tagVal != "" {
            fmt.Printf("Field %d: tag为%v\n", i, tagVal)
        }
    }

    numOfMethod := val.NumMethod()
    fmt.Printf("struct has %d methods\n", numOfMethod)
    // var params []reflect.Value
    val.Method(1).Call(nil)

    // 调用结构体的第1个方法, 反射使用的方法0和方法1排序使用的是对比结构体方法首字母是默认按照ASCII码比较
    var params []reflect.Value
    params = append(params, reflect.ValueOf(10))
    params = append(params, reflect.ValueOf(40))
    res := val.Method(0).Call(params) // 传入参数是 []reflect.Value
    fmt.Println("res=", res[0].Int())
}

func main() {
    var a Monster = Monster{
        Name:  "黄精怪",
        Age:   800,
        Score: 90.35,
    }
    TestStruct(a)
}
```

### 还有4个例子在P291

# 第十八章 TCP编程

## 1. 基本介绍

网络编程有2种：

(1) TCP socket编程，网络编程主流。底层基于TCP/IP协议

(2) b/s结构的http编程，使用浏览器访问服务器，使用的就是http协议，底层仍使用TCP/IP协议实现。

- ip地址
  
  每个internet上的主机和服务器都有一个ip地址，包括网络号和主机号，ip地址分ipv4和ipv6

- 端口
  
  0是保留端口
  
  1-1024是固定端口
  
  1025-65535是动态端口
  
  **注意事项**：
  
  1. 计算机尽可能少开端口
  
  2. 一个端口只能被1个程序监听
  
  3. 如果使用netstat -an可以查看本机哪些端口在监听
  
  4. 可以使用netstat -anb查看监听端口的pid，再结合任务管理器关闭不安全的端口

## 2.快速入门

### **服务端处理流程**

(1) 监听端口

(2) 接收客户端的tcp链接，建立客户端和服务端的链接

(3)创建goroutine，处理该链接的请求(通常客户端会发送)

### **客户端处理流程**

建立与服务端的链接

发送请求数据，接收服务端返回的结果数据

关闭链接

client.go

```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.11:8888")
	if err != nil {
		fmt.Println("Dial() err= ", err)
		return
	}

	for {
		// 功能1：客户端可以发送单行数据，然后就退出
		reader := bufio.NewReader(os.Stdin) // os.Stdin, 代表标准输入[终端]

		// 从终端读取一行用户输入，并准备发送给服务器
		line, err1 := reader.ReadString('\n')
		if err1 != nil {
			fmt.Println("ReadString() err= ", err1)
		}

		// 4. 如果输入exit就退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端服务退出...")
			return
		}
		// 将line发给服务器
		n, err2 := conn.Write([]byte(line + "\n"))
		if err2 != nil {
			fmt.Println("conn.Write() err= ", err2)
		}
		fmt.Printf("客户端发送了 %d 个字节数据，并退出了\n", n)

	}
}
```

server.go

```go
package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	// 循环接收客户端发送的数据
	defer conn.Close() // 接受完即关闭

	for {
		// 创建1个新的切片
		buf := make([]byte, 1024)

		// conn.Read(buf)
		// 1. 等待客户端通过conn发送消息
		// 2. 如果客户端没有write[发送], 协程就会阻塞在这个地方
		fmt.Printf("服务器在等待客户端%s输入\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("远程客户端已关闭 ")
			return
		}
		if err != nil {
			fmt.Println("conn.Read() err1= ", err)
			return
		}
		// 3. 显示客户端发送的内容到服务器终端
		fmt.Print(string(buf[:n])) // 这个地方一定要写:n, 非常重要！
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	// 1. tcp表示使用网络协议是tcp
	// 2. 0.0.0.0:8888表示本地监听8888端口
	if err != nil {
		fmt.Println("listen err= ", err)
		return
	}
	// fmt.Printf("listen suc=%T\n", listen) // listen suc=*net.TCPListener

	defer listen.Close() //演示关闭listen

	// 循环等待客户端连接
	for {
		fmt.Println("等待客户端链接...")
		conn, err1 := listen.Accept()
		if err1 != nil {
			fmt.Println("Accept() err1= ", err1)
		} else {
			fmt.Printf("Accept() suc con=%v, 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}

		// 准备起一个协程
		go process(conn)
	}

}
```

## 海量用户即时通讯系统

# 第19章 Redis的使用

## 1. 基本介绍

1. Redis是NoSQL数据库，不是传统的关系型数据库

2. Redis：REmote Dictionary Server(远程字典服务器)，Redis性能非常高，单机能够达到15w qps，通常适合做缓存，也可以持久化。

3. <font color="red">高性能分布式内存数据库</font>，基于内存运行并<font color="blue">支持持久化</font>的NoSql数据库。也称为数据结构服务器。

![](C:\Users\12143\AppData\Roaming\marktext\images\2023-03-27-18-27-35-image.png)

## 2. 基本使用

Redis安装好后，默认16个数据库，初始默认使用0号库，编号0...15

1. 添加key-val [set]

2. 查看当前Redis的所有key [keys *]

3. 获取key对应的值[get key]

4. 如何查看当前数据库的key-val数量[dbsize]

5. 清空当前数据库的key-val和清空所有数据库的key-val[flushdb flushall]

## 3.Redis数据类型和CRUD

- Redis五大数据类型
  
  String(字符串)、Hash(哈希)、List(列表)、Set(集合)、zset(sorted set：有序集合)

### 3.1 string

string是redis最基本的类型，一个key对应一个value

string类型是二进制安全的，除普通字符串外，也可以存放图片等数据。

- CRUD
  
  set(如果key存在相当于修改)/get/del

- 使用细节和注意事项
  
  - **setex(set with expire)键秒值**
  
  SETEX key seconds value
  
  将值value关联到key，并将key的生存时间设为secons(以秒为单位)
  
  如果key已经存在，SETEX命令将覆写旧值
  
  类似于以下2个命令
  
  SET key value
  
  EXPIRE key seconds # 设置生存时间
  
  - **mset[同时设置1个或多个key-value对]**
    
    如果某个给定的key已存在，MSET会用新值覆盖旧值
  
  - **mget[同时获取1个或多个key-value]**

### 3.2 Hash（哈希，类似golang里的map)

redis hash是一个键值对集合。var user1 map[string]string

redis hash是一个string类型的field和value的映射表，hash特别适用存储对象

如：存放一个user信息

user1 name zhangsan age 30

key:user1

name zhangsan 和 age 30两对field-value

- CRUD
  
  hset/hget/hgetall/hdel

- 使用细节和注意事项
  
  给user设置name和age时，前面是一步步设置，使用hmset和hmget可以一次性设置多个filed的值和返回多个field的值。

- hlen统计一个hash有多少个元素

- hexists key field
  
  查看哈希表key中，给定域field是否存在

### 3.3 List

列表是简单的字符串列表，按照插入顺序排序。可以添加一个元素到列表头部(左边)或者尾部(右边)

List本质是个链表，List的元素是有序的，元素的值可以重复。

举例：

lpush city beijing shanghai suzhou

lrange city 0 -1 

- CRUD
  
  lpush/rpush/lrange/rpop/del
  
  - LRANGE key start stop
    
    返回列表key中指定区间内的元素，区间以偏移量start和stop指定。下标(index)参数start和stop都以0为底，即以0表示列表第一个元素。
    
    也可以使用负数下标，-1表示列表最后一个元素，-2表示倒数第2个元素

- 可以把I想像为一根管道

**基本使用**

```
127.0.0.1:6379> lpush herolist aaa bbb ccc
(integer) 3
127.0.0.1:6379> lrange herolist 0 -1
1) "ccc"
2) "bbb"
3) "aaa"
127.0.0.1:6379> rpush herolist ddd eee
(integer) 5
127.0.0.1:6379> lrange herolist 0 -1
1) "ccc"
2) "bbb"
3) "aaa"
4) "ddd"
5) "eee"
127.0.0.1:6379> lpop herolist
"ccc"
127.0.0.1:6379> lrange herolist 0 -1
1) "bbb"
2) "aaa"
3) "ddd"
4) "eee"
127.0.0.1:6379> rpop herolist 0 -1
(error) ERR wrong number of arguments for 'rpop' command
127.0.0.1:6379> rpop herolist
"eee"
127.0.0.1:6379> del herolist
(integer) 1
127.0.0.1:6379> lrange herolist 0 -1
(empty list or set)
```

**使用细节和注意事项**

1. lindex，按照索引下标获取元素(从左到右，编号从0开始)

2. LLEN key
   
   返回列表key的长度，如果key不存在，则key被解释为1个空列表，返回0

3. List其他说明
   
   List数据，可以从左或从右插入添加
   
   如果值全移除，键也就消失了

### 3.4 集合

Redis的Set是string类型的无序集合

底层是HashTable的数据结构，Set也是存放很多字符串元素，字符串元素是无序的。而且元素值不能重复.

举例：

sadd email sgg@sohu.com tome@sohu.com

- CRUD
  
  sadd
  
  smemebers[取出所有值]
  
  sismember[判断值是否是成员]
  
  srem[删除指定值]
