package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/riddle", getRiddle)

	listenPort := os.Getenv("HTTP_PORT")
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
	io.WriteString(w, "<p>Hello, Golang!</p>")
}

func getRiddle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /riddle request\n")
	io.WriteString(w, "<p><i>What is always in front of you but can't be seen?</i><br/><br/>...</p>")
}
