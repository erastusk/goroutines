package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
)

func DoRequest() (*http.Response, error) {
	resp, err := http.Get("http://localhost:9999/")
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func GetUrlWithBAckof(r time.Duration) (*http.Response, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Second * r

	main_resp, err := DoRequest()
	if err != nil {
		log.Printf("%+v", err)
	}

	retryFunc := func() error {
		resp, err := DoRequest()
		if err != nil {
			log.Printf("%+v", err)
			return err
		}
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			main_resp = resp
			// 2xx all done

			return nil
		}
		return fmt.Errorf("received http status code %d", resp.StatusCode)

	}
	err = backoff.Retry(retryFunc, bo)

	return main_resp, err

}

func GetBook(wg *sync.WaitGroup, c chan any, r time.Duration) {
	defer wg.Done()
	resp, err := GetUrlWithBAckof(r)
	if err != nil {
		c <- "Could Not Read Endpoint, Failing after Retrying...." + r.String()
		return
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
	RetryMaxtime := time.Second * 25

	wg.Add(1)
	go GetBook(wg, ch, RetryMaxtime)
	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		log.Println(i)
	}

}
