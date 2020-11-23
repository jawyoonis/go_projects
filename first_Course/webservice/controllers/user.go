package controllers

import (
	"net/http"
	"regexp"
)

type UserController struct {
	UserIDPattern *regexp.Regexp
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the user Controller!"))

}

func newUserController() *UserController {

	return &UserController{
		UserIDPattern: regexp.MustCompile("^/users/(\\d+)/?"),
	}

}
