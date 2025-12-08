//Parallel BFS for IT327 Advanced program
// Author: Jake CLouse
// Date: 12/6/2025

//Edited by Brady McCue to have paramaters match with generated maze

package main

import (
	"sync"
	"sync/atomic"
)

type Pair struct {
	Row, Col int
}

func concurrentCheckNeighbor(cell Pair, nextFrontier *[]Pair, visited *[][]Pair, mu_visited *sync.Mutex, mu_frontier *sync.Mutex, parentCell Pair, targetRow int, targetCol int, success *int32) {
	if cell.Row == targetRow && cell.Col == targetCol {
		atomic.AddInt32(success, 1)
	}
	mu_visited.Lock()
	if (*visited)[cell.Row][cell.Col] == (Pair{-1, -1}) {
		(*visited)[cell.Row][cell.Col] = Pair{Row: parentCell.Row, Col: parentCell.Col}
		mu_visited.Unlock()
		mu_frontier.Lock()
		*nextFrontier = append(*nextFrontier, Pair{Row: cell.Row, Col: cell.Col})
		mu_frontier.Unlock()
	} else {
		mu_visited.Unlock()
	}

}

func concurrentOperation(maze [][]*Cell, cell Pair, nextFrontier *[]Pair, visited *[][]Pair, wg *sync.WaitGroup, mu_visited *sync.Mutex, mu_frontier *sync.Mutex, targetRow int, targetCol int, success *int32) {
	defer wg.Done()
	if !maze[cell.Row][cell.Col].Top {
		concurrentCheckNeighbor(Pair{Row: cell.Row - 1, Col: cell.Col}, nextFrontier, visited, mu_visited, mu_frontier, cell, targetRow, targetCol, success)
	}
	if !maze[cell.Row][cell.Col].Right {
		concurrentCheckNeighbor(Pair{Row: cell.Row, Col: cell.Col + 1}, nextFrontier, visited, mu_visited, mu_frontier, cell, targetRow, targetCol, success)
	}
	if !maze[cell.Row][cell.Col].Bottom {
		concurrentCheckNeighbor(Pair{Row: cell.Row + 1, Col: cell.Col}, nextFrontier, visited, mu_visited, mu_frontier, cell, targetRow, targetCol, success)
	}
	if !maze[cell.Row][cell.Col].Left {
		concurrentCheckNeighbor(Pair{Row: cell.Row, Col: cell.Col - 1}, nextFrontier, visited, mu_visited, mu_frontier, cell, targetRow, targetCol, success)
	}
}

func reconstructPath(visited [][]Pair, start Pair, goal Pair) []Pair {
	var path []Pair
	current := goal
	for current != start {
		path = append(path, current)
		current = visited[current.Row][current.Col]
	}
	path = append(path, start)
	return path
}

func ParallelBFS(maze [][]*Cell, numRows int, numCols int) ([]Pair, [][]Pair) {

	visited := make([][]Pair, len(maze))
	for i := range visited {
		visited[i] = make([]Pair, len(maze[0]))
	}
	for i := range visited {
		for j := range visited[0] {
			visited[i][j] = Pair{Row: -1, Col: -1}
		}
	}
	visited[0][0] = Pair{Row: 0, Col: 0}
	var frontier []Pair
	frontier = append(frontier, Pair{Row: 0, Col: 0})
	nextFrontier := []Pair{}
	var mu_visited sync.Mutex
	var mu_frontier sync.Mutex
	var wg sync.WaitGroup
	var success int32 = 0
	for (len(frontier) > 0) && atomic.LoadInt32(&success) == 0 {

		for i := range frontier {
			wg.Add(1)
			go concurrentOperation(maze, frontier[i], &nextFrontier, &visited, &wg, &mu_visited, &mu_frontier, numRows-1, numCols-1, &success)
		}
		wg.Wait()
		frontier = nextFrontier
		nextFrontier = []Pair{}
	}
	if success > 0 {
		println("Goal Found!, terminating BFS.")
	}
	return reconstructPath(visited, Pair{Row: 0, Col: 0}, Pair{Row: numRows - 1, Col: numCols - 1}), visited
}

func checkEnd(cell Pair, targetRow int, targetCol int, success *int) {
	if cell.Row == targetRow && cell.Col == targetCol {
		*success = 1
	}
}

func SequentialFindNeighbors(maze [][]*Cell, cell Pair, nextFrontier *[]Pair, visited *[][]Pair, targetRow int, targetCol int, success *int) {
	if !maze[cell.Row][cell.Col].Top && (*visited)[cell.Row-1][cell.Col] == (Pair{-1, -1}) {
		checkEnd(Pair{Row: cell.Row - 1, Col: cell.Col}, targetRow, targetCol, success)
		(*visited)[cell.Row-1][cell.Col] = Pair{Row: cell.Row, Col: cell.Col}
		*nextFrontier = append(*nextFrontier, Pair{Row: cell.Row - 1, Col: cell.Col})
	}
	if !maze[cell.Row][cell.Col].Right && (*visited)[cell.Row][cell.Col+1] == (Pair{-1, -1}) {
		checkEnd(Pair{Row: cell.Row, Col: cell.Col + 1}, targetRow, targetCol, success)
		(*visited)[cell.Row][cell.Col+1] = Pair{Row: cell.Row, Col: cell.Col}
		*nextFrontier = append(*nextFrontier, Pair{Row: cell.Row, Col: cell.Col + 1})
	}
	if !maze[cell.Row][cell.Col].Bottom && (*visited)[cell.Row+1][cell.Col] == (Pair{-1, -1}) {
		checkEnd(Pair{Row: cell.Row + 1, Col: cell.Col}, targetRow, targetCol, success)
		(*visited)[cell.Row+1][cell.Col] = Pair{Row: cell.Row, Col: cell.Col}
		*nextFrontier = append(*nextFrontier, Pair{Row: cell.Row + 1, Col: cell.Col})
	}
	if !maze[cell.Row][cell.Col].Left && (*visited)[cell.Row][cell.Col-1] == (Pair{-1, -1}) {
		checkEnd(Pair{Row: cell.Row, Col: cell.Col - 1}, targetRow, targetCol, success)
		(*visited)[cell.Row][cell.Col-1] = Pair{Row: cell.Row, Col: cell.Col}
		*nextFrontier = append(*nextFrontier, Pair{Row: cell.Row, Col: cell.Col - 1})
	}
}

func SequentialBFS(maze [][]*Cell, numRows int, numCols int) ([]Pair, [][]Pair) {

	visited := make([][]Pair, len(maze))
	for i := range visited {
		visited[i] = make([]Pair, len(maze[0]))
	}
	for i := range visited {
		for j := range visited[0] {
			visited[i][j] = Pair{Row: -1, Col: -1}
		}
	}
	visited[0][0] = Pair{Row: 0, Col: 0}
	var frontier []Pair
	frontier = append(frontier, Pair{Row: 0, Col: 0})
	nextFrontier := []Pair{}

	var success int = 0
	for (len(frontier) > 0) && success == 0 {

		for i := range frontier {
			SequentialFindNeighbors(maze, frontier[i], &nextFrontier, &visited, numRows-1, numCols-1, &success)
		}
		frontier = nextFrontier
		nextFrontier = []Pair{}
	}
	if success > 0 {
		println("Goal Found!, terminating BFS.")
	}
	return reconstructPath(visited, Pair{Row: 0, Col: 0}, Pair{Row: numRows - 1, Col: numCols - 1}), visited
}
