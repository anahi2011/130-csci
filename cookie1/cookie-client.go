package main

import  (
	"fmt"
	"log"
	"github.com/nu7hatch/gouuid"
	"net/http"
)
func handler(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("session-id")
	if err != nil{

		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-id",
			Value: id.String(),
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Println(cookie)
}
func main(){
	http.HandleFunc("/", handler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
