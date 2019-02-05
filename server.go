package main

import (
	"bitbucket.org/shadowfaxtech/negroni"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-lock-cache/handlers"
	"io"
	"log"
	"net/http"
	"os"
)

type Configuration struct {
	LicenseKey string
	AppName    string
	Dsn        string
	LogFile    string
	Broker     string
	Client     string
}

var Config Configuration

var ServerLogger *log.Logger

var LogFileObj io.Writer

func init() {
	configFile, _ := os.Open("config.json")
	err := json.NewDecoder(configFile).Decode(&Config)
	if err != nil {
		fmt.Println("error:", err)
	}
	file, err := os.OpenFile(Config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", file, ":", err)
	}
	LogFileObj = file
	logger := log.New(file,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	ServerLogger = logger
}

func main() {
	router := mux.NewRouter()
	n := negroni.New()
	n.Use(negroni.NewLogger(LogFileObj))
	router.HandleFunc("/data", handlers.HandleIndex).Methods("POST")
	n.UseHandler(router)
	ServerLogger.Println("Starting the Server")
	http.ListenAndServe("0.0.0.0:8081", n)
}
