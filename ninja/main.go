package main

import (
	"context"
	"fmt"
	"time"
)

const num = 5

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	go func() {
		time.Sleep(time.Millisecond * 100)
		cancel()
	}()
	defer cancel()
	parse(ctx)
}

func parse(ctx context.Context) {
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
