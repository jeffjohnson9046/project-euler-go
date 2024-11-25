package types

type IntSet []int

func (set *IntSet) Contains(searchTerm int) bool {
	for _, value := range *set {
		if value == searchTerm {
			return true
		}
	}

	return false
}

func (set *IntSet) Add(element int) {
	if set.Contains(element) {
		return
	}

	*set = append(*set, element)
}
