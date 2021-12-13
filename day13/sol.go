package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type paper struct {
	points        map[point]bool
	width, length int
}

type point struct {
	x, y int
}

func parseArray(s string) paper {
	rows := strings.Split(s, "\n")
	points := make(map[point]bool)
	maxX, maxY := 0, 0
	for _, row := range rows {
		split := strings.Split(row, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		points[point{x, y}] = true
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
	return paper{points, maxX, maxY}
}

func foldAlongX(x int, pap *paper) {
	for p, noSkip := range pap.points {
		if !noSkip {
			continue
		}
		if p.x >= x {
			newX := (x) - (p.x - x)
			pap.points[p] = false
			pap.points[point{newX, p.y}] = true
		}
	}
}

func foldAlongY(y int, pap *paper) {
	for p, noSkip := range pap.points {
		if !noSkip {
			continue
		}
		if p.y >= y {
			newY := (y) - (p.y - y)
			pap.points[p] = false
			pap.points[point{p.x, newY}] = true
		}
	}
}

func parseFolds(s string) [][]int {
	ret := make([][]int, 0)
	for _, row := range strings.Split(s, "\n") {
		garbage := strings.Fields(row)
		split := strings.Split(garbage[2], "=")
		num, _ := strconv.Atoi(split[1])
		first := 0
		if split[0] == "y" {
			first = 1
		}
		ret = append(ret, []int{first, num})
	}
	return ret
}

func countTrues(p *paper) int {
	ans := 0
	for _, v := range p.points {
		if v {
			ans++
		}
	}
	return ans
}

func getString(p paper) string {
	var sb strings.Builder
	maxX, maxY := 0, 0
	for k, v := range p.points {
		if !v {
			continue
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.y > maxY {
			maxY = k.y
		}
	}
	arr := make([][]string, maxY+1)
	for i := range arr {
		arr[i] = make([]string, maxX+1)
		for j := range arr[i] {
			arr[i][j] = "."
		}
	}
	for k, v := range p.points {
		if !v {
			continue
		}
		arr[k.y][k.x] = "#"
	}
	for i := range arr {
		for j := range arr[0] {
			sb.WriteString(arr[i][j])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func main() {
	iPtr := flag.String("input", "input.txt", "Input filename to read the puzzle input from.")
	flag.Parse()
	dat, _ := ioutil.ReadFile("day13/" + *iPtr)
	split := strings.Split(string(dat), "\n\n")
	p := parseArray(split[0])
	folds := parseFolds(split[1])
	if folds[0][0] == 0 {
		foldAlongX(folds[0][1], &p)
	} else {
		foldAlongY(folds[0][1], &p)
	}
	fmt.Println("Part 1:", countTrues(&p))

	p = parseArray(split[0])
	for _, fold := range folds {
		if fold[0] == 0 {
			foldAlongX(fold[1], &p)
		} else {
			foldAlongY(fold[1], &p)
		}
	}
	fmt.Println("Part 2:")
	fmt.Println(getString(p))
}
