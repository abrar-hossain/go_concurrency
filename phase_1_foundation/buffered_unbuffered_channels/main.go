package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	fmt.Println("Sending 10")
	ch <- 10
	fmt.Println("Sending 20")
	ch <- 20
	fmt.Println("Sending 30")
	ch <- 30
	// fmt.Println("Sending 30")
	// ch <- 40

	fmt.Println("All values sent to buffer")

	// Receiving values with delay
	time.Sleep(1 * time.Second)
	fmt.Println("Receiving:", <-ch)

	time.Sleep(1 * time.Second)
	fmt.Println("Receiving:", <-ch)

	time.Sleep(1 * time.Second)
	fmt.Println("Receiving:", <-ch)

	// time.Sleep(1 * time.Second)
	// fmt.Println("Receiving:", <-ch)
}
