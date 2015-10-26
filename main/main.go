package main

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type MainText struct {
	ResponseText string
}
type PasswordInput struct {
	Password string
}

func newMainText() *MainText {
	return &MainText{
		ResponseText: "5 Second Delay",
	}
}

func timehandler(w http.ResponseWriter, r *http.Request) {
	m := newMainText()
	//
	select {
	case <-time.After(time.Second * 5):
		renderTemplate(w, "main.html", m)
	}

}

func post(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t PasswordInput
	err := decoder.Decode(&t)
	if err != nil {
		log.Print("form parsing error: ", err)
	}
	hasher := sha512.New()
	ba := []byte(t.Password)
	hasher.Write(ba)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	js, err := json.Marshal(sha)
	w.Write(js)
}

func renderTemplate(w http.ResponseWriter, tmpl string, m *MainText) {
	tmpl = fmt.Sprintf("public/%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, m)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func main() {
	http.HandleFunc("/", timehandler)
	http.HandleFunc("/post", post)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":3000", nil)

}
