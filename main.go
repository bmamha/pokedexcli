package main

import (
	"fmt"
  "strings"
	"bufio"
	"os"
	"log"
)


type cliCommand struct {
	name string
	description string
	callback func() error 
}


func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
  fmt.Print("Pokedex > ")
  scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	text := scanner.Text()
	words := strings.Fields(strings.ToLower(text))
  command_map := map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
}


	command, ok := command_map[words[0]]
	if !ok {
		fmt.Println("Unknown command")
	} else {
		command.callback()
	}
  
}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil 

}

func commandHelp() error {
	command_map := map[string]cliCommand{
	 "help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
		"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
}

   fmt.Println("Welcome to the Pokedex!")
   fmt.Println("Usage")
	 for key, value := range command_map {
		fmt.Printf("%v: %v\n", key, value.description)
	}

	return nil 

}


func CleanInput(text string) []string {
	lowered_text := strings.ToLower(text)
	words := strings.Fields(lowered_text)

	return words  
}



