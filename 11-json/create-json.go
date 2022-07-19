package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type person struct {
	FirstName string          `json:"first_name"`      // set the field name to the name specified for the json
	Password  string          `json:"-"`               //- means ignore this field in json output
	Perms     map[string]bool `json:"perms,omitempty"` // if the field is null then avoid in json output
}

func main() {

	users := []person{

		{
			FirstName: "Roy",
			Password:  "abc",
			Perms:     map[string]bool{"admin": true},
		},

		{
			FirstName: "Raj",
			Password:  "qwe",
			Perms:     map[string]bool{"write": false},
		},

		{
			FirstName: "Pulkit",
			Password:  "rty",
		},
	}

	//users[0].FirstName

	//jsonData, err := json.Marshal(users)
	jsonData, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(jsonData))
	//err = os.WriteFile("test.json", jsonData, 0666)
	f, err := os.OpenFile("test.json", os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	_, err = f.Write(jsonData)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(users[0].Password)

}
