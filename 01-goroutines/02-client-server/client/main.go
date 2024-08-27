package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	// Connect to the server on localhost port 8000
	clientCON, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer clientCON.Close()

	// Create a buffered reader for user input (from stdin)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Connected to the server. Type a message to send:")
	// Read user input from the terminal
	fmt.Print(">")
	inputUser, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Send the user input to the server

		_, err = clientCON.Write([]byte(inputUser))
		if err != nil {
			log.Fatal(err)
		}

		// Receive the server's response
		serverResponse, err := bufio.NewReader(clientCON).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Print the server's response to the terminal

		fmt.Print("Server: ", serverResponse)
	}

}

// package main

// import (
// 	"io"
// 	"log"
// 	"net"
// 	"os"
// )

// func main() {
// 	//  connect to server on localhost port 8000

// 	//  conect to server
// 	clientConn, err := net.Dial("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer clientConn.Close()

// mustCopy(os.Stdout ,clientConn )

// }
// // mustCopy - utility function
// func mustCopy(dst io.Writer, src io.Reader) {
// 	if _, err := io.Copy(dst, src); err != nil {
// 		log.Fatal(err)
// 	}
// }
