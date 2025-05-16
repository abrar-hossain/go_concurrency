package main

import (
	"fmt"
	//"sync"
)

// func doTask(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Task #%d complete\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// for i := 0; i < 3; i++ {
// 	// 	wg.Add(1)
// 	// 	go doTask(i, &wg)
// 	// }
// 	wg.Add(1)
// 	go doTask(1, &wg)

// 	wg.Add(1)
// 	go doTask(2, &wg)

// 	wg.Add(1)
// 	go doTask(3, &wg)

// 	wg.Wait()
// }

func doTask(id int, done chan struct{}) {
	fmt.Printf("Task #%d complete\n", id)
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})

	go doTask(1, done)
	<-done // wait for task 1

	go doTask(2, done)
	<-done // wait for task 2

	go doTask(3, done)
	<-done // wait for task 3
}
