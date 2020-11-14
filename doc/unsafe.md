### 底层编程 ###

高级语言为了便于使用以及安全性会进行方法封装,这会隐藏很多底层细节.但很多时候为了性能或更直接的使用会摆脱这种语言规则的限制.这时就用到了unsafe包的内容.    
unsafe包采用特殊方式实现,由编译器提供一些访问语言内部特性的方法,特别是内存布局相关的细节.它被封装在独立的包中,只适合极少场景使用,同时也会存在一些安全性问题.    

#### unsafe包 ####

unsafe包功能比较简单,包括两个type和三个函数:
	
	func Alignof(x ArbitraryType) uintptr // 对应参数的类型需要对齐的倍数.
	func Offsetof(x ArbitraryType) uintptr //地址偏移量
	func Sizeof(x ArbitraryType) uintptr //操作数在内存中的字节大小
	type ArbitraryType // 二进制go表达式,int类型
	type Pointer		// ArbitraryType的指针


Sizeof:返回操作数在内存中的字节大小,参数可以是任意类型的表达式.返回的大小只包括数据结构中固定部分,如字符串对应结构体中的指针和字符串长度部分,但是不包含字符串内容.也因此非聚合类型(非struct)通常是固定的大小,但是不同环境下会有差异.

类型	大小
bool	1个字节
intN, uintN, floatN, complexN	N/8个字节（例如float64是8个字节）
int, uint, uintptr	1个机器字
*T	1个机器字
string	2个机器字（data、len）
[]T	3个机器字（data、len、cap）
map	1个机器字
func	1个机器字
chan	1个机器字
interface	2个机器字（type、value）

Alignof:返回对应参数的类型需要对齐的倍数.它也是常数.通常布尔和数字类型需要对齐到他们本身大小,其他类型对齐到机器字大小.

Offsetof:是字段相对于结构体的偏移量.即字段x.f返回f字段相对于x起始地址的偏移量.

Pointer:它代指任意类型的指针.它和普通指针一样可以比较,但是不能直接接收其值,因为不知道变量的具体类型.     
它的值接收可以通过两种方式:1是将它转换为普通指针;2将它转换成uintptr类型.

pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))



















