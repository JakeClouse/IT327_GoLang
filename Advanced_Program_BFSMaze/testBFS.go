// Testing BFS for IT327 Advanced program
// Author: Brady McCue
// Date: 12/7/2025
package main

import (
	"fmt"
	"time"
)

func printMaze(m *Maze) {
	for r := 0; r < m.Height; r++ {
		for c := 0; c < m.Width; c++ {
			if m.Cells[r][c].Top {
				fmt.Print(" ---")
			} else {
				fmt.Print("    ")
			}
		}
		fmt.Println()

		for c := 0; c < m.Width; c++ {
			if m.Cells[r][c].Left {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print("   ")
		}

		if m.Cells[r][m.Width-1].Right {
			fmt.Print("|")
		}
		fmt.Println()
	}

	for c := 0; c < m.Width; c++ {
		if m.Cells[m.Height-1][c].Bottom {
			fmt.Print(" ---")
		} else {
			fmt.Print("    ")
		}
	}
	fmt.Println()
}

/* bad don't use
func printPath(visitedArray [][]Pair) {

	goal := visitedArray[len(visitedArray)-1][len(visitedArray[0])-1]
	start := visitedArray[0][0]
	current := goal

	goalColumn := len(visitedArray[0]) - 1
	goalRow := len(visitedArray) - 1

	i := goalRow
	j := goalColumn
	fmt.Println("Path from Goal to start found: ")
	//backtracking from goal to start
	for current != start {
		fmt.Printf("Cell (%d, %d) came from (%d, %d)\n", i, j, visitedArray[i][j].Row, visitedArray[i][j].Col)
		i = visitedArray[i][j].Row
		j = visitedArray[i][j].Col
		current = visitedArray[i][j]
	}

}
*/

func main() {
	Maze := CreateMaze(10, 10)
	//print maze, taken from Maze.cpp reference file
	printMaze(Maze)

	startBFSTime := time.Now()
	fmt.Println("Starting Parallel BFS...")

	ParallelBFS(Maze.Cells, 0, 0)
	//Maze.DS.PrintArrayValues()

	elapsedBFSTime := time.Since(startBFSTime)

	fmt.Printf("Parallel BFS took %s\n", elapsedBFSTime)

	//print all visited cells and their parents
	//printPath(visited)

}
