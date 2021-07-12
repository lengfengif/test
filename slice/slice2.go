package main

import "fmt"

func main() {

	Slice := [] int {1, 2, 3, 4, 5, 6, 7}
	newSlice := Slice[1:3]
	newSlice[0] = 10	
	fmt.Println(Slice)
	fmt.Println(newSlice)
	
}

