package sudoku

import "slices"

type IntSet []int

func (set *IntSet) Contains(searchTerm int) bool {
	return slices.Contains(*set, searchTerm)
}

func (set *IntSet) Add(element int) {
	if set.Contains(element) {
		return
	}

	*set = append(*set, element)
}
