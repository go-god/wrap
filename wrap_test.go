package wrap

import (
	"log"
	"testing"
)

func TestWrapper(t *testing.T) {
	var wg = New()
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

/*
=== RUN   TestWrapper
2022/09/17 18:26:07 exec goroutine with recovery func
2022/09/17 18:26:07 current index: 5
2022/09/17 18:26:07 current index: 0
2022/09/17 18:26:07 current index: 2
2022/09/17 18:26:07 exec recover:runtime error: index out of range [3] with length 3
2022/09/17 18:26:07 current index: 1
2022/09/17 18:26:07 current index: 3
2022/09/17 18:26:07 current index: 4
2022/09/17 18:26:07 current index: 7
2022/09/17 18:26:07 current index: 6
2022/09/17 18:26:07 current index: 8
2022/09/17 18:26:07 current index: 9
2022/09/17 18:26:07 this is test
--- PASS: TestWrapper (0.00s)
PASS
*/

func TestNewChanWrapper(t *testing.T) {
	w := NewChanWrapper(2)

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

/*
=== RUN   TestNewChanWrapper
2022/09/18 00:15:18 exec goroutine with recovery func
2022/09/18 00:15:18 this is test
--- PASS: TestNewChanWrapper (0.00s)
2022/09/18 00:15:18 exec recover:runtime error: index out of range [3] with length 3
PASS
*/
