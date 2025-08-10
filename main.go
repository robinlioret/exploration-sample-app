package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/riddle", getRiddle)
	http.HandleFunc("/redis-test", getRedis)

	listenPort := os.Getenv("HTTP_PORT")
	fmt.Printf("starting on port %s...\n", listenPort)

	err := http.ListenAndServe(fmt.Sprintf(":%s", listenPort), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")

	io.WriteString(w, "<p>Hello, Golang !</p>")
}

func getRiddle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /riddle request\n")
	io.WriteString(w, "<p><i>What is always in front of you but can't be seen?</i><br/><br/>...</p>")
}

func getRedis(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /redis-test request\n")
	io.WriteString(w, "<h1>Redis Test</h2>")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "valkey-primary.valkey.svc.cluster.local:6379",
		Password: "QAFtPNNhoe", // no password set
		DB:       0,            // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		msg := "Cannot find Valkey at valkey-primary.valkey.svc.cluster.local:6379"
		fmt.Println(msg)
		io.WriteString(w, msg)
		return
	}

	msg := "Found Valkey at valkey-primary.valkey.svc.cluster.local:6379"
	fmt.Println(msg)
	io.WriteString(w, msg)

	io.WriteString(w, "<br/>I'm running this piece of code!")
}
