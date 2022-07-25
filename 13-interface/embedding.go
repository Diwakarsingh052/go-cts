package main

import "fmt"

type reader interface {
	read() int
}
type writer interface {
	write(data string)
}

type readWriter interface {
	reader
	writer
}

type student struct {
	name string
}

func (s student) read() int {
	return 10
}
func (s student) write(data string) {
	fmt.Println(data)
}

func readWrite(rw readWriter) {

	fmt.Println(rw.read())
	rw.write("hello ")
}

func readData(r reader) {

}

func writeData(w writer) {

}

func main() {
	var s student
	//readWrite(s)
	//writeData(s)
	readData(s)

}
