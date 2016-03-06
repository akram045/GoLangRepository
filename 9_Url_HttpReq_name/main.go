package main

import (

	"net/http"
	"io"
)

func urlName(res http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(res, "%v", req.URL.Path)

	key := "q"
	val := req.FormValue(key)
	io.WriteString(res, "Query String - Name :"+val)
}

func main() {
	http.HandleFunc("/", urlName)
	http.ListenAndServe(":8080", nil)
}