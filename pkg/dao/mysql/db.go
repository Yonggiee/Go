package mysql

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() (*sql.DB, error) {
	os.Setenv("DATABASE_LINK", "root@/startit")
	l := os.Getenv("DATABASE_LINK")
	db, err := sql.Open("mysql", l)
	if err != nil {
		return nil, err
	}
	return db, nil
}
