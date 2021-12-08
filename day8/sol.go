package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

func parseOutput(s string) [][]string {
	arr := make([][]string, 0)
	for _, line := range strings.Split(s, "\n") {
		split := strings.Split(line, "|")
		arr = append(arr, strings.Fields(split[1]))
	}
	return arr
}

func countUniqueEntries(s [][]string) int {
	ans := 0
	for _, line := range s {
		for _, entry := range line {
			if len(entry) == 2 || len(entry) == 3 || len(entry) == 4 || len(entry) == 7 {
				ans++
			}
		}
	}
	return ans
}

func makeArrayOfSets(s []string) []map[string]bool {
	ans := make([]map[string]bool, 0)
	for i, word := range s {
		ans = append(ans, make(map[string]bool))
		for _, letter := range word {
			ans[i][string(letter)] = true
		}
	}
	return ans
}

func parseLine(s string) int {
	// num segments -> possible numbers
	// 2: 1
	// 3: 7
	// 4: 4
	// 5: 2,3,5
	// 6: 0,6,9
	// 7: 8
	split := strings.Split(s, "|")
	segments := makeArrayOfSets(strings.Fields(split[0]))
	output := makeArrayOfSets(strings.Fields(split[1]))
	numSegmentsToPossibleSet := make(map[int][]map[string]bool)
	for _, set := range segments {
		numSegmentsToPossibleSet[len(set)] = append(numSegmentsToPossibleSet[len(set)], set)
	}

	numToSegmentMap := make(map[int]map[string]bool)
	numToSegmentMap[1] = numSegmentsToPossibleSet[2][0]
	numToSegmentMap[4] = numSegmentsToPossibleSet[4][0]
	numToSegmentMap[7] = numSegmentsToPossibleSet[3][0]
	numToSegmentMap[8] = numSegmentsToPossibleSet[7][0]
	fourdiff := make(map[string]bool)
	for key := range numToSegmentMap[4] {
		if !numToSegmentMap[1][key] {
			fourdiff[key] = true
		}
	}

	for _, set := range numSegmentsToPossibleSet[5] {
		if contains(set, numToSegmentMap[1]) {
			numToSegmentMap[3] = set
		} else if contains(set, fourdiff) {
			numToSegmentMap[5] = set
		} else {
			numToSegmentMap[2] = set
		}
	}
	for _, set := range numSegmentsToPossibleSet[6] {
		if contains(set, numToSegmentMap[4]) {
			numToSegmentMap[9] = set
		} else if contains(set, fourdiff) {
			numToSegmentMap[6] = set
		} else {
			numToSegmentMap[0] = set
		}
	}

	sum := make([]string, 4)
	for i, seg := range output {
		for num, m := range numToSegmentMap {
			if reflect.DeepEqual(m, seg) {
				sum[i] = strconv.Itoa(num)
				break
			}
		}
	}
	ans, _ := strconv.Atoi(sum[0] + sum[1] + sum[2] + sum[3])
	return ans
}

func contains(bigger, smaller map[string]bool) bool {
	for key := range smaller {
		if !bigger[key] {
			return false
		}
	}
	return true
}

func getSegmentA(one, seven []map[string]bool) string {
	oneSegments := one[0]
	sevenSegments := seven[0]
	for k := range sevenSegments {
		if !oneSegments[k] {
			return k
		}
	}
	panic("Couldn't find A")
}

func getNummberWithSegmentMapping(s map[string]string) int {
	segments := make(map[string]bool)
	for _, v := range s {
		segments[v] = true
	}
	segments_0 := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
		"e": true,
		"f": true,
		"g": true,
	}
	segments_1 := map[string]bool{
		"c": true,
		"f": true,
	}
	segments_2 := map[string]bool{
		"a": true,
		"c": true,
		"d": true,
		"e": true,
		"g": true,
	}
	segments_3 := map[string]bool{
		"a": true,
		"c": true,
		"d": true,
		"f": true,
		"g": true,
	}
	segments_4 := map[string]bool{
		"b": true,
		"c": true,
		"d": true,
		"f": true,
	}
	segments_5 := map[string]bool{
		"a": true,
		"b": true,
		"d": true,
		"f": true,
		"g": true,
	}
	segments_6 := map[string]bool{
		"a": true,
		"b": true,
		"d": true,
		"e": true,
		"f": true,
		"g": true,
	}
	segments_7 := map[string]bool{
		"a": true,
		"c": true,
		"f": true,
	}
	segments_8 := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
		"d": true,
		"e": true,
		"f": true,
		"g": true,
	}
	segments_9 := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
		"d": true,
		"f": true,
		"g": true,
	}
	if reflect.DeepEqual(segments, segments_0) {
		return 0
	}
	if reflect.DeepEqual(segments, segments_1) {
		return 1
	}
	if reflect.DeepEqual(segments, segments_2) {
		return 2
	}
	if reflect.DeepEqual(segments, segments_3) {
		return 3
	}
	if reflect.DeepEqual(segments, segments_4) {
		return 4
	}
	if reflect.DeepEqual(segments, segments_5) {
		return 5
	}
	if reflect.DeepEqual(segments, segments_6) {
		return 6
	}
	if reflect.DeepEqual(segments, segments_7) {
		return 7
	}
	if reflect.DeepEqual(segments, segments_8) {
		return 8
	}
	if reflect.DeepEqual(segments, segments_9) {
		return 9
	}
	panic("couldn't lookup number")
}

func main() {
	iPtr := flag.String("input", "input.txt", "Input filename to read the puzzle input from.")
	flag.Parse()
	dat, _ := ioutil.ReadFile("day8/" + *iPtr)
	output := parseOutput(string(dat))
	fmt.Println("Part 1:", countUniqueEntries(output))

	part2 := 0
	for _, line := range strings.Split(string(dat), "\n") {
		part2 = part2 + parseLine(line)
	}
	fmt.Println("Part 2:", part2)
}
