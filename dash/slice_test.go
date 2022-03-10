package dash_test

import (
	"testing"

	"github.com/merrickluo/go-dash/dash"
	"github.com/stretchr/testify/assert"
)

func TestMapInt(t *testing.T) {
	result := dash.Map([]int{1, 2, 3}, func(item int) int {
		return item * 2
	})
	assert.Equal(t, []int{2, 4, 6}, result)
}

func TestFlatMapInt(t *testing.T) {
	result := dash.FlatMap([]int{1, 2, 3}, func(it int) []int {
		return []int{it, it * 2}
	})
	assert.Equal(t, []int{1, 2, 2, 4, 3, 6}, result)
}

func TestFilterInt(t *testing.T) {
	result := dash.Filter([]int{1, 2, 3}, func(it int) bool {
		return it > 2
	})
	assert.Equal(t, []int{3}, result)
}

func TestReduceInt(t *testing.T) {
	result := dash.Reduce([]int{1, 2, 3}, func(acc int, it int) int {
		return acc + it
	}, 0)
	assert.Equal(t, 6, result)
}

func TestTakeInt(t *testing.T) {
	assert.Equal(t, []int{1, 2}, dash.Take([]int{1, 2, 3}, 2))
}

func TestDropInt(t *testing.T) {
	assert.Equal(t, []int{3}, dash.Drop([]int{1, 2, 3}, 2))
}

func TestIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 3}
	c := []int{2, 3}

	assert.ElementsMatch(t, []int{1, 2, 3}, dash.Intersection(a))
	assert.ElementsMatch(t, []int{1, 3}, dash.Intersection(a, b))
	assert.ElementsMatch(t, []int{2, 3}, dash.Intersection(a, c))
	assert.Equal(t, []int{3}, dash.Intersection(a, b, c))
}

func TestUniq(t *testing.T) {
	a := []int{1, 1, 2, 3}
	b := []int{1, 3, 3, 2}

	assert.Equal(t, []int{1, 2, 3}, dash.Uniq(a))
	assert.Equal(t, []int{1, 3, 2}, dash.Uniq(b))
}

func odd(i int) bool {
	return i%2 == 1
}

func TestEverySomeNone(t *testing.T) {
	a := []int{1, 3, 5}
	b := []int{1, 3, 3, 2}
	c := []int{4, 6, 2}

	assert.Equal(t, true, dash.Every(a, odd))
	assert.Equal(t, false, dash.Every(b, odd))
	assert.Equal(t, false, dash.Every(c, odd))

	assert.Equal(t, true, dash.Some(a, odd))
	assert.Equal(t, true, dash.Some(b, odd))
	assert.Equal(t, false, dash.Some(c, odd))

	assert.Equal(t, false, dash.None(a, odd))
	assert.Equal(t, false, dash.None(b, odd))
	assert.Equal(t, true, dash.None(c, odd))
}

func TestGroupBy(t *testing.T) {
	a := []int{1, 3, 3, 2}
	result := map[bool][]int{
		true:  []int{1, 3, 3},
		false: []int{2},
	}

	assert.Equal(t, result, dash.GroupBy(a, odd))
}

func TestShuffle(t *testing.T) {
	a := []int{1, 3, 3, 2, 5}

	b := dash.Shuffle(a)
	c := dash.Shuffle(a)
	d := dash.Shuffle(c)

	assert.ElementsMatch(t, b, a)
	assert.ElementsMatch(t, c, a)
	assert.ElementsMatch(t, d, a)
}

func TestInclude(t *testing.T) {
	a := []int{1, 3}

	assert.Equal(t, false, dash.Include(a, 2))
	assert.Equal(t, true, dash.Include(a, 1))
}

func TestCycle(t *testing.T) {
	a := []int{1, 3}

	assert.Equal(t, []int{}, dash.Cycle(a, 0))
	assert.Equal(t, []int{1, 3}, dash.Cycle(a, 1))
	assert.Equal(t, []int{1, 3, 1, 3}, dash.Cycle(a, 2))
}

func TestParallelMap(t *testing.T) {
	a := []int{1, 2, 3}

	assert.Equal(t, []int{2, 3, 4}, dash.ParallelMap(a, func(i int) int {
		return i + 1
	}))
}

func TestParallelMapN(t *testing.T) {
	a := []int{1, 2, 3}

	assert.Equal(t, []int{2, 3, 4}, dash.ParallelMapN(a, func(i int) int {
		return i + 1
	}, 2))
}
