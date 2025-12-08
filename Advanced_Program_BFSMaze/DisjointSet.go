package main

import "fmt"

type DisjointSet struct {
	numElements int
	theArray   []int
}

func NewDisjointSet(size int) *DisjointSet {
	ds := &DisjointSet{
		numElements: size,
		theArray:   make([]int, size),
	}
	for i:=0 ; i < size; i++ {
		ds.theArray[i] = -1
	}
	return ds
}

func (ds *DisjointSet) Find(index int) int {
	if(ds.theArray[index] < 0){
		return index
	} else{
		ds.theArray[ index ] = ds.Find( ds.theArray[ index ] )
		return ds.theArray[ index ]
	}
}

func (ds *DisjointSet) Union(index1 int, index2 int) bool {
	var foundNegative bool = false
	var root1 int = ds.Find(index1)
	var root2 int = ds.Find(index2)
	if(ds.theArray[root2]>=ds.theArray[root1]){
		ds.theArray[root1] += ds.theArray[root2]
		ds.theArray[root2] = root1
	} else {
		ds.theArray[root2] += ds.theArray[root1]
		ds.theArray[root1] = root2	
	}
	for i:=0; i<ds.numElements; i++{
		if (ds.theArray[i] < 0){
			if(!foundNegative){
				foundNegative = true
			}else{
				return false
			}
		}
	}
	return true
}
func (ds *DisjointSet) PrintArrayValues() {
	for i:= 0; i<ds.numElements; i++ {
		fmt.Print(ds.theArray[i]," ")
	}
	fmt.Print("\n")
}
