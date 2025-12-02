package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
    raw := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	task := Task{}
	res1, res2 := task.Solve([]string{raw})
	assert.Equal(t, 1227775554, res1)
	assert.Equal(t, 4174379265, res2)
}