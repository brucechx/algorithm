package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	input []int
	output []int
}

func cases() []Case{
	return []Case{
		{
			input:[]int{2, 5, 1, 7, 3, 8},
			output:[]int{1, 2, 3, 5, 7, 8},
		},
		{
			input:[]int{4, 6, 1, 7, 9, 8},
			output:[]int{1, 4, 6, 7, 8, 9},
		},
		{
			input:[]int{14, 26, 11, 47, 1109, 118},
			output:[]int{11, 14, 26, 47, 118, 1109},
		},
	}
}

func TestIntSlice_All_Sort(t *testing.T) {
	for _, cas := range cases(){
		intSlice := IntSlice{data:cas.input}
		for _, mod := range allMods{
			res := intSlice.Sort(mod)
			assert.Equal(t, cas.output, res)
		}
	}
}