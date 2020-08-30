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

goroutine使用需要考虑调用时是否安全.

### Channels ###
















































