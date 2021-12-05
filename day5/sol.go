package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type line struct {
	x1, y1, x2, y2 int
}

func (l line) isHorizVert() bool {
	return l.x1 == l.x2 || l.y1 == l.y2
}

func (l line) getPoints() []point {
	ans := make([]point, 0)
	if l.x1 == l.x2 {
		start, end := l.y1, l.y2
		if end < start {
			start, end = end, start
		}
		for y := start; y <= end; y++ {
			ans = append(ans, point{l.x1, y})
		}
	} else if l.y1 == l.y2 {
		start, end := l.x1, l.x2
		if end < start {
			start, end = end, start
		}
		for x := start; x <= end; x++ {
			ans = append(ans, point{x, l.y1})
		}
	} else {
		if l.y1 < l.y2 {
			for x := l.x1; x <= l.x2; x++ {
				ans = append(ans, point{x, l.y1 + (x - l.x1)})
			}
		} else {
			for x := l.x1; x <= l.x2; x++ {
				ans = append(ans, point{x, l.y1 - (x - l.x1)})
			}
		}
	}
	return ans
}

func parseLine(s string) line {
	split := strings.Split(s, " -> ")
	first := strings.Split(split[0], ",")
	x1, _ := strconv.Atoi(first[0])
	y1, _ := strconv.Atoi(first[1])
	second := strings.Split(split[1], ",")
	x2, _ := strconv.Atoi(second[0])
	y2, _ := strconv.Atoi(second[1])
	if x2 < x1 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}
	return line{x1, y1, x2, y2}
}

func main() {
	iPtr := flag.String("input", "input.txt", "Input filename to read the puzzle input from.")
	flag.Parse()
	dat, _ := ioutil.ReadFile("day5/" + *iPtr)
	raw := strings.Split(string(dat), "\n")
	lines := make([]line, len(raw))
	for i, line := range raw {
		lines[i] = parseLine(line)
	}
	points := make(map[point]int)
	for _, l := range lines {
		if !l.isHorizVert() {
			continue
		}
		for _, p := range l.getPoints() {
			points[p]++
		}
	}
	numPointsOverlapping := 0
	for _, val := range points {
		if val > 1 {
			numPointsOverlapping++
		}
	}
	fmt.Println("Part 1:", numPointsOverlapping)

	points = make(map[point]int)
	for _, l := range lines {
		for _, p := range l.getPoints() {
			points[p]++
		}
	}
	numPointsOverlapping = 0
	for _, val := range points {
		if val > 1 {
			numPointsOverlapping++
		}
	}
	fmt.Println("Part 2:", numPointsOverlapping)
}
