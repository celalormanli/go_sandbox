package sandbox

func Fibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return Fibonacci(x-1) + Fibonacci(x-2)
}
