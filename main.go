package main

import (
	"fmt"
	"runtime"

	"github.com/go-action/test/routine/change"
)
var usage = make(map[string]int64)

func main() {
	const user = "joe@example.org"
	usage[user] = 980000000
	fmt.Print(CheckQuota(user))
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
