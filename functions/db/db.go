package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var createEmployeesTableQuery = `
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    last_name VARCHAR(255) NOT NULL,
    salary INTEGER NOT NULL,
    department VARCHAR(255) NOT NULL
);

INSERT INTO employees (id, last_name, salary, department) VALUES
    (1, 'Larsson', 48000, 'Accounting'),
    (2, 'Bergstrom', 52000, 'Sales'),
    (3, 'Hakansson', 46000, 'Marketing'),
    (4, 'Svensson', 39000, 'Accounting'),
    (5, 'Lindberg', 56000, 'Sales'),
    (6, 'Nyström', 58000, 'Sales'),
    (7, 'Holm', 43000, 'IT'),
    (8, 'Engström', 50000, 'Marketing');
`

func prepareDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	fmt.Println("YEAH")
	res, err := db.Exec(createEmployeesTableQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.RowsAffected())

	res2, err := db.Query("SELECT * FROM employees")
	if err != nil {
		panic(err)
	}
	res2.Close()
	fmt.Println(res2.Columns())

	res3, err := db.Query("SELECT * FROM employees")
	if err != nil {
		panic(err)
	}
	res3.Close()
	fmt.Println(res3.Columns())

	res4, err := db.Query("SELECT * FROM employees")
	if err != nil {
		panic(err)
	}
	res4.Close()
	fmt.Println(res4.Columns())

	return db
}

func Query(db *sql.DB, query string) ([][]any, error) {
	//res2, err := db.Query("SELECT * FROM employees")
	//if err == nil {
	//	fmt.Println(res2.Columns())
	//}
	//res2.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	numCols := len(cols)

	var records [][]any

	for rows.Next() {
		// Create a slice of interfaces to hold the column values
		columnValues := make([]interface{}, numCols)
		columnPointers := make([]interface{}, numCols)
		for i := range columnValues {
			columnPointers[i] = &columnValues[i]
		}

		// Scan the result into the column pointers
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}
		records = append(records, columnValues)
	}

	return records, nil
}

type Checker struct {
	DB *sql.DB
}

func (c *Checker) Check(exerciseId string, query string) (bool, error) {
	exercise, found := exercisesMap[exerciseId]
	if !found {
		return false, errors.New("exercise not found")
	}

	//res2, err := c.DB.Query("SELECT * FROM employees")
	//if err != nil {
	//	return false, err
	//}
	//defer res2.Close()

	records, err := Query(c.DB, query)
	if err != nil {
		return false, err
	}
	fmt.Println("YEAH4")

	fmt.Println(exercise.Correct)
	fmt.Println(records)
	//for i := range exercise.Correct {
	//	if !equal(exercise.Correct[i], records[i]) {
	//		return false, nil
	//	}
	//}

	return true, nil
}

func NewChecker() *Checker {
	db := prepareDB()
	res4, err := db.Query("SELECT * FROM employees")
	if err != nil {
		panic(err)
	}
	res4.Close()

	return &Checker{DB: db}
}
