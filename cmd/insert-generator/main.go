package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please enter file name!")
	}
	fileName := os.Args[1]
	sql, err := getSqlFromFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	tables := []table{}
	queries := splitWithoutEmpty(sql, ";")
	for _, query := range queries {
		fmt.Println(query)
		query, err := parseQuery(query)
		if err != nil {
			log.Fatal(err)
		}
		tables = append(tables, *query)
	}
	fmt.Println(tables)
}
