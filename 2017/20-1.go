package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"util"
)

func main() {
	lines, _ := util.ReadLines("input20.txt")

	minA := math.MaxFloat64
	for i, line := range lines {
		acc := line[strings.Index(line, "a=<")+3 : len(line)-1]

		var ax, ay, az float64
		var acceleration float64

		for j, part := range strings.Split(acc, ",") {
			var ai, _ = strconv.Atoi(part)

			switch j {
			case 0:
				ax = float64(ai)
			case 1:
				ay = float64(ai)
			case 2:
				az = float64(ai)
			}
		}

		acceleration = math.Sqrt((ax * ax) + (ay * ay) + (az * az))

		if acceleration < minA {
			minA = acceleration
		}
		if acceleration == minA {
			fmt.Println(i, minA)
		}
	}
}
