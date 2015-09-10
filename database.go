package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Entry from database
type Entry struct {
	id    int32
	url   string
	views int32
}

// Database interface
type Database interface {
	Get(id int) (Entry, error)
	Save(url string) (int32, error)
}

type sqlite struct {
	Path string
}

func (db sqlite) Save(id int) (Entry, error) {
	return Entry{}, nil
}

func (db sqlite) Get(id int) (Entry, error) {
	return Entry{}, nil
}

func (db sqlite) Init() {
	c, err := sql.Open("sqlite3", db.Path)
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
