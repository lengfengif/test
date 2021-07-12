package main

import "fmt"

func main () {

	var ch chan
	ch = make (chan int )
	go say()
	a:= <- ch

}
func say (ch chan) int {
	ch <- 1
}
