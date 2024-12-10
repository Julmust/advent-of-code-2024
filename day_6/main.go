package main

import (
	ir "aoc_2024/inputreader"
	"fmt"
	"strings"
)

type guard struct {
	locX, locY, facingDirY, facingDirX int
	prevLocations                      map[string]int
}

func (g *guard) makeMove(board [][]string) bool {
	// func (g *guard) makeMove(board [][]string) ([][]string, bool) {
	nextY := g.locY + g.facingDirY
	nextX := g.locX + g.facingDirX

	if board[nextY][nextX] == "O" {
		return true
	} else if board[nextY][nextX] == "#" {
		if g.facingDirY == -1 && g.facingDirX == 0 {
			g.facingDirY, g.facingDirX = 0, 1
		} else if g.facingDirY == 0 && g.facingDirX == 1 {
			g.facingDirY, g.facingDirX = 1, 0
		} else if g.facingDirY == 1 && g.facingDirX == 0 {
			g.facingDirY, g.facingDirX = 0, -1
		} else if g.facingDirY == 0 && g.facingDirX == -1 {
			g.facingDirY, g.facingDirX = -1, 0
		}
	} else {
		g.locX, g.locY = nextX, nextY
	}
	// if g.facingDir == "up" {
	// 	g.locY += -1
	// }
	board[g.locY][g.locX] = "*"
	g.prevLocations[fmt.Sprintf("(%v, %v)", g.locY, g.locX)] = 1

	return false
}

func printBoard(board [][]string) {
	for idx, i := range board {
		fmt.Println(idx, i)
	}
}

func createGameBoard(data []string) [][]string {
	var board [][]string
	var emptyRow []string
	emptyCell := []string{"O"}

	for i := 0; i < (len(data[0]) + 2); i++ {
		emptyRow = append(emptyRow, "O")
	}

	// Pad top of board with O's
	board = append(board, emptyRow)

	for _, row := range data {
		cells := strings.Split(row, "")

		// Pad left and right of row with O's
		cells = append(emptyCell, cells...)
		cells = append(cells, emptyCell...)

		board = append(board, cells)
	}

	// Pad bottom of board with O's
	board = append(board, emptyRow)

	return board
}

func one(board [][]string) {
	var g guard

	for yidx, row := range board {
		for xidx, cell := range row {
			if cell == "^" {
				coords := fmt.Sprintf("(%v, %v)", yidx, xidx)
				locations := make(map[string]int)
				locations[coords] = 1
				g = guard{
					locX:          xidx,
					locY:          yidx,
					facingDirY:    -1,
					facingDirX:    0,
					prevLocations: locations,
				}
			}
		}
	}

	for {
		// fin := g.makeMove(board)
		// printBoard(board)
		if g.makeMove(board) {
			break
		}
		// fmt.Println("================")
	}
	fmt.Printf("Task 1: %v\n", len(g.prevLocations))
}

func main() {
	// data := ir.ReadText("example.txt")
	data := ir.ReadText("input.txt")

	board := createGameBoard(data)

	// for idx, i := range board {
	// 	fmt.Println(idx, i)
	// }

	one(board)
}
