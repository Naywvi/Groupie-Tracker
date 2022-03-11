package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func tracker(w http.ResponseWriter, r *http.Request) { //function that starts when tracker or artist page is loaded
	var (
		//Param Query
		param, Randparam, apparitionP, albumP, membersP, locationsP = r.URL.Query()["artist"], r.URL.Query()["RandomArtist"], r.URL.Query()["apparition"], r.URL.Query()["album"], r.URL.Query()["members"], r.URL.Query()["location"]
		art                                                         Templ
		url                                                         string //url for making requests
		selectart                                                   bool
	)
	rand.Seed(time.Now().UnixNano()) //random number to pick random artist
	art.Random = 1 + rand.Intn(51-0)
	var research bool
	var iteration int
	var generated []string //stock random integer for randomizer tracker page
	if apparitionP != nil || albumP != nil || membersP != nil || locationsP != nil {
		research = true
		iteration = 53
	} else {
		iteration = 11
	}

	for i := 1; i < iteration; i++ {
		url = "https://groupietrackers.herokuapp.com/api/artists/"
		artist.Members = nil
		//-------------------------- FETCH RESSOURCES API JSON
		if param != nil { //user wants specific artist page
			url += param[0]
			selectart = true
		} else if Randparam != nil { //user wants random artist page
			url += Randparam[0]
			selectart = true
		} else { //user wants page tracker with every artist on it
			rand.Seed(time.Now().UnixNano()) //random number to pick random artist
			random := strconv.Itoa(1 + rand.Intn(51-0))
			for stringInSlice(random, generated) {
				random = strconv.Itoa(1 + rand.Intn(51-0))
			}
			generated = append(generated, random)
			url += random
			selectart = false
		}
		s, _, _ := json.Unmarshal([]byte(request(url)), &artist),
			json.Unmarshal([]byte(request(artist.Locations)), &artist.Location),
			json.Unmarshal([]byte(request(artist.Location.Dates)), &artist.Location.DatesLoc)
		str := ""
		for _, v := range artist.Members { //prints members without []
			str += v + " "
		}
		artist.Membersstr = str
		if s != nil {
			fmt.Print("error when encoding the struct")
		}
		//-------------------------- append a la struct envoyée a l'api l'instance en cours
		if research { //if user has chosen some criteria
			if search(artist, apparitionP, albumP, membersP, locationsP) { //check whether artist in api is eligible to show
				art.Artiste = append(art.Artiste, artist) //put it on the page
			}
		} else { //no criteria, just display everything
			art.Artiste = append(art.Artiste, artist)
		}

		if selectart { //if artist is selected print the page of the artist and stop
			JscriptStr()
			(template.Must(template.ParseFiles(filepath.Join(templatesDir, "../templates/artist.html")))).Execute(w, art)
			return
		}
		url = ""
	}
	(template.Must(template.ParseFiles(filepath.Join(templatesDir, "../templates/tracker.html")))).Execute(w, art)
}
