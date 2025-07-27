package utilities

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
)

type SqlTable struct {
	Name    string
	Columns []string
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
	db.Close()
}

func CreateTable(db *sql.DB, table SqlTable) (sql.Result, error) {
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

// Credit to GPT for this function since I hate SQL. I wrote the rest with help from docs.
// FetchSingleValue fetches a single value from a given column for a row identified by the keyColumn/keyValue.
func FetchSingleValue[T any](db *sql.DB, table *SqlTable, column, keyColumn string, keyValue interface{}) (T, error) {
	// Validate column and keyColumn names
	validCol, validKey := false, false
	for _, col := range table.Columns {
		if col == column {
			validCol = true
		}
		if col == keyColumn {
			validKey = true
		}
	}
	if !validCol {
		var zero T
		return zero, fmt.Errorf("invalid column name: %s", column)
	}
	if !validKey {
		var zero T
		return zero, fmt.Errorf("invalid key column name: %s", keyColumn)
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", column, table.Name, keyColumn)
	var value T
	err := db.QueryRow(query, keyValue).Scan(&value)
	return value, err
}

// Also written by GPT, although it really didn't need to be honestly.
// ResetColumnForKey sets a specific column (e.g., "queries") to a value (e.g., 0)
// for the row(s) where keyColumn = keyValue.
func ResetColumnForKey(db *sql.DB, table *SqlTable, targetColumn string, resetValue interface{}, keyColumn string, keyValue interface{}) error {
	// Validate columns exist
	validTarget, validKey := false, false
	for _, col := range table.Columns {
		if col == targetColumn {
			validTarget = true
		}
		if col == keyColumn {
			validKey = true
		}
	}
	if !validTarget {
		return fmt.Errorf("table %s does not have a '%s' column", table.Name, targetColumn)
	}
	if !validKey {
		return fmt.Errorf("table %s does not have a '%s' column", table.Name, keyColumn)
	}

	query := fmt.Sprintf("UPDATE %s SET %s = ? WHERE %s = ?", table.Name, targetColumn, keyColumn)
	_, err := db.Exec(query, resetValue, keyValue)
	return err
}

// Also GPT due to me not understanding SQL or GO well enough to marshal data types between both, although I now understand how the code works.
// FetchAllIpInfo returns a slice of IpInfoStruct, one for each row in the table.
func FetchAllIpInfo(db *sql.DB, table *SqlTable) ([]IpInfoStruct, error) {
	// Validate columns exist
	hasIp, hasQueries, hasAggression := false, false, false
	for _, col := range table.Columns {
		if col == "ip" {
			hasIp = true
		}
		if col == "queries" {
			hasQueries = true
		}
		if col == "aggression" {
			hasAggression = true
		}
	}
	if !hasIp || !hasQueries || !hasAggression {
		return nil, fmt.Errorf("table %s must have 'ip', 'queries' and 'aggression' columns", table.Name)
	}

	query := fmt.Sprintf("SELECT ip, queries, aggression FROM %s", table.Name)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []IpInfoStruct
	for rows.Next() {
		var info IpInfoStruct
		// Marshal the data types
		if err := rows.Scan(&info.Ip, &info.Queries, &info.Aggression); err != nil {
			return nil, err
		}
		results = append(results, info)
	}
	return results, rows.Err()
}

// FetchAllUserAgentInfo Same as its sister function, except for user-agenting.
func FetchAllUserAgentInfo(db *sql.DB, table *SqlTable) ([]UserAgentInfoStruct, error) {
	// Validate columns exist
	hasUa, hasQueries, hasAggression := false, false, false
	for _, col := range table.Columns {
		if col == "useragent" {
			hasUa = true
		}
		if col == "queries" {
			hasQueries = true
		}
		if col == "aggression" {
			hasAggression = true
		}
	}
	if !hasUa || !hasQueries || !hasAggression {
		return nil, fmt.Errorf("table %s must have 'useragent', 'queries', and 'aggression' columns", table.Name)
	}

	query := fmt.Sprintf("SELECT useragent, queries, aggression FROM %s", table.Name)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []UserAgentInfoStruct
	for rows.Next() {
		var info UserAgentInfoStruct
		// Marshal the data types
		if err := rows.Scan(&info.UserAgent, &info.Queries, &info.Aggression); err != nil {
			return nil, err
		}
		results = append(results, info)
	}
	return results, rows.Err()
}

// WipeAllRows deletes all rows from the table but keeps the table structure.
func WipeAllRows(db *sql.DB, table *SqlTable) error {
	query := fmt.Sprintf("DELETE FROM %s", table.Name)
	_, err := db.Exec(query)
	return err
}

// Another GPT function
// UpsertRow performs an upsert (insert or update) for any table described by SqlTable.
// Assumes the first column is the unique key.
func UpsertRow(db *sql.DB, table SqlTable, values []interface{}) error {
	if len(table.Columns) == 0 || len(values) != len(table.Columns) {
		return fmt.Errorf("column and value count mismatch")
	}

	// Build column and placeholder lists
	columns := strings.Join(table.Columns, ", ")
	placeholders := strings.Repeat("?, ", len(table.Columns))
	placeholders = strings.TrimSuffix(placeholders, ", ")

	// Build ON CONFLICT clause for the first column (assumed unique)
	updateSet := []string{}
	for i, col := range table.Columns {
		// Don't update the unique key column
		if i == 0 {
			continue
		}
		updateSet = append(updateSet, fmt.Sprintf("%s = excluded.%s", col, col))
	}
	updateClause := strings.Join(updateSet, ", ")

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) ON CONFLICT(%s) DO UPDATE SET %s;",
		table.Name, columns, placeholders, table.Columns[0], updateClause,
	)

	_, err := db.Exec(query, values...)
	return err
}

// More GPT function
// SumQueries returns the sum of the 'queries' column for all rows in the table.
func SumQueries(db *sql.DB, table *SqlTable) (int, error) {
	// Validate that the table has a 'queries' column
	hasQueries := false
	for _, col := range table.Columns {
		if col == "queries" {
			hasQueries = true
			break
		}
	}
	if !hasQueries {
		return 0, fmt.Errorf("table %s does not have a 'queries' column", table.Name)
	}

	query := fmt.Sprintf("SELECT SUM(queries) FROM %s", table.Name)
	var sum sql.NullInt64
	err := db.QueryRow(query).Scan(&sum)
	if err != nil {
		return 0, err
	}
	if !sum.Valid {
		return 0, nil // No rows in table, so sum is 0
	}
	return int(sum.Int64), nil
}
