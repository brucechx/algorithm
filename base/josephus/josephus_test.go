package josephus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	n int
	k int
	res int
}

func cases() []Case{
	return []Case{
		{
			n: 11,
			k: 3,
			res: 6,
		},
		{
			n: 5,
			k: 3,
			res: 3,
		},
	}
}

func Test_josephus1(t *testing.T){
	for _, cas := range cases(){
		res := josephus1(cas.n, cas.k)
		assert.Equal(t, cas.res, res)
	}
}


func Test_josephus2(t *testing.T){
	for _, cas := range cases(){
		res := josephus2(cas.n, cas.k)
		assert.Equal(t, cas.res, res)
	}
}

func Test_josephus3(t *testing.T){
	for _, cas := range cases(){
		res := josephus3(cas.n, cas.k)
		assert.Equal(t, cas.res, res)
	}
}
