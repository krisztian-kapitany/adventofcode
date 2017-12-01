package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func nextNum(n int, l int) int {
	var r = (n + (l / 2)) % l

	return r
}

func main() {
	data, err := ioutil.ReadFile("input1-1.txt")
	check(err)

	var str = string(data)
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")

	var sum int

	for i, r := range str {
		var c1 int
		var c2 int

		c1 = int(r - '0')
		c2 = int([]rune(str)[nextNum(i, len(str))] - '0')

		if c1 == c2 {
			sum += c1
		}

	}

	fmt.Println(sum)

}
