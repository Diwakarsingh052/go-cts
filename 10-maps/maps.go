package main

import "fmt"

func main() {

	dictionary := map[string]string{ // map[key]value {}
		"up":   "above",
		"down": "below",
	}
	fmt.Println(dictionary["up"])

	dictionary["abc"] = "xyz" // if the key doesn't exist then value would be added to the map otherwise it would update the value at the key
	dictionary["up"] = "egf"  // this is update here as the key already exist in the map

	fmt.Println(dictionary)

	//var users map[int]string // default is nil

	users := make(map[int]string) // providing memory to the map // only required to do once

	users[100] = "raj"
	fmt.Printf("%v\n", users)

	for k, v := range dictionary {
		fmt.Println(k, v)
	}

}
