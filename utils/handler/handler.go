package handler

import (
	"fmt"
	"github.com/spf13/afero"
	"os"
	"path/filepath"
)

func Generate(cmd Command) error {
	fs := afero.NewOsFs()
	// Check if the command has enough arguments
	if len(cmd.Args) < 1 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	baseDir := cmd.Args[0]

	// Define the directory structure
	directories := []string{
		filepath.Join(baseDir, "content", "blog"),
		filepath.Join(baseDir, "content", "projects"),
		filepath.Join(baseDir, "static", "css"),
		filepath.Join(baseDir, "static", "js"),
		filepath.Join(baseDir, "static", "images"),
		filepath.Join(baseDir, "templates"),
		filepath.Join(baseDir, "output"),
		filepath.Join(baseDir, "output", "blog"),
		filepath.Join(baseDir, "output", "projects"),
	}

	// Create the directories
	for _, dir := range directories {
		err := fs.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Printf("Failed to create directory %s: %v\n", dir, err)
			continue
		}
		fmt.Printf("Created directory: %s\n", dir)
	}

	// Write initial content files
	contentFiles := map[string]string{
		filepath.Join(baseDir, "content", "blog", "post1.md"):        "# Post 1\n\nThis is the first blog post.",
		filepath.Join(baseDir, "content", "blog", "post2.md"):        "# Post 2\n\nThis is the second blog post.",
		filepath.Join(baseDir, "content", "projects", "project1.md"): "# Project 1\n\nDetails about project 1.",
		filepath.Join(baseDir, "content", "projects", "project2.md"): "# Project 2\n\nDetails about project 2.",
	}

	for path, content := range contentFiles {
		err := afero.WriteFile(fs, path, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", path, err)
			continue
		}
		fmt.Printf("Created file: %s\n", path)
	}
	// Copy default index
	sourceIndexPath := "./assets/index.md"
	targetIndexPath := filepath.Join(baseDir, "content", "index.md")

	indexContent, err := afero.ReadFile(fs, sourceIndexPath)
	if err != nil {
		fmt.Printf("Failed to read default index.md: %v\n", err)
		return err
	}

	err = afero.WriteFile(fs, targetIndexPath, indexContent, 0644)
	if err != nil {
		fmt.Printf("Failed to create index.md: %v\n", err)
		return err
	}
	serveFilePath := filepath.Join(baseDir, "serve.go")
	serveFileContent := `package main

import (
	"log"
	"net/http"
)

func main() {
	// Define the directory to serve
	fs := http.FileServer(http.Dir("./output"))
	http.Handle("/", fs)

	log.Println("Serving on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
`
	err = afero.WriteFile(fs, serveFilePath, []byte(serveFileContent), 0644)
	if err != nil {
		fmt.Printf("Failed to create serve.go: %v\n", err)
		return err
	}
	fmt.Printf("Created file: %s\n", serveFilePath)
	return nil
}
