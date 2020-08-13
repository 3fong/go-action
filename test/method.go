package main

import (
	"fmt"
	"math"
)

// Point 结构体
type Point struct{ X, Y float64 }

// Distance function
func Distance(p, q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance but as a method of the Point type
func (p Point) Distance(q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// I 基础类型
type I int
// Sum 求和
func (i I) Sum(q I) int {
    return int(i) + int(q)
}

func main(){
	printDistance()
}

func printDistance() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call
}