package controllers

import (
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseGlob("templates/*.html"))

func Home(w http.ResponseWriter, r *http.Request) {

	t.ExecuteTemplate(w, "Index", nil)
}

func New(w http.ResponseWriter, r *http.Request) {

}

func Insert(w http.ResponseWriter, r *http.Request) {

}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {

}
