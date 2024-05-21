package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func loadAsset(w http.ResponseWriter, r *http.Request, file string) {
	data, err := ioutil.ReadFile("." + file)
	if err != nil {
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(file)
	var mimeType string

	switch {
	case ext == ".css":
		mimeType = "text/css"
	case ext == ".js":
		mimeType = "application/javascript"
	default:
		mimeType = http.DetectContentType(data)
	}

	w.Header().Set("Content-Type", mimeType+" charset=utf-8")
	w.Write(data)
}

func executeTemplate(w http.ResponseWriter, file string, data interface{}) {
	files := []string{
		"./templates/base.html",
		"./templates/search.html",
		// "./templates/nav.html",
	}
	files = append(files, file)

	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	t.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func getArtistList() []Artist {
	var artists []Artist

	resp, err := callApi(Api["artists"])
	if err != nil {

	}
	json.Unmarshal([]byte(resp), &artists)

	for i, a := range artists {
		artists[i].MembersCount = len(artists[i].Members)
		artists[i].AuthorMemb = "Author"
		if len(artists[i].Members) > 1 {
			artists[i].AuthorMemb = "Member"
		}
		artists[i].AllNames = artists[i].Name
		for j := range artists[i].Members {
			artists[i].AllNames += "|" + artists[i].Members[j]
		}
		artists[i].Concerts = getConcerts(a.Relations)
	}
	return artists
}

func getArtist(id int) Artist {
	var artist Artist

	resp, err := callApi(Api["artists"] + "/" + strconv.Itoa(id))
	if err != nil {

	}
	json.Unmarshal([]byte(resp), &artist)

	artist.Concerts = getConcerts(artist.Relations)
	// fmt.Println(artist)

	return artist
}

func getConcerts(url string) []Concert {
	var w http.ResponseWriter
	var data map[string]interface{}

	resp, err := callApi(url)
	if err != nil {

		return nil
	}
	err = json.Unmarshal([]byte(resp), &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	var concerts []Concert

	for i, dateloc := range data["datesLocations"].(map[string]interface{}) {
		// fmt.Println(i, dateloc)
		location := strings.Split(i, "-")

		var dates []string
		for _, date := range dateloc.([]interface{}) {
			dates = append(dates, date.(string))
		}

		concerts = append(concerts, Concert{
			Country: location[1],
			City:    location[0],
			Dates:   dates,
		})

	}

	return concerts
}

func home(w http.ResponseWriter, r *http.Request) {
	// var artist Artist

	// for i := 1; i <= 9; i++ {
	// 	// artist, err = getArtist(i)
	// 	artists = append(artists, getArtist(i))
	// }

	artists := getArtistList()

	sort.Slice(artists, func(i, j int) bool {
		return artists[i].Name < artists[j].Name
	})
	// fmt.Println(artists)

	executeTemplate(w, "./templates/home.html", artists)
}

func all(w http.ResponseWriter, r *http.Request) {
	artists := getArtistList()

	sort.Slice(artists, func(i, j int) bool {
		return artists[i].Name < artists[j].Name
	})
	// fmt.Println(artists)
	executeTemplate(w, "./templates/all.html", artists)
}

func artist(w http.ResponseWriter, r *http.Request, id int) {
	artist := getArtist(id)
	// fmt.Println(artist)

	executeTemplate(w, "./templates/artist.html", artist)
}
