package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//	inline declaration of struct
	book := Book{
		Title:  "Lions",
		Author: "Myself",
		Pages:  321,
	}
	json.NewEncoder(w).Encode(book)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 style='color:steelblue'>Hello World..~</h1>"))
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/book", GetBook)
	// func(rw http.ResponseWriter, r *http.  Request  --> Inline func.
	//{
	// rw.Write([]byte("Hello"))
	// })

	log.Fatal(http.ListenAndServe(":8080", nil))
}
