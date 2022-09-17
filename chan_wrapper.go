package wrap

// chWrapper chan wrapper
type chWrapper struct {
	bufCap int
	buf    chan struct{}
}

var est = struct{}{}

// NewChanWrapper create chan wgWrapper instance
// to execute goroutine through channel synchronization, the capacity must be specified in advance
func NewChanWrapper(bufCap int) *chWrapper {
	if bufCap <= 0 {
		panic("the buf cap must be greater than 0")
	}

	w := &chWrapper{bufCap: bufCap}
	w.buf = make(chan struct{}, w.bufCap)
	return w
}

// Wrap fn func in goroutine to run without recovery func
func (c *chWrapper) Wrap(fn func()) {
	go func() {
		defer c.done()
		fn()
	}()
}

// WrapWithRecovery fn func in goroutine to run with customized recovery func
func (c *chWrapper) WrapWithRecovery(fn func(), recoveryFunc func(r interface{})) {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				recoveryFunc(e)
			}
		}()

		defer c.done()
		fn()
	}()
}

// Wait blocks until all goroutine finished.
func (c *chWrapper) Wait() {
	for i := 0; i < c.bufCap; i++ {
		<-c.buf
	}
}

func (c *chWrapper) done() {
	c.buf <- est
}
