package database

import "fmt"

type configMysql struct {
	db string
}

func NewMysql(conn string) *configMysql {
	return &configMysql{db: conn}
}

type configDb interface {
	read() string
	update() string
	close()
}
type Config struct {
	DB configDb // we can store any struct which impl the configDb interface // with that we would be able to call the methods of that type
	i  int
}

func (c Config) Login() {
	fmt.Println("we are checking for the password from the db")
	fmt.Println(c.DB.read()) // read method would be called according to the type passed in the config struct

}
func (c Config) UpdateDetails() {
	fmt.Println(c.DB.update())
}

func (c *configMysql) read() string { // impl read method of the interface
	return "reading from mysql"
}

func (c *configMysql) update() string {
	return "updating in mysql"
}

func (c *configMysql) close() {
	c.db = ""
	fmt.Println("closing mysql")
}
