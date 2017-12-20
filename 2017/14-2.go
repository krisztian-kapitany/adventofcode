package main

import (
	"fmt"
	"strconv"
)

func getHash(inputStr string) [16]int {
	const cycleLen = 256

	var inputs []int
	for _, r := range inputStr {
		var i = int(r)
		inputs = append(inputs, i)
	}
	inputs = append(inputs, []int{17, 31, 73, 47, 23}...)

	var cycle [cycleLen]int
	for i := 0; i < cycleLen; i++ {
		cycle[i] = i
	}

	var skipsize = 0
	var position = 0
	for j := 0; j < 64; j++ {
		for _, input := range inputs {
			for i := 0; i < input/2; i++ {
				var tmp = cycle[(position+i)%cycleLen]
				cycle[(position+i)%cycleLen] = cycle[(position+input-1-i)%cycleLen]
				cycle[(position+input-1-i)%cycleLen] = tmp
			}
			position += skipsize + input
			position = position % cycleLen

			skipsize++
		}
	}

	var denseHash [16]int

	for k := 0; k < 16; k++ {
		var xor = cycle[(k*16)] ^ cycle[(k*16)+1]
		for l := 2; l < 16; l++ {
			xor = xor ^ cycle[(k*16)+l]
		}
		denseHash[k] = xor
	}

	return denseHash

}

func zeroOutBlocks(memMatrix *[128][128]int, j int, k int) {
	if (*memMatrix)[j][k] == 0 {
		return
	}

	(*memMatrix)[j][k] = 0
	if j < 127 {
		zeroOutBlocks(memMatrix, j+1, k)
	}
	if j > 0 {
		zeroOutBlocks(memMatrix, j-1, k)
	}
	if k < 127 {
		zeroOutBlocks(memMatrix, j, k+1)
	}
	if k > 0 {
		zeroOutBlocks(memMatrix, j, k-1)
	}

}

func main() {

	var input = "hwlqcszp"

	var memMatrix [128][128]int

	//var cnt int
	for i := 0; i < 128; i++ {
		var inputN = input + "-" + strconv.Itoa(i)

		var s string
		for _, r := range getHash(inputN) {
			s += fmt.Sprintf("%08b", r)
		}

		for j, r := range s {
			if r == 49 {
				memMatrix[i][j] = 1
			}
		}
	}

	var blockCnt int
	for i := 0; i < 16384; i++ {
		var j = i % 128
		var k = i / 128

		if memMatrix[j][k] == 1 {
			blockCnt++

			zeroOutBlocks(&memMatrix, j, k)

		}

	}

	fmt.Println(blockCnt)

}
