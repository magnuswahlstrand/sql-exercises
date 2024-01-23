package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var createEmployeesTableQuery = `
CREATE TABLE employees (
    last_name VARCHAR(255) NOT NULL,
    salary INTEGER NOT NULL,
    department VARCHAR(255) NOT NULL
);

INSERT INTO employees (last_name, salary, department) VALUES
    ('Larsson', 48000, 'Accounting'),
    ('Bergstrom', 52000, 'Sales'),
    ('Hakansson', 46000, 'Marketing'),
    ('Svensson', 39000, 'Accounting'),
    ('Lindberg', 56000, 'Sales'),
    ('Nyström', 58000, 'Sales'),
    ('Holm', 43000, 'IT'),
    ('Engström', 50000, 'Marketing');
`

func prepareDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	if _, err = db.Exec(createEmployeesTableQuery); err != nil {
		panic(err)
	}

	// Check that the table was created
	res, err := db.Query("SELECT * FROM employees")
	if err != nil {
		panic(err)
	}
	defer res.Close()

	return db
}

func Query(db *sql.DB, query string) ([]string, [][]any, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, nil, err
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
			return nil, nil, err
		}
		records = append(records, columnValues)
	}

	return cols, records, nil
}

type Checker struct {
	DB *sql.DB
}

type Response struct {
	Success bool
	Headers []string
	Rows    [][]any
}

func (c *Checker) Check(exerciseId string, query string) (*Response, error) {
	exercise, found := exercisesMap[exerciseId]
	if !found {
		return nil, errors.New("exercise not found")
	}

	headers, records, err := Query(c.DB, query)
	if err != nil {
		return nil, err
	}

	response := &Response{
		Success: isCorrect(exercise, records),
		Headers: headers,
		Rows:    records,
	}

	return response, nil
}

func isCorrect(exercise Exercise, records [][]any) bool {
	if len(exercise.Correct) != len(records) {
		return false
	}
	for i := range exercise.Correct {
		if len(exercise.Correct[i]) != len(records[i]) {
			return false
		}
		for j := range exercise.Correct[i] {
			if exercise.Correct[i][j] != records[i][j] {
				fmt.Println(i, j, exercise.Correct[i][j], records[i][j])
				return false
			}
		}
	}
	return true
}

func NewChecker() *Checker {
	db := prepareDB()
	return &Checker{DB: db}
}
