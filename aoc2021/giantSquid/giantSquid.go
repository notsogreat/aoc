package giantsquid

import (
	"aoc/helper"
	"bufio"
	"fmt"
	"reflect"
)

func Match(grid [][]int, value int, matchPosition [5][5]int) [5][5]int {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if grid[i][j] == value {
				matchPosition[i][j] = 1
			}
		}
	}
	return matchPosition
}

func CheckRowDone(grid [5][5]int) bool {
	for _, v := range grid {
		if reflect.DeepEqual(v, [5]int{1, 1, 1, 1, 1}) {
			return true
		}
	}
	return false
}

func CheckColumnDone(grid [5][5]int) bool {
	for i := 0; i < 5; i++ {
		var myColumn []int
		for j := 0; j < 5; j++ {
			myColumn = append(myColumn, grid[j][i])
		}
		if reflect.DeepEqual(myColumn, []int{1, 1, 1, 1, 1}) {
			return true
		}
	}
	return false
}

func Bingo(grid [][]int, seq []int) ([5][5]int, [][]int, int, bool, []int) {

	var matchOutput [5][5]int
	var matchedSeq []int

	for i, v := range seq {
		matchedSeq = append(matchedSeq, v)
		matchOutput = Match(grid, v, matchOutput)
		if i > 4 {
			if CheckRowDone(matchOutput) || CheckColumnDone(matchOutput) {
				return matchOutput, grid, v, true, matchedSeq
			}
		}
	}
	return [5][5]int{}, [][]int{}, 0, false, matchedSeq
}

func RunAllBingo(inputs [][][]int, seq []int) int {

	var overallSum int

	var newMatchOutput [5][5]int
	var matchOutput [5][5]int

	var outputGrid [][]int
	var newGrid [][]int

	var number int
	var newNumber int

	var isMatched bool
	var matchedSeq []int
	var pre int

	for _, grid := range inputs {
		matchOutput, outputGrid, number, isMatched, matchedSeq = Bingo(grid, seq)
		if isMatched && (len(matchedSeq) > pre || pre == 0) {
			newMatchOutput = matchOutput
			newGrid = outputGrid
			newNumber = number
			pre = len(matchedSeq)
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if newMatchOutput[i][j] == 0 {
				overallSum = overallSum + newGrid[i][j]
			}
		}
	}
	return overallSum * newNumber
}

func ReadInput(scanInput *bufio.Scanner) {
	var lineSlice [][]int

	for scanInput.Scan() {
		line := scanInput.Text()
		if line != "" {
			values := helper.ConvertLineIntoSliceOfInt(line)
			lineSlice = append(lineSlice, values)
		}
	}

	var seqInput = lineSlice[0]
	grids := [][][]int{}

	for i := 1; i <= len(lineSlice)-5; i = i + 5 {
		miniGrid := [][]int{}
		miniGrid = append(miniGrid, lineSlice[i])
		miniGrid = append(miniGrid, lineSlice[i+1])
		miniGrid = append(miniGrid, lineSlice[i+2])
		miniGrid = append(miniGrid, lineSlice[i+3])
		miniGrid = append(miniGrid, lineSlice[i+4])
		grids = append(grids, miniGrid)
	}

	fmt.Println(RunAllBingo(grids, seqInput))
}
