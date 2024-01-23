package main

import (
	"fmt"
	"github.com/magnuswahlstrand/sql-exercises/functions/app"
	"time"
)

func main() {
	before := time.Now()
	a := app.SetupApp()
	fmt.Println("Setup took", time.Since(before))
	a.Listen(":3000")
}
