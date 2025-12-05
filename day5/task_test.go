package day5

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
    raw := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	task := Task{}
	res1, res2 := task.Solve(strings.Split(raw, "\n"))
	assert.Equal(t, 3, res1)
	assert.Equal(t, 14, res2)
}