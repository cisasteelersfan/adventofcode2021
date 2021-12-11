package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type stack [][]int

func (s *stack) Push(arr []int) {
	*s = append(*s, arr)
}

func (s *stack) Pop() []int {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

type point struct {
	energy, flashes int
}

func parseInput(s string) [][]point {
	lines := strings.Split(s, "\n")
	points := make([][]point, len(lines))
	for i, line := range lines {
		points[i] = make([]point, len(line))
		for j, char := range line {
			num, _ := strconv.Atoi(string(char))
			points[i][j] = point{energy: num, flashes: 0}
		}
	}
	return points
}

func step(p [][]point) {
	flashes := stack{}
	rows, cols := len(p), len(p[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			p[row][col].energy++
			if p[row][col].energy == 10 {
				p[row][col].energy = -1 // flashed during this step
				p[row][col].flashes++
				flashes.Push([]int{row, col})
			}
		}
	}
	for len(flashes) > 0 {
		f := flashes.Pop()
		row, col := f[0], f[1]
		for _, n := range [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
			rowAdd, colAdd := n[0], n[1]
			if row+rowAdd >= 0 && row+rowAdd < rows && col+colAdd >= 0 && col+colAdd < cols {
				if p[row+rowAdd][col+colAdd].energy == -1 {
					continue
				}
				p[row+rowAdd][col+colAdd].energy++
				if p[row+rowAdd][col+colAdd].energy == 10 {
					p[row+rowAdd][col+colAdd].energy = -1
					p[row+rowAdd][col+colAdd].flashes++
					flashes.Push([]int{row + rowAdd, col + colAdd})
				}
			}
		}
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if p[row][col].energy == -1 {
				p[row][col].energy = 0
			}
		}
	}
}

func countFlashes(p [][]point) int {
	flashes := 0
	rows, cols := len(p), len(p[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			flashes = flashes + p[row][col].flashes
		}
	}
	return flashes
}

func printPoints(p [][]point) string {
	var sb strings.Builder
	rows, cols := len(p), len(p[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sb.WriteString(strconv.Itoa(p[row][col].energy))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func allZero(p [][]point) bool {
	rows, cols := len(p), len(p[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if p[row][col].energy != 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	iPtr := flag.String("input", "input.txt", "Input filename to read the puzzle input from.")
	flag.Parse()
	dat, _ := ioutil.ReadFile("day11/" + *iPtr)
	points := parseInput(string(dat))
	for i := 0; i < 100; i++ {
		step(points)
	}
	fmt.Println("Part 1:", countFlashes(points))

	points = parseInput(string(dat))
	for i := 0; i < 100000; i++ {
		if allZero(points) {
			fmt.Println("Part 2:", i)
			break
		}
		step(points)
	}
}
