package main

import (
	"fmt"
	"log"
)

func main() {
	ch := make(chan string, 10)

	for i := 0; i < 10; i++ {
		go goroutines(ch)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	close(ch)
}
func goroutines(ch chan string) {
	log.Println("working....")
	ch <- "goroutine done..."

}
