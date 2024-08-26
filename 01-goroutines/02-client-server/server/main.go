package main

import (
	"io"
	"log"
	"net"
)

func main() {

	//  write server program to handle concurrent client connections.

	// 1 - start create server
	tcpListener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

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

	//  io.ReadAll to read all data from the connection
	data, err := io.ReadAll(clientCon)
	if err != nil {
		log.Println("Error reading data:", err)
		return
	}

	//  log recived data from client
	log.Printf("Received data %v", string(data))

	//  Create Respose to client
	response := "Response from server: " + string(data)
	// Write the response back to the client
	_, err = clientCon.Write([]byte(response))
	if err != nil {
		log.Println("Error writing response:", err)
		return
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
