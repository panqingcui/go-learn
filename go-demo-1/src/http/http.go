package main

import (
	"fmt"
	"net/http"
)

//http 监听
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Print(w, "Hello world")
}

func main() {
	fmt.Println("vim-go")
	http.HandleFunc("/", Hello)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
