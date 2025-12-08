package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	// recursively split and merge left and right halves
	left := mergeSort(arr[:mid])  // arr[:mid] gets the left half
	right := mergeSort(arr[mid:]) // arr[mid:] gets the right half
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right)) // set cap of result slice to legnth of left + right
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
