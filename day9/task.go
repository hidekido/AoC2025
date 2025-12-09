package day9

import (
	"strconv"
	"strings"
)

type Task struct{}

func (t *Task) Solve(input []string) (any, any) {
	return solveOne(input), solveTwo(input)
}

func (t *Task) Name() string {
	return "task9"
}

type dist struct {
	value float64
	p1    int
	p2    int
}

func solveOne(input []string) int {
	res := 0
	dots := make([][]int, 0)
	for _, line := range input {
		x, y := strings.Split(line, ",")[0], strings.Split(line, ",")[1]
		xi, _ := strconv.Atoi(x)
		yi, _ := strconv.Atoi(y)
		for _, d := range dots {
			res = max(res, (abs(d[0] - xi) + 1) * (abs(d[1] - yi) + 1))
		}
		dots = append(dots, []int{xi, yi})
	}
	return res
}

func abs(x int) int {
	return max(x, -x)
}

func solveTwo(input []string) int {
	res := 0
	dots := make([][]int, 0)
	n, m := 100000, 100000
	for _, line := range input {
		x, y := strings.Split(line, ",")[0], strings.Split(line, ",")[1]
		xi, _ := strconv.Atoi(x)
		yi, _ := strconv.Atoi(y)
		dots = append(dots, []int{xi, yi})
	}
	field := make([][]int, n + 1)
	for i := range n + 1 {
		field[i] = make([]int, m + 1)
	}

	prev := dots[0]
	for _, p := range dots {
		connect(prev, p, field)
		prev = p
	}
	connect(prev, dots[0], field)
	for i := 0; i < len(dots); i++ {
		for j := i + 1; j < len(dots); j++ {
			if verify(dots[i], dots[j], field, dots) {
				val := (abs(dots[i][0] - dots[j][0]) + 1) * (abs(dots[i][1] - dots[j][1]) + 1)
				if val > res {
					res = val
				}
			}
		}
	}
	
	return res
}

func connect(from, to []int, field [][]int) {
	if from[0] == to[0] {
		if from[1] > to[1] {
			from, to = to, from
		}
		for j := from[1]; j <= to[1]; j++ {
			field[from[0]][j] = 1
		}
	} else {
		if from[0] > to[0] {
			from, to = to, from
		}
		for i := from[0]; i <= to[0]; i++ {
			field[i][from[1]] = 1
		}
	}
}

func fill(i, j int, field [][]int) bool {
	if i == 0 || i == len(field) - 1 || j == 0 || j == len(field[0]) - 1 {
		return field[i][j] == 1
	}
	if field[i][j] == 1 {
		return true
	}
	field[i][j] = 1
	for x := i - 1; x <= i + 1; x++ {
		for y := j - 1; y <= j + 1; y++ {
			if !fill(x, y, field) {
				return false
			}
		}
	}
	return true
}

func verify(d1, d2 []int, field [][]int, dots [][]int) bool {
	if field[d1[0]][d2[1]] != 1 && field[d2[0]][d1[1]] != 1 {
		return false
	}
	ok, _, _, _, _ := inside(d1[0], d2[1], field)
	if !ok {
		return false
	}
	ok, _, _, _, _ = inside(d2[0], d1[1], field)
	if !ok {
		return false
	}
	mn_i, mx_i := min(d1[0], d2[0]), max(d1[0], d2[0])
	mn_j, mx_j := min(d1[1], d2[1]), max(d1[1], d2[1])
	ok, bot, top, right, left := inside(mn_i + (mx_i - mn_i) / 2, mn_j + (mx_j - mn_j) / 2, field)
	if !ok {
		return false
	}
	if bot < mx_i {
		return false
	}
	if top > mn_i {
		return false
	}
	if right < mx_j {
		return false
	}
	if left > mn_j {
		return false
	}
	for _, dot := range dots {
		if dot[0] > min(d1[0], d2[0]) && dot[0] < max(d1[0], d2[0]) && dot[1] > min(d1[1], d2[1]) && dot[1] < max(d1[1], d2[1]) {
			return false
		}
	}
	return true
}

func inside(x, y int, field [][]int) (bool, int, int, int, int) {
	bot := 0
	for i := x; i < len(field); i++ {
		if field[i][y] == 1 {
			bot = i
			break
		}
		if i == len(field) - 1 {
			return false, 0, 0, 0, 0
		}
	}
	top := 0
	for i := x; i >= 0; i-- {
		if field[i][y] == 1 {
			top = i
			break
		}
		if i == 0 {
			return false, 0, 0, 0, 0
		}
	}

	right := 0
	for j := y; j < len(field[0]); j++ {
		if field[x][j] == 1 {
			right = j
			break
		}
		if j == len(field[0]) - 1 {
			return false, 0, 0, 0, 0
		}
	}
	left := 0
	for j := y; j >= 0; j-- {
		if field[x][j] == 1 {
			left = j
			break
		}
		if j == 0 {
			return false, 0, 0, 0, 0
		}
	}
	return true, bot, top, right, left
}