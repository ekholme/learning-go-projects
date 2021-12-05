package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ekholme/learning-go-projects/gophercises-adventure/cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story) //%+v will print out a struct
}

//RESUME AT BUILDING HTTP HANDLER
