package wrap

import (
	"sync"
)

// wrapper wrap go WaitGroup
type wrapper struct {
	wg sync.WaitGroup
}

// New create wrapper instance
func New() *wrapper {
	w := &wrapper{}
	return w
}

// Wrap fn func in goroutine to run without recovery func
func (w *wrapper) Wrap(fn func()) {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		fn()
	}()
}

// WrapWithRecovery fn func in goroutine to run with customized recovery func
func (w *wrapper) WrapWithRecovery(fn func(), recoveryFunc func(r interface{})) {
	w.wg.Add(1)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				recoveryFunc(e)
			}
		}()

		defer w.wg.Done()
		fn()
	}()
}

// Wait blocks until the WaitGroup counter is zero.
func (w *wrapper) Wait() {
	w.wg.Wait()
}
