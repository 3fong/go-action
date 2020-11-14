package main

import (
	"fmt"
	"go-action/practice/myreflect"
	"go-action/practice/myreflect/parse"
	"go-action/practice/routine/change"
	"runtime"
	"time"
	"unsafe"
)

func main() {

	point()
}

func Myreflect() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(myreflect.Any(x))          // "1"
	fmt.Println(myreflect.Any(d))          // "1"
	fmt.Println(myreflect.Any([]int64{x})) // "[]int64 0x8202b87b0"
	fmt.Println(myreflect.Any([]time.Duration{d}))
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

func search() {
	// practice.Fetch()
	fmt.Printf("%d %s\n", "hello", 42)
	parse.Print(time.Hour)
}

func point() {
	var x struct {
		a bool
		b int16
		c []int
	}

	// equivalent to pb := &x.b
	pb := (*bool)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.a)))
	*pb = false

	fmt.Println(x) // "42"
	var str = "1234567899999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999"
	fmt.Println(unsafe.Sizeof(str))

	// NOTE: subtly incorrect!
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pp := (*int16)(unsafe.Pointer(tmp))
	*pp = 42
	fmt.Println(*pp)
	fmt.Println(x.b)
	fmt.Println(x)

	pT := uintptr(unsafe.Pointer(new(bool)))
	fmt.Println(pT)
}
