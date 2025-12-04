package day4

type Task struct{}

func (t *Task) Solve(input []string) (any, any) {
	return solveOne(input), solveTwo(input)
}

func (t *Task) Name() string {
	return "task4"
}

func solveOne(input []string) int {
	res := 0
	n, m := len(input), len(input[0])
	field := make([][]byte, n)
	for i := range n {
		field[i] = []byte(input[i])
	}
	for row := range n {
		for column := range m {
			if input[row][column] != '@' {
				continue
			}
			if count(row, column, field) < 4 {
				res++
			}
		}
	}
	return res
}

func count(r, c int, input [][]byte) int {
	res := 0
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if i < 0 || i >= len(input) || j < 0 || j >= len(input[0]) {
				continue
			}
			if input[i][j] == '@' {
				res++
			}
		}
	}
	return res - 1
}

func solveTwo(input []string) int {
	res := 0
	n, m := len(input), len(input[0])
	field := make([][]byte, n)
	for i := range n {
		field[i] = []byte(input[i])
	}
	prev := -1
	for prev < res {
		prev = res
		for row := range n {
			for column := range m {
				if field[row][column] != '@' {
					continue
				}
				if count(row, column, field) < 4 {
					field[row][column] = 'x'
					res++
				}
			}
		}
	}
	return res
}
