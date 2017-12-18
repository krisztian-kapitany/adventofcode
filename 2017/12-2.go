package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	id    int
	pipes []*node

	visited bool
}

func newNode(i int) *node {
	return &node{id: i, visited: false}
}

func getNodePtr(nodes []*node, i int) *node {
	for _, r := range nodes {
		if r.id == i {
			return r
		}
	}
	return nil
}

func addPipe(from *node, to *node) {
	from.pipes = append(from.pipes, to)
	to.pipes = append(to.pipes, from)
}

func walkFrom(start *node) {
	if start.pipes == nil {
		return
	}

	for _, pipe := range start.pipes {
		if pipe.visited == true {
			continue
		}

		pipe.visited = true
		walkFrom(pipe)
	}

	return
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	var lines, _ = readLines("input12.txt")

	var nodes []*node

	//Fill node array
	for _, line := range lines {
		var from, _ = strconv.Atoi(line[:strings.Index(line, " <")])
		nodes = append(nodes, newNode(from))
	}

	//Plumbing
	for _, line := range lines {
		var fromi, _ = strconv.Atoi(line[:strings.Index(line, " <")])
		var from = getNodePtr(nodes, fromi)

		var toStr = line[strings.Index(line, "> ")+2:]

		for _, r := range strings.Split(toStr, ", ") {
			var toi, _ = strconv.Atoi(r)
			var to = getNodePtr(nodes, toi)

			addPipe(from, to)
		}
	}

	//depth first
	var counter int
	for _, startPoint := range nodes {
		if startPoint.visited == true {
			continue
		}

		counter++
		walkFrom(startPoint)
	}

	fmt.Println(counter)

}
