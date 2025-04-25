package handler

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

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
		filepath.Join(baseDir, "output", "static", "css"),
		filepath.Join(baseDir, "output", "static", "js"),
		filepath.Join(baseDir, "output", "static", "images"),
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
		filepath.Join(baseDir, "content", "index.md"):                homeIndex,
		filepath.Join(baseDir, "content", "blog", "index.md"):        blogIndex,
		filepath.Join(baseDir, "content", "blog", "post1.md"):        basePost,
		filepath.Join(baseDir, "content", "projects", "index.md"):    projectIndex,
		filepath.Join(baseDir, "content", "projects", "project1.md"): baseProject,
		filepath.Join(baseDir, "templates", "base.html"):             baseTemplate,
	}

	for path, content := range contentFiles {
		err := afero.WriteFile(fs, path, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Failed to create file %s: %v\n", path, err)
			continue
		}
		fmt.Printf("Created file: %s\n", path)
	}

	// Copy static files
	err := copyFiles("./assets/images", filepath.Join(baseDir, "output", "static", "images"))
	if err != nil {
		fmt.Printf("Failed to copy static files: %v\n", err)
		return err
	}

	serveFilePath := filepath.Join(baseDir, "serve.go")
	serveFileContent := serveFile

	err = afero.WriteFile(fs, serveFilePath, []byte(serveFileContent), 0644)
	if err != nil {
		fmt.Printf("Failed to create serve.go: %v\n", err)
		return err
	}
	fmt.Printf("Created file: %s\n", serveFilePath)
	return nil
}

func copyFiles(srcDir, destDir string) error {
	// Walk through the source directory
	err := filepath.WalkDir(srcDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking directory %s: %w", path, err)
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Construct destination file path
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return fmt.Errorf("failed to compute relative path: %w", err)
		}
		destPath := filepath.Join(destDir, relPath)

		// Ensure the destination directory for the file exists
		err = os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory for file %s: %w", destPath, err)
		}

		// Copy the file
		err = copyFile(path, destPath)
		if err != nil {
			return fmt.Errorf("failed to copy file %s to %s: %w", path, destPath, err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// Helper function to copy a single file
func copyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	// Copy the file contents
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy contents: %w", err)
	}

	// Ensure the copied file has the same permissions
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to get source file info: %w", err)
	}
	err = os.Chmod(dest, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to set file permissions: %w", err)
	}

	return nil
}
