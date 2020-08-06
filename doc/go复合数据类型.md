### 复合数据结构 ###

复合数据结构:就是组合基本类型而构造出来的复合数据类型.

数组:由同构的元素组成,固定内存大小;    
结构体:由异构的元素组成,固定内存大小;    
slice,map是动态数据结构,根据需要动态增长;    

#### 数组 ####

数组是一个由**固定长度**的特定类型元素组成的序列,一个数组可以由零个或多个元素组成.Slice是基于数组的可变长度数据结构,实际中主要用Slice进行数据处理,数组较少直接使用.

数组的每个元素可通过索引下标访问.范围:0-(数组长度-1);

默认情况下,数组的每个元素都被初始化为元素类型对应的零值.


	var q [3]int = [3]int{1, 2, 3}    
	// 通过实际数据来推定数组长度,定义了含有100个元素的数组r,通过脚标来指定默认值为-1.其他元素初始化值为0    
	r := [...]int{99: -1}    


数组的长度是数组类型的一部分,因此[3]int和[4]int是两种不同的数组类型.数组的长度必须是常量表达式,因为数组的长度需要在编译期确定.

如果一个数组的元素类型相同,那么数组类型也是可以项目比较的,可以通过==判断数组是否相等.

实际比较中可以通过sha的消息摘要来判断值是否有可能相同


     	c1 := sha256.Sum256([]byte("x"))
		c2 := sha256.Sum256([]byte("X"))
   	    fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1) 
    	%x:指定以十六进制打印数组或slice全部元素
    	%t:用于打印布尔值
    	%T:显示值对应的数据类型

局部变量和引用变量:    
当函数调用时,函数的调用参数会被赋值给函数内部的参数变量,函数参数变量接收的是一个复制的副本,并不是原始调用的变量,这会造成修改内部变量不会改变引用变量值.因此很多语言区分局部变量和引用变量,引用变量就是语言会隐式的将参数引用或指针对象传入被调用的函数,达到函数内部修改外部变量的效果.go的处理比较直接,需要显示传入参数指针,函数才会进行外部变量修改.

    func zero(ptr *[32]byte) {
       for i := range ptr {
          ptr[i] = 0
       }
    }

	// 函数指针调用
	a := [32]byte
	zero(&a)

*数据类型代表指针,&变量名代表传入变量的指针

#### Slice ####

Slice(切片):变长的序列.它的底层实现是数组         
slice类型一般写作[]T.T代表元素类型;slice写法和数组很像,只是没有固定长度而已.    

slice有三部分组成:指针,长度和容量.    
指针指向第一个slice元素对应的底层数组元素的地址,注意slice的第一个元素不一定是数组的第一个元素.
长度不能超过容量,容量一般是从slice的开始位置到底层数据的结尾位置.内置的len,cap函数分别返回slice的长度和容量    

多个slice之间可以共享底层的数据,并引用数组部分区间可能重叠.


	months := [...]{1:"January", /* ... */, 12:"December"}

x[m:n]代表切片的切片,它是切片的子序列.    
因为slice值包含指向第一个slice元素的指针,因此向函数传递slice将允许在函数内部修改底层数组的元素.    
slice的初始化和数组类似,只是不用指明序列的长度.

	arr := [...]{1,2,3} // 数组初始化   
	sli := []{1,2,3}  // slice初始化

slice不能用==操作符判断slice是否相等.原因是:1slice的元素是间接引用,处理复杂;2数据结构更安全     
判断slice是否为空,使用len(s)==0,nil包含已经初始化但是未赋值的情况   

	var s []int // len(s)==0,s == nil

创建匿名的数组变量,返回slice:    

	make([]T,len,[cap])


append函数

append函数用于向slice追加元素.

	var runes []rune
	for _, r := range "Hello, 世界" {
	    runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"

slice的结构:

	type IntSlice struct {
	    ptr      *int
	    len, cap int
	}

组成结构参数:结构体指针,长度和容量,虽然底层数组的元素是间接访问的,但是slice的组成结构参数是直接访问的.

#### Map ####

hash表是一种巧妙且实用的数据结构.它是无序的key/value对的集合,其中key是不同的,通过key可以在常数时间复杂度内检索,更新或删除对应的value.

格式:

	map[K]V
	ages := make(map[string]int)
	ages["alice"]=31
	ages["bob"]=60
	===
	ages := map[string]int{
		"alice":31
		"bob":60
	}

map标是类型;K为key;V为value;
map中所有key的类型相同;所有value的类型相同;但是key和value可以是不同类型;key必须支持==比较;所以slice类型为key时必须进行类型转换.

无法对map元素进行取址操作,因为map可能随着元素的数量增长而重新分配内存空间,造成原先的地址无效

map迭代顺序不确定,遍历的顺序是随机的,每次遍历的顺序都不相同;

	age, ok := ages["tom"]
	if !ok { /* "tom" is not a key in this map; age == 0. */ }

如果key不存在,map将返回零值,所以map取值可以通过第二次布尔参数判断是否值存在,不存在时返回false.

go中没有set类型,可以使用map来实现set的功能.map[string]bool,key就是实际set要存在的值

#### 结构体 ####

结构体是一种聚合的数据类型,是由零个或多个任意类型的值聚合成的实体.每个值称为结构体的成员.

	// 定义
	type Employee struct {
	    ID        int
	    Name      string
	    Address   string
	    DoB       time.Time
	    Position  string
	    Salary    int
	    ManagerID int
	}
	// 使用
	var dilbert Employee
	dilbert.Salary -= 5000


EmployeeByID(id).Salary = 0

这种写法,EmployeeByID(id)必须要返回实体的指针,因为调用函数返回的是值,而不是一个变量.

如果EmployeeByID(id)不返回指针,则需要写成如下格式:

	var aa = EmployeeByID(id)
	aa.Salary = 0

结构体中不能包含自己,但是可以包含自己的指针;

	type tree struct {
		value       int
		left, right *tree
	}

空结构体:没有任何成员的结构体,写作struct{},大小是0.可以用于结构占位,如map模拟set数据就可以写成:

	seen := make(map[string]struct{})

但是struct的空间占用比较大,一般不这么写


结构体字面值

	type Point struct{ X, Y int }

方法一:

	p := Point{1, 2}

通过顺序来直接赋值,这种方式很容易出错,一般用于比较简单的结构

方法二:

	p := Point{x:1, y:2}

通过参数名来赋值,更直观,突破顺序的束缚

结构体可作为函数的参数和返回值;    
如果要在函数内部修改结构体成员的话,必须传入结构体指针.因为go的所有函数参数都是值拷贝传入,函数参数将不再是函数调用时的原始变量;      
结构体的全部成员都可以比较,那么结构体也可以比较.    

匿名成员:声明一个成员对应的数据类型而不指明成员的名称.匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针.匿名成员并不要求是结构体类型,任何类型都可以是匿名成员.匿名成员不仅简化了调用格式,同时也可以使用匿名成员类型的方法,这是它的核心价值


	type Circle struct {
  	   Point
 	   Radius int
	}

	// 赋值.因为Point未命名,所以可以直接赋值,
	var c Circle
	c.X=1
	c.Y=1
	c.Radius=5

但是它的结构体字面值还需要使用定义结构:

	c = Circle{Point{1,2},5}

因为匿名成员有一个隐式的名字,因此不能同时包含两个类型相同的匿名成员,这会导致名字冲突.

#### JSON ####

JSON(JavaScript Object Notation)JavaScript对象表示法:是一种用于发送和接收结构化信息的标准协议.特点:简洁,易读,流行

go的编码和解码标准库:encoding/json,encoding/xml,encoding/asn1等

基本JSON类型:数字(十进制或科学计数法),布尔值,字符串

数据类型转换时,JSON只对导出的结构体进行编码,即大写字母开头的才会进行编码.

go的JSON tag使用:

	type Movie struct {
 	   Title  string
  	   Color  bool `json:"color,omitempty"`
 	   Actors []string
	}

`json:"color,omitempty"`即是JSON的tag,它用于知名json格式时,如果为true则字段名为color,否则不显示该字段.

json和slice转换类:

data, err := json.Marshal(movies) // 编码:将slice转为json
err := json.Unmarshal(data, &titles) // 解码:将json转为slice

#### 文本和HTML模板 ####

text/template和html/template等模板包提供了将数据和模板分开管理,增强了数据展示的灵活性和安全性.

模板可以是一个字符串或文件,里面包含了一个或多个双花括号的对象{{action}}.action部分用于动态展示数据,每个action都包含了一个用模板语言书写的表达式,模板语言包含通过选择结构体的成员,调用函数或方法,表达式控制流if-else语句和range循环语句,还有其他实例化模板等诸多特性.

模板字符串:

	const templ = `{{.TotalCount}} issues:
	{{range .Items}}----------------------------------------
	Number: {{.Number}}
	User:   {{.User.Login}}
	Title:  {{.Title | printf "%.64s"}}
	Age:    {{.CreatedAt | daysAgo}} days
	{{end}}`

每个action都有一个当前值的概念,用"."表示.{{range .Items}}和{{end}}对应一个循环action;"|"操作符类似于UNIX的管道,它将前一个表达式的结果作为后一个函数的输入.daysAgo是一个函数.

template.Must辅助函数可以简化这个致命错误的处理：它接受一个模板和一个error类型的参数，检测error是否为nil（如果不是nil则发出panic异常），然后返回传入的模板
