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

type instruction struct {
	command string
	target  string
	source  string
}

func modifyRegister(inst instruction, exec func(int, int) int, m *map[string]int) {
	var amount, err = strconv.Atoi(inst.source)
	if err != nil {
		if val, ok := (*m)[inst.source]; ok {
			amount = val
		} else {
			amount = 0
		}
	}

	(*m)[inst.target] = exec((*m)[inst.target], amount)
}

func main() {
	var lines, _ = readLines("input18.txt")

	var instructions []instruction

	for _, line := range lines {
		var parts = strings.Fields(line)

		var inst = instruction{command: parts[0], target: parts[1]}

		if len(parts) > 2 {
			inst.source = parts[2]
		}
		instructions = append(instructions, inst)
	}

	var registers = make(map[string]int)

	set := func(a int, b int) int {
		_ = a
		return b
	}

	add := func(a int, b int) int {
		return a + b
	}

	mul := func(a int, b int) int {
		return a * b
	}

	mod := func(a int, b int) int {
		return a % b
	}

	var lastsound int
	var finished = false
	for i := 0; ; {
		if finished {
			break
		}

		inst := instructions[i]

		switch command := inst.command; command {
		case "set":
			modifyRegister(inst, set, &registers)

		case "add":
			modifyRegister(inst, add, &registers)

		case "mul":
			modifyRegister(inst, mul, &registers)

		case "mod":
			modifyRegister(inst, mod, &registers)

		case "snd":
			if registers[inst.target] != 0 {
				lastsound = registers[inst.target]
			}

		case "rcv":
			if registers[inst.target] != 0 {
				finished = true
				fmt.Println(lastsound)
				break
			}

		case "jgz":
			if registers[inst.target] > 0 {
				var amount, err = strconv.Atoi(inst.source)
				if err != nil {
					if val, ok := registers[inst.source]; ok {
						amount = val
					} else {
						amount = 0
					}
				}

				i += amount
				continue
			}

		default:
			fmt.Printf("Unsupported command: %s. \n", command)
		}

		i++

	}

}
