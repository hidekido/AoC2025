package day1

import "strconv"

type Task struct {} 

func (t *Task) Solve(input []string) (any, any) {
	return solveOne(input), solveTwo(input)
}

func (t *Task) Name() string {
	return "task1"
}

func solveOne(input []string) int {
	pos := 50
	res := 0
	for _, line := range input {
		sign := 1
		if line[0] == 'L' {
			sign = -1
		}
		val, _ := strconv.Atoi(line[1:])
		pos = (pos + sign * val) % 100
		if pos == 0 {
			res++
		}
	}
	return res
}

func solveTwo (input []string) int {
	pos := 50
	res := 0
	for _, line := range input {
		sign := 1
		if line[0] == 'L' {
			sign = -1
		}
		val, _ := strconv.Atoi(line[1:])
		for range val {
			pos += sign
			if pos % 100 == 0 {
				res++
			}
		}
		pos = pos % 100
	}
	return res
}