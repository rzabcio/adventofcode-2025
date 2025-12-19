// Package day01
package day01

import (
	"fmt"
	"math"
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

func Day01_3(filename string) (result int) {
	current := 50
	for line := range utils.InputRows(filename) {
		parts := strings.SplitAfterN(line, "", 2)
		num, _ := strconv.Atoi(parts[1])
		fmt.Printf("%s%s: %d", parts[0], parts[1], current)
		if parts[0] == "R" {
			current += num
			fmt.Printf(" + %d -> %d", num, current)
		} else if parts[0] == "L" {
			current -= num
			fmt.Printf(" - %d -> %d", num, current)
		}
		for current >= 100 {
			current -= 100
			result++
			fmt.Printf(" - 100 -> %d [%d]", current, result)
		}
		for current < 0 {
			current += 100
			result++
			fmt.Printf(" + 100 -> %d [%d]", current, result)
		}
		fmt.Println()
	}
	return result
}

func Day01_2(filename string) (result int) {
	current := 50
	for line := range utils.InputRows(filename) {
		parts := strings.SplitAfterN(line, "", 2)
		num, _ := strconv.Atoi(parts[1])
		fmt.Printf("%s%s: %d", parts[0], parts[1], current)
		if parts[0] == "R" {
			if current == 0 {
				result++
			}
			current += num
			fmt.Printf(" + %d -> %d", num, current)
		} else if parts[0] == "L" {
			current -= num
			fmt.Printf(" - %d -> %d", num, current)
		}
		currentMod := current % 100
		currentDiv := current / 100
		fmt.Printf(" => mod: %d, div: %d", currentMod, currentDiv)
		if currentMod < 0 {
			currentMod += 100
			currentDiv--
			fmt.Printf(" => mod: %d, div: %d", currentMod, currentDiv)
		}
		current = currentMod
		result += int(math.Abs(float64(currentDiv)))
		fmt.Println()
	}
	return result
}
