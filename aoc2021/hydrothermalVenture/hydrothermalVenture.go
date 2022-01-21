package hydrothermalventure

import (
	"bufio"
	"strconv"
	"strings"
)

type points struct {
	x int
	y int
}

type lines struct {
	start points
	end   points
}

func GetPoints(start points, end points) []points {

	var pointsinLine []points

	i, j := start.x, start.y
	pointsinLine = append(pointsinLine, points{x: start.x, y: start.y})

	for {
		if i == end.x && j == end.y {
			break
		} else if i < end.x {
			i++
			if j < end.y {
				j++
			} else if j > end.y {
				j--
			}
		} else if i > end.x {
			i--
			if j < end.y {
				j++
			} else if j > end.y {
				j--
			}
		} else if i == end.x {
			if j < end.y {
				j++
			} else if j > end.y {
				j--
			}
		}
		pointsinLine = append(pointsinLine, points{x: i, y: j})
	}
	return pointsinLine
}

func PointsOverLap(input []lines, diagonalLine bool) int {
	var output int
	var allPoints = make(map[points]int)

	for _, l := range input {
		if l.start.x != l.end.x && l.start.y != l.end.y && !diagonalLine {
			continue
		}

		linePoints := GetPoints(l.start, l.end)
		for _, p := range linePoints {
			allPoints[p] += 1
		}
	}

	for _, v := range allPoints {
		if v > 1 {
			output++
		}
	}
	return output
}

func ReadInput(scanInput *bufio.Scanner) []lines {
	var input []lines

	var linepoints lines
	var startPoints, endPoints []string

	for scanInput.Scan() {
		line := scanInput.Text()
		if line != "" {
			points := strings.Split(line, "->")
			startPoints = strings.Split(points[0], ",")
			endPoints = strings.Split(points[1], ",")
		}
		x1, _ := strconv.Atoi(strings.TrimSpace(startPoints[0]))
		y1, _ := strconv.Atoi(strings.TrimSpace(startPoints[1]))
		x2, _ := strconv.Atoi(strings.TrimSpace(endPoints[0]))
		y2, _ := strconv.Atoi(strings.TrimSpace(endPoints[1]))

		linepoints = lines{
			start: points{x: x1, y: y1},
			end:   points{x: x2, y: y2},
		}
		input = append(input, linepoints)
	}
	return input
}

func HydrathermalVenture(scanInput *bufio.Scanner) (int, int) {
	part1 := PointsOverLap(ReadInput(scanInput), false)
	part2 := PointsOverLap(ReadInput(scanInput), true)
	return part1, part2
}

// Example to run this code in main.go

// func main() {
// 	scanner := helper.ReadFile("aoc2021/hydrothermalventure/input.txt")
// 	outputPart1, outputPart2 := hydrothermalventure.HydrathermalVenture(scanner)
// 	fmt.Println(outputPart1, outputPart2)
// }
