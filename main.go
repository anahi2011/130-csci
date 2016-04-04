//serves a form and allows the user to enter their name. Once a user has entered their name, show
//their name on the webpage. Use req.FormValue to do this
package main

import(
	"net/http"
	"io"
)

func url_path(res http.ResponseWriter, req *http.Request){
	key:="n"
	res.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(res,`<form method="GET">
		 <input type="text" name="n">
		 <input type="submit">
		</form>`+ req.FormValue(key))
}

func main(){
	http.HandleFunc("/", url_path)
	http.ListenAndServe(":8080", nil)
}