// Package day03
package day03

import (
	"github.com/rzabcio/adventofcode-2025/utils"
)

func Day03_1(filename string) (result int) {
	for bank := range utils.InputRowsAsInts(filename) {
		maxJoltage := 0
		for i := 0; i < len(bank)-1; i++ {
			battery := bank[i]
			for j := i + 1; j < len(bank); j++ {
				otherBattery := bank[j]
				joltage := battery*10 + otherBattery
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}
		}
		result += maxJoltage
	}
	return result
}

func Day03_2(filename string) (result int) {
	return result
}
