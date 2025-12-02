package day2

import (
	"strconv"
	"strings"
)

type Task struct{}

func (t *Task) Solve(input []string) (any, any) {
	return solveOne(input), solveTwo(input)
}

func (t *Task) Name() string {
	return "task2"
}

func solveOne(input []string) int {
	res := 0
	for r := range strings.SplitSeq(input[0], ",") {

		vals := strings.Split(r, "-")
		from, _ := strconv.Atoi(vals[0])
		to, _ := strconv.Atoi(vals[1])
		for val := from; val <= to; val++ {
			if invalid(val, 2) {
				res += val
			}
		}
	}
	return res
}

func solveTwo(input []string) int {
	res := 0
	for r := range strings.SplitSeq(input[0], ",") {
		vals := strings.Split(r, "-")
		from, _ := strconv.Atoi(vals[0])
		to, _ := strconv.Atoi(vals[1])
		for val := from; val <= to; val++ {
			if invalid(val, 100) {
				res += val
			}
		}
	}
	return res
}

func invalid(val int, cap int) bool {
	str := strconv.Itoa(val)
	for parts := 2; parts <= min(len(str), cap); parts++ {
		if len(str)%parts != 0 {
			continue
		}
		size := len(str) / parts
		base := str[:size]
		ok := true
		for i := range parts {
			if base != str[i * size : (i + 1) * size] {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}

	return false
}
