package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type user struct {
	FirstName string          `json:"first_name"`
	Perms     map[string]bool `json:"perms"`
}

func main() {

	jsonData, err := os.ReadFile("test.json") // reading the file and putting data in jsonData var

	if err != nil {
		log.Fatalln(err)
	}
	var u []user

	err = json.Unmarshal(jsonData, &u) //converting json into the struct
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(u)

	for i, v := range u {
		fmt.Println(i, v)
	}
}
