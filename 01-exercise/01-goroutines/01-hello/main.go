package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("Go routine 01")
	// goroutine with anonymous function
	go func() {
		fun("Go routine 02")
	}()
	// goroutine with function value call
	fv := fun
	go fv("Go routine 03")
	// wait for goroutines to end
	fmt.Println("waiting for goroutines to complete..")
	time.Sleep(3 * time.Second)
	fmt.Println("done..")
}
