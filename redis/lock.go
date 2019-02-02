package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func AcquireLock(key string, value string, expiry int) (bool, error) {
	set, err := client.SetNX(key, value, time.Duration(expiry)*time.Second).Result()
	if err == redis.Nil {
		err = nil
	}
	return set, err
}

func ReleaseLock(key string, value string) (bool, error) {
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		err = nil
	}
	if val == value {
		client.Del(key)
		fmt.Println("deleted lock")
	}
	return true, nil
}
