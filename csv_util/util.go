package csv_util

import (
	"fmt"
	"strings"
)

/*
*   Function to print out to the Terminal in formated form
*/

func PrintRowsFancy(table [][]string, offset int, widthT int) {
	// calculate the max length of symbols for each column:
	maxLen := maxLenEachColumn(table)
	// if were longer than our screensize we cutoff:
	table, maxLen = cutoffToBigForScreen(table, offset, widthT)

	// print the (remaining)table out:
	for _, len := range maxLen {
		fmt.Printf("+%s", strings.Repeat("-", len))
	}
	fmt.Printf("+")
	for idx, row := range table {
		fmt.Printf("\n|")
		for i, str := range row {
			fmt.Printf("%s%s|", str, strings.Repeat(" ", maxLen[i]-len(str)))
		}
		fmt.Printf("  %v", idx+offset)
	}
	fmt.Printf("\n")
	for _, len := range maxLen {
		fmt.Printf("+%s", strings.Repeat("-", len))
	}
	fmt.Printf("+\n")
}

// max-count of chars of each column
func maxLenEachColumn(table [][]string) []int {
	maxLen := make([]int, len(table[0]))
	for _, row := range table {
		for i, str := range row {
			len := len(str)
			if len > maxLen[i] {
				maxLen[i] = len
			}
		}
	}
	return maxLen
}

// sum up how much space is needed
func sumLenAllWidth(maxLen []int, offset int) int {
	sumLen := countDigits(offset+len(maxLen)) + 2 // add index and static edges
	for _, len := range maxLen {
		sumLen += len + 1
	}
	return sumLen
}

// cutoff columns till dataset fits on terminal width
func cutoffToBigForScreen(table [][]string, offset, widthT int) ([][]string, []int) {
	maxLen := maxLenEachColumn(table)
	sumLen := sumLenAllWidth(maxLen, offset)

	if widthT <= sumLen {
		maxLen = maxLen[:len(maxLen)-1]
		for i, row := range table {
			table[i] = row[:len(maxLen)]
		}
		return cutoffToBigForScreen(table, offset, widthT)
	}
	return table, maxLen
}

// helper funcs
func countDigits(n int) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}
