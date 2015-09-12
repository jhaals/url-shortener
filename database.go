package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Entry from database
type Entry struct {
	id    int64
	url   string
	views int64
}

// Database interface
type Database interface {
	Get(id int) (Entry, error)
	Save(url string) (int64, error)
}

type sqlite struct {
	Path string
}

func (s sqlite) Save(url string) (int64, error) {
	db, err := sql.Open("sqlite3", s.Path)
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into urls(url) values(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(url)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	tx.Commit()

	return id, nil
}

func (s sqlite) Get(id int) (Entry, error) {
	return Entry{}, nil
}

func (s sqlite) Init() {
	c, err := sql.Open("sqlite3", s.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	sqlStmt := `create table if not exists urls (id integer not null primary key, url text);`
	_, err = c.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}
