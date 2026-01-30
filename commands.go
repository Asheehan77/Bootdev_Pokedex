package main

import(
	"fmt"
	"os"
	"errors"
	"github.com/Asheehan77/Bootdev_Pokedex/internal"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n")
	for _,command := range getCommands(){
		fmt.Printf("%s: %s\n",command.name,command.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config) error {
	fmt.Println("Getting map")
	var locations internal.LocationList
	var err error
	if cfg.nextLocUrl != nil {
		locations,err = cfg.pokeapiClient.GetLocations(cfg.nextLocUrl)
	}else{
		locations,err = cfg.pokeapiClient.GetLocations(nil)
	}
	if err != nil {
		fmt.Println("err1 map")
		return err
	}

	for _,loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	cfg.nextLocUrl = locations.Next
	cfg.prevLocUrl = locations.Previous
	return nil
}

func commandMapb(cfg *config) error {
	var locations internal.LocationList
	var err error
	if cfg.prevLocUrl != nil {
		locations,err = cfg.pokeapiClient.GetLocations(cfg.prevLocUrl)
	}else{
		err := errors.New("you're on the first page")
		return err
	}
	if err != nil {
		return err
	}

	for _,loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	cfg.nextLocUrl = locations.Next
	cfg.prevLocUrl = locations.Previous
	return nil
}