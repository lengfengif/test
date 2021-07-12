package main

import "fmt"

func main () {


	type girl struct {
		name string

		hair string

		body string

		leg  string

		stars int
 
	}

	c := girl {

		name : "kate",

		hair : "yellow",

		body : "white",

		leg  : "long",

		stars : 5,
 
	} 
	e := &c

	e.name = "oo"
	fmt.Println(c)
}