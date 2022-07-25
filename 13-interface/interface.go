package main

import (
	"fmt"
)

// Polymorphism means that a piece of code changes its behavior depending on the
//concrete data it’s operating on // Tom Kurtz, Basic inventor

// "Don’t design with interfaces, discover them". - Rob Pike

// Bigger the interface weaker the abstraction // Rob Pike

//interfaces are abstract type // it does not store anything of their own
type reader interface {
	read(b []byte) (int, error) // methods signature
	//abc()  // all the methods should be impl over the struct to use the interface
}

type file struct {
	name string
}

func (f file) read(b []byte) (int, error) {
	s := "hello go dev"
	fmt.Println("inside file read")
	copy(b, s)
	return len(b), nil
}

func (f file) del() {

}

type jsonObject struct {
	location string
}

func (j jsonObject) read(b []byte) (int, error) {
	s := `{name:"abc"}`
	fmt.Println("inside json read")
	copy(b, s)
	return len(s), nil
}

//
//func FetchData(f file) {
//	data := make([]byte, 50)
//	f.read(data)
//	fmt.Println(string(data))
//}
//
//func FetchDataJson(j jsonObject) {
//	json := make([]byte, 50)
//	j.read(json)
//	fmt.Println(string(json))
//}

// Fetch will accept any type of value that implements the reader interface
func Fetch(r reader) {
	data := make([]byte, 50)
	fmt.Printf("%T\n", r)
	r.read(data) // read would be called according to the concrete data (struct) passed to it
	fmt.Println(string(data))

}

func main() {
	fi := file{name: "abc.txt"}
	j := jsonObject{location: "localhost:8080"}
	//FetchData(fi)
	//FetchDataJson(j)
	//var r reader // interface default values is nil which means that it is not storing any concrete type
	var r reader = fi

	b := make([]byte, 100)
	r.read(b)

	//r.del() // we can't call del method here as this method is not the part of the reader interface signature
	Fetch(fi)
	Fetch(j)
	//f, _ := os.Open("abc.txt")
	//log.New(os.Stdout, "", 0)
	//log.New(f, "", 0)

}
