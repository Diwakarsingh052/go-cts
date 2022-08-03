package main

import (
	"context"
	"fmt"
)

type ctxKey int

const K ctxKey = 1 // use this to fetch req id

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, K, "1001") // storing the value in the context
	FetchValue(ctx)

}

func FetchValue(ctx context.Context) {
	v := ctx.Value(K) // this will fetch the value out of the context

	//if v == nil {
	//}
	s, ok := v.(string)

	if !ok {
		fmt.Println("value not there or of a different type")
		return
	}

	fmt.Println(s)

}
