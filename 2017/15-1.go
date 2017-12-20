package main

import (
	"fmt"
)

func calc(a int, f int) int {
	return (a * f) % 2147483647
}

func main() {
	//Generator A starts with 873
	//Generator B starts with 583
	var startA = 873
	var startB = 583

	//generator A uses 16807
	//generator B uses 48271
	var factorA = 16807
	var factorB = 48271

	var prevA = startA
	var prevB = startB

	var cnt int
	for i := 0; i < 40000000; i++ {

		prevA = calc(prevA, factorA)
		prevB = calc(prevB, factorB)

		//2^16=65536
		if (prevA % 65536) == (prevB % 65536) {
			cnt++
		}
	}

	fmt.Println(cnt)
}
