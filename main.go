package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	extensions := flag.String("ext", "", "Comma-separated list of file extensions to include (e.g., go,py,tsx)")
	targetDir := flag.String("target", "", "Target directory to search for files (default is current working directory)")
	outputFile := flag.String("output", "output.txt", "Output file to write the contents (default is output.txt in the current working directory)")

	flag.Parse()

	// Get the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	// Set default target directory to the current working directory
	if *targetDir == "" {
		*targetDir = workingDir
	}

	// Set default output file path to the current working directory if not specified
	if !filepath.IsAbs(*outputFile) {
		*outputFile = filepath.Join(workingDir, *outputFile)
	}

	exts := []string{}
	if *extensions != "" {
		exts = strings.Split(*extensions, ",")
	}
	extMap := make(map[string]bool)
	for _, ext := range exts {
		extMap["."+ext] = true
	}

	gitignorePatterns := readGitignore(*targetDir)

	var files []string
	err = filepath.Walk(*targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(*targetDir, path)
		if err != nil {
			return err
		}

		// Exclude hidden files and directories
		if strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() && info.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}

		// Exclude files based on .gitignore patterns
		if matchesGitignore(relativePath, gitignorePatterns) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if !info.IsDir() {
			// If extensions are provided, filter by extensions
			if len(extMap) > 0 {
				if extMap[filepath.Ext(path)] {
					files = append(files, path)
				}
			} else {
				// If no extensions are provided, exclude .txt files
				if filepath.Ext(path) != ".txt" {
					files = append(files, path)
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", *targetDir, err)
		os.Exit(1)
	}

	var output []byte
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %v: %v\n", file, err)
			continue
		}
		relativePath, err := filepath.Rel(*targetDir, file)
		if err != nil {
			fmt.Printf("Error getting relative path for file %v: %v\n", file, err)
			continue
		}
		output = append(output, []byte(fmt.Sprintf("%s\n%s\n\n", relativePath, string(data)))...)
	}

	err = os.WriteFile(*outputFile, output, 0644)
	if err != nil {
		fmt.Printf("Error writing to output file %v: %v\n", *outputFile, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully wrote to %v\n", *outputFile)
}

// readGitignore reads the .gitignore file in the target directory and returns a list of patterns
func readGitignore(targetDir string) []string {
	var patterns []string
	gitignorePath := filepath.Join(targetDir, ".gitignore")
	file, err := os.Open(gitignorePath)
	if err != nil {
		return patterns
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Ignore comments and empty lines
		if strings.TrimSpace(line) != "" && !strings.HasPrefix(line, "#") {
			patterns = append(patterns, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading .gitignore file: %v\n", err)
	}

	return patterns
}

// matchesGitignore checks if a file path matches any of the .gitignore patterns
func matchesGitignore(path string, patterns []string) bool {
	for _, pattern := range patterns {
		matched, _ := filepath.Match(pattern, path)
		if matched {
			return true
		}
		// Check for directory patterns
		if strings.HasSuffix(pattern, "/") && strings.HasPrefix(path, strings.TrimSuffix(pattern, "/")) {
			return true
		}
	}
	return false
}
