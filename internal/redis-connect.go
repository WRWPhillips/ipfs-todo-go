package internal

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func GetCID() (string, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	redis_uri := os.Getenv("REDIS_URI")

	opt, err := redis.ParseURL(redis_uri)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)

	return rdb.Do(ctx, "get", "foo").Text()

	//the cid of the todo should be consistently stored here, and we can get the cid on launch from it.
}
