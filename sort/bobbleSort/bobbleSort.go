package main

import "fmt"

func main () {
	arr:=[10] int {10, 5, 6, 11, 10, 39, 23, 53, 10 ,2}
	
		for i := 1; i < len(arr) ; i++ {
			for j := 0; j < len(arr)-i ; j++ {
				if arr[j] > arr [j+1]{
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				}
			}
		}
	fmt.Println(arr)
}
