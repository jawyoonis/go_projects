package main

import (
	"fmt"
)

func main() {
	port := 8080
	port, err := startWebServer(port)
	fmt.Println(port, err)

}

func startWebServer(port int) (int, error) {
	fmt.Println("starting webservices...")
	fmt.Println("webservices started on port ", port)
	return port, nil
}
