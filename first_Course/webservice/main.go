package main

import (
	"fmt"
	"net/http"

	"github.com/go_projects/first_Course/webservice/controllers"
)

func main() {
	controllers.RegisterController()

	port := ":5000"
	http.ListenAndServe(port, nil)
	fmt.Println("The server is listening on %s", port)
}
