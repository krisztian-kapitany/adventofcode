package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = "0	5	10	0	11	14	13	4	11	8	8	7	1	4	12	11"

func getbiggest(array []int) (int, int) {
	var max = 0
	var index = 0

	for i, r := range array {
		if r > max {
			max = r
			index = i
		}
	}

	return index, max
}

func redist(a *[]int, cnt int, startFrom int) {

	(*a)[startFrom] = 0

	for i := 1; i < cnt; i++ {
		k := (startFrom + i) % len(*a)
		(*a)[k]++
	}
}

func main() {
	var bins []int
	var cnt = 0

	for _, r := range strings.Fields(input) {
		ri, _ := strconv.Atoi(r)
		bins = append(bins, ri)
	}

	index, biggest := getbiggest(bins)
	fmt.Println(bins)
	redist(&bins, biggest, index)
	index, biggest = getbiggest(bins)
	fmt.Println(bins)

	fmt.Println(cnt)
	fmt.Println(index)
	fmt.Println(biggest)

}
