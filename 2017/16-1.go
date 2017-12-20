package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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

	for _, move := range strings.Split(moves, ",") {

		if strings.HasPrefix(move, "s") {
			size, _ := strconv.Atoi(move[1:])
			size = len(order) - size

			order = append(order[size:], order[:size]...)
			continue
		}

		if strings.HasPrefix(move, "x") {
			perIndex := strings.Index(move, "/")

			posA, _ := strconv.Atoi(move[1:perIndex])
			posB, _ := strconv.Atoi(move[perIndex+1:])

			order[posA], order[posB] = order[posB], order[posA]
			continue
		}

		if strings.HasPrefix(move, "p") {
			runeA := rune(move[1])
			runeB := rune(move[3])

			posA := getIndexOf(order, runeA)
			posB := getIndexOf(order, runeB)

			order[posA], order[posB] = order[posB], order[posA]
			continue
		}

	}

	for _, c := range order {
		fmt.Printf("%c", c)
	}

}
