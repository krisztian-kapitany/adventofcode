package main

import (
	"fmt"
	"math"
)

func main() {

	var in float64 = 312051

	var root = math.Sqrt(in)
	var k = int(root)
	var largestbaseonlayer int

	if float64(k) == root && k%2 != 0 {
		largestbaseonlayer = k
	} else if k%2 != 0 {
		largestbaseonlayer = k + 2
	} else {
		largestbaseonlayer = k + 1
	}

	var layernr = (largestbaseonlayer - 1) / 2
	var offset = in - math.Pow(float64(largestbaseonlayer-2), 2)
	var layerdistance = layernr - (int(offset) % layernr)

	fmt.Printf("Root: %f \n", root)
	fmt.Printf("Largest base on layer: %d \n", largestbaseonlayer)
	fmt.Printf("Offset from smaller base: %.0f \n", offset)
	fmt.Printf("Nth layer: %d \n", layernr)
	fmt.Printf("Layer distance: %d \n", layerdistance)

	var md = layernr + layerdistance

	fmt.Printf("Sum: %d \n", md)

}
