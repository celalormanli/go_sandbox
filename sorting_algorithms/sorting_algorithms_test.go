package sorting_algorithms

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	if !reflect.DeepEqual(SelectionSort([]int{10, 8, 48}), []int{8, 10, 48}) {
		t.Errorf("Result should be {8, 10, 48}")
	}
}

func TestBubbleSort(t *testing.T) {
	if !reflect.DeepEqual(BubbleSort([]int{10, 8, 48}), []int{8, 10, 48}) {
		t.Errorf("Result should be {8, 10, 48}")
	}
}

func TestRecursiveBubbleSort(t *testing.T) {
	if !reflect.DeepEqual(RecursiveBubbleSort([]int{10, 8, 48, 1}, -1), []int{1, 8, 10, 48}) {
		t.Errorf("Result should be {1, 8, 10, 48}")
	}
}
