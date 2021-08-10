package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Read SQL from file
	if len(os.Args) < 2 {
		panic("Please enter file name!")
	}
	fileName := os.Args[1]
	sql, err := getSqlFromFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Create SQL tables
	tables := []table{}
	for _, query := range *splitWithoutEmpty(sql, ";") {
		fmt.Println(query)
		table, err := parseQuery(query)
		if err != nil {
			log.Fatal(err)
		}
		tables = append(tables, *table)
	}

	// Print generated inserts
	mock := generateInserts(&tables)
	fmt.Println(mock)
}
