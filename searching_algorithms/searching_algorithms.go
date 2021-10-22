package searching_algorithms

func LinearSearch(inputArray []int, searchedElement int) bool {
	for _, i := range inputArray {
		if i == searchedElement {
			return true
		}
	}
	return false
}

func BinarySearch(inputArray []int, le int, ri int, x int) bool {
	for le <= ri {
		var mid int
		mid = le + (ri-1)/2
		if inputArray[mid] == x {
			return true
		} else if inputArray[mid] < x {
			le = mid + 1
		} else {
			ri = mid - 1
		}
	}
	return false
}
