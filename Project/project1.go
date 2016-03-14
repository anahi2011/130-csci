package main
import(
	"net/http"
	"html/template"
        "log"
)
// this just serves the executed template so we could have the good stuff of the first part of the project
func upload(res http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", upload)
	http.ListenAndServe(":8000", nil)
}

