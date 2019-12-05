package main

import (
	"bufio"
	"fmt"
	"log"
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

		if opcode > 99 {
			opcode, steps = checkForParameterMode(program, opcode, i)
			// Parameter mode could contain a 99 opcode
			if opcode == 99 {
				break
			}
		} else {
			switch opcode {
			case 1:
				param1Pos, param2Pos, resPos := findPointers(program, i)
				program[resPos] = program[param1Pos] + program[param2Pos]
				steps = 4
			case 2:
				param1Pos, param2Pos, resPos := findPointers(program, i)
				program[resPos] = program[param1Pos] * program[param2Pos]
				steps = 4
			case 3:
				paramPos := program[i+1]
				program[paramPos] = input
				steps = 2
			case 4:
				paramPos := program[i+1]
				fmt.Println(program[paramPos])
				steps = 2
			case 99:
				break
			default:
				log.Fatalf("Unsupported opcode: %d\n", opcode)
			}
		}
	}
}

func findPointers(program []int, instrPointer int) (int, int, int) {
	param1Pos := program[instrPointer+1]
	param2Pos := program[instrPointer+2]
	resPos := program[instrPointer+3]
	return param1Pos, param2Pos, resPos
}

func checkForParameterMode(program []int, opcode, instrPointer int) (int, int) {
	split := strings.Split(strconv.Itoa(opcode), "")
	// Merge the two right-most digit to get the opcode
	opcode, err := strconv.Atoi(split[len(split)-2] + split[len(split)-1])
	if err != nil {
		log.Fatalf("Cannot parse int: %v\n", err)
	}

	// Corner cases for which parameter mode has no meaning
	switch opcode {
	case 4: // Print value and continue
		paramPos := program[instrPointer+1]
		fmt.Println(program[paramPos])
		return 4, 2
	case 99: // Program will stop
		return 99, 0
	}

	val1 := 0
	val2 := 0
	switch len(split) {
	case 3:
		val1 = chooseMode(split[0], program, instrPointer)
		// Implied zero, so position mode
		paramPos := program[instrPointer+2]
		val2 = program[paramPos]
	case 4:
		val1 = chooseMode(split[1], program, instrPointer)
		val2 = chooseMode(split[0], program, instrPointer)
	default:
		log.Fatalf("Unexpected parameters: %v - len: %d\n", split, len(split))
	}

	resPos := program[instrPointer+3] // Never in immediate mode

	switch opcode {
	case 1:
		program[resPos] = val1 + val2
	case 2:
		program[resPos] = val1 * val2
	default:
		log.Fatalf("Unsupported opcode in parameter mode: %d\n", opcode)
	}

	return opcode, 4
}

func chooseMode(mode string, program []int, i int) int {
	val := 0
	switch mode {
	case "1": // Immediate mde
		val = program[i+1]
	case "0": // Position mode
		paramPos := program[i+1]
		val = program[paramPos]
	default:
		log.Fatalf("Unexpected mode: %v\n", mode)
	}

	return val
}
