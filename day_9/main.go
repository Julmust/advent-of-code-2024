package main

import (
	ir "aoc_2024/inputreader"
	"fmt"
	"strconv"
)

func printData(data []int) {
	for _, i := range data {
		if i == -1 {
			fmt.Print(".")
			continue
		}

		fmt.Print(i)
	}
	fmt.Print("\n")
}

func parseData(data []string) []int {
	var tmp []int
	// everything's on one row
	d := data[0]

	for _, ch := range d {
		i, _ := strconv.Atoi(string(ch))
		tmp = append(tmp, i)
	}

	var output []int
	var addInt int
	for idx, val := range tmp {
		if idx%2 == 0 {
			addInt = idx / 2
		} else {
			addInt = -1
		}
		for i := 0; i < val; i++ {
			output = append(output, addInt)
		}
	}

	return output
}

func one(data []int) {
	l, r := 0, len(data)-1

	for {
		// locate first dot, represented as -1
		if data[l] != -1 {
			l++
			continue
		}
		//locate last digit
		if data[r] == -1 {
			r--
			continue
		}

		if l > r {
			break
		}
		data[l], data[r] = data[r], data[l]

	}

	var res int
	for idx, i := range data {
		// When we hit the first "dot" we're done
		if i == -1 {
			break
		}

		res += idx * i
	}

	fmt.Printf("Task 1: %v\n", res)
}

func main() {
	// data := ir.ReadText("example.txt")
	data := ir.ReadText("input.txt")

	parsedData := parseData(data)
	one(parsedData)
}
