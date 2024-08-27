package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//os.Stdin, you're telling your program to read data from the terminal where the program is running.
	// This is useful when you want to get input directly from the user, such as when the user types a command or a string.

	// Create a buffered reader that reads from standard input (the terminal)

	reader := bufio.NewReader(os.Stdin)
	// Prompt the user for input
	fmt.Print("Enter your name: ")

	// Read the input until a newline character is encountered
	name, _ := reader.ReadString('\n')

	// // Print the user's input
	// fmt.Println("Hello, " + name)

	// os.Stdout, you're telling your program to write data to the terminal.
	// This is useful for displaying results, messages, or logs to the user.
	fmt.Fprintln(os.Stdout, "hello "+name)

}
