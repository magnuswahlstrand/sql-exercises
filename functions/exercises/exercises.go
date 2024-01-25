package exercises

type Exercise struct {
	ID             string
	CorrectHeaders []string
	Correct        [][]any
	Description    string
	AnswerQuery    string
	Title          string
}

type ExtenderExercise struct {
	Exercise
	Previous string
	Next     string
}

var Exercises = []Exercise{
	{
		Title:       "Select everything!",
		ID:          "select_all",
		Description: "Select all columns from the employees table.",
		CorrectHeaders: []string{
			"last_name",
			"salary",
			"department",
		},
		Correct: [][]any{
			{"Larsson", int64(48000), "Accounting"},
			{"Bergstrom", int64(52000), "Sales"},
			{"Hakansson", int64(46000), "Marketing"},
			{"Svensson", int64(39000), "Accounting"},
			{"Lindberg", int64(56000), "Sales"},
			{"Nyström", int64(58000), "Sales"},
			{"Holm", int64(43000), "IT"},
			{"Engström", int64(50000), "Marketing"},
		},
		AnswerQuery: "SELECT * FROM employees",
	},

	{
		ID:          "select_columns",
		Description: "Select last_name and department from the employees table.",
		CorrectHeaders: []string{
			"last_name",
			"department",
		},
		Correct: [][]any{
			{"Larsson", "Accounting"},
			{"Bergstrom", "Sales"},
			{"Hakansson", "Marketing"},
			{"Svensson", "Accounting"},
			{"Lindberg", "Sales"},
			{"Nyström", "Sales"},
			{"Holm", "IT"},
			{"Engström", "Marketing"},
		},
		AnswerQuery: "SELECT last_name, department FROM employees",
	},
	{
		ID:          "select_order_by",
		Description: "Select all columns from the employees table, ordered by salary in descending order.",
		CorrectHeaders: []string{
			"last_name",
			"salary",
			"department",
		},
		Correct: [][]any{
			{"Nyström", 58000, "Sales"},
			{"Lindberg", 56000, "Sales"},
			{"Bergstrom", 52000, "Sales"},
			{"Engström", 50000, "Marketing"},
			{"Larsson", 48000, "Accounting"},
			{"Hakansson", 46000, "Marketing"},
			{"Holm", 43000, "IT"},
			{"Svensson", 39000, "Accounting"},
		},
		AnswerQuery: "SELECT * FROM employees ORDER BY salary DESC",
	},
}

var ExercisesMap map[string]ExtenderExercise

func init() {
	ExercisesMap = make(map[string]ExtenderExercise)
	for i, e := range Exercises {
		previous, next := "", ""
		if i > 0 {
			previous = Exercises[i-1].ID
		}
		if i < len(Exercises)-1 {
			next = Exercises[i+1].ID
		}

		ExercisesMap[e.ID] = ExtenderExercise{
			Exercise: e,
			Previous: previous,
			Next:     next,
		}
	}
}
