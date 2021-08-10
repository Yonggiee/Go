package main

import (
	"fmt"
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
		panic(err)
	}

	// Open DB
	db, err := initTempDb()
	if err != nil {
		panic(err)
	}

	// Create SQL tables
	tables := []table{}
	for _, query := range *splitWithoutEmpty(sql, ";") {
		fmt.Println(query)
		table, err := parseQuery(query, db)
		if err != nil {
			panic(err)
		}
		tables = append(tables, *table)
	}

	// Print generated inserts
	// mock := generateInserts(&tables)
	// fmt.Println(mock)
}
