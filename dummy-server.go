package main

import (
	"bitbucket.org/shadowfaxtech/negroni"
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"time"
)

func handleDummy(w http.ResponseWriter, r *http.Request) {
	result := map[string]interface{}{
		"hello": "world",
		"life":  42,
		"embedded": map[string]string{
			"yes": "of course!",
		},
	}
	timer1 := time.NewTimer(time.Duration(rand.Intn(400)) * time.Millisecond)
	<-timer1.C
	js, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	router := mux.NewRouter()
	n := negroni.New()
	router.HandleFunc("/data", handleDummy).Methods("GET")
	n.UseHandler(router)
	http.ListenAndServe("0.0.0.0:8082", n)
}
