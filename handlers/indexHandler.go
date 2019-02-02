package handlers

import (
	"fmt"
	_ "github.com/gorilla/mux"
	"go-lock-cache/redis"
	"net/http"
)

type ResponseStruct struct {
	Message string
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	//ServerLogger.Println("thank god")
	lock, err := redis.AcquireLock("alok", "anand", 100)
	if err != nil {
		fmt.Println("failed to acquire lock")
	} else {
		fmt.Println(lock)
	}
	if lock {
		fmt.Println("trying to release lock")
		redis.ReleaseLock("alok", "nand")
	}
	fmt.Fprint(w, "Welcome! \n")

}
