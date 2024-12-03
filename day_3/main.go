package main

import (
	ir "aoc_2024/inputreader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	items []string
	total int
}

func (s *Stack) push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) clear() {
	s.items = nil
}

func (s *Stack) attemptMul() (int, error) {
	str := strings.Join(s.items, "")
	reValidate, _ := regexp.Compile(`mul\(\d*,\d*\)`)
	reExtract, _ := regexp.Compile(`\d*,\d*`)

	if reValidate.MatchString(str) {
		opStr := reExtract.FindString(str)
		opSlice := strings.Split(opStr, ",")

		// fmt.Printf("Op: %v * %v\n", opSlice[0], opSlice[1])

		f, _ := strconv.Atoi(opSlice[0])
		s, _ := strconv.Atoi(opSlice[1])

		return f * s, nil
	}

	return 0, fmt.Errorf("not a valid operation")
}

func isNumeric(char string) bool {
	_, err := strconv.Atoi(char)

	return err == nil
}

func isAllowedAlphaCharacter(char string) bool {
	allowedAlphaCharacters := []string{"m", "u", "l", ",", "("}

	for _, allowed := range allowedAlphaCharacters {
		if char == allowed {
			return true
		}
	}

	return false
}

func one(data string) {
	s := Stack{total: 0}
	isActive := false

	for _, char := range data {
		strChr := string(char)

		if !isActive && strChr == "m" {
			isActive = true
		}

		if isActive {
			if isAllowedAlphaCharacter(strChr) || isNumeric(strChr) {
				s.push(strChr)
			} else {
				if strChr == ")" {
					s.push(strChr)
				}
				isActive = false
				val, err := s.attemptMul()
				if err == nil {
					s.total += val
				}
				s.clear()
			}
		}
	}

	fmt.Printf("Task 1: %v\n", s.total)
}

func main() {
	// data := ir.ReadText("example.txt")
	data := ir.ReadText("input.txt")

	// fmt.Println(data)
	one(data[0])
}
