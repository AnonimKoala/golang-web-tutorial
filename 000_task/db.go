package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type person struct {
	firstName string
	lastName  string
	card      string
	points    int
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

	query := `SELECT first_name, last_name,card, points FROM people`
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

		err := rows.Scan(&p.firstName, &p.lastName, &p.card, &p.points)
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

	query := `SELECT people.first_name, people.last_name, logs.time, logs.date FROM logs,people WHERE people.card = logs.card`
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

	if err := db.QueryRow(query, 1).Scan(&num); err != nil {
		log.Fatal(err)
	}

	return num
}

func main() {

	//getPeople(db)

	//{ // Insert a new person
	//	firstName := "johndoe"
	//	lastName := "secret"
	//	card := time.Now()
	//
	//	result, err := db.Exec(`INSERT INTO users (firstName, lastName, created_at) VALUES (?, ?, ?)`, firstName, lastName, card)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	id, err := result.LastInsertId()
	//	fmt.Println(id)
	//}

	//{ // Query a single person
	//	var (
	//		id        int
	//		firstName  string
	//		lastName  string
	//		card time.Time
	//	)
	//
	//	query := "SELECT id, firstName, lastName, created_at FROM users WHERE id = ?"
	//	if err := db.QueryRow(query, 1).Scan(&id, &firstName, &lastName, &card); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(id, firstName, lastName, card)
	//}

	{ // Query all people

	}

	//{
	//	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
}
