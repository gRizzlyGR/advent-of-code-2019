package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	start, err1 := strconv.Atoi(args[1])
	if err1 != nil {
		log.Fatalf("Cannot parse first argument: %v\n", err1)
	}

	end, err2 := strconv.Atoi(args[2])
	if err2 != nil {
		log.Fatalf("Cannot parse second argument: %v\n", err1)
	}

	numDigits := countDigits(start)

	howManyPws := 0
	// The value is within the range given in your puzzle input.
	for i := start; i <= end; i++ {
		if isValidPassword(i, numDigits) {
			howManyPws++
		}
	}

	fmt.Println(howManyPws)
}

func isValidPassword(n, numDigits int) bool {
	var notDecreasing bool
	var isDouble bool

	// Handle the power of 10
	// 123 -> 3 digits, so we want i € {2, 1, 0} to get 10^2, 10^1 and 10^0
	for i := numDigits - 1; i > 0; i-- {
		// Find the digit in the current position (i) and its successor
		// To get 1 from 123 -> 123 / 10^2 = 1; 1 % 10 == 1
		// Don't need mod in the first loop, but is needed for the next ones
		// To get 2 from 123 -> 123 / 10^1 = 12; 12 % 10 == 2
		// To get 3 from 124 -> 123 / 10^0 = 123; 123 % 10 == 3
		d1 := n / int(math.Pow10(i)) % 10
		d2 := n / int(math.Pow10(i-1)) % 10

		// If still haven't found a double, keep looking for it in the next loops
		if !isDouble {
			// Two adjacent digits are the same (like 22 in 122345)
			isDouble = d1 == d2
		}

		// Going from left to right, the digits never decrease;
		// they only ever increase or stay the same (like 111123 or 135679).
		notDecreasing = d1 <= d2

		// If the second digit is lower than the first one, don't need to keep on looping
		if !notDecreasing {
			return false
		}
	}

	return notDecreasing && isDouble
}

func countDigits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}

	return count
}
