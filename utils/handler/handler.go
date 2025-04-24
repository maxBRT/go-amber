package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/spf13/afero"
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
		filepath.Join(baseDir, "content", "index.md"): `# Welcome to My Website

Welcome to the homepage of my static website. Here, you can find information about my projects, blog posts, and more.

## Features
- **Simple**: A clean and minimalistic layout.
- **Static**: No servers, just fast-loading static pages.
- **Flexible**: Easily customizable content.

## About Me
I'm passionate about web development and love creating beautiful, fast, and accessible websites.

## Explore
- [Projects](projects): View my work and experiments.
- [Blog](blog): Read my thoughts and tutorials.

---`,
		filepath.Join(baseDir, "content", "blog", "index.md"):        "# My Articles",
		filepath.Join(baseDir, "content", "blog", "post1.md"):        "# Post 1\n\nThis is the first blog post.",
		filepath.Join(baseDir, "content", "blog", "post2.md"):        "# Post 2\n\nThis is the second blog post.",
		filepath.Join(baseDir, "content", "projects", "index.md"):    "# My Projects",
		filepath.Join(baseDir, "content", "projects", "project1.md"): "# Project 1\n\nDetails about project 1.",
		filepath.Join(baseDir, "content", "projects", "project2.md"): "# Project 2\n\nDetails about project 2.",
		filepath.Join(baseDir, "templates", "base.html"): `<!-- base.html -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
</head>
<body>
    <header>
        <nav>
            <!-- Your navbar content -->
            <a href="/">Home</a>
            <a href="/blog">Blog</a>
            <a href="/projects">Projects</a>
        </nav>
    </header>
    <main>
        {{.Content}}
		{{.List}}
    </main>
    <footer>
        <!-- Your footer content -->
        <p>&copy; 2025 My Website</p>
    </footer>
</body>
</html>
`,
	}

	for path, content := range contentFiles {
		err := afero.WriteFile(fs, path, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", path, err)
			continue
		}
		fmt.Printf("Created file: %s\n", path)
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
	err := afero.WriteFile(fs, serveFilePath, []byte(serveFileContent), 0644)
	if err != nil {
		fmt.Printf("Failed to create serve.go: %v\n", err)
		return err
	}
	fmt.Printf("Created file: %s\n", serveFilePath)
	return nil
}

func MarkdownToHtml(cmd Command) error {
	fs := afero.NewOsFs()
	inputIndexPath := filepath.Join("content")
	outputIndexPath := filepath.Join("output")
	inputProjectIndexFile := filepath.Join("content", "projects", "index.md")
	inputBlogIndexFile := filepath.Join("content", "blog", "index.md")
	inputProjectsPath := filepath.Join("content", "projects")
	outputProjectsPath := filepath.Join("output", "projects")
	inputBlogPath := filepath.Join("content", "blog")
	outputBlogPath := filepath.Join("output", "blog")

	// Process the projects directory
	if err := processContent(fs, inputProjectsPath, outputProjectsPath); err != nil {
		fmt.Printf("Error processing projects directory: %v\n", err)
		return err
	}
	// Process the blog directory
	if err := processContent(fs, inputBlogPath, outputBlogPath); err != nil {
		fmt.Printf("Error processing blog directory: %v\n", err)
		return err
	}

	// Process the index file
	if err := processContent(fs, inputIndexPath, outputIndexPath); err != nil {
		fmt.Printf("Error processing index directory: %v\n", err)
		return err
	}

	// Process the Projects index file
	if err := processIndex(fs, inputProjectIndexFile, outputProjectsPath); err != nil {
		fmt.Printf("Error processing index file: %v\n", err)
		return err
	}
	// Process the Blog index file
	if err := processIndex(fs, inputBlogIndexFile, outputBlogPath); err != nil {
		fmt.Printf("Error processing index file: %v\n", err)
		return err
	}

	return nil
}

func generateFileListHTML(files []os.FileInfo) string {
	var buffer bytes.Buffer
	for _, file := range files {
		if file.Name() == "index.html" {
			continue // Skip the index.html file
		}
		if !file.IsDir() { // Skip directories
			// Start the list item
			buffer.WriteString(`<li><a href="`)
			// Add the file path to the href attribute
			buffer.WriteString(filepath.Join(file.Name()))
			buffer.WriteString(`">`)
			// Add the file name as the link text
			buffer.WriteString(strings.TrimSuffix(file.Name(), ".html"))
			buffer.WriteString(`</a></li>`)
		}
	}
	return buffer.String() // Return the complete HTML string
}

func processIndex(fs afero.Fs, inputFile, outputPath string) error {
	// Load the base template
	baseHtmlPath := filepath.Join("templates", "base.html")
	baseTemplate, err := template.ParseFiles(baseHtmlPath)
	if err != nil {
		return fmt.Errorf("failed to load base template: %w", err)
	}
	outputFilePath := filepath.Join(outputPath, "index.html")
	// Read all files in the input directory
	files, err := afero.ReadDir(fs, outputPath)
	if err != nil {
		return fmt.Errorf("failed to read input directory: %w", err)
	}

	// Generate HTML for the file list
	fileListHTML := generateFileListHTML(files)

	// Read Markdown file
	mdContent, err := afero.ReadFile(fs, inputFile)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", inputFile, err)
	}
	// Convert Markdown to HTML
	htmlContent := markdown.ToHTML(mdContent, nil, nil)

	// Prepare data for the template
	data := struct {
		Title   string
		Content template.HTML
		List    template.HTML
	}{
		Title:   "List",
		Content: template.HTML(htmlContent),
		List:    template.HTML(fileListHTML),
	}

	// Apply the template
	var outputBuffer bytes.Buffer
	if err := baseTemplate.Execute(&outputBuffer, data); err != nil {
		return fmt.Errorf("failed to execute template for file %s: %w", inputFile, err)
	}

	// Write HTML to output file
	if err := afero.WriteFile(fs, outputFilePath, outputBuffer.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", outputPath, err)
	}

	return nil
}

func processContent(fs afero.Fs, inputPath, outputPath string) error {
	// Load the base template
	baseHtmlPath := filepath.Join("templates", "base.html")
	baseTemplate, err := template.ParseFiles(baseHtmlPath)
	if err != nil {
		return fmt.Errorf("failed to load base template: %w", err)
	}

	// Read all files in the input directory
	files, err := afero.ReadDir(fs, inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input directory: %w", err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			inputFilePath := filepath.Join(inputPath, file.Name())
			outputFilePath := filepath.Join(outputPath, strings.TrimSuffix(file.Name(), ".md")+".html")

			// Read Markdown file
			mdContent, err := afero.ReadFile(fs, inputFilePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", inputFilePath, err)
			}

			// Convert Markdown to HTML
			htmlContent := markdown.ToHTML(mdContent, nil, nil)

			// Prepare data for the template
			data := struct {
				Title   string
				Content template.HTML
				List    string
			}{
				Title:   strings.TrimSuffix(file.Name(), ".md"),
				Content: template.HTML(htmlContent),
				List:    "",
			}

			// Apply the template
			var outputBuffer bytes.Buffer
			if err := baseTemplate.Execute(&outputBuffer, data); err != nil {
				return fmt.Errorf("failed to execute template for file %s: %w", inputFilePath, err)
			}

			// Write HTML to output file
			if err := afero.WriteFile(fs, outputFilePath, outputBuffer.Bytes(), 0644); err != nil {
				return fmt.Errorf("failed to write file %s: %w", outputFilePath, err)
			}
		}
	}

	return nil
}
