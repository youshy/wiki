package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const wikiAddress = "https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro&explaintext&redirects=1&titles="

const address = "https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro&explaintext&redirects=1&titles=Stack%20Overflow"

type Data struct {
	Batchcomplete string
	Query         struct {
		Pages struct {
			Number struct {
				Pageid  int
				Ns      int
				Title   string
				Extract string
			}
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`How can I find anything about someone who doesn't exist?`)
		os.Exit(1)
	}

	arguments := os.Args
	person := arguments[1]
	fmt.Println(person)

	// url := wikiAddress + person
	url := address
	fmt.Println(url)

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Println(string(body))

	var output Data
	json.Unmarshal(body, &output)

	fmt.Println(output.Query)
	fmt.Println("%+v \n", output)
	fmt.Println("%+v \n", output.Query.Pages.Number.Extract)
}
