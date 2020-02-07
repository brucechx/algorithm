package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var Input1 = IntSlice{
	data: []int{1, 3, 5, 7, 8, 10, 12},
}

func TestIntSlice_BinarySearch1(t *testing.T) {
	res := Input1.BinarySearch1(10)
	assert.Equal(t, 5, res)
}

func TestIntSlice_BinarySearch2(t *testing.T) {
	res := Input1.BinarySearch2(8)
	assert.Equal(t, 4, res)
}
