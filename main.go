package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const wikiAddress = "https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro&explaintext&redirects=1&titles="

func GetWiki(url string) []Page {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	pl := response.PageSlice()

	return pl
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`How can I find anything about someone who doesn't exist?`)
		os.Exit(1)
	}

	arguments := os.Args[1:]
	person := strings.Replace(arguments[0], " ", "%20", -1)

	url := wikiAddress + person

	response := GetWiki(url)

	fmt.Println("Searching for " + arguments[0] + "\n")
	fmt.Println(response[0].Extract)

}
