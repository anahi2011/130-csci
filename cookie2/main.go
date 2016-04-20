package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"crypto/hmac"

)

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("hey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func foo(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	cookie, err := r.Cookie("session-id")
	if err != nil {
		cookie = &http.Cookie{
			Name:     "session-id",
			Value:    "cookie-value",
			HttpOnly: true,
		}
	}

	if r.FormValue("email") != "" {
		email := r.FormValue("email")
		cookie.Value = email + `|` + getCode(email)
	}
	
}
func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8086", nil)
}