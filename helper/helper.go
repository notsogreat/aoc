package helper

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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

func ConvertLineIntoSliceOfInt(line string) []int {
	// var num int
	output := []int{}
	var value []string
	// var err error

	if strings.Contains(line, ",") {
		value = strings.Split(line, ",")
	} else {
		value = strings.Split(line, " ")
	}

	for _, v := range value {
		num, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		output = append(output, num)
	}

	return output
}
