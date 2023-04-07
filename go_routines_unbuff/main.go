package main

import (
	"log"
)

func main() {
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go goroutines(ch)
	}
	//Using for range throws a deadlock error, need to check ok, if not break
	// and close channel. Otherwise loop through the exact number of goroutines deployed above.

	for i := range ch {
		_, ok := <-ch
		if !ok {
			return
		}
		log.Println(i)
	}
	close(ch)
}
func goroutines(ch chan string) {
	log.Println("working....")
	ch <- "goroutine done..."

}
