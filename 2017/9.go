package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

type myState struct {
	escaped bool
	garbage bool
	level   int
}

func main() {
	var lines, _ = readLines("input9.txt")
	var line = lines[0]

	var cnt int
	var garbageCnt int
	state := myState{false, false, 0}

	for _, r := range line {

		if state.escaped {
			state.escaped = false
			continue
		}

		switch ch := strconv.QuoteRune(r); ch {
		case "'!'":
			if state.garbage {
				state.escaped = true
			}

		case "'<'":
			if state.garbage {
				garbageCnt++
			}
			state.garbage = true

		case "'>'":
			state.garbage = false

		case "'{'":
			if !state.garbage {
				state.level++
				cnt += state.level
			} else {
				garbageCnt++
			}

		case "'}'":
			if !state.garbage {
				if state.level > 0 {
					state.level--
				}
			} else {
				garbageCnt++
			}

		default:
			if state.garbage {
				garbageCnt++
			}
			continue
		}

	}

	fmt.Println(cnt)
	fmt.Println(garbageCnt)
}
