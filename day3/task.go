package day3

import (
	"math"
	"strconv"
)

type Task struct{}

func (t *Task) Solve(input []string) (any, any) {
	return solveOne(input), solveTwo(input)
}

func (t *Task) Name() string {
	return "task3"
}

func solveOne(input []string) int {
	res := 0
	for _, bank := range input {
		first, second := '0', '0'
		for index, val := range bank {
			if val > first && index != len(bank)-1 {
				first = val
				second = '0'
			} else if val > second {
				second = val
			}
		}
		res += int(first-'0')*10 + int(second-'0')
	}
	return res
}

func solveTwo(input []string) int {
	res := 0
	for _, bank := range input {
		val := dp(0, 0, bank, make(map[[2]int]int))
		res += val
	}
	return res
}

func dp(pos, taken int, bank string, memo map[[2]int]int) int {
	if taken == 12 {
		return 0
	}
	key := [2]int{pos, taken}
	if val, ok := memo[key]; ok {
		return val
	}
	if len(bank) - pos == 12 - taken {
		res, _ := strconv.Atoi(bank[pos:])
		return res
	}
	cur := int(bank[pos] - '0') * int(math.Pow(10, float64(12 - taken - 1)))
	take := cur + dp(pos + 1, taken + 1, bank, memo)
	skip := dp(pos + 1, taken, bank, memo)
	memo[key] = max(take, skip)
	return max(take, skip)
}
