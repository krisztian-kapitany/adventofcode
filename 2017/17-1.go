package main

import (
	"fmt"
)

func step(c []int, current int, steps int) int {
	return (current + steps) % len(c)
}

func main() {

	const input = 348

	var cycle []int
	cycle = append(cycle, 0)

	current := 0
	for i := 1; i < 2018; i++ {
		current = step(cycle, current, input)

		tail := make([]int, len(cycle)-current-1)
		copy(tail, cycle[current+1:])

		current++

		cycle = append(cycle[:current], i)
		cycle = append(cycle, tail...)
	}

	fmt.Println(cycle[current+1])
}
