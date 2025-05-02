package handler

import (
	"bytes"
	"fmt"
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
		if !file.IsDir() { // Skip directories
			if strings.Contains(outputPath, "blog") {
				pathToFile := filepath.Join("content", "blog", strings.TrimSuffix(file.Name(), ".html")+".md")
				metadata, err := extractYAMLFromMD(pathToFile)
				if err != nil {
					panic(err)
				}
				if metadata.Draft {
					continue // Skip draft files
				}

				buffer.WriteString(fmt.Sprintf(`
	<div class="content-item">
		<a class="content-link" href="%s">%s</a>
		<p class="description">%s</p>
		<p class="date">%s</p>
	</div>
`, filepath.Join(file.Name()), metadata.Title, metadata.Description, metadata.Date))
			}

			if strings.Contains(outputPath, "projects") {
				pathToFile := filepath.Join("content", "projects", strings.TrimSuffix(file.Name(), ".html")+".md")
				metadata, err := extractYAMLFromMD(pathToFile)
				if err != nil {
					panic(err)
				}
				buffer.WriteString(fmt.Sprintf(`
	<img src="%s" class="project-image" alt="project">
	<div class="content-item">
		<a  href="%s">%s</a>
		<p>%s</p>
		<p>%s</p>
	</div>
`, metadata.Image, filepath.Join(file.Name()), metadata.Title, metadata.Description, metadata.Date))
			}
		}
	}
	return buffer.String() // Return the complete HTML string
}
