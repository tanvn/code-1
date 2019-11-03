package main

import "fmt"

func main() {
	// Declare a two dimensional integer array of four elements
	// by two elements.
	var array [4][2]int

	fmt.Println(array)
	// Use an array literal to declare and initialize a two
	// dimensional integer array.
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array)

	// Declare and initialize index 1 and 3 of the outer array.
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	fmt.Println(array)

	// Declare and initialize individual elements of the outer
	// and inner array.
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(array)

}
