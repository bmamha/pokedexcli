package main

import (
"bufio"
"fmt"
"os"
"strings"
"log"
)




func startRepl(cfg *config) {

		scanner := bufio.NewScanner(os.Stdin)
		
  	for {
        fmt.Print("Pokedex > ")
        scanner.Scan()
	      err := scanner.Err()
	     
				if err != nil {
		     log.Fatal(err)
	       }
	text := scanner.Text()
	words := CleanInput(text) 
  
	command, ok := getCommands()[words[0]]
	
	if !ok {
		fmt.Println("Unknown command")
		continue 
	} else {
		err := command.callback(cfg) 
		if err != nil {
			fmt.Println(err)
		}
		continue

	}
  
}
 
}

type cliCommand struct {
	name string
	description string
	callback func(c *config) error 
}

func CleanInput(text string) []string {
	lowered_text := strings.ToLower(text)
	words := strings.Fields(lowered_text)

	return words  
}


func getCommands() map[string]cliCommand {
	return  map[string]cliCommand{
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
	"map": {
		name: "map",
		description: "Displays location area of pokemon",
		callback: commandMap,
	},
 "mapb": {
	 name: "mapb",
	 description: "Displays location of previous page",
	 callback: commandMapb,
 },

}
}
 
