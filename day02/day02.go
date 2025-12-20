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

// Day02_2 - the idea here is to repeatedly modulo divide by powerst of 10 - in the first divide
// pattern will be created and in the next divides we will search for this pattern
// if not, then the parttern does not repeat, but if found, move to next power of 10
// if after all powers of 10 pattern is not found, then the number does not have repeating parts
// works fine for test data, but for the final data the result is too high...
func Day02_2(filename string) (result int) {
	for r := range nextRange(filename) {
	nextID:
		for id := range r.nextID() {
			l := len(fmt.Sprintf("%d", id))
		nextPattern:
			for i := 1; i <= l/2; i++ {
				mod := int(math.Pow(10, float64(i)))
				pattern := (id % mod)
				rest := (id / mod)
				for rest > 0 {
					searchPart := rest % mod
					if searchPart != pattern {
						continue nextPattern
					}
					rest /= mod
				}
				result += id
				continue nextID
			}
		}
	}
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
