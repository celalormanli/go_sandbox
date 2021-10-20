package sandbox

func Fibonacci(x int) int {
	if x <= 1 {
		return x
	}
	return Fibonacci(x-1) + Fibonacci(x-2)
}

func TowerOfHanoi(height int, from_pole *[]int, to_pole *[]int, with_pole *[]int) {
	if height >= 1 {
		TowerOfHanoi(height-1, from_pole, with_pole, to_pole)

		pop := (*from_pole)[(len(*from_pole) - 1)]
		*from_pole = (*from_pole)[:len(*from_pole)-1]
		*to_pole = append(*to_pole, pop)

		TowerOfHanoi(height-1, with_pole, to_pole, from_pole)
	}

}
