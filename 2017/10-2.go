package main

import (
	"fmt"
)

func main() {
	var inputStr = "46,41,212,83,1,255,157,65,139,52,39,254,2,86,0,204"
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
			//reverse cut
			for i := 0; i < input/2; i++ {
				var tmp = cycle[(position+i)%cycleLen]
				cycle[(position+i)%cycleLen] = cycle[(position+input-1-i)%cycleLen]
				cycle[(position+input-1-i)%cycleLen] = tmp
			}

			//move position
			position += skipsize + input
			position = position % cycleLen

			//printC(cycle, position)
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

	for _, r := range denseHash {
		fmt.Printf("%02x", r)
	}

}
