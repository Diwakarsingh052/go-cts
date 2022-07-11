package calc

import "fmt"

var Sum int

// TotalSum func is exported as the first letter of the func is uppercase
func TotalSum() {
	a, b := 10, 20
	Sum = a + b
	fmt.Println("total Sum", a+b)
	sub()
}

//cannot redeclare a func
//func sub() {
//	fmt.Println("i am a sub func from calc package")
//	fmt.Println(Sum)
//	//fmt.Println() // check code of this func
//}
