package main

import (
	ir "aoc_2024/inputreader"
	"fmt"
)

func parseData(data []string) map[string][][]int {
	m := make(map[string][][]int)

	for yidx := range data {
		for xidx, char := range data[yidx] {
			if string(char) != "." {
				m[string(char)] = append(m[string(char)], []int{xidx, yidx})
			}
		}
	}

	return m
}

func createEmptyBoard(w, h int) [][]string {
	var row []string
	var board [][]string

	for i := 0; i < w; i++ {
		row = append(row, ".")
	}

	for j := 0; j < h; j++ {
		nrow := make([]string, len(row))
		copy(nrow, row)
		board = append(board, nrow)
	}

	return board
}

func printResultBoard(resultBoard [][]string) {
	for _, i := range resultBoard {
		fmt.Println(i)
	}
}

func checkBounds(x, y int, rB [][]string) bool {
	if x >= 0 && x < len(rB[0]) && y >= 0 && y < len(rB[1]) {
		return true
	}
	return false
}

func calcResults(rB [][]string) int {
	var o int
	for _, row := range rB {
		for _, v := range row {
			if v == "#" {
				o += 1
			}
		}
	}

	return o
}

func one(parsedData map[string][][]int, resultBoard [][]string) {
	fmt.Println(parsedData)
	for _, coords := range parsedData {
		for idx := range coords {
			start := coords[idx]

			for i := (idx + 1); i < len(coords); i++ {
				x, y := start[0]-coords[i][0], start[1]-coords[i][1]

				bX := start[0] + x
				bY := start[1] + y

				if checkBounds(bX, bY, resultBoard) {
					resultBoard[bY][bX] = "#"
				}

				bX = coords[i][0] + (x * -1)
				bY = coords[i][1] + (y * -1)
				if checkBounds(bX, bY, resultBoard) {
					resultBoard[bY][bX] = "#"
				}
			}
		}
	}
	printResultBoard(resultBoard)
	fmt.Printf("Task 1: %v\n", calcResults(resultBoard))
}

func main() {
	// data := ir.ReadText("example.txt")
	data := ir.ReadText("input.txt")

	parsedData := parseData(data)
	eb := createEmptyBoard(len(data[0]), len(data))

	one(parsedData, eb)
}
