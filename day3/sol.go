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
	all := make(map[string]bool)
	posToOnesCount := make(map[int]int)
	for _, s := range raw {
		all[s] = true
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

	oxygenRating := getOxygenRating(all, width)
	fmt.Println("oxygen", oxygenRating)
	scrubberRating := getScrubberRating(all, width)
	fmt.Println("scrubber", scrubberRating)
	fmt.Println("Part 2:", oxygenRating*scrubberRating)
}

func copyMap(arr map[string]bool) map[string]bool {
	ret := make(map[string]bool)
	for k, v := range arr {
		ret[k] = v
	}
	return ret
}

func getOxygenRating(arr map[string]bool, width int) int64 {
	all := copyMap(arr)
	for pos := 0; pos < width; pos++ {
		mostCommonBit := getMcb(all, pos)
		//fmt.Println("all:", all, "pos:", pos, "mcb:", mostCommonBit)
		newAll := make(map[string]bool)
		for k, _ := range all {
			if string(k[pos]) == mostCommonBit {
				newAll[k] = true
			}
		}
		all = newAll
	}
	//fmt.Println("len(all):", len(all))
	for k := range all {
		rating, _ := strconv.ParseInt(k, 2, 0)
		return rating
	}
	panic("len(all) = " + string(len(all)))
}

func getMcb(m map[string]bool, pos int) string {
	numLines := len(m)
	onesCount := 0
	for line := range m {
		if string(line[pos]) == "1" {
			onesCount++
		}
	}
	//fmt.Println("onesCount:", onesCount)
	if onesCount >= (numLines - onesCount) {
		return "1"
	} else {
		return "0"
	}
}

func getLcb(m map[string]bool, pos int) string {
	numLines := len(m)
	zeroCount := 0
	for line := range m {
		if string(line[pos]) == "1" {
			zeroCount++
		}
	}
	if zeroCount >= (numLines - zeroCount) {
		return "0"
	} else {
		return "1"
	}
}

func getScrubberRating(arr map[string]bool, width int) int64 {
	all := copyMap(arr)
	for pos := 0; pos < width; pos++ {
		leastCommonBit := getLcb(all, pos)
		//fmt.Println("all:", all, "pos:", pos, "lcb:", leastCommonBit)
		newAll := make(map[string]bool)
		for k, _ := range all {
			if string(k[pos]) == leastCommonBit {
				newAll[k] = true
			}
		}
		all = newAll
		if len(all) == 1 {
			break
		}
	}
	//fmt.Println("len(all):", len(all))
	for k := range all {
		rating, _ := strconv.ParseInt(k, 2, 0)
		return rating
	}
	panic("len(all) = " + string(len(all)))
}

func toStringArray(arr []int) []string {
	s := make([]string, len(arr))
	for i, num := range arr {
		s[i] = strconv.Itoa(num)
	}
	return s
}
