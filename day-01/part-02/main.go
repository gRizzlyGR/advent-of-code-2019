package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		log.Fatalf("Cannot open file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fuel := 0

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Cannot read integer: %v", err)
		}

		fuel += computeFuel(mass)
	}

	fmt.Println(fuel)

}

func computeFuel(n int) int {
	fuel := int(math.Floor(float64(n/3))) - 2
	if fuel <= 0 {
		return 0
	}

	return fuel + computeFuel(fuel)
}
