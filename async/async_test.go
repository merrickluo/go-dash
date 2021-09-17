package async_test

import (
	"testing"

	"github.com/merrickluo/go-dash/async"
	"github.com/stretchr/testify/assert"
)

func TestCollectInt(t *testing.T) {
	ch := make(chan int, 8)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	assert.Equal(t, []int{1, 2, 3}, async.Collect(ch))
}

func TestIntoInt(t *testing.T) {
	ch1 := make(chan int, 10)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)

	res := []int{}
	async.Into(ch1, &res)
	assert.Equal(t, []int{1, 2, 3}, res)
}

func TestTakeInt2(t *testing.T) {
	ch := make(chan int, 5)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	taken := async.Take(ch, 2)
	assert.Equal(t, []int{1, 2}, async.Collect(taken))
}

func TestDrop2(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	dropped := async.Drop(ch, 2)
	assert.Equal(t, []int{3}, async.Collect(dropped))
}

func TestMergeInt(t *testing.T) {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	ch1 <- 1
	ch2 <- 2
	ch1 <- 3
	ch2 <- 4
	ch1 <- 5
	close(ch1)
	close(ch2)

	merged := async.Merge(ch1, ch2)
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, async.Collect(merged))
}

func TestSplitInt(t *testing.T) {
	ch := make(chan int)
	sp1, sp2 := async.Split(ch, func(i int) bool {
		return i > 2
	})

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)

	assert.ElementsMatch(t, []int{3, 4}, async.Collect(sp1))
	assert.ElementsMatch(t, []int{1, 2}, async.Collect(sp2))
}

func TestMapInt(t *testing.T) {
	ch := make(chan int, 10)
	mappedCh := async.Map(ch, func(it int) int {
		return it * 2
	})

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	assert.Equal(t, []int{2, 4, 6}, async.Collect(mappedCh))
}
func TestFilterInt(t *testing.T) {
	ch := make(chan int, 10)
	filtered := async.Filter(ch, func(i int) bool {
		return i%2 == 0
	})

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	assert.Equal(t, []int{2}, async.Collect(filtered))

}

func TestReduceInt(t *testing.T) {
	ch := make(chan int, 10)
	reduced := async.Reduce(ch, func(acc int, i int) int {
		return acc + i
	}, 0)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	assert.Equal(t, []int{6}, async.Collect(reduced))
}

func TestSlidingBuffer(t *testing.T) {
	ch := make(chan int)
	sb := async.SlidingBuffer(ch, 2)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	assert.Equal(t, []int{2, 3}, async.Collect(sb))
}

func TestDroppingBuffer(t *testing.T) {
	ch := make(chan int)
	db := async.DroppingBuffer(ch, 2)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	assert.Equal(t, []int{1, 2}, async.Collect(db))
}
