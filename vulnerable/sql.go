package vulnerable

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	_ "github.com/mattn/go-sqlite3"
)

const tables = `
CREATE TABLE IF NOT EXISTS user (
   id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
   name  text NOT NULL,
   email text NOT NULL,
   password text NOT NULL
);
CREATE TABLE IF NOT EXISTS product (
   id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
   name  text NOT NULL,
   category  text NOT NULL,
   price  int NOT NULL
);
`

func PrepareSQLDB(nbEntries int) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		log.Fatalln("unexpected sql.Open error:", err)
	}

	if _, err := db.Exec(tables); err != nil {
		return nil, err
	}

	for i := 0; i < nbEntries; i++ {
		_, err := db.Exec(
			"INSERT INTO user (name, email, password) VALUES (?, ?, ?)",
			fmt.Sprintf("User %d", i),
			fmt.Sprintf("user%d@mail.com", i),
			fmt.Sprintf("secret password %d", i))
		if err != nil {
			return nil, err
		}

		_, err = db.Exec(
			"INSERT INTO product (name, category, price) VALUES (?, ?, ?)",
			fmt.Sprintf("Product %d", i),
			"sneaker",
			rand.Intn(500))
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

type Product struct {
	Id       int
	Name     string
	Category string
	Price    string
}

func GetProducts(ctx context.Context, db *sql.DB, category string) ([]Product, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM product WHERE category='"+category+"'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Category, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
