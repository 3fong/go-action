package main

import (
	"fmt"
)



func main(){

	fmt.Println(structTest())
}

func sliceTest(){
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)  
}

func nonempty(strings []string) []string {
    i := 0
    for _, s := range strings {
        if s != "" {
            strings[i] = s
            i++
        }
    }
    return strings[:i]
}

func reverseSliceTest(){
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse2(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
}

func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// 重写reverse函数，使用数组指针代替slice。
func reverse2(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// Point 定义结构体
type Point struct{ X, Y int }

func structTest() *Point {
    p := Point{1, 2}
    return &p
}
