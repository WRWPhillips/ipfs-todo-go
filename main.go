package main

import (
	"fmt"
	"os"
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/joho/godotenv"
	// "github.com/ipfs-todo-go/cmd"
	// "github.com/gonuts/commander"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load(".env")
	
	if err != nil {
		panic(err)
	}

	redis_uri := os.Getenv("REDIS_URI")
	fmt.Println("url ->", redis_uri)

	opt, err := redis.ParseURL(redis_uri)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)

	get := rdb.Get(ctx, "foo")
	fmt.Println(get.Val(), get.Err())
}
