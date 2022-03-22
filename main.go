package main

import (
	"context"
	"flink_chalenge/api"
	"flink_chalenge/configs"
	"flink_chalenge/repository"
	"flink_chalenge/service"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	conf := configs.InitConfig()
	repo := repository.NewRepo()
	locationService := service.NewLocation(&repo)

	locHandler := api.NewRouteHandler(&locationService)

	server := &http.Server{Addr: conf.Port, Handler: locHandler}
	go func() {
		if err := server.ListenAndServe(); err != nil {
		}
	}()
	ch := make(chan os.Signal, 1)
	cleanup := make(chan bool)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		var once sync.Once
		for range ch {
			once.Do(func() {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if err := server.Shutdown(ctx); err != nil {
					//log error
				}
				cleanup <- true
			})
		}
	}()
	<-cleanup
}
