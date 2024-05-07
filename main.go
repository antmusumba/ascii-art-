package main

import (
	"ascii/utils"
	"fmt"
	"os"
	"strings"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Println("Provide enough arguments") 
		return
	}
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading the file")
		return
	}
	contentlines := utils.SplitFile(string(content))
	if len(contentlines) != 856 {
		fmt.Println("invalid file input")
		return
	}
	utils.DisplayText(strings.Join(os.Args[1:], " "),contentlines)
}

