package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(`    __.-._`)
	fmt.Println(`    '-._"7'`)
	fmt.Println(`     /'.-c`)
	fmt.Println(`     |  /T`)
	fmt.Println(`    _)_/LI`)
	fmt.Print("\nWelcome to the Ultimate CLI Tool, may the Force be with You\n\n")

	// Define the 'id' string flag
	id := flag.String("id", "1", "ID")

	flag.Parse() // <-- Must be called after all the flags are defined

	nonFlagArguments := flag.Args() // <-- Get the non-flag arguments
	if len(nonFlagArguments) > 0 {
		switch nonFlagArguments[0] {
		case "swapi":
			fmt.Printf("SWAPI mode enabled, fetching person with id '%s'\n\n", *id)
			handleHttpGet("https://swapi.co/api/people/" + *id + "/")
		case "todo":
			fmt.Printf("To-do mode enabled, fetching to-do with id '%s'\n\n", *id)
			handleHttpGet("https://jsonplaceholder.typicode.com/todos/" + *id)
		}
	} else {
		fmt.Print("No mode provided, please read the README\n\n")
	}
}

func handleHttpGet(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	var pretty bytes.Buffer
	if err = json.Indent(&pretty, body, "", "    "); err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Print(string(pretty.Bytes()))
}
