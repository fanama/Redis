package cache

import (
	"fmt"

	"github.com/go-redis/redis"
)

// var ctx context.Context

func init() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set("running", "true", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get("running").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	err = rdb.Set("running", "false", 0).Err()
	if err != nil {
		panic(err)
	}

}
