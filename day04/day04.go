// Package day04
package day04

import (
	"github.com/rzabcio/adventofcode-2025/utils"
)

func Day04_1(filename string) (result int) {
	store := NewPaperRollsStore(filename)
	result = store.CalculateAccessibleRolls()
	return result
}

func Day04_2(filename string) (result int) {
	store := NewPaperRollsStore(filename)
	count := store.CalculateAccessibleRolls()
	result += count

	for count > 0 {
		store.RemoveAccessibleRolls()
		count = store.CalculateAccessibleRolls()
		result += count
	}
	return result
}

type PaperRollsStore struct {
	Rolls      [][]int
	RollCounts [][]int
}

func NewPaperRollsStore(filename string) *PaperRollsStore {
	store := &PaperRollsStore{
		Rolls:      make([][]int, 0),
		RollCounts: make([][]int, 0),
	}
	x := 0
	for line := range utils.InputRowsWithRuneMapper(filename, map[rune]int{'@': 1, '.': 0}) {
		store.Rolls = append(store.Rolls, make([]int, len(line)))
		store.RollCounts = append(store.RollCounts, make([]int, len(line)))
		for y, val := range line {
			if val == 1 {
				store.Rolls[x][y]++
			}
		}
		x++
	}

	return store
}

func (store *PaperRollsStore) CalculateAccessibleRolls() (result int) {
	for x := range store.Rolls {
		for y := range store.Rolls[x] {
			if store.Rolls[x][y] == 1 {
				count := store.countNeighbours(x, y)
				store.RollCounts[x][y] = count
				if count >= 0 && count < 4 {
					result++
				}
			}
		}
	}
	return result
}

func (store *PaperRollsStore) countNeighbours(x, y int) int {
	if store.Rolls[x][y] == 0 {
		return -1
	}
	count := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if i < 0 || j < 0 || i >= len(store.RollCounts) || j >= len(store.RollCounts[0]) {
				continue
			}
			if store.Rolls[i][j] == 0 {
				continue
			}
			count++
		}
	}
	return count
}

func (store *PaperRollsStore) RemoveAccessibleRolls() {
	for x := range store.Rolls {
		for y := range store.Rolls[x] {
			if store.RollCounts[x][y] >= 0 && store.RollCounts[x][y] < 4 {
				store.Rolls[x][y] = 0
			}
		}
	}
}
