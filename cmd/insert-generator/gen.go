package main

import (
//"fmt"
)

func generateInserts(tables *[]table) string {
	inserts := ""
	for _, table := range *tables {
		inserts += generateInsertsForATable(&table)
	}
	return inserts
}

func generateInsertsForATable(t *table) string {
	table := *t
	//tableName := table.name
	insertCount := table.insertCount

	inserts := ""
	for i := 0; i < insertCount; i++ {
		inserts += ("sad" + "\n")
	}

	return inserts
}
