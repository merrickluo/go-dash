package async

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	g := make(chan int)
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	m := Mult(g)
	Tap(m, ch1)
	Tap(m, ch2)

	g <- 10

	Untap(m, ch1)

	g <- 20
	close(g)
	close(ch1)
	close(ch2)

	assert.Equal(t, []int{10}, Collect(ch1))
	assert.Equal(t, []int{10, 20}, Collect(ch2))
}
