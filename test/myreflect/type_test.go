package myreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrintType(t *testing.T) {
	PrintType()
}

func TestDisplay(t *testing.T) {

	Display("strangelove", GetObj())
}

func TestShow(t *testing.T) {
	x := 2                   // value   type    variable?
	a := reflect.ValueOf(2)  // 2       int     no
	b := reflect.ValueOf(x)  // 2       int     no
	c := reflect.ValueOf(&x) // &x      *int    no
	d := c.Elem()            // 2       int     yes (x)
	fmt.Println(a.CanAddr()) // "false"
	fmt.Println(b.CanAddr()) // "false"
	fmt.Println(c.CanAddr()) // "false"
	fmt.Println(d.CanAddr()) // "true"
}

func TestReflectValue(t *testing.T) {
	x := 2
	d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
	px := d.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3
	fmt.Println(x)                    // "3"
}

func TestPrintObj(t *testing.T) {
	Print(GetObj())
}
