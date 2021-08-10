package main

import (
	"fmt"
)

func generateInserts(tables *[]table) string {
	for _, table := range *tables {
		fmt.Println(table)
	}
	return ""
}
