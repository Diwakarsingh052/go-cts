package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type con struct {
	*sql.DB
	data string
}

var ErrNotFound = errors.New("record not found")
var ErrNOConnection = errors.New("no connection to the database")

func (c con) searchRecord() error {
	db, err := sql.Open("postgres", "postgres") // don't use this pattern in prod code
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrNOConnection)
	}
	_, err = db.Exec("select  * from students") // don't use this pattern in prod code
	if err != nil {
		return fmt.Errorf("%v %w", err, ErrNotFound)
	}

	return nil

}

func main() {
	// errors.As //interfaces
	var c con
	err := c.searchRecord()
	fmt.Println(err)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("create the record first then search for it")
		return
	}

	if errors.Is(err, ErrNOConnection) {
		fmt.Println("create a valid connection")
		return
	}

}
