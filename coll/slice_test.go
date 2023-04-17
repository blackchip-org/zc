package coll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	type nv struct {
		Name string
		Val  int
	}
	xs := []nv{{"foo", 1}, {"bar", 2}, {"baz", 3}}
	names := Map(xs, func(x nv) string { return x.Name })
	assert.Equal(t, []string{"foo", "bar", "baz"}, names)

	vals := Map(xs, func(x nv) int { return x.Val })
	assert.Equal(t, []int{1, 2, 3}, vals)
}

func TestReverse(t *testing.T) {
	assert.Equal(t, []int{4, 3, 2, 1}, Reverse([]int{1, 2, 3, 4}))
}
