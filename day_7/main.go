package main

import (
	ir "aoc_2024/inputreader"
	"fmt"
	"strconv"
	"strings"
)

func parseData(data []string) [][]int {
	var p [][]int

	for _, row := range data {
		var values []int
		splitRow := strings.Split(row, ": ")
		key, _ := strconv.Atoi(splitRow[0])
		values = append(values, key)

		strValues := strings.Split(splitRow[1], " ")
		for _, i := range strValues {
			a, _ := strconv.Atoi(i)
			values = append(values, a)
		}

		p = append(p, values)
	}

	return p
}

func getPermutations(noOps int) []string {
	input := "+*"

	if noOps <= 0 {
		return nil
	}

	prod := make([]string, len(input))
	for i, char := range input {
		prod[i] = string(char)
	}

	for i := 1; i < noOps; i++ {
		next := make([]string, 0, len(input)*len(prod))

		for _, word := range prod {
			for _, char := range input {
				next = append(next, word+string(char))
			}
		}

		prod = next
	}

	return prod
}

func calc(a, b int, op string) int {
	if op == "*" {
		return a * b
	}
	return a + b
}

func validateEntry(target int, values []int) bool {
	permutations := getPermutations(len(values) - 1)
	for _, perm := range permutations {
		var tot = 0
		for idx := range values {
			if idx == 0 {
				tot = values[0]
				continue
			}
			op := string(perm[idx-1])
			tot = calc(tot, values[idx], op)
		}
		if tot == target {
			return true
		}
	}

	return false
}

func one(data [][]int) {
	var res int
	for _, v := range data {
		if validateEntry(v[0], v[1:]) {
			res += v[0]
		}
	}

	fmt.Printf("Part 1: %v\n", res)
}

func main() {
	// data := ir.ReadText("example.txt")
	data := ir.ReadText("input.txt")

	parsedData := parseData(data)

	one(parsedData)
}
