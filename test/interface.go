package main

import (
	"fmt"
	"time"
)


type user interface {
	Login() string
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

type file struct {
	name string
	date time.Time
}

func (f file) Login() string {
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
func measure(u user){
	var p = []byte{2,3}
	fmt.Println(u.Read(p))
}

func main(){
	f := file{name:"liu",date:time.Now()}
	var p = []byte{1,1}
	f.Read(p)
	
	measure(f)
}