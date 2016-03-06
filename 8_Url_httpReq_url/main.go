package main

import (

	"net/http"
	"strings"
	"log"
	"io"
)

func urlName(res http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(res, "%v", req.URL.Path)

	name := strings.Split(req.URL.Path, "/")
	log.Println(name)
	io.WriteString(res, "Name: " + name[1])
}

func main() {
	http.HandleFunc("/", urlName)
	http.ListenAndServe(":8080", nil)
}