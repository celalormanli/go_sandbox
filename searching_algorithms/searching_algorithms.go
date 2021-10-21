package searching_algorithms

func LinearSearch(inputArray []int, searchedElement int) bool {
	for _, i := range inputArray {
		if i == searchedElement {
			return true
		}
	}
	return false
}
