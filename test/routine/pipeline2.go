package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	// naturals := make(chan int)
	// squares := make(chan int)
	// go counter(naturals)
	// go squarer(squares, naturals)
	// printer(squares)
	// queue()
	fmt.Println(mirroredQuery())
}

func queue() {
	ch := make(chan string, 3)
	ch <- "a"
	ch <- "b"
	ch <- "c"
	fmt.Println(<-ch)
	ch <- "d"
	fmt.Println(<-ch)
	fmt.Println(len(ch))
	fmt.Println(<-ch)
	fmt.Println(len(ch))
	fmt.Println(<-ch)
	fmt.Println(len(ch))
	fmt.Println(<-ch)
	fmt.Println(len(ch))
	ch <- "c"
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()
	return <-responses // return the quickest response
}

func request(hostname string) (response string) {
	fmt.Println(hostname)
	return hostname
}
