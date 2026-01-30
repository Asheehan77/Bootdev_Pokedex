package main

import(
	"fmt"
	"bufio"
	"os"
	s "strings"
	"github.com/Asheehan77/Bootdev_Pokedex/internal"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(*config) error
}

type config struct{
	pokeapiClient 	internal.Client
	nextLocUrl		*string
	prevLocUrl		*string
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
	"map": {
        name:        "map",
        description: "Displays the next names of 20 location areas in the Pokemon world",
        callback:    commandMap,
    },
	"mapb": {
        name:        "mapb",
        description: "Displays the previous names of 20 location areas in the Pokemon world",
        callback:    commandMapb,
    },
	}
}

func runRepl(cfg *config){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)[0]
		call, real := getCommands()[cleaned]
		if real {
			err := call.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
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