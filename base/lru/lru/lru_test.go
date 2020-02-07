package lru

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type Val struct {
	name string
}

func TestNewLRU(t *testing.T) {
	l := NewLRU(3)
	// 3 --> 2 --> 1
	l.Set("1", Val{name:"cao"})
	l.Set("2", Val{name:"hai"})
	l.Set("3", Val{name:"xi"})

	// get
	// 1 --> 3 --> 2
	v := l.Get("1")
	assert.Equal(t, "cao", v.(Val).name)
	assert.Equal(t, 3, l.count)

	// 4 --> 1 --> 3
	l.Set("4", Val{name:"jing"})
	assert.Equal(t, 3, l.count)
	v2 := l.Get("2")
	assert.Equal(t, nil, v2)
}
