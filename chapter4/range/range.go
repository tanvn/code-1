package main

import "fmt"

func main() {
	// Create a slice of integers.
	// Contains a length and capacity of 4 elements.
	slice := []int{10, 20, 30, 40}
	fmt.Printf("Array-Addr: %X  First elelemt address: %X\n",
		&slice, &slice[0])
	// Iterate over each element and display the value and addresses.
	for index, value := range slice {
		fmt.Printf("Value: %d  Value-Addr: %X  ElemAddr: %X\n",
			value, &value, &slice[index])
	}
	// Iterate over each element starting at element 3.
	for index := 2; index < len(slice); index++ {
		fmt.Printf("Index: %d  Value: %d\n", index, slice[index])
	}

	// Create a slice of a slice of integers.
	slice2 := [][]int{{10}, {100, 200}}
	for index, value := range slice2 {
		fmt.Printf("Value: %d  Value-Addr: %X  ElemAddr: %X\n",
			value, &value[0], &slice2[index][0])
	}

	// Append the value of 20 to the first slice of integers.
	slice2[0] = append(slice2[0], 20)
	for index, value := range slice2 {
		fmt.Printf("Value: %d  Value-Addr: %X  ElemAddr: %X\n",
			value, &value[0], &slice2[index][0])
	}

	bigSlice := make([]int, 1e6)
	fmt.Printf("first ele = %v\n", bigSlice[0])

	// Pass the slice to the function foo.
	newSlice :=foo(bigSlice)
	fmt.Printf("newSlice = %v\n", newSlice)

}

func foo(slice []int)  []int {
	fmt.Printf("Array size = %v\n", len(slice))
	newSlice := slice[0:10:10]

	return append(newSlice, 1,2,3,4,5)
}
