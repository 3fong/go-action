package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/liulei3/go-action/test/myreflect/format"
	"github.com/liulei3/go-action/test/routine/change"
)

func main() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(format.Any(x))          // "1"
	fmt.Println(format.Any(d))          // "1"
	fmt.Println(format.Any([]int64{x})) // "[]int64 0x8202b87b0"
	fmt.Println(format.Any([]time.Duration{d}))
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
