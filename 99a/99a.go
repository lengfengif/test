package main

import "fmt"


func main() {

	var num int

	slice := make([]float32,1)
	for i := 1; i <= 9; i++ {
		for j := i; j > 0; j-- {
			num = i*j
			slice = append(slice,float32(num)/100)
		}
	}
	
	fmt.Println(slice,len(slice))
	
}