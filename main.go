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

	functions.ReadWebsite(websites)

}
