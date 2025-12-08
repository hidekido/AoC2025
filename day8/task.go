package day8

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Task struct{}

func (t *Task) Solve(input []string) (any, any) {
	return solveOne(input), solveTwo(input)
}

func (t *Task) Name() string {
	return "task8"
}

type dist struct {
	value float64
	p1    int
	p2    int
}

func solveOne(input []string) int {
	dots := make([][]int, 0)
	dists := make([]dist, 0)
	for index, row := range input {
		point := strings.Split(row, ",")
		x, _ := strconv.Atoi(point[0])
		y, _ := strconv.Atoi(point[1])
		z, _ := strconv.Atoi(point[2])
		for index2, dot := range dots {
			x1, y1, z1 := dot[0], dot[1], dot[2]
			val := math.Sqrt(math.Pow(float64(x-x1), 2) + math.Pow(float64(y-y1), 2) + math.Pow(float64(z-z1), 2))
			dists = append(dists, dist{value: val, p1: index, p2: index2})
		}
		dots = append(dots, []int{x, y, z})
	}

	slices.SortFunc(dists, func(d1, d2 dist) int {
		if d1.value < d2.value {
			return -1
		} else if d1.value > d2.value {
			return 1
		}
		return 0
	})
	parents := make([]int, len(dots))
	rank := make([]int, len(dots))
	for i := range dots {
		parents[i] = i
		rank[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if x != parents[x] {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}
	union := func(x, y int) {
		x = find(x)
		y = find(y)
		if x == y {
			return
		}
		if rank[x] > rank[y] {
			parents[y] = x
			rank[x] += rank[y]
			rank[y] = 0
		} else {
			parents[x] = y
			rank[y] += rank[x]
			rank[x] = 0
		}

	}
	for i := range 10 {
		dist := dists[i]
		union(dist.p1, dist.p2)
	}
	slices.Sort(rank)
	return rank[len(rank)-1] * rank[len(rank)-2] * rank[len(rank)-3]
}

func solveTwo(input []string) int {

	dots := make([][]int, 0)
	dists := make([]dist, 0)
	for index, row := range input {
		point := strings.Split(row, ",")
		x, _ := strconv.Atoi(point[0])
		y, _ := strconv.Atoi(point[1])
		z, _ := strconv.Atoi(point[2])
		for index2, dot := range dots {
			x1, y1, z1 := dot[0], dot[1], dot[2]
			val := math.Sqrt(math.Pow(float64(x-x1), 2) + math.Pow(float64(y-y1), 2) + math.Pow(float64(z-z1), 2))
			dists = append(dists, dist{value: val, p1: index, p2: index2})
		}
		dots = append(dots, []int{x, y, z})
	}

	slices.SortFunc(dists, func(d1, d2 dist) int {
		if d1.value < d2.value {
			return -1
		} else if d1.value > d2.value {
			return 1
		}
		return 0
	})
	parents := make([]int, len(dots))
	rank := make([]int, len(dots))
	for i := range dots {
		parents[i] = i
		rank[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if x != parents[x] {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}
	union := func(x, y int) int {
		x = find(x)
		y = find(y)
		if x == y {
			return rank[x]
		}
		if rank[x] > rank[y] {
			parents[y] = x
			rank[x] += rank[y]
			rank[y] = 0
			return rank[x]
		} else {
			parents[x] = y
			rank[y] += rank[x]
			rank[x] = 0
			return rank[y]
		}

	}
	for _, dist := range dists {
		res := union(dist.p1, dist.p2)
		if res == len(dots) {
			return dots[dist.p1][0] * dots[dist.p2][0]
		}
	}
	return 0
}
