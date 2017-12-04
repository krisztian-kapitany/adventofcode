package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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

func sContains(s []string, e string) bool {
	for _, si := range s {

		if si == e {
			fmt.Printf("%s == %s \n", si, e)
			return true

		}
	}
	return false
}

func main() {
	lines, _ := readLines("input4.txt")

	var sum = 0

	for _, line := range lines {
		var words = strings.Fields(line)

		for i, word := range words {
			if i == len(words)-1 {
				sum++
				break
			}
			if sContains(words[i+1:], string(word)) {
				break
			}
		}
	}

	fmt.Printf("There are %d valid lines", sum)

}
