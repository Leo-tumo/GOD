package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const num = 5

/*
1. context.Background() - only in the highest lever e.i. main()
2. context.TODO - when you are not sure which kind of context to use
3. context.Value - this should be rarely used and it should take only not necessary arguments
4. ctx is always the first argument in function
5. only context created function can cancel the context
*/
func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second)

	ctx = context.WithValue(ctx, "id", 1)

	parse(ctx)
}

func parse(ctx context.Context) {
 
	id := ctx.Value("id")
	fmt.Println(id.(float32))
	client := http.DefaultClient
	_ = client
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parsing completed ")
		case <-ctx.Done():
			fmt.Println("deadline exceeded")
			return
		}
	}
}
