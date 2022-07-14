package main

import (
	"log"
	"os"
)

func main() {

	f, err := os.OpenFile("test.txt", os.O_RDONLY, 0666)
	// 4 read // 2 write  // 1 exec // octal perms in linux
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	//ioutil.ReadFile() // deprecated
}
