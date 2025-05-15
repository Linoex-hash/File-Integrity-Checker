package main

import (
	"database/sql"
)

// creates a table if not exists
func create_table(db *sql.DB, table string) (sql.Result, error) {
	sql := `create table if not exists ` + table + ` (
		filepath text primary key not null,
		hash text not null
	);`

	return db.Exec(sql)
}

// Use sql.Stmt instead of db.QueryRow to avoid allocating the same sql queries for every filename
func get_hash_from_filename(hash_getter *sql.Stmt, filename string) (string, error) {
	var hash string
	err := hash_getter.QueryRow(filename).Scan(&hash)
	if err != nil {
		return "", err
	}

	return hash, nil

}

func modify_hash_in_table(modifier *sql.Stmt, filename string, hash string) error {
	_, err := modifier.Exec(filename, hash)
	return err
}

// Attempts to create prepared statements used by this program to get, add and set sql statements
func get_prepared_statements(db *sql.DB) ([3]*sql.Stmt, error) {
	get_hash_sql, err := db.Prepare("select hash from hashes where filepath = ?")
	if err != nil {
		return [3]*sql.Stmt{}, err
	}

	add_hash_sql, err := db.Prepare("insert into hashes(filepath, hash) values(?, ?)")
	if err != nil {
		return [3]*sql.Stmt{}, err
	}

	update_hash_sql, err := db.Prepare("update hashes set hash = ? where filepath = ?")
	if err != nil {
		return [3]*sql.Stmt{}, err
	}

	return [3]*sql.Stmt{get_hash_sql, add_hash_sql, update_hash_sql}, nil
}
