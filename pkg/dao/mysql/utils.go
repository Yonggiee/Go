package mysql

import (
	"database/sql"
)

func extractFromRows(t interface{}, rows *sql.Rows) []*Dao {
	var daoArr []*Dao
	for rows.Next() {
		cols, _ := rows.Columns()
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)

		dao := newDao(t)
		for i, colName := range cols {
			dao.fields[colName] = columns[i]
		}

		daoArr = append(daoArr, &dao)
	}

	return daoArr
}

func parseToMapArr(daoArr []*Dao) []*map[string]string {
	var daoMapArr []*map[string]string

	for _, dao := range daoArr {
		daoMapArr = append(daoMapArr, &dao.fields)
	}
	return daoMapArr
}
