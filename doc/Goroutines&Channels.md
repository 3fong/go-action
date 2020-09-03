### Goroutines ###

goroutine:是go的并发执行单元.简单上可以理解为其他语言的线程.当程序启动时,主函数运行在一个单独的goroutine中,我们叫它main goroutine;新的goroutine会用go语句来创建;语法:go关键字+函数或方法;go语句会使用其语句中的函数在新创建的goroutine中运行.

	f()    // call f(); wait for it to return
	go f() // create a new goroutine that calls f(); don't wait

主函数返回时,所有的goroutine都会被直接打断,程序退出.

	func main() {
	    go spinner(100 * time.Millisecond) // 在新创建的goroutine中执行spinner函数
	}
	
	func spinner(delay time.Duration) {
		fmt.Printf(delay)
	}

goroutine退出方式:

1. 主函数退出
2. 直接终止程序
3. goroutine请求触发其他goroutine自动结束执行

goroutine使用需要考虑调用时是否安全.goroutine是协程,它比线程更小,十几个goroutine可能底层只有五六个线程,内存消耗也更小,单个goroutine大概占用4-5kb栈内存.它比线程更易用,更高效,更轻便;

runtime包

Gosched 礼让协程
Goexit 协程结束执行
GOMAXPROCS 设置并行计算cpu核数最大值

### Channels ###

Channels:协程通讯通道.它是一个通讯机制,用于goroutine间发送消息.它可以发送数据的类型

创建channel
	
	ch := make(chan int) // 通道类型是int类型

channel是个值引用类型,零值为nil;也因此可以进行==比较;    
channal的通信行为包括:发送和接收;发送和接收的操作都使用<-运算符;发送时<-位于channel和值之间;接收<-位于channel前;

	ch <- x  // 发送
	x = <-ch // 接收
	<-ch     // 接收,忽略结果

通道关闭后,发送会产生panic;接收可以正常接收通道里的值,通道为空时返回零值数据;

	close(ch) // 用于关闭通道

创建通道时,容量大于零则表示为带缓存channel

	ch = make(chan int)    // unbuffered channel
	ch = make(chan int, 0) // unbuffered channel
	ch = make(chan int, 3) // buffered channel with capacity 3

#### 不带缓存的channel ####

无缓存channel也称为同步channel:它会导致两个goroutine进行同步操作,goroutine发送数据后会阻塞发送至消息被接收;反之一样,接收先发生,接收者goroutine将阻塞等待消息发送;    
happens before:接收者收到数据发生在再次唤醒发送者goroutine之前;x事件在y事件之前发生,并不是强调x发生时间比y早,而是要强调保证y前x会完成;    
并发:x事件不发生在y事件之前或之后;这里只是无法区分x,y发生的事件先后,而不是一定同时发生;    
消息事件:强调通讯发生的时刻(事实),而不是消息具体值的情形;很多时候消息事件不没有具体值;










































