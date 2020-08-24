package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

type user interface {
	login() string
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

type file struct {
	name string
	date time.Time
}

func (f file) login() string {
	return "len(p)"
}

func (f *file) login2() string {
	return "len(p)"
}
func (f file) Read(p []byte) (n int, err error) {
	fmt.Println(p, f.name, f.date)
	return len(p), nil
}
func (f file) Write(p []byte) (n int, err error) {
	fmt.Println(p, f.name, f.date)
	return len(p), nil
}

func main() {
	// f := file{name:"liu",date:time.Now()}
	// fmt.Println(f.login2)
	// var p = []byte{1,1}
	// f.Read(p)
	compare()
}

type IntSet struct{ X string }

func (i *IntSet) String() string {
	return i.X
}

var period = flag.Duration("period", 1*time.Second, "sleep period")

func receiveParam() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

func compare() {
	var w io.Writer
	fmt.Printf("%T\n", w)
	var i *io.Writer
	fmt.Printf("%T\n", i)
	w = os.Stdout
	fmt.Printf("%T\n", w)
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)
	w = nil
	fmt.Printf("%T\n", w)
	var a interface{} = 1
	fmt.Println(a == a)

	var x interface{} = []int{1, 2, 3}
	fmt.Println(x == x)
}
