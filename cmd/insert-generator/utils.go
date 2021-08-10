package main

import (
	"bufio"
	"database/sql"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func getSqlFromFile(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	sql := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sql += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	sqlLowerCase := strings.ToLower(sql)
	return sqlLowerCase, nil
}

func initTempDb() (*sql.DB, error) {
	name := "test2"
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("USE " + name)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func splitWithoutEmpty(toSplit string, delim string) *[]string {
	stringArr := strings.Split(toSplit, delim)
	var retArr []string
	for _, str := range stringArr {
		if str != "" {
			retArr = append(retArr, str)
		}
	}
	return &retArr
}
