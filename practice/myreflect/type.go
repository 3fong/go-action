package myreflect

import (
	"fmt"
	"reflect"
)

func PrintType() {
	t := reflect.TypeOf(3)  // a reflect.Type
	fmt.Println(t.String()) // "int"
	fmt.Println(t)          // "int"
}
