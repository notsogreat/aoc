package whales

import (
	"bufio"
	"strconv"
	"strings"
)

func findMinAndMax(values []int) (int, int) {
	var min, max int
	if values[0] > values[1] {
		max = values[0]
		min = values[1]
	} else {
		min = values[0]
		max = values[1]
	}

	for i := 2; i < len(values); i++ {
		if values[i] == 0 {
			continue
		}

		if min > values[i] {
			min = values[i]
		}
		if max < values[i] {
			max = values[i]
		}
	}
	return min, max
}

func exponential(v int) int {
	var output int
	for i := v; i >= 1; i-- {
		output += i
	}
	return output
}

func calculateFuel(positions []int, value int, constant bool) int {
	var sum int
	var v int

	for _, p := range positions {
		if (p - value) > 0 {
			v = (p - value)
		} else {
			v = (value - p)
		}
		if constant {
			sum = sum + v
		} else {
			sum = sum + exponential(v)
		}
	}
	return sum
}

func alignCrabs(positions []int, conRate bool) int {
	var minFuel = calculateFuel(positions, 0, conRate)
	cache := make(map[int]int)

	for i := 1; i <= positions[len(positions)-1]; i++ {
		if _, ok := cache[i]; ok {
			continue
		}
		fuel := calculateFuel(positions, i, conRate)
		cache[i] = fuel
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

func readInput(scanInput *bufio.Scanner) []int {
	var numbers []string
	var input []int
	for scanInput.Scan() {
		line := scanInput.Text()
		if line != "" {
			numbers = strings.Split(strings.TrimSpace(line), ",")
		}
	}
	for _, n := range numbers {
		num, err := strconv.Atoi(n)
		if err != nil {
			break
		}
		input = append(input, num)
	}
	return input
}

func FuelSpend(scanInput *bufio.Scanner) (int, int) {
	return alignCrabs(readInput(scanInput), true), alignCrabs(readInput(scanInput), false)
}

//RUN THIS CODE IN MAIN

// func main() {
// 	scanner := helper.ReadFile("aoc2021/whales/input.txt")
// 	fmt.Println(whales.FuelSpend(scanner))
// }
