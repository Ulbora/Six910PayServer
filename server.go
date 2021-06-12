package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	port := "3010"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	fmt.Println("Six910Pay Server (six nine ten) is running on port " + port + "!")

	//l.LogLevel = lg.OffLevel

	http.ListenAndServe(":"+port, router)

}
