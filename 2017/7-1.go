package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	name   string
	weight string

	leaves []string
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
	var lines, _ = readLines("input7.txt")

	clues := make(map[string]string)

	for _, line := range lines {
		if strings.Contains(line, " -> ") {
			clues[line[:strings.Index(line, " (")]] = line[strings.Index(line, "> ")+2:]
		}
		continue
	}

	var invalidRoots []string
	for root := range clues {
		for _, leaves := range clues {
			if strings.Contains(leaves, root) {
				//fmt.Printf("Root: %s, Leaves: %s \n", root, leaves)

				invalidRoots = append(invalidRoots, root)
			}
		}
	}

	for _, root := range invalidRoots {
		delete(clues, root)
	}

	fmt.Println("The root node is:")
	fmt.Println(clues)

}
