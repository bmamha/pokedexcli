package main

import (
	"fmt"
  "strings"
	"bufio"
	"os"
	"log"
)

func main(){
	for {
	fmt.Print("Pokedex > ")
	scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	text := scanner.Text()
	words := strings.Fields(strings.ToLower(text))
	fmt.Printf("Your command was: %v\n", words[0])

}
}


func CleanInput(text string) []string {
	lowered_text := strings.ToLower(text)
	words := strings.Fields(lowered_text)

	return words  
}




