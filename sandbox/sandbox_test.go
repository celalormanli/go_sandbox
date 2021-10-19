package sandbox

import "testing"

func TestFibonacci(t *testing.T) {
	if Fibonacci(8) != 21 {
		t.Errorf("Fibonacci should be 21")
	}
	if Fibonacci(10) != 55 {
		t.Errorf("Fibonacci should be 55")
	}
}
