
package main

import(
	"net/http"
	"io"
)

func display_url(res http.ResponseWriter, req *http.Request ){
	io.WriteString(res, req.URL.Path)

}

func main(){
	http.HandleFunc("/", display_url)
	http.ListenAndServe(":8080", nil)

}