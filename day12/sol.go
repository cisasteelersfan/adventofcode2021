package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

// adjacency lists
type graph map[string][]string

func parseInput(s string) graph {
	g := graph{}
	for _, line := range strings.Split(s, "\n") {
		split := strings.Split(line, "-")
		u, v := split[0], split[1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	return g
}

type container struct {
	paths []string
}

func (c *container) addPath(path string) {
	c.paths = append(c.paths, path)
}

func findAllPaths(g graph) []string {
	ans := container{}
	smallVisitedCaves := make(map[string]bool)
	smallVisitedCaves["start"] = true
	currentPath := []string{"start"}
	backtrack(g, currentPath, smallVisitedCaves, &ans)
	return ans.paths
}

func backtrack(g graph, currentPath []string, smallVisitedCaves map[string]bool, ans *container) {
	currentNode := currentPath[len(currentPath)-1]
	if currentNode == "end" {
		ans.addPath(strings.Join(currentPath, ","))
		return
	}
	for _, neighbor := range g[currentNode] {
		if smallVisitedCaves[neighbor] {
			continue
		}
		if IsLower(neighbor) {
			smallVisitedCaves[neighbor] = true
		}
		currentPath = append(currentPath, neighbor)
		backtrack(g, currentPath, smallVisitedCaves, ans)
		currentPath = currentPath[:len(currentPath)-1]
		if IsLower(neighbor) {
			smallVisitedCaves[neighbor] = false
		}
	}
}

func findAllPaths2(g graph) []string {
	ans := container{}
	smallVisitedCaves := make(map[string]int)
	smallVisitedCaves["start"]++
	currentPath := []string{"start"}
	backtrack2(g, currentPath, smallVisitedCaves, &ans)
	return ans.paths
}

func backtrack2(g graph, currentPath []string, smallVisitedCaves map[string]int, ans *container) {
	currentNode := currentPath[len(currentPath)-1]
	if currentNode == "end" {
		ans.addPath(strings.Join(currentPath, ","))
		return
	}
	for _, neighbor := range g[currentNode] {
		if neighbor == "start" {
			continue
		}
		if IsLower(neighbor) {
			if smallVisitedCaves[neighbor] == 0 || noTwos(smallVisitedCaves) {
				smallVisitedCaves[neighbor]++
			} else {
				continue
			}
		}
		currentPath = append(currentPath, neighbor)
		backtrack2(g, currentPath, smallVisitedCaves, ans)
		currentPath = currentPath[:len(currentPath)-1]
		if IsLower(neighbor) {
			smallVisitedCaves[neighbor]--
		}
	}
}

func noTwos(m map[string]int) bool {
	for _, v := range m {
		if v == 2 {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func main() {
	iPtr := flag.String("input", "input.txt", "Input filename to read the puzzle input from.")
	flag.Parse()
	dat, _ := ioutil.ReadFile("day12/" + *iPtr)
	g := parseInput(string(dat))
	fmt.Println("Part 1:", len(findAllPaths(g)))
	fmt.Println("Part 2:", len(findAllPaths2(g)))
}
