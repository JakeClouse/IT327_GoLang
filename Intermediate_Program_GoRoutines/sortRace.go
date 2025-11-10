package main
//Importing packages, fmt for printing, time for measuring execution time, sort for standard library sorting, and sync for goroutine synchronization
import (
	"fmt"
	"time"
	"sort"
	"sync"
)

var SortedArray []int

func main() {
	var wg sync.WaitGroup
	var backwardsArray []int
	backwardsArray = make([]int, 100000)
	SortedArray = make([]int, 100000)
	for i:=100000; i>0; i-- {
		backwardsArray[100000-i] = i
	}
	for i:=1; i<=100000; i++ {
		SortedArray[i-1] = i
	}
	wg.Add(4)
	go selectionSort(append([]int(nil), backwardsArray...), &wg)
	go insertionSort(append([]int(nil), backwardsArray...), &wg)
	go bubbleSort(append([]int(nil), backwardsArray...), &wg)
	go stdLibSort(append([]int(nil), backwardsArray...), &wg)
	wg.Wait()
}

func selectionSort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
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

func insertionSort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
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

func bubbleSort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
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

func stdLibSort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
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