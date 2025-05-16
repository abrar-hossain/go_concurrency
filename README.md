# 🐹 Go Concurrency Mastery Roadmap

This roadmap is designed to build your Go concurrency skills from beginner to expert — with daily 30-minute challenges.

---

## Phase 1: Foundation 
**Goal**: Understand goroutines, channels, and basic synchronization.

| Day | Concept                               | Practice Problem                                                  |
|-----|----------------------------------------|-------------------------------------------------------------------|
| 1   | What is concurrency? What's a goroutine? | Write a function that runs `sayHello()` 10 times using goroutines. |
| 2   | Basic goroutines and anonymous functions | Launch 5 goroutines that print their ID ("Goroutine 1", etc.)     |
| 3   | Channels: basic send/receive             | Send a value from one goroutine to another using a channel.       |
| 4   | Buffered vs unbuffered channels         | Write a buffered channel with capacity 3 and experiment with sends/receives. |
| 5   | WaitGroup: waiting for goroutines       | Launch 3 goroutines and use `sync.WaitGroup` to wait until all are done. |

---

## Phase 2: Inter-Goroutine Communication 
**Goal**: Learn to make goroutines talk to each other safely.

| Day | Concept                                | Problem                                                            |
|-----|----------------------------------------|--------------------------------------------------------------------|
| 6   | Directional channels (send-only/recv-only) | Create a producer-consumer system with typed channels.            |
| 7   | Range over channels                     | Use a `for range` loop to read from a channel until it's closed.  |
| 8   | Select statement (Go's secret weapon)   | Use `select` to read from two channels simultaneously.             |
| 9   | Timeouts and `time.After()`             | Add a timeout to a slow channel read.                              |
| 10  | Done channels and graceful shutdown     | Build a loop that exits when it receives a signal on a "done" channel. |

---

## Phase 3: Real-World Patterns 
**Goal**: Learn Go concurrency patterns used in production.

| Day | Concept               | Problem                                                         |
|-----|------------------------|------------------------------------------------------------------|
| 11  | Fan-out, fan-in pattern | Launch 5 workers that process tasks concurrently, gather results. |
| 12  | Worker pools           | Build a worker pool that processes 100 jobs with 5 workers.     |
| 13  | Pipeline pattern       | Build a 3-stage pipeline: gen → square → print.                 |
| 14  | Semaphore pattern      | Use a buffered channel to limit concurrency to 3 goroutines.    |
| 15  | Throttling with ticker | Allow 1 request per second using a ticker.                      |
| 16  | Goroutine leaks        | Simulate a goroutine leak, then fix it using cancellation.      |
| 17  | Context package        | Use `context.WithCancel` to cancel long-running work.           |

---

## Phase 4: Advanced Systems Thinking 
**Goal**: Write real systems that coordinate complex behavior.

| Day | Concept                    | Project                                                              |
|-----|----------------------------|----------------------------------------------------------------------|
| 18  | Error handling in goroutines | Have workers report errors through a channel.                       |
| 19  | Shared memory               | Use `sync.Mutex` to protect shared data in concurrent access.       |
| 20  | sync.Once, sync.Map        | Use `sync.Once` to init data only once, and `sync.Map` for concurrent maps. |
| 21  | Build a mini job scheduler | Schedule tasks to run after N seconds using `time.Timer`.           |
| 22  | Re-implement time.Sleep()  | Build your own sleep function using `select` + `time.After`.        |
| 23  | Write a mini web crawler   | Crawl URLs concurrently, avoiding duplicate visits.                 |
| 24  | Build a concurrent counter web server | Multiple users incrementing counter without race conditions. |
| 25  | Context + fan-out + error group | Build a search engine simulator with context and graceful shutdown. |

---

## Phase 5: Interview-Level Mastery 
**Goal**: Crush any Go concurrency question with confidence.

| Day | Focus                                  | Mock Interview Question                                         |
|-----|----------------------------------------|------------------------------------------------------------------|
| 26  | Debugging deadlocks and race conditions | Why does this program deadlock? Fix it.                         |
| 27  | Compare concurrency vs parallelism      | How does Go's scheduler work? Explain GOMAXPROCS.               |
| 28  | Whiteboard question: thread-safe cache  | Use goroutines, mutex, context, etc.                            |
| 29  | Real-world: log aggregator system       | Multiple producers → one consumer, rate-limited.                |
| 30  | Final project: Rebuild WaitGroup        | Rebuild `sync.WaitGroup` using channels only.                   |

---

## 🛠 Tools & Setup

- Use **Go Playground** for small experiments: https://go.dev/play/
- Use **VS Code** or **GoLand** with Go plugin for real projects.
- Run with race detection:  
  ```bash
  go run -race myfile.go
  ```

You don't need a course. You need **daily focus, discipline, and fire.** .

I will be here every step of the way.
