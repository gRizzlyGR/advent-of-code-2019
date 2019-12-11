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
	scanner.Scan()
	program := convertStringToInts(scanner.Text())

	//exec(program, 1)
	amplify(program)
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

func exec(program []int, input int) int {
	output := 0
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
			output = param
			steps = 2
		case 99:
			return output
		default:
			log.Fatalf("Unsupported opcode: %d\n", opcode)
		}
	}

	return -1
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

func amplify(program []int) {
	phases := []int{0, 1, 2, 3, 4}
	combinations := permutate(len(phases), phases)

	for _, combination := range combinations {
		fmt.Println("Phase setting:", combination)
		param := 0
		for _, phase := range combination {
			fmt.Printf("Param: %d\tPhase:%d\n", param, phase)
			// tmp := make([]int, len(program))
			// copy(tmp, program)

			// amplifier := []int{phase}
			// amplifier = append(amplifier, program...)
			amplifier := make([]int, len(program))
			copy(amplifier, program)
			fmt.Println(amplifier)
			param = exec(amplifier, param)
		}
	}
}

// Heap's algorithm
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func permutate(size int, ints []int) [][]int {
	var permutations [][]int
	if size == 1 {
		tmp := make([]int, len(ints))
		copy(tmp, ints)
		permutations = append(permutations, tmp)
	} else {
		permutations = append(permutations, permutate(size-1, ints)...)
	}

	for i := 0; i < size-1; i++ {
		if size%2 == 0 {
			ints[i], ints[size-1] = ints[size-1], ints[i]
		} else {
			ints[0], ints[size-1] = ints[size-1], ints[0]
		}
		permutations = append(permutations, permutate(size-1, ints)...)
	}

	return permutations
}
