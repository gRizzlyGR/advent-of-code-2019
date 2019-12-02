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

	var opcodes []int

	for scanner.Scan() {
		opcodes = convertStringToInts(scanner.Text())
	}

	opcodes[1] = 12
	opcodes[2] = 2

	fmt.Println(compute(opcodes))

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

func compute(ints []int) int {
	for i := 0; i < len(ints); i += 4 {
		instruction := ints[i]
		if instruction == 99 {
			break
		}

		param1Pos := ints[i+1]
		param2Pos := ints[i+2]
		resPos := ints[i+3]

		switch instruction {
		case 1:
			ints[resPos] = ints[param1Pos] + ints[param2Pos]
		case 2:
			ints[resPos] = ints[param1Pos] * ints[param2Pos]
		default:
			log.Fatalf("Unsupported instruction: %d\n", instruction)
		}
	}

	return ints[0]
}
