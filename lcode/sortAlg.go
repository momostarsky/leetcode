package leetcode

/*
说明：
算法：  https://github.com/hustcc/JS-Sorting-Algorithm
*/
func heapSort(arr []int) []int {
	arrLen := len(arr)
	buildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		swapElement(arr, 0, i)
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
	return arr
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func heapify(arr []int, i, arrLen int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swapElement(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}

func swapElement(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func HeapSort(arr []int) (res []int) {

	res = heapSort(arr)
	return
}
