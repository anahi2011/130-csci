
package main

import (
	"fmt"
	"log"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, `<!DOCTYPE HTML>
		<html>
			<head>
			<title></title>
			</head>
			<body>
				Sample secure server
			</body>
		</html>
		`)
}

func redir(res http.ResponseWriter, req *http.Request){
	http.Redirect(res, req, "https://127.0.0.1:10443/" + req.RequestURI, http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", foo)
	log.Println("Listening...on 10443.  Go to https://127.0.0.1:10443/")
	go http.ListenAndServe(":8080", http.HandlerFunc(redir))
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
