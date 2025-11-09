package main
import(
	s "strings"
)

func cleanInput(text string) []string{
	
	cleaned := s.Fields(s.ToLower(text))
	return cleaned
}