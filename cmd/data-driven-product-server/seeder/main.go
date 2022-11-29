package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sql.DB
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func main() {
	if err := findRemoveDB(); err != nil {
		log.Fatalf("Error removing DB\n%s", err)
	}

	db, err := createDB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.createTable(db.DB)
	if err != nil {
		log.Fatal()
	}

	jsonProducts, err := openJson()

	if err := seed(jsonProducts, db.DB); err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v", query(db.DB))

}

func findRemoveDB() error {
	if err := os.Remove("products.db"); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return nil
}

func createDB() (Database, error) {
	db, err := sql.Open("sqlite3", "products.db")
	if err != nil {
		return Database{}, err
	}
	fmt.Printf("database created successfully\n%v\n", err)
	return Database{DB: db}, nil
}

func (d *Database) createTable(db *sql.DB) error {
	sqlStmt := `
create table products (id integer not null primary key, name text, description text, price real);
`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	fmt.Printf("table created successfully\n%v\n", err)
	return nil
}

func openJson() ([]byte, error) {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	file, err := os.ReadFile("products.json")
	if err != nil {
		return nil, err
	}
	return file, nil
}

func seed(b []byte, db *sql.DB) error {
	var prods []Product

	if err := json.Unmarshal(b, &prods); err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into products(id, name, description, price) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	for _, p := range prods {
		_, err := stmt.Exec(p.ID, p.Name, p.Description, p.Price)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func query(db *sql.DB) []Product {
	var prod []Product
	rows, err := db.Query("select * from products limit 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var description string
		var price float64
		err = rows.Scan(&id, &name, &description, &price)
		if err != nil {
			log.Fatal(err)
		}

		res := Product{
			ID:          id,
			Name:        name,
			Description: description,
			Price:       price,
		}

		prod = append(prod, res)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return prod
}
