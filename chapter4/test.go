package main

import "fmt"

const length int = 1E6

func main() {
	var arr = [5]int{10,20,30,40}
	var arr2 = [...]int{24,50,34}
	var arr3 = [6]int{0 : 100, 5: 50, 3:20}
	var arr4 = [...]*int{0: new(int), 1: new(int)}
	fmt.Printf("type is %T", length)
	fmt.Println(arr)
	fmt.Println(len(arr2))
	fmt.Println(arr3)
	arr[4] = 9
	arr[2] = 100
	fmt.Println(arr)
	*arr4[0] = 13
	*arr4[1] = 14
	fmt.Println(arr4)

	var strArr1 = [4]string{"Hello", "world", "it", "me"}
	var strArr2 [4]string
	strArr2 = strArr1
	fmt.Println(strArr2)
	strArr2[0] = "good morning"
	fmt.Println(strArr1)
	fmt.Println(strArr2)
	fmt.Printf("Addr of strArr2 :  %p\n", &strArr2)
	fmt.Printf("Addr of strArr1 :  %p\n", &strArr1)
	var largeArr [length]int
	foo(largeArr)
	fmt.Printf("Before fooPointer :  0: %v, 1: %v\n", largeArr[0], largeArr[1])
	fooPointer(&largeArr)
	fmt.Printf("Outside fooPointer :  0: %v, 1: %v\n", largeArr[0], largeArr[1])

}

func foo(array [length]int){
	fmt.Println(len(array))
}

func fooPointer(array *[length]int) {
	array[0] = 10 // why * can be omitted ? 
	// array is a pointer to an array, so (*array)[1] will dereference to the element at index 1
	(*array)[1] = 11
	//*array[2] = 12 // this will not work, because it means *(array[2]) not (*array)[2]
	fmt.Printf("Inside fooPointer : 0: %v, 1: %v\n", array[0], array[1])
}