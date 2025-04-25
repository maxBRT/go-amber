package handler

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
)

func generateFileListHTML(files []os.FileInfo, outputPath string) string {
	var buffer bytes.Buffer
	for _, file := range files {
		if file.Name() == "index.html" {
			continue // Skip the index.html file
		}
		buffer.WriteString("<div class=content-list>")
		if !file.IsDir() { // Skip directories
			if strings.Contains(outputPath, "blog") {
				pathToFile := filepath.Join("content", "blog", strings.TrimSuffix(file.Name(), ".html")+".md")
				metadata, err := extractYAMLFromMD(pathToFile)
				if err != nil {
					panic(err)
				}
				// Start the list item
				buffer.WriteString(`<div class="list-item-blog">`)
				buffer.WriteString(`<li><a href="`)
				// Add the file path to the href attribute
				buffer.WriteString(filepath.Join(file.Name()))
				buffer.WriteString(`">`)
				// Add the file name as the link text
				buffer.WriteString(metadata.Title)
				buffer.WriteString(`</a>`)
				buffer.WriteString("<p class=description>")
				buffer.WriteString(metadata.Description)
				buffer.WriteString("</p>")
				buffer.WriteString("<p class=date>")
				buffer.WriteString(metadata.Date)
				buffer.WriteString("</p>")
				buffer.WriteString("</li>")
				buffer.WriteString("</div>")
			}

			if strings.Contains(outputPath, "projects") {
				pathToFile := filepath.Join("content", "projects", strings.TrimSuffix(file.Name(), ".html")+".md")
				metadata, err := extractYAMLFromMD(pathToFile)
				if err != nil {
					panic(err)
				}

				// Start the list item
				buffer.WriteString(`<div class="list-item-project">`)
				buffer.WriteString("<img src=" + metadata.Image + " alt=Project class=project-image>")
				buffer.WriteString(`<li><a href="`)
				// Add the file path to the href attribute
				buffer.WriteString(filepath.Join(file.Name()))
				buffer.WriteString(`">`)
				// Add the file name as the link text
				buffer.WriteString(metadata.Title)
				buffer.WriteString(`</a>`)
				buffer.WriteString("<p class=description>")
				buffer.WriteString(metadata.Description)
				buffer.WriteString("</p>")
				buffer.WriteString("<p class=date>")
				buffer.WriteString(metadata.Date)
				buffer.WriteString("</p>")
				buffer.WriteString("</li>")
				buffer.WriteString("</div>")
			}
		}
	}
	buffer.WriteString("</div>")
	return buffer.String() // Return the complete HTML string
}
