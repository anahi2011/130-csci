package main

import (
	"io"
	"net/http"
	"strconv"
)

func userCookie(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("user-info")
	if err == http.ErrNoCookie{
		cookie = &http.Cookie{
			Name: "user-info",
			Value: "0",
		}
	}
	count, _ := strconv.Atoi(cookie.Value)
	count ++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(res, cookie)
	io.WriteString(res, cookie.Value)
}


func main() {
	http.HandleFunc("/", userCookie)
	http.ListenAndServe(":8080", nil)
}
