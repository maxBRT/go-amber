package handler

import (
	"fmt"
)

// Command{Name: "login", Args: []string{"username"}}
type Command struct {
	Name string   // The name of the Command (e.g., "login")
	Args []string // Arguments passed to the Command
}

// Commands implements the Command registry and execution system.
// It provides a flexible way to register and execute CLI Commands
// while maintaining clean separation of concerns.
type Commands struct {
	// Map of Command names to their handler functions
	// This allows for dynamic dispatch of Commands
	cmdList map[string]func(Command) error
}

// Register adds a new Command handler to the registry.
// Parameters:
//   - name: The Command name to register
//   - f: The handler function for the Command
//
// The handler function receives the application state and Command arguments.
func (c *Commands) Register(name string, f func(Command) error) {
	// Initialize the map if this is the first registration
	// This ensures we don't need a separate initialization step
	if c.cmdList == nil {
		c.cmdList = make(map[string]func(Command) error)
	}
	c.cmdList[name] = f
}

// Run executes a Command if it exists in the registered Commands.
// It looks up the Command by name and invokes the corresponding handler.
// Returns any error from the Command execution.
func (c *Commands) Run(cmd Command) error {
	if f, ok := c.cmdList[cmd.Name]; ok {
		return f(cmd)
	}
	// Provide feedback when a Command is not recognized
	fmt.Println("Command not found.")
	return nil
}
