package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("1.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	left, right, err := GetTwoLists(lines)
	if err != nil {
		panic(err)
	}

	// result := CountDistanceApart(left, right)
	result := CountSimilarity(left, right)
	fmt.Println(result)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func GetTwoLists(lines []string) (left []int, right []int, err error) {
	left = make([]int, 0)
	right = make([]int, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		nums := strings.Split(line, "   ")
		if len(nums) != 2 {
			fmt.Println("line: ", line)
			return left, right, fmt.Errorf("error read input")
		}

		numLeft, err := strconv.Atoi(nums[0])
		if err != nil {
			return left, right, fmt.Errorf("error strconv")
		}

		numRight, err := strconv.Atoi(nums[1])
		if err != nil {
			return left, right, fmt.Errorf("error strconv")
		}

		left = append(left, numLeft)
		right = append(right, numRight)
	}

	if len(left) != len(right) {
		return left, right, fmt.Errorf("error read input 2")
	}

	return left, right, nil
}

func CountDistanceApart(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	result := 0
	for i := 0; i < len(left); i++ {
		distance := Abs(left[i] - right[i])
		result += distance
	}

	return result
}

func CountSimilarity(left, right []int) int {
	rightMap := make(map[int]int)
	for _, num := range right {
		rightMap[num] += 1
	}

	result := 0
	for _, num := range left {
		rightCnt := rightMap[num]
		result += rightCnt * num
	}

	return result
}
