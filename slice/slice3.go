package main

import "fmt"

func sliceOps () {

s1 := []int{2, 4, 6, 8}

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	fmt.Println(s2) 
	fmt.Println(s3) 
	fmt.Println("Copying slice")
	copy(s2, s1)
	fmt.Println(s2) 

	fmt.Println("删除掉中间的元素")
	s2 = append(s2[:3], s2[4:]...) 
	
	fmt.Println(s2) 

	fmt.Println("移除第一个元素")
	s2 = s2[1:]
	fmt.Println(s2) 

	fmt.Println("移除最后一个元素")
	s2 = s2[:len(s2)-1]
	fmt.Println(s2)  
}
func main() {

	sliceOps()

}