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

func jump(next int, cnt *int, nums *[]int) int {
	//if (next > len(*nums)) || (next < 0) || (*cnt > 10) {
	if next >= len(*nums) || next < 0 {
		return -1
	}

	fmt.Printf("current=%d cnt=%d \n", next, *cnt)

	(*cnt)++
	next += (*nums)[next]
	(*nums)[next]++

	fmt.Printf("jump=%d cnt=%d \n", next, *cnt)
	fmt.Printf("-------------\n")
	jump(next, cnt, nums)

	return 0
}

func main() {
	var nums, _ = readLines("input5.txt")

	cnt := 0
	jump(0, &cnt, &nums)

	fmt.Print(cnt)

}
