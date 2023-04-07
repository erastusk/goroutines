package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

type Book struct {
	Title       string
	Book_Number string
}

func DummyHttpServer(w http.ResponseWriter, r *http.Request) {

	random_book := strconv.Itoa(rand.Intn(100))
	BookName := strings.Join([]string{"Book_Version", random_book}, "_")
	Book := Book{Title: BookName, Book_Number: random_book}
	json.NewEncoder(w).Encode(Book)
}

func main() {
	sigchan := make(chan os.Signal, 1)
	http.HandleFunc("/", DummyHttpServer)

	log.Println("Starting server on port 9999")
	go func() {
		log.Fatal(http.ListenAndServe(":9999", nil))
	}()
	signal.Notify(sigchan, os.Interrupt)
	sig := <-sigchan
	log.Println("Server received a os.Interrupt command terminating.", sig)

}
