package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

//BUFFERSIZE : size for message queues
const BUFFERSIZE = 1000000000

//QUEUETIMEOUT : time in seconds while message queues are empty
const QUEUETIMEOUT = 1

var wg sync.WaitGroup

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

type program struct {
	id          int
	registers   map[string]int
	msgQ        chan int
	sentNumbers int
}

func getVal(input string, m map[string]int) int {
	var amount, err = strconv.Atoi(input)
	if err != nil {
		if val, ok := m[input]; ok {
			amount = val
		} else {
			amount = 0
		}
	}
	return amount
}

func modifyRegister(inst instruction, exec func(int, int) int, m map[string]int) {
	m[inst.target] = exec(m[inst.target], getVal(inst.source, m))
}

func step(instructions []instruction, p *program, sendQ chan int) {
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

mainloop:
	for i := 0; i < len(instructions); {
		inst := instructions[i]

		switch command := inst.command; command {
		case "set":
			modifyRegister(inst, set, p.registers)

		case "add":
			modifyRegister(inst, add, p.registers)

		case "mul":
			modifyRegister(inst, mul, p.registers)

		case "mod":
			modifyRegister(inst, mod, p.registers)

		case "snd":
			sendQ <- getVal(inst.target, p.registers)
			p.sentNumbers++
			//fmt.Printf("(%d) SND=My new send queue len is: %d, sent(%s): %d #%d \n", p.id, len(sendQ), inst.target, getVal(inst.target, p.registers), p.sentNumbers)

		case "rcv":
			select {
			case msg := <-p.msgQ:
				p.registers[inst.target] = msg
				//fmt.Printf("(%d) RCV=My new queue len is: %d, received: %d\n", p.id, len(p.msgQ), msg)
			case <-time.After(time.Second * QUEUETIMEOUT):
				fmt.Printf("(%d) RCV=TIMEOUT, sent: %d\n", p.id, p.sentNumbers)
				break mainloop
			}

		case "jgz":
			tar := getVal(inst.target, p.registers)
			if tar > 0 {
				i += getVal(inst.source, p.registers)
				continue
			}

		default:
			fmt.Printf("(%d) Unsupported command: \"%s\" \n", p.id, command)
		}

		i++

	}
	wg.Done()
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

	var r0 = make(map[string]int)
	r0["p"] = 0
	var r1 = make(map[string]int)
	r1["p"] = 1

	var q0 = make(chan int, BUFFERSIZE)
	var q1 = make(chan int, BUFFERSIZE)

	var program0 = program{id: 0, registers: r0, msgQ: q0, sentNumbers: 0}
	var program1 = program{id: 1, registers: r1, msgQ: q1, sentNumbers: 0}

	wg.Add(2)
	go step(instructions, &program0, program1.msgQ)
	go step(instructions, &program1, program0.msgQ)
	wg.Wait()
}
