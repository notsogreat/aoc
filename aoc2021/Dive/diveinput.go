package dive

import (
	"bufio"
	"strconv"
	"strings"
)

//
func ConverInputIntoDataStruct(scanInput *bufio.Scanner) []Data {
	var data []Data
	for scanInput.Scan() {
		line := strings.Split(scanInput.Text(), " ")
		v, _ := strconv.Atoi(line[1])
		data = append(data, Data{
			position: line[0],
			value:    v,
		})
	}
	return data
}
