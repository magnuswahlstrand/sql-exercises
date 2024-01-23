package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExercises(t *testing.T) {
	db := prepareDB()
	defer db.Close()

	for _, tc := range exercises {
		t.Run(tc.Test.Name, func(t *testing.T) {
			_, records, err := Query(db, tc.Test.Query)
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

	records, err := Query(db, `SELECT * FROM employees`)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 8, len(records))
}
