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
git clone https://github.com/shahinrahimi/goExtractor.git
cd goExtractor
make build
```
This will create an executable file named goextractor in the ./bin directory.

## Usage
Run the executable with the following options:
```sh
./bin/goextractor -ext=<extensions> -target=<directory> -output=<output_file>
```
## Options
- -ext: Comma-separated list of file extensions to include (e.g., go,py,tsx). If not provided, all files except .txt files are included.
- -target: Target directory to search for files. Default is the current working directory.
- -output: Output file to write the contents. Default is output.txt in the current working directory.

## Examples
Extract all .go and .py files
```sh
./bin/goextractor -ext=go,py
```
Specify a target directory
```sh
./bin/goextractor -ext=go,py -target=/path/to/directory
```
Specify an output file
```sh
./bin/goextractor -ext=go,py -target=/path/to/directory -output=results.txt
```
## Building the Project
To build the project, simply run:
```sh
make build
```
This will compile the goextractor executable and place it in the ./bin directory.
## Running Project
To build and run the project in one step, use:
```sh
make run
```
## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.
1. Fork the repository.
2. Create a new branch (git checkout -b feature-branch).
3. Commit your changes (git commit -am 'Add new feature').
4. Push to the branch (git push origin feature-branch).
5. Create a new Pull Request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.



