package main_test

import (
	"fmt"
	"testing"

	d "github.com/merrickluo/godash"
	"github.com/stretchr/testify/assert"
)

func TestMapInt(t *testing.T) {
	result := d.Map([]int{1, 2, 3}, func(item int) int {
		return item * 2
	})
	assert.Equal(t, result, []int{2, 4, 6})
}

func TestMapString(t *testing.T) {
	result := d.Map([]int{1, 2, 3}, func(i int) string {
		return fmt.Sprintf("%d", i)
	})
	assert.Equal(t, result, []string{"1", "2", "3"})
}
