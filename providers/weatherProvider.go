package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WeatherProvider struct {
	Url string
}

var client http.Client

var weatherProvider WeatherProvider;
func init()  {
	c := http.Client{
		Timeout: 3 * time.Second,
	}
	client = c
	weatherProvider = WeatherProvider{
		Url: "http://localhost:8082/data"}
}

func (WeatherProvider) GetData(request map[string] interface{}) map[string] interface{} {
	fmt.Println("making api call for")
	resp, err := client.Get(weatherProvider.Url)
	if err != nil {
		return map[string]interface{}{}
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	resp.Body.Close()
	return result
}
