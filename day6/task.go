package day6

import (
	"strconv"
	"strings"
)


type Task struct{}

func (t *Task) Solve(input []string) (any, any) {
	return solveOne(input), solveTwo(input)
}

func (t *Task) Name() string {
	return "task6"
}

func solveOne(input []string) int {
	res := 0
	ops := make([]string, 0)
	for _, val := range strings.Split(input[len(input)-1], " ") {
		val = strings.TrimSpace(val)
		if val == "" {
			continue
		}
		ops = append(ops, val)
	}
	cols := make([]int, len(ops))
	for index := 0; index < len(input) - 1; index++ {
		line := strings.Split(input[index], " ")
		i := 0
		for _, val := range line {
			val = strings.TrimSpace(val)
			if val == "" {
				continue
			}
			valInt, _ := strconv.Atoi(val)
			if index == 0 {
				cols[i] = valInt
			} else if ops[i] == "*" {
				cols[i] *= valInt
			} else {
				cols[i] += valInt
			}
			i++
		}
	}
	for _, val := range cols {
		res += val
	}
	return res
}

func solveTwo(input []string) int {
	res := 0
	op := ""
	vals := make([]int, 0)
	for j := range input[0] {
		seen := false
		val := 0
		for i := range input {
			if input[i][j] == '*' {
				op = "*"
				continue
			} else if input[i][j] == '+' {
				op = "+"
				continue
			}
			if input[i][j] == ' ' {
				continue
			}
			dig := int(input[i][j] - '0')
			val = val * 10 + dig
			seen = true
		}
		if !seen {
			res += calc(vals, op)
			vals = make([]int, 0)
		} else {
			vals = append(vals, val)
		}
	}
	res += calc(vals, op)
	return res
}

func calc(vals []int, op string) int {
	res := 0 
	if op == "*" {
		res = 1
	}
	for _, val := range vals {
		if op == "*" {
			res *= val
		} else {
			res += val
		}
	}
	return res
}
