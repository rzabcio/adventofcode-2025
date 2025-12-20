// Package day02
package day02

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func Day02_1(filename string) (result int) {
	for r := range nextRange(filename) {
		fmt.Println("Range:", r)
		for id := range r.nextID() {
			l := len(fmt.Sprintf("%d", id))
			// if lenght is odd, skip (cannot split in half, to compare numbers
			if l%2 != 0 {
				continue
			}
			mod := int(math.Pow(10, float64(l/2)))
			firstHalf := id / mod
			secondHalf := id % mod
			if firstHalf == secondHalf {
				result += id
			}
		}
	}
	return result
}

func Day02_2(filename string) (result int) {
	return result
}

type Range struct {
	start int
	end   int
}

func (r Range) nextID() (id chan int) {
	id = make(chan int)
	go func() {
		defer close(id)
		for i := r.start; i <= r.end; i++ {
			id <- i
		}
	}()
	return id
}

func nextRange(filename string) (ch chan Range) {
	ch = make(chan Range)
	go func() {
		defer close(ch)
		file, err := os.Open(filename)
		if err != nil {
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println("Line:", scanner.Text())
			for rStr := range strings.SplitSeq(scanner.Text(), ",") {
				var r Range
				fmt.Sscanf(rStr, "%d-%d", &r.start, &r.end)
				ch <- r
			}
		}
	}()
	return ch
}
