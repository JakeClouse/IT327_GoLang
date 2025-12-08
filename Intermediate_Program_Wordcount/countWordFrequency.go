/*  Reads a text file and outputs every word present
*   Author:Marcos Villalobos
*	Date: 11/7/2025
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Ensure proper use
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run countWordFrequency.go <filename>")
		return
	}
	filename := os.Args[1]

	//Open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer file.Close()

	// Create map
	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(word)
			wordCount[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	fmt.Printf("Read %d lines from %s:\n", len(wordCount), filename)
	for key, frequency := range wordCount {
		fmt.Printf("word \"%s\" appeared %d many time(s)\n", key, frequency)
	}
}
