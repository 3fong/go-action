package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Commencing countdown.")
	// tick := time.Tick(1 * time.Second)
	// for countdown := 10; countdown > 0; countdown-- {
	// 	fmt.Println(countdown)
	// 	<-tick
	// }
	// launch()
	ch := make(chan int, 1)
for i := 0; i < 10; i++ {
    select {
    case x := <-ch:
        fmt.Println(i) // "0" "2" "4" "6" "8"
        fmt.Println(x) // "0" "2" "4" "6" "8"
    case ch <- i:
    }
}
}

func launch() {
	fmt.Println("Lift off!")
}
