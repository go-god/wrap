package main

import (
	"log"

	"github.com/go-god/wrap"
)

func main() {
	// through wg to exec goroutine
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

	// through chan to exec goroutine
	w := wrap.NewChanWrapper(2)
	w.Wrap(func() {
		log.Println("this is test")
	})

	w.WrapWithRecovery(func() {
		log.Println("exec goroutine with recovery func")
		var s = []string{"a", "b", "c"}
		log.Printf("s[3] = %v", s[3])
	}, func(r interface{}) {
		// exec recover:runtime error: index out of range [3] with length 3
		log.Printf("exec recover:%v", r)
	})

	w.Wait()
}
