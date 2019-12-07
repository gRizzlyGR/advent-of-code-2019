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
	scanner := bufio.NewScanner(os.Stdin)

	var program []int

	for scanner.Scan() {
		program = convertStringToInts(scanner.Text())
	}

	exec(program, 5)
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
	ip := 0
	for {
		opcode := program[ip]

		switch opcode % 100 {
		case 1:
			param1 := getParam(program, opcode, ip, 1)
			param2 := getParam(program, opcode, ip, 2)
			resPos := program[ip+3]
			program[resPos] = param1 + param2
			ip += 4
		case 2:
			param1 := getParam(program, opcode, ip, 1)
			param2 := getParam(program, opcode, ip, 2)
			resPos := program[ip+3]
			program[resPos] = param1 * param2
			ip += 4
		case 3:
			paramPos := program[ip+1]
			program[paramPos] = input
			ip += 2
		case 4:
			param := getParam(program, opcode, ip, 1)
			fmt.Println(param)
			ip += 2
		// Opcode 5 is jump-if-true: if the first parameter
		// is non-zero, it sets the instruction pointer to
		// the value from the second parameter. Otherwise,
		// it does nothing.
		case 5:
			param1 := getParam(program, opcode, ip, 1)
			param2 := getParam(program, opcode, ip, 2)
			if param1 != 0 {
				ip = param2
			} else {
				ip += 3
			}
		// Opcode 6 is jump-if-false: if the first parameter
		// is zero, it sets the instruction pointer to the
		// value from the second parameter. Otherwise, it does
		// nothing.
		case 6:
			param1 := getParam(program, opcode, ip, 1)
			param2 := getParam(program, opcode, ip, 2)
			if param1 == 0 {
				ip = param2
			} else {
				ip += 3
			}
		// Opcode 7 is less than: if the first parameter is
		// less than the second parameter, it stores 1 in the
		// position given by the third parameter. Otherwise,
		// it stores 0.
		case 7:
			param1 := getParam(program, opcode, ip, 1)
			param2 := getParam(program, opcode, ip, 2)
			param3 := program[ip+3]
			if param1 < param2 {
				program[param3] = 1
			} else {
				program[param3] = 0
			}
			ip += 4
		// Opcode 8 is equals: if the first parameter is equal
		// to the second parameter, it stores 1 in the position
		// given by the third parameter. Otherwise, it stores 0.
		case 8:
			param1 := getParam(program, opcode, ip, 1)
			param2 := getParam(program, opcode, ip, 2)
			param3 := program[ip+3]
			if param1 == param2 {
				program[param3] = 1
			} else {
				program[param3] = 0
			}
			ip += 4
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
