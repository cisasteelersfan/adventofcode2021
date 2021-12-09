package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

func getLowPoints(arr [][]int) []int {
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	rows, cols := len(arr), len(arr[0])
	lowPoints := make([]int, 0)
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
				lowPoints = append(lowPoints, arr[row][col])
			}
		}
	}
	return lowPoints
}

func sum(arr []int) int {
	ans := 0
	for _, x := range arr {
		ans = ans + x
	}
	return ans
}

func main() {
	iPtr := flag.String("input", "input.txt", "Input filename to read the puzzle input from.")
	flag.Parse()
	dat, _ := ioutil.ReadFile("day9/" + *iPtr)
	heightmap := parseInput(string(dat))
	lowPoints := getLowPoints(heightmap)
	fmt.Println("Part 1:", sum(lowPoints)+len(lowPoints))
}
