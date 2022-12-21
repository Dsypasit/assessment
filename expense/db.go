package expense

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

type db_temp struct {
	DB
	db *sql.DB
}

type DB interface {
	GetAll() ([]Expense, error)
	GetByID(id int) (Expense, error)
	Update(id int) (Expense, error)
	Create() (Expense, error)
}

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

func InitDBTemp() DB {
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

	return db_temp{db: db}
}

func (db db_temp) GetAll() ([]Expense, error) {
	st, err := db.db.Prepare("SELECT * FROM expenses")
	if err != nil {
		return nil, errors.New("can't prepare db statement")
	}

	rows, err := st.Query()
	if err != nil {
		return nil, errors.New("can't query information")
	}

	var expenses []Expense
	for rows.Next() {
		var ex Expense
		err := rows.Scan(&ex.ID, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, ex)
	}

	return expenses, nil
}

func (db db_temp) GetByID(id int) (Expense, error) {
	st, err := db.db.Prepare("SELECT * FROM expenses WHERE id=$1")
	if err != nil {
		return Expense{}, errors.New("can't prepare db statement")
	}

	var expense Expense
	err = st.QueryRow(id).Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
	if err != nil {
		return Expense{}, errors.New("can't query information")
	}

	return expense, nil
}
