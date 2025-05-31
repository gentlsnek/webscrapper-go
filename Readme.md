# webscrapper-go

A simple command-line tool in Go to read a list of websites from a file and (optionally) process them for HTTP connectivity.

## Features

- **Reads website URLs from a file** (comma-separated)
- **Easy extensibility**: Functions are modular for further web processing, such as scraping or status checking
- **Uses Go's concurrency** for potential parallel processing of URLs (see `ReadWebsite.go`)

---

## Getting Started

### Prerequisites

- Go 1.24.3 or later

### Installation

Clone the repository and build:

```bash
git clone https://github.com/gentlsnek/webscrapper-go.git
cd webscrapper-go
go build
```

---

## Usage

By default, the tool reads a file containing website addresses (comma-separated) and prints the list.

1. Prepare a file (e.g., `websites.txt`) with content like:
    ```
    example.com,google.com,github.com
    ```
2. Run the program:

    ```bash
    go run main.go
    ```

3. When prompted, enter the file path (e.g.):
    ```
    Enter file path
    websites.txt
    ```

4. The program prints the list of websites as a Go string slice.

---

## Project Structure

- **main.go**: Entry point. Reads file path from user, loads sites using `functions.Readtext()`.
- **functions/Readfile.go**: Defines `Readtext`, which reads a file and splits its content by commas to produce a string slice.
- **functions/ReadWebsite.go**: Contains a function stub for checking website connectivity (can be extended to do more advanced scraping or data extraction).

---

## Code Overview

### main.go

```go
package main

import (
    "fmt"
    "webscrapper-go/functions"
)

func main() {
    var filepath string
    fmt.Println("Enter file path")
    fmt.Scan(&filepath)
    websites := functions.Readtext(filepath)
    fmt.Println(websites)
}
```

### functions/Readfile.go

```go
func Readtext(s string) []string {
    content, err := os.ReadFile(s)
    if err != nil {
        fmt.Println("error reading file", err.Error())
        return []string{}
    }
    return strings.Split(string(content), ",")
}
```

### functions/ReadWebsite.go

```go
func ReadWebsite(sites []string) {
    c := make(chan string)
    for _, list := range sites {
        go checkLink("http://"+list, c)
    }
}
func checkLink(link string, c chan string) {
    status, err := http.Get(link)
    if err != nil {
        fmt.Println("could not establish connection to", err)
        c <- link
        return
    }
    fmt.Println(status)
    c <- link
}
```

---

## Extending

- Implement logic in `functions/ReadWebsite.go` to scrape content, check statuses, or gather metrics from each site.
- Update `main.go` to invoke `ReadWebsite()` after reading the sites, if desired.

---

## Dependencies

- [golang.org/x/net](https://pkg.go.dev/golang.org/x/net) (indirect)
- Go standard library

---

## License

MIT

---

> **Note:** This README was generated based on the discovered files. For more files or updates, check the [GitHub webscrapper-go repository](https://github.com/gentlsnek/webscrapper-go).
