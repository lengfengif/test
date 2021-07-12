package main

import "fmt"

func main () {

	var k = make([] string, 0)
	fmt.Println(k, len(k), cap(k))
	k = [] string {"g", "o", "l", "a", "n", "g"}
	fmt.Println(k[2:4])
	
} 
