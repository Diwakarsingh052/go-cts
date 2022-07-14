package main

import "fmt"

func main() {

	//a := 20
	//b := a
	//b = 100
	//fmt.Println(a)
	//fmt.Println(b)

	x := 20
	update(&x)
	fmt.Println("from main", &x)

}

//update func which makes the copy of the var of another func and updates it
//so no changes would be reflected back as all the changes are local to this func
//func update(i int) {
//	i = 40
//	fmt.Println("from update", i, &i)
//}

func update(i *int) {
	*i = 40
	fmt.Println("from update", i)
}
