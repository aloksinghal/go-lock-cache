package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"go-lock-cache/providers"
	"go-lock-cache/redis"
	"time"
)

func GetWeather(request map[string] interface{}) map[string] interface{} {
	var result  = map[string]interface{}{};
	var signal = make(chan struct{})
	var signal2 = make(chan struct {})
	timer1 := time.NewTimer(1000 * time.Millisecond)
	ticker := time.NewTicker(5 * time.Millisecond)
	go func() {
		for t:= range(ticker.C) {
			cacheData, _ := redis.GetData(request["key"].(string));
			if cacheData != "" {
				json.Unmarshal([]byte (cacheData), &result)
				fmt.Println("serving data from cache")
				close(signal)
			}
			break
			fmt.Println(t)
		}
	}()


	ticker2 := time.NewTicker(20 * time.Millisecond)

	go func() {
		for t := range ticker2.C {
			value, _ := uuid.NewRandom();
			lock, err := redis.AcquireLock(request["key"].(string) + "_lock", value.String(), 10)
			if err != nil {
				fmt.Println("failed to acquire lock")
			}
			if lock {
				fmt.Println("acquired lock, getting from API")
				var provider providers.Provider = providers.WeatherProvider{Url: "https://httpbin.org/get"}
				result = provider.GetData(request)
				js, err := json.Marshal(result)
				if err != nil {
					panic(err)
				}
				redis.SaveData(request["key"].(string), string(js), 10)
				close(signal2)
				redis.ReleaseLock(request["key"].(string), value.String())
				break
			}
			fmt.Println(t)
		}
	}()

	select {
		case <-signal:
			fmt.Println("got from cache first")
			ticker.Stop()
			ticker2.Stop()
			timer1.Stop()
		case <-signal2:
			fmt.Println("got from API first")
			ticker.Stop()
			ticker2.Stop()
			timer1.Stop()
		case <-timer1.C:
			fmt.Println("timedout")
			ticker.Stop()
			ticker2.Stop()
	}
	return result
}
