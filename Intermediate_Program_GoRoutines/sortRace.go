package main

import (
	"fmt"
	"time"
	"sort"
)

var SortedArray []int

func main() {
	var backwardsArray []int
	backwardsArray = make([]int, 10000)
	SortedArray = make([]int, 10000)
	for i:=10000; i>0; i-- {
		backwardsArray[10000-i] = i
	}
	for i:=1; i<=10000; i++ {
		SortedArray[i-1] = i
	} 
	go selectionSort(backwardsArray)
	go insertionSort(backwardsArray)
	go bubbleSort(backwardsArray)
	go stdLibSort(backwardsArray)
	time.Sleep(5 * time.Second)
}

func selectionSort(arr []int) {
	startTime := time.Now()
	for i:=0;i<(len(arr)-1);i++ {
		minIndex := i
		for j:=i;j<(len(arr));j++ {
			if arr[j]<arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	endTime := time.Since(startTime)
	if(checkSorted(arr)){
		fmt.Println("Selection Sort Completed in ", endTime)
	}else{
		fmt.Println("Selection Sort Failed")
	}
}

func insertionSort(arr []int) {
	startTime := time.Now()
	for i:=1;i<len(arr);i++ {
		currVal := arr[i]
		j := i - 1
		for j>=0 && currVal<arr[j]{
			arr[j+1] = arr[j]
			j-=1
		}
		arr[j+1] = currVal
	}
	endTime := time.Since(startTime)
	if(checkSorted(arr)){
		fmt.Println("Insertion Sort Completed in ", endTime)
	}else{
		fmt.Println("Insertion Sort Failed")
	}
}

func bubbleSort(arr []int) {
	startTime := time.Now()
	for i:=0;i<len(arr);i++{
		valSwapped := false
		for j:=0;j<len(arr)-i-1;j++{
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				valSwapped = true
			}
		}
		if (valSwapped == false){
			break
		}
	}
	endTime := time.Since(startTime)
	if(checkSorted(arr)){
		fmt.Println("Bubble Sort Completed in ", endTime)
	}else{
		fmt.Println("Bubble Sort Failed")
	}
}

func stdLibSort(arr []int) {
	startTime := time.Now()
	sort.Ints(arr)
	endTime := time.Since(startTime)
	if(checkSorted(arr)){
		fmt.Println("Std Sort Completed in ", endTime)
	}else{
		fmt.Println("Std Sort Failed")
	}
}

func checkSorted(arr []int) bool {
	for i:=0;i<len(arr);i++ {
		if arr[i] != SortedArray[i] {
			return false
		}
	}
	return true
}