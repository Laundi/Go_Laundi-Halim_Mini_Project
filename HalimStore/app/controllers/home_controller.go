package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponceWriter, r *http.Request) {
	fmt.Printf(w, format: "Welcome to GoToko Home Page")
}

