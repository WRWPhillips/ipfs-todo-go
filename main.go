package main

import (
	"fmt"

	"github.com/WRWPhillips/ipfs-todo-go/internal"
	"github.com/redis/go-redis/v9"
	// "github.com/ipfs-todo-go/cmd"
	// "github.com/gonuts/commander"
)

func main() {

	cid, err := internal.GetCID()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}
	fmt.Println(cid)

	//the cid of the todo should be consistently stored here, and we can get the cid on launch from it.
}
