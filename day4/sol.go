package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type board struct {
	nums   [][]int
	marked [][]bool
}

func (b board) isWin() bool {
	// check horizontal
	for i := 0; i < len(b.marked); i++ {
		horizWin := true
		for j := 0; j < len(b.marked[0]); j++ {
			if !b.marked[i][j] {
				horizWin = false
				break
			}
		}
		if horizWin {
			return true
		}
	}
	// vertical
	for i := 0; i < len(b.marked); i++ {
		vertWin := true
		for j := 0; j < len(b.marked[0]); j++ {
			if !b.marked[j][i] {
				vertWin = false
				break
			}
		}
		if vertWin {
			return true
		}
	}
	return false
}

func (b board) mark(num int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.nums[i][j] == num {
				b.marked[i][j] = true
			}
		}
	}
}

func (b board) getSumUnmarkedNums() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				sum = sum + b.nums[i][j]
			}
		}
	}
	return sum
}

func (b board) String() string {
	var sb strings.Builder
	for r, row := range b.nums {
		for c, col := range row {
			if b.marked[r][c] {
				sb.WriteString(fmt.Sprintf("*%d*,", col))
			} else {
				sb.WriteString(fmt.Sprintf("%d,", col))
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func makeBoard(s string) board {
	// s is the raw string for a board.
	rows := make([][]int, 5)
	raw := strings.Split(s, "\n")
	for i := 0; i < 5; i++ {
		rows[i] = getNumArrayFields(raw[i])
	}
	marked := make([][]bool, 5)
	for i := 0; i < 5; i++ {
		marked[i] = make([]bool, 5)
	}
	return board{nums: rows, marked: marked}
}

func getNumArrayFields(s string) []int {
	strNums := strings.Fields(s)
	nums := make([]int, len(strNums))
	for i, strNum := range strNums {
		nums[i], _ = strconv.Atoi(strNum)
	}
	return nums
}

func getNumArray(s string, sep string) []int {
	strNums := strings.Split(s, sep)
	nums := make([]int, len(strNums))
	for i, strNum := range strNums {
		nums[i], _ = strconv.Atoi(strNum)
	}
	return nums
}

func main() {
	iPtr := flag.String("input", "input.txt", "Input filename to read the puzzle input from.")
	flag.Parse()
	dat, _ := ioutil.ReadFile("day4/" + *iPtr)
	raw := strings.Split(string(dat), "\n\n")
	calledNumbers := getNumArray(raw[0], ",")
	boards := make([]board, 0)
	for i := 1; i < len(raw); i++ {
		boards = append(boards, makeBoard(raw[i]))
	}
	winningNumber := -1
	winningBoard := board{}
	for _, calledNumber := range calledNumbers {
		for _, b := range boards {
			b.mark(calledNumber)
			if b.isWin() {
				winningNumber = calledNumber
				winningBoard = b
				fmt.Println("Winner after calling:", winningNumber)
				fmt.Println("Winning board:", winningBoard)
				break
			}
		}
		if winningNumber > -1 {
			break
		}
	}
	fmt.Println("Part 1:", winningBoard.getSumUnmarkedNums()*winningNumber)
}
