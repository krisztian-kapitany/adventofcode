package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

const (
	up    = iota
	right = iota
	down  = iota
	left  = iota
)

func step(l int, i int, dir int) (int, int) {
	switch dir {
	case up:
		return l - 1, i

	case right:
		return l, i + 1

	case down:
		return l + 1, i

	case left:
		return l, i - 1

	default:
		fmt.Println("No such direction", dir)
	}

	fmt.Println("Impossibru!")
	return -1, -1
}

func isLetter(s string) bool {
	rx, _ := regexp.Compile("[A-Z]")
	return rx.MatchString(s)
}

func isValidDirection(r byte, d int) bool {
	char := fmt.Sprintf("%c", r)

	if char == "|" && (d == right || d == left) {
		return true
	}
	if char == "-" && (d == up || d == down) {
		return true
	}

	return false
}

func main() {
	lines, _ := readLines("input19.txt")

	dir := down
	lineIndex := 0
	runeIndex := strings.Index(lines[lineIndex], "|")

	steps := 0

mainloop:
	for {
		steps++

		char := fmt.Sprintf("%c", lines[lineIndex][runeIndex])

		if char == "+" {
			for i := 1; i < 4; i++ {
				li, ri := step(lineIndex, runeIndex, (dir+i)%4)
				if isValidDirection(lines[li][ri], dir) {
					lineIndex, runeIndex = li, ri
					dir = (dir + i) % 4
					continue mainloop
				}
			}
			break mainloop

		}

		if isLetter(char) {
			fmt.Print(char)
			li, ri := step(lineIndex, runeIndex, dir)
			if isValidDirection(lines[li][ri], (dir+1)%4) {
				lineIndex, runeIndex = step(lineIndex, runeIndex, dir)
				continue mainloop
			}
			break mainloop

		}

		lineIndex, runeIndex = step(lineIndex, runeIndex, dir)
	}

	fmt.Println()
	fmt.Printf("Steps taken: %d", steps)
}
