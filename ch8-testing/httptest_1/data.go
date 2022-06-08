package main

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		log.Println(err)
		return
	}
}

func Retrieve(id int) (post Post, err error) {
	post = Post{}
	err = DB.QueryRow("SELECT id, content, author from posts where id =$1", post)
}
