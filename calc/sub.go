package calc

import "fmt"

// sub func can be used anywhere inside the calc package
func sub() {
	fmt.Println("i am a sub func from calc package")
	fmt.Println(Sum)
	//fmt.Println() // check code of this func
}
