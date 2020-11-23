package main

import (
	"net/http"

	"github.com/go_projects/first_Course/controllers"
)

func main() {
	controllers.RegisterController()

	http.ListenAndServe(":8080", nil)
}
