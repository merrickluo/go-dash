package main

type mMult[T any] struct {
	source chan T
	taps   map[chan T]bool
	tap    chan chan T
	untap  chan chan T
}

func mult[T any](source chan T) *mMult[T] {
	m := &mMult[T]{
		source: source,
		taps:   make(map[chan T]bool),
		tap:    make(chan chan T),
		untap:  make(chan chan T),
	}
	go m.run()
	return m
}

func tap[T any](m *mMult[T], ch chan T) {
	m.tap <- ch
}

func untap[T any](m *mMult[T], ch chan T) {
	m.untap <- ch
}

func (m *mMult[T]) run() {
	for {
		select {
		case it := <-m.source:
			for ch, _ := range m.taps {
				ch <- it
			}
		case t := <-m.tap:
			m.taps[t] = true
		case t := <-m.untap:
			delete(m.taps, t)
		}
	}
}