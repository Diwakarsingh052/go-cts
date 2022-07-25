package database

import "fmt"

type configPostgres struct {
	db string
}

func NewConfigPostgres(conn string) *configPostgres {
	return &configPostgres{db: conn}
}

func (c *configPostgres) read() string {
	return "reading from postgres"
}

func (c *configPostgres) update() string {
	return "updating in postgres"
}

func (c *configPostgres) close() {
	c.db = ""
	fmt.Println("closing postgres")
}
