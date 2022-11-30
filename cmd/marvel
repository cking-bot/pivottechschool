package main

import (
	"fmt"
	"log"

	"github.com/cking-bot/pivottechschool/marvel"
)

func main() {

	client := marvel.NewClient(marvel.BaseURL)
	characters, err := client.GetCharacters()
	if err != nil {
		log.Println(err)
	}
	for _, char := range characters {
		fmt.Println(char.Name)
		fmt.Println(char.Description)
	}

}
