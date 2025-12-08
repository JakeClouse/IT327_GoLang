// Maze Generator for IT327 Advanced program
// Author: Marcos Villalobos
// Date: 12/6/2025

// Modifications made by Jason Bliss using helper methods from Marcos to actually generate the maze, and modifying the implementation of Disjoint Set.

package MazeGeneration

import (
	"fmt"
	"math/rand"
)

type DisjointSet struct {
	numElements int
	theArray    []int
}

type Cell struct {
	Row, Col int
	Top      bool
	Right    bool
	Bottom   bool
	Left     bool
}

type Maze struct {
	Width, Height int
	Cells         [][]*Cell
	DS            *DisjointSet
}

func (m *Maze) removeWall(r1, c1, r2, c2 int) {
	cell := m.Cells[r1][c1]
	neighbor := m.Cells[r2][c2]

	if r1 == r2 {
		// horizontal neighbors
		if c1 < c2 {
			// cell is left of neighbor
			cell.Right = false
			neighbor.Left = false
		} else {
			cell.Left = false
			neighbor.Right = false
		}
	} else if c1 == c2 {
		// vertical neighbors
		if r1 < r2 {
			// cell is above neighbor
			cell.Bottom = false
			neighbor.Top = false
		} else {
			cell.Top = false
			neighbor.Bottom = false
		}
	}
}

func CreateMaze(w, h int) *Maze {
	m := NewMaze(w, h)
	populateMaze(m)
	return m
}

func NewMaze(w, h int) *Maze {
	cells := make([][]*Cell, h)
	for r := 0; r < h; r++ {
		cells[r] = make([]*Cell, w)
		for c := 0; c < w; c++ {
			cells[r][c] = &Cell{
				Row:    r,
				Col:    c,
				Top:    true,
				Right:  true,
				Bottom: true,
				Left:   true,
			}
		}
	}

	ds := newDisjointSet(w * h)

	return &Maze{
		Width:  w,
		Height: h,
		Cells:  cells,
		DS:     ds,
	}
}

func populateMaze(M *Maze) {

	w := M.Width
	h := M.Height

	//begin maze generation here
	var mazeComplete bool = false

	for !mazeComplete {
		//randomly select a cell
		randRow := rand.Intn(h)
		randCol := rand.Intn(w)
		cellIndex := randRow*w + randCol

		//randomly select a direction
		direction := rand.Intn(4) //0=up, 1=right, 2=down, 3=left

		switch direction {
		case 0: //up
			mazeComplete = up(cellIndex, M)
		case 1: //right
			mazeComplete = right(cellIndex, M)
		case 2: //down
			mazeComplete = down(cellIndex, M)
		case 3: //left
			mazeComplete = left(cellIndex, M)
		}

	}

}

func left(cellIndex int, M *Maze) bool {
	//out of bounds check
	neighborCellIndex := 0
	if (cellIndex % M.Width) == 0 {
		return right(cellIndex, M)
	} else { //actually going left
		neighborCellIndex = cellIndex - 1
		if M.DS.Find(cellIndex) != M.DS.Find(neighborCellIndex) {
			M.removeWall(cellIndex/M.Width, cellIndex%M.Width, neighborCellIndex/M.Width, neighborCellIndex%M.Width)
			return M.DS.Union(cellIndex, neighborCellIndex)
		}
	}

	return false

}

func right(cellIndex int, M *Maze) bool {
	//out of bounds check
	neighborCellIndex := 0
	if (cellIndex % M.Width) == M.Width-1 {
		return left(cellIndex, M)
	} else { //actually going right
		neighborCellIndex = cellIndex + 1
		if M.DS.Find(cellIndex) != M.DS.Find(neighborCellIndex) {
			M.removeWall(cellIndex/M.Width, cellIndex%M.Width, neighborCellIndex/M.Width, neighborCellIndex%M.Width)
			return M.DS.Union(cellIndex, neighborCellIndex)
		}
	}

	return false
}

func up(cellIndex int, M *Maze) bool {
	neighborCellIndex := 0
	if cellIndex < M.Width {
		return down(cellIndex, M)
	} else { //actually going up
		neighborCellIndex = cellIndex - M.Width
		if M.DS.Find(cellIndex) != M.DS.Find(neighborCellIndex) {
			M.removeWall(cellIndex/M.Width, cellIndex%M.Width, neighborCellIndex/M.Width, neighborCellIndex%M.Width)
			return M.DS.Union(cellIndex, neighborCellIndex)
		}
	}

	return false

}

func down(cellIndex int, M *Maze) bool {
	neighborCellIndex := 0
	if (cellIndex / M.Width) == M.Height-1 {
		return up(cellIndex, M)
	} else { //actually going down
		neighborCellIndex = cellIndex + M.Width
		if M.DS.Find(cellIndex) != M.DS.Find(neighborCellIndex) {
			M.removeWall(cellIndex/M.Width, cellIndex%M.Width, neighborCellIndex/M.Width, neighborCellIndex%M.Width)
			return M.DS.Union(cellIndex, neighborCellIndex)
		}
	}

	return false

}

func newDisjointSet(size int) *DisjointSet {
	ds := &DisjointSet{
		numElements: size,
		theArray:    make([]int, size),
	}
	for i := 0; i < size; i++ {
		ds.theArray[i] = -1
	}
	return ds
}

func (ds *DisjointSet) Find(index int) int {
	if ds.theArray[index] < 0 {
		return index
	} else {
		ds.theArray[index] = ds.Find(ds.theArray[index])
		return ds.theArray[index]
	}
}

func (ds *DisjointSet) Union(index1 int, index2 int) bool {
	var foundNegative bool = false
	var root1 int = ds.Find(index1)
	var root2 int = ds.Find(index2)
	if ds.theArray[root2] >= ds.theArray[root1] {
		ds.theArray[root1] += ds.theArray[root2]
		ds.theArray[root2] = root1
	} else {
		ds.theArray[root2] += ds.theArray[root1]
		ds.theArray[root1] = root2
	}
	for i := 0; i < ds.numElements; i++ {
		if ds.theArray[i] < 0 {
			if !foundNegative {
				foundNegative = true
			} else {
				return false
			}
		}
	}
	return true
}
func (ds *DisjointSet) PrintArrayValues() {
	for i := 0; i < ds.numElements; i++ {
		fmt.Print(ds.theArray[i], " ")
	}
	fmt.Print("\n")
}
