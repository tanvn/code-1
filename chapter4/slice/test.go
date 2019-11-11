package main

import "fmt"

func main() {
	slice := make([]string, 5)
	var stringSlices = []string{"Hello", "world", "this", "is", "me"}
	fmt.Printf("stringSlices= %v %v\n", cap(stringSlices), len(stringSlices))

	var slice2 = []string{100: "Clear"}
	fmt.Printf("%v %v\n", cap(slice2), len(slice2))
	stringSlices[1] = "Earth"
	fmt.Println(stringSlices)

	var intSlice = []int{10,20,30,40,50, 9:0}
	var slicedSlice = intSlice[1:3]
	fmt.Println(slicedSlice[1])
	fmt.Printf("slicedSlice= %v %v\n", len(slicedSlice), cap(slicedSlice))

	//fmt.Println(slicedSlice[2]) // out of range error

	// take  element at 2 and 3 index, set the capacity = 8-2 = 6
	var sliceWithCapacity = intSlice[2:4:8]
	fmt.Println(sliceWithCapacity)
	fmt.Printf("sliceWithCapacity : %v %v\n", cap(sliceWithCapacity), sliceWithCapacity)


	fmt.Printf("slicedSlice cap = %v |intSlice cap = %v \n",cap(slicedSlice), cap(intSlice))

	//change the shared element of 2 slices
	slicedSlice[0] = 22
	//see how it affects both slices
	fmt.Println(intSlice)
	fmt.Println(slicedSlice)

	slicedSlice = append(slicedSlice, 45)
	fmt.Println(intSlice)
	fmt.Println(slicedSlice)
	//invalidSlice := make([]int, 4, 2)  //length must be smaller than or equal to cap
	//fmt.Println(invalidSlice)
	fmt.Printf("slice length = %v\n",len(slice))

	// Create a slice of strings.
	// Contains a length and capacity of 5 elements.
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	// Slice the third element and restrict the capacity.
	// Contains a length and capacity of 1 element.
	detachedSlice := source[1:3:3]

	// Append a new string to the slice.
	detachedSlice = append(detachedSlice, "Kiwi")
	fmt.Println(source)
	fmt.Printf("detachedSlice = %v %v\n", detachedSlice, cap(detachedSlice))


	s1 := []int{1, 2}
	s2 := []int{3, 4}

	// Append the two slices together and display the results.
	fmt.Printf("%v\n", append(s1, s2...))
	for index, value := range source{
		fmt.Printf("index = %v, value=%v \n", index,value)
	}

}
