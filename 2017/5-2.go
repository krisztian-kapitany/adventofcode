package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var i, _ = strconv.Atoi(scanner.Text())
		lines = append(lines, i)
	}
	return lines, scanner.Err()
}

func main() {
	var nums, _ = readLines("input5.txt")

	cnt := 0
	current := 0
	next := 0

	for {
		if next >= len(nums) || next < 0 {
			break
		}
		cnt++

		current = next
		next += nums[current]

		if nums[current] < 3 {
			nums[current]++
		} else {
			nums[current]--
		}
	}

	fmt.Print(cnt)

}
