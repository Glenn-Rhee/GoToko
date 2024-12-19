package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to GoToko Home Page")
}

func Profile(w http.ResponseWriter, r * http.Request){
	fmt.Fprintf(w, "Welcome to GoToko Profile Page")
}