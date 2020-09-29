package test

import (
	"fmt"
	"image/color"
	"math"
)

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

func printDistance() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call
}

// Point 定义结构体
type Point struct{ X, Y float64 }

func structTest() *Point {
	p := Point{1, 2}
	return &p
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func print() {
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p)

	q := Point{1, 2}
	(&q).ScaleBy(2)
	fmt.Println(q)

	q.ScaleBy(2)
	fmt.Println(q)
	var cp = ColoredPoint{Point{1, 2}, color.RGBA{255, 0, 0, 255}}

	fmt.Println(cp.RGBA)
}

type ColoredPoint struct {
	Point
	color.RGBA
}

func bindValue() {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(p.Distance(q))
	distanceFromP := p.Distance   // method value
	fmt.Println(distanceFromP(q)) // "5"
	var origin Point              // {0, 0}
	fmt.Println(distanceFromP(origin))

	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"
}
