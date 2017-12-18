package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func distance(x int, y int, z int) int {

	var max = int(math.Abs(float64(x)))

	if math.Abs(float64(y)) > float64(max) {
		max = y
	}

	if math.Abs(float64(z)) > float64(max) {
		max = y
	}

	return max
}

func getMax(m int, i int) int {
	if math.Abs(float64(i)) > float64(m) {
		return i
	}

	return m
}

func main() {
	data, _ := ioutil.ReadFile("input11.txt")

	var inputs = strings.Split(string(data), ",")

	var x, y, z int
	var max int

	for _, input := range inputs {
		switch d := input; d {
		case "n":
			y++
			z--
			max = getMax(max, y)
			max = getMax(max, z)
		case "ne":
			z--
			x++
			max = getMax(max, y)
			max = getMax(max, z)
		case "nw":
			x--
			y++
			max = getMax(max, x)
			max = getMax(max, y)
		case "s":
			y--
			z++
			max = getMax(max, y)
			max = getMax(max, z)
		case "se":
			y--
			x++
			max = getMax(max, x)
			max = getMax(max, y)
		case "sw":
			x--
			z++
			max = getMax(max, x)
			max = getMax(max, z)
		default:
			panic("Not supported")
		}
	}

	fmt.Printf("Current distance: %d\n", distance(x, y, z))
	fmt.Printf("Max distance is: %d\n", max)
}
