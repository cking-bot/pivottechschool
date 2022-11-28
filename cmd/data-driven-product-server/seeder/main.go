package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "products.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, description, price FROM products")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var name, description, price string
		err = rows.Scan(&id, &name, &description, &price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, description, price)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
