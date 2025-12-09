package day9

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
    raw := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`
	task := Task{}
	res1, res2 := task.Solve(strings.Split(raw, "\n"))
	assert.Equal(t, 50, res1)
	assert.Equal(t, 24, res2)
}