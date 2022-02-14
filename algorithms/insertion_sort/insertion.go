package insertion

import "fmt"

func InsertionSort(unsort []int) {
	fmt.Printf("Insertion sorting: %v\n", sortingInsertion(unsort))
}

func sortingInsertion(orig []int) []int {
	s := make([]int, len(orig))
	copy(s, orig)
	for i := 1; i < len(s); i++ {
		current := s[i]
		for j := 0; j < i; j++ {
			if current <= s[j] {
				copy(s[j+1:i+1], s[j:i])
				s[j] = current
				break
			}
		}
	}
	return s
}
