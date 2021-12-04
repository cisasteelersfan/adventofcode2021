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
	dat, _ := ioutil.ReadFile("day3/" + *iPtr)
	raw := strings.Split(string(dat), "\n")
	numLines := len(raw)
	width := len(raw[0])
	posToOnesCount := make(map[int]int)
	for _, s := range raw {
		for pos, bit := range s {
			if string(bit) == "1" {
				posToOnesCount[pos]++
			}
		}
	}
	mostCommonBits := make([]int, width)
	for i := 0; i < width; i++ {
		if posToOnesCount[i] > (numLines / 2) {
			mostCommonBits[i] = 1
		} else {
			mostCommonBits[i] = 0
		}
	}
	leastCommonBits := make([]int, width)
	for i, b := range mostCommonBits {
		leastCommonBits[i] = 1 - b
	}
	mcbString := strings.Join(toStringArray(mostCommonBits), "")
	lcbString := strings.Join(toStringArray(leastCommonBits), "")
	mcbNum, _ := strconv.ParseInt(mcbString, 2, 0)
	lcbNum, _ := strconv.ParseInt(lcbString, 2, 0)
	fmt.Println("Part 1:", mcbNum*lcbNum)
}

func toStringArray(arr []int) []string {
	s := make([]string, len(arr))
	for i, num := range arr {
		s[i] = strconv.Itoa(num)
	}
	return s
}
