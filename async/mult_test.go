package async_test

import (
	"testing"

	"github.com/merrickluo/go-dash/async"
	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	g := make(chan int)
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	m := async.Mult(g)
	async.Tap(m, ch1)
	async.Tap(m, ch2)

	g <- 10

	async.Untap(m, ch1)

	g <- 20
	close(g)
	close(ch1)
	close(ch2)

	assert.Equal(t, []int{10}, async.Collect(ch1))
	assert.Equal(t, []int{10, 20}, async.Collect(ch2))
}
