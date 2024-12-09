package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DiskUnit struct {
	ID     int
	Length int
	Free   bool
}

func main() {
	input, err := os.ReadFile("9.in")
	if err != nil {
		panic(err)
	}

	line := string(input)
	line = strings.ReplaceAll(line, "\n", "")
	units, err := LoadDiskUnit(line)
	if err != nil {
		panic(err)
	}

	units = MoveFile(units, 0)
	res := ChecksumUnits(units)
	fmt.Println(res)
}

func LoadDiskUnit(line string) ([]DiskUnit, error) {
	r := []rune(line)
	units := make([]DiskUnit, 0)

	diskID := 0
	for i, ch := range r {
		num, err := strconv.Atoi(string(ch))
		if err != nil {
			return nil, err
		}

		var unit DiskUnit
		if i%2 != 0 {
			unit = DiskUnit{
				ID:     0,
				Free:   true,
				Length: num,
			}
		} else {
			unit = DiskUnit{
				ID:     diskID,
				Free:   false,
				Length: num,
			}

			diskID++
		}

		units = append(units, unit)
	}

	return units, nil
}

func Checksum(units []DiskUnit) int {
	checksum := 0
	mul := 0
	for i := 0; i < len(units); i++ {
		if !units[i].Free {
			checksum += UnitChecksum(units[i], mul)
			mul += units[i].Length

			continue
		}

		for j := len(units) - 1; j > i; j-- {
			if units[j].Free {
				continue
			}

			if units[j].Length >= units[i].Length {
				units[j].Length -= units[i].Length
				if units[j].Length == 0 {
					units[j].Free = true
				}

				units[i].Free = false
				units[i].ID = units[j].ID

				checksum += UnitChecksum(units[i], mul)
				mul += units[i].Length
			} else {
				units[j].Free = true

				newUnit := DiskUnit{
					ID:     units[j].ID,
					Length: units[j].Length,
					Free:   false,
				}

				units[i].Length -= units[j].Length

				checksum += UnitChecksum(newUnit, mul)
				mul += newUnit.Length
				i--
			}
			break
		}
	}

	return checksum
}

func UnitChecksum(unit DiskUnit, mul int) int {
	checksum := 0
	for i := 0; i < unit.Length; i++ {
		checksum += unit.ID * (mul + i)
		fmt.Println(unit.ID, " x ", mul + i)
	}


	return checksum
}

// part 2
func MoveFile(units []DiskUnit, rightIndex int) []DiskUnit {
	freeIndex := 0
	fileIndex := 0

LOOP:
	for j := len(units) - 1 - rightIndex; j >= 0; j-- {
		if units[j].Free {
			continue
		}

		fileIndex = j
		for i := 0; i < j; i++ {
			if !units[i].Free {
				continue
			}

			if units[i].Length >= units[j].Length {
				freeIndex = i
				break LOOP
			}
		}
	}

	if fileIndex == 0 {
		return units
	}

	if freeIndex == 0 {
		return MoveFile(units, rightIndex+1)
	}

	newUnit := DiskUnit{
		ID:     0,
		Length: units[freeIndex].Length - units[fileIndex].Length,
		Free:   true,
	}

	units[freeIndex] = units[fileIndex]
	units[fileIndex].ID = 0
	units[fileIndex].Free = true

	if newUnit.Length == 0 {
		return MoveFile(units, rightIndex + 1)
	}

	newUnits := make([]DiskUnit, 0)
	for i, unit := range units {
		newUnits = append(newUnits, unit)
		if i == freeIndex {
			newUnits = append(newUnits, newUnit)
		}
	}

	return MoveFile(newUnits, rightIndex + 1)
}

func ChecksumUnits(units []DiskUnit) int {
	checksum := 0
	mul := 0
	for _, unit := range units {
		if unit.Free {
			mul += unit.Length
			continue
		}

		checksum += UnitChecksum(unit, mul)
		mul += unit.Length
	}

	return checksum
}
