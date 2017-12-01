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

		if i != 0 {
			c1 = int(r - '0')
			c2 = int([]rune(str)[i-1] - '0')

		} else {
			c1 = int(r - '0')
			c2 = int([]rune(str)[len(str)-1] - '0')
		}

		if c1 == c2 {
			sum += c1
		}

	}

	fmt.Println(sum)

}
