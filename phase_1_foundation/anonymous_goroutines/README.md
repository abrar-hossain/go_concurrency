# Anonymous Goroutines with WaitGroup

Today, I learned how to launch multiple goroutines using anonymous functions inside a loop. I also used `sync.WaitGroup` to ensure that all goroutines finished before the main function exited.

## What I Learned

- An anonymous function is a function without a name, defined and called simultaneously.
- I used a loop to launch 5 goroutines, each printing its own ID.
- `wg.Add(1)` is called before launching each goroutine to track it.
- Inside each goroutine, I used `defer wg.Done()` to mark it as finished.
- `wg.Wait()` in the main function blocks until all goroutines call `Done()`.

## Code Behaviour

- Each loop iteration starts a new goroutine with an anonymous function.
- The current loop value `i` is passed into the function as `id`, so each goroutine prints the correct number.
- All 5 goroutines run concurrently and finish printing their messages.

## Mistake I Could Have Made

If I didn't pass `i` as a parameter to the anonymous function (e.g. used just `fmt.Println(i)`), all goroutines might print the same or the wrong value. Passing `i` as an argument fixes this.

## Summary

Today I learned how to use anonymous goroutines with WaitGroup and how to safely capture loop variables by passing them as function arguments. This helps avoid common bugs in concurrent code.
