package redis

import (
	"fmt"
	"time"
)

func SaveData(key string, value string, expiry int) {
	fmt.Println(key)
	fmt.Println(value)
	client.Set(key,value,time.Duration(expiry) * time.Second)
}

func GetData(key string) (string, error) {
	val, err := client.Get(key).Result()
	return val,err
}
