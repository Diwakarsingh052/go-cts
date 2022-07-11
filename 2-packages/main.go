package main

import (
	"go-cts/calc/data"
)

func main() {

	//calc.TotalSum()
	//calc.sum = 988
	//fmt.Println("main", calc.sum)
	data.DB = "mysql"
	data.GetData()
	//rand.Float64()

}
