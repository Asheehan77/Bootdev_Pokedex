package main

import(
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n")
	for _,command := range getCommands(){
		fmt.Printf("%s: %s\n",command.name,command.description)
	}
	fmt.Println()
	return nil
}