package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"lucasnevespereira/src/configs"
	"os"
	"path/filepath"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
	"github.com/yuin/goldmark"
)

func main() {
	// Load configuration
	config, err := configs.LoadSiteConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Load and parse content
	content, err := os.ReadFile("src/content/index.md")
	if err != nil {
		log.Fatalf("Error reading content: %v", err)
	}

	htmlContent, err := markdownToHTML(content)
	if err != nil {
		log.Fatalf("Error converting Markdown to HTML: %v", err)
	}

	// Generate HTML
	err = generateHTML(config, htmlContent)
	if err != nil {
		log.Fatalf("Error generating HTML: %v", err)
	}

	// Copy and minify CSS, JS, and images
	err = copyAndMinifyAssets(config.Theme)
	if err != nil {
		log.Fatalf("Error copying and minifying assets: %v", err)
	}

	fmt.Println("Site generated successfully!")
}

func markdownToHTML(content []byte) (string, error) {
	var buf bytes.Buffer
	md := goldmark.New()
	err := md.Convert(content, &buf)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func generateHTML(config *configs.SiteConfig, content string) error {
	templateFile := fmt.Sprintf("src/templates/%s/index.html", config.Theme)

	// Load HTML template
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	// Open output file
	outputFile, err := os.Create("index.html")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Define data to pass to the template
	data := struct {
		Config  *configs.SiteConfig
		Content template.HTML
	}{
		Config:  config,
		Content: template.HTML(content),
	}

	// Execute template with data
	return tmpl.Execute(outputFile, data)
}

func copyAndMinifyAssets(templateFolder string) error {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("application/javascript", js.Minify)

	// Create assets directories if they don't exist
	if err := os.MkdirAll("assets/css", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create assets/css directory: %w", err)
	}
	if err := os.MkdirAll("assets/js", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create assets/js directory: %w", err)
	}
	if err := os.MkdirAll("assets/img", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create assets/img directory: %w", err)
	}
	if err := os.MkdirAll("assets/icons", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create assets/icons directory: %w", err)
	}

	// Minify and combine CSS files
	cssFiles, err := filepath.Glob(fmt.Sprintf("src/templates/%s/assets/css/*.css", templateFolder))
	if err != nil {
		return fmt.Errorf("failed to list CSS files: %w", err)
	}
	if err := combineAndMinifyFiles(cssFiles, "assets/css/styles.css", "text/css", m); err != nil {
		return fmt.Errorf("failed to minify CSS files: %w", err)
	}

	// Minify and combine JS files
	jsFiles, err := filepath.Glob(fmt.Sprintf("src/templates/%s/assets/js/*.js", templateFolder))
	if err != nil {
		return fmt.Errorf("failed to list JS files: %w", err)
	}
	if err := combineAndMinifyFiles(jsFiles, "assets/js/script.js", "application/javascript", m); err != nil {
		return fmt.Errorf("failed to minify JS files: %w", err)
	}

	// Copy images
	imgFiles, err := filepath.Glob(fmt.Sprintf("src/templates/%s/assets/img/*", templateFolder))
	if err != nil {
		return fmt.Errorf("failed to list image files: %w", err)
	}
	if err := copyFiles(imgFiles, "assets/img"); err != nil {
		return fmt.Errorf("failed to copy image files: %w", err)
	}

	// Copy favicons
	iconsFiles, err := filepath.Glob(fmt.Sprintf("src/templates/%s/assets/icons/*", templateFolder))
	if err != nil {
		return fmt.Errorf("failed to list icon files: %w", err)
	}
	if err := copyFiles(iconsFiles, "assets/icons"); err != nil {
		return fmt.Errorf("failed to copy icon files: %w", err)
	}

	return nil
}

func combineAndMinifyFiles(files []string, outputFile, mimeType string, m *minify.M) error {
	var combinedData bytes.Buffer
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", file, err)
		}
		combinedData.Write(data)
	}

	// Create the output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", outputFile, err)
	}
	defer outFile.Close()

	// Minify and write the combined data to the output file
	return m.Minify(mimeType, outFile, &combinedData)
}

func copyFiles(files []string, outputDir string) error {
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", file, err)
		}

		outputPath := filepath.Join(outputDir, filepath.Base(file))
		outFile, err := os.Create(outputPath)
		if err != nil {
			return fmt.Errorf("failed to create %s: %w", outputPath, err)
		}
		defer outFile.Close()

		if _, err := io.Copy(outFile, bytes.NewReader(data)); err != nil {
			return fmt.Errorf("failed to copy data to %s: %w", outputPath, err)
		}
	}

	return nil
}
