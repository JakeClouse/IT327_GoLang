// Testing BFS for IT327 Advanced program
// Author: Brady McCue
// Date: 12/7/2025
package main

import (
	"fmt"
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

func main() {
	Maze := CreateMaze(10, 10)
	//print maze, taken from Maze.cpp reference file
	printMaze(Maze)
	visited := ParallelBFS(Maze.Cells, 0, 0)
	//show all maze array values
	Maze.DS.PrintArrayValues()
	//print all visited cells and their parents?
	for i := range visited {
		for j := range visited[0] {
			fmt.Printf("Cell (%d, %d) came from (%d, %d)\n", i, j, visited[i][j].Row, visited[i][j].Col)
		}
	}
}
