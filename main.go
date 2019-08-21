package main

import (
	"fmt"
	"framework/config"
	"framework/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.HTTPPort),
		Handler:        router,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
