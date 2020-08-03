### 复合数据结构 ###

复合数据结构:就是组合基本类型而构造出来的复合数据类型.

数组:由同构的元素组成,固定内存大小;    
结构体:由异构的元素组成,固定内存大小;    
slice,map是动态数据结构,根据需要动态增长;    


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











