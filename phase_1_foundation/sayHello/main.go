package main

import (
	"fmt"
	"sync"
)

func sayHello(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it finishes
	fmt.Printf("Hello from goroutine #%d\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)           // Add 1 to WaitGroup counter
		go sayHello(i, &wg) // Launch goroutine
	}

	wg.Wait() // Wait until all goroutines call Done()
}
