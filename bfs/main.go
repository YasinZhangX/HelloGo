package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (p point) add(dir point) point {
	return point{p.i + dir.i, p.j + dir.j}
}

func (p point) getValue(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[0]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func (p point) setValue(grid [][]int, v int) {
	grid[p.i][p.j] = v
}

func main() {
	maze := readMaze("maze.in")

	printMatrix(maze)

	fmt.Println()

	minStep, steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	printMatrix(steps)

	fmt.Println("Minimun step: ", minStep)
}

func printMatrix(m [][]int) {
	for _, row := range m {
		for _, val := range row {
			fmt.Printf("%2d ", val)
		}
		fmt.Println()
	}
}

func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		eatLF(file)
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

func eatLF(file *os.File) {
	var t rune
	fmt.Fscanf(file, "\n", &t)
}

func walk(maze [][]int, start, end point) (int, [][]int) {
	max := len(maze) * len(maze[0])

	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
		for j := range steps[i] {
			steps[i][j] = max
		}
	}
	start.setValue(steps, 0)

	var cur point
	queue := []point{start}

	for len(queue) > 0 {
		cur, queue = popFirst(queue)
		curStep := steps[cur.i][cur.j]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			val, ok := next.getValue(maze)
			step, _ := next.getValue(steps)

			if ok && val != 1 && curStep+1 < step {
				queue = append(queue, next)
				next.setValue(steps, curStep+1)
			}
		}
	}

	result, _ := end.getValue(steps)

	return result, steps
}

func popFirst(path []point) (point, []point) {
	cur := path[0]
	path = path[1:]
	return cur, path
}
