package main

import "fmt"


func main() {

	slice := make([]int,3,10)
	slice = append(slice, 2)	

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	
}