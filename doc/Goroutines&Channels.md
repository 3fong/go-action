### Goroutines ###

#### 进程,线程和协程区别:   ####  

进程和线程:内核进行调度,有cpu时间片的概念,进行**抢占式调度**;相互间更公平,但是资源占用高;    
协程:用户级线程,对内核透明,程序自行调度,通常只能进行**协作式调度**,需要协程主动让出控制权;相互间执行不公平,无法直接利用多核优势,但资源占用低;

goroutine(协程):是go的并发执行单元.简单上可以理解为其他语言的线程.协程相对独立,有自己的上下文,切换由自己控制(线程是由系统控制);当程序启动时,主函数运行在一个单独的goroutine中,我们叫它main goroutine;新的goroutine会用go语句来创建;    
go中所有系统调用操作都会出让cpu给其他goroutine,这使得goroutine切换不依赖系统线程和进程,也不依赖cpu的核数;    

#### 协程原理 ####

多线程编程:线程多则上下文切换频繁,cpu时间消耗大;    
异步编程:针对上面一个线程一个socket连接的消耗问题,通过少了线程来服务大量的网络连接和I/O操作;这样会让代码复杂,易出错(执行无序);    
协程:应用模拟线程,避免上下文切换,降低并发复杂度;兼顾了并发和性能;    
原理:和线程一样,维护一个线程栈,当a线程切换到b线程是,需要将a线程的相关执行进度压入栈,然后将b线程的执行进度出栈,进入b线程的执行序列;只不过协程是在应用层实现这一点;    

问题:
1 应用程序没有cpu调用权限,无法直接操作线程入栈,出栈执行问题?    
解决:协程是基于线程的,内部维护一组数据结构和n个线程,真正执行还是线程,协程执行代码被扔进一个待执行队列,由n个线程从队列拉取执行;    
2 协程如何切换,即异步执行问题?    
解决:利用操作系统的异步函数实现,包括linux的epoll,select,windows的iocp,event等.go通过封装各种io函数,这些io函数调用了操作系统底层的异步io函数,当异步函数返回busy或blocking时,go就进行现有执行序列压栈,让线程拉取另一个协程代码执行.    

语法:go关键字+函数或方法;go语句会使用其语句中的函数在新创建的goroutine中运行.

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

#### pipeline(串联的channels) ####

pipeline:是将多个channel连接在一起.由于是同步channel,channel越多越可能因为阻塞造成死循环;如果channel数据发送完毕可以通过close关闭,避免阻塞发生;但是关闭会造成发送panic,接收无限获取nil值;实际不需要关闭每一个channel,当channel没有被引用时,go会进行垃圾自动回收;

	go func() {
        for x := range naturals {
            squares <- x * x
        }
        close(squares)
    }()

实际只需要在业务执行完再进行关闭即可

#### 单方向的channel ####

单方向的channel:语法上限制channel只能用于发送或接收.通过这种限制方式,来避免使用混乱问题,也可以避免位置的panic;    

	发送: chan <- int 表示只发送int的channel    
	接收: <- chan int 表示只能接收int的channel    

因为关闭操作只用于断言不再向channel发送新数据,所以只有发送方的goroutine才会调用close函数;    
可以将channel转换为单方向channel,但是不能将单方向channel转成正常的channel,这种转换是单向的;    

#### 带缓存的channels ####

带缓存的Channel内部持有一个元素队列.

	ch = make(chan string, 3)

向缓存channel发送操作就是向内部缓存队列的尾部插入元素,接收操作时从队列的头部删除元素;如果队列是满的,则阻塞发送;如果队列是空的,则阻塞接收;    

查询channel容量

	cap(ch)

查询数量

	len(ch)

channel和goroutine的调度器机制紧密相连,如果没有其他goroutine从channel发送或接收,将会有永远阻塞的风险(deadlock)

注意事项:

1. 无缓存的channel,存在无接收方引起的goroutines泄漏问题;因为发送一直被占用;
2. 无缓存channel用于保证每个发送操作与相应的同步接收操作;带缓存channel用于解耦通信操作;
3. 带缓存channel的容量规划也很关键,因为容量大的话,会造成发送缓存(接收闲置,资源利用不佳);容量小时,又会造成程序死锁(发送堆积);
4. 带缓存channel可能影响程序性能;发送和接收的效率不一致,会造成程序闲置,性能不佳;


#### goroutine与channel ####

并发中,当goroutine执行时由于异步执行的原因,并不会等待执行结果,它会在触发异步后直接返回,可能会造成执行中断或无法获取结果的问题;这时就需要channel来进行goroutine间通信;


- select多路复用     

它和switch语句有点像,它用于选择通信操作(channel的发送和接收);当通道接收时,可以根据执行块决定是否声明接收变量;当满足条件时,select会通信并执行case后的语句;这时其他通信不会执行;一个没有任何case的select语句写作select{},标识永远等待下去;

	select {
	case <-ch1:
	    // ...
	case x := <-ch2:
	    // ...use x...
	case ch3 <- y:
	    // ...
	default:
	    // ...
	}


如果多个case同时就绪时,select会随机地选择一个执行,以保证channel执行公平;    
default中设置其他操作都能处理的逻辑;     
nil的channel表示永远阻塞;select语句中操作nil的channel永远都不会被执行到;可以通过nil来激活或禁用case,实现额外事件超时和取消的逻辑;

#### goroutine退出 ####

程序中我们需要当业务执行异常或某些原因中途退出goroutine执行,这就需要goroutine可以退出.但是go并没有提供一个goroutine中终止另一个goroutine的方法,因为这会造成goroutine间的共享变量状态不可控;     

退出方式有如下几种:

	1 向abort的channel中发送和goroutine同样数量的退出消息;这种做法很理想,因为消息发送接收都可能阻塞,同时goroutine的实际数量也无法准确统计;    
	2 广播机制;消费掉所有channel发送值并关闭channel;这样操作channel之后的代码会立即被执行;

基于信号管理:

	适合协程具备层级关系的情形;主协程下有子协程,主协程发送信息,触发子协程的关闭;

问题:

	接收者和发送者,缺少并行机制,消息订阅模式;

builtin包的close方法

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

创建一个goroutine,用于执行匿名函数,输入任意字符触发close方法,close可以接收一个发送者通道,用于关闭通道并进行关闭广播;它会向所有接收者发送消息,用于触发接收者后面的代码;    

当主协程返回时,一个程序会退出,这样你无法在主函数退出后确认是否所有资源都进行了释放;这里可以使用panic让runtime把每个goroutine的栈dump下来,如果主协程时唯一剩下的goroutine,证明资源释放成功,否则就想办法进行资源释放;

核心问题:调用链关系.     
资源请求和管理有层级关系;

全局机制:上下文共享数据结构,只读,规避风险;具备通过具体点进行上下的层管理.








