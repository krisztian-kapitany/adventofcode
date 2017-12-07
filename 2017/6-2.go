package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = "0	5	10	0	11	14	13	4	11	8	8	7	1	4	12	11"

//var input = "0	2	7	0"

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

	for i := 1; i <= cnt; i++ {
		k := (startFrom + i) % len(*a)
		(*a)[k]++
	}
}

func gethash(a []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), ","), "[]")
}

func containsS(a []string, s string) bool {
	for _, r := range a {
		if r == s {
			return true
		}
	}
	return false
}

func main() {
	var bins []int
	var cnt = 0

	states := make(map[string]int)

	for _, r := range strings.Fields(input) {
		ri, _ := strconv.Atoi(r)
		bins = append(bins, ri)
	}

	for {
		cnt++

		if val, ok := states[gethash(bins)]; ok {
			cnt -= val
			break
		}

		states[gethash(bins)] = cnt
		index, biggest := getbiggest(bins)
		redist(&bins, biggest, index)

	}

	fmt.Println("====================")
	fmt.Println(cnt)

}
