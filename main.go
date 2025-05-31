package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var filepath string
	fmt.Println("Enter file path")
	fmt.Scan(&filepath)

	content, err := os.ReadFile(filepath)

	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	text := string(content)

	websites := strings.Split(text, ",")

	fmt.Println(websites)

}
