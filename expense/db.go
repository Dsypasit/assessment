package expense

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func InitDB() {
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	createTB := `
CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`

	_, err = db.Exec(createTB)
	if err != nil {
		log.Fatal("can't create table", err)
	}

	log.Println("Ok")
}
