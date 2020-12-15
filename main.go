package main



import (
	"context"
	"fmt" 
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ExampleClient() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    err := rdb.Set(ctx, "test", "1234", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "test").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("test", val)

    val2, err := rdb.Get(ctx, "test2").Result()
    if err == redis.Nil {
        fmt.Println("test2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("test2", val2)
    }
    // Output: key value
    // key2 does not exist
}


func main() {

	ExampleClient();

}