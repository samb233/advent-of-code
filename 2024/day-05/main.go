package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("5.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	updateOrder, err := LoadUpdateOrder(lines)
	updateLines := GetUpdateLines(lines)

	res := 0
	for _, line := range updateLines {
		if len(line) == 0 {
			continue
		}

		n, err := CheckUpdatePart2(line, updateOrder)
		if err != nil {
			panic(err)
		}

		// fmt.Println(n)
		res += n
	}

	fmt.Println(res)
}

func LoadUpdateOrder(lines []string) (map[int][]int, error) {
	updateOrder := make(map[int][]int)
	for _, line := range lines {
		if len(line) == 0 {
			break
		}

		numStrs := strings.Split(line, "|")
		if len(numStrs) != 2 {
			return nil, fmt.Errorf("parse error: ", line)
		}

		num1Str := numStrs[0]
		num2Str := numStrs[1]

		num1, err := strconv.Atoi(num1Str)
		if err != nil {
			return nil, err
		}

		num2, err := strconv.Atoi(num2Str)
		if err != nil {
			return nil, err
		}

		order, ok := updateOrder[num1]
		if !ok {
			order = make([]int, 0)
		}

		order = append(order, num2)
		updateOrder[num1] = order
	}

	return updateOrder, nil

}

func GetUpdateLines(lines []string) []string {
	for i, line := range lines {
		if len(line) == 0 {
			return lines[i+1:]
		}
	}

	return lines
}

func NumInSlice(num int, slice []int) bool {
	for _, s := range slice {
		if s == num {
			return true
		}
	}

	return false
}

// part 1
func CheckUpdate(line string, updateOrder map[int][]int) (int, error) {
	res := 0
	orderOk := true

	numStrs := strings.Split(line, ",")
	nums := make([]int, 0)
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, err
		}
		nums = append(nums, num)
	}

	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		order, ok := updateOrder[num]
		if !ok {
			continue
		}

		for j := 0; j < i; j++ {
			if NumInSlice(nums[j], order) {
				orderOk = false
				break
			}
		}
	}

	if orderOk {
		res = nums[len(nums)/2]
	}

	return res, nil
}

// part 2
func CheckUpdatePart2(line string, updateOrder map[int][]int) (int, error) {
	numStrs := strings.Split(line, ",")
	nums := make([]int, 0)
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, err
		}
		nums = append(nums, num)
	}

	return CheckUpdateOnlyError(nums, updateOrder, false)
}


func CheckUpdateOnlyError(nums []int, updateOrder map[int][]int, getNum bool) (int, error) {
	orderOk := true

	x := 0
	y := 0

	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		order, ok := updateOrder[num]
		if !ok {
			continue
		}

		for j := 0; j < i; j++ {
			if NumInSlice(nums[j], order) {
				orderOk = false
				x = i
				y = j
				break
			}
		}
	}

	if orderOk {
		if getNum {
			return nums[len(nums)/2], nil
		} else {
			return 0, nil
		}
	}

	nums[x], nums[y] = nums[y], nums[x]
	return CheckUpdateOnlyError(nums, updateOrder, true)
}
