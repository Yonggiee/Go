package main

import (
	"database/sql"
	"errors"
	// "fmt"
	"strconv"
	"strings"
)

func parseQuery(query string, db *sql.DB) (*table, error) {
	if !checkCreateTableSyntax(query) {
		return nil, errors.New("Invalid table syntax")
	}
	tableName := strings.Fields(query)[2]
	stripCreateSql, insertCount, err := getTableDetails(query)
	if err != nil {
		return nil, err
	}
	colsSql := stripCreateSql[1 : len(stripCreateSql)-1]

	t := table{name: tableName, insertCount: insertCount}
	for _, colSql := range *splitWithoutEmpty(colsSql, ",") {
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
	isMatch := CREATETABLEREGEX.MatchString(query)
	if !isMatch {
		return false
	}
	return true
}

func getTableDetails(query string) (string, int, error) {
	start := strings.Index(query, "(")
	end := strings.LastIndex(query, ")")
	insertCount, err := strconv.Atoi(strings.TrimSpace(query[end+1:]))
	if err != nil {
		return "", 0, err
	}
	return query[start:end], insertCount, nil
}
