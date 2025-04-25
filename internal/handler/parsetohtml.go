package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/spf13/afero"
)

func ParseToHtml(cmd Command) error {
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
	fileListHTML := generateFileListHTML(files, outputPath)

	// Read Markdown file
	mdContent, err := afero.ReadFile(fs, inputFile)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", inputFile, err)
	}
	// Extract metadata from the Markdown file
	mdContent = removeYAMLMetaData(mdContent)

	// Convert Markdown to HTML
	htmlContent := markdown.ToHTML(mdContent, nil, nil)

	// Prepare data for the template
	data := struct {
		Title   string
		Content template.HTML
		List    template.HTML
	}{
		Title:   strings.TrimPrefix(outputPath, "/output/"),
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
			// Create the output directory if it doesn't exist
			metadata, err := extractYAMLFromMD(inputFilePath)
			if err != nil {
				return fmt.Errorf("failed to extract metadata from file %s: %w", inputFilePath, err)
			}

			// Read Markdown file
			mdContent, err := afero.ReadFile(fs, inputFilePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", inputFilePath, err)
			}

			// Remove YAML metadata from content
			mdContent = removeYAMLMetaData(mdContent)

			// Convert Markdown to HTML
			htmlContent := markdown.ToHTML(mdContent, nil, nil)

			// Prepare data for the template
			data := struct {
				Title   string
				Content template.HTML
				List    string
			}{
				Title:   metadata.Title,
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
