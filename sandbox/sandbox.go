package sandbox

func Fibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return Fibonacci(x-1) + Fibonacci(x-2)
}

func TowerOfHanoi(height int, fromPole *[]int, toPole *[]int, withPole *[]int) {
	if height >= 1 {
		TowerOfHanoi(height-1, fromPole, withPole, toPole)

		pop := (*fromPole)[(len(*fromPole) - 1)]
		*fromPole = (*fromPole)[:len(*fromPole)-1]
		*toPole = append(*toPole, pop)

		TowerOfHanoi(height-1, withPole, toPole, fromPole)
	}

}
