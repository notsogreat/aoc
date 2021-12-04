package SonarSweep

func CalculateDepthMeasureIncrease(input []int) int {
	var increaseDepthMeasure int

	for i := 0; i < len(input)-1; i++ {
		if input[i+1] > input[i] {
			increaseDepthMeasure++
		}
	}
	return increaseDepthMeasure
}

func CalculateDepthMeasureSlidingWindow(input []int) ([]int, int) {
	var newData []int

	for i := 0; i < len(input)-2; i++ {
		total := input[i] + input[i+1] + input[i+2]
		newData = append(newData, total)
	}
	return newData, CalculateDepthMeasureIncrease(newData)
}
