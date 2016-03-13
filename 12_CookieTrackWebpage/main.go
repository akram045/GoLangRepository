package main

import (

	"net/http"
	"strconv"
	"log"
	"io"
)

func urlName(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/favicon.ico" {
		http.NotFound(res, req)
		return
	}

	cookie, err := req.Cookie("Counter")
	logError(err)
	if cookie == nil {
		http.SetCookie(res, &http.Cookie{
			Name:  "Counter",
			Value: "1"})
	} else {
		counter, err := strconv.Atoi(cookie.Value)
		logError(err)
		counter++
		cookie.Value = strconv.Itoa(counter)
		http.SetCookie(res,cookie )
	}
	io.WriteString(res, cookie.Value)
}
func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", urlName)

	http.ListenAndServe(":8080", nil)
}