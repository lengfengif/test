package main

import "fmt"

func main() {

	var a[9] int
	var b[2][3] int
	a[6] = 1
	b[1][2] = 3
	arr := [5] int { 1, 2, 3, 4, 5 }
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(arr)

	p := &a

	fmt.Println(p)

	
}