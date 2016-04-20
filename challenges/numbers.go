package main

import "fmt"

func main() {
	var numOne int
	var numTwo int
	fmt.Print("Please enter your moms age: ")
	fmt.Scan(&numOne)
	fmt.Print("Please enter your age: ")
	fmt.Scan(&numTwo)
	fmt.Println(numOne, "*", numTwo, " = ", numOne*numTwo)
}