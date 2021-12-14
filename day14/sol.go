package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func parseRules(s string) map[string]string {
	rules := make(map[string]string)
	for _, line := range strings.Split(s, "\n") {
		str := strings.Fields(line)
		rules[str[0]] = str[2]
	}
	return rules
}

func parseArray(s string) []string {
	arr := make([]string, len(s))
	for i, char := range s {
		arr[i] = string(char)
	}
	return arr
}

func step(arr []string, rules map[string]string) []string {
	insertions := make([]string, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		pair := strings.Join([]string{arr[i], arr[i+1]}, "")
		insertions[i] = rules[pair]
	}
	merged := make([]string, 2*len(arr)-1)
	for i, j := 0, 0; i < len(arr)-1; i++ {
		merged[j] = arr[i]
		j++
		merged[j] = insertions[i]
		j++
	}
	merged[len(merged)-1] = arr[len(arr)-1]
	return merged
}

func getMostCommonMinusLeast(arr []string) int {
	m := make(map[string]int)
	for _, s := range arr {
		m[s]++
	}
	mostCommon, leastCommon := m[arr[0]], m[arr[0]]
	for _, v := range m {
		if v > mostCommon {
			mostCommon = v
		}
		if v < leastCommon {
			leastCommon = v
		}
	}
	return mostCommon - leastCommon
}

func main() {
	iPtr := flag.String("input", "input.txt", "Input filename to read the puzzle input from.")
	flag.Parse()
	dat, _ := ioutil.ReadFile("day14/" + *iPtr)
	split := strings.Split(string(dat), "\n\n")
	arr := parseArray(split[0])
	rules := parseRules(split[1])

	for i := 0; i < 10; i++ {
		arr = step(arr, rules)
	}
	fmt.Println("Part 1:", getMostCommonMinusLeast(arr))

	arr = parseArray(split[0])
	for i := 0; i < 40; i++ {
		arr = step(arr, rules)
	}
	fmt.Println("Part 2:", getMostCommonMinusLeast(arr))
}
