package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "panel.gohtml", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})

	r.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		type data struct {
			l5        []person
			a5        []person
			last10    []person
			aC        int
			lC        int
			thisMonth int
		}

		var d data
		d.l5 = getPeople(" WHERE degree = 'L' ORDER BY points DESC LIMIT 5")
		d.a5 = getPeople(" WHERE degree = 'M' ORDER BY points DESC LIMIT 5")
		d.last10 = getPeople(" ORDER By logs.date DESC, logs.time desc LIMIT 10")
		d.aC = getQueryInt("SELECT COUNT(DISTINCT(card)) FROM `people` WHERE degree = 'M'")
		d.lC = getQueryInt("SELECT COUNT(DISTINCT(card)) FROM `people` WHERE degree = 'L'")

		//todo
		d.thisMonth = getQueryInt("")

		err := tpl.ExecuteTemplate(w, "home.gohtml", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})

	r.HandleFunc("/api/card/{card_id}/{passwd}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cardId := vars["card_id"]

		fmt.Fprintf(w, "Your card id is: %s", cardId)
	})

	http.ListenAndServe(":8080", r)

}
