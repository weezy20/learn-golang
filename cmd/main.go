package main

import (
	"fmt"
	"learn-golang/pkg/sorter"
)

func main() {
	v := []int{213, 12, 3, 153, 1, 5, 124, 122}
	b_int := sorter.InsertionSort[int]{}
	b_int.Sort(v)
	b_char := sorter.InsertionSort[rune]{}
	// declare a array c of 10 runes that are unsorted
	c := []rune{'g', 'a', 'f', 'c', 'b', 'd', 'e', 'h', 'i', 'j'}
	// print a line for unsorted array , print each rune as characters
	fmt.Printf("Unsorted array : %c\n", c)
	b_char.Sort(c)
	fmt.Printf("Sorted array : %c\n", c)
}
