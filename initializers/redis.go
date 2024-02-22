package initializers

import (
	"os"

	"github.com/redis/go-redis/v9"
)

// Initialization of the redis client itself
func Redis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"),
		Password: "", // no password set
        DB:		  0,  // use default DB
	})
}