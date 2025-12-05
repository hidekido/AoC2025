package day5

import (
	"slices"
	"strconv"
	"strings"
)

type Task struct{}

func (t *Task) Solve(input []string) (any, any) {
	return solveOne(input), solveTwo(input)
}

func (t *Task) Name() string {
	return "task5"
}

func solveOne(input []string) int {
	res := 0
	ranges := make([][]int, 0)
	index := 0
	for input[index] != "" {
		r := strings.Split(strings.TrimSpace(input[index]), "-")
		from, _ := strconv.Atoi(r[0])
		to, _ := strconv.Atoi(r[1])
		ranges = append(ranges, []int{from, to})
		index++
	}
	index++
	for index < len(input) {
		val, _ := strconv.Atoi(input[index])
		for _, r := range ranges {
			if val >= r[0] && val <= r[1] {
				res++
				break
			}
		}
		index++
	}
	return res
}

func solveTwo(input []string) int {
	res := 0
	ranges := make([][]int, 0)
	index := 0
	for input[index] != "" {
		r := strings.Split(strings.TrimSpace(input[index]), "-")
		from, _ := strconv.Atoi(r[0])
		to, _ := strconv.Atoi(r[1])
		ranges = append(ranges, []int{from, to})
		index++
	}
	slices.SortFunc(ranges, func(a, b []int) int {
		return a[0] - b[0]
	})
	last := 0
	for _, r := range ranges {
		res += max(r[1] - max(r[0], last) + 1, 0)
		last = max(last, r[1] + 1)
	}

	return res
}
