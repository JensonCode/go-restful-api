package routes

import (
	"fmt"
	"net/http"
)


func NewRouter() http.Handler{
	mux:= http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api", apiHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r* http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func apiHandler(w http.ResponseWriter, r* http.Request) {
	fmt.Fprintln(w, "Calling API")
}
