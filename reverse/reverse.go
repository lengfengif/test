package main

import "fmt"

func main () {
	arr:=[7] int {7, 6, 5, 4, 3, 2, 1}
	    i := 0
		j := len(arr)-1
		for ;i < j; i++{
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
			j--
		}
	fmt.Println(arr)
}
