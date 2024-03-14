package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Row struct {
	Date     string
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
	AdjClose string
}

func createRows(data [][]string) []Row {
	var rows []Row

	for i, line := range data {
		if i == 0 {
			continue
		}

		var row Row
		for j, field := range line {
			if j == 0 {
				row.Date = field
			} else if j == 1 {
				row.Open = field
			} else if j == 2 {
				row.High = field
			} else if j == 3 {
				row.Low = field
			} else if j == 4 {
				row.Close = field
			} else if j == 5 {
				row.Volume = field
			} else if j == 6 {
				row.AdjClose = field
			} else {
				print("Unexpected field!")
			}
		}

		rows = append(rows, row)
	}
	return rows
}

func main() {
	f, err := os.Open("table.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	rows := createRows(data)

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", rows)
	if err != nil {
		log.Fatal(err)
	}

}
