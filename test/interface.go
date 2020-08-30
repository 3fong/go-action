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
	// compare()
	// plot();
	// typeJudge()
	var arr = []string{"a","b","c","d"}
	fmt.Println(arr[1:])

	forValue()
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

func typeJudge(){
	var w io.Writer
w = os.Stdout
f := w.(*os.File)      // success: f == os.Stdout
// c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer

fmt.Println(f)
// fmt.Println(c)

	var w2 io.Writer
w2 = os.Stdout
rw := w2.(io.ReadWriter) // success: *os.File has both Read and Write

w2 = new(ByteCounter)
rw = w2.(io.ReadWriter) // panic: *ByteCounter has no Read method
fmt.Println(rw)
}

type ByteCounter struct  {
	
}

func (b ByteCounter) Write(p []byte) (n int, err error) {
	return 1,nil
}

func forValue(){
	x := []string{"html","body","div","div","h2","p"} 
	y := []string{"div","div","h2"} 
	fmt.Println(containsAll(x,y))
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}