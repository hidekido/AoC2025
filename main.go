package main

import (
	"bufio"
	"fmt"
	"main/day1"
	"main/day2"
	"main/day3"
	"main/day4"
	"main/day5"
	"main/day6"
	"os"
	"slices"
)

type Task interface {
	Solve(input []string) (any, any)
	Name() string
}

// registry holds all registered tasks.
var registry = make(map[string]Task)

func register() {
	tasks := []Task{
		&day1.Task{},
		&day2.Task{},
		&day3.Task{},
		&day4.Task{},
		&day5.Task{},
		&day6.Task{},
	}
	for _, task := range tasks {
		name := task.Name()
		registry[name] = task
	}
}

func main() {
	register()
	if len(os.Args) < 2 {
		fmt.Println("Usage: run <task> [args...]")
		fmt.Println("Available tasks:", tasks())
		return
	}

	taskName := os.Args[1]
	filepath := os.Args[2]

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Failed to open a file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Println(registry[taskName].Solve(input))
}

func tasks() []string {
	keys := make([]string, 0)
	for k := range registry {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
