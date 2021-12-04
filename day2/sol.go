package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	iPtr := flag.String("input", "input.txt", "Input filename to read the puzzle input from.")
	flag.Parse()
	dat, _ := ioutil.ReadFile("day2/" + *iPtr)
	raw := strings.Split(string(dat), "\n")
	curX := 0
	curDepth := 0
	for _, s := range raw {
		r := strings.Split(s, " ")
		command := r[0]
		distance, _ := strconv.Atoi(r[1])
		switch command {
		case "forward":
			curX = curX + distance
			break
		case "down":
			curDepth = curDepth + distance
			break
		case "up":
			curDepth = curDepth - distance
			break
		}
	}
	fmt.Println("Part 1:", curX*curDepth)

	curX = 0
	curDepth = 0
	aim := 0
	for _, s := range raw {
		r := strings.Split(s, " ")
		command := r[0]
		distance, _ := strconv.Atoi(r[1])
		switch command {
		case "forward":
			curX = curX + distance
			curDepth = curDepth + (aim * distance)
			break
		case "down":
			aim = aim + distance
			break
		case "up":
			aim = aim - distance
			break
		}
	}
	fmt.Println("Part 2:", curX*curDepth)
}
