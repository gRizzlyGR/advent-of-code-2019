package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	set1 := getPointsSet(scanner.Text())

	scanner.Scan()
	set2 := getPointsSet(scanner.Text())

	min := math.MaxInt32
	for p, steps1 := range set1 {
		if steps2, ok := set2[p]; ok {
			man := steps1 + steps2
			if man < min {
				min = man
			}
		}
	}

	fmt.Println(min)
}

func manhattan(a, b point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func getPointsSet(wire string) map[point]int {
	path := strings.Split(wire, ",")
	set := make(map[point]int)

	var c point
	totSteps := 0
	for _, step := range path {
		var next point
		distance, _ := strconv.Atoi(step[1:])

		for i := 0; i < distance; i++ {
			switch step[0] {
			case 'R':
				next = point{c.x + 1, c.y}
			case 'L':
				next = point{c.x - 1, c.y}
			case 'U':
				next = point{c.x, c.y + 1}
			case 'D':
				next = point{c.x, c.y - 1}
			}
			totSteps++
			set[next] = totSteps 
			c = next
		}
	}

	return set
}
