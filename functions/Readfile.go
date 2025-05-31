package functions

import (
	"fmt"
	"os"
	"strings"
)

func Readtext(s string) []string {
	content, err := os.ReadFile(s)

	if err != nil {
		fmt.Println("error reading file", err.Error())
		return []string{}
	}
	return strings.Split(string(content), ",")

}
