package main

import (
	"fmt"
	"github.com/maxBRT/go-amber/internal/handler"
	"os"
)

func checkArgs() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
}

// runCommandEntered parses the command-line arguments into a Command structure
// and passes it to the command handler for execution.
// Exits with status code 1 if the command execution fails.
func runCommandEntered(commands handler.Commands) {
	// Create a Command from command-line arguments
	cmdEntered := handler.Command{
		Name: os.Args[1],  // First argument is the command name
		Args: os.Args[2:], // Remaining arguments are passed to the command
	}

	// Execute the command and handle any errors
	err := commands.Run(cmdEntered)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	checkArgs()

	// Register available commands
	cmd := handler.Commands{}

	cmd.Register("generate", handler.Generate)
	cmd.Register("parse", handler.ParseToHtml)
	cmd.Register("new", handler.NewFile)

	runCommandEntered(cmd)

}
