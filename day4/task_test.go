package day4

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
    raw := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`
	task := Task{}
	res1, res2 := task.Solve(strings.Split(raw, "\n"))
	assert.Equal(t, 13, res1)
	assert.Equal(t, 43, res2)
}