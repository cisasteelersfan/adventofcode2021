package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}
	raw := strings.Split(string(dat), "\n")
	n := []int{}
	for _, r := range raw {
		num, _ := strconv.Atoi(r)
		n = append(n, num)
	}
	c := 0
	for i := 1; i < len(n); i++ {
		if n[i] > n[i-1] {
			c++
		}
	}
	fmt.Println("Part 1:", c)

	c = 0
	lastThreeTotal := 0
	for i := 0; i < 3; i++ {
		lastThreeTotal = lastThreeTotal + n[i]
	}
	for i := 3; i < len(n); i++ {
		prev := lastThreeTotal
		lastThreeTotal = lastThreeTotal - n[i-3] + n[i]
		if lastThreeTotal > prev {
			c++
		}
	}
	fmt.Println("Part 2:", c)
}
