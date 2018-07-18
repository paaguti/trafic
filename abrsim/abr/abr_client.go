package abr

import (
	"net"
	"fmt"
	"log"
	"io"
	// "flag"
	"time"
	"math/rand"
)

func mkTransfer (conn net.Conn, iter int, total int, tsize int) {
	// send to socket
	fmt.Fprintf(conn, fmt.Sprintf("GET %d/%d %d\n", iter, total, tsize))
	// listen for reply
	readBuffer := make([]byte, tsize)
	fmt.Printf("Trying to read %d bytes back...", len(readBuffer))
	readBytes, err := io.ReadFull(conn, readBuffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Effectively read %d bytes\n", readBytes)
}

func Client(host string, port int, iter int, interval int, burst int) {
	// connect to this socket
	serverAddr := fmt.Sprintf("%s:%d",host,port)
	fmt.Printf("Talking to %s\n",serverAddr)
	conn, _ := net.Dial("tcp", serverAddr)

    r := rand.New(rand.NewSource(33))
	initWait := r.Intn(interval * 50) / 50.0
	time.Sleep(time.Duration(initWait) * time.Second)

	currIter := 1
	go mkTransfer (conn , currIter, iter, burst)

	ticker := time.Tick(time.Duration(interval) * time.Second)
	for now := range ticker {
		// read in input from stdin

		currIter ++
		if currIter > iter {
			break
		}
		fmt.Printf("Launching at %v\n", now)
		go mkTransfer (conn , currIter, iter, burst)
	}
	conn.Close()
	fmt.Printf("\nFinished...\n\n")
}

// func main() {
// 	ipPtr    := flag.String("ip", "127.0.0.1", "The IP address to bind to")
// 	portPtr  := flag.Int("port", 8081, "The port to use")
// 	burstPtr := flag.Int("burst", 100000, "The size of the burst")
// 	iterPtr  := flag.Int("iter", 6, "The number of repetitions")
// 	timePtr  := flag.Int("interval", 10, "Wait time between bursts in secs")

// 	flag.Parse()

// 	abrclient(*ipPtr, *portPtr, *iterPtr, *timePtr, *burstPtr)
// }
