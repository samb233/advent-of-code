package main

import (
	"fmt"
	"os"
	"strconv"
	// "strings"
)

func main() {
	input, err := os.ReadFile("3.in")
	if err != nil {
		panic(err)
	}

	// lines := strings.Split(string(input), "\n")
	result := MultiplyF(string(input))
	fmt.Println(result)
}

func Multiply(line string) int {
	lineRune := []rune(line)
	res := 0
	for i, c := range lineRune {
		if c == 'm' {
			isMul, n1, n2 := GetMultiplyNums(lineRune, i)
			if isMul {
				res += n1 * n2
			}
		}
	}

	return res
}

func MultiplyF(line string) int {
	lineRune := []rune(line)
	res := 0
	do := true
	for i, c := range lineRune {
		if c == 'd' {
			ifDo := DoOrDont(lineRune, i)
			if do && ifDo == -1 {
				do = false
			}

			if !do && ifDo == 1 {
				do = true
			}
		}

		if c == 'm' && do {
			isMul, n1, n2 := GetMultiplyNums(lineRune, i)
			if isMul {
				res += n1 * n2
			}
		}
	}

	return res
}

func DoOrDont(line []rune, index int) int {
	if index + 6 < len(line) && line[index] == 'd' && line[index+1] == 'o' && line[index+2] == 'n' &&
		line[index+3] == '\'' && line[index+4] == 't' && line[index+5] == '(' &&
		line[index+6] == ')' {
		return -1
	}

	if index + 3 < len(line) && line[index] == 'd' && line[index+1] == 'o' && line[index+2] == '(' &&
		line[index+3] == ')' {
		return 1
	}

	return 0

}

func GetMultiplyNums(line []rune, index int) (bool, int, int) {
	if index+4 > len(line) {
		return false, 0, 0
	}

	n1 := 0
	n2 := 0
	var err error

	if line[index] == 'm' && line[index+1] == 'u' && line[index+2] == 'l' && line[index+3] == '(' {
		num1 := make([]rune, 0)
		num2 := make([]rune, 0)
		num2Index := 0
		for i := index + 4; i < len(line); i++ {
			if line[i] == ',' {
				num2Index = i + 1
				break
			}

			if line[i] > '9' || line[i] < '0' {
				return false, 0, 0
			}

			num1 = append(num1, line[i])
		}

		if len(num1) == 0 {
			return false, 0, 0
		}

		for i := num2Index; i < len(line); i++ {
			if line[i] == ')' {
				break
			}

			if line[i] > '9' || line[i] < '0' {
				return false, 0, 0
			}

			num2 = append(num2, line[i])
		}

		if len(num2) == 0 {
			return false, 0, 0
		}

		n1, err = strconv.Atoi(string(num1))
		if err != nil {
			fmt.Println(num1)
			return false, 0, 0
		}

		n2, err = strconv.Atoi(string(num2))
		if err != nil {
			fmt.Println(num2)
			return false, 0, 0
		}
	}

	return true, n1, n2
}
