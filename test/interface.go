package main

import (
	"fmt"
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
	fmt.Println(p,f.name,f.date)
	return len(p),nil
}
func (f file) Write(p []byte) (n int, err error) {
	fmt.Println(p,f.name,f.date)
	return len(p),nil
}

func main(){
	f := file{name:"liu",date:time.Now()}
	fmt.Println(f.login2)
	// var p = []byte{1,1}
	// f.Read(p)
	
	var db interface{}
	fmt.Println(db)
}

type IntSet struct { X string }
func (i *IntSet) String() string {
	return i.X
}


type db interface {}

func (d db) aa() string {
	return "aa"
}