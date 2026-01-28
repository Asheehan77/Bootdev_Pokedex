package main

import(
	"fmt"
	"bufio"
	"os"
	s "strings"
)

func runRepl(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		fmt.Println("Your command was:",cleaned[0])
	}
}
func cleanInput(text string) []string {
	lower := s.ToLower(text)
	words := s.Fields(lower)
	return words
}