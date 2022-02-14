package main

import (
	"fmt"

	ins "github.com/aanikolaev5/test-and-pets/algorithms/insertion_sort"
	m "github.com/aanikolaev5/test-and-pets/algorithms/merge_sort"
)

func main() {
	var orig []int = []int{9, 3, 7, 8, 2, 4, 6, 1, 5}
	fmt.Printf("Original slice: %v\n", orig)
	m.MergeSort(orig)
	ins.InsertionSort(orig)
}
