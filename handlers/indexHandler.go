package handlers

import (
	"encoding/json"
	_ "github.com/gorilla/mux"
	"go-lock-cache/helpers"
	"net/http"
)

type ResponseStruct struct {
	Message string
	Data    map[string]interface{}
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	//ServerLogger.Println("thank god")
	var request map[string]interface{};
	var result map[string]interface{};
	json.NewDecoder(r.Body).Decode(&request)



	result = helpers.GetWeather(request)

	js, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
