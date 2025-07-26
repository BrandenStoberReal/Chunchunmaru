package utilities

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

type SqlTable struct {
	Name    string
	Columns []string
	Values  []interface{}
}

func OpenDatabase(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func CloseDatabase(db *sql.DB) {
	defer db.Close()
}

func CreateTable(db *sql.DB, table *SqlTable) (sql.Result, error) {
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS `" + table.Name + "` (" + strings.Join(table.Columns, ", ") + ")")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func DeleteTable(db *sql.DB, table *SqlTable) (sql.Result, error) {
	result, err := db.Exec("DROP TABLE IF EXISTS `" + table.Name + "`")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func InsertIntoTable(db *sql.DB, table *SqlTable, values []string) (sql.Result, error) {
	result, err := db.Exec("insert into " + table.Name + " (" + strings.Join(table.Columns, ", ") + ") values(" + strings.Join(values, ", ") + ")")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func DeleteFromTable(db *sql.DB, table *SqlTable) (sql.Result, error) {
	result, err := db.Exec("delete from " + table.Name)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func FetchValuesFromColumn(db *sql.DB, table *SqlTable, column string) ([]string, error) {
	rows, err := db.Query("select " + column + " from " + table.Name)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var values []string
	for rows.Next() {
		var value string
		scanerr := rows.Scan(&value)
		if scanerr != nil {
			log.Fatal(err)
			return nil, scanerr
		}
		values = append(values, value)
	}
	return values, nil
}
