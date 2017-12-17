package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputStr = "46,41,212,83,1,255,157,65,139,52,39,254,2,86,0,204"
	const cycleLen = 256

	var inputs []int
	for _, r := range strings.Split(inputStr, ",") {
		var i, _ = strconv.Atoi(r)
		inputs = append(inputs, i)
	}

	var cycle [cycleLen]int
	for i := 0; i < cycleLen; i++ {
		cycle[i] = i
	}

	var skipsize = 0
	var position = 0

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

		skipsize++
	}

	fmt.Printf("%d \n", cycle[0]*cycle[1])

}
