package day11

import (
	"strings"
)

type Task struct{}

func (t *Task) Solve(input []string) (any, any) {
	return solveOne(input), solveTwo(input)
}

func (t *Task) Name() string {
	return "task11"
}

func solveOne(input []string) int {
	graph := make(map[string]map[string]struct{})
	for _, line := range input {
		from := strings.Split(line, ":")[0]
		toSlice := strings.Split(line[5:], " ")
		graph[from] = make(map[string]struct{})
		for _, to := range toSlice {
			graph[from][to] = struct{}{}
		}
	}
	return dfs("you", "out", graph)
}

func dfs(cur, target string, graph map[string]map[string]struct{}) int {
	if cur == target {
		return 1
	}
	res := 0
	for next := range graph[cur] {
		res += dfs(next, target, graph)
	}
	return res
}

func solveTwo(input []string) int {
		graph := make(map[string]map[string]struct{})
	for _, line := range input {
		from := strings.Split(line, ":")[0]
		toSlice := strings.Split(line[5:], " ")
		graph[from] = make(map[string]struct{})
		for _, to := range toSlice {
			graph[from][to] = struct{}{}
		}
	}
	svrDac := dfs2("svr", "dac", graph, make(map[string]int))
	dacFft := dfs2("dac", "fft", graph, make(map[string]int))
	fftOut := dfs2("fft", "out", graph, make(map[string]int))
	svrFft := dfs2("svr", "fft", graph, make(map[string]int))
	fftDac := dfs2("fft", "dac", graph, make(map[string]int))
	dacOut := dfs2("dac", "out", graph, make(map[string]int))

	return svrDac * dacFft * fftOut + svrFft * fftDac * dacOut
}


func dfs2(cur, target string, graph map[string]map[string]struct{}, memo map[string]int) int {
	if cur == target {
		return 1
	}
	if val, ok := memo[cur]; ok {
		return val
	}
	res := 0
	for next := range graph[cur] {
		res += dfs2(next, target, graph, memo)
	}
	memo[cur] = res
	return res
}