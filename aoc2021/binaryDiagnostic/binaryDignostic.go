package binarydiagnostic

import (
	"bufio"
	"strconv"
	"strings"
)

// InputSetup will setup the input for day 3 from text file
func InputSetup(scanInput *bufio.Scanner) []string {
	data := []string{}

	for scanInput.Scan() {
		line := scanInput.Text()
		data = append(data, line)
	}
	return data
}

func PowerConsumption(input []string) (int64, error) {
	var mostCommon, listCommon string

	for i := 0; i < len(input[0]); i++ {
		var numberOfZero, numberOfOnes int

		for j := 0; j < len(input); j++ {
			s := strings.Split(input[j], "")
			if s[i] == "0" {
				numberOfZero++
			} else {
				numberOfOnes++
			}
		}
		if numberOfOnes > numberOfZero {
			mostCommon = mostCommon + "1"
			listCommon = listCommon + "0"
		} else {
			mostCommon = mostCommon + "0"
			listCommon = listCommon + "1"
		}
	}

	gammaRate, err := strconv.ParseInt(mostCommon, 2, 64)
	if err != nil {
		return 0, err
	}
	epsilonRate, err := strconv.ParseInt(listCommon, 2, 64)
	if err != nil {
		return 0, err
	}
	return gammaRate * epsilonRate, nil
}

func MatchValue(input []string, value string, position int) []string {

	var newData []string
	for i := 0; i < len(input); i++ {
		s := strings.Split(input[i], "")
		if s[position] == value {
			newData = append(newData, input[i])
		}
	}
	return newData
}

func CalculateCommonBits(input []string, position int) (string, string) {
	var numberOfZero, numberOfOnes int
	for j := 0; j < len(input); j++ {
		s := strings.Split(input[j], "")
		if s[position] == "0" {
			numberOfZero++
		} else {
			numberOfOnes++
		}
	}
	if numberOfZero > numberOfOnes {
		return "0", "1"
	} else if numberOfZero <= numberOfOnes {
		return "1", "0"
	}
	return "0", "0"
}

func CalculateMostCommon(input []string) []string {
	var newMostCommonInput []string
	var mostCommon string
	for i := 0; i <= len(input[0]); i++ {
		if i == 0 {
			mostCommon, _ = CalculateCommonBits(input, i)
			newMostCommonInput = input
		} else {
			newMostCommonInput = MatchValue(newMostCommonInput, mostCommon, i-1)
			if len(newMostCommonInput) == 1 {
				return newMostCommonInput
			}
			mostCommon, _ = CalculateCommonBits(newMostCommonInput, i)
		}
	}
	return nil
}

func CalculateListCommon(input []string) []string {
	var listCommon string
	var newListCommonInput []string
	for i := 0; i <= len(input[0]); i++ {
		if i == 0 {
			_, listCommon = CalculateCommonBits(input, i)
			newListCommonInput = input
		} else {
			newListCommonInput = MatchValue(newListCommonInput, listCommon, i-1)
			if len(newListCommonInput) == 1 {
				return newListCommonInput
			}
			_, listCommon = CalculateCommonBits(newListCommonInput, i)
		}
	}
	return nil
}

func LifeSupport(input []string) (int64, error) {
	mostC := CalculateMostCommon(input)
	ListC := CalculateListCommon(input)

	oxygenGeneratorRating, err := strconv.ParseInt(mostC[0], 2, 64)
	if err != nil {
		return 0, err
	}
	Co2ScrubberRating, err := strconv.ParseInt(ListC[0], 2, 64)
	if err != nil {
		return 0, err
	}
	return oxygenGeneratorRating * Co2ScrubberRating, nil
}
