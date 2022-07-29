package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	PrintData(ctx, "hello this is a test func")

}

func PrintData(ctx context.Context, input string) { // ctx should be the first param in func

	select {
	case <-ctx.Done(): // it will tell when a ctx is cancelled or finished
		fmt.Println("timeout")
		fmt.Println(ctx.Err()) // to check err inside the context

	case <-time.After(2 * time.Second): // run the code after specified amt of time
		fmt.Println(input)

	}
}
