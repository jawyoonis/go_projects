package main

import "fmt"
func main() {
	firstName:="jamal"
	ptr:= &firstName
	fmt.Println(ptr, *ptr)
}