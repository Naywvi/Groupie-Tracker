package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var templatesDir = os.Getenv("TEMPLATES_DIR")
var artist Artist //creation instance struct artist

//Search
func search(artist Artist, apparitionP []string, albumP []string, membersP []string, locationsP []string) bool { //handles the comparisons with the users criteria to the artists
	if apparitionP != nil {
		if strconv.Itoa(artist.CreationDate) == apparitionP[0] && apparitionP[0] != "1945" {
			return true
		}
	}
	if albumP != nil {
		if artist.FirstAlbum == albumP[0] && albumP[0] != "1945" {
			return true
		}
	}
	if membersP != nil {
		if strconv.Itoa(len(artist.Members)) == membersP[0] && membersP[0] != "" {
			return true
		}
	}
	if locationsP != nil {
		if strings.Contains(artist.Location.Locations[0], locationsP[0]) {
			return true
		}
	}
	return false
}

//Marshal API > Strcut
func request(url string) []byte { //simple get request to get the api data
	req, _ := http.Get(url)
	body, _ := ioutil.ReadAll(req.Body)
	return body
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)
	http.HandleFunc("/pages/tracker", tracker)
	http.HandleFunc("/pages/artist", tracker)
	fmt.Printf("Started server successfully on http://localhost:8089/\n")
	http.ListenAndServe(":8089", nil)
}
