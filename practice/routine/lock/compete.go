package lock

import "fmt"

func Compete() {
	var x, y int
	done := make(chan int)
	go func() {
		done <- 0
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()
	go func() {
		done <- 0
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()
	<-done
	<-done
}