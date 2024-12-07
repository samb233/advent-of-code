package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("7.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	reses, numses, err := LoadInput(lines)
	if err != nil {
		panic(err)
	}

	res := 0
	for i := 0; i < len(reses); i++ {
		if EvaluatePart2(reses[i], numses[i]) {
			res += reses[i]
			// fmt.Println(reses[i])
		}
	}

	fmt.Println(res)
}

func LoadInput(lines []string) ([]int, [][]int, error) {
	reses := make([]int, 0)
	numses := make([][]int, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		nums := make([]int, 0)
		line = strings.ReplaceAll(line, ":", "")
		s := strings.Split(line, " ")

		for i, numStr := range s {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, nil, err
			}

			if i == 0 {
				reses = append(reses, num)
			} else {
				nums = append(nums, num)
			}
		}

		numses = append(numses, nums)
	}

	return reses, numses, nil
}

// part 1
func Evaluate(res int, nums []int) bool {
	if len(nums) == 1 {
		return res == nums[0]
	}

	numsCpy := make([]int, len(nums) - 1)
	copy(numsCpy, nums[1:])

	mul := nums[0] * nums[1]
	plus := nums[0] + nums[1]

	numsCpy[0] = mul
	if Evaluate(res, numsCpy) {
		return true
	}

	numsCpy[0] = plus
	return Evaluate(res, numsCpy)
}

// part 2
func EvaluatePart2(res int, nums []int) bool {
	if len(nums) == 1 {
		return res == nums[0]
	}

	numsCpy := make([]int, len(nums) - 1)
	copy(numsCpy, nums[1:])

	mul := nums[0] * nums[1]
	plus := nums[0] + nums[1]
	conStr := fmt.Sprintf("%d%d", nums[0], nums[1])
	con, _ := strconv.Atoi(conStr)

	numsCpy[0] = mul
	if EvaluatePart2(res, numsCpy) {
		return true
	}

	numsCpy[0] = plus
	if EvaluatePart2(res, numsCpy) {
		return true
	}

	numsCpy[0] = con
	return EvaluatePart2(res, numsCpy)
}
