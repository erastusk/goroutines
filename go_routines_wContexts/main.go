package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/context"
)

func GetBook(c chan any, ctx context.Context) {
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
	ctx := context.Background()
	timeout := time.Duration(time.Microsecond * 4)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	start := time.Now()
	go GetBook(ch, ctx)
	log.Println("Took:", time.Since(start))

	select {
	case m := <-ch: //Response from called function/Goroutine
		log.Println("Go routine returned a response within the Timelimit")
		log.Println(m)
	case <-ctx.Done(): // ctx.Done is read/unblocks when timeout is exceeded.
		log.Printf("%+v. Context deadline =  %+v", ctx.Err().Error(), timeout)
	}

}
