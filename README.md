# goExtractor

`goExtractor` is a Go-based command-line tool that allows you to search for files with specific extensions in a target directory (and its subdirectories) while excluding files and directories specified in a `.gitignore` file. It compiles the contents of these files into a single output file.

## Features

- Specify file extensions to include.
- Define a target directory to search.
- Exclude files and directories based on `.gitignore` patterns.
- Combine contents of matching files into a single output file.
- Exclude hidden files and directories by default.

## Requirements

- Go 1.22.4 or later.

## Installation

Clone the repository and build the executable:

```sh
git clone https://github.com/yourusername/goExtractor.git
cd goExtractor
make build
