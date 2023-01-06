package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Expense struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Amount int      `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}
type Err struct {
	Message string `json:"message"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error")
	}
	defer db.Close()
	createTb := `
	CREATE TABLE IF NOT EXISTS expenses ( id SERIAL PRIMARY KEY, title TEXT,amount FLOAT,note TEXT,tags TEXT[]);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}
	fmt.Println("create table success")

	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
}
