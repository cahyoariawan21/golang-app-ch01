package main

import (
	"fmt"
	"log"
	"net/http"
)

// this function is to handle Form that have Method POST
// First create request ParseForm  and give the handle error with check if error give the return response error
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	// If error is nil give the response POST request is success
	fmt.Fprintf(w, "POST request success \n\n")

	// POST data
	name := r.FormValue("name")
	address := r.FormValue("address")
	// Print response value
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

// This function is create just for Print Hello to user
// Method available just GET if Method not GET give the return error
// if user not access with route /hello then give the return error not found page
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if URL Path route is not /hello then return error: route not found
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// Check if Method is not GET then return error: method is not supported!
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	// If 2 Check is qualified then print this
	fmt.Fprintf(w, "Hello User!")
}
func main() {
	// fmt.Println("This is main setup")
	fileServer := http.FileServer(http.Dir("./static"))

	// Create route handler function
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8080\n")
	// Handle Error Check
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
