package main

import "fmt"

func main() {

	a := 3 
	var p *int
	p = &a
	fmt.Println(*p)
	fmt.Println(p)

}


