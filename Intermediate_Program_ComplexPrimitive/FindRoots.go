package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	//collecting user input
	var inputBuf string = ""
	fmt.Println("How many terms does your function have?(up to 3): ")
	fmt.Scanln(&inputBuf)

	numTerms, err := strconv.Atoi(inputBuf)
	if err != nil {
		os.Exit(1)
	}

	var i int = numTerms
	function := make([]float64, numTerms)

	for i >= 0 {
		fmt.Println("Input term " + strconv.Itoa(i) + " of your function(highest order first): ")
		fmt.Scanln(&inputBuf)

		value, err := strconv.ParseFloat(inputBuf, 64)
		if err != nil {
			os.Exit(1)
		}

		function[i] = value
		i++
	}

	//starting function processing.

}
