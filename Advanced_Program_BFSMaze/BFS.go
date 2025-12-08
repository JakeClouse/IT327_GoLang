//Parallel BFS for IT327 Advanced program
// Author: Jake CLouse
// Date: 12/6/2025

//Edited by Brady McCue to have paramaters match with generated maze

package main

import (
	"sync"
)

type Pair struct {
	Row, Col int
}

func checkNeigbor(cell Pair, nextFrontier *[]Pair, visited *[][]Pair, mu_visited *sync.Mutex, mu_frontier *sync.Mutex, parentCell Pair) {
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

func concurrentOperation(maze [][]*Cell, cell Pair, nextFrontier *[]Pair, visited *[][]Pair, wg *sync.WaitGroup, mu_visited *sync.Mutex, mu_frontier *sync.Mutex) {
	defer wg.Done()
	if !maze[cell.Row][cell.Col].Top {
		checkNeigbor(Pair{Row: cell.Row - 1, Col: cell.Col}, nextFrontier, visited, mu_visited, mu_frontier, cell)
	}
	if !maze[cell.Row][cell.Col].Right {
		checkNeigbor(Pair{Row: cell.Row, Col: cell.Col + 1}, nextFrontier, visited, mu_visited, mu_frontier, cell)
	}
	if !maze[cell.Row][cell.Col].Bottom {
		checkNeigbor(Pair{Row: cell.Row + 1, Col: cell.Col}, nextFrontier, visited, mu_visited, mu_frontier, cell)
	}
	if !maze[cell.Row][cell.Col].Left {
		checkNeigbor(Pair{Row: cell.Row, Col: cell.Col - 1}, nextFrontier, visited, mu_visited, mu_frontier, cell)
	}
}

func ParallelBFS(maze [][]*Cell, startRow int, startCol int) [][]Pair {

	visited := make([][]Pair, len(maze))
	for i := range visited {
		visited[i] = make([]Pair, len(maze[0]))
	}
	for i := range visited {
		for j := range visited[0] {
			visited[i][j] = Pair{Row: -1, Col: -1}
		}
	}
	visited[startRow][startCol] = Pair{Row: startRow, Col: startCol}
	var frontier []Pair
	frontier = append(frontier, Pair{Row: startRow, Col: startCol})
	nextFrontier := []Pair{}
	var mu_visited sync.Mutex
	var mu_frontier sync.Mutex
	var wg sync.WaitGroup
	for len(frontier) > 0 {

		for i := range frontier {
			wg.Add(1)
			go concurrentOperation(maze, frontier[i], &nextFrontier, &visited, &wg, &mu_visited, &mu_frontier)
		}
		wg.Wait()
		frontier = nextFrontier
		nextFrontier = []Pair{}
	}
	return visited
}
