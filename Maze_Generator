// Maze Generator for IT327 Advanced program
// Author: Marcos Villalobos
// Date: 12/6/2025

import (
	
	"math/rand"
	"time"
)

//
type DisjointSet struct {
	parent []int
	rank   []int
}

//
type Cell struct {
	Row, Col int
	Top 	 bool
	Right    bool
	Bottom   bool
	Left 	 bool
}

//
type Maze struct {
	Width, Height int
	Cells         [][]*Cell
	DS            *DisjointSet
}

//
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

// 
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

	ds := NewDisjointSet(w * h)

	return &Maze{
		Width:  w,
		Height: h,
		Cells:  cells,
		DS:     ds,
	}
}

func NewDisjointSet(n int) *DisjointSet {
	
	for i := 0; i < n; i++ {
		
	}
	return ds
}

// 
func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] != x {
		ds.parent[x] = ds.Find(ds.parent[x]) // path compression
	}
	return ds.parent[x]
}

// 
func (ds *DisjointSet) Union(x, y int) bool {
	rx, ry := ds.Find(x), ds.Find(y)
	if rx == ry {
		return false
	}
	if ds.rank[rx] < ds.rank[ry] {
		ds.parent[rx] = ry
	} else if ds.rank[rx] > ds.rank[ry] {
		ds.parent[ry] = rx
	} else {
		ds.parent[ry] = rx
		ds.rank[rx]++
	}
	return true
}
