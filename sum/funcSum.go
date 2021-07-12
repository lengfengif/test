package main

import "fmt"

func main () {

	var a int 
	var b int
	a = 3
	b = 2
	s := sum (a, b)
	fmt.Println(s)

}
func sum (num1 int, num2 int) int {

	return num1 + num2 

}

