package main

import (
	"testing"

	"github.com/rzabcio/adventofcode-2025/day01"
	"github.com/rzabcio/adventofcode-2025/day03"
	"github.com/rzabcio/adventofcode-2025/day04"
	"github.com/rzabcio/adventofcode-2025/day05"
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

func TestDay04(t *testing.T) {
	got, want := day04.Day04_1("input-files/day04-test1"), 13
	if got != want {
		t.Errorf("Day04_1(test1) = %d; want %d", got, want)
	}
	got, want = day04.Day04_2("input-files/day04-test1"), 43
	if got != want {
		t.Errorf("Day04_2(test1) = %d; want %d", got, want)
	}
}

func TestDay05(t *testing.T) {
	got, want := day05.Day05_1("input-files/day05-test1"), 3
	if got != want {
		t.Errorf("Day05_1(test1) = %d; want %d", got, want)
	}
	got, want = day05.Day05_2("input-files/day05-test1"), 14
	if got != want {
		t.Errorf("Day05_2(test1) = %d; want %d", got, want)
	}
}
