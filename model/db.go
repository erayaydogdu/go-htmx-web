package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
    db *sql.DB
)

func Setup() {
    fmt.Println("Setup DB starting...")
    var err error
    db, err = sql.Open("postgres", "host=localhost port=6432 user=todouser password=todo1234! dbname=todo sslmode=disable")
    
    fmt.Println("Conn Open finised!")

    if err != nil {
        fmt.Println("Could not connect to db", err)
    }
    err = db.Ping()
    if err != nil {
        fmt.Println("Could not ping db", err)
    }

    fmt.Println("DB connection finished!");
}
