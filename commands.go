package main

import(
	"fmt"
	"os"
	"errors"
	"github.com/Asheehan77/Bootdev_Pokedex/internal/pokeapi"
	"math/rand"
	"time"
)

func commandExit(cfg *config, param []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, param []string) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n")
	for _,command := range getCommands(){
		fmt.Printf("%s: %s\n",command.name,command.description)
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config, param []string) error {
	var locations pokeapi.LocationList
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

func commandMapb(cfg *config, param []string) error {
	var locations pokeapi.LocationList
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

func commandExplore(cfg *config, param []string) error {
	if len(param) < 2 {
		return errors.New("No location name")
	}
	loc,err := cfg.pokeapiClient.GetLocationInfo(&param[1])

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _,pok := range loc.PokemonEncounters {
		fmt.Println(pok.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, param []string) error {
	if len(param) < 2 {
		return errors.New("No Pokemon to catch")
	}

	pok,err := cfg.pokeapiClient.GetPokemonInfo(&param[1])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n",pok.Name)
	time.Sleep(1*time.Second)
	fmt.Printf("The ball shakes\n")
	time.Sleep(1*time.Second)
	fmt.Printf("The ball shakes\n")
	time.Sleep(1*time.Second)
	fmt.Printf("The ball shakes\n")

	catch_chance := rand.Intn(1000)
	fmt.Println("Catch Chance: ",catch_chance)
	if catch_chance >= pok.BaseExperience {
		fmt.Printf("%s was caught!\n",pok.Name)
		cfg.pokeapiClient.AddPokemon(pok)
	}else{
		fmt.Printf("%s escaped!\n",pok.Name)
	}
	return nil
}

func commandInspect(cfg *config, param []string) error {
	if len(param) < 2 {
		return errors.New("No Pokemon to inspect")
	}

	pok,err := cfg.pokeapiClient.GetPokemon(param[1])
	if err != nil {
		return err
	}
	if pok.Name == param[1] {
		fmt.Println("Height: ",pok.Height)
		fmt.Println("Weight: ",pok.Weight)
		fmt.Printf("Stats:\n")
		for _,s := range pok.Stats {
			fmt.Printf("  -%s: %v\n",s.Stat.Name,s.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _,t := range pok.Types {
			fmt.Printf("  -%s\n",t.Type.Name)
		}
		return nil
	}

	return errors.New("You haven't caught that pokemon")
}

func commandPokedex(cfg *config, param []string) error {
	poklist := cfg.pokeapiClient.GetPokemonList()
	fmt.Println("Your Pokedex:")
	for _,pok := range poklist {
		fmt.Println(" - ",pok.Name)
	}
	return nil
}