package handler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

func NewFile(cmd Command) error {

	fs := afero.NewOsFs()
	// Check if the command has enough arguments
	if len(cmd.Args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	switch cmd.Args[0] {
	case "blog":
		// Create a new blog post
		outputPath := filepath.Join("content", "blog", cmd.Args[1]+".md")
		err := createNewPost(fs, basePost, outputPath)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", outputPath, err)
			return err
		}
	case "project":
		// Create a new project
		outputPath := filepath.Join("content", "projects", cmd.Args[1]+".md")
		err := createNewPost(fs, baseProject, outputPath)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", outputPath, err)
			return err
		}
	}

	return nil
}

func createNewPost(fs afero.Fs, file, outputPath string) error {
	err := afero.WriteFile(fs, outputPath, []byte(file), 0644)
	if err != nil {
		fmt.Printf("Failed to create file %s: %v\n", outputPath, err)
		return err
	}
	fmt.Printf("Created file: %s\n", outputPath)
	return nil
}
