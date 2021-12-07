package main

import (
	dive "aoc/aoc2021/Dive"
	"aoc/helper"
	"fmt"
)

func main() {
	scanner := helper.ReadFile("aoc2021/Dive/input.txt")
	input := dive.ConverInputIntoDataStruct(scanner)
	fmt.Println(dive.CorrectSubmarineDepth(input))
}
