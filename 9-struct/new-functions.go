package main

import (
	"database/sql"
	"fmt"
	"log"
)

type config struct {
	dbConn string
	data   string
}

var Db string

func NewConfig(dbConn string) (config, error) {
	if dbConn == "" {
		return config{}, sql.ErrConnDone
	}
	c := config{
		dbConn: dbConn,
	}

	return c, nil

}

func (x config) read() {
	if x.dbConn == "" {
		log.Println("not allowed to read")
		return
	}
	fmt.Println(x.data)
}

func (x *config) update(data string) {
	if x.dbConn == "" {
		log.Println("not allowed to update")
		return
	}
	x.data = data
}

func main() {
	//var x config
	con, err := NewConfig("localhost:mysql")
	if err != nil {
		log.Println(err)
	}
	con.update("some data")
	con.read()
	//x.update("some names are stored")
	//x.read()

}
