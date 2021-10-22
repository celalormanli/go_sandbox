package searching_algorithms

import "testing"

func TestLinearSearch(t *testing.T) {
	if LinearSearch([]int{10, 8, 48}, 5) == true {
		t.Errorf("LinearSearch should return false")
	}

	if LinearSearch([]int{10, 8, 48}, 48) == false {
		t.Errorf("LinearSearch should return true")
	}
}

func TestBinarySearch(t *testing.T) {
	if BinarySearch([]int{8, 13, 58, 96, 101}, 0, 4, 5) == true {
		t.Errorf("BinarySearch should return false")
	}

	if BinarySearch([]int{8, 13, 58, 96, 101}, 0, 4, 96) == false {
		t.Errorf("BinarySearch should return true")
	}
}
