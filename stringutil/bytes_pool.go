package stringutil

// ByteChanPool struct
//
// Usage:
//	bp := stringutil.NewByteChanPool(500, 1024, 1024)
// 	buf:=bp.Get()
// 	defer bp.Put(buf)
//	// use buf do something ...
type ByteChanPool struct {
	c    chan []byte
	w    int
	wcap int
}

// NewByteChanPool instance
func NewByteChanPool(maxSize int, width int, capWidth int) *ByteChanPool {
	return &ByteChanPool{
		c:    make(chan []byte, maxSize),
		w:    width,
		wcap: capWidth,
	}
}

// Get gets a []byte from the BytePool, or creates a new one if none are
// available in the pool.
func (bp *ByteChanPool) Get() (b []byte) {
	select {
	case b = <-bp.c:
	// reuse existing buffer
	default:
		// create new buffer
		if bp.wcap > 0 {
			b = make([]byte, bp.w, bp.wcap)
		} else {
			b = make([]byte, bp.w)
		}
	}
	return
}

// Put returns the given Buffer to the BytePool.
func (bp *ByteChanPool) Put(b []byte) {
	select {
	case bp.c <- b:
		// buffer went back into pool
	default:
		// buffer didn't go back into pool, just discard
	}
}

// Width returns the width of the byte arrays in this pool.
func (bp *ByteChanPool) Width() (n int) {
	return bp.w
}

// WidthCap returns the cap width of the byte arrays in this pool.
func (bp *ByteChanPool) WidthCap() (n int) {
	return bp.wcap
}
