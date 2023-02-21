[![GitHub release (latest by date)](https://img.shields.io/github/v/release/shufo/find-config)](https://github.com/shufo/find-config/releases)
[![codecov](https://codecov.io/gh/shufo/find-config/branch/main/graph/badge.svg?token=PeApmbPdsM)](https://codecov.io/gh/shufo/find-config)

# find-config

Recursive Directory Search for Config File

This is a simple Go program that recursively searches up directories to find a config file. The program starts the search from a specified directory and continues searching up the directory tree until it finds the config file or reaches the root directory.

## Installation

To install and run this program, you must have [Go](https://golang.org/) installed on your machine.
Once you have Go installed, you can install and run the program using the following commands:

```bash
$ go get github.com/shufo/find-config
```

## Usage

To start the search, use `config.Find(filename, startDir)`.
The program will return the absolute path of the config file if it is found, or an error message if the file is not found.

```golang
import "github.com/shufo/find-config"

path, err := config.Find(".eslintrc.json", ".")

if err != nil {
	log.Fatalln(err)
}

fmt.Println(path)
// => /home/user/.eslintrc.json
```

## License

This program is licensed under the MIT License. Feel free to use, modify, and distribute this code as you see fit.
