package main

import "go-cts/database"

func main() {

	m := database.NewMysql("localhost")
	c := database.Config{DB: m}
	c.Login()
	c.UpdateDetails()
	p := database.NewConfigPostgres("localhost")
	c = database.Config{DB: p}
	c.Login()
	c.UpdateDetails()
}
