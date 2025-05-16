package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine #%d\n", id)
		}(i)

	}

	wg.Wait()
}
