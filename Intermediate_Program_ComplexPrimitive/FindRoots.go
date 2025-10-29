package main

import (
	"fmt"
	"math/cmplx"
	"os"
	"strconv"
)

func getUserInput() [3]float64 {

	//collecting user input
	var inputBuf string = ""
	fmt.Println("Please Input you function terms, Starting with x^2 term, in the form y = (3)x^2 + (2)x + (1): ")

	var function [3]float64

	var i int = 2
	for i >= 0 {
		fmt.Println("Input term " + strconv.Itoa(i+1) + " of your function: ")
		fmt.Scanln(&inputBuf)

		value, err := strconv.ParseFloat(inputBuf, 64)
		if err != nil {
			os.Exit(1)
		}

		function[i] = value
		i--
	}

	return function

}

func getRoots(function [3]float64) {
	//does the quadratic formula with the variable names from the equation
	var a float64 = function[2]
	var b float64 = function[1]
	var c float64 = function[0]

	var root1 complex128
	var root2 complex128

	var determinant complex128 = cmplx.Sqrt(complex((b*b)-(4*a*c), 0))

	root1 = (-complex(b, 0) + determinant) / (2 * complex(a, 0))
	root2 = (-complex(b, 0) - determinant) / (2 * complex(a, 0))

	fmt.Println("First Root: ", root1)
	fmt.Println("Second Root: ", root2)

}

func main() {
	var function [3]float64 = getUserInput()
	getRoots(function)
}
