package main

import (
	"fmt"
	"runtime"

	"github.com/liulei3/go-action/test/myreflect"
	"github.com/liulei3/go-action/test/routine/change"
)

func main() {
	myreflect.PrintType()
	myreflect.Print(myreflect.GetObj())
}

func PrintData() {
	var v string = "5"
	go a(v)
	go b(v)
	runtime.Gosched()
	fmt.Println("c" + v)
	change.Multipy()
}

func a(v string) {
	v += "a"
	fmt.Println(v)
}

func b(v string) {
	v += "b"
	fmt.Println(v)
}
