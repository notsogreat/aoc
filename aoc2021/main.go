package main

import (
	"aoc2021/pkg/SonarSweep"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := []int{}

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		input = append(input, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	_, d := SonarSweep.CalculateDepthMeasureSlidingWindow(input)
	fmt.Println(d)
}
