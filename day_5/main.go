package main

import (
	ir "aoc_2024/inputreader"
	"fmt"
	"strconv"
	"strings"
)

type rule struct {
	base_page  string
	downstream []string
}

func parseData(data []string) ([]string, []string) {
	var rules []string
	var pages []string

	for idx, v := range data {
		if v == "" {
			rules = append(rules, data[:idx]...)
			pages = append(pages, data[idx+1:]...)
		}
	}

	return rules, pages
}

func parseRules(rawRules []string) map[string]rule {
	m := make(map[string]rule)

	for idx := range rawRules {
		ruleSlice := strings.Split(rawRules[idx], "|")
		base, ref := ruleSlice[0], ruleSlice[1]

		if val, exists := m[base]; exists {
			val.downstream = append(val.downstream, ref)
			m[base] = val
		} else {
			r := rule{base_page: base, downstream: []string{ref}}
			m[base] = r
		}
	}

	return m
}

func checkOneVal(curr string, after []string, m map[string]rule) bool {
	for _, val := range after {
		isValid := false

		// Check if the value we're looking at is in the downstream of the
		//  base value. If it's not we've found a rule violation
		for _, ds := range m[curr].downstream {
			if val == ds {
				isValid = true
			}
		}
		if !isValid {
			return false
		}
	}

	return true
}

func getMid(pages []string) int {
	l := len(pages)
	i, _ := strconv.Atoi(pages[l/2])
	return i
}

func one(rawRules, rawPages []string) {
	m := parseRules(rawRules)
	op := 0
	// for _, v := range m {
	// 	fmt.Println(v)
	// }

	for _, entry := range rawPages {
		valid := true
		pages := strings.Split(entry, ",")
		for idx := range pages {
			// curr, before, after := pages[idx], pages[:idx], pages[idx+1:]
			// idk why but I dont have to check previous pages?
			curr, after := pages[idx], pages[idx+1:]
			validPage := checkOneVal(curr, after, m)

			if !validPage {
				valid = false
				break
			}
		}
		if valid {
			op += getMid(pages)
		}
	}

	fmt.Printf("Part 1: %v\n", op)
}

func main() {
	// data := ir.ReadText("example.txt")
	data := ir.ReadText("input.txt")
	rules, pages := parseData(data)

	one(rules, pages)
}
