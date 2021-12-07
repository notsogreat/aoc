package helper

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFile(filepath string) *bufio.Scanner {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	return scanner
}

// filepath example: aoc2021/SonarSweep/input.txt
func ConvertInputIntoSliceOfInt(scanInput *bufio.Scanner) []int {
	output := []int{}
	for scanInput.Scan() {
		num, _ := strconv.Atoi(scanInput.Text())
		output = append(output, num)
	}

	if err := scanInput.Err(); err != nil {
		log.Fatal(err)
	}
	return output
}
