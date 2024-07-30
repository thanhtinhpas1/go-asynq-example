package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
)

func main() {
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/", // RootPath specifies the root for asynqmon app
		RedisConnOpt: asynq.RedisClientOpt{Addr: ":6379"},
	})

	r := mux.NewRouter()
	r.PathPrefix(h.RootPath()).Handler(h)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8081",
	}

	log.Println("Dashboard is running at: http://localhost:8081")

	// Go to http://localhost:8080/monitoring to see asynqmon homepage.
	log.Fatal(srv.ListenAndServe())
}
