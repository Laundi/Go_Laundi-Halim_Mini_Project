package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponceWriter, r *http.Request) {
	fmt.FPrintf(w, format: "Welcome to HalimStore Home Page")
}
