package lanternfish

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func internalTimerCheck(fishState []int) []int {
	var newFishState []int

	for _, fs := range fishState {
		if fs == 0 {
			newFishState = append(newFishState, 6)
			newFishState = append(newFishState, 8)
			continue
		}
		fs--
		newFishState = append(newFishState, fs)
	}
	return newFishState
}

func internalTimerCheckOptimized(fishState []int) []int {
	l := len(fishState)
	for i := 0; i < l; i++ {
		if fishState[i] == 0 {
			fishState[i] = 6
			fishState = append(fishState, 8)
			continue
		}
		fishState[i]--
	}
	return fishState
}

func CalculateNumberOfFish(InitialState []int, days int) int {
	var fishState = InitialState
	for i := 1; i <= days; i++ {
		min := getMin(fishState)
		if min > 0 {
			for j := range fishState {
				fishState[j] = fishState[j] - min
			}
		}
		i = i + min
		fishState = internalTimerCheckOptimized(fishState)
		fmt.Println(i)
	}
	return len(fishState)
}

func getMin(seq []int) int {
	var min = seq[0]
	for i := 1; i < len(seq); i++ {
		if min > seq[i] {
			min = seq[i]
		}
	}
	return min
}

func ReadInput(scanInput *bufio.Scanner) []int {
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
			fmt.Println(err)
			break
		}
		input = append(input, num)
	}
	return input
}

func LaternFish(input *bufio.Scanner) int {
	return CalculateNumberOfFish(ReadInput(input), 256)
}

// Efficiant code

// func CalculateNumberOfFish1(InitialState []int, days int) int {
// 	for i := 1; i <= days; i++ {
// 		for j, f := range InitialState {
// 			remFishLife, newFish := internalTimerCheck1(f, days)
// 			InitialState[j] = remFishLife
// 			InitialState = append(InitialState, newFish...)
// 		}
// 		fmt.Println(InitialState, i)
// 	}
// 	return len(InitialState)
// }

// func internalTimerCheck1(fishLife int, days int) (int, []int) {
// 	var addState []int

// 	for i := 1; i <= days; i++ {
// 		if fishLife == 0 {
// 			fishLife = 6
// 			addState = append(addState, 8)
// 			continue
// 		}
// 		fishLife--
// 	}
// 	return fishLife, addState
// }
