package sorting_algorithms

func SelectionSort(inputArray []int) []int {
	for x := 0; x < len(inputArray); x++ {
		minNumId := x
		for y := x; y < len(inputArray); y++ {
			if inputArray[minNumId] > inputArray[y] {
				minNumId = y
			}
		}
		temp := inputArray[x]
		inputArray[x] = inputArray[minNumId]
		inputArray[minNumId] = temp
	}
	return inputArray
}
