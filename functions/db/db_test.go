package db

import (
	"context"
	"github.com/magnuswahlstrand/sql-exercises/functions/exercises"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExercises(t *testing.T) {
	db := prepareDB()
	defer db.Close()

	for _, tc := range exercises.Exercises {
		tc := tc
		t.Run(tc.ID, func(t *testing.T) {
			_, records, err := Query(context.Background(), db, tc.AnswerQuery)
			if err != nil {
				t.Fatal(err)
			}

			correct := isCorrect(tc, records)
			if !correct {
				t.Errorf("expected %v, got %v", tc.Correct, records)
			}
		})
	}
}

func TestBasic(t *testing.T) {
	db := prepareDB()
	defer db.Close()

	_, records, err := Query(context.Background(), db, `SELECT * FROM employees`)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 8, len(records))
}

func TestDatabaseIsReadonly(t *testing.T) {
	db := prepareDB()
	defer db.Close()

	_, _, err := Query(context.Background(), db, `DROP TABLE employees`)
	if err != nil {
		t.Fatal(err)
	}

	_, records, err := Query(context.Background(), db, `SELECT * FROM employees`)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 8, len(records))
}

func TestTimeout(t *testing.T) {
	db := prepareDB()
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	_, _, err := Query(ctx, db, `SELECT * FROM employees`)
	assert.Equal(t, context.DeadlineExceeded, err)
}
