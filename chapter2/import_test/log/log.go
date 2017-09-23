package log

import "fmt"


func init() {
	fmt.Println("log init")
	LogInt()
}

func LogInt(){
	fmt.Println("log int")
}