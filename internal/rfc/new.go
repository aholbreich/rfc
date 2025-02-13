package rfc

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"text/template"
	"time"
)

func New(title string) {
	// Ensure RFC directory exists
	rfcDir := "rfc"
	if err := os.MkdirAll(rfcDir, os.ModePerm); err != nil {
		fmt.Println("Error creating RFC directory:", err)
		return
	}

	// Get the next counter value
	counter, err := getNextCounter(rfcDir)
	if err != nil {
		fmt.Println("Error determining next counter:", err)
		return
	}

	// Generate RFC metadata
	metadata := RFCMetadata{
		Title:     title,
		Author:    "Your Name", // Replace with actual author or pass as argument
		Status:    "Draft",
		CreatedAt: time.Now().Format("2006-01-02"),
	}

	templatePath := filepath.Join("templates", "basic_rfc.md")
	// Load the template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic("can't find template:" + templatePath)
	}

	// Render the template
	content := new(bytes.Buffer)
	err = tmpl.Execute(content, TemplateData{
		Title:     metadata.Title,
		Author:    metadata.Author,
		Status:    metadata.Status,
		CreatedAt: metadata.CreatedAt,
	})

	// Generate RFC ID and filename
	filename := fmt.Sprintf("%03d-%s.md", counter, title)

	// Write to file
	filePath := filepath.Join(rfcDir, filename)
	if err := os.WriteFile(filePath, []byte(content.Bytes()), 0644); err != nil {
		fmt.Println("Error writing RFC:", err)
		return
	}

	fmt.Printf("RFC '%s' created successfully at %s\n", title, filePath)
}

// getNextCounter determines the next counter value by scanning existing RFC files
func getNextCounter(rfcDir string) (int, error) {
	files, err := os.ReadDir(rfcDir)
	if err != nil {
		return 0, err
	}

	// Regex to match filenames like "001-Some title.md"
	re := regexp.MustCompile(`^(\d{3})-.+\.md$`)

	maxCounter := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		matches := re.FindStringSubmatch(file.Name())
		if len(matches) > 1 {
			counter, err := strconv.Atoi(matches[1])
			if err != nil {
				return 0, err
			}
			if counter > maxCounter {
				maxCounter = counter
			}
		}
	}

	// Return the next counter value
	return maxCounter + 1, nil
}
