package db

import "database/sql"

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

	_, err = db.Exec(createEmployeesTableQuery)
	if err != nil {
		panic(err)
	}
	return db
}

func Query(db *sql.DB, query string) ([][]any, error) {
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
