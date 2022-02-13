package main

import "fmt"

func main() {
	var orig []int = []int{3, 7, 8, 2, 4, 6, 1, 5}
	fmt.Println("Start!")
	fmt.Printf("Original slice: %v\n", orig)
	mergeSort(orig)
}

func mergeSort(unsort []int) {
	fmt.Printf("Merge sorting: %v\n", sortingMerge(unsort))
}

func sortingMerge(s []int) []int {
	if len(s) == 1 {
		return s
	}

	middle := len(s) / 2
	left := sortingMerge(s[:middle])
	right := sortingMerge(s[middle:])
	return merge(left, right)
}

func merge(l, r []int) (result []int) {
	var lind, rind int
	for i := 0; i < len(l)+len(r); i++ {

		if lind == len(l) {
			result = append(result, r[rind])
			rind++
		} else if rind == len(r) {
			result = append(result, l[lind])
			lind++
		} else if l[lind] >= r[rind] {
			result = append(result, r[rind])
			rind++
		} else {
			result = append(result, l[lind])
			lind++
		}

	}
	return result
}
