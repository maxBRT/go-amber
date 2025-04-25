package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type MetaData struct {
	Title       string `yaml:"title"`
	Date        string `yaml:"date"`
	Author      string `yaml:"author"`
	Description string `yaml:"description"`
	Image       string `yaml:"image"`
	Draft       bool   `yaml:"draft"`
}

func extractYAMLFromMD(filePath string) (*MetaData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var yamlBuffer bytes.Buffer
	scanner := bufio.NewScanner(file)
	isInYAML := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "---" {
			if isInYAML {
				break
			}
			isInYAML = true
			continue
		}
		if isInYAML {
			yamlBuffer.WriteString(line + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if yamlBuffer.Len() == 0 {
		return nil, fmt.Errorf("no YAML front matter found")
	}

	var metaData MetaData
	if err := yaml.Unmarshal(yamlBuffer.Bytes(), &metaData); err != nil {
		return nil, err
	}

	return &metaData, nil
}

func removeYAMLMetaData(content []byte) []byte {
	re := regexp.MustCompile(`(?s)^---\n(.*?)\n---\n`)
	return re.ReplaceAll(content, []byte{})
}
