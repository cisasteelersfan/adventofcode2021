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

func parseMap(s string, rules map[string]string) map[string]int {
	m := make(map[string]int)
	for i := 1; i < len(s); i++ {
		m[s[i-1:i+1]]++
	}
	return m
}

func step2(m map[string]int, rules map[string]string) map[string]int {
	nm := make(map[string]int)
	for k, v := range m {
		first := string(k[0]) + rules[k]
		second := rules[k] + string(k[1])
		nm[first] = nm[first] + v
		nm[second] = nm[second] + v
	}
	return nm
}

func getPart2(m map[string]int, start, end string) int {
	letterMap := make(map[string]int)
	for k, v := range m {
		letterMap[string(k[0])] = letterMap[string(k[0])] + v
		letterMap[string(k[1])] = letterMap[string(k[1])] + v
	}
	letterMap[start]++
	letterMap[end]++
	for k, v := range letterMap {
		letterMap[k] = v / 2
	}
	mostCommon := 0
	leastCommon := letterMap["N"]
	for _, v := range letterMap {
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

	m := parseMap(split[0], rules)
	for i := 0; i < 40; i++ {
		m = step2(m, rules)
	}
	fmt.Println("Part 2:", getPart2(m, string(split[0][0]), string(split[0][len(split[0])-1])))
}
