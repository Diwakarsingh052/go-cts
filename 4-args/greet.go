package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	greet()
	fmt.Println("main func ending")
}
func greet() {
	details := os.Args[1:] // from the first index till the end of the slice
	fmt.Println(details, "len = ", len(details))
	if len(details) < 2 {
		log.Println("please provide your name, age, total marks")
		return // stop the exec of the current func // note :- it doesn't stop the program
		//log.Fatalln() // quits the program not recommended
	}
	name := details[0]
	ageString := details[1]
	//access the mark here from the list
	//var err error // default val is nil which indicates no error
	//fmt.Println(err)

	age, err := strconv.Atoi(ageString)
	if err != nil {
		log.Println("please provide a valid age", err)
		return
	}

	fmt.Println(name, age) // ,marks
}

// checking for success
//if err == nil {
//		log.Println("converted", age, err)
//	}

//checking for error
//if err != nil {
//		log.Println("issue", err)
//	}
