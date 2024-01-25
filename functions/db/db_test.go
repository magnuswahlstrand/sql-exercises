package db

import (
	"github.com/magnuswahlstrand/sql-exercises/functions/exercises"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExercises(t *testing.T) {
	db := prepareDB()
	defer db.Close()

	for _, tc := range exercises.Exercises {
		tc := tc
		t.Run(tc.ID, func(t *testing.T) {
			t.Parallel()
			_, records, err := Query(db, tc.AnswerQuery)
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

	_, records, err := Query(db, `SELECT * FROM employees`)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 8, len(records))
}
