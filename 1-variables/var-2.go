package main

import (
	"fmt"
)

func main() {

	var i = 10
	i++
	i *= 2 // i = i*2
	//++i // not in go
	i = 20
	x := "hello"
	_, _ = i, x // ignore values // don't use in real code
	var (
		name  = "ajay"
		class int
		pass  bool
	)
	_, _, _ = name, class, pass
	const PI = 3.14
	//os.OpenFile()
	//os.O_WRONLY

	fmt.Println()

}
