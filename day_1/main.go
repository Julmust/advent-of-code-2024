package main

import (
	ir "aoc_2024/inputreader"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parseData(data []string) ([]int, []int) {
	var left []int
	var right []int

	for _, d := range data {
		rawdata := strings.ReplaceAll(d, "   ", ",")
		numstrings := strings.Split(rawdata, ",")
		// fmt.Println(numstrings)
		leftint, _ := strconv.Atoi(numstrings[0])
		rightint, _ := strconv.Atoi(numstrings[1])

		left = append(left, leftint)
		right = append(right, rightint)
	}

	return left, right
}

func sortSlice(sl []int) []int {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})

	return sl
}

// Writing my own Abs func that takes integers instead of floats :P
func myAbs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func two(ldata []int, rdata []int) {
	var res int

	itemCounts := make(map[int]int)
	for _, i := range rdata {
		itemCounts[i] += 1
	}

	for _, i := range ldata {
		res += i * itemCounts[i]
	}

	fmt.Printf("Task 2: %v\n", res)
}

func one(ldata []int, rdata []int) {
	ldata = sortSlice(ldata)
	rdata = sortSlice(rdata)
	var deltas []int
	for idx := range ldata {
		deltas = append(deltas, myAbs(ldata[idx], rdata[idx]))
	}

	var res int
	for _, v := range deltas {
		res += v
	}

	fmt.Printf("Task 1: %v\n", res)
}

func main() {
	data := ir.ReadText("input.txt")
	// data := ir.ReadText("example.txt")

	ldata, rdata := parseData(data)
	one(ldata, rdata)
	two(ldata, rdata)
}
