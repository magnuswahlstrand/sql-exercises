package exercises

type Exercise struct {
	ID             string
	CorrectHeaders []string
	Correct        [][]any
	Test           Test
}

type Test struct {
	Name  string
	Query string
}

var Exercises = []Exercise{
	{
		ID: "select_all",
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
		Test: Test{
			Name:  "Exercise 1",
			Query: "SELECT * FROM employees",
		},
	},

	//{
	//	ID: "select all",
	//	Correct: [][]any{
	//		{"Larsson", "Accounting"},
	//		{"Bergstrom", "Sales"},
	//		{"Hakansson", "Marketing"},
	//		{"Svensson", "Accounting"},
	//		{"Lindberg", "Sales"},
	//		{"Nyström", "Sales"},
	//		{"Holm", "IT"},
	//		{"Engström", "Marketing"},
	//	},
	//	Test: Test{
	//		Name:  "Exercise 2",
	//		Query: "SELECT last_name, department FROM employees",
	//	},
	//},
	//{
	//	ID: "select all",
	//	Correct: [][]any{
	//		{6, "Nyström", 58000, "Sales"},
	//		{5, "Lindberg", 56000, "Sales"},
	//		{2, "Bergstrom", 52000, "Sales"},
	//		{8, "Engström", 50000, "Marketing"},
	//		{1, "Larsson", 48000, "Accounting"},
	//		{3, "Hakansson", 46000, "Marketing"},
	//		{7, "Holm", 43000, "IT"},
	//		{4, "Svensson", 39000, "Accounting"},
	//	},
	//	Test: Test{
	//		Name:  "Exercise 3",
	//		Query: "SELECT * FROM employees ORDER BY salary DESC",
	//	},
	//},
}

var ExercisesMap = map[string]Exercise{
	Exercises[0].ID: Exercises[0],
}
