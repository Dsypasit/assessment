package expense

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB
var err error

type DBError struct {
	error
	s      string
	status int
}

type db_temp struct {
	DB
	db *sql.DB
}

type DB interface {
	GetAll() ([]Expense, error)
	GetByID(id int) (Expense, error)
	Update(id int, ex Expense) error
	Create(ex *Expense) error
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

func (db db_temp) Update(id int, ex Expense) error {
	st, err := db.db.Prepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1")
	if err != nil {
		return errors.New("can't prepare statement")
	}

	result, err := st.Exec(id, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))
	if err != nil {
		return errors.New("can't update information")
	}

	if rowAffected, err := result.RowsAffected(); err != nil {
		return errors.New("can't get row affect")
	} else if rowAffected == 0 {
		return errors.New("id missmatch")
	}
	return nil
}

func (db db_temp) Create(ex *Expense) error {
	st, err := db.db.Prepare("INSERT INTO expenses (title, amount, note, tags) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return errors.New("can't prepare statement")
	}

	row := st.QueryRow(&ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))

	if err := row.Scan(&ex.ID); err != nil {
		return errors.New("can't insert information")
	}
	return nil
}
