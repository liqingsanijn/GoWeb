package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main()  {
	fmt.Println("main")
	http.HandleFunc("/", index)
	http.ListenAndServe("localhost:8080", nil)
}

func index(w http.ResponseWriter, r *http.Request)  {
	bytes, err := ioutil.ReadFile("webapp/index.html")
	if (err != nil) {
		fmt.Fprintln(w, err)
		fmt.Println(err)
	}
	fmt.Fprintln(w, string(bytes))
}