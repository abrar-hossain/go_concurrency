# What is Concurrency? What’s a Goroutine?

This was the first step in understanding concurrency in Go. I learned how goroutines work and how to use them with sync.WaitGroup.

## What I Learned

- A goroutine is a lightweight thread managed by Go. It runs functions concurrently.
- I can start a goroutine using the `go` keyword before a function call.
- When I launch multiple goroutines, they run independently and the order in which they finish is not guaranteed.
- To make sure the main function waits for all goroutines to finish, I used a `sync.WaitGroup`.
- `wg.Add(1)` tells the WaitGroup that a new goroutine is running.
- `wg.Done()` tells it that a goroutine is finished.
- `wg.Wait()` blocks until all goroutines have called `Done()`.

## Code Explanation

I wrote a loop that launches 10 goroutines. Each one prints its own ID. I used a WaitGroup to ensure the main function didn’t exit before the goroutines finished.

## Mistake I Could Have Made

If I didn’t use `wg.Wait()`, the main function might finish before the goroutines print anything. That would cause the program to exit early.

## Summary

Today I learned how to create multiple goroutines and manage their execution using WaitGroup. This is the foundation for all future concurrency in Go.
