package main

import (
	"errors"
	// "fmt"
	"strings"
)

func parseQuery(query string) (*table, error) {
	if !checkCreateTableSyntax(query) {
		return nil, errors.New("Invalid table syntax")
	}
	stripCreateSql := innerTableRegex.FindStringSubmatch(query)[0]
	colsSql := stripCreateSql[1 : len(stripCreateSql)-1]

	t := table{}
	for _, colSql := range splitWithoutEmpty(colsSql, ",") {
		wordArr := strings.Fields(colSql)
		if len(wordArr) < 2 {
			return nil, errors.New("Invalid column syntax")
		}
		colName := wordArr[0]
		typ := wordArr[1]
		column := column{name: colName, typ: typ}
		t.cols = append(t.cols, column)
	}
	return &t, nil
}

func checkCreateTableSyntax(query string) bool {
	isMatch := createTableRegex.MatchString(query)
	if !isMatch {
		return false
	}
	return true
}
