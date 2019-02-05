package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	client := http.Client{}
	var keys [20] string
	for i := 0; i < 20; i++ {
		keys[i] = string(int('a') + i)
	}

	i := 0
	for i < 50000 {
		i += 1
		fmt.Println(i)
		message := map[string]interface{}{
			"data": "help"}
		message["key"] = keys[rand.Intn(len(keys))]
		bytesRepresentation, err := json.Marshal(message)
		fmt.Println(message)
		resp, err := client.Post("http://localhost:8081/data", "application/json", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			log.Fatalln(err)
		}

		var result map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&result)

		log.Println(result)
		log.Println(result["data"])

	}
}
