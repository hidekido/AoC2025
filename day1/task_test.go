package day1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
    raw := "L68,L30,R48,L5,R60,L55,L1,L99,R14,L82"
	test := strings.Split(raw, ",")
	task := Task{}
	res1, res2 := task.Solve(test)
	assert.Equal(t, 3, res1)
	assert.Equal(t, 6, res2)
}