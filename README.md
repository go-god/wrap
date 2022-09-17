# wrapper
wrap go `sync.WaitGroup` processing goroutine.

# why wrap
- Wrap the go `sync.WaitGroup` so that you don't sometimes forget to write `defer wg.Add(1)` before the business layer writes go func, or `defer wg.Done()` when it's Done.
- By wrapping it, you can ensure that the Go Goroutine executes successfully.
