package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Lang        string `yaml:"lang"`
	Author      string `yaml:"author"`
	SiteUrl     string `yaml:"siteUrl"`
	Bio         string `yaml:"bio"`
	AvatarUrl   string `yaml:"avatarUrl"`
	Email       string `yaml:"email"`
	Links       []struct {
		Name string `yaml:"name"`
		URL  string `yaml:"url"`
	}
	Socials []struct {
		Name string `yaml:"name"`
		URL  string `yaml:"url"`
	} `yaml:"socials"`
}

func main() {
	// Load configuration
	config, err := loadConfig("config.yaml")
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
	err = copyAndMinifyAssets()
	if err != nil {
		log.Fatalf("Error copying and minifying assets: %v", err)
	}

	fmt.Println("Site generated successfully!")
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
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

func generateHTML(config *Config, content string) error {
	// Load HTML template
	tmpl, err := template.ParseFiles("src/templates/index.html")
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
		Config  *Config
		Content template.HTML
	}{
		Config:  config,
		Content: template.HTML(content),
	}

	// Execute template with data
	return tmpl.Execute(outputFile, data)
}

func copyAndMinifyAssets() error {
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

	// Minify and combine CSS files
	cssFiles, err := filepath.Glob("src/templates/assets/css/*.css")
	if err != nil {
		return fmt.Errorf("failed to list CSS files: %w", err)
	}
	if err := combineAndMinifyFiles(cssFiles, "assets/css/styles.css", "text/css", m); err != nil {
		return fmt.Errorf("failed to minify CSS files: %w", err)
	}

	// Minify and combine JS files
	jsFiles, err := filepath.Glob("src/templates/assets/js/*.js")
	if err != nil {
		return fmt.Errorf("failed to list JS files: %w", err)
	}
	if err := combineAndMinifyFiles(jsFiles, "assets/js/script.js", "application/javascript", m); err != nil {
		return fmt.Errorf("failed to minify JS files: %w", err)
	}

	// Copy images
	imgFiles, err := filepath.Glob("src/templates/assets/img/*")
	if err != nil {
		return fmt.Errorf("failed to list image files: %w", err)
	}
	if err := copyFiles(imgFiles, "assets/img"); err != nil {
		return fmt.Errorf("failed to copy image files: %w", err)
	}

	// Copy icons (e.g., favicon)
	iconFiles, err := filepath.Glob("src/templates/assets/favicon.ico")
	if err != nil {
		return fmt.Errorf("failed to list icon files: %w", err)
	}
	if err := copyFiles(iconFiles, "assets"); err != nil {
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
