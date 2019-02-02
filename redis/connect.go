package redis

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

var client *redis.Client

var Config Configuration

type Configuration struct {
	RedisUrl      string
	RedisPassword string
	LogFile       string
}

func init() {
	configFile, _ := os.Open("config.json")
	err := json.NewDecoder(configFile).Decode(&Config)
	if err != nil {
		fmt.Println("error:", err)
	}

	file, err := os.OpenFile(Config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file", file, ":", err)
	}
	c := createNewClient(Config.RedisUrl, Config.RedisPassword)
	client = c
}

func createNewClient(url string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		// no password se
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
	// Output: PONG <nil>
}
