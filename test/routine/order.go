package main

import (
	"fmt"
	"runtime"
)


func main() {
	var v string = "5"
	go a(v)
	go b(v)
	runtime.Gosched()
	fmt.Println("c"+v)
}

func a(v string) {
	v += "a"
	fmt.Println(v)
}

func b(v string) {
	v += "b"
	fmt.Println(v)
}