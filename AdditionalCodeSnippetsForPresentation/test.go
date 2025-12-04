package main

import "fmt"

func main() {

	var a int = 10

	addTen(&a)

	fmt.Println(a)

}

func addTen(b *int) int {
	*b = *b + 10
	return *b
}

/*
func incrementAndAdd(a int) func(b int) int {
	return func(b int) int {
		a++
		return a + b
	}
}

func main() {

	fmt.Println("Function Currying: ")

	newFunction := incrementAndAdd(5)
	fmt.Println(newFunction(10))

	newFunction2 := incrementAndAdd(10)
	fmt.Println(newFunction2(20))
}
*/

/*
	var x int = 2
	var y float64 = 2.5
	var z float64 = y + x
	fmt.Println(z)
*/

/*
	x := 2
	y := 2.5
	z := x+y
	fmt.Println(z)
*/
/*
	var x rune = 'ඞ'
	fmt.Printf("%c", x)
*/
