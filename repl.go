package main

import(
	"fmt"
	"bufio"
	"os"
	s "strings"
	"github.com/Asheehan77/Bootdev_Pokedex/internal/pokeapi"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(*config,[]string) error
}

type config struct{
	pokeapiClient 	pokeapi.Client
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
	"explore": {
        name:        "explore",
        description: "Displays the pokemon in the given area",
        callback:    commandExplore,
    },
	"catch": {
        name:        "catch",
        description: "Attempts to catch and add the given pokemon to your collection",
        callback:    commandCatch,
    },
	"inspect": {
        name:        "inspect",
        description: "Displays the stats of a caught pokemon",
        callback:    commandInspect,
    },
	"pokedex": {
        name:        "pokedex",
        description: "Displays your caught pokemon",
        callback:    commandPokedex,
    },
	}
}

func runRepl(cfg *config){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		call, real := getCommands()[cleaned[0]]
		if real {
			err := call.callback(cfg,cleaned)
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