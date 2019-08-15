package main

// Taken from amazing sadbox/mediawiki (with some changes ofc)
// Response is a struct used for unmarshaling the MediaWiki JSON response.
type Response struct {
	Query struct {
		// The JSON response for this part of the struct is dumb.
		// It will return something like { '23': { 'pageid': 23 ...
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
// This is a modified struct to work with the return type from Wiki API
// Fields left for readability
type Page struct {
	//	Pageid  int
	//	Ns      int
	//  Title   string
	Extract string
}
