package main

import (
	"fmt"
)

func main() {
	// Declare a string pointer array of three elements.
	var array1 [3]*string

	// Declare a second string pointer array of three elements.
	// Initialize the array with string pointers.
	array2 := [3]*string{new(string), new(string), new(string)}

	// Add colors to each element
	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "Green"

	// Copy the values from array2 into array1.
	array1 = array2
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Printf("array1= %p | array2= %p\n", &array1, &array2)
}
