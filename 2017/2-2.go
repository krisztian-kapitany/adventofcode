package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("input2.txt")
	check(err)

	var sum int

	for _, line := range lines {
		var nums []int

		for _, r := range strings.Fields(line) {
			n, err := strconv.Atoi(r)
			if err != nil {
				panic(err)
			}

			nums = append(nums, n)
		}

		sort.Ints(nums)

		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[j]%nums[i] == 0 {
					sum += nums[j] / nums[i]
					break
				}

			}
		}

	}

	fmt.Println(sum)
}
