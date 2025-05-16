# Basic Channels — Send and Receive

In this program, I learned how to send multiple values from a goroutine to the main function using a channel, and how to receive them using a `for range` loop.

## What I Learned

- A channel lets one goroutine send data and another goroutine receive it.
- I used `make(chan int)` to create a channel that carries integers.
- The sender goroutine used a loop to send 5 values through the channel.
- I used `close(ch)` to tell the receiver that no more values will be sent.
- The `for value := range ch` loop in `main()` keeps receiving values until the channel is closed.

## Code Behavior

- A goroutine runs `sendNumber()` and sends values 1 to 5 into the channel.
- The `main()` function listens to the channel and prints each value it receives.
- Once all values are received and the channel is closed, the loop ends.

## Mistake I Could Have Made

If I had not closed the channel, the `range` loop in `main()` would wait forever, causing the program to hang. Closing the channel is important when using `range`.

## Summary

Today I practiced sending multiple values through a channel from a goroutine and receiving them in the main function using `range`. I also learned the importance of closing the channel after sending is done.
