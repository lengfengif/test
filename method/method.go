package main

import "fmt"

type cat struct {
		
	color string
	gender byte
	age int

}
func (d cat) say() {
	
	//fmt.Printf("%s %d %d", d.color, d.gender, d.age)
	fmt.Println(d.color, d.gender, d.age)

}

func main() {

	var b cat 
	b.color = "yellow"
	b.say()
	
}	

