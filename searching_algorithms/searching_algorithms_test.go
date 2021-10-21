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
