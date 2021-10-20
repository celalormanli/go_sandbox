package sandbox

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	if Fibonacci(8) != 21 {
		t.Errorf("Fibonacci should be 21")
	}
	if Fibonacci(10) != 55 {
		t.Errorf("Fibonacci should be 55")
	}
}

func TestTowerOfHanoi(t *testing.T) {
	src := []int{4, 3, 2, 1}
	aux := []int{}
	dest := []int{}

	var s *[]int = &src
	var a *[]int = &aux
	var d *[]int = &dest
	TowerOfHanoi(4, s, d, a)
	if !reflect.DeepEqual(*d, []int{4, 3, 2, 1}) {
		t.Errorf("d should be [4, 3, 2, 1]")
	}
	if !reflect.DeepEqual(*a, []int{}) {
		t.Errorf("a should be []")
	}
	if !reflect.DeepEqual(*s, []int{}) {
		t.Errorf("s should be []")
	}
}
