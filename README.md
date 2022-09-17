# wrapper
wrap go `sync.WaitGroup` processing goroutine.

# why wrap
- Wrap the go `sync.WaitGroup` so that you don't sometimes forget to write `defer wg.Add(1)` before the business layer writes go func, or `defer wg.Done()` when it's Done.
- By wrapping it, you can ensure that the Go Goroutine executes successfully.

# example
```go
package main

import (
	"log"

	"github.com/go-god/wrap"
)

func main() {
	var wg = wrap.New()
	wg.Wrap(func() {
		log.Println("this is test")
	})

	for i := 0; i < 10; i++ {
		num := i
		// wrap go goroutine without recovery func
		wg.Wrap(func() {
			log.Println("current index:", num)
		})
	}

	wg.WrapWithRecovery(func() {
		log.Println("exec goroutine with recovery func")
		var s = []string{"a", "b", "c"}
		log.Printf("s[3] = %v", s[3])
	}, func(r interface{}) {
		// exec recover:runtime error: index out of range [3] with length 3
		log.Printf("exec recover:%v", r)
	})

	wg.Wait()
}
```
