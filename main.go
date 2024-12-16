package main

import (
	"fmt"
  "strings"
)

func main(){
	fmt.Println("Hello, World!")
}


func CleanInput(text string) []string {
	lowered_text := strings.ToLower(text)
	words := strings.Fields(lowered_text)

	return words  
}

