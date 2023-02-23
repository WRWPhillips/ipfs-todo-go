package internal

import(
	"fmt"
	"os"
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/joho/godotenv"
)


var ctx = context.Background()

func GetCID() string, err {
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

	return rdb.Get(ctx, "foo")

	//the cid of the todo should be consistently stored here, and we can get the cid on launch from it. 
}