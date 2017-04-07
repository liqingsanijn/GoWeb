package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main()  {
	fmt.Println("main")
	http.Handle("/", indexMiddleWare(http.HandlerFunc(index)))
	http.Handle("/test", testMiddleWare(http.HandlerFunc(test)))
	http.ListenAndServe("localhost:8080", nil)
}

func indexMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("index")
		next.ServeHTTP(w, r)
	})
}

func testMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test")
		next.ServeHTTP(w, r)
	})
}

func index(w http.ResponseWriter, r *http.Request)  {
	bytes, err := ioutil.ReadFile("webapp/index.html")
	if err != nil {
		fmt.Fprintln(w, err)
		fmt.Println(err)
	}
	fmt.Fprintln(w, string(bytes))
}

func test(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "test")
}
