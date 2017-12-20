package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

type scanner struct {
	layer  int
	srange int
}

func main() {
	var lines, _ = readLines("input13.txt")
	var inputs []scanner

	for _, line := range lines {
		var parts = strings.Split(line, ": ")

		var l, _ = strconv.Atoi(parts[0])
		var sr, _ = strconv.Atoi(parts[1])

		inputs = append(inputs, scanner{l, sr})
	}

	for delay := 0; ; delay++ {
		var boom = false
		for _, input := range inputs {
			if (input.layer+delay)%((2*input.srange)-2) == 0 {
				boom = true
				break
			}
		}

		if !boom {
			fmt.Println(delay)
			break
		}
	}
}
