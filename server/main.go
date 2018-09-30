package main

import (
	"net/http"
	"io"
	)


func main() {
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		io.WriteString(w,"hello world!")
	})
	http.ListenAndServe(":8086",nil)
}
