package main

import (
	"testing"

	"github.com/rzabcio/adventofcode-2025/day01"
	"github.com/rzabcio/adventofcode-2025/day03"
)

func TestDay01(t *testing.T) {
	got, want := day01.Day01_1("input-files/day01-test1"), 3
	if got != want {
		t.Errorf("Day01_1(test1) = %d; want %d", got, want)
	}
	// got, want = day01.Day01_2("input-files/day01-test1"), 6
	// if got != want {
	// 	t.Errorf("Day01_2(test1) = %d; want %d", got, want)
	// }
}

func TestDay03(t *testing.T) {
	got, want := day03.Day03_1("input-files/day03-test1"), 357
	if got != want {
		t.Errorf("Day03_1(test1) = %d; want %d", got, want)
	}
	got, want = day03.Day03_2("input-files/day03-test1"), 3121910778619
	if got != want {
		t.Errorf("Day03_2(test1) = %d; want %d", got, want)
	}
}
