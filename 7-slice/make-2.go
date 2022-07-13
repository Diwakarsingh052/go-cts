package main

import "fmt"

func main() {
	var data = make([]int, 0, 100000) // len, cap
	var count int
	lastCap := cap(data)

	for r := 0; r < 1000000; r++ {
		data = append(data, r)

		if lastCap != cap(data) {
			count++

			capCh := float64(cap(data)-lastCap) / float64(lastCap) * 100

			lastCap = cap(data)
			fmt.Printf("Add %p Cap %d CapCh %v\n", data, cap(data), capCh)
		}

	}

	fmt.Println(count)
}
