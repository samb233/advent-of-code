package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("11.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	line := lines[0]

	nums := strings.Split(line, " ")
	// for i := 0; i < 75; i++ {
	// 	fmt.Println(i)

	// 	nums, err = BlinkPart1(nums)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	res, _ := BlinkPart2(nums, 75)
	fmt.Println(res)
}

func BlinkPart1(nums []string) ([]string, error) {
	newNums := make([]string, 0)
	for _, num := range nums {
		newNum, err := Blink(num)
		if err != nil {
			return nil, err
		}

		newNums = append(newNums, newNum...)
	}

	return newNums, nil
}

func Blink(num string) ([]string, error) {
	if num == "0" {
		return []string{"1"}, nil
	}

	if len(num)%2 == 0 {
		n := len(num) / 2
		num1, num2 := num[0:n], num[n:]

		num1Int, err := strconv.Atoi(num1)
		if err != nil {
			return nil, err
		}

		num2Int, err := strconv.Atoi(num2)
		if err != nil {
			return nil, err
		}

		return []string{
			fmt.Sprintf("%d", num1Int),
			fmt.Sprintf("%d", num2Int),
		}, nil
	}

	numInt, err := strconv.Atoi(num)
	if err != nil {
		return nil, err
	}

	return []string{
		fmt.Sprintf("%d", numInt*2024),
	}, nil
}

func BlinkPart2(nums []string, t int) (int, error) {
	type Key struct {
		Stone   string
		StepNum int
	}

	m := make(map[Key]int)
	// x: num
	// t: step
	var DPBlink func(Key) int
	DPBlink = func(k Key) int {
		t, ok := m[k]
		if ok {
			return t
		}

		if k.StepNum == 0 {
			return 1
		}

		res := 0
		if k.Stone == "0" {
			res = DPBlink(Key{"1", k.StepNum - 1})
		} else if len(k.Stone)%2 == 0 {
			n := len(k.Stone) / 2
			num1, num2 := k.Stone[0:n], k.Stone[n:]
			num1Int, _ := strconv.Atoi(num1)
			num2Int, _ := strconv.Atoi(num2)

			k1 := Key{
				Stone: fmt.Sprintf("%d", num1Int),
				StepNum: k.StepNum - 1,
			}

			k2 := Key{
				Stone: fmt.Sprintf("%d", num2Int),
				StepNum: k.StepNum - 1,
			}
			res = DPBlink(k1) + DPBlink(k2)

		} else {
			num, _ := strconv.Atoi(k.Stone)
			k1 := Key{
				Stone: fmt.Sprintf("%d", num*2024),
				StepNum: k.StepNum - 1,
			}

			res = DPBlink(k1)
		}

		m[k] = res
		return res
	}

	sum := 0
	for _, num := range nums {
		k := Key{
			Stone: num,
			StepNum: t,
		}

		sum += DPBlink(k)
	}

	return sum, nil
}
