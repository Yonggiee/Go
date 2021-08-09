package mysql

import (
	"fmt"
	"reflect"
)

type Dao struct {
	fields map[string]string
}

func newDao(t interface{}) Dao {
	v := reflect.TypeOf(t)
	dao := Dao{make(map[string]string)}

	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Field(i).Name
		dao.fields[fieldName] = ""
	}
	return dao
}

func ReadAll(t interface{}) ([]*map[string]string, error) {
	tableName := reflect.TypeOf(t).Name()
	db, err := initDB()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT * FROM %s;", tableName)
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	daoArr := extractFromRows(t, rows)

	return parseToMapArr(daoArr), nil
}

// TODO
// func create() {}
// func update() {}
// func delete() {}
