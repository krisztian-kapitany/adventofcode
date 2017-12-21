package main

import (
	"fmt"
)

func main() {

	const input = 348

	current := 0
	result := 0
	for i := 1; i < 50000001; i++ {
		current = ((current + input) % i) + 1

		if current == 1 {
			result = i
		}
	}

	fmt.Println(result)
}
