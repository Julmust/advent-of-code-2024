package main

import (
	ir "aoc_2024/inputreader"
	"fmt"
	"strconv"
	"strings"
)

func parseData(data []string) [][]int {
	var opStr [][]string
	var opInt [][]int

	for _, v := range data {
		opStr = append(opStr, strings.Split(v, " "))
	}

	for _, ss := range opStr {
		var tmpArr []int
		for _, s := range ss {
			toInt, _ := strconv.Atoi(s)
			tmpArr = append(tmpArr, toInt)
		}
		opInt = append(opInt, tmpArr)
	}

	return opInt
}

func myAbs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func checkSlope(row []int) bool {
	row_max_idx := len(row)
	dir := 0

	for i := 1; i < row_max_idx; i++ {
		delta := row[i-1] - row[i]
		absDelta := myAbs(delta)

		// Checking that slope is not even and does not
		// increase too quickly
		if absDelta <= 0 || absDelta > 3 {
			return false
		}

		// Setting expected direction of slope
		if dir == 0 {
			if delta < 0 {
				dir = -1
			} else {
				dir = 1
			}
		}

		if dir*delta < 0 {
			return false
		}
	}

	return true
}

func one(data [][]int) {
	ans := 0
	for _, v := range data {
		if checkSlope(v) {
			ans++
		}
	}

	fmt.Printf("Task 1: %v\n", ans)

}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func generatePermutations(d []int) [][]int {
	var op [][]int
	baseSliceLen := len(d)

	for i := 0; i < baseSliceLen; i++ {
		os := make([]int, baseSliceLen)
		copy(os, d)
		tmpSlice := remove(os, i)
		op = append(op, tmpSlice)
	}

	return op
}

func two(data [][]int) {
	ans := 0
	for _, v := range data {
		success := checkSlope(v)
		if !success {
			// Generate permutations
			permutations := generatePermutations(v)

			for _, permutation := range permutations {
				success = checkSlope(permutation)
				if success {
					break
				}
			}
		}

		if success {
			ans++
		}
	}

	fmt.Printf("Task 2: %v\n", ans)

}

func main() {
	// data := ir.ReadText("example.txt")
	data := ir.ReadText("input.txt")

	parsedData := parseData(data)

	one(parsedData)
	two(parsedData)
}
