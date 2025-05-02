package handler

import (
	"fmt"
	"os"
	"os/exec"
)

func GenerateRepo(cmd Command) error {
	if len(cmd.Args) < 1 {
		os.Exit(1)
		return fmt.Errorf("not enough arguments")
	}
	cloneCommand := exec.Command("git", "clone", "git@github.com:maxBRT/basic-website.git", cmd.Args[0])
	cloneCommand.Stdout = os.Stdout
	cloneCommand.Stderr = os.Stderr
	return cloneCommand.Run()
}
