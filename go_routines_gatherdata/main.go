package main

import (
	"io"
	"log"
	"net/http"
	"sync"
)

func GetBook(wg *sync.WaitGroup, c chan any, job int) {
	defer wg.Done()
	resp, err := http.Get("http://localhost:9999/")
	if err != nil {
		log.Fatalf("Could not Read endpoint %+v", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	c <- bodyString

}

func main() {
	ch := make(chan any)
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go GetBook(wg, ch, i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		log.Println(i)
	}

}
