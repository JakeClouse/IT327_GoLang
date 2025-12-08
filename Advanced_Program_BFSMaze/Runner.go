// Testing BFS for IT327 Advanced program
// Author: Brady McCue
// Date: 12/7/2025

package main

import (
	"fmt"
	"time"
)

func checkSolutionPath(path []Pair, point Pair) bool {
	for i := range path {
		if path[i] == point {
			return true
		}
	}
	return false
}
func printMaze(m *Maze, path []Pair) {
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
			if (checkSolutionPath(path, Pair{Row: r, Col: c})) {
				fmt.Print(" X ")

			} else {
				fmt.Print("   ")
			}
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

func main() {
	Maze := CreateMaze(100, 100)
	//print maze, taken from Maze.cpp reference file
	printMaze(Maze, nil)

	startBFSTime := time.Now()
	fmt.Println("\nStarting Parallel BFS...")

	path, _ := ParallelBFS(Maze.Cells, 100, 100)
	//Maze.DS.PrintArrayValues()

	elapsedBFSTime := time.Since(startBFSTime)

	fmt.Printf("Parallel BFS took %s\n", elapsedBFSTime)

	startSequentialBFSTime := time.Now()
	fmt.Println("Starting Sequential BFS...")

	SequentialBFS(Maze.Cells, 100, 100)
	//Maze.DS.PrintArrayValues()

	elapsedSequentialBFSTime := time.Since(startSequentialBFSTime)

	fmt.Printf("Sequential BFS took %s\n\n", elapsedSequentialBFSTime)

	printMaze(Maze, path)

}
