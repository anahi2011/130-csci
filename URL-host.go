package main

import(
	"net/http"
	"io"

)

func url_path(res http.ResponseWriter, req *http.Request){
	key:="n"
	x:=req.FormValue(key)
	io.WriteString(res, "Name:"+ x)

}

func main(){
	http.HandleFunc("/", url_path)
	http.ListenAndServe(":8080", nil)
}