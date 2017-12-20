package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	mode  int
	size  int
	posA  int
	posB  int
	runeA rune
	runeB rune
}

func getIndexOf(arr []rune, a rune) int {
	for i, r := range arr {
		if r == a {
			return i
		}
	}
	return -1
}

func main() {
	inputStr := "abcdefghijklmnop"
	order := []rune(inputStr)

	data, _ := ioutil.ReadFile("input16.txt")
	var moves = string(data)

	var instructions []instruction
	for _, move := range strings.Split(moves, ",") {

		if strings.HasPrefix(move, "s") {
			size, _ := strconv.Atoi(move[1:])
			size = len(order) - size

			instructions = append(instructions, instruction{mode: 0, size: size})
			continue
		}

		if strings.HasPrefix(move, "x") {
			perIndex := strings.Index(move, "/")

			posA, _ := strconv.Atoi(move[1:perIndex])
			posB, _ := strconv.Atoi(move[perIndex+1:])

			instructions = append(instructions, instruction{mode: 1, posA: posA, posB: posB})
			continue
		}

		if strings.HasPrefix(move, "p") {
			runeA := rune(move[1])
			runeB := rune(move[3])

			instructions = append(instructions, instruction{mode: 2, runeA: runeA, runeB: runeB})
			continue
		}

	}

	var states = make(map[string]int)
	var cycleLen int

	for dancei := 0; dancei < 150; dancei++ {
		for _, inst := range instructions {

			if inst.mode == 0 {
				order = append(order[inst.size:], order[:inst.size]...)
				continue
			}

			if inst.mode == 1 {
				order[inst.posA], order[inst.posB] = order[inst.posB], order[inst.posA]
				continue
			}

			if inst.mode == 2 {
				posA := getIndexOf(order, inst.runeA)
				posB := getIndexOf(order, inst.runeB)

				order[posA], order[posB] = order[posB], order[posA]
				continue
			}

		}

		var orderi string
		for _, c := range order {
			orderi += fmt.Sprintf("%c", c)
		}

		fmt.Println(orderi, dancei)

		if val, ok := states[orderi]; ok {
			cycleLen = dancei - val

			var leftover = (1000000000 - 1) % cycleLen

			fmt.Println(cycleLen)
			fmt.Println(leftover)

			for k, v := range states {
				if v == leftover {
					fmt.Println(k)
				}
			}
			break

		} else {
			states[orderi] = dancei
		}
	}
}
