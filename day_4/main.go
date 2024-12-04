package main

import (
	ir "aoc_2024/inputreader"
	"fmt"
)

func expandInput(data []string, yMax int) []string {
	var emptyRow string
	var output []string

	for i := 0; i < (yMax + 6); i++ {
		emptyRow += "O"
	}

	for j := 0; j < 3; j++ {
		output = append(output, emptyRow)
	}

	for idx := range data {
		output = append(output, "000"+data[idx]+"000")
	}

	for j := 0; j < 3; j++ {
		output = append(output, emptyRow)
	}

	return output
}

func searchSubstr(data []string, yStart, xStart int) int {
	// right, left, down, up, right-up, right-down, left-up, left-down
	substrings := [8]string{"", "", "", "", "", "", "", ""}
	noFound := 0

	for i := 0; i < 4; i++ {
		substrings[0] += string(data[yStart][xStart+i])
		substrings[1] += string(data[yStart][xStart-i])
		substrings[2] += string(data[yStart+i][xStart])
		substrings[3] += string(data[yStart-i][xStart])
		substrings[4] += string(data[yStart+i][xStart+i])
		substrings[5] += string(data[yStart-i][xStart+i])
		substrings[6] += string(data[yStart+i][xStart-i])
		substrings[7] += string(data[yStart-i][xStart-i])
	}

	for _, substring := range substrings {
		if substring == "XMAS" {
			noFound += 1
		}
	}

	return noFound
}

func one(data []string) {
	var total int
	// var cnt int

	for yidx := range data {
		for xidx := range data[yidx] {
			if string(data[yidx][xidx]) == "X" {
				total += searchSubstr(data, yidx, xidx)
			}
		}
	}

	fmt.Printf("Task 1: %v\n", total)
}

func main() {
	// data := ir.ReadText("example.txt")
	data := ir.ReadText("input.txt")

	parsedData := expandInput(data, len(data))

	one(parsedData)
}
