package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

type stack []string

func (s *stack) Push(v string) {
	*s = append(*s, v)
}

func (s *stack) Pop() string {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *stack) Peek() string {
	if len(*s) == 0 {
		return ""
	}
	return (*s)[len(*s)-1]
}

func parseLine(l string) int {
	s := stack{}
	for _, char := range l {
		str := string(char)
		switch str {
		case "[":
			fallthrough
		case "(":
			fallthrough
		case "{":
			fallthrough
		case "<":
			s.Push(str)
			break
		case ")":
			if s.Peek() == "(" {
				s.Pop()
				break
			} else {
				return getPoints(str)
			}
		case "]":
			if s.Peek() == "[" {
				s.Pop()
				break
			} else {
				return getPoints(str)
			}
		case "}":
			if s.Peek() == "{" {
				s.Pop()
				break
			} else {
				return getPoints(str)
			}
		case ">":
			if s.Peek() == "<" {
				s.Pop()
			} else {
				return getPoints(str)
			}
			break
		}
	}
	fmt.Println("No illegal lines")
	return 0
}

func getPoints(s string) int {
	switch s {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}
	panic("wrong char")
}

func main() {
	iPtr := flag.String("input", "input.txt", "Input filename to read the puzzle input from.")
	flag.Parse()
	dat, _ := ioutil.ReadFile("day10/" + *iPtr)
	split := strings.Split(string(dat), "\n")
	totalScore := 0
	for _, line := range split {
		totalScore = totalScore + parseLine(line)
	}
	fmt.Println("Part 1:", totalScore)
}
