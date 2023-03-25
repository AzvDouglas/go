package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	key := ctx.Value("key")
	fmt.Println(key)
}
