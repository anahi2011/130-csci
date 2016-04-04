//Create a webpage that serves a form and allows the user to upload a txt file. You do not need to
//check if the file is a txt; bad programming but just trust the user to follow the instructions. Once
//a user has uploaded a txt file, copy the text from the file and display it on the webpage. Use
//req.FormFile and io.Copy to do this
package main


import (
	"os"
	"io"
	"path/filepath"
	"net/http"
	"html/template"
	"log"
)
func fileP(res http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(res, nil)

}

func upload(res http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		file, _, err := req.FormFile("n")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer file.Close()

		src := io.LimitReader(file, 400)

		dst, err := os.Create(filepath.Join(".", "file.txt"))
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()

		io.Copy(dst, src)
	}

}

func main() {

	http.HandleFunc("/", fileP)
	http.ListenAndServe(":8020", nil)

}