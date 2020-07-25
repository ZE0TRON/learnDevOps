package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func users(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{users: []}")
	fmt.Println("Users requested")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", users)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
