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

func BubbleSort(inputArray []int) []int {
	for x := 0; x < len(inputArray)-1; x++ {
		for y := 0; y < len(inputArray)-x-1; y++ {
			if inputArray[y] > inputArray[y+1] {
				temp := inputArray[y]
				inputArray[y] = inputArray[y+1]
				inputArray[y+1] = temp
			}
		}
	}
	return inputArray
}

func RecursiveBubbleSort(inputArray []int, counter int) []int {
	if counter == -1 {
		counter = len(inputArray)
	}
	if counter == 1 {
		return inputArray
	}
	for x := 0; x < counter-1; x++ {
		if inputArray[x] > inputArray[x+1] {
			temp := inputArray[x]
			inputArray[x] = inputArray[x+1]
			inputArray[x+1] = temp
		}
	}
	RecursiveBubbleSort(inputArray, counter-1)
	return inputArray
}
