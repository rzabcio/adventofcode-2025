// Package day01
package day01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2025/utils"
)

func Day01_1(filename string) (result int) {
	current := 50
	for line := range utils.InputRows(filename) {
		parts := strings.SplitAfterN(line, "", 2)
		num, _ := strconv.Atoi(parts[1])
		if parts[0] == "R" {
			current += num
			for current >= 100 {
				current -= 100
			}
		} else if parts[0] == "L" {
			current -= num
			for current < 0 {
				current += 100
			}
		}
		if current == 0 {
			result++
		}
	}
	return result
}

func Day01_2(filename string) (result int) {
	for line := range utils.InputRows(filename) {
		// split line to letter and number, example: "R14" -> "R" and "14"
		parts := strings.SplitAfterN(line, "", 2)
		num, _ := strconv.Atoi(parts[1])
		fmt.Println(parts[0], num)
	}

	return result
}
