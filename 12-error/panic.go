package main

import (
	"fmt"
	"runtime/debug"
)

func main() {

	searchData()

	showData("hello this is a sample func")
	fmt.Println("end of the main func")
}

func showData(data string) {
	fmt.Println(data)
}

func searchData() {
	defer recoverFunc()
	var i []int
	i[19] = 100
	fmt.Println(i)

}

func recoverFunc() {
	r := recover() // nil means no panic // otherwise r would be the msg of the panic

	if r != nil {
		fmt.Println("recovered from the panic", r)
		fmt.Println(string(debug.Stack()))
	}
}
