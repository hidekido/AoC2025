package day7

type Task struct{}

func (t *Task) Solve(input []string) (any, any) {
	return solveOne(input), solveTwo(input)
}

func (t *Task) Name() string {
	return "task7"
}

func solveOne(input []string) int {
	res := 0
	n, m := len(input), len(input[0])
	field := make([][]byte, n)
	for i := range n {
		field[i] = []byte(input[i])
	}
	for i := range n - 1 {
		for j := range m {
			if field[i][j] == '.' || field[i][j] == '^' {
				continue
			}
			if field[i+1][j] == '.' {
				field[i+1][j] = '|'
			}
			if field[i+1][j] == '^' {
				res++
				if j > 0 {
					field[i+1][j-1] = '|'
				}
				if j < m-1 {
					field[i+1][j+1] = '|'
				}
			}
		}
	}
	return res
}

func solveTwo(input []string) int {
	n, m := len(input), len(input[0])
	field := make([][]byte, n)
	for i := range n {
		field[i] = []byte(input[i])
	}
	for i := range n - 1 {
		for j := range m {
			if field[i][j] == '.' || field[i][j] == '^' {
				continue
			}
			if field[i+1][j] == '.' {
				field[i+1][j] = '|'
			}
			if field[i+1][j] == '^' {
				if j > 0 {
					field[i+1][j-1] = '|'
				}
				if j < m-1 {
					field[i+1][j+1] = '|'
				}
			}
		}
	}
	counts := make([]int, m)
	for j := range m {
		if field[n - 1][j] == '|' {
			counts[j] = 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		next := make([]int, m)
		for j := range m {
			if field[i][j] == '|' {
				next[j] = counts[j]
			}
			if field[i][j] == '^' {
				if j > 0 {
					next[j] += counts[j - 1]
				} 
				if j < m - 1 {
					next[j] += counts[j + 1]
				}
			}
			if field[i][j] == 'S' {
				return counts[j]
			}
		}
		counts = next
	}
	return 0
}
