package quick

import "fmt"

func QuickSort(unsort []int) {
	s := make([]int, len(unsort))
	copy(s, unsort)
	sortingQuick(s)
	fmt.Printf("Quick sorting: %v\n", s)
}

func sortingQuick(s []int) {
	if len(s) <= 1 {
		return
	}

	separation := partition(s)
	sortingQuick(s[:separation])
	sortingQuick(s[separation:])
}

func partition(s []int) int {
	middle := len(s) / 2
	pivot := s[middle]
	i, j := 0, len(s)-1
	for {
		for s[i] < pivot {
			i++
		}

		for s[j] > pivot {
			j--
		}

		if i >= j {
			return j
		}

		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
