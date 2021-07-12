package main

import "fmt"

func main () {

    var s = make([] int, 10)    
    fmt.Println(s, len(s),cap(s))  
    s = [] int {1,2}        
    fmt.Println(s, len(s),cap(s))
    s = append(s, 5)           
    fmt.Println(s, len(s),cap(s))
    var s2 = s
    fmt.Println(s2) 
    s2[2]=6
    fmt.Println(s)

} 
