/*
---> Start listen Server
--------> Start Listening
------------> Accept Connection with channel - go routine
---------------> Close connection, continue listening.
*/

package main

import (
	"log"
	"net"
)

func startserver() {
	log.Println("sTarted server on port 9999")
	ch := make(chan any)
	defer close(ch)
	func() {
		lst, err := net.Listen(
			"tcp",
			":9999",
		)
		if err != nil {
			log.Fatal(err)
		}
		go Acceptedconn(lst)

		<-ch
	}()

}

func Acceptedconn(l net.Listener) {
	for {
		con, err := l.Accept()
		log.Println("connection accpeted from ", con.RemoteAddr())
		if err != nil {
			log.Fatalf("Unable to accept connections")

		}
		go ReadMessages(con)
	}
}

func ReadMessages(c net.Conn) {
	for {
		buff := make([]byte, 2048)
		m, err := c.Read(buff)
		if err != nil {
			log.Printf("Connection closed by %+v Goodbye....", c.RemoteAddr())
			return
		}
		log.Printf("Returning received message from %+v, %s", c.RemoteAddr(), string(buff[:m]))
	}

}

func main() {
	startserver()

}
