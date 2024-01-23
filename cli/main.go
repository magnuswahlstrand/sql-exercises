package main

import "github.com/magnuswahlstrand/sql-exercises/functions/app"

func main() {
	app.SetupApp().Listen(":3000")
}
