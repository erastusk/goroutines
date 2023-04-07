package main

import (
	"log"
	"sync"
)

func main() {
	ch := make(chan string)
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go goroutines(ch, wg)
	}
	//Somehow creating a wait and close in a goroutine ensures channel closure occurs gracefully
	//When using an unbuffered channel.Making it save to range over channel.
	//Not required if using a buffered channel. Can use main go routine
	go func() {
		wg.Wait()
		log.Println("All go routines finished executing")
		close(ch)
		log.Println("Closed channel")
	}()
	// A closed channel can be read.
	for i := range ch {
		log.Println(i)
	}
}
func goroutines(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("working....")
	ch <- "goroutine done..."

}
