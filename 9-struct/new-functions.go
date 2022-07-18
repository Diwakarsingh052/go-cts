package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type config struct {
	dbConn string
	data   string
}

func NewConfig(con string) (*config, error) {
	if con == "" {
		return nil, sql.ErrConnDone // if not config in return type is not pointer // config{} , err
	}
	//c := config{
	//	dbConn: dbConn,
	//}

	return &config{
		dbConn: con,
	}, nil

}

func (x *config) read() {
	if x.dbConn == "" {
		log.Println("not allowed to read")
		return
	}
	fmt.Println(x.data)
}

// update method receiver is a pointer which means any changes done inside here would be reflected back to the caller
func (x *config) update(data string) {
	if x.dbConn == "" {
		log.Println("you don't have a db connection so you are not allowed to update")
		return
	}
	x.data = data
}

func main() {
	l := log.New(os.Stdout, "sales-app ", log.LstdFlags)
	l.Println("this is a main func")

	//var x config
	con, err := NewConfig("localhost:mysql")
	if err != nil {
		log.Println(err)
	}
	con.update("some data")
	con.read()
	//x.update("some names are stored")
	//x.read()
	LogData(l)

}

func LogData(l *log.Logger) {
	l.Println("inside log data")
}
