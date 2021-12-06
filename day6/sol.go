package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getFishMap(s string) map[int]int {
	numStr := strings.Split(s, ",")
	arr := make([]int, len(numStr))
	for i := range numStr {
		num, _ := strconv.Atoi(numStr[i])
		arr[i] = num
	}
	fm := make(map[int]int)
	for _, num := range arr {
		fm[num]++
	}
	return fm
}

func getSum(fm map[int]int) int {
	sum := 0
	for _, count := range fm {
		sum = sum + count
	}
	return sum
}

func main() {
	iPtr := flag.String("input", "input.txt", "Input filename to read the puzzle input from.")
	flag.Parse()
	dat, _ := ioutil.ReadFile("day6/" + *iPtr)
	fishMap := getFishMap(string(dat))

	ansAt80 := 0
	for i := 0; i < 256; i++ {
		newFishMap := make(map[int]int)
		for d := 1; d <= 8; d++ {
			newFishMap[d-1] = fishMap[d]
		}
		newFishMap[6] = newFishMap[6] + fishMap[0]
		newFishMap[8] = newFishMap[8] + fishMap[0]
		fishMap = newFishMap
		if i == 79 {
			ansAt80 = getSum(fishMap)
		}
	}
	fmt.Println("Part 1:", ansAt80)
	fmt.Println("Part 2:", getSum(fishMap))
}
