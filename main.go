package main

import (
	"fmt"
	"log"
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprintf(w, "wrong handler")
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/welcome" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hi There, welcome %s", r.URL.Path[1:])
}

func view(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/view" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "view end point")
}

func edit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "edit end point")
}

func main() {
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/view", view)
	fmt.Printf("Server started Succesfully")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
