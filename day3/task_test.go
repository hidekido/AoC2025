package day3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
    raw := `987654321111111
811111111111119
234234234234278
818181911112111`
	task := Task{}
	res1, res2 := task.Solve(strings.Split(raw, "\n"))
	assert.Equal(t, 357, res1)
	assert.Equal(t, 3121910778619, res2)
}