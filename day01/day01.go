// Package day01
package day01

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/rzabcio/adventofcode-2025/utils"
)

func Day01_1(filename string) (result int) {
	l, r := readFile(filename)

	fmt.Println(l)
	fmt.Println(r)

	return result
}

func Day01_2(filename string) (result int) {
	l, r := readFile(filename)

	fmt.Println(l)
	fmt.Println(r)

	return result
}

func readFile(filename string) (lList, rList []int) {
	for line := range utils.InputRows(filename) {
		s := strings.Split(line, "   ")
		lInt, _ := strconv.Atoi(s[0])
		rInt, _ := strconv.Atoi(s[1])
		lList = append(lList, lInt)
		rList = append(rList, rInt)
	}
	sort.Ints(lList)
	sort.Ints(rList)
	return lList, rList
}
