package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getPositions(s string) []int {
	strPos := strings.Split(string(s), ",")
	arr := make([]int, len(strPos))
	for i, strNum := range strPos {
		num, _ := strconv.Atoi(strNum)
		arr[i] = num
	}
	return arr
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getFuelCost(arr []int, pos int) int {
	cost := 0
	for _, num := range arr {
		cost = cost + abs(pos-num)
	}
	return cost
}

func getMaxPos(arr []int) int {
	max := arr[0]
	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}

func getMinPos(arr []int) int {
	min := arr[0]
	for _, i := range arr {
		if i < min {
			min = i
		}
	}
	return min
}

func main() {
	iPtr := flag.String("input", "input.txt", "Input filename to read the puzzle input from.")
	flag.Parse()
	dat, _ := ioutil.ReadFile("day7/" + *iPtr)
	positions := getPositions(string(dat))

	// Use binary search to find the minimum.
	minPos, maxPos := getMinPos(positions), getMaxPos(positions)
	for minPos <= maxPos {
		mid := (minPos + maxPos) / 2
		cost := getFuelCost(positions, mid)
		left := getFuelCost(positions, mid-1)
		right := getFuelCost(positions, mid+1)
		if cost < left && cost < right {
			fmt.Println("Part 1:", cost)
			break
		}
		if cost < left {
			minPos = mid + 1
		} else if cost < right {
			maxPos = mid - 1
		} else {
			panic("weird ordering")
		}
	}
}
