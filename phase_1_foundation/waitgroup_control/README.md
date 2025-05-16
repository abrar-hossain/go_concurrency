# WaitGroup with Manual Goroutine Control

Today I learned how to use `sync.WaitGroup` to manage goroutines, both with and without loops. I also explored how to control the order in which goroutines run.

## What I Learned

- You don't have to use a loop to launch multiple goroutines — you can call them manually.
- Each time I launch a goroutine, I must call `wg.Add(1)` before and `wg.Done()` at the end of the function.
- Without proper `wg.Add()`, the program may exit before the goroutines run.
- WaitGroup doesn't control the order of execution — it only waits for all tasks to finish.
- To run goroutines one after another, I used a `chan struct{}` to wait for a signal from each task before starting the next.

## Code Experiment

I first launched 3 goroutines separately (without a loop) and waited using `WaitGroup`. The order was not guaranteed.

Then I used a `chan struct{}` to launch and wait for each goroutine one at a time. This forced a strict order: task 1, task 2, task 3.

## Mistake I Could Have Made

If I had used WaitGroup alone and expected the tasks to run in order, I would have been confused. WaitGroup waits for all to finish but does not enforce any order.

## Summary

WaitGroup helps manage goroutine lifecycle, but not order. If I need ordered execution, I can use channels to control when each goroutine starts and ends.
