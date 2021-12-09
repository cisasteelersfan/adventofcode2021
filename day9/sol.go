package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func parseInput(s string) [][]int {
	split := strings.Split(s, "\n")
	arr := make([][]int, len(split))
	for i, line := range split {
		arr[i] = make([]int, len(line))
		for j, col := range line {
			num, _ := strconv.Atoi(string(col))
			arr[i][j] = num
		}
	}
	return arr
}

func getLowPoints(arr [][]int) [][]int {
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	rows, cols := len(arr), len(arr[0])
	lowPoints := make([][]int, 0)
	for row, r := range arr {
		for col := range r {
			isLowest := true
			for i := range directions {
				rowAdd, colAdd := directions[i][0], directions[i][1]
				if row+rowAdd >= 0 && row+rowAdd < rows && col+colAdd >= 0 && col+colAdd < cols {
					if arr[row][col] >= arr[row+rowAdd][col+colAdd] {
						isLowest = false
						break
					}
				}
			}
			if isLowest {
				lowPoints = append(lowPoints, []int{row, col})
			}
		}
	}
	return lowPoints
}

type point struct {
	row, col int
}

func getBasinSize(arr [][]int, row int, col int, seen map[point]bool) int {
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	rows, cols := len(arr), len(arr[0])
	if seen[point{row, col}] {
		return 0
	}
	seen[point{row, col}] = true
	if row < 0 || row >= rows || col < 0 || col >= cols || arr[row][col] == 9 {
		return 0
	}
	ans := 1
	for i := range directions {
		rowAdd, colAdd := directions[i][0], directions[i][1]
		ans = ans + getBasinSize(arr, row+rowAdd, col+colAdd, seen)
	}
	return ans
}

func main() {
	iPtr := flag.String("input", "input.txt", "Input filename to read the puzzle input from.")
	flag.Parse()
	dat, _ := ioutil.ReadFile("day9/" + *iPtr)
	heightmap := parseInput(string(dat))
	lowPoints := getLowPoints(heightmap)
	sum := 0
	for _, lp := range lowPoints {
		sum = sum + heightmap[lp[0]][lp[1]]
	}
	fmt.Println("Part 1:", sum+len(lowPoints))
	basinSizes := make([]int, len(lowPoints))
	for i, lp := range lowPoints {
		basinSizes[i] = getBasinSize(heightmap, lp[0], lp[1], make(map[point]bool))
	}
	sort.Ints(basinSizes)
	fmt.Println("Part 2:", basinSizes[len(basinSizes)-1]*basinSizes[len(basinSizes)-2]*basinSizes[len(basinSizes)-3])
}
