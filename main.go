package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

const wikiAddress = "https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro&explaintext&redirects=1&titles="

const address = "https://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro&explaintext&redirects=1&titles=Stack%20Overflow"

// Taken from amazing sadbox/mediawiki
// Response is a struct used for unmarshaling the MediaWiki JSON response.
type Response struct {
	Query struct {
		// The JSON response for this part of the struct is dumb.
		// It will return something like { '23': { 'pageid': 23 ...
		//
		// As a workaround you can use PageSlice which will create
		// a list of pages from the map.
		Pages map[string]Page
	}
}

// PageSlice generates a slice from Pages to work around the sillyness in
// the MediaWiki API.
func (r *Response) PageSlice() []Page {
	pl := []Page{}
	for _, page := range r.Query.Pages {
		pl = append(pl, page)
	}
	return pl
}

// A Page represents a MediaWiki page and its metadata.
type Page struct {
	Pageid    int
	Ns        int
	Title     string
	Touched   string
	Lastrevid int
	// Mediawiki will return '' for zero, this makes me sad.
	// If for some reason you need this value you'll have to
	// do some type assertion sillyness.
	Counter   interface{}
	Length    int
	Edittoken string
	Revisions []struct {
		Revid         int       `json:"revid"`
		Parentid      int       `json:"parentid"`
		Minor         string    `json:"minor"`
		User          string    `json:"user"`
		Userid        int       `json:"userid"`
		Timestamp     time.Time `json:"timestamp"`
		Size          int       `json:"size"`
		Sha1          string    `json:"sha1"`
		ContentModel  string    `json:"contentmodel"`
		Comment       string    `json:"comment"`
		ParsedComment string    `json:"parsedcomment"`
		ContentFormat string    `json:"contentformat"`
		Body          string    `json:"*"` // Take note, MediaWiki literally returns { '*':
	}
	Imageinfo []struct {
		Url            string
		Descriptionurl string
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println(`How can I find anything about someone who doesn't exist?`)
		os.Exit(1)
	}

	arguments := os.Args
	person := strings.Replace(arguments[1], " ", "%20", -1)

	url := wikiAddress + person

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	bodyToString := string(body)

	re := regexp.MustCompile(`\b(\w*extract\w*)\b`)

	result := re.Split(bodyToString, -1)

	fmt.Println(result[1])

	// fmt.Println(strings.TrimLeft(bodyToString, "extract"))

	// this will return the JSON
	// needs a lot of work
	// var response Response
	// err = json.Unmarshal(body, &response)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pl := response.PageSlice()
}
