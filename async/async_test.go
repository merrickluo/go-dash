package async

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectInt(t *testing.T) {
	ch := make(chan int, 8)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	assert.Equal(t, []int{1, 2, 3}, Collect(ch))
}

func TestIntoInt(t *testing.T) {
	ch1 := make(chan int, 10)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)

	res := []int{}
	Into(ch1, &res)
	assert.Equal(t, []int{1, 2, 3}, res)
}

func TestTakeInt2(t *testing.T) {
	ch := make(chan int, 5)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	taken := Take(ch, 2)
	assert.Equal(t, []int{1, 2}, Collect(taken))
}

func TestDrop2(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	dropped := Drop(ch, 2)
	assert.Equal(t, []int{3}, Collect(dropped))
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

	merged := Merge(ch1, ch2)
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5}, Collect(merged))
}

func TestSplitInt(t *testing.T) {
	ch := make(chan int)
	sp1, sp2 := Split(ch, func(i int) bool {
		return i > 2
	})

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)

	assert.ElementsMatch(t, []int{3, 4}, Collect(sp1))
	assert.ElementsMatch(t, []int{1, 2}, Collect(sp2))
}

func TestMapInt(t *testing.T) {
	ch := make(chan int, 10)
	mappedCh := MMap(ch, func(it int) int {
		return it * 2
	})

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	assert.Equal(t, []int{2, 4, 6}, Collect(mappedCh))
}