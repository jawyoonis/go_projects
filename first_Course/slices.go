package main

import "fmt"

func main() {

	arr := [4]int{1, 2, 3, 4} // arrays in go are limited to the specified size of the array
	slice := arr[:]
	arr[2]=34
	slice[1]=34
	fmt.Println(arr, slice)

	slice = append(slice, 3, 4, 5 ,6,7) // slice does not have size limit unlike arrays
	fmt.Println(slice)

	s2 :=slice[1:]

	s3 :=slice[:2]
	s4 :=slice[1:2]
	
	
	fmt.Println(s2, s3,s4)

}