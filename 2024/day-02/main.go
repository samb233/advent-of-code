package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("2.in")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	// result, err := CountSafeReport(lines)
	result, err := CountSafeReportWithDampener(lines)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func CountSafeReportWithDampener(lines []string) (int, error) {
	result := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		nums := strings.Split(line, " ")
		report := make([]int, 0)
		for _, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return -1, err
			}

			report = append(report, num)
		}

		isSafe, cnt := IsReportSafe(report)
		if isSafe {
			result++
			continue
		}

		newReport := make([]int, 0)
		for i, num := range report {
			if i != cnt {
				newReport = append(newReport, num)
			}
		}

		isSafe, _ = IsReportSafe(newReport)
		if isSafe {
			result++
			continue
		}

		newReport2 := make([]int, 0)
		for i, num := range report {
			if i != cnt-1 {
				newReport2 = append(newReport2, num)
			}
		}

		isSafe, _ = IsReportSafe(newReport2)
		if isSafe {
			result++
			continue
		}

		if cnt > 1 {
			newReport3 := make([]int, 0)
			for i, num := range report {
				if i != 0 {
					newReport3 = append(newReport3, num)
				}
			}

			isSafe, _ = IsReportSafe(newReport3)
			if isSafe {
				result ++
			}
		}
	}

	return result, nil
}

func IsReportSafe(report []int) (bool, int) {
	last := 0
	n := 1
	for i, num := range report {
		if i == 0 {
			last = num
			continue
		}

		diff := num - last
		if i == 1 {
			if diff > 0 {
				n = 1
			} else if diff < 0 {
				n = -1
			} else {
				return false, i
			}
		}

		if diff*n <= 0 {
			return false, i
		}

		if Abs(diff) > 3 {
			return false, i
		}

		last = num
	}

	return true, 0
}

func CountSafeReport(lines []string) (int, error) {
	result := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		nums := strings.Split(line, " ")

		last := 0
		n := 1
		isSafe := true

		for i, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return -1, err
			}

			if i == 0 {
				last = num
				continue
			}

			diff := num - last
			if i == 1 {
				if diff > 0 {
					n = 1
				} else if diff < 0 {
					n = -1
				} else {
					isSafe = false
					fmt.Println(diff)
					break
				}
			}

			if diff*n <= 0 {
				isSafe = false
				fmt.Println(diff)
				break
			}

			if Abs(diff) > 3 {
				isSafe = false
				fmt.Println(diff)
				break
			}

			last = num
		}

		if isSafe {
			result++
		} else {
			fmt.Println(line)
		}
	}

	return result, nil
}
