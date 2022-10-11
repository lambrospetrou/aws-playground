package redisdb

import (
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
)

func NewNative() *redis.Client {
	redisUrl := strings.TrimSpace(os.Getenv("UPSTASH_REDIS_URL"))
	if redisUrl == "" {
		log.Fatalln("Required env UPSTASH_REDIS_URL not set!")
	}
	opt, _ := redis.ParseURL(redisUrl)
	redisDb := redis.NewClient(opt)

	return redisDb
}
