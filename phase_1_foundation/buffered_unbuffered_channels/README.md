#  Buffered vs Unbuffered Channels

Today I learned the difference between unbuffered and buffered channels in Go.

## What I Learned

- By default, channels in Go are unbuffered. This means that sending blocks until the value is received, and receiving blocks until there is something to receive.
- Buffered channels allow you to send a limited number of values without needing a receiver right away. You can create a buffered channel with `make(chan int, size)` where `size` is the buffer capacity.
- I learned that if I try to send more values than the buffer size without a receiver, the program will block (freeze).
- I also learned that if I try to receive from a channel when there is nothing to receive and no sender, it will block and cause a deadlock.

## Code Experiment

I created a buffered channel with a size of 3 and sent three values to it. Then I used `time.Sleep()` between receives to see how the values stayed in the channel buffer until I read them. This showed that sending to a buffered channel can be instant, and receiving can happen later.

## Mistake I Made

I tried to receive more values than I sent. That caused a deadlock because the program waited for a value that never came. This helped me understand that receivers block if there's nothing in the channel and no one is sending.

## Summary

Buffered channels let you send and receive at different times, but both sides must still respect the buffer limit and content availability. Sending more than the buffer without receiving will block, and receiving more than was sent will also block.
