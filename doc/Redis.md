# example of cache with redis

## Golang

1. Init connection

```go
rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

```

2. set/get a value

```go

	err := rdb.Set("key", "value", time.Minute()*time).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
```

3. get/set list

```go

lentgh, err := rdb.LPush("mylist", []string{"test", "2"}).Result()

	if err != nil {
		panic(err)
	}

	fmt.Println("length", lentgh)

```