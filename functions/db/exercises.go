package db

type Exercise struct {
	Name    string
	Correct [][]any
	Test    Test
}

type Test struct {
	Name  string
	Query string
}

var exercises = []Exercise{
	{
		Name: "select all",
		Correct: [][]any{
			{1, "Larsson", 48000, "Accounting"},
			{2, "Bergstrom", 52000, "Sales"},
			{3, "Hakansson", 46000, "Marketing"},
			{4, "Svensson", 39000, "Accounting"},
			{5, "Lindberg", 56000, "Sales"},
			{6, "Nyström", 58000, "Sales"},
			{7, "Holm", 43000, "IT"},
			{8, "Engström", 50000, "Marketing"},
		},
		Test: Test{
			Name:  "Exercise 1",
			Query: "SELECT * FROM employees",
		},
	},
	{
		Name: "select all",
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
		Test: Test{
			Name:  "Exercise 2",
			Query: "SELECT last_name, department FROM employees",
		},
	},
	{
		Name: "select all",
		Correct: [][]any{
			{6, "Nyström", 58000, "Sales"},
			{5, "Lindberg", 56000, "Sales"},
			{2, "Bergstrom", 52000, "Sales"},
			{8, "Engström", 50000, "Marketing"},
			{1, "Larsson", 48000, "Accounting"},
			{3, "Hakansson", 46000, "Marketing"},
			{7, "Holm", 43000, "IT"},
			{4, "Svensson", 39000, "Accounting"},
		},
		Test: Test{
			Name:  "Exercise 3",
			Query: "SELECT * FROM employees ORDER BY salary DESC",
		},
	},
}
