package async

type mult struct {
	source chan int
	taps   map[chan int]bool
	tap    chan chan int
	untap  chan chan int
}

func Mult(source chan int) *mult {
	m := &mult{
		source: source,
		taps:   make(map[chan int]bool),
		tap:    make(chan chan int),
		untap:  make(chan chan int),
	}
	go m.run()
	return m
}

func Tap(m *mult, ch chan int) {
	m.tap <- ch
}

func Untap(m *mult, ch chan int) {
	m.untap <- ch
}

func (m *mult) run() {
	for {
		select {
		case int := <-m.tap:
			m.taps[int] = true
		case int := <-m.untap:
			delete(m.taps, int)
		case it, ok := <-m.source:
			if !ok {
				m.source = nil
				break
			}
			for ch := range m.taps {
				ch <- it
			}
		}
		if m.source == nil {
			close(m.tap)
			close(m.untap)
			break
		}
	}
}
