// Package day05
package day05

import (
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2025/utils"
)

func Day05_1(filename string) (result int) {
	il := NewIngredientsList(filename)
	result = il.CountFreshAvailable()
	return result
}

func Day05_2(filename string) (result int) {
	return result
}

type IngredientsList struct {
	Fresh       []int
	FreshRanges []FreshRange
	Available   []int
}

func NewIngredientsList(filename string) IngredientsList {
	il := IngredientsList{}
	readingRanges := true
	for line := range utils.InputCh(filename) {
		if line == "" {
			readingRanges = false
			continue
		}
		if readingRanges {
			// il.AddFresh(line)
			il.AddFreshRange(line)
			continue
		} else {
			il.AddAvailable(line)
		}
	}
	return il
}

// AddFresh -- OPTION 1 -- array of ints - does not work, out of memory
func (il *IngredientsList) AddFresh(line string) {
	ranges := strings.Split(line, "-")
	start, _ := strconv.Atoi(ranges[0])
	end, _ := strconv.Atoi(ranges[1])

	// update Fresh list to have at lease 'end'-size elements
	if len(il.Fresh) == 0 {
		il.Fresh = make([]int, end+1)
	}
	if len(il.Fresh) < end {
		newFresh := make([]int, end+1-len(il.Fresh))
		il.Fresh = append(il.Fresh, newFresh...)
	}
	for i := start; i <= end; i++ {
		il.Fresh[i]++
	}
}

// AddFreshRange -- OPTION 2 -- array of structs
func (il *IngredientsList) AddFreshRange(line string) {
	ranges := strings.Split(line, "-")
	start, _ := strconv.Atoi(ranges[0])
	end, _ := strconv.Atoi(ranges[1])

	if il.FreshRanges == nil {
		il.FreshRanges = make([]FreshRange, 0)
	}
	fr := FreshRange{From: start, To: end}
	// fmt.Println("Adding fresh range:", fr)
	il.FreshRanges = append(il.FreshRanges, fr)
}

func (il *IngredientsList) AddAvailable(line string) {
	available, _ := strconv.Atoi(line)
	il.Available = append(il.Available, available)
}

func (il IngredientsList) CountFreshAvailable() int {
	count := 0
	for _, v := range il.Available {
		// OPTION 1: array of ints
		// if v >= len(il.Fresh) {
		// 	continue
		// }
		// if il.Fresh[v] > 0 {
		// 	count++
		// }
		// OPTION 2: map
		for _, fr := range il.FreshRanges {
			// fmt.Println("Checking available ingredient", v, "in range", fr)
			if fr.IsIngredientFresh(v) {
				count++
				break
			}
		}
	}
	return count
}

type FreshRange struct {
	From int
	To   int
}

func (fr FreshRange) IsIngredientFresh(ingredient int) bool {
	return ingredient >= fr.From && ingredient <= fr.To
}
