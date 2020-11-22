package main

import "fmt"

func main(){

	var arr [3]int //first way to declare arrays in go 
	arr[1]=2
	arr[2]=3
	fmt.Println(arr)
	// fmt.Println(arr[5])
	arr2 := [3]int{1,2,3} //second way to declare arrays in go
	fmt.Println(arr2)

}