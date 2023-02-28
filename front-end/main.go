package main

import (
	"net/http"
	"text/template"
)

// integração com html5
var temp = template.Must(template.ParseGlob("templates/*.html"))

// funcão princial do projeto bntv
func main(){
	http.HandleFunc("/", index) 
	http.ListenAndServe(":8000", nil)


}

// Funçao de intregação do themes 

func index(w http.ResponseWriter,r *http.Request) {
	temp.ExecuteTemplate(w, "index", nil)

}
