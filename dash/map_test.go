package dash_test

import (
	"testing"

	"github.com/merrickluo/go-dash/dash"
	"github.com/stretchr/testify/assert"
)

func TestMapKeys(t *testing.T) {
	m := map[string]int{"hello": 1, "world": 2}
	assert.ElementsMatch(t, []string{"hello", "world"}, dash.Keys(m))
}

func TestMapValues(t *testing.T) {
	m := map[string]int{"hello": 1, "world": 2}
	assert.ElementsMatch(t, []int{1, 2}, dash.Values(m))
}

func TestMergeWithPlus(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 10, "b": 20}
	m3 := map[string]int{"c": 3, "a": 100}

	merged := dash.MergeWith(func(n1 int, n2 int) int {
		return n1 + n2
	}, m1, m2, m3)
	assert.Equal(t, map[string]int{"a": 111, "b": 22, "c": 3}, merged)
}
