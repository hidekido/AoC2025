package day6

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
    raw := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	task := Task{}
	res1, res2 := task.Solve(strings.Split(raw, "\n"))
	assert.Equal(t, 4277556, res1)
	assert.Equal(t, 3263827, res2)
}