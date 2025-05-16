package main

import "fmt"

func sendNumber(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	channel := make(chan int)

	go sendNumber(channel)

	for value := range channel {
		fmt.Println("Received:", value)
	}

}
