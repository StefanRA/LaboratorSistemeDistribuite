package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strconv"
	"fmt"
	"strings"
)

const numberOfSeats int64 = 10
var seats [numberOfSeats]int

func SocketServer(port int) {
	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	defer listen.Close()
	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
		os.Exit(1)
	}
	log.Printf("Begin listen port: %d", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
		message = strings.TrimSuffix(message, "\n")
		seatNumber, _ := strconv.ParseInt(message, 10,32)
		seatNumber--
		if seatNumber >= numberOfSeats {
			conn.Write([]byte("Locul cu acel numar nu exista!\n"))
		} else {
			if seats[seatNumber] == 0 {
				seats[seatNumber] = 1
				conn.Write([]byte("Locul a fost rezervat pentru dumneavoastra!\n"))
			} else {
				conn.Write([]byte("Locul este deja rezervat!\n"))
			}
		}
	}
}

func main() {
	port := 3333
	SocketServer(port)
}