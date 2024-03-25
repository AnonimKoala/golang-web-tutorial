package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type person struct {
	FirstName string
	LastName  string
	Card      string
	Points    int
}

type activity struct {
	firstName string
	lastName  string
	time      string
	date      string
}

func getPeople(param string) []person {
	db, err := sql.Open("mysql", "log:123@(localhost:3306)/logmachine?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	query := `SELECT people.first_name, people.last_name,people.Card, people.Points FROM people`
	if param != "" {
		query += param
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var people []person
	for rows.Next() {
		var p person

		err := rows.Scan(&p.FirstName, &p.LastName, &p.Card, &p.Points)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, p)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return people
}

func getActivities(param string) []activity {
	db, err := sql.Open("mysql", "log:123@(localhost:3306)/logmachine?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	query := `SELECT people.first_name, people.last_name, logs.time, logs.date FROM logs,people WHERE people.Card = logs.Card`
	if param != "" {
		query += param
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var activities []activity
	for rows.Next() {
		var a activity

		err := rows.Scan(&a.firstName, &a.lastName, &a.time, &a.date)
		if err != nil {
			log.Fatal(err)
		}
		activities = append(activities, a)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return activities
}

func getQueryInt(query string) int {
	db, err := sql.Open("mysql", "log:123@(localhost:3306)/logmachine?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var num int

	if err := db.QueryRow(query).Scan(&num); err != nil {
		log.Fatal(err)
	}

	return num
}

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "panel.gohtml", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})

	r.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		type data struct {
			L5        []person
			A5        []person
			LAST10    []person
			AC        int
			LC        int
			ThisMonth int
		}

		var d data
		d.L5 = getPeople(" WHERE degree = 'L' ORDER BY Points DESC LIMIT 5")
		d.A5 = getPeople(" WHERE degree = 'M' ORDER BY Points DESC LIMIT 5")
		d.LAST10 = getPeople(" JOIN logs ON people.Card = logs.Card ORDER BY logs.date DESC, logs.time DESC LIMIT 10")
		d.AC = getQueryInt("SELECT COUNT(DISTINCT(people.Card)) FROM `people` WHERE people.degree = 'M'")
		d.LC = getQueryInt("SELECT COUNT(DISTINCT(people.Card)) FROM `people` WHERE people.degree = 'L'")

		d.ThisMonth = getQueryInt("select COUNT(*) from logs where  logs.date > CURRENT_DATE - INTERVAL 1 MONTH")

		err := tpl.ExecuteTemplate(w, "home.gohtml", d)
		if err != nil {
			log.Fatalln(err)
		}
	})

	r.HandleFunc("/api/Card/{card_id}/{passwd}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cardId := vars["card_id"]

		fmt.Fprintf(w, "Your Card id is: %s", cardId)
	})

	http.ListenAndServe(":8080", r)

}
