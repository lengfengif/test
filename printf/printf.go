package main

import "fmt"

func main () {
	
	s := "sss"
	i := 1
	var p *int
	p = &i
	fmt.Printf("%s, %d, %p",s, i, p)

}