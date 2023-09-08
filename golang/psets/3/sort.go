package main

import "fmt"

func main() {
	arr := []int{6, 3, 10, 9, 1, 5, 2, 7, 4}
	fmt.Println("unsorted :", arr)
	mergeSort(arr)
	fmt.Println("sorted :", arr)

}

// merge sort
func mergeSort(arr []int) {

	if len(arr) <= 1 {
		return
	} else {
		mid := len(arr) / 2
		leftArr := make([]int, mid)
		rightArr := make([]int, len(arr)-mid)

		copy(leftArr, arr[:mid])
		copy(rightArr, arr[mid:])

		mergeSort(leftArr)
		mergeSort(rightArr)

		merge(arr, leftArr, rightArr)
	}
}

func merge(arr, left, right []int) {
	i, leftIndex, rightIndex := 0, 0, 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			arr[i] = left[leftIndex]
			leftIndex++
		} else {
			arr[i] = right[rightIndex]
			rightIndex++
		}
		i++
	}

	for leftIndex < len(left) {
		arr[i] = left[leftIndex]
		leftIndex++
		i++
	}

	for rightIndex < len(right) {
		arr[i] = right[rightIndex]
		rightIndex++
		i++
	}
}
