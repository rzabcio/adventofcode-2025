// Package utils provides utility functions for file input and common operations.
package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func InputCh(filename string) (ch chan string) {
	return InputRows(filename)
}

func InputRows(filename string) (ch chan string) {
	ch = make(chan string)
	go func() {
		file, err := os.Open(filename)
		if err != nil {
			close(ch)
			return
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()
	return ch
}

func InputRowsWithRuneMapper(filename string, mapper map[rune]int) (ch chan []int) {
	ch = make(chan []int)
	go func() {
		file, err := os.Open(filename)
		if err != nil {
			close(ch)
			return
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			row := []int{}
			for _, c := range scanner.Text() {
				row = append(row, mapper[c])
			}
			ch <- row
		}
		close(ch)
	}()
	return ch
}

func InputRowsAsInts(filename string) (ch chan []int) {
	ch = make(chan []int)
	go func() {
		file, err := os.Open(filename)
		if err != nil {
			close(ch)
			return
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			row := rowAsInts(scanner.Text())
			ch <- row
		}
		close(ch)
	}()
	return ch
}

func rowAsInts(row string) (result []int) {
	for _, s := range strings.Fields(row) {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		result = append(result, i)
	}
	return result
}

func InputCols(filename string) (result []string) {
	matrix := [][]string{}
	for row := range InputRows(filename) {
		for i, c := range row {
			if len(matrix) <= i {
				matrix = append(matrix, []string{})
			}
			matrix[i] = append(matrix[i], string(c))
		}
	}
	for _, col := range matrix {
		result = append(result, strings.Join(col, ""))
	}
	return result
}

func MinMax(array []int) (int, int) {
	max := array[0]
	min := array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func Min(array []int) int {
	min, _ := MinMax(array)
	return min
}

func Max(array []int) int {
	_, max := MinMax(array)
	return max
}

// func ReverseString(s string) string {
// 	r := []rune(s)
// 	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return string(r)
// }

func ReverseString(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
