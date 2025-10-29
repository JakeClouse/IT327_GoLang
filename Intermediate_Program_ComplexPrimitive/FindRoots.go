package main

import "fmt"
import "strconv"

func main() {

	//collecting user input
	var input String = ""
	fmt.Println("How many terms does your function have?(up to 3): ")
	fmt.ScanLn(&input)

	var numTerms int = strconv.Atoi(input)

	var i int = numTerms
	var function [numTerms]double = []

	while (i >= 0){
		fmt.Println("Input term " + i + " of your function(highest order first): ")
		fmt.ScanLn(&input[i])
	}


	//starting function processing.













}
