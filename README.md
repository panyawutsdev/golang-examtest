# Golang Basic Exam

Solutions to 5 Go programming challenges covering concurrency, thread-safety, interfaces, algorithms, and HTTP.

## Structure

| Folder | Topic |
|--------|-------|
| [01-worker-pool](./01-worker-pool/) | Basic Concurrency: Worker Pool |
| [02-safe-counter](./02-safe-counter/) | Thread-Safety: Safe Counter |
| [03-shapes](./03-shapes/) | Interfaces: Shape Area System |
| [04-two-sum](./04-two-sum/) | Logic & Map: Two Sum |
| [05-http-handler](./05-http-handler/) | HTTP Handler: JSON API |

## Running Each Solution

```bash
cd <folder>
go run main.go
```

---

### 1. Worker Pool
Creates `numWorkers` concurrent goroutines that process `numJobs` jobs from a channel. Uses `sync.WaitGroup` to block until all jobs complete.

### 2. Safe Counter
A `SafeCounter` struct with `sync.Mutex` ensuring `Inc()` and `Value()` are race-condition-free across 1,000+ concurrent goroutines.

### 3. Shapes
`Shape` interface with `Area() float64`. Implemented by `Rectangle` (Width × Height) and `Circle` (π × r²). `PrintArea(s Shape)` accepts any shape.

### 4. Two Sum
Returns indices of two numbers summing to `target`. Uses a hash map for O(n) time complexity — no nested loops.

### 5. HTTP Handler
POST `/hello` endpoint on `:8080`. Accepts `{"name": "..."}`, responds `{"message": "Hello ..."}`. Returns `405` for wrong method, `400` for bad JSON.

#### Test with curl:
```bash
curl -X POST http://localhost:8080/hello \
  -H "Content-Type: application/json" \
  -d '{"name":"Somchai"}'
# {"message":"Hello Somchai"}
```
