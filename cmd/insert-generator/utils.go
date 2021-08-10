package main

import (
	"bufio"
	"os"
	"strings"
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
