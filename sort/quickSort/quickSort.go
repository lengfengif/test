package main

import "fmt"

func main () {

	var number1 = make([] int, 10)
	number1 =[] int {7, 3, 4, 2, 1, 5, 8, 7, 9, 11}
	quicksort(number1, 0, len(number1))
	fmt.Println(number1)

}

func quicksort (s [] int,low int, high int ) {
	if low >= high{
		return
	}

	i := low
	j := high 
	base := s[low]
	fmt.Println(i, j, base)
	for i < j {
		for s[j] <= base && i < j {
			j--
		}
		for s [i] <= base && i < j {
			i++
		} 
		swap(s, i, j)
	}
		swap(s, low, high)
}
func swap(s [] int, i int, j int){

	temp := s[i]
	s[i] = s[j]
	s[j] = temp

}