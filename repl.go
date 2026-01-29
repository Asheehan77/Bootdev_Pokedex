package main

import(
	"fmt"
	"bufio"
	"os"
	s "strings"
)

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

func getCommands() map[string]cliCommand{
    return map[string]cliCommand{
	"exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
	"help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    },
	}
}

func runRepl(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)[0]
		call, real := getCommands()[cleaned]
		if real {
			call.callback()
		}else{
			fmt.Println("Unknown command")
		}
		
	}
}
func cleanInput(text string) []string {
	lower := s.ToLower(text)
	words := s.Fields(lower)
	return words
}