// Package day03
package day03

import (
	"math"

	"github.com/rzabcio/adventofcode-2025/utils"
)

func Day03_1(filename string) (result int) {
	for bank := range utils.InputRowsAsInts(filename) {
		result += maxJoltage(bank, 2)
	}
	return result
}

func Day03_2(filename string) (result int) {
	for bank := range utils.InputRowsAsInts(filename) {
		result += maxJoltage(bank, 12)
	}
	return result
}

func maxJoltage(bank []int, length int) int {
	maxJoltage := 0
	maxBattery := 0
	lastIdx := 0
	idx := 0
	for i := length; i > 0; i-- {
		maxBattery, idx = utils.MaxIdx(bank[lastIdx : len(bank)-i+1])
		lastIdx += idx + 1
		maxJoltage += maxBattery * int(math.Pow(10, float64(i-1)))
	}
	return maxJoltage
}
