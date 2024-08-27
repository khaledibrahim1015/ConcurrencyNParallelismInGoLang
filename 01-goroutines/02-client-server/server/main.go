package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func main() {

	//  write server program to handle concurrent client connections.

	// 1 - start create server
	tcpListener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server start on localhost:9000")

	// 2 - start infinite loop to accept clients connection

	for {
		clientCon, err := tcpListener.Accept()

		if err != nil {
			continue //  to get next client connction
		}

		// Call HandleClientConnection and passed clientConnection to process request and write response
		//  and run HandleClientConnection  in multi goRoutines (threads ) to handle multiple client connectionn
		go HandleClientConnection(clientCon)
	}

}

// HandleClientConnection processes the client request and sends a response
func HandleClientConnection(clientCon net.Conn) {
	defer clientCon.Close()
	reader := bufio.NewReader(clientCon)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("Error reading data:", err)
			}
			return
		}
		data = strings.TrimSpace(data)
		log.Printf("Received data: %s", data)
		response := "Response from server: " + data + "\n"
		time.Sleep(100 * time.Millisecond)
		_, err = clientCon.Write([]byte(response))
		if err != nil {
			log.Println("Error writing response:", err)
			return
		}
	}
}

// // HandleClientConnection processes the client request and sends a response
// func HandleClientConnection(clientCon net.Conn) {
// 	defer clientCon.Close()

// 	// Create a buffered reader to read the client’s request
// 	reader := bufio.NewReader(clientCon)

// 	for {
// 		// Read the client's request
// 		clientRequest, err := reader.ReadString('\n')
// 		if err != nil {
// 			log.Println("Client disconnected:", err)
// 			return // Exit the loop if the client disconnects or an error occurs
// 		}

// 		// Log the received request from the client
// 		log.Printf("Received request: %s", clientRequest)

// 		// Create a response message
// 		response := "Response from server: " + clientRequest

// 		// Write the response back to the client
// 		_, err = clientCon.Write([]byte(response))
// 		if err != nil {
// 			log.Println("Error writing response:", err)
// 			return
// 		}
// 	}
// }
