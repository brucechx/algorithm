package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	cases := []struct{
		arr []int
		target int
		index int
	}{
		{
			arr:[]int{1, 2, 2, 2, 3, 4, 5, 6},
			target: 3,
			index: 4,
		},
		{
			arr: []int{1, 3, 5, 7, 8, 10, 12},
			target: 10,
			index: 5,
		},
	}
	for i, cas := range cases{
		assert.Equal(t, cas.index, BinarySearch(cas.arr, cas.target))
		assert.Equal(t, cas.index, BinarySearch2(cas.arr, cas.target))
		if i == 0{
			assert.Equal(t, 1, BinarySearchLeftBound(cas.arr, 2))
			assert.Equal(t, 1, BinFirst(cas.arr, 2))
			assert.Equal(t, 3, BinLast(cas.arr, 2))
			assert.Equal(t, 3, BinarySearchRightBound(cas.arr, 2))
		}
	}
}
