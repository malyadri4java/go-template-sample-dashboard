package main

import (
	"log"
	"net/http"
	"user_managment/routing"
)

func init() {
	log.Println("User Managment Application Starting....")
}

func main() {
	r := routing.Router()
	fs := http.FileServer(http.Dir("./static/"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	//r.PathPrefix("/").Handler(fs)
	log.Println("Application Started on : ", "localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
