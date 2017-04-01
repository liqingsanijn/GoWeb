package main

import (
	"fmt"
	"net/http"
)

func main()  {
	fmt.Println("main")
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "hello, word")
}