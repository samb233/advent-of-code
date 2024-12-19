package main

import (
	"fmt"
	"os"
	// "sort"
	// "strconv"
	"strings"
)

// part 1
func main() {
	input, err := os.ReadFile("19.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n\n")

	towels := strings.Split(lines[0], ", ")
	targets := strings.Split(lines[1], "\n")

	res := 0
	for _, target := range targets {
		if len(target) == 0 {
			continue
		}

		// if CanMake(towels, target) {
		// 	res++
		// }
		res += CanMakeNums(towels, target)
	}

	fmt.Println(res)
}

var DP map[string]bool = make(map[string]bool)

func CanMake(words []string, target string) bool {
	if len(target) == 0 {
		return true
	}

	can, ok := DP[target]
	if ok {
		return can
	}

	for _, word := range words {
		if len(target) < len(word) {
			continue
		}

		if word == target[:len(word)] {
			can = CanMake(words, target[len(word):])
			if can {
				DP[target] = true
				break
			}
		}
	}

	DP[target] = can
	return can
}

var DPNum map[string]int = make(map[string]int)

func CanMakeNums(words []string, target string) int {
	if len(target) == 0 {
		return 1
	}

	num, ok := DPNum[target]
	if ok {
		return num
	}

	for _, word := range words {
		if len(target) < len(word) {
			continue
		}

		if word == target[:len(word)] {
			t := CanMakeNums(words, target[len(word):])
			if t > 0 {
				num += t
			}
		}
	}

	DPNum[target] = num
	return num
}
