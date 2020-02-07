package sort_list

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSortList(t *testing.T){
	// 排序前： [4 2 1 3]
	expectedRes := []int{1, 2, 3, 4}
	cases := []struct{
		input *ListNode
		sortType string
	}{
		{initList(), "insert1"}, // 插入
		{initList(), "insert2"},
		{initList(), "select"},  // 选择
		{initList(), "select2"},
		{initList(), "quick"}, // 快排
		{initList(), "quick2"},
		{initList(), "merge"}, // 归并
		{initList(), "bubble"}, // 冒泡

	}

	for _, cas := range cases{
		assert.Equal(t, expectedRes, cas.input.sortList(cas.sortType))
	}
}
