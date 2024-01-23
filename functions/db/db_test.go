package db

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExercises(t *testing.T) {
	db := prepareDB()
	defer db.Close()

	for _, tc := range exercises {
		t.Run(tc.Test.Name, func(t *testing.T) {
			records, err := Query(db, tc.Test.Query)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, len(tc.Correct), len(records))
			for i := range tc.Correct {
				assert.EqualValues(t, fmt.Sprint(tc.Correct[i]), fmt.Sprint(records[i]))
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
