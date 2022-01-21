package main

import (
	hydrothermalventure "aoc/aoc2021/hydrothermalVenture"
	"aoc/helper"
	"fmt"
)

func main() {
	scanner := helper.ReadFile("aoc2021/hydrothermalventure/input.txt")
	outputPart1, outputPart2 := hydrothermalventure.HydrathermalVenture(scanner)
	fmt.Println(outputPart1, outputPart2)
}
