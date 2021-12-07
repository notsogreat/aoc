package dive

import "fmt"

type Data struct {
	position string
	value    int
}

//Part1 SubmarineDepth calculation
func SubmarineDepth(input []Data) int {
	var horizontalPosition, verticlePosition int
	for _, v := range input {
		switch v.position {
		case "forward":
			horizontalPosition = horizontalPosition + v.value
		case "down":
			verticlePosition = verticlePosition + v.value
		case "up":
			verticlePosition = verticlePosition - v.value
		default:
			fmt.Printf("Not able to indentify position %s", v.position)
		}
	}
	return horizontalPosition * verticlePosition
}

//Part 2 CorrectSubmarineDepth after reading manual
func CorrectSubmarineDepth(input []Data) int {
	var horizontalPosition, verticlePosition, aim int
	for _, v := range input {
		switch v.position {
		case "forward":
			horizontalPosition = horizontalPosition + v.value
			if aim != 0 {
				verticlePosition = verticlePosition + (v.value * aim)
			}
		case "down":
			aim = aim + v.value
		case "up":
			aim = aim - v.value
		}
	}
	return horizontalPosition * verticlePosition
}
