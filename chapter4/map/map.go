package main

import "fmt"

func main() {
	dict := make(map[string]int)

	initialDict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

	fmt.Printf("dict = %v\n", dict)
	fmt.Printf("initialDict = %v\n", initialDict)

	//compileErrorDict := map[[]string]int{} // compile error
	//fmt.Printf("compileErrorDict = %v\n", compileErrorDict)
	stringMap := map[int][]string{}

	stringMap[1] = []string{"hello","world" }
	stringMap[10] = []string{"take","it","easy" }

	stringMap[100] = []string{"this","is","golang" }

	fmt.Printf("stringMap = %v\n", stringMap)

	//// Create a nil map by just declaring the map.
	//var colors map[string]string
	//// Add the Red color code to the map.
	//colors["Red"] = "#da1337" // panic: assignment to entry in nil map
	fmt.Printf("stringMap[1] pointer = %p, length = %d, cap = %v\n", stringMap[1], len(stringMap[1]), cap(stringMap[1]))

	firstText, found := stringMap[1]
	if found{
		fmt.Printf("firstText pointer = %p, length = %d, cap = %v\n", &firstText, len(firstText), cap(firstText))

	}

	secondText := stringMap[2]
	if secondText != nil {
		fmt.Printf("secondText = %v\n", secondText)

	}

	//change element in array is visible to the map
	// firstText is a slice, so append will change it length and cap, but the
	firstText = append(firstText, "from", "Tokyo")
	fmt.Printf("firstText (after append) pointer = %p, length = %d, cap = %v\n", &firstText, len(firstText), cap(firstText))

	firstTextNew, foundNew := stringMap[1]
	if foundNew {
		fmt.Printf("firstTextNew = %p\n", &firstTextNew)

	}

	for key,value := range stringMap{
		fmt.Printf("key= %d, value = %v\n", key,value)

	}

	delete(stringMap, 100)
	fmt.Println("---------- after deleting key = 100----------")

	for key,value := range stringMap{
		fmt.Printf("key= %d, value = %v\n", key,value)

	}
	editMap(stringMap)

	fmt.Println("---------- after editing map----------")
	for key,value := range stringMap{
		fmt.Printf("key= %d, value = %v\n", key,value)

	}

}

func editMap(stringMap map[int][]string)  {
	delete(stringMap, 1)
	stringMap[2] = []string{"replaced", "text"}
	stringMap[10][0] = "trust me, take"

}
