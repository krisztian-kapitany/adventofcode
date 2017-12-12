package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func eval(a string, test func(int, string) bool, b string, m map[string]int) bool {
	if v, ok := m[a]; ok {
		return test(v, b)
	}
	return test(0, b)
}

func execCommand(target string, command string, amount int, m *map[string]int) {
	if command == "dec" {
		amount *= -1
	}

	(*m)[target] += amount
}

func main() {
	var lines, _ = readLines("input8.txt")

	gt := func(a int, b string) bool {
		var bi, _ = strconv.Atoi(b)
		return (a > bi)
	}
	lt := func(a int, b string) bool {
		var bi, _ = strconv.Atoi(b)
		return (a < bi)
	}
	ge := func(a int, b string) bool {
		var bi, _ = strconv.Atoi(b)
		return (a >= bi)
	}
	le := func(a int, b string) bool {
		var bi, _ = strconv.Atoi(b)
		return (a <= bi)
	}
	eq := func(a int, b string) bool {
		var bi, _ = strconv.Atoi(b)
		return (a == bi)
	}
	ne := func(a int, b string) bool {
		var bi, _ = strconv.Atoi(b)
		return (a != bi)
	}

	var valueRegisters = make(map[string]int)

	for _, line := range lines {
		var parts = strings.Fields(line)

		var targetRegister = parts[0]
		var command = parts[1]
		var amount, _ = strconv.Atoi(parts[2])
		var conditionL = parts[4]
		var operator = parts[5]
		var conditionR = parts[6]

		switch op := operator; op {
		case ">":
			if eval(conditionL, gt, conditionR, valueRegisters) {
				execCommand(targetRegister, command, amount, &valueRegisters)
			}
		case "<":
			if eval(conditionL, lt, conditionR, valueRegisters) {
				execCommand(targetRegister, command, amount, &valueRegisters)
			}
		case ">=":
			if eval(conditionL, ge, conditionR, valueRegisters) {
				execCommand(targetRegister, command, amount, &valueRegisters)
			}
		case "<=":
			if eval(conditionL, le, conditionR, valueRegisters) {
				execCommand(targetRegister, command, amount, &valueRegisters)
			}
		case "==":
			if eval(conditionL, eq, conditionR, valueRegisters) {
				execCommand(targetRegister, command, amount, &valueRegisters)
			}
		case "!=":
			if eval(conditionL, ne, conditionR, valueRegisters) {
				execCommand(targetRegister, command, amount, &valueRegisters)
			}
		default:
			fmt.Printf("Unsupported condition: %s. \n", op)
		}

	}

	var max = 0

	for _, v := range valueRegisters {
		if v > max {
			max = v
		}
	}

	fmt.Println(max)

}
