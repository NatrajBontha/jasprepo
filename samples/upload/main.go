package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Field struct {
	Firstname  string
	Secondname string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Index Template Parse Error: ", err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Index Template Execution Error: ", err)
	}

}

func main() {
	http.HandleFunc("/", RootHandler) // sets router
	http.HandleFunc("/welcome", WelcomeHandler)
	err := http.ListenAndServe(":4000", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {

	Firstname := r.FormValue("FirstName")
	Secondname := r.FormValue("LastName")
	fmt.Println(Firstname)

	f := new(Field)
	f.Firstname = Firstname
	f.Secondname = Secondname
	fmt.Println(*f)
	tmpl, err := template.ParseFiles("tmpl/welcome.tmpl")
	if err != nil {
		fmt.Println("Index Template Parse Error: ", err)
	}
	err = tmpl.Execute(w, &f)
	if err != nil {
		fmt.Println("Index Template Execution Error: ", err)
	}

}
