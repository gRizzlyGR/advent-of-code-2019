package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatalf("Cannot open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var program []int

	for scanner.Scan() {
		program = convertStringToInts(scanner.Text())
	}

	exec(program, 1)
}

func convertStringToInts(str string) []int {
	var ints []int
	split := strings.Split(str, ",")
	for _, s := range split {
		n, _ := strconv.Atoi(s)
		ints = append(ints, n)
	}

	return ints
}

func exec(program []int, input int) {
	steps := 0
	for i := 0; i < len(program); i += steps {
		opcode := program[i]

		switch opcode % 100 {
		case 1:
			param1 := getParam(program, opcode, i, 1)
			param2 := getParam(program, opcode, i, 2)
			resPos := program[i+3]
			program[resPos] = param1 + param2
			steps = 4
		case 2:
			param1 := getParam(program, opcode, i, 1)
			param2 := getParam(program, opcode, i, 2)
			resPos := program[i+3]
			program[resPos] = param1 * param2
			steps = 4
		case 3:
			paramPos := program[i+1]
			program[paramPos] = input
			steps = 2
		case 4:
			param := getParam(program, opcode, i, 1)
			fmt.Println(param)
			steps = 2
		case 99:
			return
		default:
			log.Fatalf("Unsupported opcode: %d\n", opcode)
		}
	}
}

func getParam(program []int, opcode, instrPointer, offset int) int {
	param := 0
	if isImmediate(opcode, offset) {
		param = program[instrPointer+offset]
	} else {
		param = program[program[instrPointer+offset]]
	}

	return param
}

// Pos == 1 --> hundreds
// Pos == 2 --> thousands
// Pos == 3 --> tens of thousands
func isImmediate(opcode, pos int) bool {
	// Get the unit
	return opcode/int(math.Pow10(pos+1))%10 == 1
}
