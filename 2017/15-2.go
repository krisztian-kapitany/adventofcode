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

	var accA []int
	var accB []int

	for {

		prevA = calc(prevA, factorA)
		prevB = calc(prevB, factorB)

		if prevA%4 == 0 {
			accA = append(accA, prevA)
		}
		if prevB%8 == 0 {
			accB = append(accB, prevB)
		}

		if len(accA) >= 5000000 && len(accB) >= 5000000 {
			break
		}
	}

	var cnt int
	for i, r := range accA {
		//2^16=65536
		if i > len(accB)-1 {
			break
		}

		if (r % 65536) == (accB[i] % 65536) {
			cnt++
		}
	}

	fmt.Println(cnt)
}
